---
transition: fade-out
---

# Anatomi JWT

**JSON Web Token (JWT)** standar terdiri dari tiga bagian yang dipisahkan oleh titik `(.)`

<v-click>

<div
  v-motion
  :initial="{ x: -40, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { duration: 500 } }"
  class="mt-6 p-4 rounded-lg border border-red-400/40 bg-red-900/10 flex items-start gap-3"
>
  <div class="text-2xl">📄</div>
  <div>
    <div class="text-red-400 font-bold">Header</div>
    <div class="text-gray-300 text-sm">Berisi metadata seperti tipe token dan algoritma penandatanganan (misal: <code>HS256</code>, <code>RS256</code>).</div>
  </div>
</div>

</v-click>

<v-click>

<div
  v-motion
  :initial="{ x: -40, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { duration: 500 } }"
  class="mt-3 p-4 rounded-lg border border-purple-400/40 bg-purple-900/10 flex items-start gap-3"
>
  <div class="text-2xl">📦</div>
  <div>
    <div class="text-purple-400 font-bold">Payload</div>
    <div class="text-gray-300 text-sm">Berisi klaim — pernyataan tentang suatu entitas dan data tambahan.</div>
  </div>
</div>

</v-click>

<v-click>

<div
  v-motion
  :initial="{ x: -40, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { duration: 500 } }"
  class="mt-3 p-4 rounded-lg border border-blue-400/40 bg-blue-900/10 flex items-start gap-3"
>
  <div class="text-2xl">🔏</div>
  <div>
    <div class="text-blue-400 font-bold">Tanda Tangan (Signature)</div>
    <div class="text-gray-300 text-sm">Digunakan untuk memverifikasi bahwa token tidak diubah dalam perjalanan.</div>
  </div>
</div>

</v-click>

<div v-click class="mt-8 p-4 bg-gray-800/80 rounded-xl shadow-lg font-mono text-center text-xl tracking-wider border border-gray-600/40">
  <span class="text-red-400 font-bold">HEADER</span><span class="text-gray-500">.</span><span class="text-purple-400 font-bold">PAYLOAD</span><span class="text-gray-500">.</span><span class="text-blue-400 font-bold">SIGNATURE</span>
</div>
