<script setup>
import { RouterLink } from 'vue-router';
</script>

<template>
  <div class="max-w-md mx-auto mt-10 bg-white p-6 rounded-lg shadow-lg">
    <h2 class="text-2xl font-bold mb-4 text-gray-700">404 Not Found</h2>
    <p class="text-gray-600">The page you are looking for does not exist.</p>
    <RouterLink to="/" class="text-blue-500 hover:underline">Go back to home</RouterLink>
  </div>
</template>