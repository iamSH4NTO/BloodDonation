<template>
  <div class="space-y-6 lg:space-y-8 max-w-7xl mx-auto pb-10">
    
    <!-- Header Section -->
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-end gap-4 bg-white p-6 rounded-3xl shadow-sm border border-gray-100">
      <div>
        <h1 class="text-2xl sm:text-3xl font-black text-gray-900 tracking-tight">System Overview</h1>
        <p class="text-sm font-medium text-gray-500 mt-1">Monitor blood bank performance and user engagement</p>
      </div>
      <div class="flex items-center gap-3">
        <!-- Date Range Filter -->
        <select 
          v-model="dateRange" 
          class="bg-gray-50 border border-gray-100 text-gray-700 text-sm rounded-xl focus:ring-[#FF3D3D] focus:border-[#FF3D3D] block p-2 font-medium cursor-pointer shadow-sm hover:bg-gray-100 transition-colors outline-none"
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
        <button @click="fetchData" class="w-10 h-10 shrink-0 rouned-xl bg-gray-50 hover:bg-gray-100 flex items-center justify-center text-gray-500 hover:text-gray-900 transition-colors border border-gray-100 shadow-sm rounded-xl" title="Refresh">
          <span class="material-icons text-sm" :class="{'animate-spin': isLoading}">refresh</span>
        </button>
      </div>
    </div>

    <!-- Quick Stats Grid -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 sm:gap-6">
      
      <!-- Total Donors -->
      <div class="relative overflow-hidden bg-white rounded-3xl shadow-sm border border-gray-100 p-6 group hover:shadow-xl hover:shadow-red-500/5 transition-all duration-300 transform hover:-translate-y-1">
        <div class="absolute -right-6 -top-6 w-24 h-24 bg-red-50 rounded-full blur-2xl group-hover:bg-red-100 transition-colors"></div>
        <div class="relative">
            <div class="flex justify-between items-start mb-6">
                <div class="w-12 h-12 rounded-2xl bg-red-50 flex items-center justify-center text-[#FF3D3D] border border-red-100 group-hover:scale-110 group-hover:rotate-3 transition-transform duration-300 shadow-sm">
                    <span class="material-icons text-xl">people</span>
                </div>
                <!-- Growth Indicator -->
                <span v-if="stats.growth?.donors" class="inline-flex items-center gap-1 text-[11px] font-black px-2 py-1 rounded-lg" :class="parseFloat(stats.growth.donors) >= 0 ? 'bg-emerald-50 text-emerald-600 border border-emerald-100' : 'bg-red-50 text-red-600 border border-red-100'">
                    <span class="material-icons text-[12px]">{{ parseFloat(stats.growth.donors) >= 0 ? 'trending_up' : 'trending_down' }}</span> 
                    {{ Math.abs(parseFloat(stats.growth.donors)) }}%
                </span>
            </div>
            <div>
              <p class="text-xs font-bold text-gray-400 uppercase tracking-wider mb-1">Registered Donors</p>
              <h3 class="text-4xl font-black text-gray-900 tracking-tight">{{ stats.total_donors }}</h3>
            </div>
        </div>
      </div>

      <!-- Total Donations -->
      <div class="relative overflow-hidden bg-white rounded-3xl shadow-sm border border-gray-100 p-6 group hover:shadow-xl hover:shadow-orange-500/5 transition-all duration-300 transform hover:-translate-y-1">
        <div class="absolute -right-6 -top-6 w-24 h-24 bg-orange-50 rounded-full blur-2xl group-hover:bg-orange-100 transition-colors"></div>
        <div class="relative">
             <div class="flex justify-between items-start mb-6">
                <div class="w-12 h-12 rounded-2xl bg-orange-50 flex items-center justify-center text-orange-500 border border-orange-100 group-hover:scale-110 group-hover:-rotate-3 transition-transform duration-300 shadow-sm">
                    <span class="material-icons text-xl">volunteer_activism</span>
                </div>
                <!-- Growth Indicator -->
                <span v-if="stats.growth?.donations" class="inline-flex items-center gap-1 text-[11px] font-black px-2 py-1 rounded-lg" :class="parseFloat(stats.growth.donations) >= 0 ? 'bg-emerald-50 text-emerald-600 border border-emerald-100' : 'bg-red-50 text-red-600 border border-red-100'">
                    <span class="material-icons text-[12px]">{{ parseFloat(stats.growth.donations) >= 0 ? 'trending_up' : 'trending_down' }}</span> 
                    {{ Math.abs(parseFloat(stats.growth.donations)) }}%
                </span>
            </div>
            <div>
              <p class="text-xs font-bold text-gray-400 uppercase tracking-wider mb-1">Total Donations</p>
              <h3 class="text-4xl font-black text-gray-900 tracking-tight">{{ stats.total_donations }}</h3>
            </div>
        </div>
      </div>

       <!-- Lives Saved Estimate -->
      <div class="relative overflow-hidden bg-white rounded-3xl shadow-sm border border-gray-100 p-6 group hover:shadow-xl hover:shadow-pink-500/5 transition-all duration-300 transform hover:-translate-y-1">
        <div class="absolute -right-6 -top-6 w-24 h-24 bg-pink-50 rounded-full blur-2xl group-hover:bg-pink-100 transition-colors"></div>
        <div class="relative">
             <div class="flex justify-between items-start mb-6">
                <div class="w-12 h-12 rounded-2xl bg-pink-50 flex items-center justify-center text-pink-500 border border-pink-100 group-hover:scale-110 group-hover:rotate-12 transition-transform duration-300 shadow-sm">
                    <span class="material-icons text-xl">favorite</span>
                </div>
            </div>
            <div>
              <p class="text-xs font-bold text-gray-400 uppercase tracking-wider mb-1">Est. Lives Saved</p>
              <h3 class="text-4xl font-black text-pink-500 tracking-tight">{{ stats.total_donations * 3 }}</h3>
            </div>
        </div>
      </div>

      <!-- Total Users -->
      <div class="relative overflow-hidden bg-linear-to-br from-[#1e1e2d] to-[#2a2a3c] rounded-3xl shadow-lg p-6 group hover:shadow-2xl hover:shadow-gray-900/20 transition-all duration-300 transform hover:-translate-y-1">
        <div class="absolute right-0 top-0 h-full w-1/2 bg-linear-to-l from-white/10 to-transparent pointer-events-none rounded-r-3xl"></div>
        <div class="relative z-10">
             <div class="flex justify-between items-start mb-6">
                <div class="w-12 h-12 rounded-2xl bg-white/10 flex items-center justify-center text-white border border-white/20 backdrop-blur-sm group-hover:scale-110 transition-transform duration-300 shadow-sm">
                    <span class="material-icons text-xl">group</span>
                </div>
                <!-- Growth Indicator -->
                <span v-if="stats.growth?.users" class="inline-flex items-center gap-1 text-[11px] font-black px-2 py-1 rounded-lg bg-white/10 text-white border border-white/20 backdrop-blur-sm">
                    <span class="material-icons text-[12px]">{{ parseFloat(stats.growth.users) >= 0 ? 'trending_up' : 'trending_down' }}</span> 
                    {{ Math.abs(parseFloat(stats.growth.users)) }}%
                </span>
            </div>
            <div>
              <p class="text-xs font-bold text-gray-300 uppercase tracking-wider mb-1">Total Active Accounts</p>
              <h3 class="text-4xl font-black text-white tracking-tight">{{ stats.total_users }}</h3>
            </div>
        </div>
      </div>
    </div>

    <!-- Charts Section -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        
        <!-- Main Line Chart: Registrations vs Donations -->
        <div class="lg:col-span-2 bg-white rounded-3xl shadow-sm border border-gray-100 p-6 flex flex-col h-[400px]">
            <div class="flex justify-between items-center mb-6 shrink-0">
                <div>
                  <h2 class="text-lg font-black text-gray-900 tracking-tight">Activity Trends</h2>
                  <p class="text-[11px] font-bold text-gray-400 uppercase mt-0.5">Last 6 Months Comparison</p>
                </div>
                <div class="px-3 py-1.5 bg-gray-50 border border-gray-100 rounded-lg text-xs font-bold text-gray-500">
                  Monthly
                </div>
            </div>
            <div class="grow relative w-full h-full pb-2">
                <Line v-if="chartDataLoaded" :data="lineChartData" :options="lineChartOptions" />
                <div v-else class="absolute inset-0 flex items-center justify-center text-gray-400">
                    <span class="material-icons animate-spin mr-2">refresh</span> Loading chart...
                </div>
            </div>
        </div>

        <!-- Doughnut Chart: Blood Group Dist -->
        <div class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 flex flex-col h-[400px]">
             <div class="flex justify-between items-center mb-4 shrink-0">
                <div>
                  <h2 class="text-lg font-black text-gray-900 tracking-tight">Blood Inventory</h2>
                  <p class="text-[11px] font-bold text-gray-400 uppercase mt-0.5">Donor Distribution by Type</p>
                </div>
                <span class="material-icons text-red-200 text-3xl">bloodtype</span>
            </div>
            <div class="grow relative w-full h-full flex items-center justify-center pb-4">
                 <Doughnut v-if="chartDataLoaded" :data="doughnutChartData" :options="doughnutChartOptions" />
                 <div v-else class="absolute inset-0 flex items-center justify-center text-gray-400">
                    <span class="material-icons animate-spin mr-2">refresh</span> Loading chart...
                </div>
            </div>
        </div>
    </div>
    
    <!-- Bottom Grid -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        
        <!-- Activity Timeline -->
        <div class="lg:col-span-2 bg-white rounded-3xl shadow-sm border border-gray-100 p-6">
            <div class="flex justify-between items-center mb-6">
                <div>
                  <h2 class="text-lg font-black text-gray-900 tracking-tight">Recent Activity Feed</h2>
                  <p class="text-[11px] font-bold text-gray-400 uppercase mt-0.5">Latest system interactions</p>
                </div>
                <router-link to="/admin/logs" class="text-xs font-bold text-[#FF3D3D] hover:text-red-700 bg-red-50 hover:bg-red-100 px-4 py-2 rounded-xl transition-colors flex items-center gap-1 shadow-sm border border-red-50">
                View detailed logs <span class="material-icons text-sm">arrow_forward</span>
                </router-link>
            </div>
            
            <div class="space-y-4">
                <div v-for="log in recentLogs" :key="log.id" class="flex items-start gap-4 p-4 bg-gray-50/50 rounded-2xl hover:bg-gray-50 transition-colors border border-transparent hover:border-gray-100">
                <div class="w-10 h-10 rounded-xl bg-blue-50 flex items-center justify-center text-blue-500 border border-blue-100 shrink-0 shadow-sm">
                    <span class="material-icons text-[18px]">person_search</span>
                </div>
                <div class="flex-1 min-w-0 pt-0.5">
                    <div class="flex items-center flex-wrap gap-x-1 gap-y-1 text-sm">
                        <span class="font-bold text-gray-900">{{ log.viewer_name }}</span>
                        <span class="text-gray-400 font-medium">viewed the profile of</span>
                        <button @click="viewProfile(log.target_unique_id)" class="font-bold text-[#FF3D3D] hover:underline hover:text-red-700 transition-colors truncate max-w-[150px] sm:max-w-xs">{{ log.target_name }}</button>
                    </div>
                    <div class="flex items-center gap-2 mt-1.5">
                        <span class="inline-flex items-center gap-1 text-[10px] font-bold px-2 py-0.5 rounded-md bg-white border border-gray-200 text-gray-500">
                            <span class="material-icons text-[10px] text-gray-400">schedule</span> {{ formatTimeAgo(log.created_at) }}
                        </span>
                        <span class="text-[10px] font-mono text-gray-400">ID: {{ log.target_unique_id }}</span>
                    </div>
                </div>
                </div>
                
                <div v-if="recentLogs.length === 0" class="flex flex-col items-center justify-center py-10 text-gray-400">
                   <div class="w-16 h-16 rounded-full bg-gray-50 flex items-center justify-center mb-3 border border-gray-100 shadow-inner">
                       <span class="material-icons text-3xl opacity-50">data_array</span>
                   </div>
                   <p class="text-sm font-bold">No activity recorded yet</p>
                </div>
            </div>
        </div>

        <!-- Quick Actions Panel -->
        <div class="bg-indigo-50/50 rounded-3xl shadow-sm border border-indigo-100 p-6 flex flex-col justify-between relative overflow-hidden group">
             <div class="absolute -right-8 -bottom-8 w-40 h-40 bg-indigo-500/10 rounded-full blur-3xl group-hover:bg-indigo-500/20 transition-colors"></div>
             <div class="relative z-10 space-y-6">
                 <div>
                    <div class="w-12 h-12 rounded-2xl bg-indigo-100 text-indigo-600 flex items-center justify-center mb-4 shadow-sm border border-indigo-200">
                        <span class="material-icons">flash_on</span>
                    </div>
                    <h2 class="text-xl font-black text-gray-900 tracking-tight">Admin Actions</h2>
                    <p class="text-xs font-medium text-gray-500 mt-2 leading-relaxed">Instantly manage platform users and review flagged activities.</p>
                 </div>
                 
                 <div class="space-y-3">
                     <router-link to="/admin/donors" class="w-full bg-white hover:bg-gray-50 text-indigo-700 border border-indigo-100 px-4 py-3.5 rounded-2xl font-bold text-sm shadow-sm transition-all flex items-center justify-between group/btn">
                         <div class="flex items-center gap-3">
                             <div class="w-8 h-8 rounded-lg bg-indigo-50 flex items-center justify-center"><span class="material-icons text-[16px]">manage_accounts</span></div>
                             Manage Users
                         </div>
                         <span class="material-icons text-sm text-indigo-300 group-hover/btn:translate-x-1 transition-transform">arrow_forward</span>
                     </router-link>
                     <button @click="$router.push('/admin/logs')" class="w-full bg-white hover:bg-gray-50 text-gray-700 border border-gray-200 px-4 py-3.5 rounded-2xl font-bold text-sm shadow-sm transition-all flex items-center justify-between group/btn">
                         <div class="flex items-center gap-3">
                             <div class="w-8 h-8 rounded-lg bg-gray-50 flex items-center justify-center"><span class="material-icons text-[16px] text-gray-500">security</span></div>
                             Review Security Logs
                         </div>
                         <span class="material-icons text-sm text-gray-300 group-hover/btn:translate-x-1 transition-transform">arrow_forward</span>
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
                backgroundColor: 'rgba(255, 61, 61, 0.1)',
                borderColor: '#FF3D3D',
                pointBackgroundColor: '#ffffff',
                pointBorderColor: '#FF3D3D',
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
                backgroundColor: 'rgba(34, 197, 94, 0.05)',
                borderColor: '#22C55E',
                borderDash: [5, 5],
                pointBackgroundColor: '#ffffff',
                pointBorderColor: '#22C55E',
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
                font: {
                    family: "'Inter', sans-serif",
                    weight: 'bold' as const,
                    size: 11
                },
                color: '#9CA3AF'
            }
        },
        tooltip: {
            backgroundColor: 'rgba(17, 24, 39, 0.9)',
            titleFont: { family: "'Inter', sans-serif", size: 13 },
            bodyFont: { family: "'Inter', sans-serif", size: 12 },
            padding: 12,
            cornerRadius: 8,
            displayColors: true,
            boxPadding: 4
        }
    },
    scales: {
        y: {
            beginAtZero: true,
            grid: {
                color: '#F3F4F6',
                drawBorder: false,
            },
            ticks: {
                font: { family: "'Inter', sans-serif", size: 11 },
                color: '#9CA3AF',
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
                font: { family: "'Inter', sans-serif", size: 11, weight: 'bold' as const },
                color: '#6B7280',
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
    
    // Aesthetic color palette for blood groups
    const backgroundColors = [
        '#FF3D3D', // A+ 
        '#FCA5A5', // A-
        '#3B82F6', // B+
        '#93C5FD', // B-
        '#8B5CF6', // AB+
        '#C4B5FD', // AB-
        '#10B981', // O+
        '#6EE7B7'  // O-
    ];

    return {
        labels,
        datasets: [{
            data,
            backgroundColor: backgroundColors.slice(0, data.length),
            borderWidth: 0,
            hoverOffset: 4
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
                    family: "'Inter', sans-serif",
                    weight: 'bold' as const,
                    size: 11
                },
                color: '#6B7280'
            }
        },
        tooltip: {
            backgroundColor: 'rgba(17, 24, 39, 0.9)',
            titleFont: { family: "'Inter', sans-serif", size: 13 },
            bodyFont: { family: "'Inter', sans-serif", size: 12 },
            padding: 12,
            cornerRadius: 8,
            displayColors: true,
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
