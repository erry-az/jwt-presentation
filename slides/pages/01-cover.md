# JWT Best Practices and DPoP

<div class="text-2xl text-gray-300 mt-2">Demonstrating Proof-of-Possession</div>

<div
  v-motion
  :initial="{ opacity: 0, y: 40 }"
  :enter="{ opacity: 1, y: 0, transition: { delay: 600, duration: 800 } }"
  class="mt-12"
>
  <div @click="$slidev.nav.next" class="px-6 py-2 rounded-full border border-white/30 cursor-pointer inline-block hover:bg-white/10 transition-all">
    Press Space to begin <carbon:arrow-right class="inline" />
  </div>
</div>

<div
  v-motion
  :initial="{ opacity: 0 }"
  :enter="{ opacity: 1, transition: { delay: 1200 } }"
  class="abs-br m-6 text-sm text-gray-400"
>
  RFC 7519 · RFC 8725 · RFC 9068 · RFC 9449
</div>

<!--
Welcome! Today we'll explore JSON Web Tokens — from the standard, through the vulnerabilities, to the cutting-edge DPoP solution.
-->
