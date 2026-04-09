---
transition: slide-up
---

# JWS vs JWE

<div class="grid grid-cols-2 gap-8 mt-4">

<div class="p-5 rounded-xl border border-blue-400/30 bg-blue-900/10">
<div class="text-blue-400 font-bold mb-3 flex items-center gap-2">
  <carbon:pen /> JWS (Signed)
</div>

- **Struktur**: 3 Bagian (Header.Payload.Sig)
- **Konten**: Base64 Encoded (Bisa dibaca langsung)
- **Tujuan**: <span v-mark.underline.blue>Keaslian & Integritas</span>
- **Dipakai Buat Apa?**: Login, Otorisasi, Klaim Open ID.
</div>

<div class="p-5 rounded-xl border border-purple-400/30 bg-purple-900/10">
<div class="text-purple-400 font-bold mb-3 flex items-center gap-2">
  <carbon:security /> JWE (Encrypted)
</div>

- **Struktur**: 5 Bagian (Isinya kunci & metadata)
- **Konten**: Terenkripsi (Biner/Ciphertext)
- **Tujuan**: <span v-mark.underline.purple>Kerahasiaan & Privasi</span>
- **Dipakai Buat Apa?**: Data PII Sensitif, Perbankan, atau privasi total.
</div>

</div>

<div v-click class="mt-8 p-4 rounded bg-gray-800/40 border border-gray-600/30 text-xs text-center text-gray-400 italic">
  "Gampangnya: JWS buat buktiin siapa yang kirim. JWE buat mastiin cuma penerima yang bisa baca."
</div>
