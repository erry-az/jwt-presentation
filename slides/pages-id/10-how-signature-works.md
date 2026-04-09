---
layout: center
transition: slide-up
---

# Rahasia di Balik Signature JWT

<div class="grid grid-cols-2 gap-8 items-start">

<div v-click
  v-motion
  :initial="{ x: -30, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { duration: 600 } }"
>
<h3 class="text-blue-400 font-bold mb-4 flex items-center gap-2">
  <carbon:edit /> 1. Penandatanganan (Signing)
</h3>

<div class="flex flex-col gap-2">
  <div class="p-2 border border-blue-400/30 rounded bg-blue-900/10 text-xs font-mono">
    base64Url(Header) + "." + base64Url(Payload)
  </div>
  <div class="text-center text-gray-500">⬇️</div>
  <div class="p-2 border border-purple-400/30 rounded bg-purple-900/10 text-xs font-mono text-center">
    HASH (Algoritma + Kunci Privat)
  </div>
  <div class="text-center text-gray-500">⬇️</div>
  <div class="p-2 border border-green-400/30 rounded bg-green-900/20 text-xs font-mono font-bold text-center">
    SIGNATURE
  </div>
</div>

<div class="mt-4 text-xs text-gray-400 italic">
  "Signature itu nempel erat sama isi Header dan Payload."
</div>
</div>

<div v-click
  v-motion
  :initial="{ x: 30, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { duration: 600, delay: 200 } }"
>
<h3 class="text-green-400 font-bold mb-4 flex items-center gap-2">
  <carbon:checkmark-outline /> 2. Validasi (Verification)
</h3>

<div class="flex flex-col gap-2">
  <div class="p-2 border border-gray-400/30 rounded bg-gray-900/30 text-xs font-mono">
    Pisahkan JWT ➔ [H, P, S]
  </div>
  <div class="text-center text-gray-500">⬇️</div>
  <div class="p-2 border border-blue-400/30 rounded bg-blue-900/10 text-xs font-mono">
    Hitung ulang Sig dari H, P + Kunci Publik
  </div>
  <div class="text-center text-gray-500">⬇️</div>
  <div class="p-2 border border-green-400/30 rounded bg-green-900/20 text-xs font-mono font-bold text-center">
    Bandingkan Sig yang Datang == Sig Hasil Hitung?
  </div>
</div>

<div class="mt-4 text-xs text-gray-400 italic">
  "Kalau ada satu huruf aja yang berubah, Signature-nya auto-gagal."
</div>
</div>

</div>
