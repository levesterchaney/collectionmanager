<template>
  <main class="min-h-screen bg-gray-50 flex items-center justify-center">
    <div class="max-w-md w-full bg-white rounded-lg shadow-md p-8">
      <h1 class="text-3xl font-bold text-gray-900 mb-4">About</h1>
      <p class="text-gray-600 mb-6">
        This is a Vue 3 application with TypeScript, Vue Router, Pinia, and Tailwind CSS.
      </p>
      <div class="space-y-4">
        <router-link
          to="/"
          class="block w-full bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors text-center"
        >
          Back to Home
        </router-link>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
// About page component
</script>