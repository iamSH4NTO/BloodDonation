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
      <div class="space-y-2">
        <div v-for="log in logs" :key="log.id" class="bg-white rounded-xl p-3 shadow-sm border border-gray-100 hover:shadow-md transition-shadow">
          <div class="flex items-start gap-2 sm:gap-3">
            <!-- Viewer Avatar -->
            <div class="w-9 h-9 sm:w-10 sm:h-10 rounded-lg bg-blue-100 flex items-center justify-center text-blue-600 shrink-0">
              <span class="material-icons text-base sm:text-lg">person_search</span>
            </div>

            <!-- Log Details -->
            <div class="flex-1 min-w-0">
              <div class="flex flex-col sm:flex-row sm:items-center gap-1 sm:gap-2">
                <button @click="viewProfile(log.viewer_unique_id)" class="text-sm sm:text-base font-bold text-gray-900 hover:text-[#FF3D3D] transition-colors text-left truncate">
                  {{ log.viewer_name }}
                </button>
                <span class="text-xs text-gray-400">viewed</span>
                <button @click="viewProfile(log.target_unique_id)" class="text-sm sm:text-base font-bold text-gray-900 hover:text-[#FF3D3D] transition-colors text-left truncate">
                  {{ log.target_name }}
                </button>
              </div>

              <!-- Meta Info -->
              <div class="flex flex-wrap items-center gap-1.5 mt-1">
                <span class="text-[10px] sm:text-xs text-gray-400">{{ formatDate(log.created_at) }}</span>
                <span class="text-[10px] text-gray-300">•</span>
                <span class="text-[10px] sm:text-xs text-gray-400 font-mono">{{ log.ip_address }}</span>
              </div>
              
              <!-- Viewer Details -->
              <div class="mt-2 p-2 bg-blue-50 rounded-lg">
                <div class="text-[9px] sm:text-[10px] font-bold text-blue-900 uppercase mb-1">Viewer Info</div>
                <div class="grid grid-cols-1 sm:grid-cols-3 gap-1">
                  <div class="text-[10px] sm:text-xs text-blue-700 truncate">📞 {{ log.viewer_phone || 'N/A' }}</div>
                  <div class="text-[10px] sm:text-xs text-blue-700 truncate">✉️ {{ log.viewer_email || 'N/A' }}</div>
                  <div class="text-[10px] sm:text-xs text-blue-700 truncate">📍 {{ log.viewer_city || 'N/A' }}, {{ log.viewer_district || 'N/A' }}</div>
                </div>
              </div>
              
              <!-- Target Details -->
              <div class="mt-1.5 p-2 bg-gray-50 rounded-lg">
                <div class="text-[9px] sm:text-[10px] font-bold text-gray-900 uppercase mb-1">Target Info</div>
                <div class="grid grid-cols-1 sm:grid-cols-3 gap-1">
                  <div class="text-[10px] sm:text-xs text-gray-700 truncate">📞 {{ log.target_phone || 'N/A' }}</div>
                  <div class="text-[10px] sm:text-xs text-gray-700 truncate">✉️ {{ log.target_email || 'N/A' }}</div>
                  <div class="text-[10px] sm:text-xs text-gray-700 truncate">📍 {{ log.target_city || 'N/A' }}, {{ log.target_district || 'N/A' }}</div>
                </div>
              </div>
            </div>

            <!-- Arrow Icon -->
            <div class="shrink-0 text-gray-300">
              <span class="material-icons text-lg">arrow_forward</span>
            </div>
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
import api from '../../lib/axios';

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
  router.push(`/admin/donors/${userId}/edit`);
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
