<template>
  <div class="min-h-screen bg-[#F8F9FA] font-sans pb-6 pt-4">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
      
      <!-- Top Search Bar Container -->
      <div class="bg-white rounded-2xl shadow-sm border border-gray-100 p-3 mb-6">
        <DonorSearchForm 
          @search="handleSearch"
          :initial-filters="filters"
        />
      </div>

       <!-- Results Header -->
       <div class="mb-4">
        <h2 class="text-xl font-bold text-gray-800">Search Results</h2>
        <p class="text-xs text-gray-500">Please perform a search to see donors</p>
      </div>

      <!-- Main Content Area -->
      <div class="bg-white rounded-4xl sm:rounded-[2.5rem] p-6 sm:p-10 min-h-[400px] border border-dashed border-gray-200 relative">
          
          <!-- Loading State -->
          <div v-if="loading" class="absolute inset-0 flex items-center justify-center bg-white/80 z-10 rounded-4xl sm:rounded-[2.5rem]">
             <div class="flex flex-col items-center gap-4">
                 <div class="animate-spin rounded-full h-12 w-12 border-4 border-gray-100 border-t-[#FF3D3D]"></div>
                 <p class="text-gray-500 text-sm font-bold animate-pulse">Searching for donors...</p>
             </div>
          </div>

          <!-- Empty State (No Search Yet) -->
          <div v-if="!hasSearched" class="h-full flex flex-col items-center justify-center py-12 sm:py-20 text-center">
              <div class="w-24 h-24 bg-red-50 rounded-full flex items-center justify-center mb-6 animate-float">
                  <span class="material-icons text-5xl text-[#FF3D3D]/40">person_search</span>
              </div>
              <h3 class="text-xl font-bold text-gray-800 mb-2">Ready to find a donor?</h3>
              <p class="text-gray-400 text-sm max-w-xs mx-auto mb-10 leading-relaxed font-medium">Enter location and blood group to start your search.</p>
              
              <div class="flex flex-col sm:flex-row justify-center items-center gap-4 sm:gap-6 text-xs font-bold text-gray-400">
                  <div class="flex items-center gap-2 bg-gray-50 px-4 py-2 rounded-xl">
                      <span class="material-icons text-[#FF3D3D] text-sm">location_on</span>
                      <span>Select Location</span>
                  </div>
                  <span class="material-icons text-gray-300 text-sm rotate-90 sm:rotate-0">arrow_forward</span>
                  <div class="flex items-center gap-2 bg-gray-50 px-4 py-2 rounded-xl">
                       <span class="material-icons text-[#FF3D3D] text-sm">bloodtype</span>
                      <span>Pick Blood Group</span>
                  </div>
                  <span class="material-icons text-gray-300 text-sm rotate-90 sm:rotate-0">arrow_forward</span>
                  <div class="flex items-center gap-2 px-4 py-2 rounded-xl border border-red-100 bg-red-50 text-[#FF3D3D]">
                       <span class="material-icons text-sm">search</span>
                      <span>Click Search</span>
                  </div>
              </div>
          </div>

          <!-- Empty State (No Results) -->
          <div v-else-if="donors.length === 0 && !loading" class="h-full flex flex-col items-center justify-center py-20 text-center">
             <div class="w-20 h-20 bg-gray-50 rounded-full flex items-center justify-center mb-4">
                  <span class="material-icons text-4xl text-gray-300">search_off</span>
              </div>
              <h3 class="text-lg font-bold text-gray-800 mb-1">No donors found</h3>
              <p class="text-sm text-gray-500 leading-relaxed">Try adjusting your filters or expanding your search area.</p>
          </div>

          <!-- Results Grid -->
          <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
             <div v-for="donor in donors" :key="donor.user_id" class="bg-white rounded-2xl border border-gray-100 p-5 hover:shadow-xl hover:-translate-y-1 transition-all duration-300 group">
                 <div class="flex items-start gap-4 mb-4">
                     <div class="relative shrink-0">
                         <img :src="`https://i.pravatar.cc/150?u=${donor.user_id}`" class="w-14 h-14 rounded-xl object-cover bg-gray-100 shadow-inner" />
                         <span class="absolute -bottom-2 -right-2 bg-[#FF3D3D] text-white text-[11px] font-black px-2 py-0.5 rounded shadow-lg border-2 border-white">
                             {{ donor.blood_group }}
                         </span>
                     </div>
                     <div class="flex-1 overflow-hidden">
                         <h4 class="font-bold text-gray-900 leading-tight text-base truncate">{{ donor.name }}</h4>
                         <p class="text-xs text-gray-500 mt-1 flex items-center gap-1.5">
                             <span class="w-1.5 h-1.5 rounded-full" :class="donor.is_available ? 'bg-green-500' : 'bg-gray-300'"></span>
                             {{ donor.is_available ? 'Available' : 'Unavailable' }}
                         </p>
                          <div class="flex flex-wrap gap-1.5 mt-2">
                             <span v-if="donor.gender" class="px-2 py-0.5 bg-gray-50 border border-gray-100 rounded-md text-[10px] font-bold text-gray-600 uppercase tracking-wide">{{ donor.gender }}</span>
                             <span v-if="donor.privacy_settings?.verified" class="px-2 py-0.5 bg-blue-50 border border-blue-100 rounded-md text-[10px] font-bold text-blue-600 uppercase tracking-wide">Verified</span>
                         </div>
                     </div>
                 </div>

                 <div class="space-y-2 mb-6">
                     <div class="flex items-start gap-2 text-xs font-semibold text-gray-600 leading-relaxed">
                         <span class="material-icons text-gray-400 text-sm mt-0.5">location_on</span>
                         <span class="line-clamp-2">
                             {{ [donor.area_village, donor.city, donor.upazila, donor.district].filter(Boolean).join(', ') }}
                         </span>
                     </div>
                      <div class="flex items-center gap-2 text-xs font-semibold text-gray-600">
                         <span class="material-icons text-gray-400 text-sm">history</span>
                         <span>Last: <span class="text-gray-900 font-bold">{{ donor.last_donation_date ? formatDateRelative(donor.last_donation_date) : 'Never' }}</span></span>
                     </div>
                 </div>

              <button @click="$router.push(`/donors/${donor.user_id}`)" class="w-full bg-white border-2 border-gray-100 text-gray-700 font-black py-3 rounded-xl hover:bg-[#FF3D3D] hover:text-white hover:border-[#FF3D3D] transition-all flex items-center justify-center gap-2 text-xs group-hover:shadow-lg group-hover:shadow-red-500/20">
                     <span>View Full Profile</span>
                     <span class="material-icons text-sm opacity-60">arrow_forward</span>
                 </button>
             </div>
          </div>

          <!-- Pagination UI -->
          <div v-if="donors.length > 0 && totalDonors > itemsPerPage" class="mt-12 flex flex-col sm:flex-row items-center justify-between gap-6 pt-8 border-t border-gray-100">
              <div class="text-xs font-bold text-gray-400 order-2 sm:order-1">
                  Showing <span class="text-gray-900">{{ (currentPage - 1) * itemsPerPage + 1 }}</span> to 
                  <span class="text-gray-900">{{ Math.min(currentPage * itemsPerPage, totalDonors) }}</span> of 
                  <span class="text-gray-900">{{ totalDonors }}</span> donors
              </div>
              
              <div class="flex items-center gap-2 order-1 sm:order-2">
                  <button 
                      @click="handlePageChange(currentPage - 1)"
                      :disabled="currentPage === 1"
                      class="p-2 rounded-xl border-2 border-gray-100 text-gray-400 hover:text-[#FF3D3D] hover:border-[#FF3D3D] disabled:opacity-30 disabled:hover:text-gray-400 disabled:hover:border-gray-100 transition-all"
                  >
                      <span class="material-icons">chevron_left</span>
                  </button>
                  
                  <div class="flex items-center gap-1">
                      <button 
                          v-for="p in Math.ceil(totalDonors / itemsPerPage)" 
                          :key="p"
                          @click="handlePageChange(p)"
                          :class="p === currentPage ? 'bg-[#FF3D3D] text-white border-[#FF3D3D]' : 'bg-white text-gray-400 border-gray-100 hover:border-[#FF3D3D] hover:text-[#FF3D3D]'"
                          class="w-10 h-10 rounded-xl border-2 font-black text-xs transition-all flex items-center justify-center"
                      >
                          {{ p }}
                      </button>
                  </div>

                  <button 
                      @click="handlePageChange(currentPage + 1)"
                      :disabled="currentPage >= Math.ceil(totalDonors / itemsPerPage)"
                      class="p-2 rounded-xl border-2 border-gray-100 text-gray-400 hover:text-[#FF3D3D] hover:border-[#FF3D3D] disabled:opacity-30 disabled:hover:text-gray-400 disabled:hover:border-gray-100 transition-all"
                  >
                      <span class="material-icons">chevron_right</span>
                  </button>
              </div>
          </div>

      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, computed, onMounted, onUnmounted } from 'vue';
import api from '@/lib/axios';
import { useRouter, useRoute } from 'vue-router';
import DonorSearchForm from '@/components/DonorSearchForm.vue';
import { useToastStore } from '@/stores/toast';

const toastStore = useToastStore();

// State
const loading = ref(false);
const hasSearched = ref(false);
const donors = ref<any[]>([]);
const totalDonors = ref(0);
const currentPage = ref(1);
const itemsPerPage = ref(12); // Matching grid layout (3 or 4 per row)
const route = useRoute();
const router = useRouter();

// Filters state to sync with URL or component
const filters = ref({
  district: '',
  upazila: '',
  group: '',
  availableOnly: false,
  locationQuery: ''
});

// Methods
const fetchDonors = async (searchFilters: any, page = 1) => {
  if (!searchFilters.group) {
    toastStore.show("Please select a blood group to search.", "info");
    return;
  }

  loading.value = true;
  hasSearched.value = true;
  donors.value = []; // Clear previous results

  try {
    const params: any = {};
    params.group = searchFilters.group;
    
    if (searchFilters.district) params.district = searchFilters.district;
    if (searchFilters.upazila) params.upazila = searchFilters.upazila;
    
    if (searchFilters.gender) params.gender = searchFilters.gender;
    params.available_only = searchFilters.availableOnly;
    if (searchFilters.locationQuery) params.q = searchFilters.locationQuery;
    
    params.page = page;
    params.limit = itemsPerPage.value;

    const res = await api.get('/donors', { params });
    // Handle paginated response
    if (res.data.donors) {
        donors.value = res.data.donors;
        totalDonors.value = res.data.total;
        currentPage.value = res.data.page;
    } else {
        // Fallback for non-paginated (unlikely now)
        donors.value = res.data;
        totalDonors.value = res.data.length;
    }
  } catch (error) {
    console.error("Search failed", error);
  } finally {
    loading.value = false;
  }
};

const handleSearch = (searchFilters: any) => {
    // Update local state
    filters.value = { ...filters.value, ...searchFilters };
    currentPage.value = 1; // Reset to first page on new search
    
    // Update URL to reflect search state
    router.replace({
        query: {
            ...route.query,
            group: searchFilters.group,
            district: searchFilters.district,
            upazila: searchFilters.upazila,
            available: searchFilters.availableOnly.toString(),
            q: searchFilters.locationQuery,
            page: '1'
        }
    });

    fetchDonors(searchFilters, 1);
};

const handlePageChange = (page: number) => {
    if (page < 1 || page > Math.ceil(totalDonors.value / itemsPerPage.value)) return;
    
    currentPage.value = page;
    
    // Update URL
    router.replace({
        query: {
            ...route.query,
            page: page.toString()
        }
    });

    fetchDonors(filters.value, page);
};

const formatDateRelative = (dateString: string) => {
    const date = new Date(dateString);
    const now = new Date();
    const diffTime = Math.abs(now.getTime() - date.getTime());
    const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24));
    
    if (diffDays < 30) return `${diffDays} days ago`;
    if (diffDays < 365) return `${Math.floor(diffDays / 30)} mos ago`;
    return `${Math.floor(diffDays / 365)} yr ago`;
};

// Initialize from URL query params
onMounted(() => {
    const query = route.query;
    if (query.group) {
        filters.value = {
            district: (query.district as string) || '',
            upazila: (query.upazila as string) || '',
            group: query.group as string,
            availableOnly: query.available === 'true',
            locationQuery: (query.q as string) || ''
        };
        
        const page = parseInt(query.page as string) || 1;
        currentPage.value = page;

        // Auto search if group is present
        fetchDonors(filters.value, page);
    }
});
</script>
