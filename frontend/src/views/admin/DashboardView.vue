<template>
  <div class="space-y-6 lg:space-y-8 max-w-7xl mx-auto pb-10">
    
    <!-- Header Section -->
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-end gap-4 bg-gradient-to-br from-white to-gray-50 p-6 rounded-2xl shadow-sm border border-gray-200">
      <div>
        <h1 class="text-2xl sm:text-3xl font-black text-slate-800 tracking-tight">
          System Overview
        </h1>
        <p class="text-sm font-medium text-gray-500 mt-1">Monitor blood bank performance and user engagement</p>
      </div>
      <div class="flex items-center gap-3">
        <!-- Date Range Filter -->
        <select 
          v-model="dateRange" 
          class="bg-gray-50 border border-gray-200 text-gray-700 text-sm rounded-xl focus:ring-[#FF3D3D] focus:border-[#FF3D3D] block p-2 font-medium cursor-pointer hover:bg-gray-100 transition-colors outline-none"
        >
          <option value="7">Last 7 Days</option>
          <option value="30">Last 30 Days</option>
          <option value="90">Last 3 Months</option>
          <option value="180">Last 6 Months</option>
          <option value="365">Last 1 Year</option>
          <option value="1825">All Time</option>
        </select>

        <span class="hidden sm:inline-flex items-center gap-1.5 px-3 py-1.5 rounded-lg bg-red-50 text-[#FF3D3D] text-xs font-bold border border-red-100">
          <span class="w-2 h-2 rounded-full bg-[#FF3D3D] animate-pulse"></span>
          Live Sync
        </span>
        <button @click="fetchData" class="w-10 h-10 shrink-0 rouned-xl bg-gray-50 hover:bg-gray-100 flex items-center justify-center text-gray-500 hover:text-slate-800 transition-all border border-gray-200 rounded-xl" title="Refresh">
          <span class="material-icons text-sm" :class="{'animate-spin text-[#FF3D3D]': isLoading}">refresh</span>
        </button>
      </div>
    </div>

    <!-- Quick Stats Grid -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 sm:gap-6">
      
      <!-- Total Donors -->
      <div class="relative overflow-hidden bg-gradient-to-br from-white to-gray-50 rounded-2xl shadow-sm border border-gray-200 p-6 group hover:shadow-lg transition-all duration-300">
        <div class="absolute -right-6 -top-6 w-32 h-32 bg-red-50 rounded-full blur-3xl group-hover:bg-red-100 transition-colors duration-500"></div>
        <div class="relative">
            <div class="flex justify-between items-start mb-6">
                <div class="w-12 h-12 rounded-xl bg-red-50 flex items-center justify-center text-[#FF3D3D] border border-red-100 group-hover:scale-105 transition-transform duration-300">
                    <span class="material-icons text-xl">people</span>
                </div>
                <!-- Growth Indicator -->
                <span v-if="stats.growth?.donors" class="inline-flex items-center gap-1 text-[11px] font-bold px-2 py-1 rounded-lg" :class="parseFloat(stats.growth.donors) >= 0 ? 'bg-emerald-50 text-emerald-600 border border-emerald-100' : 'bg-red-50 text-red-600 border border-red-100'">
                    <span class="material-icons text-[12px]">{{ parseFloat(stats.growth.donors) >= 0 ? 'trending_up' : 'trending_down' }}</span> 
                    {{ Math.abs(parseFloat(stats.growth.donors)) }}%
                </span>
            </div>
            <div>
              <p class="text-xs font-bold text-gray-400 uppercase tracking-wider mb-1">Registered Donors</p>
              <h3 class="text-4xl font-black text-slate-800 tracking-tight">{{ stats.total_donors }}</h3>
            </div>
        </div>
      </div>

      <!-- Total Donations -->
      <div class="relative overflow-hidden bg-gradient-to-br from-white to-gray-50 rounded-2xl shadow-sm border border-gray-200 p-6 group hover:shadow-lg transition-all duration-300">
        <div class="absolute -right-6 -top-6 w-32 h-32 bg-rose-50 rounded-full blur-3xl group-hover:bg-rose-100 transition-colors duration-500"></div>
        <div class="relative">
             <div class="flex justify-between items-start mb-6">
                <div class="w-12 h-12 rounded-xl bg-rose-50 flex items-center justify-center text-rose-500 border border-rose-100 group-hover:scale-105 transition-transform duration-300">
                    <span class="material-icons text-xl">volunteer_activism</span>
                </div>
                <!-- Growth Indicator -->
                <span v-if="stats.growth?.donations" class="inline-flex items-center gap-1 text-[11px] font-bold px-2 py-1 rounded-lg" :class="parseFloat(stats.growth.donations) >= 0 ? 'bg-emerald-50 text-emerald-600 border border-emerald-100' : 'bg-red-50 text-red-600 border border-red-100'">
                    <span class="material-icons text-[12px]">{{ parseFloat(stats.growth.donations) >= 0 ? 'trending_up' : 'trending_down' }}</span> 
                    {{ Math.abs(parseFloat(stats.growth.donations)) }}%
                </span>
            </div>
            <div>
              <p class="text-xs font-bold text-gray-400 uppercase tracking-wider mb-1">Total Donations</p>
              <h3 class="text-4xl font-black text-slate-800 tracking-tight">{{ stats.total_donations }}</h3>
            </div>
        </div>
      </div>

       <!-- Lives Saved Estimate -->
      <div class="relative overflow-hidden bg-gradient-to-br from-white to-gray-50 rounded-2xl shadow-sm border border-gray-200 p-6 group hover:shadow-lg transition-all duration-300">
        <div class="absolute -right-6 -top-6 w-32 h-32 bg-pink-50 rounded-full blur-3xl group-hover:bg-pink-100 transition-colors duration-500"></div>
        <div class="relative">
             <div class="flex justify-between items-start mb-6">
                <div class="w-12 h-12 rounded-xl bg-pink-50 flex items-center justify-center text-pink-500 border border-pink-100 group-hover:scale-105 transition-transform duration-300">
                    <span class="material-icons text-xl">favorite</span>
                </div>
            </div>
            <div>
              <p class="text-xs font-bold text-gray-400 uppercase tracking-wider mb-1">Est. Lives Saved</p>
              <h3 class="text-4xl font-black text-[#FF3D3D] tracking-tight">{{ stats.total_donations * 3 }}</h3>
            </div>
        </div>
      </div>

      <!-- Total Users -->
      <div class="relative overflow-hidden bg-gradient-to-br from-indigo-50/80 to-indigo-100/50 rounded-2xl shadow-sm border border-indigo-200 p-6 group hover:shadow-lg transition-all duration-300">
        <div class="absolute right-0 top-0 h-full w-1/2 bg-gradient-to-l from-white/50 to-transparent pointer-events-none rounded-r-2xl"></div>
        <div class="relative z-10">
             <div class="flex justify-between items-start mb-6">
                <div class="w-12 h-12 rounded-xl bg-white flex items-center justify-center text-indigo-600 border border-indigo-100 group-hover:scale-105 transition-transform duration-300 shadow-sm">
                    <span class="material-icons text-xl">account_circle</span>
                </div>
                <!-- Growth Indicator -->
                <span v-if="stats.growth?.users" class="inline-flex items-center gap-1 text-[11px] font-bold px-2 py-1 rounded-lg bg-white border border-indigo-100 text-indigo-600 shadow-sm">
                    <span class="material-icons text-[12px]">{{ parseFloat(stats.growth.users) >= 0 ? 'trending_up' : 'trending_down' }}</span> 
                    {{ Math.abs(parseFloat(stats.growth.users)) }}%
                </span>
            </div>
            <div>
              <p class="text-xs font-bold text-indigo-400 uppercase tracking-wider mb-1">System Accounts</p>
              <h3 class="text-4xl font-black text-indigo-900 tracking-tight">{{ stats.total_users }}</h3>
            </div>
        </div>
      </div>
    </div>

    <!-- Charts Section -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        
        <!-- Main Line Chart: Registrations vs Donations -->
        <div class="lg:col-span-2 bg-gradient-to-br from-white to-gray-50 rounded-2xl shadow-sm border border-gray-200 p-6 flex flex-col h-[400px] hover:shadow-md transition-shadow duration-300">
            <div class="flex justify-between items-center mb-6 shrink-0">
                <div>
                  <h2 class="text-lg font-black text-slate-800 tracking-tight">Activity Trends</h2>
                  <p class="text-[11px] font-bold text-gray-400 uppercase mt-0.5">Donations vs Users over time</p>
                </div>
                <div class="px-3 py-1.5 bg-gray-50 border border-gray-100 rounded-lg text-xs font-bold text-gray-500">
                  Dynamics
                </div>
            </div>
            <div class="grow relative w-full h-full pb-2">
                <Line v-if="chartDataLoaded" :data="lineChartData" :options="lineChartOptions" />
                <div v-else class="absolute inset-0 flex items-center justify-center text-gray-400 text-sm">
                    <span class="material-icons animate-spin mr-2">refresh</span> Loading data...
                </div>
            </div>
        </div>

        <!-- Doughnut Chart: Blood Group Dist -->
        <div class="bg-gradient-to-br from-white to-gray-50 rounded-2xl shadow-sm border border-gray-200 p-6 flex flex-col h-[400px] hover:shadow-md transition-shadow duration-300">
             <div class="flex justify-between items-center mb-4 shrink-0">
                <div>
                  <h2 class="text-lg font-black text-slate-800 tracking-tight">Blood Inventory</h2>
                  <p class="text-[11px] font-bold text-gray-400 uppercase mt-0.5">Donor Distribution by Type</p>
                </div>
                <span class="material-icons text-red-200 text-3xl">bloodtype</span>
            </div>
            <div class="grow relative w-full h-full flex items-center justify-center pb-4">
                 <Doughnut v-if="chartDataLoaded" :data="doughnutChartData" :options="doughnutChartOptions" />
                 <div v-else class="absolute inset-0 flex items-center justify-center text-gray-400 text-sm">
                    <span class="material-icons animate-spin mr-2">refresh</span> Loading chart...
                </div>
            </div>
        </div>
    </div>
    
    <!-- Bottom Grid -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        
        <!-- Activity Timeline -->
        <div class="lg:col-span-2 bg-gradient-to-br from-white to-gray-50 rounded-2xl shadow-sm border border-gray-200 p-6 hover:shadow-md transition-shadow duration-300">
            <div class="flex justify-between items-center mb-6">
                <div>
                  <h2 class="text-lg font-black text-slate-800 tracking-tight">Recent Activity Feed</h2>
                  <p class="text-[11px] font-bold text-gray-400 uppercase mt-0.5">Latest system interactions</p>
                </div>
                <router-link to="/admin/logs" class="text-xs font-bold text-[#FF3D3D] hover:text-red-700 bg-red-50 hover:bg-red-100 px-4 py-2 rounded-xl transition-colors flex items-center gap-1 border border-red-100 shadow-sm">
                View Log <span class="material-icons text-sm">arrow_forward</span>
                </router-link>
            </div>
            
            <div class="space-y-4">
                <div v-for="log in recentLogs" :key="log.id" class="flex items-start gap-4 p-4 bg-gray-50/50 rounded-2xl hover:bg-gray-50 transition-colors border border-transparent hover:border-gray-200 group">
                <div class="w-10 h-10 rounded-full bg-blue-50 flex items-center justify-center text-blue-500 border border-blue-100 shrink-0 shadow-sm group-hover:scale-105 transition-transform">
                    <span class="material-icons text-[18px]">person_search</span>
                </div>
                <div class="flex-1 min-w-0 pt-0.5">
                    <div class="flex items-center flex-wrap gap-x-1 gap-y-1 text-sm">
                        <span class="font-bold text-slate-800">{{ log.viewer_name }}</span>
                        <span class="text-gray-500 font-medium text-sm">viewed the profile of</span>
                        <button @click="viewProfile(log.target_unique_id)" class="font-bold text-[#FF3D3D] hover:underline hover:text-red-700 transition-colors truncate max-w-[150px] sm:max-w-xs">{{ log.target_name }}</button>
                    </div>
                    <div class="flex items-center gap-2 mt-1.5">
                        <span class="inline-flex items-center gap-1 text-[10px] font-bold px-2 py-0.5 rounded-md bg-white border border-gray-200 text-gray-500 shadow-sm">
                            <span class="material-icons text-[10px]">schedule</span> {{ formatTimeAgo(log.created_at) }}
                        </span>
                    </div>
                </div>
                </div>
                
                <div v-if="recentLogs.length === 0" class="flex flex-col items-center justify-center py-10 text-gray-400">
                   <div class="w-16 h-16 rounded-full bg-gray-50 flex items-center justify-center mb-3 border border-gray-100 shadow-sm">
                       <span class="material-icons text-3xl opacity-50">data_array</span>
                   </div>
                   <p class="text-sm font-bold">No activity recorded yet</p>
                </div>
            </div>
        </div>

        <!-- Quick Actions Panel -->
        <div class="bg-gradient-to-br from-indigo-50/80 to-indigo-100/50 rounded-2xl shadow-sm border border-indigo-200 p-6 flex flex-col justify-between relative overflow-hidden group hover:shadow-md transition-shadow duration-300">
             <div class="absolute -right-8 -bottom-8 w-40 h-40 bg-indigo-500/10 rounded-full blur-3xl group-hover:bg-indigo-500/20 transition-colors duration-500"></div>
             <div class="relative z-10 space-y-6">
                 <div>
                    <div class="w-12 h-12 rounded-xl bg-indigo-100 text-indigo-600 flex items-center justify-center mb-4 border border-indigo-200 shadow-sm">
                        <span class="material-icons">flash_on</span>
                    </div>
                    <h2 class="text-xl font-black text-slate-800 tracking-tight">Admin Actions</h2>
                    <p class="text-xs text-gray-500 font-medium mt-2 leading-relaxed">Direct access to core system management sections.</p>
                 </div>
                 
                 <div class="space-y-3">
                     <router-link to="/admin/donors" class="w-full bg-white hover:bg-gray-50 text-indigo-600 border border-indigo-100 hover:border-indigo-300 px-4 py-3.5 rounded-2xl font-bold text-sm shadow-sm transition-all flex items-center justify-between group/btn">
                         <div class="flex items-center gap-3">
                             <div class="w-8 h-8 rounded-lg bg-indigo-50 flex items-center justify-center"><span class="material-icons text-[16px]">manage_accounts</span></div>
                             Manage Users
                         </div>
                         <span class="material-icons text-sm text-indigo-300 group-hover/btn:translate-x-1 group-hover/btn:text-indigo-600 transition-all">arrow_forward</span>
                     </router-link>
                     <button @click="$router.push('/admin/logs')" class="w-full bg-white hover:bg-gray-50 text-slate-700 border border-gray-200 hover:border-gray-300 px-4 py-3.5 rounded-2xl font-bold text-sm shadow-sm transition-all flex items-center justify-between group/btn">
                         <div class="flex items-center gap-3">
                             <div class="w-8 h-8 rounded-lg bg-gray-50 flex items-center justify-center border border-gray-100"><span class="material-icons text-[16px]">security</span></div>
                             Security Logs
                         </div>
                         <span class="material-icons text-sm text-gray-300 group-hover/btn:translate-x-1 group-hover/btn:text-slate-600 transition-all">arrow_forward</span>
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
                backgroundColor: 'rgba(255, 61, 61, 0.12)', // Soft red gradient
                borderColor: '#FF3D3D',
                pointBackgroundColor: '#FF3D3D',
                pointBorderColor: '#fff',
                pointBorderWidth: 2,
                pointRadius: 4,
                pointHoverRadius: 6,
                borderWidth: 3,
                fill: true,
                tension: 0.4, // smooth curves
                data: donationData
            },
            {
                label: 'New Users',
                backgroundColor: 'transparent',
                borderColor: '#6366F1',
                borderDash: [5, 5],
                pointBackgroundColor: '#6366F1',
                pointBorderColor: '#fff',
                pointBorderWidth: 2,
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
                padding: 20,
                font: {
                    family: "'Plus Jakarta Sans', sans-serif",
                    weight: 'bold' as const,
                    size: 13
                },
                color: '#475569'
            }
        },
        tooltip: {
            backgroundColor: 'rgba(255, 255, 255, 0.98)',
            titleColor: '#0F172A',
            bodyColor: '#334155',
            titleFont: { family: "'Plus Jakarta Sans', sans-serif", size: 14, weight: 'bold' as const },
            bodyFont: { family: "'Plus Jakarta Sans', sans-serif", size: 13, weight: 'normal' as const },
            padding: 16,
            cornerRadius: 12,
            displayColors: true,
            boxPadding: 6,
            borderColor: '#E2E8F0',
            borderWidth: 1,
            boxShadow: '0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)'
        }
    },
    scales: {
        y: {
            beginAtZero: true,
            grid: {
                color: '#F1F5F9', // light slate gray
                drawBorder: false,
            },
            border: { display: false },
            ticks: {
                font: { family: "'Plus Jakarta Sans', sans-serif", size: 12, weight: 'normal' as const },
                color: '#64748B',
                padding: 12,
                stepSize: 1
            }
        },
        x: {
            grid: {
                display: false,
                drawBorder: false,
            },
            border: { display: false },
            ticks: {
                font: { family: "'Plus Jakarta Sans', sans-serif", size: 12, weight: 'bold' as const },
                color: '#64748B',
                padding: 12
            }
        }
    },
    elements: {
        line: {
            borderJoinStyle: 'round' as const
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
    
    // Professional color palette for blood groups matching the app
    const backgroundColors = [
        '#FF3D3D', // A+   (Primary Red)
        '#FCA5A5', // A-   (Light Red)
        '#4F46E5', // B+   (Indigo)
        '#A5B4FC', // B-   (Light Indigo)
        '#059669', // AB+  (Emerald)
        '#D1FAE5', // AB-  (Light Emerald)
        '#0284C7', // O+   (Sky Blue)
        '#E0F2FE'  // O-   (Light Sky)
    ];

    return {
        labels,
        datasets: [{
            data,
            backgroundColor: backgroundColors.slice(0, data.length),
            borderWidth: 2,
            borderColor: '#ffffff', // Clean white border to separate segments
            hoverOffset: 8
        }]
    };
});

const doughnutChartOptions = {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
        legend: {
            position: 'right' as const,
            labels: {
                usePointStyle: true,
                pointStyle: 'circle',
                boxWidth: 8,
                padding: 16,
                font: {
                    family: "'Plus Jakarta Sans', sans-serif",
                    weight: 'bold' as const,
                    size: 12
                },
                color: '#475569'
            }
        },
        tooltip: {
            backgroundColor: 'rgba(255, 255, 255, 0.95)',
            titleColor: '#1E293B',
            bodyColor: '#475569',
            titleFont: { family: "'Plus Jakarta Sans', sans-serif", size: 13, weight: 'bold' as const },
            bodyFont: { family: "'Plus Jakarta Sans', sans-serif", size: 12, weight: 'normal' as const },
            padding: 12,
            cornerRadius: 8,
            displayColors: true,
            borderColor: '#E2E8F0',
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
