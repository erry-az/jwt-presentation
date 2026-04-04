package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"example/internal/auth"
	"example/internal/dpop"
	"example/internal/models"

	"github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	// 1. Setup JWT Provider
	provider, err := auth.NewJWTProvider("./keys/private.pem", "./keys/public.pem")
	if err != nil {
		log.Fatalf("failed to initialize JWT provider: %v", err)
	}

	// 2. Setup JWT Middleware (Part 1 Standard)
	jwtValidator, err := validator.New(
		func(ctx context.Context) (interface{}, error) {
			return provider.PublicKey(), nil
		},
		validator.RS256,
		models.Issuer,
		[]string{models.Audience},
		validator.WithCustomClaims(func() validator.CustomClaims {
			return &models.CustomClaims{}
		}),
	)
	if err != nil {
		log.Fatalf("failed to set up the validator: %v", err)
	}

	// Custom extractor to support both Bearer and DPoP schemes (RFC 9449 Section 7)
	extractor := func(r *http.Request) (string, error) {
		authHeader := r.Header.Get(models.HeaderAuthorization)
		if authHeader == "" {
			return "", nil
		}
		parts := strings.Fields(authHeader)
		if len(parts) != 2 {
			return "", nil
		}
		scheme := strings.ToLower(parts[0])
		if scheme != "bearer" && scheme != strings.ToLower(models.HeaderDPoP) {
			return "", nil
		}
		return parts[1], nil
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithTokenExtractor(extractor),
	)

	// 3. Define Handlers
	mux := http.NewServeMux()

	// --- Part 1: Standard Handlers ---
	mux.HandleFunc("POST "+models.PathAuthStandard, func(w http.ResponseWriter, r *http.Request) {
		token, err := provider.IssueToken("user_123", "John Doe", false)
		if err != nil {
			http.Error(w, "could not issue token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"access_token": token})
	})

	mux.Handle("GET "+models.PathDataClassic, middleware.CheckJWT(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		customClaims := claims.CustomClaims.(*models.CustomClaims)
		log.Printf("authorized access by %s (admin: %v)", customClaims.Name, customClaims.Admin)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Welcome to the classic protected resource!",
			"user":    customClaims.Name,
			"admin":   customClaims.Admin,
		})
	})))

	// --- Part 2: Vulnerabilities Demo Handlers ---

	// Issue the 'alg: none' token
	mux.HandleFunc("POST "+models.PathAuthVulnNone, func(w http.ResponseWriter, r *http.Request) {
		token, _ := provider.IssueTokenUnsigned("attacker_111", "Hacker Joe", true)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"access_token": token})
	})

	// Public endpoint giving the key, common for RSA but dangerous if mishandled in HMAC code
	mux.HandleFunc("GET "+models.PathPublicKey, func(w http.ResponseWriter, r *http.Request) {
		w.Write(provider.PublicKeyPEM())
	})

	// Vulnerable endpoint for 'alg: none' bypass
	mux.HandleFunc("GET "+models.PathDataVulnNone, func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get(models.HeaderAuthorization)
		if len(tokenString) > 7 {
			tokenString = tokenString[7:] // Remove "Bearer "
		}

		// VULNERABLE: Uses a Parser that allows "none" algorithm
		parser := jwt.NewParser(jwt.WithValidMethods([]string{"none", "RS256"}))
		token, err := parser.ParseWithClaims(tokenString, &models.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
			if t.Method.Alg() == "none" {
				return jwt.UnsafeAllowNoneSignatureType, nil
			}
			return provider.PublicKey(), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(*models.CustomClaims)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "REACHED VULNERABLE ENDPOINT via alg: none! 💀",
			"user":    claims.Name,
			"admin":   claims.Admin,
		})
	})

	// Vulnerable endpoint for Key Confusion (HMAC vs RSA)
	mux.HandleFunc("GET "+models.PathDataVulnConfuse, func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get(models.HeaderAuthorization)
		if len(tokenString) > 7 {
			tokenString = tokenString[7:]
		}

		// VULNERABLE: Incorrectly uses the Public Key (bytes) for HMAC validation if the client sends HS256
		token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
			// A common MISTAKE: Not checking the actual alg before returning the key
			// The developer thinks: "I'm using RSA, so here is my public key"
			// But if t.Method is HMAC (HS256), the library uses these bytes as the HMAC secret!
			return provider.PublicKeyPEM(), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(*models.CustomClaims)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "REACHED VULNERABLE ENDPOINT via Key Confusion! 🔑",
			"user":    claims.Name,
			"admin":   claims.Admin,
		})
	})

	// --- Part 3: Official Fixes (RFC 8725 & RFC 9068) ---

	// 1. Setup Strict Validator (RFC 8725)
	strictValidator, err := validator.New(
		func(ctx context.Context) (interface{}, error) {
			return provider.PublicKey(), nil
		},
		validator.RS256, // Rule 1: Strict algorithm whitelist
		models.Issuer,
		[]string{models.Audience},
		validator.WithCustomClaims(func() validator.CustomClaims {
			return &models.CustomClaims{}
		}),
	)
	if err != nil {
		log.Fatalf("failed to set up strict validator: %v", err)
	}

	strictMiddleware := jwtmiddleware.New(func(ctx context.Context, tokenString string) (interface{}, error) {
		// Rule 2: Explicit Typing check (RFC 8725 Section 3.11)
		// We can check it manually before passing to the validator
		parser := jwt.NewParser()
		token, _, err := parser.ParseUnverified(tokenString, jwt.MapClaims{})
		if err == nil {
			if typ, ok := token.Header["typ"].(string); !ok || typ != models.TypeAccessToken {
				return nil, log.Output(2, "invalid typ header: expected "+models.TypeAccessToken)
			}
		}

		return strictValidator.ValidateToken(ctx, tokenString)
	})

	// JWKS Endpoint
	mux.HandleFunc("GET "+models.PathJWKS, func(w http.ResponseWriter, r *http.Request) {
		jwks, _ := provider.JWKS()
		w.Header().Set("Content-Type", "application/json")
		w.Write(jwks)
	})

	// Auth Server: Issue Strict Token
	mux.HandleFunc("POST "+models.PathAuthSecure, func(w http.ResponseWriter, r *http.Request) {
		token, err := provider.IssueTokenStrict("user_456", "Jane BCP", false)
		if err != nil {
			http.Error(w, "could not issue strict token", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"access_token": token})
	})

	// Resource API: Strict Enforcement
	mux.Handle("GET "+models.PathDataStrict, strictMiddleware.CheckJWT(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		customClaims := claims.CustomClaims.(*models.CustomClaims)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "REACHED SECURE ENDPOINT via RFC 8725 standards! ✅",
			"user":    customClaims.Name,
		})
	})))

	// --- Part 4: DPoP Evolution (RFC 9449) ---

	dpopValidator := dpop.NewValidator()

	// Auth Server: Issue DPoP bound token
	mux.HandleFunc("POST "+models.PathAuthDPoP, func(w http.ResponseWriter, r *http.Request) {
		proof := r.Header.Get(models.HeaderDPoP)
		jkt, err := dpopValidator.ValidateProof(r, proof)
		if err != nil {
			http.Error(w, "invalid DPoP proof: "+err.Error(), http.StatusBadRequest)
			return
		}

		token, err := provider.IssueTokenDPoP("user_789", "DPoP User", jkt)
		if err != nil {
			http.Error(w, "could not issue dpop token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"access_token": token})
	})

	// Resource API: DPoP Enforcement
	mux.Handle("GET "+models.PathDataDPoP, middleware.CheckJWT(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Regular JWT check is already done by middleware (RS256 signature etc)
		claims := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		customClaims := claims.CustomClaims.(*models.CustomClaims)

		// 2. CRYPTOGRAPHIC BINDING: Verify DPoP Proof
		proof := r.Header.Get(models.HeaderDPoP)
		jkt, err := dpopValidator.ValidateProof(r, proof)
		if err != nil {
			http.Error(w, "DPoP proof invalid: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// 3. MATCH: Does the token cnf.jkt match the current proof's public key?
		if customClaims.Confirmation == nil || customClaims.Confirmation.JKT != jkt {
			http.Error(w, "DPoP key binding mismatch! Stolen token detected? 🚨", http.StatusForbidden)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "REACHED SECURE ENDPOINT via DPoP Proof-of-Possession! 🚀",
			"user":    customClaims.Name,
			"jkt":     jkt,
		})
	})))

	log.Println("Server starting on :8080 (Vulnerabilities, Fixes & DPoP Enabled)...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
