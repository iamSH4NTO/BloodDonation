<template>
  <div class="min-h-screen bg-gray-50 flex font-sans">
    <!-- Sidebar -->
    <!-- Sidebar Drawer (Mobile) & Persistent Sidebar (Desktop) -->
    <aside 
        class="w-72 backdrop-blur-lg bg-white/95 border-r border-gray-100/50 text-gray-600 fixed top-0 bottom-0 left-0 z-50 flex flex-col transition-transform duration-300 ease-in-out lg:translate-x-0" 
        :class="{'translate-x-0': isSidebarOpen, '-translate-x-full': !isSidebarOpen}"
    >
      <!-- Logo Section (matches navbar height) -->
      <div class="h-24 flex items-center px-4 sm:px-6 lg:px-8 border-b border-gray-100/50">
        <div class="flex items-center gap-2.5 cursor-pointer group" @click="$router.push('/')">
          <div class="bg-[#FF3D3D] p-1.5 rounded-lg shadow-md group-hover:rotate-12 transition-transform duration-300">
            <span class="material-icons text-white text-2xl">bloodtype</span>
          </div>
          <span class="text-2xl font-bold text-[#1A1A1A] tracking-tight">BloodLink</span>
        </div>
      </div>

      <div class="space-y-10 px-4 py-6 flex-1 overflow-y-auto">
        <!-- Navigation Groups -->
        <div class="space-y-4">
          <p class="px-4 text-[10px] font-black uppercase tracking-[0.2em] text-gray-400">Menu</p>
          <nav class="space-y-1">
            <router-link to="/admin" class="flex items-center gap-4 px-4 py-3 rounded-xl transition-all group" :class="[$route.path === '/admin' ? 'bg-[#FF3D3D] text-white shadow-lg shadow-red-500/20' : 'text-gray-500 hover:text-[#FF3D3D] hover:bg-gray-50']" @click="isSidebarOpen = false">
              <div class="flex items-center justify-center w-6 h-6">
                <span class="material-icons text-[20px]">dashboard</span>
              </div>
              <span class="text-sm font-bold tracking-wide">Dashboard</span>
              <span v-if="$route.path === '/admin'" class="ml-auto w-1.5 h-1.5 rounded-full bg-white animate-pulse"></span>
            </router-link>
            
            <router-link to="/admin/donors" class="flex items-center gap-4 px-4 py-3 rounded-xl transition-all group" :class="[$route.path === '/admin/donors' ? 'bg-[#FF3D3D] text-white shadow-lg shadow-red-500/20' : 'text-gray-500 hover:text-[#FF3D3D] hover:bg-gray-50']" @click="isSidebarOpen = false">
              <div class="flex items-center justify-center w-6 h-6">
                <span class="material-icons text-[20px]">people_alt</span>
              </div>
              <span class="text-sm font-bold tracking-wide">Donors List</span>
              <span v-if="$route.path === '/admin/donors'" class="ml-auto w-1.5 h-1.5 rounded-full bg-white animate-pulse"></span>
            </router-link>
          </nav>
        </div>
      </div>
    </aside>

    <!-- Overlay for Mobile -->
    <div 
      v-if="isSidebarOpen" 
      @click="isSidebarOpen = false" 
      class="fixed inset-0 bg-black/60 backdrop-blur-sm z-20 lg:hidden transition-opacity duration-300"
    ></div>

    <!-- Main Content -->
    <div class="flex-1 flex flex-col lg:ml-72 min-h-[calc(100vh-6rem)] transition-all duration-300">
      
      <!-- Mobile Sidebar Toggle (only visible on mobile) -->
       <div class="lg:hidden p-4 bg-white border-b border-gray-200 flex items-center justify-between">
            <!-- Left: Brand/Logo (Mobile toggle) -->
            <div class="flex items-center gap-4">
                <button @click="isSidebarOpen = !isSidebarOpen" class="lg:hidden p-2 rounded-lg hover:bg-gray-100 text-gray-500 transition-colors">
                    <span class="material-icons text-2xl">{{ isSidebarOpen ? 'close' : 'menu' }}</span>
                </button>
                <h2 class="font-bold text-gray-800">Admin Menu</h2>
            </div>
       </div>
      
      <!-- Content Area -->
      <main class="flex-1 p-6 sm:p-8 lg:p-10 overflow-x-hidden">
        <div class="max-w-7xl mx-auto">
            <router-view v-slot="{ Component }">
                <transition name="fade" mode="out-in">
                    <component :is="Component" />
                </transition>
            </router-view>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router';

const isSidebarOpen = ref(false);

const closeSidebarOnMobile = () => {
  if (window.innerWidth < 1024) {
    isSidebarOpen.value = false;
  }
};
const authStore = useAuthStore();
const router = useRouter();

const logout = () => {
  authStore.logout();
  router.push('/login');
};
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: rgba(0, 0, 0, 0.1);
  border-radius: 4px;
}
.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: rgba(0, 0, 0, 0.2);
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
