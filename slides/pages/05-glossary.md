---
transition: slide-up
---

# Glossary of Terms

<div class="grid grid-cols-2 gap-x-8 gap-y-4 mt-6">

<div v-click>
  <div class="text-blue-400 font-bold">JWT (JSON Web Token)</div>
  <div class="text-xs text-gray-400">A compact, URL-safe means of representing claims to be transferred between two parties.</div>
</div>

<div v-click>
  <div class="text-purple-400 font-bold">JWS / JWE</div>
  <div class="text-xs text-gray-400">JWS is Signed JWT (Integrity). JWE is Encrypted JWT (Confidentiality).</div>
</div>

<div v-click>
  <div class="text-green-400 font-bold">DPoP</div>
  <div class="text-xs text-gray-400">Demonstrating Proof-of-Possession. Mechanism to bind token with asymmetric key.</div>
</div>

<div v-click>
  <div class="text-orange-400 font-bold">JWKS (Key Set)</div>
  <div class="text-xs text-gray-400">A set of JSON-encoded cryptographic keys used by servers to distribute public keys.</div>
</div>

<div v-click>
  <div class="text-red-400 font-bold">Bearer Token</div>
  <div class="text-xs text-gray-400">A token that allows access to whoever holds it (vulnerable if stolen).</div>
</div>

<div v-click>
  <div class="text-teal-400 font-bold">Proof-of-Possession</div>
  <div class="text-xs text-gray-400">Requirement for the sender to prove they hold the private key associated with the token.</div>
</div>

<div v-click>
  <div class="text-yellow-400 font-bold">XSS / CSRF</div>
  <div class="text-xs text-gray-400">XSS (Script injection to steal tokens). CSRF (Cross-site fraud to use tokens).</div>
</div>

<div v-click>
  <div class="text-pink-400 font-bold">Symmetric / Asymmetric</div>
  <div class="text-xs text-gray-400">Symmetric (Shared secret, e.g. HS256). Asymmetric (Private/Public, e.g. RS256).</div>
</div>

</div>
