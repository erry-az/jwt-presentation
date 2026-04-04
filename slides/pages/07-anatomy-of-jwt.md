---
transition: fade-out
---

# Anatomy of a JWT

A standard **JSON Web Token (JWT)** consists of three parts separated by dots `(.)`

<v-click>

<div
  v-motion
  :initial="{ x: -40, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { duration: 500 } }"
  class="mt-6 p-4 rounded-lg border border-red-400/40 bg-red-900/10 flex items-start gap-3"
>
  <div class="text-2xl">📄</div>
  <div>
    <div class="text-red-400 font-bold">Header</div>
    <div class="text-gray-300 text-sm">Contains metadata like token type and signing algorithm (e.g., <code>HS256</code>, <code>RS256</code>).</div>
  </div>
</div>

</v-click>

<v-click>

<div
  v-motion
  :initial="{ x: -40, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { duration: 500 } }"
  class="mt-3 p-4 rounded-lg border border-purple-400/40 bg-purple-900/10 flex items-start gap-3"
>
  <div class="text-2xl">📦</div>
  <div>
    <div class="text-purple-400 font-bold">Payload</div>
    <div class="text-gray-300 text-sm">Contains the claims — statements about an entity and additional data.</div>
  </div>
</div>

</v-click>

<v-click>

<div
  v-motion
  :initial="{ x: -40, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { duration: 500 } }"
  class="mt-3 p-4 rounded-lg border border-blue-400/40 bg-blue-900/10 flex items-start gap-3"
>
  <div class="text-2xl">🔏</div>
  <div>
    <div class="text-blue-400 font-bold">Signature</div>
    <div class="text-gray-300 text-sm">Used to verify the token hasn't been altered in transit.</div>
  </div>
</div>

</v-click>

<div v-click class="mt-8 p-4 bg-gray-800/80 rounded-xl shadow-lg font-mono text-center text-xl tracking-wider border border-gray-600/40">
  <span class="text-red-400 font-bold">HEADER</span><span class="text-gray-500">.</span><span class="text-purple-400 font-bold">PAYLOAD</span><span class="text-gray-500">.</span><span class="text-blue-400 font-bold">SIGNATURE</span>
</div>
