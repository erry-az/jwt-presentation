# JWT Security Lab (Go Example)

This directory contains a full end-to-end laboratory implementation for testing JWT security across the evolutionary stages covered in the presentation.

## 🚀 Key Evolutionary Stages

- **Stage 1 (Standard)**: Basic JWT flow using RFC 7519 (RS256).
- **Stage 2 (Vulnerabilities)**: Demonstrates `alg: none` bypasses and Key Confusion mistakes.
- **Stage 3 (Official Fixes)**: Implementing Best Current Practices (BCP) from RFC 8725 and RFC 9068.
- **Stage 4 (DPoP)**: Full Proof-of-Possession binding (RFC 9449) for preventing token theft/replay.

---

## 🛠️ Components

1. **`cmd/server/`**: A unified Auth Server & Resource API with multiple endpoints for each stage.
2. **`cmd/client/`**: A full-featured CLI client with automatic RSA key generation and DPoP proof signing.
3. **`internal/`**:
   - `auth/`: JWT provider for issuing various token types.
   - `dpop/`: Cryptographic validator for RFC 9449 proofs.
   - `models/`: Shared constants, paths, and claim structures.
4. **`JWT-Security-Workshop.postman_collection.json`**: For visual testing and exploit demonstration.

---

## ⚙️ Setup

### **1. Key Generation**
The server requires a 2048-bit RSA key pair to sign JWS tokens.
```bash
mkdir -p keys
openssl genrsa -out keys/private.pem 2048
openssl rsa -in keys/private.pem -pubout -out keys/public.pem
```

### **2. Dependencies**
```bash
go mod tidy
```

---

## 🚦 How to Run & Test

### **A. Unified Run (CLI Flow)**
Start the server in one terminal:
```bash
go run cmd/server/main.go
```
Run the automated laboratory flow in another terminal:
```bash
go run cmd/client/main.go
```
This client will automatically:
- Register a local key for DPoP.
- Authenticate and get a standard token.
- Authenticate and get a DPoP-bound token.
- Access the classic, secure, and DPoP-enforced endpoints.

### **B. Postman Testing (Manual & Exploit Flow)**
1. Import `JWT-Security-Workshop.postman_collection.json` into Postman.
2. The collection is organized into Parts 1-4.
3. Use the **Part 2** folder to witness how an `alg: none` token can bypass security on vulnerable endpoints.
4. Use the **Part 3** folder to see how strict claim validation (aud, iss, typ) blocks invalid tokens.

---

## 🔍 Laboratory Checklist

| Objective | Endpoint | Test Mechanism |
|---|---|---|
| **Test `alg: none`** | `/api/data/vuln/alg-none` | Postman (Exploit folder) |
| **Test Key Confusion** | `/api/data/vuln/key-confusion` | Postman (Key Confusion folder) |
| **Verify `typ` header** | `/api/data/strict` | Postman (Official Fixes folder) |
| **Test Token Theft** | `/api/data/dpop` | CLI Client (Automated binding) |

---

## 🔒 Security Disclaimer
This code contains intentional vulnerabilities for educational purposes (marked with `VULNERABLE` comments). **Do not use the vulnerable handler logic in production.**
