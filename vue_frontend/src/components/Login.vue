<script setup>
import { ref } from 'vue';

// State for email and password
const email = ref('');
const password = ref('');
const error = ref('');

// Handle form submission
const handleSubmit = () => {
  error.value = '';

  // Basic validation
  if (!email.value || !password.value) {
    error.value = 'Email and password are required';
    return;
  }

  // Example: Log in the user (replace this with actual logic)
  console.log('Logging in with:', { email: email.value, password: password.value });

  // Simulate login success or failure
  setTimeout(() => {
    if (email.value === 'test@example.com' && password.value === 'password') {
      alert('Login successful! 🚀');
    } else {
      error.value = 'Invalid email or password';
    }
  }, 500);
};
</script>

<template>
  <div class="max-w-md mx-auto mt-10 bg-white p-6 rounded-lg shadow-lg">
    <h2 class="text-2xl font-bold mb-4 text-gray-700">Login</h2>

    <!-- Error message -->
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
          placeholder="Enter your password" 
          required
        />
      </div>

      <!-- Submit Button -->
      <button 
        type="submit" 
        class="w-full bg-blue-500 text-white py-2 rounded-lg hover:bg-blue-600 transition duration-300"
      >
        Log In
      </button>
    </form>

    <!-- Link to Sign Up -->
    <p class="mt-4 text-sm text-gray-500">
      Don't have an account? 
      <router-link to="/signup" class="text-blue-500 hover:underline">
        Sign up
      </router-link>
    </p>
  </div>
</template>

<style scoped>
/* Optional extra styling */
</style>
