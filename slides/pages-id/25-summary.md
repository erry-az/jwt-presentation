---
layout: center
class: text-center
---

<div
  v-motion
  :initial="{ y: 30, opacity: 0 }"
  :enter="{ y: 0, opacity: 1, transition: { delay: 400, duration: 700 } }"
>

# Ringkasan & Poin Penting

</div>

<div class="mt-10 space-y-6 max-w-3xl mx-auto text-left">

<div v-click
  v-motion
  :initial="{ x: -30, opacity: 0 }"
  :enter="{ x: 0, opacity: 1 }"
  class="flex gap-3 items-center"
>
  <span class="text-xl">✅</span>
  <span class="text-gray-300">JWT itu standar stateless yang simpel buat bawa data kita.</span>
</div>

<div v-click
  v-motion
  :initial="{ x: -30, opacity: 0 }"
  :enter="{ x: 0, opacity: 1 }"
  class="flex gap-3 items-center"
>
  <span class="text-xl">🛡️</span>
  <span class="text-gray-300">Wajib ikutin <b>RFC 8725</b> kalau mau implementasi JWT yang aman.</span>
</div>

<div v-click
  v-motion
  :initial="{ x: -30, opacity: 0 }"
  :enter="{ x: 0, opacity: 1 }"
  class="flex gap-3 items-center"
>
  <span class="text-xl">⚡</span>
  <span class="text-gray-300">Bearer token emang rentan — kalau dicuri, penyerang dapet akses penuh.</span>
</div>

<div v-click
  v-motion
  :initial="{ x: -30, opacity: 0 }"
  :enter="{ x: 0, opacity: 1 }"
  class="flex gap-3 items-center"
>
  <span class="text-xl">🚀</span>
  <span class="text-gray-300"><span v-mark.underline.purple>DPoP "ngunci" token ke klien</span> — biarpun dicuri, tokennya nggak bakal guna.</span>
</div>

</div>

<div
  v-motion
  :initial="{ opacity: 0 }"
  :enter="{ opacity: 1, transition: { delay: 2000 } }"
  class="mt-10 text-gray-500 text-sm"
>
  RFC 7519 · RFC 8725 · RFC 9068 · RFC 9449
</div>

<PoweredBySlidev
  v-motion
  :initial="{ opacity: 0 }"
  :enter="{ opacity: 1, transition: { delay: 2500 } }"
  class="mt-6"
/>
