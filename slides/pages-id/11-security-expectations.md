---
transition: slide-up
---

# Harapan Kita Soal Keamanan

<div class="mt-8 space-y-6">

<div
  v-motion
  :initial="{ x: -50, opacity: 0 }"
  :enter="{ x: 0, opacity: 1 }"
  class="flex items-center gap-4 p-4 rounded-xl border border-yellow-400/30 bg-yellow-900/10"
>
  <div class="text-3xl">⚠️</div>
  <div>
    <div class="font-bold text-yellow-300">Nggak Ada Kerahasiaan (JWS)</div>
    <div class="text-gray-400 text-sm">JWT standar (JWS) cuma di-**base64url encoded**, bukan dienkripsi. Buat data sensitif, pakai **JWE (JSON Web Encryption)**.</div>
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
    <div class="font-bold text-green-300">Data Dijamin Nggak Berubah</div>
    <div class="text-gray-400 text-sm"><span v-mark.underline.green>Signature</span> mastiin kalau tokennya nggak dimanipulasi orang lain.</div>
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
    <div class="font-bold text-blue-300">Stateless & Enteng</div>
    <div class="text-gray-400 text-sm">Gampang dipakai di <span v-mark.underline.blue>berbagai mikroservis</span> tanpa perlu bagi-bagi status sesi.</div>
  </div>
</div>

</div>
