<template>
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
                       <span class="material-icons text-xs" v-else-if="loc.type === 'district' || loc.type === 'city'">location_city</span>
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
          @click="emitSearch"
          class="w-full bg-[#FF3D3D] text-white font-bold py-2.5 px-6 rounded-xl shadow-lg shadow-red-500/30 hover:shadow-red-500/40 hover:bg-red-600 transition-all transform active:scale-95 flex items-center justify-center gap-2 text-sm"
      >
          <span class="material-icons text-sm">search</span>
          <span>Search</span>
      </button>
     </div>

  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted } from 'vue';
import { divisions } from '@/lib/locations';
import api from '@/lib/axios';

const props = defineProps<{
  initialFilters?: {
    district: string;
    upazila: string;
    group: string;
    availableOnly: boolean;
    locationQuery?: string; // Optional for preserving search text
  }
}>();

const emit = defineEmits<{
  (e: 'search', filters: any): void
}>();

const selectedDivision = ref('');
const locationSearch = ref('');
const showLocationDropdown = ref(false);
const backendLocations = ref<any[]>([]);
const selectedLocationObj = ref<any>(null);
const filters = ref({
  district: '',
  upazila: '',
  group: '',
  availableOnly: false,
});

const bloodGroups = ['A+', 'A-', 'B+', 'B-', 'AB+', 'AB-', 'O+', 'O-'];

// Initialize logic
const initializeFromProps = () => {
  if (props.initialFilters) {
    filters.value.group = props.initialFilters.group || '';
    filters.value.district = props.initialFilters.district || '';
    filters.value.upazila = props.initialFilters.upazila || '';
    
    if (props.initialFilters.availableOnly !== undefined) {
        filters.value.availableOnly = props.initialFilters.availableOnly;
    }
    
    // Reconstruct location text
    if (props.initialFilters.locationQuery) {
        locationSearch.value = props.initialFilters.locationQuery;
    } else if (filters.value.upazila) {
         const loc = allLocations.value.find(l => l.upazila === filters.value.upazila && l.district === filters.value.district);
         if (loc) locationSearch.value = loc.name;
    } else if (filters.value.district) {
         const loc = allLocations.value.find(l => l.district === filters.value.district && l.type === 'district');
         if (loc) locationSearch.value = loc.name;
    }
  }
};

onMounted(() => {
  initializeFromProps();
});

watch(() => props.initialFilters, () => {
    initializeFromProps();
}, { deep: true });

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

const fetchBackendLocations = async (query: string) => {
    if (query.length < 2) {
        backendLocations.value = [];
        return;
    }
    try {
        const res = await api.get('/donors/locations/search', { params: { q: query } });
        backendLocations.value = res.data;
    } catch (error) {
        console.error("Failed to fetch locations", error);
    }
};

watch(locationSearch, (newVal) => {
    if (showLocationDropdown.value) {
        fetchBackendLocations(newVal);
    }
    // If the user manually changed the text away from the selected location, clear the object
    if (selectedLocationObj.value && newVal !== selectedLocationObj.value.name && !newVal.startsWith(selectedLocationObj.value.name)) {
        selectedLocationObj.value = null;
    }
});

const filteredLocations = computed(() => {
    const query = locationSearch.value.toLowerCase().trim();
    
    // Static locations
    const staticLocs = query ? allLocations.value.filter(loc => 
        loc.name.toLowerCase().includes(query) || 
        loc.subtitle.toLowerCase().includes(query)
    ).slice(0, 5) : [];

    // Backend locations
    const dynamicLocs = backendLocations.value;

    // Merge and remove duplicates
    const combined = [...staticLocs, ...dynamicLocs];
    const seen = new Set();
    return combined.filter(loc => {
        const key = `${loc.name}-${loc.type}-${loc.district}`;
        if (seen.has(key)) return false;
        seen.add(key);
        return true;
    }).slice(0, 10);
});

const selectLocation = (loc: any) => {
    // Set internal filters
    filters.value.district = loc.district || '';
    filters.value.upazila = loc.upazila || '';
    selectedLocationObj.value = loc;
    
    // Set display text
    locationSearch.value = loc.name;

    showLocationDropdown.value = false;
};

const clearLocation = () => {
    locationSearch.value = '';
    selectedDivision.value = '';
    filters.value.district = '';
    filters.value.upazila = '';
    selectedLocationObj.value = null; // Clear selected object
    showLocationDropdown.value = false;
};

const closeLocationDropdown = () => {
    setTimeout(() => {
        showLocationDropdown.value = false;
    }, 200);
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

const emitSearch = () => {
    // If we have a selected object, use its base name for the query
    // Otherwise use the raw search text
    const query = selectedLocationObj.value && locationSearch.value.includes(selectedLocationObj.value.name)
        ? selectedLocationObj.value.name 
        : locationSearch.value;

    emit('search', {
        ...filters.value,
        locationQuery: query
    });
};
</script>

<style scoped>
.scrollbar-hide::-webkit-scrollbar {
    display: none;
}
.scrollbar-hide {
    -ms-overflow-style: none;
    scrollbar-width: none;
}
</style>
