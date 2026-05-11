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

    <!-- Bento Box Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-5">
      
      <!-- Total Donors (1x1) -->
      <div class="relative overflow-hidden bg-gradient-to-br from-white to-gray-50 rounded-[2rem] shadow-sm border border-gray-100 p-6 group transition-all duration-300">
        <div class="absolute -right-6 -top-6 w-32 h-32 bg-red-50 rounded-full blur-3xl group-hover:bg-red-100 transition-colors duration-500"></div>
        <div class="relative">
            <div class="flex justify-between items-start mb-6">
                <div class="w-12 h-12 rounded-2xl bg-red-50 flex items-center justify-center text-[#FF3D3D] group-hover:scale-105 transition-transform duration-300">
                    <span class="material-icons text-xl">people</span>
                </div>
                <!-- Growth Indicator -->
                <span v-if="stats.growth?.donors" class="inline-flex items-center gap-1 text-[11px] font-bold px-2.5 py-1 rounded-xl" :class="parseFloat(stats.growth.donors) >= 0 ? 'bg-emerald-50 text-emerald-600' : 'bg-red-50 text-red-600'">
                    <span class="material-icons text-[12px]">{{ parseFloat(stats.growth.donors) >= 0 ? 'trending_up' : 'trending_down' }}</span> 
                    {{ Math.abs(parseFloat(stats.growth.donors)) }}%
                </span>
            </div>
            <div>
              <p class="text-xs font-bold text-gray-400 tracking-wide mb-1">Registered Donors</p>
              <h3 class="text-4xl font-black text-slate-800 tracking-tighter">{{ stats.total_donors }}</h3>
            </div>
        </div>
      </div>

      <!-- Total Donations (1x1) -->
      <div class="relative overflow-hidden bg-gradient-to-br from-white to-gray-50 rounded-[2rem] shadow-sm border border-gray-100 p-6 group transition-all duration-300">
        <div class="absolute -right-6 -top-6 w-32 h-32 bg-rose-50 rounded-full blur-3xl group-hover:bg-rose-100 transition-colors duration-500"></div>
        <div class="relative">
             <div class="flex justify-between items-start mb-6">
                <div class="w-12 h-12 rounded-2xl bg-rose-50 flex items-center justify-center text-rose-500 group-hover:scale-105 transition-transform duration-300">
                    <span class="material-icons text-xl">volunteer_activism</span>
                </div>
                <!-- Growth Indicator -->
                <span v-if="stats.growth?.donations" class="inline-flex items-center gap-1 text-[11px] font-bold px-2.5 py-1 rounded-xl" :class="parseFloat(stats.growth.donations) >= 0 ? 'bg-emerald-50 text-emerald-600' : 'bg-red-50 text-red-600'">
                    <span class="material-icons text-[12px]">{{ parseFloat(stats.growth.donations) >= 0 ? 'trending_up' : 'trending_down' }}</span> 
                    {{ Math.abs(parseFloat(stats.growth.donations)) }}%
                </span>
            </div>
            <div>
              <p class="text-xs font-bold text-gray-400 tracking-wide mb-1">Total Donations</p>
              <h3 class="text-4xl font-black text-slate-800 tracking-tighter">{{ stats.total_donations }}</h3>
            </div>
        </div>
      </div>

       <!-- Lives Saved Estimate (1x1) -->
      <div class="relative overflow-hidden bg-gradient-to-br from-white to-gray-50 rounded-[2rem] shadow-sm border border-gray-100 p-6 group transition-all duration-300">
        <div class="absolute -right-6 -top-6 w-32 h-32 bg-pink-50 rounded-full blur-3xl group-hover:bg-pink-100 transition-colors duration-500"></div>
        <div class="relative">
             <div class="flex justify-between items-start mb-6">
                <div class="w-12 h-12 rounded-2xl bg-pink-50 flex items-center justify-center text-pink-500 group-hover:scale-105 transition-transform duration-300">
                    <span class="material-icons text-xl">favorite</span>
                </div>
            </div>
            <div>
              <p class="text-xs font-bold text-gray-400 tracking-wide mb-1">Est. Lives Saved</p>
              <h3 class="text-4xl font-black text-[#FF3D3D] tracking-tighter">{{ stats.total_donations * 3 }}</h3>
            </div>
        </div>
      </div>

      <!-- Total Users (1x1) -->
      <div class="relative overflow-hidden bg-gradient-to-br from-indigo-50/80 to-indigo-100/50 rounded-[2rem] shadow-sm border border-indigo-100/50 p-6 group transition-all duration-300">
        <div class="absolute right-0 top-0 h-full w-1/2 bg-gradient-to-l from-white/50 to-transparent pointer-events-none rounded-r-[2rem]"></div>
        <div class="relative z-10">
             <div class="flex justify-between items-start mb-6">
                <div class="w-12 h-12 rounded-2xl bg-white flex items-center justify-center text-indigo-600 group-hover:scale-105 transition-transform duration-300 shadow-sm">
                    <span class="material-icons text-xl">account_circle</span>
                </div>
                <!-- Growth Indicator -->
                <span v-if="stats.growth?.users" class="inline-flex items-center gap-1 text-[11px] font-bold px-2.5 py-1 rounded-xl bg-white text-indigo-600 shadow-sm">
                    <span class="material-icons text-[12px]">{{ parseFloat(stats.growth.users) >= 0 ? 'trending_up' : 'trending_down' }}</span> 
                    {{ Math.abs(parseFloat(stats.growth.users)) }}%
                </span>
            </div>
            <div>
              <p class="text-xs font-bold text-indigo-400/90 tracking-wide mb-1">System Accounts</p>
              <h3 class="text-4xl font-black text-indigo-900 tracking-tighter">{{ stats.total_users }}</h3>
            </div>
        </div>
      </div>

        
      <!-- Main Line Chart: Registrations vs Donations (2x2) -->
      <div class="md:col-span-2 lg:col-span-2 lg:row-span-2 bg-gradient-to-br from-white to-gray-50 rounded-[2rem] shadow-sm border border-gray-100 p-6 flex flex-col min-h-[400px]">
          <div class="flex justify-between items-center mb-6 shrink-0">
              <div>
                <h2 class="text-xl font-black text-slate-800 tracking-tight">Activity Trends</h2>
                <p class="text-xs font-bold text-gray-400 mt-1">Donations vs Users over time</p>
              </div>
              <div class="px-4 py-2 bg-gray-50 border border-gray-100 rounded-xl text-xs font-bold text-gray-500">
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

      <!-- Doughnut Chart: Blood Group Dist (1x2) -->
      <div class="lg:col-span-1 lg:row-span-2 bg-gradient-to-br from-white to-gray-50 rounded-[2rem] shadow-sm border border-gray-100 p-6 flex flex-col min-h-[400px]">
            <div class="flex justify-between items-center mb-4 shrink-0">
              <div>
                <h2 class="text-xl font-black text-slate-800 tracking-tight">Inventory</h2>
                <p class="text-xs font-bold text-gray-400 mt-1">Distribution by Type</p>
              </div>
          </div>
          <div class="grow relative w-full h-full flex items-center justify-center pb-4">
                <PolarArea v-if="chartDataLoaded" :data="polarAreaChartData" :options="polarAreaChartOptions" />
                <div v-else class="absolute inset-0 flex items-center justify-center text-gray-400 text-sm">
                  <span class="material-icons animate-spin mr-2">refresh</span> Loading chart...
              </div>
          </div>
      </div>

      <!-- Quick Actions Panel (1x2) -->
      <div class="lg:col-span-1 lg:row-span-2 bg-gradient-to-br from-slate-800 to-slate-900 rounded-[2rem] shadow-md border border-slate-700/50 p-6 flex flex-col justify-between relative overflow-hidden group">
            <div class="absolute -right-8 -bottom-8 w-48 h-48 bg-indigo-500/20 rounded-full blur-3xl group-hover:bg-indigo-500/30 transition-colors duration-500"></div>
            <div class="relative z-10 space-y-6 flex-1 flex flex-col">
                <div>
                  <div class="w-14 h-14 rounded-2xl bg-white/10 text-white flex items-center justify-center mb-5 backdrop-blur-sm border border-white/5">
                      <span class="material-icons text-2xl">flash_on</span>
                  </div>
                  <h2 class="text-2xl font-black text-white tracking-tight">Admin Actions</h2>
                  <p class="text-sm text-slate-400 font-medium mt-2 leading-relaxed">Direct access to core system tools.</p>
                </div>
                
                <div class="space-y-4 mt-auto">
                    <router-link to="/admin/donors" class="w-full bg-white/10 hover:bg-white/15 text-white border border-white/10 px-5 py-4 rounded-[1.5rem] font-bold text-sm backdrop-blur-sm transition-all flex items-center justify-between group/btn">
                        <div class="flex items-center gap-4">
                            <span class="material-icons text-[20px] text-indigo-300">manage_accounts</span>
                            Manage Users
                        </div>
                        <span class="material-icons text-sm opacity-50 group-hover/btn:translate-x-1 group-hover/btn:opacity-100 transition-all">arrow_forward</span>
                    </router-link>
                    <button @click="$router.push('/admin/logs')" class="w-full bg-white/10 hover:bg-white/15 text-white border border-white/10 px-5 py-4 rounded-[1.5rem] font-bold text-sm backdrop-blur-sm transition-all flex items-center justify-between group/btn">
                        <div class="flex items-center gap-4">
                            <span class="material-icons text-[20px] text-gray-300">security</span>
                            Security Logs
                        </div>
                        <span class="material-icons text-sm opacity-50 group-hover/btn:translate-x-1 group-hover/btn:opacity-100 transition-all">arrow_forward</span>
                    </button>
                </div>
            </div>
      </div>

      <!-- Activity Timeline (2x1) -->
      <div class="md:col-span-2 lg:col-span-2 bg-gradient-to-br from-white to-gray-50 rounded-[2rem] shadow-sm border border-gray-100 p-6 lg:row-start-4">
          <div class="flex justify-between items-center mb-6">
              <div>
                <h2 class="text-xl font-black text-slate-800 tracking-tight">Activity Feed</h2>
                <p class="text-xs font-bold text-gray-400 mt-1">Latest system interactions</p>
              </div>
              <router-link to="/admin/logs" class="shrink-0 w-10 h-10 flex items-center justify-center rounded-2xl bg-red-50 text-[#FF3D3D] hover:bg-[#FF3D3D] hover:text-white transition-colors">
                 <span class="material-icons text-[20px]">open_in_new</span>
              </router-link>
          </div>
          
          <div class="space-y-3">
              <div v-for="log in recentLogs.slice(0,3)" :key="log.id" class="flex items-start gap-4 p-4 bg-white rounded-[1.25rem] border border-gray-100 shadow-sm hover:border-gray-200 transition-colors group">
                <div class="w-10 h-10 rounded-2xl bg-blue-50 flex items-center justify-center text-blue-500 shrink-0 group-hover:scale-105 transition-transform">
                    <span class="material-icons text-[18px]">person_search</span>
                </div>
                <div class="flex-1 min-w-0 pt-0.5">
                    <div class="flex items-center flex-wrap gap-x-1 gap-y-1 text-sm">
                        <span class="font-bold text-slate-800">{{ log.viewer_name }}</span>
                        <span class="text-gray-500 font-medium">viewed</span>
                        <button @click="viewProfile(log.target_unique_id)" class="font-bold text-[#FF3D3D] hover:underline transition-colors truncate max-w-[150px] sm:max-w-xs">{{ log.target_name }}</button>
                    </div>
                     <div class="text-[10px] font-bold text-gray-400 mt-1 flex items-center gap-1">
                        <span class="material-icons text-[10px]">schedule</span> {{ formatTimeAgo(log.created_at) }}
                    </div>
                </div>
              </div>
              
              <div v-if="recentLogs.length === 0" class="flex flex-col items-center justify-center py-8 text-gray-400 bg-white rounded-[1.25rem] border border-gray-100">
                  <span class="material-icons text-3xl opacity-50 mb-2">data_array</span>
                  <p class="text-sm font-bold">No activity recorded</p>
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
  RadialLinearScale,
  Filler
} from 'chart.js';
import { Line, PolarArea } from 'vue-chartjs';

// Register Chart.js components
ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  PointElement,
  LineElement,
  ArcElement,
  RadialLinearScale,
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

// 2. Polar Area Chart (Blood Group Dist)
const polarAreaChartData = computed(() => {
    const labels = rawBloodGroupDist.value.map(item => item.blood_group);
    const data = rawBloodGroupDist.value.map(item => item.count);
    
    // Professional color palette for blood groups matching the app with transparency
    const backgroundColors = [
        'rgba(255, 61, 61, 0.75)', // A+   (Primary Red)
        'rgba(252, 165, 165, 0.75)', // A-   (Light Red)
        'rgba(79, 70, 229, 0.75)', // B+   (Indigo)
        'rgba(165, 180, 252, 0.75)', // B-   (Light Indigo)
        'rgba(5, 150, 105, 0.75)', // AB+  (Emerald)
        'rgba(209, 250, 229, 0.75)', // AB-  (Light Emerald)
        'rgba(2, 132, 199, 0.75)', // O+   (Sky Blue)
        'rgba(224, 242, 254, 0.75)'  // O-   (Light Sky)
    ];

    return {
        labels,
        datasets: [{
            data,
            backgroundColor: backgroundColors.slice(0, data.length),
            borderWidth: 1.5,
            borderColor: '#ffffff', // Clean white border 
            hoverOffset: 4
        }]
    };
});

const polarAreaChartOptions = {
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
            backgroundColor: 'rgba(255, 255, 255, 0.98)',
            titleColor: '#0F172A',
            bodyColor: '#334155',
            titleFont: { family: "'Plus Jakarta Sans', sans-serif", size: 14, weight: 'bold' as const },
            bodyFont: { family: "'Plus Jakarta Sans', sans-serif", size: 13, weight: 'normal' as const },
            padding: 16,
            cornerRadius: 12,
            displayColors: true,
            borderColor: '#E2E8F0',
            borderWidth: 1,
            boxPadding: 6,
            boxShadow: '0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)',
            callbacks: {
                label: function(context: any) {
                    let label = context.label || '';
                    if (label) {
                        label += ': ';
                    }
                    if (context.parsed !== null) {
                        label += context.parsed.r + ' donors';
                    }
                    return label;
                }
            }
        }
    },
    scales: {
        r: {
             ticks: {
                display: false,
            },
            grid: {
                color: '#F1F5F9',
                circular: true
            },
            border: { display: false },
             angleLines: {
                color: '#E2E8F0',
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
