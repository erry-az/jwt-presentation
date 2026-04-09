---
transition: slide-up
---

# JWKS: Ngurus Kunci Jarak Jauh

<div class="mt-8">
  <div class="p-6 rounded-2xl border border-gray-600/30 bg-gray-800/40">
    <div class="text-blue-300 font-bold mb-4 flex items-center gap-3 text-xl">
      <carbon:cloud-download /> JSON Web Key Set (RFC 7517)
    </div>
    <div class="space-y-4 text-gray-300">
      <div v-click class="flex gap-4 items-start">
        <div class="p-2 rounded-lg bg-blue-900/40 text-blue-400 shrink-0">1.</div>
        <div>Bagi-bagi **Kunci Publik** aja, biar backend bisa cek token sendiri secara mandiri.</div>
      </div>
      <div v-click class="flex gap-4 items-start">
        <div class="p-2 rounded-lg bg-blue-900/40 text-blue-400 shrink-0">2.</div>
        <div>Mendukung **Rotasi Kunci** otomatis tanpa perlu restart-restart server backend.</div>
      </div>
      <div v-click class="flex gap-4 items-start">
        <div class="p-2 rounded-lg bg-blue-900/40 text-blue-400 shrink-0">3.</div>
        <div>Punya endpoint standar biar gampang dicari: <code>/.well-known/jwks.json</code></div>
      </div>
    </div>
  </div>
</div>
