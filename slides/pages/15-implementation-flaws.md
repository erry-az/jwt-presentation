---
transition: slide-up
---

# Implementation Flaws

<div class="mt-8 space-y-8">

<div
  v-motion
  :initial="{ x: 80, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { type: 'spring', stiffness: 100 } }"
>
  <div class="flex items-start gap-4">
    <div class="text-3xl mt-1">💀</div>
    <div>
      <div class="font-bold text-red-300 text-xl mb-1">The <code>alg: none</code> Bypass</div>
      <div class="text-gray-400 mb-3">If a backend <span v-mark.circle.red>blindly trusts</span> the algorithm specified in the header, attackers can send a payload with <code>alg: none</code>, bypass the signature check, and elevate privileges.</div>

```json
{ "alg": "none", "typ": "JWT" }
```
    </div>
  </div>
</div>

<div
  v-motion
  :initial="{ x: 80, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { type: 'spring', stiffness: 100 } }"
>
  <div class="flex items-start gap-4">
    <div class="text-3xl mt-1">🔑</div>
    <div>
      <div class="font-bold text-orange-300 text-xl mb-1">Key Confusion Attack (HMAC vs RSA)</div>
      <div class="text-gray-400">Sending an HMAC-signed token <span v-mark.circle.orange>using a public RSA key</span> can trick insecure backends into validating fake tokens.</div>
    </div>
  </div>
</div>

</div>
