package auth

import (
	"crypto/rsa"
	"encoding/json"
	"example/internal/models"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/pkg/errors"
)

type JWTProvider struct {
	privateKey *rsa.PrivateKey
	publicKeys map[string]*rsa.PublicKey
}

func NewJWTProvider(privateKeyPath, publicKeyPath string) (*JWTProvider, error) {
	// Default active private key (v1)
	privateKeyData, err := os.ReadFile("./keys/private_v1.pem")
	if err != nil {
		// Fallback to legacy path if v1 not found
		privateKeyData, err = os.ReadFile(privateKeyPath)
	}
	if err != nil {
		return nil, errors.Wrap(err, "read private key")
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		return nil, errors.Wrap(err, "parse private key")
	}

	pubKeys := make(map[string]*rsa.PublicKey)
	// Try loading v1, v2, v3
	for i := 1; i <= 3; i++ {
		path := fmt.Sprintf("./keys/public_v%d.pem", i)
		pubKeyData, err := os.ReadFile(path)
		if err != nil {
			continue // Skip if missing
		}
		pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyData)
		if err == nil {
			pubKeys[fmt.Sprintf("v%d", i)] = pubKey
		}
	}

	// Also load the default public key for backward compatibility
	if _, ok := pubKeys["v1"]; !ok {
		publicKeyData, _ := os.ReadFile(publicKeyPath)
		if publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyData); err == nil {
			pubKeys["v1"] = publicKey
		}
	}

	return &JWTProvider{
		privateKey: privateKey,
		publicKeys: pubKeys,
	}, nil
}

func (p *JWTProvider) IssueToken(userID, userName string, isAdmin bool) (string, error) {
	claims := models.CustomClaims{
		Name:  userName,
		Admin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			Issuer:    models.Issuer,
			Audience:  []string{models.Audience},
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(p.privateKey)
	if err != nil {
		return "", errors.Wrap(err, "sign token")
	}

	return tokenString, nil
}

// IssueTokenStrict implements RFC 8725 & RFC 9068 recommendations
func (p *JWTProvider) IssueTokenStrict(userID, userName string, isAdmin bool) (string, error) {
	claims := models.CustomClaims{
		Name:  userName,
		Admin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        "unique-jti-" + userID,
			Subject:   userID,
			Issuer:    "jwt-presentation-auth-server",
			Audience:  []string{"jwt-presentation-resource-api"},
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)), // Shorter expiry
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	// Mandatory Type for RFC 8725 Section 3.11
	token.Header["typ"] = models.TypeAccessToken
	token.Header["kid"] = models.KID_V1 // Key ID for discovery in JWKS

	tokenString, err := token.SignedString(p.privateKey)
	if err != nil {
		return "", errors.Wrap(err, "sign token strict")
	}

	return tokenString, nil
}

// JWKS returns the public key set in JSON format
func (p *JWTProvider) JWKS() ([]byte, error) {
	set := jwk.NewSet()

	for kid, pubKey := range p.publicKeys {
		key, err := jwk.FromRaw(pubKey)
		if err != nil {
			return nil, errors.Wrap(err, "create jwk from raw rsa")
		}

		_ = key.Set(jwk.KeyIDKey, kid)
		_ = key.Set(jwk.AlgorithmKey, "RS256")
		_ = key.Set(jwk.KeyUsageKey, "sig")

		_ = set.AddKey(key)
	}

	buf, err := json.MarshalIndent(set, "", "  ")
	if err != nil {
		return nil, errors.Wrap(err, "marshal jwks")
	}

	return buf, nil
}

// IssueTokenDPoP implements RFC 9449 key binding
func (p *JWTProvider) IssueTokenDPoP(userID, userName string, jkt string) (string, error) {
	claims := models.CustomClaims{
		Name: userName,
		Confirmation: &models.Confirmation{
			JKT: jkt,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			Issuer:    "jwt-presentation-auth-server",
			Audience:  []string{"jwt-presentation-resource-api"},
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["typ"] = models.TypeAccessToken // Recommended for transparency
	tokenString, err := token.SignedString(p.privateKey)
	if err != nil {
		return "", errors.Wrap(err, "sign token dpop")
	}

	return tokenString, nil
}

// IssueTokenUnsigned demonstrates the "alg: none" bypass by returning an unsigned token.
func (p *JWTProvider) IssueTokenUnsigned(userID, userName string, isAdmin bool) (string, error) {
	claims := models.CustomClaims{
		Name:  userName,
		Admin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			Issuer:    "jwt-presentation-auth-server",
			Audience:  []string{"jwt-presentation-resource-api"},
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodNone, claims)
	// SigningMethodNone requires the secret to be 'jwt.UnsafeAllowNoneSignatureType' to actually output it.
	tokenString, err := token.SignedString(jwt.UnsafeAllowNoneSignatureType)
	if err != nil {
		return "", errors.Wrap(err, "create unsigned token")
	}

	return tokenString, nil
}

func (p *JWTProvider) PublicKey() *rsa.PublicKey {
	return p.publicKeys["v1"]
}

// PublicKeyPEM returns the public key in PEM format for simulation purposes.
func (p *JWTProvider) PublicKeyPEM() []byte {
	// Simple helper to get the public key for Key Confusion demo
	publicKeyPEM, _ := os.ReadFile("./keys/public.pem")
	return publicKeyPEM
}
