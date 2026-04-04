# JWT Best Practices and DPoP Presentation

[![Slidev](https://img.shields.io/badge/Presentation-Slidev-blue.svg)](https://sli.dev/)
[![Go](https://img.shields.io/badge/Example-Golang-00ADD8.svg)](https://go.dev/)
[![Postman](https://img.shields.io/badge/Tests-Postman-FF6C37.svg)](https://www.postman.com/)

An end-to-end technical workshop exploring the evolution of **JSON Web Token (JWT)** security from the initial RFC 7519 standard to modern state-of-the-art **DPoP (RFC 9449)** Proof-of-Possession.

---

## 🚀 Presentation Features

The presentation is built with **Slidev**, offering a high-fidelity, interactive slide experience.

- **Modular Structure**: 25 individual markdown slides in `slides/pages/` for easy management.
- **Deep Technical Content**:
  - **Part 1**: JWT Anatomy & The Standard (RFC 7519).
  - **Part 2**: Common Vulnerabilities (`alg: none`, Key Confusion).
  - **Part 3**: Official Fixes & Best Practices (RFC 8725 & RFC 9068).
  - **Part 4**: DPoP Evolution (RFC 9449) — Solving Bearer token theft.

### To Run the Slides locally:
```bash
cd slides
bun install # or npm install
bun run dev
```
*(Access the presentation at `http://localhost:3031`)*

---

## 🛠️ Lab Example (Go)

The `/example` directory contains a robust Golang implementation designed for hands-on experimentation.

### **Features:**
- **Auth Server**: Issues standard, strict, and DPoP-bound tokens.
- **Resource API**: Demonstrates both vulnerable and secure enforcement patterns.
- **DPoP Validator**: Custom implementation of RFC 9449 proof check including replay protection.
- **CLI Client**: A full-featured client generating RSA keypairs and DPoP proofs for API interaction.

### **Getting Started with Lab:**
```bash
cd example
# 1. Generate RSA keys for the server
mkdir -p keys
openssl genrsa -out keys/private.pem 2048
openssl rsa -in keys/private.pem -pubout -out keys/public.pem

# 2. Run the Server
go run cmd/server/main.go

# 3. Run the CLI Client (Automated flow)
go run cmd/client/main.go
```

---

## 📮 API Testing (Postman)

A comprehensive Postman collection is provided at `example/JWT-Security-Workshop.postman_collection.json`.

- **Part 1-4 Folders**: Tests every endpoint mentioned in the slides.
- **Vulnerability Tests**: Automated steps to exploit `alg: none` and view key confusion results.
- **Environment Management**: Automatically captures and injects `access_token` into subsequent requests.
- **Strict BCP Check**: Verifies RFC 8725 adherence (`typ`, `aud`, `iss` checks).

---

## 📚 Key RFC References

| RFC | Title | Focus |
|---|---|---|
| [**7519**](https://datatracker.ietf.org/doc/html/rfc7519) | JSON Web Token (JWT) | The Core Standard |
| [**8725**](https://datatracker.ietf.org/doc/html/rfc8725) | JWT Best Current Practices | Vulnerability Prevention |
| [**9068**](https://datatracker.ietf.org/doc/html/rfc9068) | JWT Profile for OAuth 2.0 Access Tokens | Standard Interoperability |
| [**9449**](https://datatracker.ietf.org/doc/html/rfc9449) | OAuth 2.0 DPoP | Proof-of-Possession binding |

---

## ⚖️ License
This repository is provided for educational purposes. Feel free to use the slides and code as a template for your own security workshops!
