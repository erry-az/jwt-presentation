---
transition: slide-up
---

# IETF Best Current Practices

<div class="text-gray-400 mb-6">RFC 8725: Stop making the same mistakes!</div>

<v-click>

<div
  v-motion
  :initial="{ x: -40, opacity: 0 }"
  :enter="{ x: 0, opacity: 1 }"
  class="mb-4 p-4 rounded-lg border border-green-400/30 bg-green-900/10 flex gap-3 items-start"
>
  <div class="text-xl">✅</div>
  <div>
    <div class="text-green-300 font-bold">Mandatory Algorithm Verification</div>
    <div class="text-gray-400 text-sm">Hardcode expected algorithms in the backend. <span v-mark.underline.green>Never trust the <code>alg</code> header unconditionally.</span></div>
  </div>
</div>

</v-click>

<v-click>

<div
  v-motion
  :initial="{ x: -40, opacity: 0 }"
  :enter="{ x: 0, opacity: 1 }"
  class="mb-4 p-4 rounded-lg border border-blue-400/30 bg-blue-900/10 flex gap-3 items-start"
>
  <div class="text-xl">🏷️</div>
  <div>
    <div class="text-blue-300 font-bold">Explicit Typing</div>
    <div class="text-gray-400 text-sm">Use the <code>typ</code> header (e.g., <code>typ: application/at+jwt</code>) to <span v-mark.underline.blue>prevent token substitution attacks.</span></div>
  </div>
</div>

</v-click>

<v-click>

<div
  v-motion
  :initial="{ x: -40, opacity: 0 }"
  :enter="{ x: 0, opacity: 1 }"
  class="mb-4 p-4 rounded-lg border border-yellow-400/30 bg-yellow-900/10 flex gap-3 items-start"
>
  <div class="text-xl">🔍</div>
  <div>
    <div class="text-yellow-300 font-bold">Strict Claim Validation</div>
    <div class="text-gray-400 text-sm">Always validate <code>iss</code> (issuer), <code>aud</code> (audience), <code>exp</code> (expiration), and <code>nbf</code> (not before).</div>
  </div>
</div>

</v-click>

<div v-click
  v-motion
  :initial="{ scale: 0.95, opacity: 0 }"
  :enter="{ scale: 1, opacity: 1, transition: { delay: 100 } }"
  class="mt-2 p-4 rounded-lg border border-purple-400/30 bg-purple-900/10"
>
  <div class="text-purple-300 font-bold mb-1">📜 RFC 9068</div>
  <div class="text-gray-400 text-sm">Standardized layout for OAuth 2.0 access tokens. Specifies exactly which claims must be present for interoperability.</div>
</div>
