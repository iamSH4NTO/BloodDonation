<template>
  <div class="space-y-8">
    
    <!-- Stats Grid -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 sm:gap-6">
      
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
    
    <!-- Recent Activity Logs -->
    <div class="bg-white rounded-2xl shadow-sm border border-gray-100 p-4 sm:p-6">
      <div class="flex justify-between items-center mb-4">
        <h2 class="text-base sm:text-lg font-bold text-gray-900">Recent Activity</h2>
        <router-link to="/admin/logs" class="text-xs sm:text-sm font-bold text-[#FF3D3D] hover:text-red-600 transition-colors flex items-center gap-1">
          View All <span class="material-icons text-sm">arrow_forward</span>
        </router-link>
      </div>
      
      <div class="space-y-3">
        <div v-for="log in recentLogs" :key="log.id" class="flex items-center gap-3 p-3 bg-gray-50 rounded-xl hover:bg-gray-100 transition-colors">
          <div class="w-8 h-8 sm:w-10 sm:h-10 rounded-lg bg-blue-100 flex items-center justify-center text-blue-600 shrink-0">
            <span class="material-icons text-sm sm:text-base">person_search</span>
          </div>
          <div class="flex-1 min-w-0">
            <div class="flex items-center gap-1 text-xs sm:text-sm">
              <button @click="viewProfile(log.viewer_unique_id)" class="font-bold text-gray-900 hover:text-[#FF3D3D] transition-colors truncate">{{ log.viewer_name }}</button>
              <span class="text-gray-400 shrink-0">viewed</span>
              <button @click="viewProfile(log.target_unique_id)" class="font-bold text-gray-900 hover:text-[#FF3D3D] transition-colors truncate">{{ log.target_name }}</button>
            </div>
            <div class="text-[10px] sm:text-xs text-gray-400 mt-0.5">{{ formatTimeAgo(log.created_at) }}</div>
          </div>
        </div>
        
        <div v-if="recentLogs.length === 0" class="text-center py-6 text-gray-400 text-sm">
          No recent activity
        </div>
      </div>
    </div>
    
    <!-- Quick Actions / Placeholder -->
    <div class="bg-linear-to-r from-[#1e1e2d] to-[#2a2a3c] rounded-4xl p-8 sm:p-10 text-white shadow-xl relative overflow-hidden">
        <div class="absolute right-0 top-0 h-full w-1/3 bg-linear-to-l from-white/5 to-transparent pointer-events-none"></div>
        <div class="relative z-10 max-w-2xl text-center sm:text-left">
            <h2 class="text-2xl sm:text-3xl font-black mb-3 tracking-tight">Welcome Back, Admin!</h2>
            <p class="text-gray-400 text-sm sm:text-base mb-8 leading-relaxed font-medium">You have standard access to manage donors, view donation history, and maintain the platform's integrity.</p>
            <div class="flex flex-col sm:flex-row gap-4 items-center sm:items-start">
                <router-link to="/admin/donors" class="w-full sm:w-auto bg-[#FF3D3D] hover:bg-red-600 text-white px-8 py-3.5 rounded-xl font-bold text-sm shadow-lg shadow-red-500/30 transition-all flex items-center justify-center gap-2 transform hover:-translate-y-1">
                    Manage Donors <span class="material-icons text-sm">arrow_forward</span>
                </router-link>
            </div>
        </div>
    </div>

  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import api from '@/lib/axios';

const router = useRouter();

const stats = ref({
  total_donors: 0,
  total_donations: 0,
  total_users: 0
});

interface RecentLog {
  id: number;
  viewer_name: string;
  viewer_unique_id: string;
  target_name: string;
  target_unique_id: string;
  created_at: string;
}

const recentLogs = ref<RecentLog[]>([]);

const fetchStats = async () => {
  try {
    const res = await api.get('/admin/stats');
    stats.value = res.data;
  } catch (error) {
    console.error("Failed to load stats", error);
  }
};

const fetchRecentLogs = async () => {
  try {
    const res = await api.get('/admin/logs/recent');
    recentLogs.value = (res.data || []).slice(0, 5);
  } catch (error) {
    console.error("Failed to load recent logs", error);
  }
};

const formatTimeAgo = (dateString: string) => {
  const date = new Date(dateString);
  const now = new Date();
  const diff = now.getTime() - date.getTime();
  const minutes = Math.floor(diff / 60000);
  const hours = Math.floor(diff / 3600000);
  const days = Math.floor(diff / 86400000);
  
  if (minutes < 1) return 'Just now';
  if (minutes < 60) return `${minutes}m ago`;
  if (hours < 24) return `${hours}h ago`;
  if (days < 7) return `${days}d ago`;
  
  return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric' });
};

const viewProfile = (userId: string) => {
  router.push(`/admin/donors/${userId}/edit`);
};

onMounted(() => {
  fetchStats();
  fetchRecentLogs();
});
</script>
