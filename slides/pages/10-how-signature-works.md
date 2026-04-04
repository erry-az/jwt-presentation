---
layout: center
transition: slide-up
---

# How JWT Signature Works

<div class="grid grid-cols-2 gap-8 items-start">

<div v-click
  v-motion
  :initial="{ x: -30, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { duration: 600 } }"
>
<h3 class="text-blue-400 font-bold mb-4 flex items-center gap-2">
  <carbon:edit /> 1. Signing (Creation)
</h3>

<div class="flex flex-col gap-2">
  <div class="p-2 border border-blue-400/30 rounded bg-blue-900/10 text-xs font-mono">
    base64(Header) + "." + base64(Payload)
  </div>
  <div class="text-center text-gray-500">⬇️</div>
  <div class="p-2 border border-purple-400/30 rounded bg-purple-900/10 text-xs font-mono text-center">
    HASH (Algorithm + Private Key)
  </div>
  <div class="text-center text-gray-500">⬇️</div>
  <div class="p-2 border border-green-400/30 rounded bg-green-900/20 text-xs font-mono font-bold text-center">
    SIGNATURE
  </div>
</div>

<div class="mt-4 text-xs text-gray-400 italic">
  "The signature is tied to the exact content of the Header and Payload."
</div>
</div>

<div v-click
  v-motion
  :initial="{ x: 30, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { duration: 600, delay: 200 } }"
>
<h3 class="text-green-400 font-bold mb-4 flex items-center gap-2">
  <carbon:checkmark-outline /> 2. Validation (Verification)
</h3>

<div class="flex flex-col gap-2">
  <div class="p-2 border border-gray-400/30 rounded bg-gray-900/30 text-xs font-mono">
    Split JWT ➔ [H, P, S]
  </div>
  <div class="text-center text-gray-500">⬇️</div>
  <div class="p-2 border border-blue-400/30 rounded bg-blue-900/10 text-xs font-mono">
    Re-calculate Sig using H, P + Public Key
  </div>
  <div class="text-center text-gray-500">⬇️</div>
  <div class="p-2 border border-green-400/30 rounded bg-green-900/20 text-xs font-mono font-bold text-center">
    Compare Recieved Sig == Calculated Sig?
  </div>
</div>

<div class="mt-4 text-xs text-gray-400 italic">
  "If any byte in the Header or Payload changes, the signatures won't match."
</div>
</div>

</div>
