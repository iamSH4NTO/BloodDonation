<script setup lang="ts">
import { RouterView, useRoute } from 'vue-router';
import Navbar from '@/components/Navbar.vue';
import Footer from '@/components/Footer.vue';
import { computed } from 'vue';

const route = useRoute();
const hideNavFooter = computed(() => {
  return ['/login', '/register'].includes(route.path);
});
</script>

<template>
  <div class="flex flex-col min-h-screen">
    <Navbar v-if="!hideNavFooter" />
    <main class="flex-grow">
      <RouterView />
    </main>
    <Footer v-if="!hideNavFooter" />
  </div>
</template>
