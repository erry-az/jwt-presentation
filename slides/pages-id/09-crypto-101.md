---
layout: default
transition: slide-up
---

# Kriptografi 101: Apa Bedanya?

Sebelum lanjut ke Signature, kita pahamin dulu 3 konsep yang sering bikin bingung:

<div class="grid grid-cols-3 gap-4 mt-8">

<div v-click class="p-4 rounded-xl border border-blue-400/30 bg-blue-900/10 h-full">
<div class="text-blue-400 font-bold mb-3 flex items-center gap-2">
  <carbon:data-class /> Encoding
</div>
<div class="text-xs text-gray-300 space-y-2">
<p>**Tujuan**: Biar data aman dibawa-bawa.</p>
<p>**Proses**: Algoritma bolak-balik (misal: Base64URL).</p>
<div class="p-2 bg-black/30 rounded font-mono text-[10px] text-center">Plaintext ↔️ A-Z/0-9</div>
<p class="text-red-400 font-bold mt-2 pt-2 border-t border-red-400/20">🚫 BUKAN buat keamanan!</p>
<p class="text-[10px] opacity-70 italic font-mono">Siapa aja bisa buka isinya tanpa kunci.</p>
</div>
</div>

<div v-click class="p-4 rounded-xl border border-orange-400/30 bg-orange-900/10 h-full">
<div class="text-orange-400 font-bold mb-3 flex items-center gap-2">
  <carbon:direction-fork /> Hashing
</div>
<div class="text-xs text-gray-300 space-y-2">
<p>**Tujuan**: Ngejaga isi data tetep asli.</p>
<p>**Proses**: Pemetaan satu arah (misal: SHA256).</p>
<div class="p-2 bg-black/30 rounded font-mono text-[10px] text-center">Data ➡️ "Digest" unik</div>
<p class="text-green-400 font-bold mt-2 pt-2 border-t border-green-400/20">🔒 Nggak bisa dibalik.</p>
<p class="text-[10px] opacity-70 italic font-mono">Ini yang dipakai buat bikin Signature JWT.</p>
</div>
</div>

<div v-click class="p-4 rounded-xl border border-purple-400/30 bg-purple-900/10 h-full">
<div class="text-purple-400 font-bold mb-3 flex items-center gap-2">
  <carbon:locked /> Enkripsi
</div>
<div class="text-xs text-gray-300 space-y-2">
<p>**Tujuan**: Kerahasiaan data (Privasi).</p>
<p>**Proses**: Butuh kunci buat buka (misal: AES).</p>
<div class="p-2 bg-black/30 rounded font-mono text-[10px] text-center">Data + Kunci ↔️ Ciphertext</div>
<p class="text-blue-400 font-bold mt-2 pt-2 border-t border-blue-400/20">🔐 Aman & Tersembunyi.</p>
<p class="text-[10px] opacity-70 italic font-mono">Dipakai di JWE (JSON Web Encryption).</p>
</div>
</div>

</div>

<div v-click class="mt-8 p-3 rounded bg-gray-800/50 border border-gray-600/30 text-center text-sm italic text-gray-400">
  "Inget: JWT itu di-<b>Encode</b> dan di-<b>Hash</b>, tapi jarang banget di-<b>Enkripsi</b>."
</div>
