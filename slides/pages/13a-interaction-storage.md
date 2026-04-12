---
layout: center
class: text-center
transition: slide-left
---

<div
  v-motion
  :initial="{ rotate: -15, opacity: 0 }"
  :enter="{ rotate: 0, opacity: 1, transition: { duration: 600 } }"
  class="text-6xl mb-8"
>
  🗳️
</div>

<h2 class="text-4xl font-bold mb-8 text-orange-400">Audience Poll</h2>

<div class="text-3xl font-semibold mb-10 text-white">
  Where do you usually store your JWTs?
</div>

<div class="grid grid-cols-2 gap-4 max-w-xl mx-auto">
  <div v-click class="p-4 rounded-xl border border-blue-400/30 bg-blue-900/10 text-blue-300 font-mono text-lg">
    LocalStorage
  </div>
  <div v-click class="p-4 rounded-xl border border-pink-400/30 bg-pink-900/10 text-pink-300 font-mono text-lg">
    HttpOnly Cookies
  </div>
  <div v-click class="p-4 rounded-xl border border-green-400/30 bg-green-900/10 text-green-300 font-mono text-lg">
    In-Memory (State)
  </div>
  <div v-click class="p-4 rounded-xl border border-gray-400/30 bg-gray-900/10 text-gray-400 font-mono text-lg italic">
    I'm not sure...
  </div>
</div>
