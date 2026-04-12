---
layout: center
class: text-center
transition: slide-up
---

<div
  v-motion
  :initial="{ scale: 0.8, opacity: 0 }"
  :enter="{ scale: 1, opacity: 1, transition: { type: 'spring', stiffness: 200, damping: 15 } }"
  class="text-6xl mb-8"
>
  🤔
</div>

<h2 
  v-motion
  :initial="{ y: 20, opacity: 0 }"
  :enter="{ y: 0, opacity: 1, transition: { delay: 200 } }"
  class="text-4xl font-bold mb-6 text-blue-400"
>
  Quick Quiz
</h2>

<p 
  v-motion
  :initial="{ y: 20, opacity: 0 }"
  :enter="{ y: 0, opacity: 1, transition: { delay: 400 } }"
  class="text-2xl text-gray-200 max-w-2xl mx-auto leading-relaxed"
>
  Anyone knows the difference between <br>
  <span class="text-blue-300 font-mono">Encoding</span>, 
  <span class="text-orange-300 font-mono">Hashing</span>, and 
  <span class="text-purple-300 font-mono">Encryption</span>?
</p>

<div 
  v-motion
  :initial="{ opacity: 0 }"
  :enter="{ opacity: 1, transition: { delay: 800 } }"
  class="mt-12 text-gray-500 italic text-sm"
>
  Raise your hand if you think you know! 🙋‍♂️
</div>
