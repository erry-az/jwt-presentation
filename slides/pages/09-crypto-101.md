---
layout: default
transition: slide-up
---

# The Fundamentals: Crypto 101

Before deep-diving into signatures, let's clear up three concepts often confused in the JWT world:

<div class="grid grid-cols-3 gap-4 mt-8">

<div v-click class="p-4 rounded-xl border border-blue-400/30 bg-blue-900/10 h-full">
<div class="text-blue-400 font-bold mb-3 flex items-center gap-2">
  <carbon:data-class /> Encoding
</div>
<div class="text-xs text-gray-300 space-y-2">
<p>**Goal**: Data portability & transport.</p>
<p>**Process**: Reversible algorithm (e.g., Base64URL).</p>
<div class="p-2 bg-black/30 rounded font-mono text-[10px] text-center">Plaintext ↔️ A-Z/0-9</div>
<p class="text-red-400 font-bold mt-2 pt-2 border-t border-red-400/20">🚫 NOT for security.</p>
<p class="text-[10px] opacity-70 italic font-mono">Anyone can "decode" it without a key.</p>
</div>
</div>

<div v-click class="p-4 rounded-xl border border-orange-400/30 bg-orange-900/10 h-full">
<div class="text-orange-400 font-bold mb-3 flex items-center gap-2">
  <carbon:direction-fork /> Hashing
</div>
<div class="text-xs text-gray-300 space-y-2">
<p>**Goal**: Integrity & identity.</p>
<p>**Process**: One-way mapping (e.g., SHA256).</p>
<div class="p-2 bg-black/30 rounded font-mono text-[10px] text-center">Data ➡️ Fixed-length "Digest"</div>
<p class="text-green-400 font-bold mt-2 pt-2 border-t border-green-400/20">🔒 Irreversible.</p>
<p class="text-[10px] opacity-70 italic font-mono">Used to create the JWT Signature.</p>
</div>
</div>

<div v-click class="p-4 rounded-xl border border-purple-400/30 bg-purple-900/10 h-full">
<div class="text-purple-400 font-bold mb-3 flex items-center gap-2">
  <carbon:locked /> Encryption
</div>
<div class="text-xs text-gray-300 space-y-2">
<p>**Goal**: Confidentiality & privacy.</p>
<p>**Process**: Reversible ONLY with a key (e.g., AES).</p>
<div class="p-2 bg-black/30 rounded font-mono text-[10px] text-center">Data + Key ↔️ Ciphertext</div>
<p class="text-blue-400 font-bold mt-2 pt-2 border-t border-blue-400/20">🔐 Reversible WITH Key.</p>
<p class="text-[10px] opacity-70 italic font-mono">Used in JWE (JSON Web Encryption).</p>
</div>
</div>

</div>

<div v-click class="mt-8 p-3 rounded bg-gray-800/50 border border-gray-600/30 text-center text-sm italic text-gray-400">
  "JWTs are <b>Encoded</b> and <b>Hashed</b> (signed), but rarely <b>Encrypted</b>."
</div>
