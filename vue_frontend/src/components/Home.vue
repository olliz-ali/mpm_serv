<script setup>
import { ref } from 'vue';
</script>


<template>
    <div class="min-h-screen flex items-center justify-center bg-gray-100">
      <div class="text-center">
        <h1 class="text-4xl font-bold text-blue-600 mb-4">
          Welcome to MyApp 🚀
        </h1>

        <router-link to="/about">
          <button
            class="px-6 py-2 bg-blue-500 text-white rounded-lg shadow-md hover:bg-blue-600 transition duration-300"
          >
            Learn More
          </button>
        </router-link>
      </div>
    </div>
  </template>
  
