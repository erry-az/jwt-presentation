---
transition: slide-left
---

# JWT Anatomy: The Details

<div class="grid grid-cols-3 gap-4 mt-4">

<div v-click>
<div class="text-red-400 font-bold mb-2">1. Header (JSON)</div>

```json
{
  "alg": "RS256",
  "typ": "JWT",
  "kid": "key-id-123"
}
```

<div class="text-[10px] text-gray-500 mt-1 italic">Tells the server how to verify the signature.</div>
</div>

<div v-click>
<div class="text-purple-400 font-bold mb-2">2. Payload (Claims)</div>

```json
{
  "sub": "1234567890",
  "name": "John Doe",
  "admin": true,
  "iat": 1516239022,
  "exp": 1516242622
}
```

<div class="text-[10px] text-gray-500 mt-1 italic">The actual data and metadata (Registered, Public, Private claims).</div>
</div>

<div v-click>
<div class="text-blue-400 font-bold mb-2">3. Signature</div>

```javascript
RSASHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  publicKey,
  privateKey
)
```

<div class="text-[10px] text-gray-500 mt-1 italic">Ensures the message wasn't changed along the way.</div>
</div>

</div>

<div v-click class="mt-8 p-3 rounded bg-blue-900/10 border border-blue-400/20">
  <div class="text-blue-300 text-xs text-center font-mono break-all">
    eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImtleS1pZC0xMjMifQ.<br>
    eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMiwiZXhwIjoxNTE2MjQyNjIyfQ.<br>
    <span class="opacity-50">SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c...</span>
  </div>
</div>
