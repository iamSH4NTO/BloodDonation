<template>
  <div class="min-h-screen bg-[#faf5f5] flex items-center justify-center p-4 font-sans">
    <div class="max-w-md w-full bg-white rounded-2xl shadow-xl p-8 space-y-6">
      <div class="text-center space-y-2">
        <h1 class="text-2xl font-black text-gray-900 tracking-tight">Forgot Password</h1>
        <p class="text-gray-500 text-sm font-medium">Enter your email and we'll send you a link to reset your password.</p>
      </div>

      <form v-if="!submitted" @submit.prevent="handleSubmit" class="space-y-6">
        <div class="space-y-1.5">
          <label class="block text-sm font-semibold text-gray-700" for="email">Email Address</label>
          <div class="relative">
            <div class="absolute inset-y-0 left-0 flex items-center pl-3.5 pointer-events-none text-gray-400">
              <span class="material-icons text-lg">mail_outline</span>
            </div>
            <input v-model="email" type="email" id="email" required placeholder="you@example.com"
              class="block w-full rounded-lg border border-gray-200 bg-gray-50/50 focus:bg-white focus:ring-2 focus:ring-red-100 focus:border-red-500 outline-none transition-all py-2.5 pl-10 placeholder-gray-400 text-gray-800 sm:text-sm" />
          </div>
        </div>

        <button :disabled="loading" class="w-full bg-[#E53935] hover:bg-red-700 text-white font-semibold py-3 px-4 rounded-lg shadow-md hover:shadow-lg transition-all duration-300 flex items-center justify-center gap-2 disabled:opacity-50" type="submit">
          <span v-if="loading" class="animate-spin rounded-full h-5 w-5 border-b-2 border-white"></span>
          {{ loading ? 'Sending...' : 'Send Reset Link' }}
          <span v-if="!loading" class="material-icons text-sm">arrow_forward</span>
        </button>
      </form>

      <div v-else class="text-center space-y-6 animate-in fade-in zoom-in duration-300">
        <div class="bg-green-100 text-green-600 rounded-full w-16 h-16 flex items-center justify-center mx-auto">
          <span class="material-icons text-4xl">mark_email_read</span>
        </div>
        <div class="space-y-2">
          <p class="text-gray-800 font-bold">Check your email</p>
          <p class="text-gray-500 text-sm">We've sent a password reset link to <span class="font-semibold text-gray-700">{{ email }}</span></p>
        </div>
        <router-link to="/login" class="inline-block text-sm font-medium text-red-500 hover:text-red-600">
          Return to login
        </router-link>
      </div>

      <div v-if="!submitted" class="text-center">
        <router-link to="/login" class="text-sm font-medium text-gray-500 hover:text-gray-700 flex items-center justify-center gap-1">
          <span class="material-icons text-base">arrow_back</span>
          Back to login
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '@/stores/auth';

const email = ref('');
const loading = ref(false);
const submitted = ref(false);
const authStore = useAuthStore();

const handleSubmit = async () => {
  loading.value = true;
  try {
    await authStore.forgotPassword(email.value);
    submitted.value = true;
  } catch (error: any) {
    console.error(error);
    alert('Something went wrong. Please try again.');
  } finally {
    loading.value = false;
  }
};
</script>
