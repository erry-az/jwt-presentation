---
transition: slide-left
---

# JWS Algorithms (RFC 7518)

<div class="mt-4 overflow-hidden rounded-xl border border-gray-700">
<table class="w-full text-left border-collapse bg-gray-900/50">
  <thead class="bg-gray-800 text-blue-300">
    <tr>
      <th class="p-3 border-b border-gray-700">"alg" Name</th>
      <th class="p-3 border-b border-gray-700 text-xs">Full Algorithm Name</th>
      <th class="p-3 border-b border-gray-700">Implementation</th>
      <th class="p-3 border-b border-gray-700">Type</th>
    </tr>
  </thead>
  <tbody class="text-xs text-gray-300">
    <tr v-click class="hover:bg-blue-400/5 transition-colors">
      <td class="p-3 border-b border-gray-800 font-mono text-orange-400">HS256</td>
      <td class="p-3 border-b border-gray-800 italic">HMAC using SHA-256</td>
      <td class="p-3 border-b border-gray-800 font-bold text-green-400 text-[10px]">REQUIRED</td>
      <td class="p-3 border-b border-gray-800">Symmetric</td>
    </tr>
    <tr v-click class="hover:bg-blue-400/5 transition-colors">
      <td class="p-3 border-b border-gray-800 font-mono text-blue-400">RS256</td>
      <td class="p-3 border-b border-gray-800 italic">RSASSA-PKCS1-v1_5 w/ SHA-256</td>
      <td class="p-3 border-b border-gray-800 font-bold text-blue-300 text-[10px]">RECOMMENDED</td>
      <td class="p-3 border-b border-gray-800">Asymmetric</td>
    </tr>
    <tr v-click class="hover:bg-blue-400/5 transition-colors">
      <td class="p-3 border-b border-gray-800 font-mono text-teal-400">ES256</td>
      <td class="p-3 border-b border-gray-800 italic">ECDSA using P-256 and SHA-256</td>
      <td class="p-3 border-b border-gray-800 font-bold text-teal-300 text-[10px]">RECOMMENDED+</td>
      <td class="p-3 border-b border-gray-800">Asymmetric</td>
    </tr>
    <tr v-click class="hover:bg-blue-400/5 transition-colors">
      <td class="p-3 border-b border-gray-800 font-mono text-purple-400">PS256</td>
      <td class="p-3 border-b border-gray-800 italic">RSASSA-PSS using SHA-256</td>
      <td class="p-3 border-b border-gray-800 font-bold text-gray-400 text-[10px]">OPTIONAL</td>
      <td class="p-3 border-b border-gray-800">Asymmetric</td>
    </tr>
    <tr v-click class="hover:bg-red-400/5 transition-colors">
      <td class="p-3 border-b border-gray-800 font-mono text-red-500">none</td>
      <td class="p-3 border-b border-gray-800 italic text-red-400">No Digital Signature/MAC</td>
      <td class="p-3 border-b border-gray-800 font-bold text-red-500 text-[10px]">UNSAFE</td>
      <td class="p-3 border-b border-gray-800">-</td>
    </tr>
  </tbody>
</table>
</div>
