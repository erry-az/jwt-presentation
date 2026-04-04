---
transition: slide-up
---

# JWS vs JWE

<div class="grid grid-cols-2 gap-8 mt-4">

<div class="p-5 rounded-xl border border-blue-400/30 bg-blue-900/10">
<div class="text-blue-400 font-bold mb-3 flex items-center gap-2">
  <carbon:pen /> JWS (Signed)
</div>

- **Structure**: 3 Parts (Header.Payload.Sig)
- **Content**: Base64 Encoded (Readable)
- **Goal**: <span v-mark.underline.blue>Integrity & Authenticity</span>
- **Use Case**: Authentication, Authorization, Open ID Claims.
</div>

<div class="p-5 rounded-xl border border-purple-400/30 bg-purple-900/10">
<div class="text-purple-400 font-bold mb-3 flex items-center gap-2">
  <carbon:security /> JWE (Encrypted)
</div>

- **Structure**: 5 Parts (various keys/metadata)
- **Content**: Encrypted (Binary/Ciphertext)
- **Goal**: <span v-mark.underline.purple>Confidentiality & Privacy</span>
- **Use Case**: Sensitive PII, Banking data, end-to-end privacy.
</div>

</div>

<div v-click class="mt-8 p-4 rounded bg-gray-800/40 border border-gray-600/30 text-xs text-center text-gray-400 italic">
  "JWS proves who sent the data. JWE ensures only the intended receiver can see it."
</div>
