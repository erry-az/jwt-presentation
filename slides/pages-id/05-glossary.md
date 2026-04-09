---
transition: slide-up
---

# Referensi RFC yang Kita Pakai

<div class="grid grid-cols-2 gap-x-8 gap-y-6 mt-8">

<div v-click>
  <div class="text-blue-400 font-bold">Signature</div>
  <div class="text-xs text-gray-400">Bagian dari JWT yang mastiin data nggak dimanipulasi orang lain.</div>
</div>

<div v-click>
  <div class="text-red-400 font-bold">Bearer Token</div>
  <div class="text-xs text-gray-400">Token yang kasih akses ke siapa aja yang pegang (bahaya kalau sampai kecuri).</div>
</div>

<div v-click>
  <div class="text-teal-400 font-bold">Proof-of-Possession</div>
  <div class="text-xs text-gray-400">Syarat buat pengirim buktiin kalau mereka beneran punya kunci privatnya.</div>
</div>

<div v-click>
  <div class="text-yellow-400 font-bold">XSS / CSRF</div>
  <div class="text-xs text-gray-400">XSS (Curi token lewat skrip jahat). CSRF (Pakai token orang lain tanpa izin).</div>
</div>

<div v-click>
  <div class="text-pink-400 font-bold">Simetris / Asimetris</div>
  <div class="text-xs text-gray-400">Simetris (Pakai 1 rahasia barengan). Asimetris (Pakai pasangan kunci Privat/Publik).</div>
</div>

</div>
