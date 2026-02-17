<script setup lang="ts">
import { RouterView } from 'vue-router';
import Navbar from '@/components/Navbar.vue';
import Footer from '@/components/Footer.vue';
import ToastNotification from '@/components/ToastNotification.vue';
</script>

<template>
  <div class="flex flex-col min-h-screen">
    <ToastNotification />
    <Navbar />
    <main class="grow flex flex-col">
      <RouterView />
    </main>
    <Footer />
  </div>
</template>
