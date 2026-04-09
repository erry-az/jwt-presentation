---
layout: center
class: text-center
transition: slide-left
---

<div
  v-motion
  :initial="{ scale: 0.3, opacity: 0 }"
  :enter="{ scale: 1, opacity: 1, transition: { type: 'spring', stiffness: 200, damping: 15 } }"
  class="text-6xl mb-4"
>
  🚀
</div>

<div
  v-motion
  :initial="{ y: 30, opacity: 0 }"
  :enter="{ y: 0, opacity: 1, transition: { delay: 400, duration: 700 } }"
  class="text-5xl font-bold"
>
  Part 4
</div>

<div
  v-motion
  :initial="{ y: 30, opacity: 0 }"
  :enter="{ y: 0, opacity: 1, transition: { delay: 600, duration: 700 } }"
  class="text-3xl text-purple-300 mt-2 font-semibold"
>
  The Zero-Trust Evolution
</div>

<div
  v-motion
  :initial="{ opacity: 0 }"
  :enter="{ opacity: 1, transition: { delay: 900 } }"
  class="text-gray-400 mt-4 italic text-sm max-w-2xl mx-auto"
>
  "A mechanism for sender-constraining OAuth 2.0 tokens via a proof-of-possession mechanism on the application level."<br>
  — RFC 9449
</div>
