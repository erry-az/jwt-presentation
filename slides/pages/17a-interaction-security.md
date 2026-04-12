---
layout: center
class: text-center
transition: slide-up
---

<div
  v-motion
  :initial="{ y: -20, opacity: 0 }"
  :enter="{ y: 0, opacity: 1 }"
  class="text-6xl mb-8"
>
  🛡️
</div>

<h2 class="text-4xl font-bold mb-6 text-green-400">Reflection</h2>

<p class="text-2xl text-gray-200 max-w-2xl mx-auto leading-relaxed mb-8">
  "I'm using a popular library, so it must be secure by default."
</p>

<div 
  v-motion
  :initial="{ opacity: 0, scale: 0.9 }"
  :enter="{ opacity: 1, scale: 1, transition: { delay: 500 } }"
  class="text-3xl font-bold text-red-400"
>
  Do you trust your library blindly?
</div>

<div class="mt-12 text-gray-500 text-sm flex justify-center gap-8">
  <span v-click>Verifying <b>alg</b> header?</span>
  <span v-click>Validating <b>iss</b> / <b>aud</b>?</span>
  <span v-click>Handling <b>exp</b> correctly?</span>
</div>
