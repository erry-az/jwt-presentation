---
layout: default
transition: slide-left
---

# JWT: The 4W Introduction

<div class="grid grid-cols-2 gap-8 mt-6">

<div v-click>
<h3 class="text-blue-400 mb-2 flex items-center gap-2">
  <carbon:help class="inline" /> What is it?
</h3>
<p class="text-sm text-gray-300 italic">
  "JSON Web Token (JWT) is a <strong>compact</strong>, <strong>URL-safe</strong> means of representing claims to be transferred between two parties."
  <br><br>
  <span class="text-xs text-gray-500">— RFC 7519</span>
</p>
</div>

<div v-click>
<h3 class="text-teal-400 mb-2 flex items-center gap-2">
  <carbon:rocket class="inline" /> Why use it?
</h3>
<p class="text-sm text-gray-300 italic">
  They enable claims to "be digitally signed or MACed and/or encrypted," providing self-contained, stateless verification.
  <br><br>
  <span class="text-xs text-gray-500">— RFC 7519, Section 1</span>
</p>
</div>

<div v-click class="mt-4">
<h3 class="text-purple-400 mb-2 flex items-center gap-2">
  <carbon:location class="inline" /> Where to use?
</h3>
<p class="text-sm text-gray-300 italic">
  As a "standard, interoperable format for access tokens" in HTTP Authorization headers for APIs and microservices.
  <br><br>
  <span class="text-xs text-gray-500">— RFC 9068, Section 1</span>
</p>
</div>

<div v-click class="mt-4">
<h3 class="text-orange-400 mb-2 flex items-center gap-2">
  <carbon:time class="inline" /> When to use?
</h3>
<p class="text-sm text-gray-300 italic">
  "When passing claims between parties" requiring cryptographic security, especially where central state lookups aren't feasible.
  <br><br>
  <span class="text-xs text-gray-500">— RFC 8725 / 7519</span>
</p>
</div>

</div>

<div v-click class="mt-12 p-3 bg-gray-800/40 rounded-lg border border-gray-500/20 text-center text-sm italic text-gray-400">
  "Don't just use it because it's popular; use it because you need decentralization."
</div>
