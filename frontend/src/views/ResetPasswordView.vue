<template>
  <div class="min-h-screen bg-[#faf5f5] flex items-center justify-center p-4 font-sans">
    <div class="max-w-md w-full bg-white rounded-2xl shadow-xl p-8 space-y-6">
      <div class="text-center space-y-2">
        <h1 class="text-2xl font-black text-gray-900 tracking-tight">Set New Password</h1>
        <p class="text-gray-500 text-sm font-medium">Please enter your new password below.</p>
      </div>

      <form v-if="!success" @submit.prevent="handleSubmit" class="space-y-6">
        <!-- New Password -->
        <div class="space-y-1.5">
          <label class="block text-sm font-semibold text-gray-700" for="password">New Password</label>
          <div class="relative">
            <div class="absolute inset-y-0 left-0 flex items-center pl-3.5 pointer-events-none text-gray-400">
              <span class="material-icons text-lg">lock_outline</span>
            </div>
            <input v-model="password" :type="showPassword ? 'text' : 'password'" id="password" required placeholder="••••••••"
              class="block w-full rounded-lg border border-gray-200 bg-gray-50/50 focus:bg-white focus:ring-2 focus:ring-red-100 focus:border-red-500 outline-none transition-all py-2.5 pl-10 pr-10 placeholder-gray-400 text-gray-800 sm:text-sm" />
            <button type="button" @click="showPassword = !showPassword" class="absolute inset-y-0 right-0 flex items-center px-3 text-gray-400 hover:text-gray-600 focus:outline-none transition-colors">
              <span class="material-icons text-lg">{{ showPassword ? 'visibility' : 'visibility_off' }}</span>
            </button>
          </div>
        </div>

        <!-- Confirm Password -->
        <div class="space-y-1.5">
          <label class="block text-sm font-semibold text-gray-700" for="confirmPassword">Confirm Password</label>
          <div class="relative">
            <div class="absolute inset-y-0 left-0 flex items-center pl-3.5 pointer-events-none text-gray-400">
              <span class="material-icons text-lg">lock_clock</span>
            </div>
            <input v-model="confirmPassword" :type="showConfirmPassword ? 'text' : 'password'" id="confirmPassword" required placeholder="••••••••"
              class="block w-full rounded-lg border border-gray-200 bg-gray-50/50 focus:bg-white focus:ring-2 focus:ring-red-100 focus:border-red-500 outline-none transition-all py-2.5 pl-10 pr-10 placeholder-gray-400 text-gray-800 sm:text-sm" />
            <button type="button" @click="showConfirmPassword = !showConfirmPassword" class="absolute inset-y-0 right-0 flex items-center px-3 text-gray-400 hover:text-gray-600 focus:outline-none transition-colors">
              <span class="material-icons text-lg">{{ showConfirmPassword ? 'visibility' : 'visibility_off' }}</span>
            </button>
          </div>
        </div>

        <button :disabled="loading" class="w-full bg-[#E53935] hover:bg-red-700 text-white font-semibold py-3 px-4 rounded-lg shadow-md hover:shadow-lg transition-all duration-300 flex items-center justify-center gap-2 disabled:opacity-50" type="submit">
          <span v-if="loading" class="animate-spin rounded-full h-5 w-5 border-b-2 border-white"></span>
          {{ loading ? 'Updating...' : 'Reset Password' }}
          <span v-if="!loading" class="material-icons text-sm">vps</span>
        </button>
      </form>

      <div v-else class="text-center space-y-6 animate-in fade-in zoom-in duration-300">
        <div class="bg-green-100 text-green-600 rounded-full w-16 h-16 flex items-center justify-center mx-auto">
          <span class="material-icons text-4xl">verified</span>
        </div>
        <div class="space-y-2">
          <p class="text-gray-800 font-bold">Password Updated!</p>
          <p class="text-gray-500 text-sm">Your password has been reset successfully. You can now use your new password to log in.</p>
        </div>
        <router-link to="/login" class="block w-full bg-red-600 hover:bg-red-700 text-white font-semibold py-3 px-4 rounded-lg transition-colors">
          Go to Login
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import { useAuthStore } from '@/stores/auth';
import { useToastStore } from '@/stores/toast';

const route = useRoute();
const authStore = useAuthStore();
const toastStore = useToastStore();

const password = ref('');
const confirmPassword = ref('');
const showPassword = ref(false);
const showConfirmPassword = ref(false);
const loading = ref(false);
const success = ref(false);

const handleSubmit = async () => {
  if (password.value !== confirmPassword.value) {
    toastStore.show("Passwords don't match!", 'error');
    return;
  }

  const token = route.query.token as string;
  if (!token) {
    toastStore.show('Reset token is missing.', 'error');
    return;
  }

  loading.value = true;
  try {
    await authStore.resetPassword({
      token: token,
      password: password.value
    });
    success.value = true;
    toastStore.show('Password updated successfully', 'success');
  } catch (error: any) {
    console.error(error);
    toastStore.show(error.response?.data?.error || 'Failed to update password.', 'error');
  } finally {
    loading.value = false;
  }
};
</script>
