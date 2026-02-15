<template>
  <div class="min-h-screen bg-[#F8F9FA] font-sans pb-6 pt-4">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
      
      <!-- Top Search Bar Container -->
      <div class="bg-white rounded-2xl shadow-sm border border-gray-100 p-3 mb-6">
        <div class="flex flex-col lg:flex-row gap-3 items-center">
          
          <!-- Location Search Input -->
          <div class="flex-1 w-full lg:w-auto relative group z-30" v-click-outside="closeLocationDropdown">
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <span class="material-icons text-gray-400 text-base">location_on</span>
            </div>
            <input 
                v-model="locationSearch"
                @focus="showLocationDropdown = true"
                type="text"
                placeholder="Search Division, District or Upazila..."
                class="w-full bg-gray-50 border border-gray-100 text-gray-700 text-sm font-semibold rounded-lg focus:ring-2 focus:ring-[#FF3D3D]/20 focus:border-[#FF3D3D] block pl-9 pr-9 py-2.5 transition-colors placeholder-gray-400"
            />
            <div v-if="locationSearch" @click="clearLocation" class="absolute inset-y-0 right-0 flex items-center px-3 cursor-pointer text-gray-400 hover:text-gray-600">
                <span class="material-icons text-xs">close</span>
            </div>
            <div v-else class="absolute inset-y-0 right-0 flex items-center px-3 pointer-events-none text-gray-400">
                <span class="material-icons text-xs">search</span>
            </div>

            <!-- Autocomplete Dropdown -->
            <div v-if="showLocationDropdown && filteredLocations.length > 0" class="absolute top-full left-0 right-0 mt-2 bg-white rounded-xl shadow-xl border border-gray-100 max-h-64 overflow-y-auto z-50">
                <ul>
                    <li 
                        v-for="(loc, index) in filteredLocations" 
                        :key="index"
                        @click="selectLocation(loc)"
                        class="px-4 py-2 hover:bg-red-50 cursor-pointer flex items-center gap-3 transition-colors border-b border-gray-50 last:border-0"
                    >
                        <div class="w-6 h-6 rounded-full bg-gray-50 flex items-center justify-center shrink-0 text-gray-400">
                             <span class="material-icons text-xs" v-if="loc.type === 'division'">map</span>
                             <span class="material-icons text-xs" v-else-if="loc.type === 'district'">location_city</span>
                             <span class="material-icons text-xs" v-else>holiday_village</span>
                        </div>
                        <div>
                            <p class="text-xs font-bold text-gray-800">{{ loc.name }}</p>
                            <p class="text-[10px] uppercase font-bold text-gray-400 tracking-wider">
                                {{ loc.type }} 
                                <span v-if="loc.subtitle" class="normal-case font-medium text-gray-300">• {{ loc.subtitle }}</span>
                            </p>
                        </div>
                    </li>
                </ul>
            </div>
            <div v-else-if="showLocationDropdown && locationSearch && filteredLocations.length === 0" class="absolute top-full left-0 right-0 mt-2 bg-white rounded-xl shadow-xl border border-gray-100 p-4 text-center z-50">
                <p class="text-xs text-gray-500">No locations found</p>
            </div>
          </div>

          <!-- Divider -->
          <div class="hidden lg:block w-px h-8 bg-gray-100 mx-1"></div>

           <!-- Blood Group Selection -->
           <div class="flex items-center gap-1.5 overflow-x-auto pb-2 lg:pb-0 w-full lg:w-auto scrollbar-hide">
              <button 
                  v-for="bg in bloodGroups" 
                  :key="bg"
                  type="button"
                  @click="filters.group = bg"
                  :class="[
                      'px-3 py-1.5 rounded-lg text-xs font-bold whitespace-nowrap transition-all border',
                      filters.group === bg 
                          ? 'bg-[#FF3D3D] text-white border-[#FF3D3D] shadow-md shadow-red-500/20' 
                          : 'bg-gray-50 text-gray-600 border-gray-100 hover:bg-gray-100 hover:border-gray-200'
                  ]"
              >
                  {{ bg }}
              </button>
           </div>

           <!-- Divider -->
          <div class="hidden lg:block w-px h-8 bg-gray-100 mx-1"></div>

           <!-- Availability Toggle -->
           <label class="flex items-center gap-2 cursor-pointer group select-none shrink-0 border border-gray-100 rounded-lg px-3 py-1.5 bg-gray-50 hover:bg-gray-100 transition-colors">
                <div class="relative">
                    <input type="checkbox" v-model="filters.availableOnly" class="sr-only peer">
                    <div class="w-7 h-4 bg-gray-300 peer-focus:outline-none rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-3 after:w-3 after:transition-all peer-checked:bg-[#FF3D3D]"></div>
                </div>
                <span class="text-[10px] font-bold text-gray-600 uppercase tracking-wide">Available</span>
            </label>

           <!-- Search Button -->
           <div class="w-full lg:w-auto">
             <button 
                @click="searchDonors"
                class="w-full bg-[#FF3D3D] text-white font-bold py-2.5 px-6 rounded-xl shadow-lg shadow-red-500/30 hover:shadow-red-500/40 hover:bg-red-600 transition-all transform active:scale-95 flex items-center justify-center gap-2 text-sm"
            >
                <span class="material-icons text-sm">search</span>
                <span>Search</span>
            </button>
           </div>

        </div>
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
import { useRouter } from 'vue-router';
// Import location data
import { divisions } from '@/lib/locations';

// State
const loading = ref(false);
const hasSearched = ref(false);
const donors = ref<any[]>([]);

const selectedDivision = ref('');
const locationSearch = ref('');
const showLocationDropdown = ref(false);
const filters = ref({
  district: '',
  upazila: '',
  group: '',
  gender: '', 
  availableOnly: true, // Default true
});

const bloodGroups = ['A+', 'A-', 'B+', 'B-', 'AB+', 'AB-', 'O+', 'O-'];

// Location Flattening & Search
const allLocations = computed(() => {
    const locs: any[] = [];
    divisions.forEach(div => {
        // Division
        locs.push({
            name: div.name,
            type: 'division',
            subtitle: 'Entire Division',
            division: div.name
        });
        
        div.districts.forEach(dist => {
            // District
            locs.push({
                name: dist.name,
                type: 'district',
                subtitle: div.name,
                division: div.name,
                district: dist.name
            });

            dist.upazilas.forEach(upz => {
                // Upazila
                locs.push({
                    name: upz,
                    type: 'upazila',
                    subtitle: `${dist.name}, ${div.name}`,
                    division: div.name,
                    district: dist.name,
                    upazila: upz
                });
            });
        });
    });
    return locs;
});

const filteredLocations = computed(() => {
    if (!locationSearch.value) return [];
    
    const query = locationSearch.value.toLowerCase().trim();
    return allLocations.value.filter(loc => 
        loc.name.toLowerCase().includes(query) || 
        loc.subtitle.toLowerCase().includes(query)
    ).slice(0, 10); // Limit results
});

const selectLocation = (loc: any) => {
    // Set internal filters
    selectedDivision.value = loc.division;
    filters.value.district = loc.district || '';
    filters.value.upazila = loc.upazila || '';

    // Set display text
    locationSearch.value = loc.name;
    // Optional: Add context to display text? e.g. "Savar, Dhaka"
    if (loc.type === 'upazila') locationSearch.value = `${loc.name}, ${loc.district}`;
    else if (loc.type === 'district') locationSearch.value = `${loc.name}, ${loc.division}`;

    showLocationDropdown.value = false;
};

const clearLocation = () => {
    locationSearch.value = '';
    selectedDivision.value = '';
    filters.value.district = '';
    filters.value.upazila = '';
    showLocationDropdown.value = false;
};

const closeLocationDropdown = () => {
    // Delay slightly to allow click to register
    setTimeout(() => {
        showLocationDropdown.value = false;
    }, 200);
};

// Click outside logic
const handleClickOutside = (event: MouseEvent) => {
    const target = event.target as HTMLElement;
    if (!target.closest('.group')) { // Assuming the container has class 'group' or similar unique identifier
        showLocationDropdown.value = false;
    }
};

// Register/Unregister click listener
const vClickOutside = {
    mounted(el: any, binding: any) {
        el.clickOutsideEvent = (event: Event) => {
            if (!(el === event.target || el.contains(event.target))) {
                binding.value(event);
            }
        };
        document.addEventListener('click', el.clickOutsideEvent);
    },
    unmounted(el: any) {
        document.removeEventListener('click', el.clickOutsideEvent);
    },
};

// Watchers
watch(selectedDivision, (newVal) => {
    if (!newVal) {
        // If division is cleared externally (shouldn't happen often with new UI), sync it
    }
});


// Methods
const fetchDonors = async () => {
  if (!filters.value.group) {
    alert("Please select a blood group to search.");
    return;
  }

  loading.value = true;
  hasSearched.value = true;
  try {
    const params: any = {};
    params.group = filters.value.group;
    
    if (filters.value.district) params.district = filters.value.district;
    if (filters.value.upazila) params.upazila = filters.value.upazila;
    // If Only division selected (no district), we might need to handle that if backend supports it.
    // Assuming backend filters by district/upazila primarily.
    // If division is selected but no district, we technically can't filter optimally unless backend supports "division" param.
    // For now, let's pass what we have.
    
    if (filters.value.gender) params.gender = filters.value.gender;
    params.available_only = filters.value.availableOnly;

    const res = await api.get('/donors', { params });
    donors.value = res.data;
  } catch (error) {
    console.error("Search failed", error);
  } finally {
    loading.value = false;
  }
};

const searchDonors = () => {
    fetchDonors();
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

// Formatting helpers
const getEstimatedAge = (donor: any) => {
    if (donor.birthday) {
        const birthDate = new Date(donor.birthday);
        const today = new Date();
        let years = today.getFullYear() - birthDate.getFullYear();
        const m = today.getMonth() - birthDate.getMonth();
        if (m < 0 || (m === 0 && today.getDate() < birthDate.getDate())) {
            years--;
        }
        return years;
    }
    // Fallback if no birthday (shouldn't happen for new users, but for old data)
    return 'N/A';
};

</script>
