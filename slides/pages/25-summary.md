---
layout: center
class: text-center
transition: fade-out
---

<div
  v-motion
  :initial="{ scale: 0.5, opacity: 0 }"
  :enter="{ scale: 1, opacity: 1, transition: { type: 'spring', stiffness: 200, damping: 15 } }"
  class="text-6xl mb-6"
>
  🎯
</div>

<div
  v-motion
  :initial="{ y: 20, opacity: 0 }"
  :enter="{ y: 0, opacity: 1, transition: { delay: 400, duration: 700 } }"
  class="text-4xl font-bold mb-4"
>
  Key Takeaways
</div>

<div class="text-left space-y-3 max-w-xl mx-auto mt-6">

<div v-click
  v-motion
  :initial="{ x: -30, opacity: 0 }"
  :enter="{ x: 0, opacity: 1 }"
  class="flex gap-3 items-center"
>
  <span class="text-xl">📋</span>
  <span class="text-gray-300">JWT is base64url encoded, <span v-mark.underline.yellow>not encrypted</span> — treat it accordingly</span>
</div>

<div v-click
  v-motion
  :initial="{ x: -30, opacity: 0 }"
  :enter="{ x: 0, opacity: 1 }"
  class="flex gap-3 items-center"
>
  <span class="text-xl">🛡️</span>
  <span class="text-gray-300">Always <span v-mark.underline.green>hardcode algorithms</span> and strictly validate all claims</span>
</div>

<div v-click
  v-motion
  :initial="{ x: -30, opacity: 0 }"
  :enter="{ x: 0, opacity: 1 }"
  class="flex gap-3 items-center"
>
  <span class="text-xl">⚡</span>
  <span class="text-gray-300">Bearer tokens are inherently vulnerable — <span v-mark.underline.red>stolen = full access</span></span>
</div>

<div v-click
  v-motion
  :initial="{ x: -30, opacity: 0 }"
  :enter="{ x: 0, opacity: 1 }"
  class="flex gap-3 items-center"
>
  <span class="text-xl">🚀</span>
  <span class="text-gray-300"><span v-mark.underline.purple>DPoP binds tokens</span> to the client — stolen tokens are useless</span>
</div>

</div>

<div
  v-motion
  :initial="{ opacity: 0 }"
  :enter="{ opacity: 1, transition: { delay: 2000 } }"
  class="mt-10 text-gray-500 text-sm"
>
  RFC 7519 · RFC 8725 · RFC 9068 · RFC 9449
</div>

<PoweredBySlidev
  v-motion
  :initial="{ opacity: 0 }"
  :enter="{ opacity: 1, transition: { delay: 2500 } }"
  mt-6
/>
