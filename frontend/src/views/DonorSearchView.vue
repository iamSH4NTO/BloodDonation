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
      <div class="bg-white rounded-3xl p-6 min-h-[400px] border border-dashed border-gray-200 relative">
          
          <!-- Loading State -->
          <div v-if="loading" class="absolute inset-0 flex items-center justify-center bg-white/80 z-10 rounded-3xl">
             <div class="flex flex-col items-center gap-3">
                 <div class="animate-spin rounded-full h-10 w-10 border-4 border-gray-100 border-t-[#FF3D3D]"></div>
                 <p class="text-gray-500 text-sm font-medium animate-pulse">Searching for donors...</p>
             </div>
          </div>

          <!-- Empty State (No Search Yet) -->
          <div v-if="!hasSearched" class="h-full flex flex-col items-center justify-center py-16 text-center">
              <div class="w-20 h-20 bg-red-50 rounded-full flex items-center justify-center mb-4 animate-float">
                  <span class="material-icons text-4xl text-[#FF3D3D]/40">person_search</span>
              </div>
              <h3 class="text-lg font-bold text-gray-800 mb-1">Ready to find a donor?</h3>
              <p class="text-gray-400 text-sm max-w-xs mx-auto mb-8">Enter location and blood group to start your search.</p>
              
              <div class="flex flex-wrap justify-center items-center gap-3 md:gap-6 text-xs font-medium text-gray-400">
                  <div class="flex items-center gap-1.5 bg-gray-50 px-3 py-1.5 rounded-md">
                      <span class="material-icons text-[#FF3D3D] text-xs">location_on</span>
                      <span>Select Location</span>
                  </div>
                  <span class="material-icons text-gray-300 text-xs">arrow_forward</span>
                  <div class="flex items-center gap-1.5 bg-gray-50 px-3 py-1.5 rounded-md">
                       <span class="material-icons text-[#FF3D3D] text-xs">bloodtype</span>
                      <span>Pick Blood Group</span>
                  </div>
                  <span class="material-icons text-gray-300 text-xs">arrow_forward</span>
                  <div class="flex items-center gap-1.5 px-3 py-1.5 rounded-md border border-red-100 bg-red-50 text-[#FF3D3D]">
                       <span class="material-icons text-xs">search</span>
                      <span>Click Search</span>
                  </div>
              </div>
          </div>

          <!-- Empty State (No Results) -->
          <div v-else-if="donors.length === 0 && !loading" class="h-full flex flex-col items-center justify-center py-16 text-center">
             <div class="w-16 h-16 bg-gray-50 rounded-full flex items-center justify-center mb-3">
                  <span class="material-icons text-3xl text-gray-400">search_off</span>
              </div>
              <h3 class="text-base font-bold text-gray-800">No donors found</h3>
              <p class="text-xs text-gray-500">Try adjusting your filters or expanding your search area.</p>
          </div>

          <!-- Results Grid -->
          <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
             <div v-for="donor in donors" :key="donor.user_id" class="bg-white rounded-xl border border-gray-100 p-4 hover:shadow-lg hover:-translate-y-1 transition-all duration-300 group">
                 <div class="flex items-start gap-3 mb-3">
                     <div class="relative shrink-0">
                         <img :src="`https://i.pravatar.cc/150?u=${donor.user_id}`" class="w-12 h-12 rounded-lg object-cover bg-gray-100 shadow-inner" />
                         <span class="absolute -bottom-1.5 -right-1.5 bg-white text-[#FF3D3D] text-[10px] font-black px-1.5 py-0.5 rounded shadow-sm border border-gray-100">
                             {{ donor.blood_group }}
                         </span>
                     </div>
                     <div>
                         <h4 class="font-bold text-gray-900 leading-tight text-sm">{{ donor.name }}</h4>
                         <p class="text-[10px] text-gray-500 mt-1 flex items-center gap-1">
                             <span class="w-1.5 h-1.5 rounded-full" :class="donor.is_available ? 'bg-green-500' : 'bg-gray-300'"></span>
                             {{ donor.is_available ? 'Available' : 'Unavailable' }}
                         </p>
                          <div class="flex gap-1 mt-1.5">
                             <span v-if="donor.gender" class="px-1.5 py-0.5 bg-gray-50 rounded text-[9px] font-bold text-gray-600 uppercase tracking-wide">{{ donor.gender }}</span>
                             <span v-if="donor.privacy_settings?.verified" class="px-1.5 py-0.5 bg-blue-50 rounded text-[9px] font-bold text-blue-600 uppercase tracking-wide">Verified</span>
                         </div>
                     </div>
                 </div>

                 <div class="space-y-1.5 mb-4">
                     <div class="flex items-center gap-1.5 text-[11px] font-semibold text-gray-500">
                         <span class="material-icons text-gray-300 text-xs">location_on</span>
                         <span class="truncate">{{ donor.district }}{{ donor.city ? `, ${donor.city}` : '' }}</span>
                     </div>
                      <div class="flex items-center gap-1.5 text-[11px] font-semibold text-gray-500">
                         <span class="material-icons text-gray-300 text-xs">history</span>
                         <span>Last: <span class="text-gray-700">{{ donor.last_donation_date ? formatDateRelative(donor.last_donation_date) : 'Never' }}</span></span>
                     </div>
                 </div>

                 <button @click="$router.push(`/donors/${donor.user_id}`)" class="w-full bg-white border border-gray-200 text-gray-700 font-bold py-2 rounded-lg hover:bg-[#FF3D3D] hover:text-white hover:border-[#FF3D3D] transition-all flex items-center justify-center gap-2 text-xs group-hover:shadow-md group-hover:shadow-red-500/10">
                     <span>Contact Donor</span>
                     <span class="material-icons text-xs opacity-60">arrow_forward</span>
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

// State
const loading = ref(false);
const hasSearched = ref(false);
const donors = ref<any[]>([]);
const route = useRoute();
const router = useRouter();

// Filters state to sync with URL or component
const filters = ref({
  district: '',
  upazila: '',
  group: '',
  availableOnly: true,
  locationQuery: ''
});

// Methods
const fetchDonors = async (searchFilters: any) => {
  if (!searchFilters.group) {
    alert("Please select a blood group to search.");
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

    const res = await api.get('/donors', { params });
    donors.value = res.data;
  } catch (error) {
    console.error("Search failed", error);
  } finally {
    loading.value = false;
  }
};

const handleSearch = (searchFilters: any) => {
    // Update local state
    filters.value = { ...filters.value, ...searchFilters };
    
    // Update URL to reflect search state (optional but good for UX)
    router.replace({
        query: {
            ...route.query, // preserve other params if any
            group: searchFilters.group,
            district: searchFilters.district,
            upazila: searchFilters.upazila,
            available: searchFilters.availableOnly.toString(),
            q: searchFilters.locationQuery
        }
    });

    fetchDonors(searchFilters);
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
            availableOnly: query.available !== 'false',
            locationQuery: (query.q as string) || ''
        };

        // Auto search if group is present
        fetchDonors(filters.value);
    }
});
</script>
