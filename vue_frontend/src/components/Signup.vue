<script setup>
import { ref } from 'vue';
import { RouterLink, useRoute } from 'vue-router';
// State for form inputs
const email = ref('');
const password = ref('');
const confirmPassword = ref('');
const error = ref('');

const isActiveLink = (routePath) => {
  const route = useRoute();
  return route.path === routePath;
};
// Handle form submission
const handleSubmit = () => {
  error.value = '';

  // Basic validation
  if (!email.value || !password.value || !confirmPassword.value) {
    error.value = 'All fields are required';
    return;
  }

  // Check password match
  if (password.value !== confirmPassword.value) {
    error.value = 'Passwords do not match';
    return;
  }

  // Example: Simulate signup process
  console.log('Signing up with:', { email: email.value, password: password.value });

  setTimeout(() => {
    alert('Signup successful! 🎉');
  }, 500);
};
</script>

<template>
  <div class="max-w-md mx-auto mt-10 bg-white p-6 rounded-lg shadow-lg">
    <h2 class="text-2xl font-bold mb-4 text-gray-700">Sign Up</h2>

    <!-- Error Message -->
    <div v-if="error" class="mb-4 text-red-500 text-sm">
      {{ error }}
    </div>

    <form @submit.prevent="handleSubmit">
      <!-- Email Input -->
      <div class="mb-4">
        <label for="email" class="block text-gray-600 mb-1">Email</label>
        <input 
          type="email" 
          id="email" 
          v-model="email" 
          class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-400"
          placeholder="Enter your email" 
          required
        />
      </div>

      <!-- Password Input -->
      <div class="mb-4">
        <label for="password" class="block text-gray-600 mb-1">Password</label>
        <input 
          type="password" 
          id="password" 
          v-model="password" 
          class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-400"
          placeholder="Create a password" 
          required
        />
      </div>

      <!-- Confirm Password Input -->
      <div class="mb-4">
        <label for="confirm-password" class="block text-gray-600 mb-1">Confirm Password</label>
        <input 
          type="password" 
          id="confirm-password" 
          v-model="confirmPassword" 
          class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-400"
          placeholder="Confirm your password" 
          required
        />
      </div>

      <!-- Submit Button -->
      <button 
        type="submit" 
        class="w-full bg-blue-500 text-white py-2 rounded-lg hover:bg-blue-600 transition duration-300"
      >
        Sign Up
      </button>
    </form>

    <!-- Link to Login -->
    <p class="mt-4 text-sm text-gray-500">
      Already have an account? 
      <RouterLink to="/login" class="text-blue-500 hover:underline">
        Log in
      </RouterLink>
    </p>
  </div>
</template>

<style scoped>
/* Optional extra styling */
</style>
