<script setup>
import { RouterLink,useRoute } from 'vue-router';
import { ref } from 'vue';

  const isMenuOpen = ref(false);
  const isActiveLink = (routePath) => {
    const route = useRoute();
    return route.path === routePath;
  }


    const toggleMenu = () => {
      isMenuOpen.value = !isMenuOpen.value;
    }

</script>

<template>
    <nav class="bg-white shadow-md">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <!-- Logo -->
          <div class="flex-shrink-0 flex items-center">
            <RouterLink to="/" class="text-2xl font-bold text-blue-600">MyApp</RouterLink>
          </div>
  
          <!-- Desktop Menu -->
          <div class="hidden md:flex space-x-8 items-center">
            <RouterLink to="/" :class="[
              isActiveLink('/') ? 'bg-blue-100' : 'hover:bg-gray-100',
              'text-gray-700', 'hover:text-blue-500']">Home</RouterLink>
            <RouterLink to="/login" :class="[
              isActiveLink('/login') ? 'bg-blue-100' : 'hover:bg-gray-100',
              'text-gray-700']">Login</RouterLink>
            <RouterLink to="/signup" :class="[
              isActiveLink('/signup') ? 'bg-blue-100' : 'hover:bg-gray-100',
              'text-gray-700']">Signup</RouterLink>
            <RouterLink to="/about" :class="[
              isActiveLink('/about') ? 'bg-blue-100' : 'hover:bg-gray-100',
              'text-gray-700']">About</RouterLink>
            <RouterLink to="/services" :class="[
              isActiveLink('/services') ? 'bg-blue-100' : 'hover:bg-gray-100',
              'text-gray-700']">Services</Routerlink>
            <RouterLink to="/contact" :class="[
              isActiveLink('/contact') ? 'bg-blue-100' : 'hover:bg-gray-100',
              'text-gray-700']">Contact</RouterLink>
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
          <RouterLink to="/" :class="[
            isActiveLink('/') ? 'bg-blue-100' : 'hover:bg-gray-100',
            'block text-gray-700', 'py-2']">Home</RouterLink>
          <RouterLink to="/login" :class="[
            isActiveLink('/login') ? 'bg-blue-100' : 'hover:bg-gray-100',
            'block text-gray-700', 'py-2']">Login</RouterLink>
          <RouterLink to="/signup" :class="[
            isActiveLink('/signup') ? 'bg-blue-100' : 'hover:bg-gray-100',
            'block text-gray-700', 'py-2']">Signup</RouterLink>
          <RouterLink to="/about" :class="[
            isActiveLink('/about') ? 'bg-blue-100' : 'hover:bg-gray-100',
            'block text-gray-700', 'py-2']">About</RouterLink>
          <RouterLink to="/service" :class="[
            isActiveLink('/service') ? 'bg-blue-100' : 'hover:bg-gray-100',
            'block text-gray-700', 'py-2']">Services</RouterLink>
          <RouterLink to="/contact" :class="[
            isActiveLink('/contact') ? 'bg-blue-100' : 'hover:bg-gray-100',
            'block text-gray-700', 'py-2']">Contact</RouterLink>
        </div>
      </div>
    </nav>
  </template>
  

  
  <style scoped>
  /* Optional extra styling */
  </style>
  