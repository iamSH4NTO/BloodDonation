<script setup lang="ts">
import { computed } from 'vue';
import { RouterView, useRoute } from 'vue-router';
import Navbar from '@/components/Navbar.vue';
import Footer from '@/components/Footer.vue';
import ToastNotification from '@/components/ToastNotification.vue';

const route = useRoute();
const isAdminPage = computed(() => route.path.startsWith('/admin'));
</script>

<template>
  <div class="flex flex-col min-h-screen">
    <ToastNotification />
    <Navbar />
    <main class="grow flex flex-col">
      <RouterView />
    </main>
    <Footer v-if="!isAdminPage" />
  </div>
</template>
