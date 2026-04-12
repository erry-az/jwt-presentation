---
layout: center
class: text-center
transition: fade-out
---

<div
  v-motion
  :initial="{ scale: 0.5, opacity: 0 }"
  :enter="{ scale: 1, opacity: 1, transition: { type: 'spring', stiffness: 200, damping: 15, delay: 200 } }"
>

# Agenda

</div>

<div class="grid grid-cols-2 gap-6 mt-8 text-left">

<div
  v-motion
  :initial="{ x: -60, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { delay: 400, duration: 600 } }"
  class="p-5 rounded-xl border border-blue-400/30 bg-blue-900/20"
>
  <div class="text-blue-400 text-lg font-bold mb-2">📋 Part 1</div>
  <div class="text-white font-semibold">The Stateless Dream</div>
  <div class="text-gray-400 text-sm">RFC 7519 — JWT Anatomy & Security</div>
</div>

<div
  v-motion
  :initial="{ x: 60, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { delay: 600, duration: 600 } }"
  class="p-5 rounded-xl border border-red-400/30 bg-red-900/20"
>
  <div class="text-red-400 text-lg font-bold mb-2">⚠️ Part 2</div>
  <div class="text-white font-semibold">The Bearer Token Crisis</div>
  <div class="text-gray-400 text-sm">Where Bearer Tokens Fall Short</div>
</div>

<div
  v-motion
  :initial="{ x: -60, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { delay: 800, duration: 600 } }"
  class="p-5 rounded-xl border border-green-400/30 bg-green-900/20"
>
  <div class="text-green-400 text-lg font-bold mb-2">🔒 Part 3</div>
  <div class="text-white font-semibold">The Official Fixes</div>
  <div class="text-gray-400 text-sm">RFC 8725 & RFC 9068</div>
</div>

<div
  v-motion
  :initial="{ x: 60, opacity: 0 }"
  :enter="{ x: 0, opacity: 1, transition: { delay: 1000, duration: 600 } }"
  class="p-5 rounded-xl border border-purple-400/30 bg-purple-900/20"
>
  <div class="text-purple-400 text-lg font-bold mb-2">🚀 Part 4</div>
  <div class="text-white font-semibold">The Zero-Trust Evolution</div>
  <div class="text-gray-400 text-sm">RFC 9449 — Proof-of-Possession</div>
</div>

</div>
