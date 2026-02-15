<template>
  <nav class="sticky top-0 z-50 border-b border-gray-100/50 backdrop-blur-lg bg-white/95">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center h-24">
        <!-- Logo -->
        <div class="shrink-0 flex items-center gap-2.5 cursor-pointer group" :class="$route.path.startsWith('/admin') ? 'lg:opacity-0 lg:pointer-events-none' : ''" @click="$router.push('/')">
           <div class="bg-[#FF3D3D] p-1.5 rounded-lg shadow-md group-hover:rotate-12 transition-transform duration-300">
            <span class="material-icons text-white text-2xl">bloodtype</span>
           </div>
           <span class="text-2xl font-bold text-[#1A1A1A] tracking-tight">BloodLink</span>
        </div>

        <!-- Desktop Menu -->
        <div class="hidden md:flex items-center space-x-10">
          <router-link to="/" class="text-sm font-bold text-gray-500 hover:text-[#FF3D3D] transition-colors relative group" exact-active-class="text-[#FF3D3D]">
            Home
            <span class="absolute -bottom-1 left-0 w-0 h-0.5 bg-[#FF3D3D] transition-all group-hover:w-full"></span>
          </router-link>
          <router-link to="/search" class="text-sm font-bold text-gray-500 hover:text-[#FF3D3D] transition-colors relative group" active-class="text-[#FF3D3D]">
            Find Donors
            <span class="absolute -bottom-1 left-0 w-0 h-0.5 bg-[#FF3D3D] transition-all group-hover:w-full"></span>
          </router-link>

          <router-link to="/about" class="text-sm font-bold text-gray-500 hover:text-[#FF3D3D] transition-colors relative group" active-class="text-[#FF3D3D]">
            About Us
            <span class="absolute -bottom-1 left-0 w-0 h-0.5 bg-[#FF3D3D] transition-all group-hover:w-full"></span>
          </router-link>
        </div>

        <!-- Right Side Actions -->
        <div class="hidden md:flex items-center space-x-6">
          <template v-if="authStore.isAuthenticated">
             <router-link v-if="authStore.user?.role === 'admin'" to="/admin" class="bg-[#1e1e2d] hover:bg-black text-white px-5 py-2.5 rounded-xl font-bold text-sm shadow-lg shadow-gray-500/20 hover:shadow-gray-500/30 transition-all transform hover:-translate-y-0.5 flex items-center gap-2">
                <span class="material-icons text-sm">dashboard</span> Admin
             </router-link>
             <router-link to="/profile" class="text-gray-600 hover:text-[#FF3D3D] font-bold text-sm">Profile</router-link>
             <button @click="logout" class="text-gray-400 hover:text-gray-600 font-bold text-sm">Logout</button>
          </template>
          <template v-else>
             <router-link to="/login" class="text-gray-600 hover:text-[#FF3D3D] font-bold text-sm">
              Login
             </router-link>
             <router-link to="/register" class="bg-[#FF3D3D] hover:bg-red-600 text-white px-6 py-3 rounded-full text-sm font-bold shadow-lg shadow-red-500/30 hover:shadow-red-500/40 transition-all transform hover:-translate-y-0.5">
              Register as Donor
            </router-link>
          </template>
        </div>

        <!-- Mobile Menu Button -->
        <div class="flex items-center md:hidden">
          <button @click="mobileMenuOpen = !mobileMenuOpen" class="text-gray-900 hover:text-[#FF3D3D] focus:outline-none p-2">
            <span class="material-icons text-3xl">{{ mobileMenuOpen ? 'close' : 'menu' }}</span>
          </button>
        </div>
      </div>
    </div>

    <!-- Mobile Menu -->
    <div v-show="mobileMenuOpen" class="md:hidden bg-white border-t border-gray-100 absolute w-full shadow-2xl z-50">
      <div class="px-6 pt-4 pb-8 space-y-4">
        <router-link to="/" class="block text-lg font-bold text-gray-900 hover:text-[#FF3D3D]" exact-active-class="text-[#FF3D3D]">Home</router-link>
        <router-link to="/search" class="block text-lg font-bold text-gray-900 hover:text-[#FF3D3D]" active-class="text-[#FF3D3D]">Find Donors</router-link>

        <router-link to="/about" class="block text-lg font-bold text-gray-900 hover:text-[#FF3D3D]" active-class="text-[#FF3D3D]">About Us</router-link>
        
        <div class="border-t border-gray-100 pt-6 mt-4">
           <template v-if="authStore.isAuthenticated">
              <router-link v-if="authStore.user?.role === 'admin'" to="/admin" class="text-lg font-bold text-[#1e1e2d] hover:text-[#FF3D3D] mb-4 flex items-center gap-2">
                 <span class="material-icons">dashboard</span> Admin Dashboard
              </router-link>
              <router-link to="/profile" class="block text-lg font-bold text-gray-900 hover:text-[#FF3D3D] mb-4">Profile</router-link>
              <button @click="logout" class="block w-full text-left text-lg font-bold text-gray-500 hover:text-gray-700">Logout</button>
           </template>
           <template v-else>
              <router-link to="/login" class="block text-lg font-bold text-gray-900 hover:text-[#FF3D3D] mb-4">
                Login
              </router-link>
              <router-link to="/register" class="block w-full text-center bg-[#FF3D3D] text-white px-6 py-4 rounded-xl text-lg font-bold shadow-lg">
                Register as Donor
              </router-link>
           </template>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router';

const authStore = useAuthStore();
const router = useRouter();
const mobileMenuOpen = ref(false);

const logout = () => {
  authStore.logout();
  router.push('/login');
  mobileMenuOpen.value = false;
};
</script>
