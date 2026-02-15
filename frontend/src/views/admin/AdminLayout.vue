<template>
  <div class="min-h-screen bg-gray-50 flex font-sans">
    <!-- Sidebar -->
    <aside 
        class="w-72 bg-white border-r border-gray-100 text-gray-600 fixed top-24 bottom-0 left-0 z-40 flex flex-col transition-transform duration-300 ease-in-out lg:translate-x-0" 
        :class="{'translate-x-0': sidebarOpen, '-translate-x-full': !sidebarOpen}"
    >
      <!-- Nav -->
      <nav class="flex-1 p-6 space-y-2 overflow-y-auto custom-scrollbar">
        <p class="px-4 text-xs font-bold text-gray-500 uppercase tracking-widest mb-4">Menu</p>
        
        <router-link to="/admin" class="flex items-center gap-4 px-4 py-3.5 text-gray-500 rounded-xl hover:bg-red-50 hover:text-[#FF3D3D] font-medium transition-all group" active-class="bg-red-50 text-[#FF3D3D] font-bold" exact>
          <span class="material-icons text-[22px] group-hover:scale-110 transition-transform">dashboard</span>
          Dashboard
        </router-link>
        
        <router-link to="/admin/donors" class="flex items-center gap-4 px-4 py-3.5 text-gray-500 rounded-xl hover:bg-red-50 hover:text-[#FF3D3D] font-medium transition-all group" active-class="bg-red-50 text-[#FF3D3D] font-bold">
          <span class="material-icons text-[22px] group-hover:scale-110 transition-transform">people_alt</span>
          Donors List
        </router-link>
        
        <router-link to="/admin/donations" class="flex items-center gap-4 px-4 py-3.5 text-gray-500 rounded-xl hover:bg-red-50 hover:text-[#FF3D3D] font-medium transition-all group" active-class="bg-red-50 text-[#FF3D3D] font-bold">
           <span class="material-icons text-[22px] group-hover:scale-110 transition-transform">volunteer_activism</span>
           Donations
        </router-link>
      </nav>
    </aside>

    <!-- Overlay -->
    <div v-if="sidebarOpen" @click="sidebarOpen = false" class="fixed inset-0 bg-black/60 backdrop-blur-sm z-30 lg:hidden transition-opacity"></div>

    <!-- Main Content -->
    <div class="flex-1 flex flex-col lg:ml-72 min-h-[calc(100vh-6rem)] transition-all duration-300">
      
      <!-- Mobile Sidebar Toggle (only visible on mobile) -->
       <div class="lg:hidden p-4 bg-white border-b border-gray-200 flex items-center justify-between">
            <h2 class="font-bold text-gray-800">Admin Menu</h2>
            <button @click="sidebarOpen = !sidebarOpen" class="text-gray-500 hover:text-[#FF3D3D] transition-colors p-2 rounded-lg hover:bg-gray-100">
                <span class="material-icons text-2xl">menu</span>
            </button>
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

const sidebarOpen = ref(false);
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
