
<script>
import Navbar from './components/Navbar.vue';
import { RouterView } from 'vue-router';
export default {
  components: {
    Navbar
  }
};
</script>

<template>
 <Navbar />
  <RouterView />
</template>


