---
transition: slide-up
---

# Algoritma JWS: Mana yang Harus Dipilih?

<div class="grid grid-cols-2 gap-x-8 gap-y-6 mt-8">

<div v-click>
  <div class="flex items-center gap-2 mb-2">
    <div class="w-2 h-2 rounded-full bg-blue-400"></div>
    <div class="font-bold text-blue-300 text-lg">HMAC (HS256/512)</div>
  </div>
  <div class="text-gray-400 text-sm">Symmetric (Kunci barengan). Emang cepet, tapi bahaya kalau bagi-bagi rahasia ke banyak orang.</div>
</div>

<div v-click>
  <div class="flex items-center gap-2 mb-2">
    <div class="w-2 h-2 rounded-full bg-purple-400"></div>
    <div class="font-bold text-purple-300 text-lg">RSA (RS256/512)</div>
  </div>
  <div class="text-gray-400 text-sm">Asymmetric (Kunci Privat/Publik). Standar industri, tapi butuh infrastruktur kunci publik.</div>
</div>

<div v-click>
  <div class="flex items-center gap-2 mb-2">
    <div class="w-2 h-2 rounded-full bg-green-400"></div>
    <div class="font-bold text-green-300 text-lg">ECDSA (ES256/512)</div>
  </div>
  <div class="text-gray-400 text-sm">Elliptic Curve. Signature lebih pendek dan performanya lebih kencang dibanding RSA.</div>
</div>

<div v-click>
  <div class="flex items-center gap-2 mb-2">
    <div class="w-2 h-2 rounded-full bg-orange-400"></div>
    <div class="font-bold text-orange-300 text-lg">EdDSA (Ed25519)</div>
  </div>
  <div class="text-gray-400 text-sm">Modern, aman banget, dan performanya juara.</div>
</div>

</div>
