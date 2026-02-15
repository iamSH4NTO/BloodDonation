<template>
  <div class="space-y-8">
    
    <!-- Stats Grid -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      
      <!-- Total Donors -->
      <div class="relative overflow-hidden bg-white rounded-2xl shadow-sm border border-gray-100 p-6 group hover:shadow-lg transition-all duration-300">
        <div class="absolute top-0 right-0 w-32 h-32 bg-red-50 rounded-full blur-3xl -mr-16 -mt-16 opacity-50 group-hover:opacity-100 transition-opacity"></div>
        
        <div class="relative">
            <div class="flex justify-between items-start mb-4">
                <div class="w-12 h-12 rounded-xl bg-red-50 flex items-center justify-center text-[#FF3D3D] group-hover:scale-110 transition-transform duration-300">
                    <span class="material-icons text-2xl">people</span>
                </div>
                <span class="inline-flex items-center gap-1 text-[10px] font-bold px-2 py-1 rounded bg-green-50 text-green-600">
                    <span class="material-icons text-[10px]">arrow_upward</span> 12%
                </span>
            </div>
            <h3 class="text-3xl font-black text-gray-900 mb-1">{{ stats.total_donors }}</h3>
            <p class="text-sm font-medium text-gray-500">Registered Donors</p>
        </div>
      </div>

      <!-- Total Donations -->
      <div class="relative overflow-hidden bg-white rounded-2xl shadow-sm border border-gray-100 p-6 group hover:shadow-lg transition-all duration-300">
        <div class="absolute top-0 right-0 w-32 h-32 bg-orange-50 rounded-full blur-3xl -mr-16 -mt-16 opacity-50 group-hover:opacity-100 transition-opacity"></div>
        
        <div class="relative">
             <div class="flex justify-between items-start mb-4">
                <div class="w-12 h-12 rounded-xl bg-orange-50 flex items-center justify-center text-orange-500 group-hover:scale-110 transition-transform duration-300">
                    <span class="material-icons text-2xl">volunteer_activism</span>
                </div>
                <!-- <span class="inline-flex items-center gap-1 text-[10px] font-bold px-2 py-1 rounded bg-green-50 text-green-600">
                    <span class="material-icons text-[10px]">arrow_upward</span> 5%
                </span> -->
            </div>
            <h3 class="text-3xl font-black text-gray-900 mb-1">{{ stats.total_donations }}</h3>
            <p class="text-sm font-medium text-gray-500">Total Donations</p>
        </div>
      </div>

      <!-- Total Users -->
      <div class="relative overflow-hidden bg-white rounded-2xl shadow-sm border border-gray-100 p-6 group hover:shadow-lg transition-all duration-300">
        <div class="absolute top-0 right-0 w-32 h-32 bg-blue-50 rounded-full blur-3xl -mr-16 -mt-16 opacity-50 group-hover:opacity-100 transition-opacity"></div>
        
        <div class="relative">
             <div class="flex justify-between items-start mb-4">
                <div class="w-12 h-12 rounded-xl bg-blue-50 flex items-center justify-center text-blue-500 group-hover:scale-110 transition-transform duration-300">
                    <span class="material-icons text-2xl">group</span>
                </div>
            </div>
            <h3 class="text-3xl font-black text-gray-900 mb-1">{{ stats.total_users }}</h3>
            <p class="text-sm font-medium text-gray-500">Total Users</p>
        </div>
      </div>
    </div>
    
    <!-- Quick Actions / Placeholder -->
    <div class="bg-linear-to-r from-[#1e1e2d] to-[#2a2a3c] rounded-2xl p-8 text-white shadow-xl relative overflow-hidden">
        <div class="absolute right-0 top-0 h-full w-1/3 bg-linear-to-l from-white/5 to-transparent"></div>
        <div class="relative z-10 max-w-2xl">
            <h2 class="text-2xl font-bold mb-2">Welcome Back, Admin!</h2>
            <p class="text-gray-400 mb-6">You have standard access to manage donors, view donation history, and maintain the platform's integrity.</p>
            <div class="flex gap-4">
                <router-link to="/admin/donors" class="bg-[#FF3D3D] hover:bg-red-600 text-white px-6 py-3 rounded-xl font-bold text-sm shadow-lg shadow-red-500/30 transition-all flex items-center gap-2">
                    Manage Donors <span class="material-icons text-sm">arrow_forward</span>
                </router-link>
            </div>
        </div>
    </div>

  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import api from '@/lib/axios';

const stats = ref({
  total_donors: 0,
  total_donations: 0,
  total_users: 0
});

onMounted(async () => {
    try {
        const res = await api.get('/admin/stats');
        stats.value = res.data;
    } catch (error) {
        console.error("Failed to load stats", error);
    }
});
</script>
