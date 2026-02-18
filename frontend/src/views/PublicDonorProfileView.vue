<template>
  <div class="min-h-screen bg-[#FAFAFA] font-sans pt-4 pb-4">
    
    <!-- Top Decoration -->

    <div v-if="loading" class="flex justify-center items-center h-64">
        <div class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-red-600"></div>
    </div>

    <div v-else-if="profile" class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8">
      
      <!-- 1. Profile Header Card -->
      <div class="bg-white rounded-3xl p-6 sm:p-8 shadow-sm border border-gray-100 relative overflow-visible">
         <div class="flex flex-col md:flex-row gap-6 md:gap-8 items-center md:items-center text-center md:text-left">
            <!-- Avatar -->
            <div class="relative group">
                <UserAvatar 
                  :src="profile.profile_picture" 
                  :gender="profile.gender" 
                  :name="profile.name" 
                  size="xl" 
                  class="shadow-xl"
                />
                <div class="absolute -bottom-2 -right-2 bg-[#FF3D3D] text-white w-10 h-10 md:w-12 md:h-12 rounded-full flex items-center justify-center font-bold text-base md:text-lg shadow-md border-4 border-white">
                    {{ profile.blood_group || '?' }}
                </div>
            </div>

            <!-- Info -->
            <div class="flex-1 space-y-2 w-full">
                <div class="flex items-center justify-center md:justify-start gap-3 flex-wrap">
                    <h1 class="text-2xl md:text-3xl font-black text-gray-900 tracking-tight">{{ profile.name }}</h1>
                    <span class="bg-green-100 text-green-700 px-3 py-1 rounded-full text-[10px] font-bold uppercase tracking-wide flex items-center gap-1">
                        <span class="material-icons text-sm">verified</span> Verified Donor
                    </span>
                </div>
                
                <div class="flex flex-col sm:flex-row items-center justify-center md:justify-start gap-2 sm:gap-4 text-gray-500 text-sm font-medium">
                    <span class="flex items-center gap-1.5"><span class="material-icons text-gray-400 text-sm">location_on</span> {{ profile.city }}, {{ profile.district }}</span>
                    <span v-if="profile.is_available" class="flex items-center gap-1.5 text-[#FF3D3D]"><span class="w-2 h-2 bg-[#FF3D3D] rounded-full animate-pulse"></span> Available for Emergency</span>
                </div>

                <div class="flex flex-wrap justify-center md:justify-start gap-2 pt-2">
                   <div class="px-3 py-1 bg-gray-50 rounded-lg text-[10px] font-bold text-gray-400 border border-gray-100 uppercase tracking-widest">Whole Blood</div>
                </div>
            </div>

            <!-- Actions -->
            <div class="w-full md:w-auto mt-4 md:mt-0">
                <button v-if="!contactRevealed" @click="showContact" class="w-full md:w-auto px-8 py-3.5 rounded-xl bg-[#FF3D3D] hover:bg-red-600 text-white font-bold text-sm shadow-xl shadow-red-500/30 transition-all flex items-center justify-center gap-2 transform hover:-translate-y-1">
                    <span class="material-icons text-sm">phone</span>
                    Contact Donor
                </button>
                <div v-else class="flex flex-col sm:flex-row gap-2 w-full">
                    <div class="px-6 py-3.5 rounded-xl bg-gray-50 border border-gray-200 text-gray-900 font-bold text-sm flex items-center justify-center gap-2 flex-1 md:flex-none md:min-w-[160px]">
                        <span class="material-icons text-gray-400 text-sm">phone_iphone</span>
                        {{ profile.phone }}
                    </div>
                    <a :href="'tel:' + profile.phone" class="px-6 py-3.5 rounded-xl bg-[#22C55E] hover:bg-green-600 text-white font-bold text-sm shadow-lg shadow-green-500/30 transition-all flex items-center justify-center gap-2 flex-1 md:flex-none">
                        <span class="material-icons text-sm">call</span>
                        Call Now
                    </a>
                </div>
            </div>
         </div>
      </div>

      <!-- 2. Stats Grid -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 sm:gap-6">
        <!-- Total Donations -->
        <div class="bg-white p-6 sm:p-8 rounded-3xl shadow-sm border border-gray-100 flex flex-col items-center text-center hover:shadow-md transition-shadow">
            <h3 class="text-[10px] font-black text-gray-400 uppercase tracking-widest mb-2">Total Donations</h3>
            <span class="text-4xl sm:text-5xl font-black text-gray-900 mb-2">{{ stats.total_donations }}</span>
             <span class="text-[#22C55E] text-[10px] font-black bg-green-50 px-2 py-1 rounded flex items-center gap-1 uppercase tracking-tighter">
                <span class="material-icons text-[10px]">trending_up</span> Top Donor
            </span>
        </div>

        <!-- Lives Saved -->
         <div class="bg-white p-6 sm:p-8 rounded-3xl shadow-sm border border-gray-100 flex flex-col items-center text-center hover:shadow-md transition-shadow">
            <h3 class="text-[10px] font-black text-gray-400 uppercase tracking-widest mb-2">Lives Saved</h3>
            <span class="text-4xl sm:text-5xl font-black text-[#FF3D3D] mb-2">{{ stats.lives_saved }}</span>
            <span class="text-gray-400 text-[10px] font-black uppercase tracking-widest">Est. Impact</span>
        </div>

        <!-- Last Donation -->
         <div class="bg-white p-6 sm:p-8 rounded-3xl shadow-sm border border-gray-100 flex flex-col items-center text-center hover:shadow-md transition-shadow sm:col-span-2 lg:col-span-1">
            <h3 class="text-[10px] font-black text-gray-400 uppercase tracking-widest mb-2">Last Donation</h3>
             <span class="text-3xl sm:text-4xl font-black text-gray-900 mb-1">
                {{ stats.last_donation ? formatDate(stats.last_donation) : 'N/A' }}
            </span>
            <span class="text-gray-400 text-xs font-bold uppercase tracking-tight">
                {{ stats.last_donation ? formatYear(stats.last_donation) : '' }}
            </span>
        </div>
      </div>

      <!-- 3. Main Split -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        
        <!-- Left Sidebar -->
        <div class="space-y-8">
            <!-- About -->
            <div class="bg-white p-6 rounded-3xl shadow-sm border border-gray-100">
                <h3 class="font-bold text-gray-900 mb-4">About {{ profile.name?.split(' ')[0] }}</h3>
                <p class="text-sm text-gray-500 leading-relaxed mb-6 font-medium">
                    A dedicated donor helping save lives in {{ profile.district }}.
                </p>
                
                 <div class="flex items-start gap-3">
                    <span class="material-icons text-gray-300">location_on</span>
                        <div>
                        <p class="text-xs font-bold text-gray-400 uppercase">Location</p>
                        <p class="text-sm font-bold text-gray-900">{{ profile.area_village }}</p>
                    </div>
                </div>
            </div>

            <!-- Achievements -->
             <div class="bg-white p-6 rounded-3xl shadow-sm border border-gray-100">
                <h3 class="font-bold text-gray-900 mb-4">Achievements</h3>
                <div class="flex flex-wrap gap-3">
                    <span class="inline-flex items-center gap-1.5 px-3 py-1.5 bg-yellow-50 text-yellow-700 rounded-lg text-xs font-bold border border-yellow-100">
                        <span class="material-icons text-sm">emoji_events</span> Gold Donor
                    </span>
                    <span class="inline-flex items-center gap-1.5 px-3 py-1.5 bg-blue-50 text-blue-700 rounded-lg text-xs font-bold border border-blue-100">
                        <span class="material-icons text-sm">verified_user</span> ID Verified
                    </span>
                </div>
            </div>
        </div>

        <!-- Right Content -->
        <div class="lg:col-span-2 space-y-8">
            <!-- Donation History -->
            <div class="bg-white p-8 rounded-3xl shadow-sm border border-gray-100">
                <div class="flex justify-between items-center mb-8">
                    <h3 class="font-bold text-gray-900 text-lg">Donation History</h3>
                </div>

                <!-- Timeline -->
                <div class="relative pl-4 border-l-2 border-gray-100 space-y-12">
                     <div v-if="history.length === 0" class="text-gray-500 italic text-sm">No donation history available.</div>

                     <!-- Item -->
                    <div v-for="donation in history" :key="donation.id" class="relative">
                        <span class="absolute -left-[21px] top-1 w-3 h-3 rounded-full border-2 border-white ring-2" :class="donation.verified ? 'bg-[#FF3D3D] ring-red-100' : 'bg-gray-300 ring-gray-100'"></span>
                        <div class="bg-gray-50 rounded-2xl p-6 hover:bg-gray-50 transition-colors group">
                            <div class="flex justify-between items-start mb-4">
                                <div>
                                    <h4 class="font-bold text-gray-900">{{ donation.type }}</h4>
                                    <span v-if="donation.verified" class="bg-green-100 text-[#22C55E] text-[10px] font-bold px-2 py-0.5 rounded ml-2">Verified</span> 
                                </div>
                                <span class="bg-white px-3 py-1 rounded-lg text-xs font-bold text-gray-500 shadow-sm">{{ formatDateFull(donation.date) }}</span>
                            </div>
                            
                            <div class="flex flex-wrap gap-4 text-xs font-medium text-gray-500 mb-4">
                                <span class="flex items-center gap-1"><span class="material-icons text-gray-400 text-sm">location_on</span> {{ donation.location }}</span>
                                <span v-if="donation.amount_ml" class="flex items-center gap-1"><span class="material-icons text-gray-400 text-sm">opacity</span> {{ donation.amount_ml }}ml</span>
                            </div>
                            <p v-if="donation.notes" class="text-xs text-gray-500 mb-4">{{ donation.notes }}</p>
                            <p v-else class="text-xs text-gray-400 italic mb-4">"Thank you for your generous donation."</p>

                            <!-- Donation Proof Image -->
                            <div v-if="donation.image" class="mt-2 rounded-xl overflow-hidden border border-gray-100 bg-white max-w-sm">
                                <img :src="getDonationImageUrl(donation.image)" alt="Donation Proof" class="w-full h-auto object-cover max-h-64 hover:scale-105 transition-transform duration-500" />
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import api from '@/lib/axios';
import { useToastStore } from '@/stores/toast';
import UserAvatar from '@/components/UserAvatar.vue';

const toastStore = useToastStore();

const route = useRoute();
const router = useRouter();
const loading = ref(true);
const contactRevealed = ref(false);

const profile = ref<any>(null);
const stats = ref<any>(null);
const history = ref<any[]>([]);

onMounted(async () => {
    const donorId = route.params.id;
    if (!donorId) {
        router.push('/');
        return;
    }

    try {
        const res = await api.get(`/donors/${donorId}`);
        if (res.data) {
             profile.value = res.data.profile;
             stats.value = res.data.stats;
             history.value = res.data.history;
        }
    } catch (e) {
        console.error("Could not fetch donor profile", e);
        router.push('/');
    } finally {
        loading.value = false;
    }
});

const showContact = async () => {
    if (contactRevealed.value) return;
    
    try {
        const res = await api.get(`/donors/${route.params.id}/contact`);
        // The API returns { phone: "..." }
        if (profile.value) {
            profile.value.phone = res.data.phone;
        }
        contactRevealed.value = true;
    } catch (error) {
        toastStore.show('Please login to view contact details.', 'info');
        router.push('/login');
    }
};

const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric' });
};

const formatYear = (dateString: string) => {
    const date = new Date(dateString);
    return date.getFullYear();
};

const formatDateFull = (dateString: string) => {
    const date = new Date(dateString);
    return date.toLocaleDateString('en-US', { month: 'long', day: 'numeric', year: 'numeric' });
};

const getDonationImageUrl = (path: string) => {
    if (!path) return '';
    const baseUrl = import.meta.env.VITE_API_URL 
        ? import.meta.env.VITE_API_URL.replace('/api/v1', '') 
        : 'http://localhost:4000';
    const cleanPath = path.startsWith('/') ? path : `/${path}`;
    return `${baseUrl}${cleanPath}`;
};
</script>
