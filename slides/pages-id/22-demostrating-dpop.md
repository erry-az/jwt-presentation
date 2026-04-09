---
transition: slide-up
---

# Cara Kerja DPoP

<div class="text-gray-400 text-sm mb-6 font-semibold italic">Membuktikan Kepemilikan (RFC 9449)</div>

<div class="mt-4 space-y-3">

<div v-click
  v-motion
  :initial="{ x: -50, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { duration: 500 } }"
  class="flex gap-3 items-start p-3 rounded-lg border border-gray-600/30 bg-gray-800/30"
>
  <div class="text-purple-400 font-bold text-lg w-7 shrink-0">1.</div>
  <div>
    <div class="font-bold text-purple-300">Pembuatan Kunci</div>
    <div class="text-gray-400 text-sm">Klien (misal: Aplikasi Mobile) bikin kunci privat/publik sendiri di lokal.</div>
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
    <div class="font-bold text-purple-300">Endpoint Token</div>
    <div class="text-gray-400 text-sm">Klien minta token sambil kasih liat Bukti DPoP (ditandatangani pakai kunci privat).</div>
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
    <div class="font-bold text-purple-300">Pengikatan Kunci</div>
    <div class="text-gray-400 text-sm">Server Otorisasi kasih token yang udah "dikunci" ke thumbprint kunci publik (klaim <code>jkt</code>).</div>
  </div>
</div>

<div v-click
  v-motion
  :initial="{ x: -50, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { duration: 500 } }"
  class="flex gap-3 items-start p-3 rounded-lg border border-gray-600/30 bg-gray-800/30"
>
  <div class="text-purple-400 font-bold text-lg w-7 shrink-0">4.</div>
  <div>
    <div class="font-bold text-purple-300">Panggil API</div>
    <div class="text-gray-400 text-sm">Setiap panggil API, wajib kasih bukti DPoP baru biar ketauan kalau kita pegang kunci aslinya.</div>
  </div>
</div>

</div>
