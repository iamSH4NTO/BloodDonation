<template>
  <div class="bg-[#FDF8F8] font-sans">
    
    <!-- Hero Section -->
    <section class="relative pt-4 pb-8 lg:pt-6 lg:pb-16 overflow-visible">
      <!-- Background subtle gradient -->
      <div class="absolute top-0 left-0 w-full h-full bg-linear-to-br from-red-50/50 via-white to-white pointer-events-none"></div>

      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 relative z-10">
        <div class="grid lg:grid-cols-2 gap-8 lg:gap-6 items-center">
          
          <!-- Left Content -->
          <div class="space-y-6 lg:max-w-xl">
            <!-- Urgent Badge -->
            <div class="inline-flex items-center gap-2 px-4 py-1.5 bg-[#FFEAEA] text-[#FF3D3D] rounded-full text-xs font-bold tracking-wide shadow-sm transform hover:scale-105 transition-transform cursor-default">
              <span class="w-2 h-2 bg-[#FF3D3D] rounded-full animate-pulse"></span>
              Urgent need for O- donors in Dhaka
            </div>
            
            <h1 class="text-3xl lg:text-5xl font-extrabold text-[#1A1A1A] leading-[1.1] tracking-tight">
              Donate Blood, <br />
              <span class="text-[#FF3D3D]">Save a Life</span> Today.
            </h1>
            
            <p class="text-base text-gray-500 max-w-lg leading-relaxed font-medium">
              Connect with thousands of heroes ready to help. Our platform bridges the gap between donors and those in critical need. Fast, secure, and always free.
            </p>
            
            <div class="flex flex-wrap items-center gap-4 pt-2">
              <router-link to="/search" class="bg-[#FF3D3D] hover:bg-red-600 text-white px-6 py-3 rounded-xl font-bold shadow-[0_10px_30px_rgba(255,61,61,0.3)] hover:shadow-[0_15px_35px_rgba(255,61,61,0.4)] transition-all transform hover:-translate-y-1 flex items-center gap-2 text-sm">
                <span class="material-icons text-lg">search</span>
                Find Donors
              </router-link>
              
              <button class="flex items-center gap-2 px-5 py-3 rounded-xl bg-white text-gray-700 font-bold shadow-md border border-gray-100 hover:bg-gray-50 transition-all text-sm group">
                <span class="material-icons text-[#FF3D3D] group-hover:scale-110 transition-transform">play_circle</span>
                How it works
              </button>
            </div>
            
            <div class="flex items-center gap-6 pt-4">
               <div class="flex items-center gap-2 bg-green-50 px-3 py-1.5 rounded-lg border border-green-100">
                 <span class="material-icons text-green-500 text-sm">check_circle</span>
                 <span class="text-xs text-gray-700 font-bold">Verified Donors</span>
               </div>
               <div class="flex items-center gap-2 bg-blue-50 px-3 py-1.5 rounded-lg border border-blue-100">
                 <span class="material-icons text-blue-500 text-sm">shield</span>
                 <span class="text-xs text-gray-700 font-bold">Secure Data</span>
               </div>
            </div>
          </div>

          <!-- Right Image Area -->
          <div class="relative z-10 lg:h-[450px] flex items-center justify-end">
             <div class="relative w-full max-w-md lg:max-w-[420px]">
                <!-- Main Image -->
                <div class="relative rounded-4xl overflow-hidden shadow-2xl border-[6px] border-white z-10">
                    <img src="https://images.unsplash.com/photo-1579154204601-01588f351e67?ixlib=rb-4.0.3&auto=format&fit=crop&w=1000&q=80" 
                         alt="Doctor holding heart" 
                         class="w-full h-auto object-cover transform scale-100 hover:scale-105 transition-transform duration-700" />
                </div>
                
                <!-- Floating Card -->
                <div class="absolute -bottom-6 -left-6 md:bottom-6 md:-left-8 bg-white p-3 rounded-xl shadow-[0_20px_40px_rgba(0,0,0,0.1)] flex items-center gap-3 z-20 border border-gray-50 animate-bounce-slow">
                    <div class="bg-[#FFF0F0] p-3 rounded-lg">
                        <span class="material-icons text-[#FF3D3D] text-xl">favorite</span>
                    </div>
                    <div>
                        <p class="text-lg font-extrabold text-gray-900">12k+</p>
                        <p class="text-[10px] text-gray-400 font-bold uppercase tracking-wider">Lives Saved</p>
                    </div>
                </div>
             </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Floating Search Section -->
    <div class="relative z-20 -mt-12 px-4 sm:px-6 lg:px-8">
        <div class="max-w-6xl mx-auto">
            <div class="bg-white rounded-3xl shadow-[0_20px_60px_rgba(0,0,0,0.08)] p-5 lg:p-6">
                <div class="flex items-center gap-3 mb-4">
                    <div class="w-2.5 h-2.5 rounded-full bg-[#FF3D3D] ring-4 ring-red-100"></div>
                    <h2 class="text-lg font-bold text-gray-900">Find a Donor Now</h2>
                </div>
                
                <DonorSearchForm @search="handleSearch" />
            </div>
        </div>
    </div>

    <!-- Process Section -->
    <section class="py-8 bg-white relative">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
             <div class="text-center mb-8">
                <span class="text-[#FF3D3D] font-bold tracking-[0.2em] uppercase text-[10px]">Process</span>
                <h2 class="text-3xl font-extrabold text-[#1A1A1A] mt-2">How It Works</h2>
                <p class="text-gray-500 mt-2 max-w-xl mx-auto text-base">Getting help or saving a life is simpler than you think. Just three easy steps.</p>
            </div>

            <div class="grid md:grid-cols-3 gap-5">
                <!-- Step 1 -->
                <div class="relative bg-[#FBFAFA] rounded-3xl p-5 lg:p-6 transition-all hover:-translate-y-2 hover:shadow-xl border border-transparent hover:border-blue-100 group">
                    <div class="flex justify-between items-start mb-4">
                        <div class="bg-blue-100 p-3 rounded-xl text-blue-600 group-hover:scale-110 transition-transform duration-300">
                             <span class="material-icons text-2xl">person_add</span>
                        </div>
                        <span class="text-5xl font-black text-gray-100 group-hover:text-blue-50 transition-colors -mt-2 opacity-50">01</span>
                    </div>
                    <h3 class="text-lg font-bold text-gray-900 mb-2">Register</h3>
                    <p class="text-sm text-gray-500 leading-relaxed font-medium">Sign up in minutes. Create your profile whether you want to donate blood or find a donor.</p>
                </div>

                <!-- Step 2 -->
                 <div class="relative bg-[#FBFAFA] rounded-3xl p-5 lg:p-6 transition-all hover:-translate-y-2 hover:shadow-xl border border-transparent hover:border-purple-100 group">
                    <div class="flex justify-between items-start mb-4">
                        <div class="bg-purple-100 p-3 rounded-xl text-purple-600 group-hover:scale-110 transition-transform duration-300">
                             <span class="material-icons text-2xl">filter_list</span>
                        </div>
                        <span class="text-5xl font-black text-gray-100 group-hover:text-purple-50 transition-colors -mt-2 opacity-50">02</span>
                    </div>
                    <h3 class="text-lg font-bold text-gray-900 mb-2">Search</h3>
                    <p class="text-sm text-gray-500 leading-relaxed font-medium">Use our advanced filter to find donors by specific blood group, division, and district instantly.</p>
                </div>

                <!-- Step 3 -->
                 <div class="relative bg-[#FBFAFA] rounded-3xl p-5 lg:p-6 transition-all hover:-translate-y-2 hover:shadow-xl border border-transparent hover:border-green-100 group">
                    <div class="flex justify-between items-start mb-4">
                        <div class="bg-green-100 p-3 rounded-xl text-green-600 group-hover:scale-110 transition-transform duration-300">
                             <span class="material-icons text-2xl">connect_without_contact</span>
                        </div>
                        <span class="text-5xl font-black text-gray-100 group-hover:text-green-50 transition-colors -mt-2 opacity-50">03</span>
                    </div>
                    <h3 class="text-lg font-bold text-gray-900 mb-2">Connect</h3>
                    <p class="text-sm text-gray-500 leading-relaxed font-medium">Get direct contact details of the donor and connect with them to arrange the donation.</p>
                </div>
            </div>
        </div>
    </section>

    <!-- Featured Donors -->
    <section class="py-10 bg-[#F8F9FA]">
         <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="flex flex-col sm:flex-row justify-between items-end mb-8 gap-4">
                 <div class="text-center sm:text-left">
                    <span class="text-[#FF3D3D] font-bold tracking-[0.2em] uppercase text-xs">Heroes</span>
                    <h2 class="text-4xl font-extrabold text-[#1A1A1A] mt-2">Featured Donors</h2>
                 </div>
                 <router-link to="/search" class="inline-flex items-center text-[#FF3D3D] font-bold hover:text-red-700 transition-colors group">
                    View All Donors <span class="material-icons text-sm ml-2 group-hover:translate-x-1 transition-transform">arrow_forward</span>
                 </router-link>
            </div>

            <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
                <!-- Donor cards logic repeated for variety -->
                <div v-for="(donor, index) in donors" :key="index" class="bg-white rounded-2xl overflow-hidden shadow-sm hover:shadow-[0_20px_40px_rgba(0,0,0,0.08)] transition-all duration-300 group border border-gray-100">
                    <div :class="['h-28 relative', donor.bgClass]">
                         <span class="absolute top-4 right-4 bg-[#FF3D3D] text-white text-[10px] font-black px-2.5 py-1 rounded-md shadow-sm tracking-wide">{{ donor.bloodGroup }}</span>
                    </div>
                    <div class="px-6 pb-6 relative">
                        <div class="w-20 h-20 rounded-full border-[5px] border-white shadow-lg bg-gray-200 -mt-10 mb-3 overflow-hidden mx-auto sm:mx-0">
                             <img :src="donor.image" alt="User" class="w-full h-full object-cover">
                        </div>
                        <h3 class="text-lg font-bold text-gray-900 text-center sm:text-left">{{ donor.name }}</h3>
                        <div class="flex items-center justify-center sm:justify-start text-gray-500 text-xs mb-6 mt-1 font-medium">
                            <span class="material-icons text-sm mr-1 text-[#FF3D3D]">location_on</span>
                            {{ donor.location }}
                        </div>
                        <button class="w-full py-2.5 border border-gray-200 rounded-xl text-sm font-bold text-gray-600 hover:border-[#FF3D3D] hover:bg-[#FF3D3D] hover:text-white transition-all duration-300 flex items-center justify-center gap-2 group-hover:shadow-md">
                             <span class="material-icons text-sm">call</span> Contact
                        </button>
                    </div>
                </div>
            </div>
         </div>
    </section>

    <!-- CTA Section -->
    <section class="py-8 px-4 sm:px-6 lg:px-8 bg-white">
        <div class="max-w-7xl mx-auto bg-[#111111] rounded-[2.5rem] overflow-hidden shadow-2xl relative isolate h-[400px] lg:h-[450px]">
            <!-- Background Image with Overlay -->
            <img src="https://images.unsplash.com/photo-1536856136534-bb679c52a9aa?ixlib=rb-4.0.3&auto=format&fit=crop&w=2000&q=80" alt="Hands" class="absolute inset-0 w-full h-full object-cover opacity-40 mix-blend-overlay">
            <div class="absolute inset-0 bg-linear-to-r from-black/90 via-black/60 to-transparent"></div>

            <div class="relative px-8 py-16 lg:px-20 h-full flex flex-col justify-center">
                <div class="lg:max-w-xl">
                     <h2 class="text-4xl lg:text-6xl font-black text-white tracking-tight mb-6 leading-tight">
                        Become a <br/><span class="text-[#FF3D3D]">Life Saver</span>
                     </h2>
                     <p class="text-gray-400 text-lg mb-10 max-w-lg font-medium leading-relaxed">
                        There are many requests for blood every day. Your donation can save a life. Join our community of heroes today.
                     </p>
                     <div class="flex flex-wrap gap-4">
                        <router-link to="/register" class="bg-[#FF3D3D] hover:bg-red-600 text-white px-8 py-4 rounded-xl font-bold shadow-[0_10px_20px_rgba(220,38,38,0.4)] transition-all">
                            Join as a Donor
                        </router-link>
                         <button class="px-8 py-4 rounded-xl font-bold text-white border-2 border-white/20 hover:bg-white/10 transition-colors backdrop-blur-sm">
                            Learn More
                        </button>
                     </div>
                </div>
            </div>
            
            <!-- Floating Stats Card (Bottom Right) -->
            <div class="absolute bottom-8 right-8 lg:bottom-12 lg:right-16 hidden sm:block">
                 <div class="bg-white/10 backdrop-blur-md p-6 rounded-2xl border border-white/10 text-center w-64 shadow-2xl hover:scale-105 transition-transform cursor-default">
                     <div class="bg-white w-14 h-14 rounded-full flex items-center justify-center mx-auto mb-4 shadow-lg">
                         <span class="material-icons text-[#FF3D3D] text-2xl">water_drop</span>
                     </div>
                     <p class="text-4xl font-extrabold text-white mb-1">450+</p>
                     <p class="text-red-200 text-[10px] font-bold uppercase tracking-[0.2em]">Donors Registered</p>
                 </div>
            </div>
        </div>
    </section>

  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { ref } from 'vue';
import DonorSearchForm from '@/components/DonorSearchForm.vue';

const router = useRouter();

const donors = ref([
    { name: 'Rahim Uddin', location: 'Dhaka, Mirpur', bloodGroup: 'A+', bgClass: 'bg-red-50', image: 'https://i.pravatar.cc/150?img=11' },
    { name: 'Fatima Begum', location: 'Sylhet, Sadar', bloodGroup: 'O-', bgClass: 'bg-blue-50', image: 'https://i.pravatar.cc/150?img=5' },
    { name: 'Kamal Hossain', location: 'Chittagong, Raozan', bloodGroup: 'B+', bgClass: 'bg-purple-50', image: 'https://i.pravatar.cc/150?img=3' },
    { name: 'Nusrat Jahan', location: 'Dhaka, Uttara', bloodGroup: 'AB-', bgClass: 'bg-green-50', image: 'https://i.pravatar.cc/150?img=9' }
]);

const handleSearch = (filters: any) => {
    router.push({
        path: '/search',
        query: {
            group: filters.group,
            district: filters.district,
            upazila: filters.upazila,
            available: filters.availableOnly,
            q: filters.locationQuery
        }
    });
};
</script>

<style>
/* Custom curve for hero image if needed, though border-radius works well */
.animate-bounce-slow {
    animation: bounce 3s infinite;
}
</style>
