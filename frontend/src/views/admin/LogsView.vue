<template>
  <div class="min-h-screen bg-[#FAFAFA] font-sans p-3 sm:p-4">
    <div class="max-w-7xl mx-auto space-y-3">
      
      <!-- Header -->
      <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-3">
        <div>
          <h1 class="text-xl sm:text-2xl font-black text-gray-900">Activity Logs</h1>
          <p class="text-xs sm:text-sm text-gray-500 mt-1">Monitor all profile views and contact activities</p>
        </div>
        <div class="flex items-center gap-2 text-xs sm:text-sm text-gray-500">
          <span class="material-icons text-sm">info</span>
          <span>Total: {{ totalLogs }} activities</span>
        </div>
      </div>

      <!-- Logs List -->
      <div class="space-y-1.5">
        <div v-for="log in logs" :key="log.id" class="bg-white rounded-lg p-2 sm:p-2.5 shadow-sm border border-gray-100 hover:shadow-md transition-shadow">
          <div class="flex items-center gap-2 text-[10px] sm:text-xs">
            <!-- Icon -->
            <div class="w-7 h-7 sm:w-8 sm:h-8 rounded-lg bg-blue-100 flex items-center justify-center text-blue-600 shrink-0">
              <span class="material-icons text-sm">visibility</span>
            </div>
            
            <!-- Viewer -->
            <button @click="viewProfile(log.viewer_unique_id)" class="font-bold text-gray-900 hover:text-[#FF3D3D] transition-colors truncate min-w-0">
              {{ log.viewer_name }}
            </button>
            
            <!-- Arrow -->
            <span class="material-icons text-gray-300 text-sm shrink-0">arrow_forward</span>
            
            <!-- Target -->
            <button @click="viewProfile(log.target_unique_id)" class="font-bold text-gray-900 hover:text-[#FF3D3D] transition-colors truncate min-w-0">
              {{ log.target_name }}
            </button>
            
            <!-- Divider -->
            <div class="hidden sm:block h-4 w-px bg-gray-200 shrink-0"></div>
            
            <!-- Viewer Contact (Desktop) -->
            <div class="hidden lg:flex items-center gap-1.5 text-blue-700 shrink-0">
              <span>📞 {{ log.viewer_phone || 'N/A' }}</span>
              <span class="text-gray-300">•</span>
              <span class="truncate max-w-[120px]">✉️ {{ log.viewer_email || 'N/A' }}</span>
            </div>
            
            <!-- Divider -->
            <div class="hidden lg:block h-4 w-px bg-gray-200 shrink-0"></div>
            
            <!-- Target Contact (Desktop) -->
            <div class="hidden lg:flex items-center gap-1.5 text-gray-700 shrink-0">
              <span>📞 {{ log.target_phone || 'N/A' }}</span>
              <span class="text-gray-300">•</span>
              <span class="truncate max-w-[120px]">✉️ {{ log.target_email || 'N/A' }}</span>
            </div>
            
            <!-- Spacer -->
            <div class="flex-1 min-w-0"></div>
            
            <!-- Time & IP -->
            <div class="flex items-center gap-1.5 text-gray-600 shrink-0">
              <span class="hidden sm:inline text-xs">{{ formatDate(log.created_at) }}</span>
              <span class="hidden sm:inline text-gray-300">•</span>
              <span class="font-mono text-xs font-semibold bg-gray-100 px-2 py-0.5 rounded">{{ log.ip_address || 'Unknown' }}</span>
            </div>
          </div>
          
          <!-- Mobile Contact Info (Collapsible) -->
          <div class="lg:hidden mt-1.5 pt-1.5 border-t border-gray-100 grid grid-cols-2 gap-1 text-[9px]">
            <div class="text-blue-700 truncate">👤 {{ log.viewer_phone }}</div>
            <div class="text-gray-700 truncate">🎯 {{ log.target_phone }}</div>
          </div>
        </div>

        <!-- Empty State -->
        <div v-if="logs.length === 0 && !loading" class="text-center py-12 bg-white rounded-2xl border-2 border-dashed border-gray-100">
          <span class="material-icons text-gray-300 text-4xl mb-2">event_busy</span>
          <div class="text-gray-400 text-sm">No activity logs found.</div>
        </div>

        <!-- Loading State -->
        <div v-if="loading" class="text-center py-12">
          <div class="inline-block animate-spin rounded-full h-8 w-8 border-4 border-gray-200 border-t-[#FF3D3D]"></div>
          <div class="text-gray-400 text-sm mt-2">Loading logs...</div>
        </div>
      </div>

      <!-- Pagination -->
      <div v-if="totalPages > 1" class="flex justify-center items-center gap-2 mt-6">
        <button @click="changePage(currentPage - 1)" :disabled="currentPage === 1" class="px-3 py-2 rounded-lg bg-white border border-gray-200 text-sm font-bold text-gray-700 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed">
          <span class="material-icons text-sm">chevron_left</span>
        </button>
        
        <div class="flex items-center gap-1">
          <button v-for="page in visiblePages" :key="page" @click="changePage(page)" :class="currentPage === page ? 'bg-[#FF3D3D] text-white' : 'bg-white text-gray-700 hover:bg-gray-50'" class="w-8 h-8 sm:w-10 sm:h-10 rounded-lg border border-gray-200 text-xs sm:text-sm font-bold transition-colors">
            {{ page }}
          </button>
        </div>

        <button @click="changePage(currentPage + 1)" :disabled="currentPage === totalPages" class="px-3 py-2 rounded-lg bg-white border border-gray-200 text-sm font-bold text-gray-700 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed">
          <span class="material-icons text-sm">chevron_right</span>
        </button>
      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRouter } from 'vue-router';
import api from '@/lib/axios';

const router = useRouter();

interface Log {
  id: number;
  viewer_id: string;
  viewer_name: string;
  viewer_unique_id: string;
  viewer_blood: string;
  viewer_phone: string;
  viewer_email: string;
  viewer_district: string;
  viewer_city: string;
  target_id: string;
  target_name: string;
  target_unique_id: string;
  target_blood: string;
  target_phone: string;
  target_email: string;
  target_district: string;
  target_city: string;
  created_at: string;
  ip_address: string;
}

const logs = ref<Log[]>([]);
const loading = ref(false);
const currentPage = ref(1);
const totalLogs = ref(0);
const totalPages = ref(0);
const limit = 20;

const fetchLogs = async (page: number = 1) => {
  loading.value = true;
  try {
    const res = await api.get(`/admin/logs?page=${page}&limit=${limit}`);
    logs.value = res.data.logs || [];
    totalLogs.value = res.data.total || 0;
    totalPages.value = res.data.pages || 0;
    currentPage.value = page;
  } catch (error) {
    console.error('Failed to fetch logs:', error);
  } finally {
    loading.value = false;
  }
};

const changePage = (page: number) => {
  if (page >= 1 && page <= totalPages.value) {
    fetchLogs(page);
  }
};

const visiblePages = computed(() => {
  const pages = [];
  const maxVisible = 5;
  let start = Math.max(1, currentPage.value - Math.floor(maxVisible / 2));
  let end = Math.min(totalPages.value, start + maxVisible - 1);
  
  if (end - start < maxVisible - 1) {
    start = Math.max(1, end - maxVisible + 1);
  }
  
  for (let i = start; i <= end; i++) {
    pages.push(i);
  }
  return pages;
});

const formatDate = (dateString: string) => {
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
  
  return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' });
};

const viewProfile = (userId: string) => {
  if (userId) {
    router.push(`/admin/donors/${userId}`);
  }
};

onMounted(() => fetchLogs(1));
</script>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.4s ease-out;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>
