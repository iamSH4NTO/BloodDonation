<template>
  <div class="min-h-screen bg-[#FAFAFA] font-sans pt-4 pb-4">
    
    <!-- Top Decoration/Breadcrumb area (Optional matches design white space) -->

    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8">
      
      <!-- 1. Profile Header Card -->
      <div class="bg-white rounded-2xl p-5 shadow-sm border border-gray-100 relative overflow-visible">
         <!-- Background decoration if needed, currently clean white -->
         
         <div class="flex flex-col md:flex-row gap-6 md:gap-8 items-start md:items-center">
            <!-- Avatar -->
            <div class="relative group">
                <UserAvatar 
                  :src="profile.profile_picture" 
                  :gender="profile.gender" 
                  :name="profile.name" 
                  size="xl" 
                  class="shadow-xl"
                />
                
                <!-- Upload Overlay on Hover -->
                <label v-if="!isUploading" class="absolute inset-0 flex items-center justify-center bg-black/40 rounded-full opacity-0 group-hover:opacity-100 transition-opacity cursor-pointer">
                    <span class="material-icons text-white">camera_alt</span>
                    <input type="file" class="hidden" accept="image/*" @change="handleImageUpload" />
                </label>
                <div v-else class="absolute inset-0 flex items-center justify-center bg-black/40 rounded-full opacity-100">
                    <span class="material-icons text-white animate-spin">refresh</span>
                </div>

                <div class="absolute -bottom-3 -right-3 bg-[#FF3D3D] text-white w-12 h-12 rounded-full flex items-center justify-center font-bold text-lg shadow-md border-4 border-white">
                    {{ profile.blood_group || '?' }}
                </div>
            </div>

            <!-- Info -->
            <div class="flex-1 space-y-2">
                <div class="flex items-center gap-3 flex-wrap">
                    <h1 class="text-2xl font-extrabold text-gray-900">{{ profile.name || 'Donor Name' }}</h1>
                    <span class="bg-green-100 text-green-700 px-3 py-1 rounded-full text-xs font-bold uppercase tracking-wide flex items-center gap-1">
                        <span class="material-icons text-sm">verified</span> Verified Donor
                    </span>
                    <button v-if="profile.profile_picture" @click="handleImageDelete" class="text-gray-400 hover:text-red-500 transition-colors" title="Remove profile picture">
                        <span class="material-icons text-sm">delete_outline</span>
                    </button>
                </div>
                
                <div class="flex items-center gap-4 text-gray-500 text-sm font-medium">
                    <span class="flex items-center gap-1"><span class="material-icons text-gray-400 text-sm">location_on</span> {{ profile.city || 'Location' }}, {{ profile.district }}</span>
                    <span v-if="profile.is_available" class="flex items-center gap-1 text-[#FF3D3D]"><span class="w-2 h-2 bg-[#FF3D3D] rounded-full animate-pulse"></span> Available for Emergency</span>
                    <span v-else class="flex items-center gap-1 text-gray-400"><span class="w-2 h-2 bg-gray-400 rounded-full"></span> Unavailable</span>
                </div>

                <div class="flex flex-wrap gap-2 pt-2">
                   <div class="px-3 py-1.5 bg-gray-50 rounded-lg text-xs font-semibold text-gray-500 border border-gray-100">Whole Blood</div>
                   <div class="px-3 py-1.5 bg-gray-50 rounded-lg text-xs font-semibold text-gray-500 border border-gray-100">Platelets</div>
                </div>

                <!-- Stats in Header -->
                <div v-if="!isEditing" class="grid grid-cols-3 gap-3 pt-3 border-t border-gray-50 mt-3 md:max-w-md">
                    <div class="text-center md:text-left">
                        <div class="text-[10px] font-bold text-gray-400 uppercase tracking-wider">Donations</div>
                        <div class="text-lg font-black text-gray-900 leading-tight">{{ stats.total_donations }}</div>
                    </div>
                    <div class="text-center md:text-left border-l border-gray-100 pl-3">
                        <div class="text-[10px] font-bold text-gray-400 uppercase tracking-wider">Lives Saved</div>
                        <div class="text-lg font-black text-[#FF3D3D] leading-tight">{{ stats.lives_saved }}</div>
                    </div>
                    <div class="text-center md:text-left border-l border-gray-100 pl-3">
                        <div class="text-[10px] font-bold text-gray-400 uppercase tracking-wider">Last Donated</div>
                        <div class="text-lg font-black text-gray-900 leading-tight">{{ stats.last_donation ? formatDate(stats.last_donation) : 'N/A' }}</div>
                    </div>
                </div>
            </div>

            <!-- Actions -->
            <div class="flex items-center gap-3 self-start md:self-auto mt-4 md:mt-0">
                <button class="px-5 py-2.5 rounded-xl border border-gray-200 text-gray-700 font-bold text-sm hover:bg-gray-50 transition-colors shadow-sm">
                    Share Profile
                </button>
                <button @click="isEditing = !isEditing" class="px-5 py-2.5 rounded-xl bg-[#FF3D3D] hover:bg-red-600 text-white font-bold text-sm shadow-lg shadow-red-500/30 transition-all flex items-center gap-2">
                    <span class="material-icons text-sm">{{ isEditing ? 'close' : 'edit' }}</span>
                    {{ isEditing ? 'Cancel Edit' : 'Edit Profile' }}
                </button>
            </div>
         </div>
      </div>

      <!-- Edit Mode Form -->
      <div v-if="isEditing" class="bg-white rounded-2xl p-6 shadow-lg border border-red-100 animate-fade-in-up">
        <h2 class="text-xl font-bold text-gray-900 mb-6 flex items-center gap-2">
            <span class="material-icons text-[#FF3D3D]">edit_note</span> Update Your Information
        </h2>
        <form @submit.prevent="updateProfile" class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="space-y-1.5">
                <label class="text-xs font-bold text-gray-500 uppercase">Full Name</label>
                <input v-model="profile.name" type="text" class="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-3 focus:ring-2 focus:ring-[#FF3D3D]/20 outline-none font-medium" />
            </div>
            <div class="space-y-1.5">
                <label class="text-xs font-bold text-gray-500 uppercase">Blood Group</label>
                 <select v-model="profile.blood_group" class="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-3 focus:ring-2 focus:ring-[#FF3D3D]/20 outline-none font-medium">
                  <option value="A+">A+</option>
                  <option value="A-">A-</option>
                  <option value="B+">B+</option>
                  <option value="B-">B-</option>
                  <option value="AB+">AB+</option>
                  <option value="AB-">AB-</option>
                  <option value="O+">O+</option>
                  <option value="O-">O-</option>
                </select>
            </div>
            <div class="space-y-1.5">
                <label class="text-xs font-bold text-gray-500 uppercase">Phone</label>
                 <!-- Added Phone field as it's critical -->
                <input v-model="profile.phone" type="text" class="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-3 focus:ring-2 focus:ring-[#FF3D3D]/20 outline-none font-medium" placeholder="017..." />
            </div>
            <div class="space-y-1.5">
                <label class="text-xs font-bold text-gray-500 uppercase">Gender</label>
                 <select v-model="profile.gender" class="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-3 focus:ring-2 focus:ring-[#FF3D3D]/20 outline-none font-medium">
                  <option value="">Select Gender</option>
                  <option value="Male">Male</option>
                  <option value="Female">Female</option>
                  <option value="Other">Other</option>
                </select>
            </div>
            <div class="space-y-1.5">
                <label class="text-xs font-bold text-gray-500 uppercase">Birthday</label>
                <input v-model="profile.birthday" type="date" class="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-3 focus:ring-2 focus:ring-[#FF3D3D]/20 outline-none font-medium" />
            </div>
            <div class="space-y-1.5">
                <label class="text-xs font-bold text-gray-500 uppercase">District</label>
                <input v-model="profile.district" type="text" class="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-3 focus:ring-2 focus:ring-[#FF3D3D]/20 outline-none font-medium" />
            </div>
            <div class="space-y-1.5">
                <label class="text-xs font-bold text-gray-500 uppercase">City/Upazila</label>
                <input v-model="profile.city" type="text" class="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-3 focus:ring-2 focus:ring-[#FF3D3D]/20 outline-none font-medium" />
            </div>
            <div class="space-y-1.5">
                <label class="text-xs font-bold text-gray-500 uppercase">Area/Village</label>
                <input v-model="profile.area_village" type="text" class="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-3 focus:ring-2 focus:ring-[#FF3D3D]/20 outline-none font-medium" />
            </div>
             <div class="space-y-1.5 md:col-span-2">
                 <label class="text-xs font-bold text-gray-500 uppercase">Availability</label>
                 <div class="flex items-center gap-3 bg-gray-50 p-4 rounded-xl border border-gray-200">
                     <input v-model="profile.is_available" type="checkbox" class="w-5 h-5 text-[#FF3D3D] rounded focus:ring-[#FF3D3D]" />
                     <span class="text-gray-700 font-medium text-sm">I am available for emergency donations</span>
                 </div>
            </div>

            <div class="space-y-1.5">
                <label class="text-xs font-bold text-gray-500 uppercase">Last Donation Date</label>
                <input v-model="profile.last_donation_date" type="date" class="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-3 focus:ring-2 focus:ring-[#FF3D3D]/20 outline-none font-medium" />
                <p class="text-xs text-gray-400">Updating this will automatically calculate your eligibility.</p>
            </div>
            
            <div class="md:col-span-2 pt-4">
                <button type="submit" class="w-full bg-[#FF3D3D] hover:bg-red-600 text-white font-bold py-3.5 rounded-xl shadow-lg shadow-red-500/20 transition-all">
                    Save Changes
                </button>
            </div>
        </form>
      </div>


      <!-- 3. Main Split -->
      <div v-if="!isEditing" class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        
        <!-- Left Sidebar -->
        <div class="space-y-8">
            <!-- Eligibility -->
            <div class="bg-white p-5 rounded-2xl shadow-sm border border-gray-100">
                <h3 class="font-bold text-gray-900 mb-4">Donation Eligibility</h3>
                <div class="flex justify-between items-center text-xs font-bold mb-2">
                    <span class="bg-green-100 text-green-700 px-2 py-0.5 rounded">READY TO DONATE</span>
                    <span class="text-gray-900">100%</span>
                </div>
                <div class="w-full bg-gray-100 rounded-full h-2 mb-4 overflow-hidden">
                    <div class="bg-[#22C55E] h-2 rounded-full" style="width: 100%"></div>
                </div>
                <p class="text-xs text-gray-500 leading-relaxed font-medium">
                    {{ profile.name }} is currently eligible for Whole Blood donation.
                </p>
            </div>

            <!-- About -->
            <div class="bg-white p-5 rounded-2xl shadow-sm border border-gray-100">
                <h3 class="font-bold text-gray-900 mb-4">About {{ profile.name?.split(' ')[0] || 'Donor' }}</h3>
                <p class="text-sm text-gray-500 leading-relaxed mb-6 font-medium">
                    I started donating blood back in college and have tried to keep it up ever since. I'm {{ profile.blood_group }} which means I can help a lot of people!
                </p>
                
                <div class="space-y-4">
                    <div class="flex items-start gap-3">
                        <span class="material-icons text-gray-300">cake</span>
                        <div>
                            <p class="text-xs font-bold text-gray-400 uppercase">Age</p>
                            <p class="text-sm font-bold text-gray-900">{{ age }}</p>
                        </div>
                    </div>
                     <div class="flex items-start gap-3">
                        <span class="material-icons text-gray-300">person</span>
                        <div>
                            <p class="text-xs font-bold text-gray-400 uppercase">Gender</p>
                            <p class="text-sm font-bold text-gray-900">{{ profile.gender || 'Not Specified' }}</p>
                        </div>
                    </div>
                     <div class="flex items-start gap-3">
                        <span class="material-icons text-gray-300">work</span>
                         <div>
                            <p class="text-xs font-bold text-gray-400 uppercase">Occupation</p>
                            <p class="text-sm font-bold text-gray-900">Graphic Designer</p>
                        </div>
                    </div>
                </div>

                <!-- Mini Map Mockup -->
                <div class="mt-6 rounded-xl overflow-hidden h-32 bg-gray-200 relative">
                     <img src="https://static.vecteezy.com/system/resources/previews/000/153/588/original/vector-city-map.jpg" class="w-full h-full object-cover opacity-60 grayscale" />
                     <div class="absolute inset-0 flex items-center justify-center">
                         <span class="material-icons text-[#FF3D3D] text-3xl drop-shadow-md">location_on</span>
                     </div>
                </div>
                <p class="text-xs text-gray-400 mt-2 text-center">{{ profile.area_village || 'Dhaka' }}, {{ profile.district }}</p>
            </div>

            <!-- Achievements -->
             <div class="bg-white p-5 rounded-2xl shadow-sm border border-gray-100">
                <h3 class="font-bold text-gray-900 mb-4">Achievements</h3>
                <div class="flex flex-wrap gap-3">
                    <span class="inline-flex items-center gap-1.5 px-3 py-1.5 bg-yellow-50 text-yellow-700 rounded-lg text-xs font-bold border border-yellow-100">
                        <span class="material-icons text-sm">emoji_events</span> Gold Donor
                    </span>
                    <span class="inline-flex items-center gap-1.5 px-3 py-1.5 bg-blue-50 text-blue-700 rounded-lg text-xs font-bold border border-blue-100">
                        <span class="material-icons text-sm">verified_user</span> ID Verified
                    </span>
                     <span class="inline-flex items-center gap-1.5 px-3 py-1.5 bg-red-50 text-red-700 rounded-lg text-xs font-bold border border-red-100">
                        <span class="material-icons text-sm">flash_on</span> Rapid Responder
                    </span>
                </div>
            </div>
        </div>

        <!-- Right Content -->
        <div class="lg:col-span-2 space-y-8">
            <!-- Donation History -->
            <div class="bg-white p-6 rounded-2xl shadow-sm border border-gray-100">
                <div class="flex justify-between items-center mb-8">
                    <h3 class="font-bold text-gray-900 text-lg">Donation History</h3>
                    <button class="px-3 py-1 rounded-lg border border-gray-200 text-xs font-bold text-gray-500 hover:bg-gray-50">All Time</button>
                </div>

                <!-- Timeline -->
                <div class="relative pl-4 border-l-2 border-gray-100 space-y-12">
                     <div v-if="history.length === 0" class="text-gray-500 italic text-sm">No donation history yet.</div>
                     
                     <!-- Item -->
                    <div v-for="donation in history" :key="donation.id" class="relative">
                        <span class="absolute -left-[21px] top-1 w-3 h-3 rounded-full border-2 border-white ring-2" :class="donation.verified ? 'bg-[#FF3D3D] ring-red-100' : 'bg-gray-300 ring-gray-100'"></span>
                        <div class="bg-gray-50 rounded-2xl p-6 hover:bg-red-50/50 transition-colors group">
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

                <div class="mt-8 text-center space-x-4">
                    <button class="text-[#FF3D3D] text-sm font-bold hover:text-red-700 inline-flex items-center justify-center gap-1">
                        View Older History <span class="material-icons text-sm">arrow_downward</span>
                    </button>
                    <button @click="showAddDonationModal = true" class="text-white bg-[#FF3D3D] hover:bg-red-700 px-4 py-2 rounded-lg text-sm font-bold shadow-md hover:shadow-lg transition-all inline-flex items-center justify-center gap-1">
                        <span class="material-icons text-sm">add</span> Add Donation
                    </button>
                </div>
            </div>

            <!-- Notes -->
             <div class="bg-white p-6 rounded-2xl shadow-sm border border-gray-100">
                <h3 class="font-bold text-gray-900 text-lg mb-6">Staff Notes</h3>
                
                <div class="bg-gray-50 rounded-2xl p-6 flex gap-4">
                <div class="w-10 h-10 rounded-full bg-red-100 flex items-center justify-center shrink-0 text-red-600 font-bold text-xs">MK</div>
                    <div>
                         <div class="flex justify-between items-start mb-1">
                            <h4 class="font-bold text-gray-900 text-sm">Dr. Reynolds</h4>
                            <span class="text-[10px] text-gray-400 font-bold">Aug 12, 2023</span>
                        </div>
                        <p class="text-xs text-gray-500 leading-relaxed font-medium">
                            Excellent donor. Very cooperative and follows pre-donation guidelines perfectly. Veins are easy to access.
                        </p>
                    </div>
                </div>
            </div>

        </div>
      </div>
    </div>

    <!-- Add Donation Modal -->
    <div v-if="showAddDonationModal" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm animate-fade-in">
        <div class="bg-white rounded-2xl shadow-2xl w-full max-w-md overflow-hidden animate-scale-up">
            <div class="bg-red-50 px-6 py-4 flex justify-between items-center border-b border-red-100">
                <h3 class="text-lg font-bold text-gray-900 flex items-center gap-2">
                    <span class="material-icons text-[#FF3D3D]">volunteer_activism</span> Add New Donation
                </h3>
                <button @click="showAddDonationModal = false" class="text-gray-400 hover:text-gray-600 transition-colors">
                    <span class="material-icons">close</span>
                </button>
            </div>
            
            <form @submit.prevent="submitDonation" class="p-6 space-y-4">
                <div class="space-y-1.5">
                    <label class="text-xs font-bold text-gray-500 uppercase">Date</label>
                    <input v-model="newDonation.date" type="date" required class="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-3 focus:ring-2 focus:ring-[#FF3D3D]/20 outline-none font-medium text-gray-800" />
                </div>
                
                <div class="space-y-1.5">
                    <label class="text-xs font-bold text-gray-500 uppercase">Donation Type</label>
                    <select v-model="newDonation.type" required class="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-3 focus:ring-2 focus:ring-[#FF3D3D]/20 outline-none font-medium text-gray-800">
                        <option value="Whole Blood">Whole Blood</option>
                        <option value="Platelets">Platelets</option>
                        <option value="Plasma">Plasma</option>
                    </select>
                </div>

                <div class="space-y-1.5">
                    <label class="text-xs font-bold text-gray-500 uppercase">Location</label>
                    <input v-model="newDonation.location" type="text" placeholder="e.g. Dhaka Medical College Hospital" required class="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-3 focus:ring-2 focus:ring-[#FF3D3D]/20 outline-none font-medium text-gray-800 placeholder-gray-400" />
                </div>

                <div class="grid grid-cols-2 gap-4">
                     <div class="space-y-1.5">
                        <label class="text-xs font-bold text-gray-500 uppercase">Amount (ml)</label>
                        <input v-model="newDonation.amount_ml" type="number" placeholder="450" class="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-3 focus:ring-2 focus:ring-[#FF3D3D]/20 outline-none font-medium text-gray-800 placeholder-gray-400" />
                    </div>
                </div>

                  <div class="space-y-1.5">
                    <label class="text-xs font-bold text-gray-500 uppercase">Notes (Optional)</label>
                    <textarea v-model="newDonation.notes" rows="2" placeholder="Any details..." class="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-3 focus:ring-2 focus:ring-[#FF3D3D]/20 outline-none font-medium text-gray-800 placeholder-gray-400"></textarea>
                </div>

                <div class="space-y-1.5">
                    <label class="text-xs font-bold text-gray-500 uppercase">Evidence Image (Optional)</label>
                    <div class="flex items-center gap-4">
                        <label class="flex-1 cursor-pointer group">
                            <div class="flex items-center justify-center gap-2 px-4 py-3 bg-gray-50 border-2 border-dashed border-gray-200 rounded-xl group-hover:border-[#FF3D3D]/30 group-hover:bg-red-50/30 transition-all">
                                <span class="material-icons text-gray-400 group-hover:text-[#FF3D3D] transition-colors">image</span>
                                <span class="text-sm font-medium text-gray-500 group-hover:text-gray-700">
                                    {{ selectedDonationImage ? selectedDonationImage.name : 'Choose proving image' }}
                                </span>
                            </div>
                            <input type="file" @change="onDonationImageSelected" accept="image/*" class="hidden" />
                        </label>
                        <button v-if="selectedDonationImage" @click="selectedDonationImage = null" type="button" class="p-3 bg-gray-100 text-gray-500 rounded-xl hover:bg-gray-200 transition-colors">
                            <span class="material-icons text-sm">close</span>
                        </button>
                    </div>
                </div>

                <div class="pt-2">
                    <button type="submit" class="w-full bg-[#FF3D3D] hover:bg-red-700 text-white font-bold py-3.5 rounded-xl shadow-lg shadow-red-500/30 transition-all flex items-center justify-center gap-2">
                        <span v-if="isSubmitting" class="material-icons animate-spin text-sm">refresh</span>
                        {{ isSubmitting ? 'Saving...' : 'Save Donation Record' }}
                    </button>
                </div>
            </form>
        </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import api from '@/lib/axios';
import { useToastStore } from '@/stores/toast';
import UserAvatar from '@/components/UserAvatar.vue';

const toastStore = useToastStore();

const isEditing = ref(false);
const isUploading = ref(false);

const profile = ref({
  name: '',
  blood_group: '',
  district: '',
  city: '',
  area_village: '',
  postal_code: '',
  phone: '',
  gender: '',
  birthday: '',
  last_donation_date: '', // Added last_donation_date
  is_available: true,
  profile_picture: ''
});

const stats = ref({
    total_donations: 0,
    lives_saved: 0,
    last_donation: null as string | null
});

// Computed Age
const age = computed(() => {
    if (!profile.value.birthday) return 'N/A';
    const birthDate = new Date(profile.value.birthday);
    const today = new Date();
    let years = today.getFullYear() - birthDate.getFullYear();
    const m = today.getMonth() - birthDate.getMonth();
    if (m < 0 || (m === 0 && today.getDate() < birthDate.getDate())) {
        years--;
    }
    return years + ' years old';
});

interface Donation {
    id: number;
    type: string;
    location: string;
    amount_ml: number;
    notes: string;
    image: string; // ADDED
    verified: boolean;
    date: string;
}

const history = ref<Donation[]>([]);

onMounted(async () => {
    try {
        const res = await api.get('/profile');
        if (res.data) {
             // Handle nested response from new endpoint
             if (res.data.profile) {
                 profile.value = { ...profile.value, ...res.data.profile };
                 // Format birthday for input type="date"
                 if (profile.value.birthday) {
                     const bday = new Date(profile.value.birthday);
                     if (!isNaN(bday.getTime())) {
                        profile.value.birthday = bday.toISOString().split('T')[0] || '';
                     }
                 }
                 // Format last_donation_date for input type="date"
                 if (profile.value.last_donation_date) {
                     const dateVal = profile.value.last_donation_date as unknown as string;
                     const lastDate = new Date(dateVal);
                     if (!isNaN(lastDate.getTime())) {
                        profile.value.last_donation_date = lastDate.toISOString().split('T')[0] || '';
                     }
                 }
             }
             if (res.data.stats) {
                 stats.value = res.data.stats;
             }
             if (res.data.history) {
                 history.value = res.data.history;
             }
        }
    } catch (e) {
        console.log("Could not fetch profile", e);
    }
});

const updateProfile = async () => {
  try {
    const payload = { ...profile.value };
    if (payload.birthday) {
        payload.birthday = new Date(payload.birthday).toISOString();
    }
    if (payload.last_donation_date) {
        payload.last_donation_date = new Date(payload.last_donation_date).toISOString();
    }
    const res = await api.put('/profile', payload);
    toastStore.show('Profile updated successfully!', 'success');
    isEditing.value = false;
  } catch (error) {
    toastStore.show('Failed to update profile. Please try again.', 'error');
  }
};

const handleImageUpload = async (event: Event) => {
    const target = event.target as HTMLInputElement;
    if (!target.files?.length) return;

    const file = target.files[0];
    if (!file) return;
    
    if (file.size > 5 * 1024 * 1024) {
        toastStore.show('Image too large (Max 5MB)', 'error');
        return;
    }

    isUploading.value = true;
    const formData = new FormData();
    formData.append('image', file);

    try {
        const res = await api.post('/profile/picture', formData, {
            headers: {
                'Content-Type': 'multipart/form-data'
            }
        });
        profile.value.profile_picture = res.data.profile_picture;
        toastStore.show('Profile picture updated!', 'success');
    } catch (error) {
        toastStore.show('Failed to upload image', 'error');
    } finally {
        isUploading.value = false;
    }
};

const handleImageDelete = async () => {
    if (!confirm('Are you sure you want to remove your profile picture?')) return;

    try {
        await api.delete('/profile/picture');
        profile.value.profile_picture = '';
        toastStore.show('Profile picture removed', 'success');
    } catch (error) {
        toastStore.show('Failed to remove image', 'error');
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

// Add Donation Logic
const showAddDonationModal = ref(false);
const isSubmitting = ref(false);

interface NewDonation {
    date: string;
    type: string;
    location: string;
    amount_ml: number;
    notes: string;
}

const newDonation = ref<NewDonation>({
    date: new Date().toISOString().split('T')[0] || '',
    type: 'Whole Blood',
    location: '',
    amount_ml: 450,
    notes: ''
});

const getDonationImageUrl = (path: string) => {
    if (!path) return '';
    const baseUrl = import.meta.env.VITE_API_URL 
        ? import.meta.env.VITE_API_URL.replace('/api/v1', '') 
        : 'http://localhost:4000';
    const cleanPath = path.startsWith('/') ? path : `/${path}`;
    return `${baseUrl}${cleanPath}`;
};

const selectedDonationImage = ref<File | null>(null);

const onDonationImageSelected = (e: Event) => {
    const target = e.target as HTMLInputElement;
    if (target.files && target.files[0]) {
        selectedDonationImage.value = target.files[0];
    }
};

const submitDonation = async () => {
    isSubmitting.value = true;
    try {
        const formData = new FormData();
        formData.append('date', new Date(newDonation.value.date).toISOString());
        formData.append('type', newDonation.value.type);
        formData.append('location', newDonation.value.location);
        formData.append('amount_ml', newDonation.value.amount_ml.toString());
        formData.append('notes', newDonation.value.notes);
        
        if (selectedDonationImage.value) {
            formData.append('image', selectedDonationImage.value);
        }

        const res = await api.post('/donations', formData, {
            headers: { 'Content-Type': 'multipart/form-data' }
        });
        
        // Add to local history list immediately
        if (res.data.donation) {
            history.value.unshift(res.data.donation);
            stats.value.total_donations++;
            stats.value.lives_saved = stats.value.total_donations * 3;
            // Update last donation if new one is more recent
            if (!stats.value.last_donation || (res.data.donation.date && new Date(res.data.donation.date) > new Date(stats.value.last_donation))) {
                stats.value.last_donation = res.data.donation.date;
                // Also update profile ref if needed to sync
                if (typeof newDonation.value.date === 'string') {
                    profile.value.last_donation_date = newDonation.value.date;
                    profile.value.is_available = false;
                }
            }
        }
        
        showAddDonationModal.value = false;
        // Reset form
        newDonation.value = {
            date: new Date().toISOString().split('T')[0] || '',
            type: 'Whole Blood',
            location: '',
            amount_ml: 450,
            notes: ''
        };
        selectedDonationImage.value = null; // Clear image
        toastStore.show('Donation record added successfully!', 'success');
    } catch (error) {
        console.error(error);
        toastStore.show('Failed to save donation. Please try again.', 'error');
    } finally {
        isSubmitting.value = false;
    }
};
</script>

<style scoped>
.animate-fade-in-up {
    animation: fadeInUp 0.5s ease-out;
}

@keyframes fadeInUp {
    from {
        opacity: 0;
        transform: translateY(20px);
    }
    to {
        opacity: 1;
        transform: translateY(0);
    }
}

.animate-fade-in {
    animation: fadeIn 0.2s ease-out;
}

.animate-scale-up {
    animation: scaleUp 0.3s cubic-bezier(0.16, 1, 0.3, 1);
}

@keyframes fadeIn {
    from { opacity: 0; }
    to { opacity: 1; }
}

@keyframes scaleUp {
    from { transform: scale(0.95); opacity: 0; }
    to { transform: scale(1); opacity: 1; }
}
</style>
