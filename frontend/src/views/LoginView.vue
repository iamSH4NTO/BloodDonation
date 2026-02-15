<template>
  <div class="min-h-screen bg-[#faf5f5] flex flex-col font-sans relative">
    <!-- Top Navigation -->
    <div class="absolute top-0 left-0 w-full px-6 py-6 flex justify-between items-center">
      <div class="flex items-center gap-2">
         <!-- Logo Icon -->
        <div class="bg-red-600 p-1.5 rounded-lg shadow-sm">
            <span class="material-icons text-white text-lg">bloodtype</span>
        </div>
        <span class="text-xl font-bold text-gray-900 tracking-tight">BloodLink</span>
      </div>
      <a href="#" class="text-sm text-gray-500 hover:text-gray-900 font-medium">Help Center</a>
    </div>

    <!-- Main Content Area -->
    <main class="flex-grow flex items-center justify-center px-4 sm:px-6 lg:px-8 relative z-10">
      <div class="w-full max-w-[520px] bg-white rounded-2xl shadow-xl p-8 md:p-12 space-y-8">
        
        <!-- Header -->
        <div class="text-center space-y-2">
          <h1 class="text-2xl font-bold text-gray-900">
            Welcome Back, Donor
          </h1>
          <p class="text-gray-500 text-sm">
            Please enter your details to access your donation history.
          </p>
        </div>

        <form @submit.prevent="handleLogin" class="space-y-6">
            <!-- Email -->
            <div class="space-y-1.5">
            <label class="block text-sm font-semibold text-gray-700" for="email">Email Address</label>
            <div class="relative">
                <div class="absolute inset-y-0 left-0 flex items-center pl-3.5 pointer-events-none text-gray-400">
                <span class="material-icons text-lg">mail_outline</span>
                </div>
                <input v-model="email" autocomplete="email" 
                    class="block w-full rounded-lg border border-gray-200 bg-gray-50/50 focus:bg-white focus:ring-2 focus:ring-red-100 focus:border-red-500 outline-none transition-all py-3 pl-10 placeholder-gray-400 text-gray-800 sm:text-sm" 
                    id="email" name="email" placeholder="you@example.com" type="email" required />
            </div>
            </div>

            <!-- Password -->
            <div class="space-y-1.5">
            <label class="block text-sm font-semibold text-gray-700" for="password">Password</label>
            <div class="relative">
                <div class="absolute inset-y-0 left-0 flex items-center pl-3.5 pointer-events-none text-gray-400">
                <span class="material-icons text-lg">lock_outline</span>
                </div>
                <input v-model="password" :type="showPassword ? 'text' : 'password'" 
                    class="block w-full rounded-lg border border-gray-200 bg-gray-50/50 focus:bg-white focus:ring-2 focus:ring-red-100 focus:border-red-500 outline-none transition-all py-3 pl-10 pr-10 placeholder-gray-400 text-gray-800 sm:text-sm tracking-widest" 
                    id="password" name="password" placeholder="••••••••" required />
            </div>
            </div>

            <!-- Remember & Forgot -->
            <div class="flex items-center justify-between">
            <div class="flex items-center">
                <input id="remember-me" name="remember-me" type="checkbox" 
                    class="h-4 w-4 rounded border-gray-300 text-red-600 focus:ring-red-500 focus:ring-2 cursor-pointer" />
                <label for="remember-me" class="ml-2 block text-sm text-gray-600 cursor-pointer">Remember me</label>
            </div>
            <div class="text-sm">
                <a href="#" class="font-medium text-red-500 hover:text-red-600">Forgot password?</a>
            </div>
            </div>

            <!-- Sign In Button -->
            <button class="w-full bg-[#E53935] hover:bg-red-700 text-white font-semibold py-3 px-4 rounded-lg shadow-md hover:shadow-lg transition-all duration-300 flex items-center justify-center gap-2" type="submit">
            Sign In
            <span class="material-icons text-sm">arrow_forward</span>
            </button>
        </form>

        <!-- Divider -->
        <div class="relative">
            <div class="absolute inset-0 flex items-center">
                <div class="w-full border-t border-gray-200"></div>
            </div>
            <div class="relative flex justify-center text-sm">
                <span class="px-4 bg-white text-gray-500">New to BloodLink?</span>
            </div>
        </div>

        <!-- Register Link Button -->
        <router-link to="/register" class="w-full bg-white border border-gray-200 hover:bg-gray-50 text-gray-700 font-medium py-3 px-4 rounded-lg transition-all duration-200 flex items-center justify-center shadow-sm">
            Register for an account
        </router-link>

      </div>
    </main>

    <!-- Footer -->
    <footer class="py-6 text-center text-xs text-gray-400 bg-[#faf5f5]">
        <p>&copy; 2023 BloodLink Healthcare Systems. All rights reserved.</p>
        <div class="mt-2 space-x-4">
            <a href="#" class="hover:text-gray-600 transition-colors">Privacy Policy</a>
            <a href="#" class="hover:text-gray-600 transition-colors">Terms of Service</a>
        </div>
    </footer>

    <!-- Background Decoration (Optional subtle gradient based on image) -->
    <div class="absolute top-0 right-0 w-1/2 h-full bg-gradient-to-l from-red-50/50 to-transparent pointer-events-none"></div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router';

const email = ref('');
const password = ref('');
const showPassword = ref(false); // In design it's dots, but good to have toggle conceptually even if not visual in mock
const authStore = useAuthStore();
const router = useRouter();

const handleLogin = async () => {
  try {
    await authStore.login({ email: email.value, password: password.value });
    resetForm();
    router.push('/');
  } catch (error) {
    console.error(error);
    alert('Failed to login. Please check your credentials.');
  }
};

const resetForm = () => {
  email.value = '';
  password.value = '';
};
</script>
