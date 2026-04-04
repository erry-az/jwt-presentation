---
transition: slide-up
---

# Safe vs Unsafe JWT

<div class="grid grid-cols-2 gap-8 mt-4">

<div v-click class="p-4 rounded-xl border border-red-500/30 bg-red-900/10">
<div class="text-red-400 font-bold mb-2 flex items-center gap-2">
  <carbon:warning /> Unsafe (Vulnerable)
</div>

```json
// Header
{ "alg": "none", "typ": "JWT" }

// Payload (Tampered)
{ "sub": "user123", "role": "admin" }
```
<div class="text-xs text-red-300/70 mt-2">
  ✖️ No signature verification<br>
  ✖️ Trusts Header algorithm blindly<br>
  ✖️ Easy to impersonate any user
</div>
</div>

<div v-click class="p-4 rounded-xl border border-green-500/30 bg-green-900/10">
<div class="text-green-400 font-bold mb-2 flex items-center gap-2">
  <carbon:checkmark-outline /> Safe (Best Practice)
</div>
```json
// Header
{ "alg": "RS256", "typ": "JWT", "kid": "v1.2" }

// Payload
{ "sub": "user123", "role": "user", "exp": 1712150400 }
```
<div class="text-xs text-green-300/70 mt-2">
  ✔️ Cryptographic signature check<br>
  ✔️ Short expiration time<br>
  ✔️ Explicit algorithm enforcement
</div>
</div>

</div>
