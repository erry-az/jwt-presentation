.PHONY: keys clean

# Generate Multiple RSA Keys for rotation demo
keys:
	@mkdir -p keys
	@ssh-keygen -t rsa -b 4096 -m PEM -f keys/private_v1.pem -N ""
	@ssh-keygen -t rsa -b 4096 -m PEM -f keys/private_v2.pem -N ""
	@ssh-keygen -t rsa -b 4096 -m PEM -f keys/private_v3.pem -N ""
	@ssh-keygen -f keys/private_v1.pem -e -m PKCS8 > keys/public_v1.pem
	@ssh-keygen -f keys/private_v2.pem -e -m PKCS8 > keys/public_v2.pem
	@ssh-keygen -f keys/private_v3.pem -e -m PKCS8 > keys/public_v3.pem
	@cp keys/private_v1.pem keys/private.pem
	@cp keys/public_v1.pem keys/public.pem
	@echo "Multiple RSA keys (v1, v2, v3) generated in ./keys/"

# Remove the keys directory
clean:
	rm -rf keys
	@echo "Keys removed."

# Run the server
run-server:
	go run cmd/server/main.go

# Run the client (Part 1 Standard)
run-client:
	go run cmd/client/main.go --auth-type standard --api-endpoint /api/data/classic

# Run the 'alg: none' attack demo
attack-alg-none:
	go run cmd/client/main.go --auth-type vuln/alg-none --api-endpoint /api/data/vuln/alg-none

# Run the 'Key Confusion' attack demo
attack-key-confusion:
	go run cmd/client/main.go --attack key-confusion

# Run the Part 3 Secure Flow (RFC 8725)
run-secure:
	go run cmd/client/main.go --auth-type secure --api-endpoint /api/data/strict

# Show the JWKS endpoint
show-jwks:
	curl -s http://localhost:8080/.well-known/jwks.json | jq .

# Run the Part 4 DPoP Flow (RFC 9449)
run-dpop:
	go run cmd/client/main.go --auth-type dpop

# Run the DPoP Stolen Token simulation (Should FAIL with 403)
attack-dpop-stolen:
	go run cmd/client/main.go --attack dpop-stolen
