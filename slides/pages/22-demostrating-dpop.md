---
transition: slide-up
---

# Demonstrating Proof-of-Possession

<div class="text-gray-400 mb-2">DPoP solves the fundamental flaw of Bearer tokens — it <span v-mark.underline.purple>binds an access token to a specific client</span>.</div>

<div class="mt-4 space-y-3">

<div v-click
  v-motion
  :initial="{ x: -50, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { duration: 500 } }"
  class="flex gap-3 items-start p-3 rounded-lg border border-gray-600/30 bg-gray-800/30"
>
  <div class="text-purple-400 font-bold text-lg w-7 shrink-0">1.</div>
  <div>
    <div class="font-bold text-purple-300">Key Generation</div>
    <div class="text-gray-400 text-sm">Mobile app generates a key pair (e.g., in a Secure Enclave).</div>
  </div>
</div>

<div v-click
  v-motion
  :initial="{ x: -50, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { duration: 500 } }"
  class="flex gap-3 items-start p-3 rounded-lg border border-gray-600/30 bg-gray-800/30"
>
  <div class="text-purple-400 font-bold text-lg w-7 shrink-0">2.</div>
  <div>
    <div class="font-bold text-purple-300">Token Endpoint</div>
    <div class="text-gray-400 text-sm">App requests a token, sending a DPoP proof (signed with private key) containing the public key.</div>
  </div>
</div>

<div v-click
  v-motion
  :initial="{ x: -50, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { duration: 500 } }"
  class="flex gap-3 items-start p-3 rounded-lg border border-gray-600/30 bg-gray-800/30"
>
  <div class="text-purple-400 font-bold text-lg w-7 shrink-0">3.</div>
  <div>
    <div class="font-bold text-purple-300">Key Binding</div>
    <div class="text-gray-400 text-sm">The Authorization Server issues an access token bound to the public key thumbprint (<code>jkt</code> claim).</div>
  </div>
</div>

<div
  v-motion
  :initial="{ x: -50, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { duration: 500 } }"
  class="flex gap-3 items-start p-3 rounded-lg border border-gray-600/30 bg-gray-800/30"
>
  <div class="text-purple-400 font-bold text-lg w-7 shrink-0">4.</div>
  <div>
    <div class="font-bold text-purple-300">API Requests</div>
    <div class="text-gray-400 text-sm">Every request to the API must include a <span v-mark.underline.purple>fresh DPoP proof</span>, proving possession of the private key.</div>
  </div>
</div>

</div>
