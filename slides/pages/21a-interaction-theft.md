---
layout: center
class: text-center
transition: slide-left
---

<div
  v-motion
  :initial="{ scale: 0.5, opacity: 0 }"
  :enter="{ scale: 1, opacity: 1 }"
  class="text-6xl mb-8"
>
  💸
</div>

<h2 class="text-4xl font-bold mb-6 text-purple-400">The Worst Case Scenario</h2>

<div class="text-3xl font-semibold mb-10 text-white max-w-2xl mx-auto">
  If an attacker steals a user's token right now... <br>
  <span class="text-red-400">how do you stop them?</span>
</div>

<div class="space-y-4 text-left max-w-md mx-auto">
  <div v-click class="flex gap-4 items-center p-3 rounded bg-gray-800/50 border border-gray-700">
    <div class="w-8 h-8 rounded-full bg-red-500/20 text-red-400 flex items-center justify-center font-bold">1</div>
    <div class="text-gray-300">Wait for the token to expire?</div>
  </div>
  <div v-click class="flex gap-4 items-center p-3 rounded bg-gray-800/50 border border-gray-700">
    <div class="w-8 h-8 rounded-full bg-orange-500/20 text-orange-400 flex items-center justify-center font-bold">2</div>
    <div class="text-gray-300">Blacklist the token (Stateful)?</div>
  </div>
  <div v-click class="flex gap-4 items-center p-3 rounded bg-gray-800/50 border border-gray-700">
    <div class="w-8 h-8 rounded-full bg-green-500/20 text-green-400 flex items-center justify-center font-bold">3</div>
    <div class="text-green-300 font-bold">Sender Constraining (DPoP)?</div>
  </div>
</div>
