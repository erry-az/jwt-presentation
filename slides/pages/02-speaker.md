---
layout: two-cols
layoutClass: gap-8
transition: slide-up
---

# About the Speaker

<div
  v-motion
  :initial="{ opacity: 0, y: 20 }"
  :enter="{ opacity: 1, y: 0, transition: { delay: 200, duration: 600 } }"
>

<div class="mt-4 flex items-center gap-4">
  <div class="w-16 h-16 rounded-full bg-gradient-to-br from-blue-400 to-purple-600 flex items-center justify-center text-3xl font-bold text-white shadow-lg">
    EA
  </div>
  <div>
    <div class="text-2xl font-bold text-white">Erry Azhari</div>
    <div class="text-blue-300 font-semibold">Technical Lead</div>
    <div class="text-gray-400 text-sm">📍 Jakarta, Indonesia</div>
  </div>
</div>

</div>

<div
  v-click
  :initial="{ opacity: 0, x: -20 }"
  :enter="{ opacity: 1, x: 0, transition: { duration: 500 } }"
  class="mt-5 p-3 rounded-lg bg-gray-800/50 border border-gray-600/30 text-sm text-gray-300 italic"
>
  Highly motivated Technical Lead with <span v-mark.underline.blue>13+ years</span> of experience building robust, scalable, and performant backend systems. Expert in complex architectures, leadership, and mentorship.
</div>

<div v-click class="mt-5 space-y-2">

<div class="flex items-center gap-2 text-sm text-gray-300">
  <carbon:location class="text-blue-400" /> PT Pakar Digital Global (paper.id)
</div>
<div class="flex items-center gap-2 text-sm text-gray-300">
  <carbon:email class="text-blue-400" /> erry.az94@gmail.com
</div>
<div class="flex items-center gap-2 text-sm text-gray-300">
  <carbon:logo-github class="text-blue-400" /> github.com/erry-az
</div>
<div class="flex items-center gap-2 text-sm text-gray-300">
  <carbon:logo-linkedin class="text-blue-400" /> erry-azhari-b446a910a
</div>

</div>

::right::

<div class="mt-14">

<div v-click
  v-motion
  :initial="{ opacity: 0, x: 30 }"
  :enter="{ opacity: 1, x: 0, transition: { duration: 500 } }"
>

**Expanded Tech Stack**

<div class="grid grid-cols-2 gap-2 mt-3">

<div class="p-2 rounded bg-blue-900/30 border border-blue-400/20 text-center text-xs text-blue-300">
  Go · PHP · Java · C#
</div>
<div class="p-2 rounded bg-purple-900/30 border border-purple-400/20 text-center text-xs text-purple-300">
  GCP · AWS · DigitalOcean
</div>
<div class="p-2 rounded bg-green-900/30 border border-green-400/20 text-center text-xs text-green-300">
  gRPC · GQL · Kafka · NSQ
</div>
<div class="p-2 rounded bg-orange-900/30 border border-orange-400/20 text-center text-xs text-orange-300">
  MySQL · Postgres · Redis
</div>

</div>

</div>

<div v-click
  v-motion
  :initial="{ opacity: 0, x: 30 }"
  :enter="{ opacity: 1, x: 0, transition: { delay: 100, duration: 500 } }"
  class="mt-4"
>

**Notable Experience**

<div class="mt-2 space-y-2 text-sm">
<div class="flex gap-2 items-center text-gray-400">
  <span class="text-yellow-400">▸</span> <span class="text-gray-200">Tokopedia</span> — Senior Software Engineer
</div>
<div class="flex gap-2 items-center text-gray-400">
  <span class="text-yellow-400">▸</span> <span class="text-gray-200">Smeshub</span> — Principal Engineer
</div>
<div class="flex gap-2 items-center text-gray-400">
  <span class="text-yellow-400">▸</span> <span class="text-gray-200">paper.id</span> — Technical Lead <span class="text-blue-400 text-xs ml-1">(current)</span>
</div>
</div>

</div>

<div v-click
  v-motion
  :initial="{ opacity: 0, y: 10 }"
  :enter="{ opacity: 1, y: 0, transition: { delay: 200, duration: 500 } }"
  class="mt-4 p-3 rounded-lg border border-teal-400/30 bg-teal-900/20 text-xs text-gray-300"
>
  🎓 **Budi Luhur University** — Bachelor of Computer Science <br>
  🏫 **Vocational High School Telkom** — Jakarta
</div>

</div>
