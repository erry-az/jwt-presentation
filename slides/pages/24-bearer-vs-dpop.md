---
layout: two-cols
layoutClass: gap-8
transition: slide-left
---

# Bearer vs DPoP

<div class="mt-2 text-gray-400 text-sm">Side-by-Side Comparison</div>

<div v-click class="mt-4">

### 🔴 Standard Bearer Request

```http
GET /api/data HTTP/1.1
Host: api.example.com
Authorization: Bearer eyJhbG...
```

<div
  v-motion
  :initial="{ opacity: 0 }"
  :enter="{ opacity: 1, transition: { delay: 400 } }"
  class="mt-3 p-3 rounded-lg bg-red-900/20 border border-red-500/30 text-red-300 text-sm italic"
>
  ⚡ If stolen, anyone can replay this request from anywhere in the world.
</div>

</div>

::right::

<div class="mt-14">

<div v-click>

### 🟢 DPoP Request

```http
GET /api/data HTTP/1.1
Host: api.example.com
Authorization: DPoP eyJhbG...
DPoP: eyJ0eXAiOiJkcG9wK2p3dCIs...
```

<div
  v-motion
  :initial="{ opacity: 0 }"
  :enter="{ opacity: 1, transition: { delay: 400 } }"
  class="mt-3 p-3 rounded-lg bg-green-900/20 border border-green-500/30 text-green-300 text-sm italic"
>
  ✅ The DPoP header is a fresh JWT covering <code>htm</code>=GET, <code>htu</code>=/api/data. Stolen tokens are useless without the private key!
</div>

</div>

</div>
