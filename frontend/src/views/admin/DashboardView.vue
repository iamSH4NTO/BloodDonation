<template>
  <div class="space-y-6 lg:space-y-8 max-w-7xl mx-auto pb-10">
    
    <!-- Header Section -->
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-end gap-4 bg-gray-900/50 backdrop-blur-xl p-6 rounded-3xl shadow-[0_8px_32px_rgba(0,0,0,0.12)] border border-gray-800/50">
      <div>
        <h1 class="text-2xl sm:text-3xl font-black text-white tracking-tight flex items-center gap-3">
          <span class="w-2 h-8 bg-[#00F0FF] rounded-full shadow-[0_0_15px_rgba(0,240,255,0.7)]"></span>
          Command Center
        </h1>
        <p class="text-sm font-medium text-gray-400 mt-1">Real-time blood bank telemetry & user engagement</p>
      </div>
      <div class="flex items-center gap-3">
        <!-- Date Range Filter -->
        <select 
          v-model="dateRange" 
          class="bg-gray-800 border border-gray-700 text-gray-300 text-sm rounded-xl focus:ring-[#00F0FF] focus:border-[#00F0FF] block p-2 font-medium cursor-pointer shadow-sm hover:bg-gray-700 transition-colors outline-none"
        >
          <option value="7">Last 7 Days</option>
          <option value="30">Last 30 Days</option>
          <option value="90">Last 3 Months</option>
          <option value="180">Last 6 Months</option>
          <option value="365">Last 1 Year</option>
          <option value="1825">All Time</option>
        </select>

        <span class="hidden sm:inline-flex items-center gap-1.5 px-3 py-1.5 rounded-lg bg-[#00F0FF]/10 text-[#00F0FF] text-xs font-bold border border-[#00F0FF]/20 shadow-[0_0_10px_rgba(0,240,255,0.2)]">
          <span class="w-2 h-2 rounded-full bg-[#00F0FF] animate-pulse shadow-[0_0_8px_rgba(0,240,255,1)]"></span>
          Live Sync
        </span>
        <button @click="fetchData" class="w-10 h-10 shrink-0 rouned-xl bg-gray-800 hover:bg-gray-700 flex items-center justify-center text-gray-400 hover:text-[#00F0FF] hover:shadow-[0_0_15px_rgba(0,240,255,0.3)] transition-all border border-gray-700 rounded-xl" title="Refresh">
          <span class="material-icons text-sm" :class="{'animate-spin text-[#00F0FF]': isLoading}">refresh</span>
        </button>
      </div>
    </div>

    <!-- Quick Stats Grid -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 sm:gap-6">
      
      <!-- Total Donors -->
      <div class="relative overflow-hidden bg-gray-900/40 backdrop-blur-xl rounded-3xl shadow-[0_8px_32px_rgba(0,0,0,0.2)] border border-[#00F0FF]/10 p-6 group hover:shadow-[0_0_30px_rgba(0,240,255,0.15)] transition-all duration-300 transform hover:-translate-y-1">
        <div class="absolute -right-6 -top-6 w-32 h-32 bg-[#00F0FF]/5 rounded-full blur-3xl group-hover:bg-[#00F0FF]/15 transition-colors duration-500"></div>
        <div class="relative">
            <div class="flex justify-between items-start mb-6">
                <div class="w-12 h-12 rounded-2xl bg-gray-800/80 flex items-center justify-center text-[#00F0FF] border border-[#00F0FF]/20 group-hover:scale-110 group-hover:rotate-3 transition-transform duration-300 shadow-[0_0_15px_rgba(0,240,255,0.1)]">
                    <span class="material-icons text-xl">people</span>
                </div>
                <!-- Growth Indicator -->
                <span v-if="stats.growth?.donors" class="inline-flex items-center gap-1 text-[11px] font-black px-2 py-1 rounded-lg" :class="parseFloat(stats.growth.donors) >= 0 ? 'bg-[#00FF66]/10 text-[#00FF66] border border-[#00FF66]/20' : 'bg-[#FF003C]/10 text-[#FF003C] border border-[#FF003C]/20'">
                    <span class="material-icons text-[12px]">{{ parseFloat(stats.growth.donors) >= 0 ? 'trending_up' : 'trending_down' }}</span> 
                    {{ Math.abs(parseFloat(stats.growth.donors)) }}%
                </span>
            </div>
            <div>
              <p class="text-xs font-bold text-gray-400 uppercase tracking-widest mb-1 font-mono">Registered Donors</p>
              <h3 class="text-4xl font-black text-white tracking-tight drop-shadow-[0_0_8px_rgba(255,255,255,0.3)]">{{ stats.total_donors }}</h3>
            </div>
        </div>
      </div>

      <!-- Total Donations -->
      <div class="relative overflow-hidden bg-gray-900/40 backdrop-blur-xl rounded-3xl shadow-[0_8px_32px_rgba(0,0,0,0.2)] border border-[#FF00FF]/10 p-6 group hover:shadow-[0_0_30px_rgba(255,0,255,0.15)] transition-all duration-300 transform hover:-translate-y-1">
        <div class="absolute -right-6 -top-6 w-32 h-32 bg-[#FF00FF]/5 rounded-full blur-3xl group-hover:bg-[#FF00FF]/15 transition-colors duration-500"></div>
        <div class="relative">
             <div class="flex justify-between items-start mb-6">
                <div class="w-12 h-12 rounded-2xl bg-gray-800/80 flex items-center justify-center text-[#FF00FF] border border-[#FF00FF]/20 group-hover:scale-110 group-hover:-rotate-3 transition-transform duration-300 shadow-[0_0_15px_rgba(255,0,255,0.1)]">
                    <span class="material-icons text-xl">volunteer_activism</span>
                </div>
                <!-- Growth Indicator -->
                <span v-if="stats.growth?.donations" class="inline-flex items-center gap-1 text-[11px] font-black px-2 py-1 rounded-lg" :class="parseFloat(stats.growth.donations) >= 0 ? 'bg-[#00FF66]/10 text-[#00FF66] border border-[#00FF66]/20' : 'bg-[#FF003C]/10 text-[#FF003C] border border-[#FF003C]/20'">
                    <span class="material-icons text-[12px]">{{ parseFloat(stats.growth.donations) >= 0 ? 'trending_up' : 'trending_down' }}</span> 
                    {{ Math.abs(parseFloat(stats.growth.donations)) }}%
                </span>
            </div>
            <div>
              <p class="text-xs font-bold text-gray-400 uppercase tracking-widest mb-1 font-mono">Total Donations</p>
              <h3 class="text-4xl font-black text-white tracking-tight drop-shadow-[0_0_8px_rgba(255,255,255,0.3)]">{{ stats.total_donations }}</h3>
            </div>
        </div>
      </div>

       <!-- Lives Saved Estimate -->
      <div class="relative overflow-hidden bg-gray-900/40 backdrop-blur-xl rounded-3xl shadow-[0_8px_32px_rgba(0,0,0,0.2)] border border-[#FF003C]/10 p-6 group hover:shadow-[0_0_30px_rgba(255,0,60,0.15)] transition-all duration-300 transform hover:-translate-y-1">
        <div class="absolute -right-6 -top-6 w-32 h-32 bg-[#FF003C]/5 rounded-full blur-3xl group-hover:bg-[#FF003C]/20 transition-colors duration-500"></div>
        <div class="relative">
             <div class="flex justify-between items-start mb-6">
                <div class="w-12 h-12 rounded-2xl bg-gray-800/80 flex items-center justify-center text-[#FF003C] border border-[#FF003C]/30 group-hover:scale-110 group-hover:rotate-12 transition-transform duration-300 shadow-[0_0_15px_rgba(255,0,60,0.2)]">
                    <span class="material-icons text-xl">favorite</span>
                </div>
            </div>
            <div>
              <p class="text-xs font-bold text-gray-400 uppercase tracking-widest mb-1 font-mono">Est. Lives Saved</p>
              <h3 class="text-4xl font-black text-[#FF003C] tracking-tight drop-shadow-[0_0_10px_rgba(255,0,60,0.4)]">{{ stats.total_donations * 3 }}</h3>
            </div>
        </div>
      </div>

      <!-- Total Users -->
      <div class="relative overflow-hidden bg-linear-to-br from-[#1E293B] to-[#0F172A] rounded-3xl shadow-[0_8px_32px_rgba(0,0,0,0.3)] border border-gray-700/50 p-6 group hover:shadow-[0_0_30px_rgba(255,255,255,0.05)] transition-all duration-300 transform hover:-translate-y-1">
        <div class="absolute right-0 top-0 h-full w-1/2 bg-linear-to-l from-white/5 to-transparent pointer-events-none rounded-r-3xl"></div>
        <div class="relative z-10">
             <div class="flex justify-between items-start mb-6">
                <div class="w-12 h-12 rounded-2xl bg-white/5 flex items-center justify-center text-white border border-white/10 backdrop-blur-md group-hover:scale-110 transition-transform duration-300">
                    <span class="material-icons text-xl">account_circle</span>
                </div>
                <!-- Growth Indicator -->
                <span v-if="stats.growth?.users" class="inline-flex items-center gap-1 text-[11px] font-black px-2 py-1 rounded-lg" :class="parseFloat(stats.growth.users) >= 0 ? 'bg-[#00FF66]/10 text-[#00FF66] border border-[#00FF66]/20' : 'bg-[#FF003C]/10 text-[#FF003C] border border-[#FF003C]/20'">
                    <span class="material-icons text-[12px]">{{ parseFloat(stats.growth.users) >= 0 ? 'trending_up' : 'trending_down' }}</span> 
                    {{ Math.abs(parseFloat(stats.growth.users)) }}%
                </span>
            </div>
            <div>
              <p class="text-xs font-bold text-gray-400 uppercase tracking-widest mb-1 font-mono">System Accounts</p>
              <h3 class="text-4xl font-black text-white tracking-tight">{{ stats.total_users }}</h3>
            </div>
        </div>
      </div>
    </div>

    <!-- Charts Section -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        
        <!-- Main Line Chart: Registrations vs Donations -->
        <div class="lg:col-span-2 bg-gray-900/50 backdrop-blur-xl rounded-3xl shadow-[0_8px_32px_rgba(0,0,0,0.2)] border border-[#00F0FF]/10 p-6 flex flex-col h-[400px]">
            <div class="flex justify-between items-center mb-6 shrink-0">
                <div>
                  <h2 class="text-lg font-black text-white tracking-tight drop-shadow-[0_0_5px_rgba(255,255,255,0.2)]">Activity Trends</h2>
                  <p class="text-[11px] font-bold text-[#00F0FF]/70 uppercase mt-0.5 font-mono">Real-time Data Stream</p>
                </div>
                <div class="px-3 py-1.5 bg-gray-800 border border-[#00F0FF]/20 rounded-lg text-xs font-bold text-[#00F0FF] shadow-[0_0_10px_rgba(0,240,255,0.1)]">
                  Live View
                </div>
            </div>
            <div class="grow relative w-full h-full pb-2">
                <Line v-if="chartDataLoaded" :data="lineChartData" :options="lineChartOptions" />
                <div v-else class="absolute inset-0 flex items-center justify-center text-[#00F0FF] font-mono text-sm animate-pulse">
                    <span class="material-icons animate-spin mr-2">group_work</span> Initializing Matrix...
                </div>
            </div>
        </div>

        <!-- Doughnut Chart: Blood Group Dist -->
        <div class="bg-gray-900/50 backdrop-blur-xl rounded-3xl shadow-[0_8px_32px_rgba(0,0,0,0.2)] border border-[#FF00FF]/10 p-6 flex flex-col h-[400px]">
             <div class="flex justify-between items-center mb-4 shrink-0">
                <div>
                  <h2 class="text-lg font-black text-white tracking-tight drop-shadow-[0_0_5px_rgba(255,255,255,0.2)]">Inventory Scan</h2>
                  <p class="text-[11px] font-bold text-[#FF00FF]/70 uppercase mt-0.5 font-mono">Blood Group Distribution</p>
                </div>
                <span class="material-icons text-[#FF00FF]/50 text-3xl shadow-[0_0_15px_rgba(255,0,255,0.2)] rounded-full">radar</span>
            </div>
            <div class="grow relative w-full h-full flex items-center justify-center pb-4">
                 <Doughnut v-if="chartDataLoaded" :data="doughnutChartData" :options="doughnutChartOptions" />
                 <div v-else class="absolute inset-0 flex items-center justify-center text-[#FF00FF] font-mono text-sm animate-pulse">
                    <span class="material-icons animate-spin mr-2">radar</span> Scanning...
                </div>
            </div>
        </div>
    </div>
    
    <!-- Bottom Grid -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        
        <!-- Activity Timeline -->
        <div class="lg:col-span-2 bg-gray-900/50 backdrop-blur-xl rounded-3xl shadow-[0_8px_32px_rgba(0,0,0,0.2)] border border-gray-700/50 p-6">
            <div class="flex justify-between items-center mb-6">
                <div>
                  <h2 class="text-lg font-black text-white tracking-tight drop-shadow-[0_0_5px_rgba(255,255,255,0.2)]">System Log</h2>
                  <p class="text-[11px] font-bold text-gray-400 uppercase mt-0.5 font-mono">Recent Interactions Output</p>
                </div>
                <router-link to="/admin/logs" class="text-xs font-bold text-[#00FF66] hover:text-[#00FF66] bg-[#00FF66]/10 hover:bg-[#00FF66]/20 px-4 py-2 rounded-xl transition-colors flex items-center gap-1 border border-[#00FF66]/30 shadow-[0_0_10px_rgba(0,255,102,0.1)]">
                Access Terminal <span class="material-icons text-sm">terminal</span>
                </router-link>
            </div>
            
            <div class="space-y-4">
                <div v-for="log in recentLogs" :key="log.id" class="flex items-start gap-4 p-4 bg-gray-800/40 rounded-2xl hover:bg-gray-800 transition-colors border border-gray-700/50 hover:border-[#00F0FF]/30 group">
                <div class="w-10 h-10 rounded-xl bg-[#00F0FF]/10 flex items-center justify-center text-[#00F0FF] border border-[#00F0FF]/20 shrink-0 shadow-[0_0_10px_rgba(0,240,255,0.1)] group-hover:shadow-[0_0_15px_rgba(0,240,255,0.3)] transition-shadow">
                    <span class="material-icons text-[18px]">memory</span>
                </div>
                <div class="flex-1 min-w-0 pt-0.5">
                    <div class="flex items-center flex-wrap gap-x-1 gap-y-1 text-sm">
                        <span class="font-bold text-gray-200">{{ log.viewer_name }}</span>
                        <span class="text-gray-500 font-mono text-xs">intercepted</span>
                        <button @click="viewProfile(log.target_unique_id)" class="font-bold text-[#FF00FF] hover:underline hover:text-[#FF00FF] transition-colors truncate max-w-[150px] sm:max-w-xs">{{ log.target_name }}</button>
                    </div>
                    <div class="flex items-center gap-2 mt-1.5">
                        <span class="inline-flex items-center gap-1 text-[10px] font-mono px-2 py-0.5 rounded-md bg-gray-900 border border-gray-700 text-[#00FF66]">
                            > {{ formatTimeAgo(log.created_at) }}
                        </span>
                        <span class="text-[10px] font-mono text-gray-500">SYS_ID: {{ log.target_unique_id }}</span>
                    </div>
                </div>
                </div>
                
                <div v-if="recentLogs.length === 0" class="flex flex-col items-center justify-center py-10 text-gray-500 font-mono">
                   <div class="w-16 h-16 rounded-full bg-gray-800 flex items-center justify-center mb-3 border border-gray-700 shadow-inner">
                       <span class="material-icons text-3xl opacity-50">data_array</span>
                   </div>
                   <p class="text-sm">Empty Buffer</p>
                </div>
            </div>
        </div>

        <!-- Quick Actions Panel -->
        <div class="bg-gray-900/80 backdrop-blur-xl rounded-3xl shadow-[0_8px_32px_rgba(0,0,0,0.4)] border border-[#00F0FF]/20 p-6 flex flex-col justify-between relative overflow-hidden group">
             <div class="absolute -right-8 -bottom-8 w-40 h-40 bg-[#00F0FF]/10 rounded-full blur-[40px] group-hover:bg-[#00F0FF]/20 transition-colors duration-700"></div>
             <div class="relative z-10 space-y-6">
                 <div>
                    <div class="w-12 h-12 rounded-2xl bg-[#00F0FF]/10 text-[#00F0FF] flex items-center justify-center mb-4 border border-[#00F0FF]/30 shadow-[0_0_20px_rgba(0,240,255,0.2)]">
                        <span class="material-icons">bolt</span>
                    </div>
                    <h2 class="text-xl font-black text-white tracking-tight drop-shadow-[0_0_5px_rgba(255,255,255,0.2)]">Override Protocols</h2>
                    <p class="text-xs font-mono text-gray-400 mt-2 leading-relaxed">Direct access to core system management sequences.</p>
                 </div>
                 
                 <div class="space-y-3">
                     <router-link to="/admin/donors" class="w-full bg-gray-800 hover:bg-gray-700 text-[#00F0FF] border border-[#00F0FF]/20 hover:border-[#00F0FF]/50 px-4 py-3.5 rounded-2xl font-mono text-sm shadow-[0_0_10px_rgba(0,0,0,0.5)] transition-all flex items-center justify-between group/btn">
                         <div class="flex items-center gap-3">
                             <div class="w-8 h-8 rounded-lg bg-gray-900 flex items-center justify-center border border-gray-700"><span class="material-icons text-[16px]">manage_accounts</span></div>
                             INIT_USERS
                         </div>
                         <span class="material-icons text-sm text-[#00F0FF]/50 group-hover/btn:translate-x-1 group-hover/btn:text-[#00F0FF] transition-all">chevron_right</span>
                     </router-link>
                     <button @click="$router.push('/admin/logs')" class="w-full bg-gray-800 hover:bg-gray-700 text-[#FF00FF] border border-[#FF00FF]/20 hover:border-[#FF00FF]/50 px-4 py-3.5 rounded-2xl font-mono text-sm shadow-[0_0_10px_rgba(0,0,0,0.5)] transition-all flex items-center justify-between group/btn">
                         <div class="flex items-center gap-3">
                             <div class="w-8 h-8 rounded-lg bg-gray-900 flex items-center justify-center border border-gray-700"><span class="material-icons text-[16px]">verified_user</span></div>
                             READ_LOGS
                         </div>
                         <span class="material-icons text-sm text-[#FF00FF]/50 group-hover/btn:translate-x-1 group-hover/btn:text-[#FF00FF] transition-all">chevron_right</span>
                     </button>
                 </div>
             </div>
        </div>

    </div>

  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue';
import { useRouter } from 'vue-router';
import api from '@/lib/axios';
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  BarElement,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  ArcElement,
  Filler
} from 'chart.js';
import { Line, Doughnut } from 'vue-chartjs';

// Register Chart.js components
ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  PointElement,
  LineElement,
  ArcElement,
  Title,
  Tooltip,
  Legend,
  Filler
);

const router = useRouter();

const isLoading = ref(false);
const chartDataLoaded = ref(false);
const dateRange = ref('180'); // Default to 6 months

const stats = ref({
  total_donors: 0,
  total_donations: 0,
  total_users: 0,
  growth: {
      donors: '0',
      donations: '0',
      users: '0'
  }
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

// Chart Data References
const rawDonationTrends = ref<{month: string, count: number}[]>([]);
const rawUserTrends = ref<{month: string, count: number}[]>([]);
const rawBloodGroupDist = ref<{blood_group: string, count: number}[]>([]);

const fetchData = async () => {
  isLoading.value = true;
  chartDataLoaded.value = false;
  try {
    const res = await api.get(`/admin/stats?range=${dateRange.value}`);
    stats.value = res.data;
    
    // Store chart data
    rawDonationTrends.value = res.data.monthly_donations || [];
    rawUserTrends.value = res.data.monthly_users || [];
    rawBloodGroupDist.value = res.data.blood_group_distribution || [];
    
    chartDataLoaded.value = true;
  } catch (error) {
    console.error("Failed to load stats", error);
  } finally {
      isLoading.value = false;
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

onMounted(() => {
  fetchData();
  fetchRecentLogs();
});

watch(dateRange, () => {
  fetchData();
});

// --- Chart Configurations ---

// 1. Line Chart (Activity Trends)
const lineChartData = computed(() => {
    // Merge labels from both datasets to handle missing months gracefully
    const labelsSet = new Set<string>();
    rawDonationTrends.value.forEach(d => labelsSet.add(d.month));
    rawUserTrends.value.forEach(u => labelsSet.add(u.month));
    
    const labels = Array.from(labelsSet).sort();
    
    // Map data to sorted labels
    const donationData = labels.map(label => {
        const found = rawDonationTrends.value.find(d => d.month === label);
        return found ? found.count : 0;
    });
    
    const userData = labels.map(label => {
        const found = rawUserTrends.value.find(u => u.month === label);
        return found ? found.count : 0;
    });

    // Format labels nicely (e.g. "2023-10" to "Oct '23")
    const formattedLabels = labels.map(l => {
        const d = new Date(l + '-01');
        return d.toLocaleDateString('en-US', { month: 'short', year: '2-digit' });
    });

    return {
        labels: formattedLabels,
        datasets: [
            {
                label: 'New Donations',
                backgroundColor: 'rgba(0, 240, 255, 0.15)', // Neon Cyan Fill
                borderColor: '#00F0FF',
                pointBackgroundColor: '#00F0FF',
                pointBorderColor: '#fff',
                pointBorderWidth: 1,
                pointRadius: 4,
                pointHoverRadius: 6,
                borderWidth: 3,
                fill: true,
                tension: 0.4, // smooth curves
                data: donationData
            },
            {
                label: 'New Users',
                backgroundColor: 'rgba(255, 0, 255, 0.05)', // Neon Magenta
                borderColor: '#FF00FF',
                borderDash: [5, 5],
                pointBackgroundColor: '#FF00FF',
                pointBorderColor: '#fff',
                pointBorderWidth: 1,
                pointRadius: 4,
                borderWidth: 2,
                tension: 0.4,
                data: userData
            }
        ]
    };
});

const lineChartOptions = {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
        legend: {
            position: 'top' as const,
            align: 'end' as const,
            labels: {
                usePointStyle: true,
                boxWidth: 8,
                font: {
                    family: "'JetBrains Mono', 'Courier New', monospace",
                    weight: 'bold' as const,
                    size: 11
                },
                color: '#00F0FF'
            }
        },
        tooltip: {
            backgroundColor: 'rgba(15, 23, 42, 0.95)',
            titleFont: { family: "'JetBrains Mono', monospace", size: 13, weight: 'bold' as const },
            bodyFont: { family: "'JetBrains Mono', monospace", size: 12 },
            padding: 12,
            cornerRadius: 4,
            displayColors: true,
            boxPadding: 4,
            borderColor: '#00F0FF',
            borderWidth: 1
        }
    },
    scales: {
        y: {
            beginAtZero: true,
            grid: {
                color: 'rgba(255, 255, 255, 0.05)',
                drawBorder: false,
            },
            ticks: {
                font: { family: "'JetBrains Mono', monospace", size: 11 },
                color: '#64748B',
                padding: 8,
                stepSize: 1
            }
        },
        x: {
            grid: {
                display: false,
                drawBorder: false,
            },
            ticks: {
                font: { family: "'JetBrains Mono', monospace", size: 11, weight: 'bold' as const },
                color: '#64748B',
                padding: 8
            }
        }
    },
    interaction: {
        intersect: false,
        mode: 'index' as const,
    },
};

// 2. Doughnut Chart (Blood Group Dist)
const doughnutChartData = computed(() => {
    const labels = rawBloodGroupDist.value.map(item => item.blood_group);
    const data = rawBloodGroupDist.value.map(item => item.count);
    
    // Sci-fi Neon color palette for blood groups
    const backgroundColors = [
        '#FF003C', // A+   (Cyberpunk Red)
        '#FF8A00', // A-   (Neon Orange)
        '#00F0FF', // B+   (Cyan)
        '#0080FF', // B-   (Deep Blue)
        '#FF00FF', // AB+  (Magenta)
        '#8A2BE2', // AB-  (Blue Violet)
        '#00FF66', // O+   (Matrix Green)
        '#DFFF00'  // O-   (Chartreuse)
    ];

    return {
        labels,
        datasets: [{
            data,
            backgroundColor: backgroundColors.slice(0, data.length),
            borderWidth: 2,
            borderColor: '#0F172A', // Match Dark background to look like cutouts
            hoverOffset: 12
        }]
    };
});

const doughnutChartOptions = {
    responsive: true,
    maintainAspectRatio: false,
    cutout: '70%',
    plugins: {
        legend: {
            position: 'right' as const,
            labels: {
                usePointStyle: true,
                pointStyle: 'circle',
                boxWidth: 8,
                padding: 16,
                font: {
                    family: "'JetBrains Mono', monospace",
                    weight: 'bold' as const,
                    size: 11
                },
                color: '#94A3B8'
            }
        },
        tooltip: {
            backgroundColor: 'rgba(15, 23, 42, 0.95)',
            titleFont: { family: "'JetBrains Mono', monospace", size: 13, weight: 'bold' as const },
            bodyFont: { family: "'JetBrains Mono', monospace", size: 12 },
            padding: 12,
            cornerRadius: 4,
            displayColors: true,
            borderColor: '#FF00FF',
            borderWidth: 1,
            boxPadding: 4,
            callbacks: {
                label: function(context: any) {
                    let label = context.label || '';
                    if (label) {
                        label += ': ';
                    }
                    if (context.parsed !== null) {
                        label += context.parsed + ' donors';
                    }
                    return label;
                }
            }
        }
    }
};


// Utils
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
  router.push(`/admin/donors/${userId}`);
};
</script>
