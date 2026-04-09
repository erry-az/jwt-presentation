---
transition: slide-up
---

# Tips Biar JWT Makin Secure (RFC 8725)

<div class="mt-8 space-y-4">

<div v-click class="p-4 rounded-lg bg-gray-800/40 border-l-4 border-blue-400">
  <div class="font-bold text-blue-300">Wajib Pakai Explicit Typing</div>
  <div class="text-gray-400 text-sm">Pakai <code>typ: at+jwt</code> buat cegah serangan tuker-tukar token.</div>
</div>

<div v-click class="p-4 rounded-lg bg-gray-800/40 border-l-4 border-purple-400">
  <div class="font-bold text-purple-300">Whitelist Algoritma</div>
  <div class="text-gray-400 text-sm">Cuma bolehin algoritma aman (RS256 atau ES256). Matiin <code>none</code> sama <code>HS256</code> kalau pakai RSA.</div>
</div>

<div v-click class="p-4 rounded-lg bg-gray-800/40 border-l-4 border-green-400">
  <div class="font-bold text-green-300">Validasi Semua Klaim</div>
  <div class="text-gray-400 text-sm">Wajib cek <code>iss</code>, <code>aud</code>, <code>exp</code>, dan <code>nbf</code> secara teliti.</div>
</div>

<div v-click class="p-4 rounded-lg bg-gray-800/40 border-l-4 border-orange-400">
  <div class="font-bold text-orange-300">Pakai Kunci yang Kuat</div>
  <div class="text-gray-400 text-sm">Minimal 2048-bit buat RSA dan 256-bit buat kunci simetris.</div>
</div>

</div>
