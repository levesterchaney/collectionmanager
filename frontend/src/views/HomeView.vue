<template>
  <main class="min-h-screen bg-gray-50 flex items-center justify-center">
    <div class="max-w-md w-full bg-white rounded-lg shadow-md p-8">
      <h1 class="text-3xl font-bold text-gray-900 mb-4">Collection Manager</h1>
      <p class="text-gray-600 mb-6">Welcome to your collection management system.</p>
      <div class="space-y-4">
        <button
          @click="counter.increment()"
          class="w-full bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700 transition-colors"
        >
          Count: {{ counter.count }}
        </button>
        <router-link
          to="/about"
          class="block w-full bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors text-center"
        >
          About
        </router-link>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { useCounterStore } from '@/stores/counter'

const counter = useCounterStore()
</script>