---
layout: center
class: text-center
transition: slide-left
---

<div
  v-motion
  :initial="{ scale: 0.3, rotate: -10, opacity: 0 }"
  :enter="{ scale: 1, rotate: 0, opacity: 1, transition: { type: 'spring', stiffness: 150, damping: 12 } }"
  class="text-6xl mb-4"
>
  ⚠️
</div>

<div
  v-motion
  :initial="{ y: 30, opacity: 0 }"
  :enter="{ y: 0, opacity: 1, transition: { delay: 400, duration: 700 } }"
  class="text-5xl font-bold"
>
  Part 2
</div>

<div
  v-motion
  :initial="{ y: 30, opacity: 0 }"
  :enter="{ y: 0, opacity: 1, transition: { delay: 600, duration: 700 } }"
  class="text-2xl text-red-300 mt-2"
>
  The Vulnerabilities
</div>

<div
  v-motion
  :initial="{ opacity: 0 }"
  :enter="{ opacity: 1, transition: { delay: 900 } }"
  class="text-gray-400 mt-2"
>
  Where Bearer Tokens Fall Short
</div>
