---
layout: default
transition: slide-left
---

# JWT: Pengenalan 4W

<div class="grid grid-cols-2 gap-8 mt-6">

<div v-click>
<h3 class="text-blue-400 mb-2 flex items-center gap-2">
  <carbon:help class="inline" /> Apa itu?
</h3>
<p class="text-sm text-gray-300 italic">
  "JSON Web Token (JWT) adalah cara yang <strong>ringkas</strong> dan <strong>aman bagi URL</strong> untuk merepresentasikan klaim yang ditransfer antara dua pihak."
  <br><br>
  <span class="text-xs text-gray-500">— RFC 7519</span>
</p>
</div>

<div v-click>
<h3 class="text-teal-400 mb-2 flex items-center gap-2">
  <carbon:rocket class="inline" /> Mengapa menggunakannya?
</h3>
<p class="text-sm text-gray-300 italic">
  JWT memungkinkan klaim untuk "ditandatangani secara digital atau di-MAC dan/atau dienkripsi," menyediakan verifikasi stateless yang mandiri.
  <br><br>
  <span class="text-xs text-gray-500">— RFC 7519, Bagian 1</span>
</p>
</div>

<div v-click class="mt-4">
<h3 class="text-purple-400 mb-2 flex items-center gap-2">
  <carbon:location class="inline" /> Di mana digunakan?
</h3>
<p class="text-sm text-gray-300 italic">
  Sebagai "format interoperabel standar untuk access token" dalam header HTTP Authorization untuk API dan mikroservis.
  <br><br>
  <span class="text-xs text-gray-500">— RFC 9068, Bagian 1</span>
</p>
</div>

<div v-click class="mt-4">
<h3 class="text-orange-400 mb-2 flex items-center gap-2">
  <carbon:time class="inline" /> Kapan digunakan?
</h3>
<p class="text-sm text-gray-300 italic">
  "Saat melewatkan klaim antar pihak" yang membutuhkan keamanan kriptografi, terutama di mana pencarian status pusat tidak memungkinkan.
  <br><br>
  <span class="text-xs text-gray-500">— RFC 8725 / 7519</span>
</p>
</div>

</div>

<div v-click class="mt-12 p-3 bg-gray-800/40 rounded-lg border border-gray-500/20 text-center text-sm italic text-gray-400">
  "JSON Web Token digunakan saat melewatkan klaim antar pihak dalam lingkungan dengan ruang terbatas."<br>
  — RFC 7519 (Pendahuluan)
</div>
