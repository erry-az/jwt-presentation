<script setup lang="ts">
import { computed } from 'vue'

// In Slidev, we can check the current path or a global state.
// However, since we are using separate entry points (different build targets),
// we can simply check the window.location or use a prop if we want to be more dynamic.
// For now, let's assume it's based on the URL path.

const isIndonesian = computed(() => {
  if (typeof window === 'undefined') return false
  return window.location.pathname.includes('/id/')
})

const toggleLanguage = (lang: 'en' | 'id') => {
  if (typeof window === 'undefined') return
  if (lang === 'id') {
    // If we are in dev, this might not work perfectly without the separate process,
    // but for production builds it will navigate / -> /id/
    window.location.href = window.location.origin + '/id/' + window.location.hash
  } else {
    window.location.href = window.location.origin + '/' + window.location.hash
  }
}
</script>

<template>
  <div class="fixed bottom-4 right-4 z-50 flex gap-2 p-1.5 rounded-full bg-white/10 backdrop-blur-md border border-white/20 shadow-lg transition-all hover:bg-white/20">
    <button 
      @click="toggleLanguage('en')"
      class="w-10 h-10 flex items-center justify-center rounded-full transition-all duration-300 transform"
      :class="!isIndonesian ? 'bg-white/20 scale-110 shadow-sm' : 'opacity-50 hover:opacity-100 hover:scale-105'"
      title="English"
    >
      <span class="text-2xl mt-[-2px]">🇺🇸</span>
    </button>
    <button 
      @click="toggleLanguage('id')"
      class="w-10 h-10 flex items-center justify-center rounded-full transition-all duration-300 transform"
      :class="isIndonesian ? 'bg-white/20 scale-110 shadow-sm' : 'opacity-50 hover:opacity-100 hover:scale-105'"
      title="Bahasa Indonesia"
    >
      <span class="text-2xl mt-[-2px]">🇮🇩</span>
    </button>
  </div>
</template>

<style scoped>
button {
  cursor: pointer;
  border: none;
  background: transparent;
  outline: none;
}
</style>
