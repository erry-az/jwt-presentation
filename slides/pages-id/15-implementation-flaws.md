---
transition: slide-up
---

# Dosa-Dosa Saat Implementasi

<div class="mt-8 space-y-8">

<div
  v-motion
  :initial="{ x: 80, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { type: 'spring', stiffness: 100 } }"
>
  <div class="flex items-start gap-4">
    <div class="text-3xl mt-1">💀</div>
    <div>
      <div class="font-bold text-red-300 text-xl mb-1">Bypass <code>alg: none</code></div>
      <div class="text-gray-400 mb-3">Kalau backend asal percaya sama algoritma di header, penyerang bisa pakai <code>alg: none</code> buat nembus sistem dan naikkan hak akses.</div>

```json
{ "alg": "none", "typ": "JWT" }
```
    </div>
  </div>
</div>

<div
  v-motion
  :initial="{ x: 80, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { type: 'spring', stiffness: 100 } }"
>
  <div class="flex items-start gap-4">
    <div class="text-3xl mt-1">🔑</div>
    <div>
      <div class="font-bold text-orange-300 text-xl mb-1">Serangan Key Confusion</div>
      <div class="text-gray-400">Nge-trick backend dengan kirim token HMAC tapi pakai kunci publik RSA. Hasilnya? Token palsu malah dianggap asli.</div>
    </div>
  </div>
</div>

</div>
