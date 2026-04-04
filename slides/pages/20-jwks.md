---
transition: slide-up
---

# JWKS (JSON Web Key Set)

<div class="grid grid-cols-2 gap-4 mt-2">

<div v-click>
<div class="text-blue-300 font-bold mb-2 flex items-center gap-2 text-sm">
  <carbon:code /> Public Keys (jwks.json)
</div>

```json
{
  "keys": [
    {
      "kty": "RSA",
      "use": "sig",
      "kid": "v1.2",
      "n": "v_Yf_vW4E-Xo7...",
      "e": "AQAB",
      "alg": "RS256"
    }
  ]
}
```
</div>

<div class="space-y-4">

<div v-click class="p-3 rounded-lg border border-blue-400/30 bg-blue-900/10">
  <div class="text-blue-200 font-bold text-xs mb-1">Standard key distribution</div>
  <div class="text-[10px] text-gray-400 italic">Allows resource servers to automatically fetch latest public keys from the Auth Server.</div>
</div>

<div v-click class="p-3 rounded-lg border border-purple-400/30 bg-purple-900/10">
  <div class="text-purple-200 font-bold text-xs mb-1">Key Rotation Support</div>
  <div class="text-[10px] text-gray-400 italic">By having multiple keys in the set, you can rotate keys with zero downtime.</div>
</div>

<div v-click class="flex items-center gap-3 mt-2">
  <carbon:renew class="text-blue-400" />
  <span class="text-xs text-blue-300">Auth servers expose <code>/.well-known/jwks.json</code></span>
</div>

</div>

</div>
