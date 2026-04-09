---
transition: slide-up
---

# Mana yang Aman, Mana yang Nggak?

<div class="grid grid-cols-2 gap-8 mt-8">

<div v-click class="p-5 rounded-xl border border-red-400/30 bg-red-900/10 h-full">
<div class="text-red-400 font-bold mb-4 flex items-center gap-2">
  <carbon:warning-alt /> Gak Aman (Bahaya)
</div>
<ul class="text-sm text-gray-300 space-y-3">
  <li class="flex items-center gap-2"><carbon:close-filled class="text-red-500 shrink-0" /> <code>alg: none</code> dibolehin</li>
  <li class="flex items-center gap-2"><carbon:close-filled class="text-red-500 shrink-0" /> Kunci pendek (misal: 1024-bit RSA)</li>
  <li class="flex items-center gap-2"><carbon:close-filled class="text-red-500 shrink-0" /> Gak cek expired (kedaluwarsa)</li>
  <li class="flex items-center gap-2"><carbon:close-filled class="text-red-500 shrink-0" /> Gak cek siapa penerbitnya (<code>iss</code>)</li>
</ul>
</div>

<div v-click class="p-5 rounded-xl border border-green-400/30 bg-green-900/10 h-full shadow-[0_0_20px_-5px_rgba(34,197,94,0.3)]">
<div class="text-green-400 font-bold mb-4 flex items-center gap-2">
  <carbon:shield-check /> Aman (Best Practice)
</div>
<ul class="text-sm text-gray-300 space-y-3">
  <li class="flex items-center gap-2"><carbon:checkmark-filled class="text-green-500 shrink-0" /> Algoritma dikunci secara eksplisit</li>
  <li class="flex items-center gap-2"><carbon:checkmark-filled class="text-green-500 shrink-0" /> Pakai kunci kuat (2048-bit+)</li>
  <li class="flex items-center gap-2"><carbon:checkmark-filled class="text-green-500 shrink-0" /> Selalu cek expired & klaim lainnya</li>
  <li class="flex items-center gap-2"><carbon:checkmark-filled class="text-green-500 shrink-0" /> Pakai <code>kid</code> (Key ID)</li>
</ul>
</div>

</div>
