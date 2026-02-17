<template>
  <div class="min-h-screen bg-gray-50 flex flex-col items-center justify-center p-4">
    <div class="max-w-md w-full bg-white rounded-2xl shadow-xl p-8 text-center space-y-6">
      <div v-if="loading" class="space-y-4">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-red-600 mx-auto"></div>
        <p class="text-gray-600 font-medium">Verifying your email...</p>
      </div>

      <div v-else-if="success" class="space-y-6">
        <div class="bg-green-100 text-green-600 rounded-full w-16 h-16 flex items-center justify-center mx-auto">
          <span class="material-icons text-4xl">check_circle</span>
        </div>
        <h1 class="text-2xl font-bold text-gray-900">Email Verified!</h1>
        <p class="text-gray-600 text-sm">
          Your email has been successfully verified. You can now log in to your account.
        </p>
        <router-link to="/login" class="block w-full bg-red-600 hover:bg-red-700 text-white font-semibold py-3 px-4 rounded-lg transition-colors">
          Go to Login
        </router-link>
      </div>

      <div v-else class="space-y-6">
        <div class="bg-red-100 text-red-600 rounded-full w-16 h-16 flex items-center justify-center mx-auto">
          <span class="material-icons text-4xl">error_outline</span>
        </div>
        <h1 class="text-2xl font-bold text-gray-900">Verification Failed</h1>
        <p class="text-gray-600 text-sm">
          {{ error || 'The verification link is invalid or has expired.' }}
        </p>
        <router-link to="/login" class="block w-full bg-gray-100 hover:bg-gray-200 text-gray-700 font-semibold py-3 px-4 rounded-lg transition-colors">
          Back to Login
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { useAuthStore } from '@/stores/auth';

const route = useRoute();
const authStore = useAuthStore();

const loading = ref(true);
const success = ref(false);
const error = ref('');

onMounted(async () => {
  const token = route.query.token as string;
  if (!token) {
    loading.value = false;
    error.value = 'Verification token is missing.';
    return;
  }

  try {
    await authStore.verifyEmail(token);
    success.value = true;
  } catch (err: any) {
    error.value = err.response?.data?.error || 'Verification failed.';
  } finally {
    loading.value = false;
  }
});
</script>
