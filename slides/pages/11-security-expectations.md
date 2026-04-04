---
transition: slide-up
---

# Security Expectations

<div class="mt-8 space-y-6">

<div
  v-motion
  :initial="{ x: -50, opacity: 0 }"
  :enter="{ x: 0, opacity: 1 }"
  class="flex items-center gap-4 p-4 rounded-xl border border-yellow-400/30 bg-yellow-900/10"
>
  <div class="text-3xl">⚠️</div>
  <div>
    <div class="font-bold text-yellow-300">No Confidentiality (JWS)</div>
    <div class="text-gray-400 text-sm">Standard JWTs (JWS) are <span v-mark.underline.orange>base64url encoded</span>, not encrypted. For sensitive data, use <span v-mark.underline.blue>JWE (JSON Web Encryption)</span>.</div>
  </div>
</div>

<div
  v-motion
  :initial="{ x: -50, opacity: 0 }"
  :enter="{ x: 0, opacity: 1 }"
  class="flex items-center gap-4 p-4 rounded-xl border border-green-400/30 bg-green-900/10"
>
  <div class="text-3xl">🔒</div>
  <div>
    <div class="font-bold text-green-300">Integrity Provided</div>
    <div class="text-gray-400 text-sm">The <span v-mark.underline.green>signature</span> ensures the token hasn't been tampered with.</div>
  </div>
</div>

<div
  v-motion
  :initial="{ x: -50, opacity: 0 }"
  :enter="{ x: 0, opacity: 1 }"
  class="flex items-center gap-4 p-4 rounded-xl border border-blue-400/30 bg-blue-900/10"
>
  <div class="text-3xl">🚀</div>
  <div>
    <div class="font-bold text-blue-300">Stateless &amp; Portable</div>
    <div class="text-gray-400 text-sm">Designed to work seamlessly <span v-mark.underline.blue>across microservices</span> without shared session state.</div>
  </div>
</div>

</div>
