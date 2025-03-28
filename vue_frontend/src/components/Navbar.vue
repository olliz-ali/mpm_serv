<script>
export default {
  data() {
    return {
      isMenuOpen: false,
    };
  },
  methods: {
    toggleMenu() {
      this.isMenuOpen = !this.isMenuOpen;
    },
  },
};
</script>

<template>
    <nav class="bg-white shadow-md">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <!-- Logo -->
          <div class="flex-shrink-0 flex items-center">
            <a href="/" class="text-2xl font-bold text-blue-600">MyApp</a>
          </div>
  
          <!-- Desktop Menu -->
          <div class="hidden md:flex space-x-8 items-center">
            <a href="/" class="text-gray-700 hover:text-blue-500">Home</a>
            <a href="/login" class="text-gray-700 hover:text-blue-500">Login</a>
            <a href="/signup" class="text-gray-700 hover:text-blue-500">Signup</a>
            <a href="/about" class="text-gray-700 hover:text-blue-500">About</a>
            <a href="/services" class="text-gray-700 hover:text-blue-500">Services</a>
            <a href="/contact" class="text-gray-700 hover:text-blue-500">Contact</a>
          </div>
  
          <!-- Mobile Menu Button -->
          <div class="-mr-2 flex md:hidden">
            <button @click="toggleMenu" class="inline-flex items-center justify-center p-2 rounded-md text-gray-700 hover:bg-gray-100 focus:outline-none">
              <svg class="h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path v-if="!isMenuOpen" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"/>
                <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
              </svg>
            </button>
          </div>
        </div>
      </div>
  
      <!-- Mobile Menu -->
      <div v-if="isMenuOpen" class="md:hidden">
        <div class="px-2 pt-2 pb-3 space-y-1 sm:px-3">
          <a href="/" class="block text-gray-700 hover:text-blue-500 py-2">Home</a>
          <a href="/login" class="block text-gray-700 hover:text-blue-500 py-2">Login</a>
          <a href="/signup" class="block text-gray-700 hover:text-blue-500 py-2">Signup</a>
          <a href="/about" class="block text-gray-700 hover:text-blue-500 py-2">About</a>
          <a href="/service" class="block text-gray-700 hover:text-blue-500 py-2">Services</a>
          <a href="/contact" class="block text-gray-700 hover:text-blue-500 py-2">Contact</a>
        </div>
      </div>
    </nav>
  </template>
  

  
  <style scoped>
  /* Optional extra styling */
  </style>
  