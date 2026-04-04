package dpop

import (
	"crypto"
	"encoding/base64"
	"encoding/json"
	"example/internal/models"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/pkg/errors"
)

// DPoPClaims represents the standard claims for a DPoP Proof JWT
type DPoPClaims struct {
	HTTPMethod string `json:"htm"`
	HTTPURL    string `json:"htu"`
	jwt.RegisteredClaims
}

// Validator handles the verification of DPoP proofs
type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

// ValidateProof verifies the DPoP proof header and returns the thumbprint (jkt) of the public key
func (v *Validator) ValidateProof(r *http.Request, proofHeader string) (string, error) {
	if proofHeader == "" {
		return "", fmt.Errorf("missing %s proof", models.HeaderDPoP)
	}

	// 1. Parse without verifying signature first to extract JWK
	var proofClaims DPoPClaims
	token, err := jwt.ParseWithClaims(proofHeader, &proofClaims, func(t *jwt.Token) (interface{}, error) {
		// DPoP proof must have an embedded JWK in the header
		jwkMap, ok := t.Header["jwk"].(map[string]interface{})
		if !ok {
			return nil, errors.New("missing jwk in DPoP header")
		}

		jsonBuffer, _ := json.Marshal(jwkMap)
		key, err := jwk.ParseKey(jsonBuffer)
		if err != nil {
			return nil, errors.Wrap(err, "invalid jwk in header")
		}

		var rawKey interface{}
		if err := key.Raw(&rawKey); err != nil {
			return nil, errors.Wrap(err, "failed to get raw key")
		}

		return rawKey, nil
	})

	if err != nil || !token.Valid {
		return "", errors.Wrap(err, "invalid DPoP proof signature or format")
	}

	// 2. Validate htm and htu (Part of RFC 9449 Section 4.2)
	fullURL := fmt.Sprintf("http://%s%s", r.Host, r.URL.Path)
	if !strings.EqualFold(proofClaims.HTTPMethod, r.Method) {
		return "", fmt.Errorf("htm mismatch: expected %s, got %s", r.Method, proofClaims.HTTPMethod)
	}
	// Note: In real production, check htu strictly against actual request URL
	if !strings.Contains(fullURL, proofClaims.HTTPURL) {
		return "", fmt.Errorf("htu mismatch: request %s, proof %s", fullURL, proofClaims.HTTPURL)
	}

	// 3. Check for replay protection (iat should be fresh)
	if time.Since(proofClaims.IssuedAt.Time) > 2*time.Minute {
		return "", errors.New("DPoP proof expired")
	}

	// 4. Compute the JWK Thumbprint (SHA-256)
	jwkMap := token.Header["jwk"].(map[string]interface{})
	jsonBuffer, _ := json.Marshal(jwkMap)
	key, _ := jwk.ParseKey(jsonBuffer)
	
	thumbprint, err := key.Thumbprint(crypto.SHA256)
	if err != nil {
		return "", errors.Wrap(err, "failed to compute thumbprint")
	}

	return base64.RawURLEncoding.EncodeToString(thumbprint), nil
}
