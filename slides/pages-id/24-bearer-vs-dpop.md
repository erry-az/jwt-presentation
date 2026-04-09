---
transition: slide-up
---

# Perbandingan: Bearer vs DPoP

<div class="grid grid-cols-2 gap-8 mt-10">

<div v-click class="p-6 rounded-2xl border border-red-500/30 bg-red-900/10 h-full relative overflow-hidden">
  <div class="absolute top-0 right-0 p-2 bg-red-500 text-white text-[10px] font-bold uppercase tracking-widest">Bearer</div>
  <div class="text-5xl mb-6">🔓</div>
  <div class="text-red-300 font-bold text-xl mb-2">Token Dicuri = Tamat</div>
  <p class="text-gray-400 text-sm">Gak ada cara buat buktiin kalau yang bawa token itu beneran pemilik aslinya.</p>
  <div class="mt-6 flex flex-col gap-2 font-mono text-[10px] text-gray-500">
    <div class="flex justify-between border-b border-gray-700 pb-1">
      <span>Konteks</span>
      <span>Apa pun</span>
    </div>
    <div class="flex justify-between border-b border-gray-700 pb-1">
      <span>Pembatasan</span>
      <span>Nggak Ada</span>
    </div>
  </div>
</div>

<div v-click class="p-6 rounded-2xl border border-purple-500/40 bg-purple-900/20 h-full relative overflow-hidden shadow-[0_0_30px_-10px_rgba(168,85,247,0.4)]">
  <div class="absolute top-0 right-0 p-2 bg-purple-500 text-white text-[10px] font-bold uppercase tracking-widest">DPoP</div>
  <div class="text-5xl mb-6">🔐</div>
  <div class="text-purple-300 font-bold text-xl mb-2">Token Dicuri = Gak Guna</div>
  <p class="text-gray-400 text-sm">Token udah "dikunci" mati ke kunci privat punya si klien.</p>
  <div class="mt-6 flex flex-col gap-2 font-mono text-[10px] text-gray-500">
    <div class="flex justify-between border-b border-gray-700 pb-1">
      <span>Konteks</span>
      <span>Dibatasi Pengirim</span>
    </div>
    <div class="flex justify-between border-b border-gray-700 pb-1">
      <span>Keamanan</span>
      <span>Anti-Pencurian</span>
    </div>
  </div>
</div>

</div>
