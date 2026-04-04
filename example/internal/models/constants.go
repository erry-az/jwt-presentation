package models

const (
	// JWT Configuration
	Issuer   = "jwt-presentation-auth-server"
	Audience = "jwt-presentation-resource-api"

	// RFC 8725 & 9068 Headers
	TypeAccessToken = "application/at+jwt"
	TypeDPoPProof   = "dpop+jwt"
	
	// Key IDs
	KID_V1 = "v1.2"

	// HTTP Headers
	HeaderAuthorization = "Authorization"
	HeaderDPoP          = "DPoP"
)

const (
	// Auth Endpoints
	PathAuthStandard = "/auth/standard"
	PathAuthSecure   = "/auth/secure"
	PathAuthDPoP     = "/auth/dpop"
	PathAuthVulnNone = "/auth/vuln/alg-none"

	// Resource Endpoints
	PathDataClassic      = "/api/data/classic"
	PathDataStrict       = "/api/data/strict"
	PathDataDPoP         = "/api/data/dpop"
	PathDataVulnNone     = "/api/data/vuln/alg-none"
	PathDataVulnConfuse  = "/api/data/vuln/key-confusion"
	
	// Metadata Endpoints
	PathJWKS      = "/.well-known/jwks.json"
	PathPublicKey = "/v1/public-key"
)
