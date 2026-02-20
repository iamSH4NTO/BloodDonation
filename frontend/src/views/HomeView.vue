<template>
  <div class="bg-[#FDF8F8] font-sans text-gray-900">
    
    <!-- Hero Section -->
    <section class="relative pt-6 pb-12 lg:pt-12 lg:pb-24 overflow-hidden">
      <!-- Background subtle gradient -->
      <div class="absolute top-0 left-0 w-full h-full bg-linear-to-br from-red-50/80 via-white to-white pointer-events-none"></div>

      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 relative z-10">
        <div class="grid lg:grid-cols-2 gap-12 lg:gap-8 items-center">
          
          <!-- Left Content -->
          <div class="space-y-8 lg:max-w-xl text-center lg:text-left">
            <!-- Urgent Badge -->
            <div class="inline-flex items-center gap-2 px-4 py-2 bg-[#FFEAEA] text-[#FF3D3D] rounded-full text-xs font-bold tracking-wide shadow-sm hover:shadow-md transition-shadow cursor-default mx-auto lg:mx-0 border border-red-100/50">
              <span class="w-2 h-2 bg-[#FF3D3D] rounded-full animate-pulse shadow-[0_0_8px_rgba(255,61,61,0.6)]"></span>
               জরুরি রক্তের প্রয়োজন
            </div>
            
            <h1 class="text-4xl sm:text-5xl lg:text-6xl font-extrabold text-[#1A1A1A] leading-[1.15] tracking-tight">
              এক ফোঁটা রক্ত, <br />
              <span class="text-[#FF3D3D] drop-shadow-sm">একটু আশা</span>
            </h1>
            
            <p class="text-base sm:text-lg text-gray-600 max-w-lg mx-auto lg:mx-0 leading-relaxed font-medium">
              সহায়তার জন্য প্রস্তুত হাজারো মানবিক নায়কের সাথে সংযুক্ত হোন। আমাদের প্ল্যাটফর্ম রক্তদাতা ও রক্তগ্রহীতার মাঝে গড়ে তোলে একটি নির্ভরযোগ্য সেতুবন্ধন। নিরাপদ এবং সম্পূর্ণ বিনামূল্যে।
            </p>
            
            <div class="flex flex-col sm:flex-row items-center justify-center lg:justify-start gap-4 pt-2">
              <router-link to="/search" class="w-full sm:w-auto bg-[#FF3D3D] hover:bg-red-600 text-white px-8 py-3.5 rounded-xl font-bold shadow-[0_8px_20px_rgba(255,61,61,0.25)] hover:shadow-[0_12px_25px_rgba(255,61,61,0.35)] transition-all transform hover:-translate-y-0.5 flex items-center justify-center gap-2 text-sm">
                <span class="material-icons text-lg">search</span>
                রক্তদাতা খুঁজুন
              </router-link>
              
              <button class="w-full sm:w-auto flex items-center justify-center gap-2 px-8 py-3.5 rounded-xl bg-white text-gray-800 font-bold shadow-sm hover:shadow-md border border-gray-100 hover:border-gray-200 transition-all text-sm group">
                <span class="material-icons text-[#FF3D3D] group-hover:scale-110 transition-transform">play_circle</span>
                কীভাবে কাজ করে
              </button>
            </div>
            
            <div class="flex flex-wrap items-center justify-center lg:justify-start gap-4 sm:gap-6 pt-4">
               <div class="flex items-center gap-2 bg-green-50/80 px-4 py-2 rounded-xl border border-green-100/50 shadow-sm">
                 <span class="material-icons text-green-500 text-sm">check_circle</span>
                 <span class="text-xs text-gray-800 font-bold">ভেরিফাইড রক্তদাতা</span>
               </div>
               <div class="flex items-center gap-2 bg-blue-50/80 px-4 py-2 rounded-xl border border-blue-100/50 shadow-sm">
                 <span class="material-icons text-blue-500 text-sm">shield</span>
                 <span class="text-xs text-gray-800 font-bold">নিরাপদ তথ্য</span>
               </div>
            </div>
          </div>

          <!-- Right Image Area -->
          <div class="relative z-10 lg:h-[500px] flex items-center justify-center lg:justify-end">
             <div class="relative w-full max-w-sm lg:max-w-md">
                <!-- Main Image -->
                <div class="relative rounded-3xl sm:rounded-[2.5rem] overflow-hidden shadow-[0_20px_50px_rgba(0,0,0,0.1)] border-8 border-white z-10 aspect-square sm:aspect-auto transform transition-transform duration-700 hover:scale-[1.02]">
                    <img src="https://images.unsplash.com/photo-1579154204601-01588f351e67?ixlib=rb-4.0.3&auto=format&fit=crop&w=1000&q=80" 
                         alt="Doctor holding heart" 
                         class="w-full h-full object-cover" />
                </div>
                
                <!-- Floating Card -->
                <div class="absolute -bottom-6 -left-4 sm:-bottom-8 sm:-left-10 bg-white/90 backdrop-blur-xl p-4 rounded-2xl shadow-[0_15px_35px_rgba(0,0,0,0.08)] flex items-center gap-4 z-20 border border-white">
                    <div class="bg-red-50 p-3.5 rounded-xl border border-red-100">
                        <span class="material-icons text-[#FF3D3D] text-2xl">favorite</span>
                    </div>
                    <div>
                        <p class="text-xl font-black text-gray-900 drop-shadow-sm">১২ হাজার+</p>
                        <p class="text-xs text-gray-500 font-bold uppercase tracking-wide">জীবন রক্ষিত</p>
                    </div>
                </div>
             </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Floating Search Section -->
    <div class="relative z-30 -mt-8 sm:-mt-14 px-4 sm:px-6 lg:px-8">
        <div class="max-w-6xl mx-auto">
            <div class="bg-white/95 backdrop-blur-xl rounded-3xl shadow-[0_15px_40px_rgba(0,0,0,0.06)] border border-white p-6 sm:p-8 relative overflow-hidden">
                <div class="absolute top-0 left-0 w-full h-1 bg-linear-to-r from-[#FF3D3D] to-red-400"></div>
                <div class="flex items-center gap-3 mb-6">
                    <div class="w-3 h-3 rounded-full bg-[#FF3D3D] shadow-[0_0_10px_rgba(255,61,61,0.5)]"></div>
                    <h2 class="text-xl font-bold text-gray-900">এখনই রক্তদাতা খুঁজুন</h2>
                </div>
                
                <DonorSearchForm @search="handleSearch" />
            </div>
        </div>
    </div>

    <!-- Process Section -->
    <section class="py-20 sm:py-32 bg-white relative">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
             <div class="text-center mb-16">
                <span class="inline-block px-3 py-1 bg-red-50 text-[#FF3D3D] rounded-full font-bold tracking-widest uppercase text-[10px] mb-3 border border-red-100">আমাদের প্রক্রিয়া</span>
                <h2 class="text-3xl sm:text-4xl font-extrabold text-[#1A1A1A]">কীভাবে কাজ করে</h2>
                <p class="text-gray-500 mt-4 max-w-xl mx-auto text-base leading-relaxed">রক্তদান করা বা সাহায্য নেওয়া এখন আরও সহজ। মাত্র তিনটি ধাপে সম্পন্ন করুন।</p>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
                <!-- Step 1 -->
                <div class="relative bg-[#FBFAFA] rounded-4xl p-8 transition-all duration-300 hover:-translate-y-2 hover:shadow-[0_20px_40px_rgba(59,130,246,0.1)] border border-gray-100 hover:border-blue-200 group">
                    <div class="flex justify-between items-start mb-8">
                        <div class="bg-blue-100 p-4 rounded-2xl text-blue-600 transition-transform duration-300 group-hover:scale-110 group-hover:bg-blue-600 group-hover:text-white shadow-sm group-hover:shadow-md">
                             <span class="material-icons text-3xl">person_add</span>
                        </div>
                        <span class="text-7xl font-black text-gray-50 group-hover:text-blue-50 transition-colors -mt-4">০১</span>
                    </div>
                    <h3 class="text-xl font-bold text-gray-900 mb-3">নিবন্ধন করুন</h3>
                    <p class="text-sm text-gray-600 leading-relaxed font-medium">খুব সহজেই নিবন্ধন করে নিজের প্রোফাইল তৈরি করুন। রক্তদাতা বা গ্রহীতা হিসেবে যুক্ত হতে পারেন।</p>
                </div>

                <!-- Step 2 -->
                 <div class="relative bg-[#FBFAFA] rounded-4xl p-8 transition-all duration-300 hover:-translate-y-2 hover:shadow-[0_20px_40px_rgba(168,85,247,0.1)] border border-gray-100 hover:border-purple-200 group">
                    <div class="flex justify-between items-start mb-8">
                        <div class="bg-purple-100 p-4 rounded-2xl text-purple-600 transition-transform duration-300 group-hover:scale-110 group-hover:bg-purple-600 group-hover:text-white shadow-sm group-hover:shadow-md">
                             <span class="material-icons text-3xl">manage_search</span>
                        </div>
                        <span class="text-7xl font-black text-gray-50 group-hover:text-purple-50 transition-colors -mt-4">০২</span>
                    </div>
                    <h3 class="text-xl font-bold text-gray-900 mb-3">খুঁজুন</h3>
                    <p class="text-sm text-gray-600 leading-relaxed font-medium">আমাদের উন্নত ফিল্টার ব্যবহার করে আপনার প্রয়োজনীয় রক্তের গ্রুপ ও স্থান অনুযায়ী রক্তদাতা খুঁজুন।</p>
                </div>

                <!-- Step 3 -->
                 <div class="relative bg-[#FBFAFA] rounded-4xl p-8 transition-all duration-300 hover:-translate-y-2 hover:shadow-[0_20px_40px_rgba(34,197,94,0.1)] border border-gray-100 hover:border-green-200 group">
                    <div class="flex justify-between items-start mb-8">
                        <div class="bg-green-100 p-4 rounded-2xl text-green-600 transition-transform duration-300 group-hover:scale-110 group-hover:bg-green-600 group-hover:text-white shadow-sm group-hover:shadow-md">
                             <span class="material-icons text-3xl">phone_in_talk</span>
                        </div>
                        <span class="text-7xl font-black text-gray-50 group-hover:text-green-50 transition-colors -mt-4">০৩</span>
                    </div>
                    <h3 class="text-xl font-bold text-gray-900 mb-3">যোগাযোগ করুন</h3>
                    <p class="text-sm text-gray-600 leading-relaxed font-medium">সরাসরি রক্তদাতার সাথে যোগাযোগ করে রক্তদানের ব্যবস্থা করুন সম্পূর্ণ নিরাপদে।</p>
                </div>
            </div>
        </div>
    </section>

    <!-- Featured Donors -->
    <section class="py-20 sm:py-32 bg-[#F8F9FA] border-t border-gray-100">
         <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="flex flex-col sm:flex-row justify-between items-center sm:items-end mb-12 gap-6 text-center sm:text-left">
                 <div>
                    <span class="inline-block px-3 py-1 bg-red-50 text-[#FF3D3D] rounded-full font-bold tracking-widest uppercase text-[10px] mb-3 border border-red-100">আমাদের হিরোরা</span>
                    <h2 class="text-3xl sm:text-4xl font-extrabold text-[#1A1A1A]">সাম্প্রতিক রক্তদাতাগণ</h2>
                 </div>
                 <router-link to="/search" class="inline-flex items-center px-6 py-3 bg-white border border-gray-200 rounded-xl text-gray-700 font-bold hover:bg-gray-50 hover:border-gray-300 transition-all shadow-sm group">
                    সব রক্তদাতা দেখুন <span class="material-icons text-sm ml-2 text-[#FF3D3D] group-hover:translate-x-1 transition-transform">arrow_forward</span>
                 </router-link>
            </div>

            <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
                <!-- Donor cards logic repeated for variety -->
                <div v-for="(donor, index) in donors" :key="index" class="bg-white rounded-3xl overflow-hidden shadow-[0_4px_20px_rgba(0,0,0,0.03)] hover:shadow-[0_15px_35px_rgba(0,0,0,0.08)] transition-all duration-300 group border border-gray-100 hover:border-gray-200 flex flex-col h-full">
                    <div :class="['h-28 relative', donor.bgClass]">
                         <span class="absolute top-4 right-4 bg-white/90 backdrop-blur-sm text-[#FF3D3D] text-[11px] font-black px-3 py-1.5 rounded-lg flex items-center justify-center shadow-sm border border-white/50">
                            {{ donor.bloodGroup }}
                         </span>
                    </div>
                    <div class="px-6 pb-6 relative flex flex-col grow">
                        <div class="relative w-24 h-24 -mt-12 mb-4 mx-auto sm:mx-0">
                             <UserAvatar 
                                :src="donor.image" 
                                :gender="donor.gender" 
                                :name="donor.name" 
                                size="xl" 
                                class="shadow-[0_8px_20px_rgba(0,0,0,0.1)] border-[6px] border-white ring-1 ring-gray-100"
                             />
                        </div>
                        <h3 class="text-xl font-bold text-gray-900 text-center sm:text-left">{{ donor.name }}</h3>
                        <div class="flex items-center justify-center sm:justify-start text-gray-500 text-sm mb-6 mt-1.5 font-medium">
                            <span class="material-icons text-[16px] mr-1 text-[#FF3D3D]">location_on</span>
                            {{ donor.location }}
                        </div>
                        <button class="w-full mt-auto py-3 bg-gray-50 border border-gray-200 rounded-xl text-sm font-bold text-gray-700 hover:border-[#FF3D3D] hover:bg-[#FF3D3D] hover:text-white transition-all duration-300 flex items-center justify-center gap-2 group/btn">
                             <span class="material-icons text-[18px] text-[#FF3D3D] group-hover/btn:text-white transition-colors">call</span> যোগাযোগ
                        </button>
                    </div>
                </div>
            </div>
         </div>
    </section>

    <!-- CTA Section -->
    <section class="py-16 sm:py-24 px-4 sm:px-6 lg:px-8 bg-white">
        <div class="max-w-7xl mx-auto bg-gray-900 rounded-[2.5rem] sm:rounded-4xl overflow-hidden shadow-2xl relative isolate min-h-[500px] sm:min-h-[450px]">
            <!-- Background Image with Overlay -->
            <img src="https://images.unsplash.com/photo-1536856136534-bb679c52a9aa?ixlib=rb-4.0.3&auto=format&fit=crop&w=2000&q=80" alt="Hands" class="absolute inset-0 w-full h-full object-cover opacity-20 mix-blend-luminosity grayscale">
            <div class="absolute inset-0 bg-linear-to-b sm:bg-linear-to-r from-gray-900/95 via-gray-900/80 sm:via-gray-900/60 to-transparent"></div>

            <div class="relative px-6 py-12 sm:px-12 lg:px-20 h-full flex flex-col justify-center text-center sm:text-left">
                <div class="lg:max-w-xl">
                     <h2 class="text-4xl sm:text-5xl lg:text-6xl font-black text-white tracking-tight mb-6 leading-tight drop-shadow-lg">
                        জীবন বাঁচানোর এই <br/><span class="text-[#FF3D3D]">উদ্যোগে যোগ দিন</span>
                     </h2>
                     <p class="text-gray-300 text-base sm:text-lg mb-10 max-w-lg mx-auto sm:mx-0 font-medium leading-relaxed">
                        আপনার এক ফোঁটা রক্ত একটি জীবন বাঁচাতে পারে। আজই আমাদের সাথে যুক্ত হোন এবং মানবিক এই কার্যক্রমে অবদান রাখুন।
                     </p>
                     <div class="flex flex-col sm:flex-row items-center justify-center sm:justify-start gap-4">
                        <router-link to="/register" class="w-full sm:w-auto bg-[#FF3D3D] hover:bg-red-600 text-white px-10 py-4 rounded-xl font-bold shadow-[0_10px_20px_rgba(255,61,61,0.3)] hover:shadow-[0_15px_30px_rgba(255,61,61,0.4)] transition-all transform hover:-translate-y-1 text-center">
                            রক্তদাতা হিসেবে যুক্ত হোন
                        </router-link>
                         <button class="w-full sm:w-auto px-10 py-4 rounded-xl font-bold text-white border-2 border-white/20 hover:bg-white/10 hover:border-white/40 transition-all backdrop-blur-sm">
                            আমাদের সম্পর্কে
                        </button>
                     </div>
                </div>
            </div>
            
            <!-- Floating Stats Card (Bottom Right) -->
            <div class="absolute bottom-12 right-12 hidden lg:block">
                 <div class="bg-white/10 backdrop-blur-xl p-8 rounded-4xl border border-white/20 text-center w-64 shadow-2xl transform transition-transform duration-500 hover:scale-105 hover:-translate-y-2 cursor-default relative overflow-hidden">
                     <div class="absolute -top-10 -right-10 w-32 h-32 bg-[#FF3D3D] rounded-full blur-3xl opacity-20"></div>
                     <div class="bg-white w-14 h-14 rounded-2xl flex items-center justify-center mx-auto mb-5 shadow-lg transform -rotate-6">
                         <span class="material-icons text-[#FF3D3D] text-2xl animate-pulse">water_drop</span>
                     </div>
                     <p class="text-5xl font-extrabold text-white mb-2 drop-shadow-md">৪৫০+</p>
                     <p class="text-red-200 text-xs font-bold uppercase tracking-[0.15em]">নিবন্ধিত রক্তদাতা</p>
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
import UserAvatar from '@/components/UserAvatar.vue';

const router = useRouter();

const donors = ref([
    { name: 'Rahim Uddin', location: 'Mirpur, Dhaka', bloodGroup: 'A+', bgClass: 'bg-red-50/80', image: 'https://i.pravatar.cc/150?img=11', gender: 'Male' },
    { name: 'Fatima Begum', location: 'Sadar, Sylhet', bloodGroup: 'O-', bgClass: 'bg-blue-50/80', image: 'https://i.pravatar.cc/150?img=5', gender: 'Female' },
    { name: 'Kamal Hossain', location: 'Raozan, Chittagong', bloodGroup: 'B+', bgClass: 'bg-purple-50/80', image: 'https://i.pravatar.cc/150?img=3', gender: 'Male' },
    { name: 'Nusrat Jahan', location: 'Uttara, Dhaka', bloodGroup: 'AB-', bgClass: 'bg-green-50/80', image: 'https://i.pravatar.cc/150?img=9', gender: 'Female' }
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
