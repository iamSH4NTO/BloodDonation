<template>
  <div class="min-h-screen bg-[#FAFAFA] font-sans pt-4 pb-8">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8">
      
      <!-- Back Header -->
      <div class="flex items-center justify-between">
        <button @click="$router.push('/admin/donors')" class="flex items-center gap-2 text-gray-500 hover:text-gray-900 font-bold transition-colors">
          <span class="material-icons">arrow_back</span>
          Back to Users
        </button>
        <div class="flex items-center gap-3">
             <span class="text-xs font-bold text-gray-400 uppercase bg-white px-3 py-1 rounded-lg border border-gray-100 shadow-sm">Admin View</span>
        </div>
      </div>

      <!-- Profile Header Card -->
      <div class="bg-white rounded-2xl p-4 sm:p-6 shadow-sm border border-gray-100 flex flex-col sm:flex-row sm:items-start justify-between gap-4">
         <div class="flex items-start gap-3 sm:gap-4 flex-1 min-w-0">
            <!-- Avatar -->
            <div class="relative shrink-0">
                <UserAvatar 
                  :src="profile.profile_picture" 
                  :gender="profile.gender" 
                  :name="profile.name" 
                  size="lg" 
                  class="shadow-lg"
                />
                
                <!-- Admin Upload Overlay -->
                <label v-if="!isUploading" class="absolute inset-0 flex items-center justify-center bg-black/40 rounded-full opacity-0 hover:opacity-100 transition-opacity cursor-pointer">
                    <span class="material-icons text-white text-sm">camera_alt</span>
                    <input type="file" class="hidden" accept="image/*" @change="handleAdminImageUpload" />
                </label>
                <div v-else class="absolute inset-0 flex items-center justify-center bg-black/40 rounded-full opacity-100">
                    <span class="material-icons text-white animate-spin text-sm">refresh</span>
                </div>

                <div class="absolute -bottom-1.5 -right-1.5 bg-[#FF3D3D] text-white w-7 h-7 sm:w-8 sm:h-8 rounded-full flex items-center justify-center font-bold text-xs sm:text-sm shadow-md border-2 border-white">
                    {{ profile.blood_group || '?' }}
                </div>
            </div>

            <!-- Info -->
            <div class="flex-1 min-w-0 space-y-2">
                <div class="flex flex-col gap-2">
                    <h1 class="text-lg sm:text-xl font-bold text-gray-900 truncate">{{ profile.name || 'Donor Profile' }}</h1>
                    <div class="flex items-center gap-1.5 flex-wrap">
                        <span :class="userAccount.is_active ? 'bg-emerald-50 text-emerald-700' : 'bg-red-50 text-red-700'" class="px-2 py-0.5 rounded-md text-[9px] sm:text-[10px] font-bold uppercase flex items-center gap-1">
                            <span class="w-1 h-1 rounded-full" :class="userAccount.is_active ? 'bg-emerald-500' : 'bg-red-500'"></span>
                            {{ userAccount.is_active ? 'Active' : 'Banned' }}
                        </span>
                        <span class="bg-purple-50 text-purple-700 px-2 py-0.5 rounded-md text-[9px] sm:text-[10px] font-bold uppercase">
                            {{ userAccount.role }}
                        </span>
                    </div>
                </div>
                
                <div class="flex flex-col gap-1 text-[11px] sm:text-xs text-gray-500">
                    <span class="flex items-center gap-1 truncate"><span class="material-icons text-[14px] shrink-0">email</span> <span class="truncate">{{ userAccount.email }}</span></span>
                    <span class="flex items-center gap-1"><span class="material-icons text-[14px] shrink-0">id_card</span> <span class="font-bold text-gray-400">{{ userAccount.id }}</span></span>
                    <span class="flex items-center gap-1 truncate"><span class="material-icons text-[14px] shrink-0">location_on</span> <span class="truncate">{{ profile.city || 'N/A' }}, {{ profile.district || 'N/A' }}</span></span>
                </div>
            </div>
         </div>
         
         <!-- Save Button (Relative positioning via Flexbox) -->
         <div class="w-full sm:w-auto shrink-0 pt-0 sm:pt-2">
             <button @click="saveChanges" :disabled="isSaving" class="w-full sm:w-auto px-4 sm:px-6 py-2.5 sm:py-3 rounded-xl bg-[#FF3D3D] hover:bg-red-600 text-white font-bold text-xs sm:text-sm shadow-lg shadow-red-500/20 transition-all flex items-center justify-center gap-2 disabled:opacity-70">
                 <span v-if="isSaving" class="material-icons animate-spin text-sm">refresh</span>
                 <span v-else class="material-icons text-sm">save</span>
                 {{ isSaving ? 'Saving...' : 'Save Changes' }}
             </button>
         </div>
      </div>

      <!-- Main Content Tabs -->
      <div class="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
        <div class="px-3 sm:px-4 border-b border-gray-100 flex gap-1 bg-white overflow-x-auto scrollbar-hide">
            <button 
                v-for="tab in ['profile', 'account', 'history', 'logs']" 
                :key="tab"
                @click="activeTab = tab"
                class="text-[11px] sm:text-xs font-bold py-3 px-3 sm:px-4 transition-all border-b-2 capitalize whitespace-nowrap"
                :class="activeTab === tab ? 'text-[#FF3D3D] border-[#FF3D3D]' : 'text-gray-400 border-transparent'"
            >
                {{ tab === 'history' ? 'History' : tab === 'logs' ? 'Logs' : tab }}
            </button>
        </div>

        <div class="p-4 sm:p-6">
            <!-- Profile Details Tab -->
            <div v-if="activeTab === 'profile'" class="space-y-6 animate-fade-in">
                <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
                    <div class="space-y-1.5">
                        <label class="text-xs font-bold text-gray-500 uppercase">Full Name</label>
                        <input v-model="profile.name" type="text" class="input-field" />
                    </div>
                    <div class="space-y-1.5">
                        <label class="text-xs font-bold text-gray-500 uppercase">Phone Number</label>
                        <input v-model="profile.phone" type="text" class="input-field" />
                    </div>
                    <div class="space-y-1.5">
                        <label class="text-xs font-bold text-gray-500 uppercase">Blood Group</label>
                        <select v-model="profile.blood_group" class="input-field">
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
                        <label class="text-xs font-bold text-gray-500 uppercase">Gender</label>
                        <select v-model="profile.gender" class="input-field">
                            <option value="Male">Male</option>
                            <option value="Female">Female</option>
                            <option value="Other">Other</option>
                        </select>
                    </div>
                    <div class="space-y-1.5">
                        <label class="text-xs font-bold text-gray-500 uppercase">Birthday</label>
                        <input v-model="profile.birthday" type="date" class="input-field" />
                    </div>
                    <div class="space-y-1.5">
                        <label class="text-xs font-bold text-gray-500 uppercase">Last Donation Date</label>
                        <input v-model="profile.last_donation_date" type="date" class="input-field" />
                    </div>
                </div>

                <div class="pt-4 border-t border-gray-100">
                    <h3 class="text-xs sm:text-sm font-bold text-gray-700 mb-3 flex items-center gap-1.5">
                        <span class="material-icons text-gray-400 text-sm">place</span> Location
                    </h3>
                    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
                        <div class="space-y-1.5">
                            <label class="text-xs font-bold text-gray-500 uppercase">District</label>
                            <input v-model="profile.district" type="text" class="input-field" />
                        </div>
                        <div class="space-y-1.5">
                            <label class="text-xs font-bold text-gray-500 uppercase">Upazila</label>
                            <input v-model="profile.upazila" type="text" class="input-field" />
                        </div>
                        <div class="space-y-1.5">
                            <label class="text-xs font-bold text-gray-500 uppercase">City</label>
                            <input v-model="profile.city" type="text" class="input-field" />
                        </div>
                        <div class="space-y-1.5">
                            <label class="text-xs font-bold text-gray-500 uppercase">Area / Village</label>
                            <input v-model="profile.area_village" type="text" class="input-field" />
                        </div>
                        <div class="space-y-1.5">
                            <label class="text-xs font-bold text-gray-500 uppercase">Postal Code</label>
                            <input v-model="profile.postal_code" type="text" class="input-field" />
                        </div>
                        <div class="space-y-1.5">
                            <label class="text-xs font-bold text-gray-500 uppercase">Availability</label>
                            <select v-model="profile.is_available" class="input-field font-bold" :class="profile.is_available ? 'text-emerald-600' : 'text-red-500'">
                                <option :value="true">Available</option>
                                <option :value="false">Unavailable</option>
                            </select>
                        </div>
                        <div class="space-y-1.5 min-w-0">
                            <label class="text-xs font-bold text-gray-500 uppercase">Verification</label>
                            <label class="flex items-center gap-2 mt-2 cursor-pointer bg-white border border-gray-100 p-2.5 rounded-xl shadow-sm hover:border-gray-300 transition-colors">
                                <div class="relative">
                                    <input type="checkbox" v-model="profile.is_admin_verified" class="sr-only">
                                    <div class="block bg-gray-200 w-10 h-6 rounded-full transition-colors" :class="{'bg-[#FF3D3D]': profile.is_admin_verified}"></div>
                                    <div class="dot absolute left-1 top-1 bg-white w-4 h-4 rounded-full transition-transform" :class="{'transform translate-x-4': profile.is_admin_verified}"></div>
                                </div>
                                <span class="text-sm font-bold text-gray-700 select-none">Admin Verified</span>
                            </label>
                        </div>
                        <div class="space-y-1.5 sm:col-span-2 mt-2 lg:col-span-3">
                            <label class="text-xs font-bold text-gray-500 uppercase">Google Map Link</label>
                            <input v-model="profile.google_map_link" type="text" class="input-field" placeholder="https://maps.google.com/..." />
                        </div>
                        
                        <!-- Social Links Panel -->
                        <div class="sm:col-span-2 lg:col-span-3 mt-4 bg-gray-50 p-4 rounded-xl border border-gray-100">
                            <h3 class="text-xs font-bold text-gray-900 mb-3 uppercase tracking-wider"><i class="fas fa-link text-gray-400 mr-2"></i> Social Profiles</h3>
                            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                                <div class="space-y-1.5">
                                    <label class="text-[10px] font-bold text-gray-500 uppercase flex items-center gap-1"><i class="fab fa-facebook text-blue-600"></i> Facebook Link</label>
                                    <input v-model="profile.facebook_link" type="url" placeholder="https://facebook.com/..." class="input-field text-sm" />
                                </div>
                                <div class="space-y-1.5">
                                    <label class="text-[10px] font-bold text-gray-500 uppercase flex items-center gap-1"><i class="fab fa-instagram text-pink-600"></i> Instagram Link</label>
                                    <input v-model="profile.instagram_link" type="url" placeholder="https://instagram.com/..." class="input-field text-sm" />
                                </div>
                                <div class="space-y-1.5">
                                    <label class="text-[10px] font-bold text-gray-500 uppercase flex items-center gap-1"><i class="fab fa-linkedin text-sky-600"></i> LinkedIn Link</label>
                                    <input v-model="profile.linkedin_link" type="url" placeholder="https://linkedin.com/in/..." class="input-field text-sm" />
                                </div>
                                <div class="space-y-1.5">
                                    <label class="text-[10px] font-bold text-gray-500 uppercase flex items-center gap-1"><i class="fab fa-youtube text-red-600"></i> YouTube Link</label>
                                    <input v-model="profile.youtube_link" type="url" placeholder="https://youtube.com/@..." class="input-field text-sm" />
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Account Settings Tab -->
            <div v-if="activeTab === 'account'" class="max-w-2xl space-y-6 animate-fade-in">
                <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                    <div class="space-y-6">
                        <div class="space-y-1.5">
                            <label class="text-xs font-bold text-gray-500 uppercase">Email Address</label>
                            <input v-model="userAccount.email" type="email" class="input-field" />
                        </div>
                        <div class="space-y-1.5">
                            <label class="text-xs font-bold text-gray-500 uppercase">Account Status</label>
                            <div class="flex gap-2">
                                <button 
                                    @click="userAccount.is_active = true"
                                    type="button"
                                    class="flex-1 py-3 rounded-xl text-xs font-bold border transition-all"
                                    :class="userAccount.is_active ? 'bg-emerald-50 text-emerald-700 border-emerald-200' : 'bg-white text-gray-400 border-gray-100'"
                                >Active</button>
                                <button 
                                    @click="userAccount.is_active = false"
                                    type="button"
                                    class="flex-1 py-3 rounded-xl text-xs font-bold border transition-all"
                                    :class="!userAccount.is_active ? 'bg-red-50 text-red-700 border-red-200' : 'bg-white text-gray-400 border-gray-100'"
                                >Banned</button>
                            </div>
                        </div>
                    </div>

                    <div class="space-y-6">
                        <div class="space-y-1.5">
                            <label class="text-xs font-bold text-gray-500 uppercase">User Role</label>
                            <select v-model="userAccount.role" class="input-field">
                                <option value="donor">Donor</option>
                                <option value="admin">Admin</option>
                            </select>
                        </div>
                        <div class="space-y-1.5">
                            <label class="text-xs font-bold text-gray-500 uppercase">New Password</label>
                            <input v-model="passwordForm" type="password" placeholder="Leave blank to keep current" class="input-field" />
                        </div>
                    </div>
                </div>
            </div>

            <!-- Donation History Tab -->
            <div v-if="activeTab === 'history'" class="animate-fade-in">
                <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-2 mb-4">
                    <h3 class="text-sm font-bold text-gray-900">Records</h3>
                    <button @click="openAddDonation" class="w-full sm:w-auto bg-[#FF3D3D] text-white px-3 py-2 rounded-lg text-[11px] sm:text-xs font-bold hover:bg-red-600 transition-colors flex items-center justify-center gap-1 shadow-md shadow-red-500/20">
                        <span class="material-icons text-sm">add</span> Add
                    </button>
                </div>

                <div class="space-y-3">
                    <div v-for="h in history" :key="h.id" class="p-3 sm:p-4 bg-gray-50 rounded-xl border border-gray-100 flex flex-col sm:flex-row items-start sm:items-center justify-between gap-2 group hover:bg-white hover:shadow-sm transition-all">
                        <div class="flex items-center gap-2 sm:gap-3 flex-1 min-w-0">
                            <div class="w-9 h-9 sm:w-10 sm:h-10 rounded-xl bg-red-100 flex items-center justify-center text-red-600 shrink-0">
                                <span class="material-icons text-lg">water_drop</span>
                            </div>
                            <div class="min-w-0 flex-1">
                                <div class="font-bold text-gray-900 text-xs sm:text-sm truncate">{{ formatDateFull(h.date) }}</div>
                                <div class="text-[11px] text-gray-500 truncate">{{ h.location }} • {{ h.type }}</div>
                                
                                <div v-if="h.notes" class="mt-1 text-[10px] text-gray-400 italic">"{{ h.notes }}"</div>

                                <!-- Donation Proof Image -->
                                <div v-if="h.image" class="mt-2 rounded-lg overflow-hidden border border-gray-100 bg-white inline-block">
                                    <a :href="getDonationImageUrl(h.image)" target="_blank" class="block">
                                        <img :src="getDonationImageUrl(h.image)" alt="Proof" class="h-12 w-auto object-cover hover:scale-105 transition-transform" />
                                    </a>
                                </div>
                            </div>
                        </div>
                        <button @click="deleteDonation(h.id)" class="sm:opacity-0 sm:group-hover:opacity-100 w-full sm:w-9 h-9 rounded-lg flex items-center justify-center text-gray-400 hover:text-red-600 hover:bg-red-50 transition-all shrink-0">
                            <span class="material-icons text-sm">delete</span>
                        </button>
                    </div>
                    <div v-if="history.length === 0" class="text-center py-8 bg-gray-50 rounded-xl border-2 border-dashed border-gray-100">
                        <span class="material-icons text-gray-300 text-3xl mb-2">event_busy</span>
                        <div class="text-gray-400 text-xs">No records found.</div>
                    </div>
                </div>
            </div>

            <!-- View Logs Tab -->
            <div v-if="activeTab === 'logs'" class="animate-fade-in">
                <h3 class="text-sm font-bold text-gray-900 mb-4">Access Logs</h3>
                <div class="space-y-3">
                    <div v-for="log in logs" :key="log.id" class="p-3 sm:p-4 bg-white border border-gray-100 rounded-xl shadow-sm hover:shadow-md transition-shadow">
                        <div class="flex flex-col sm:flex-row items-start justify-between gap-2 mb-3">
                            <div class="flex items-center gap-2 sm:gap-3 flex-1 min-w-0">
                                <div class="w-9 h-9 sm:w-10 sm:h-10 rounded-xl bg-blue-50 flex items-center justify-center text-blue-600 shrink-0">
                                    <span class="material-icons text-lg">person_search</span>
                                </div>
                                <div class="min-w-0 flex-1">
                                    <div class="font-bold text-gray-900 text-xs sm:text-sm truncate">{{ log.viewer_name }}</div>
                                    <div class="flex flex-col sm:flex-row sm:items-center gap-0.5 sm:gap-1.5 mt-0.5">
                                        <div class="text-[9px] sm:text-[10px] text-gray-400 font-bold uppercase truncate">{{ log.unique_id }}</div>
                                        <div class="hidden sm:block text-[10px] text-gray-400">•</div>
                                        <div class="text-[9px] sm:text-[10px] text-gray-400 truncate">{{ new Date(log.created_at).toLocaleString() }}</div>
                                    </div>
                                </div>
                            </div>
                            <button @click="viewUserProfile(log)" class="w-full sm:w-auto px-3 py-1.5 bg-blue-50 text-blue-600 text-[10px] sm:text-xs font-bold rounded-lg hover:bg-blue-100 transition-colors flex items-center justify-center gap-1 shrink-0">
                                <span class="material-icons text-sm">visibility</span>
                                View
                            </button>
                        </div>
                        
                        <div class="grid grid-cols-2 md:grid-cols-4 gap-2 sm:gap-3">
                            <div class="bg-gray-50 p-2 rounded-lg border border-gray-100">
                                <div class="text-[9px] text-gray-400 font-bold uppercase mb-0.5">Blood</div>
                                <div class="text-[11px] sm:text-xs font-bold text-red-600">{{ log.blood_group || 'N/A' }}</div>
                            </div>
                            <div class="bg-gray-50 p-2 rounded-lg border border-gray-100">
                                <div class="text-[9px] text-gray-400 font-bold uppercase mb-0.5">Phone</div>
                                <div class="text-[11px] sm:text-xs font-bold text-gray-700 truncate">{{ log.phone || 'N/A' }}</div>
                            </div>
                            <div class="bg-gray-50 p-2 rounded-lg border border-gray-100">
                                <div class="text-[9px] text-gray-400 font-bold uppercase mb-0.5">Location</div>
                                <div class="text-[11px] sm:text-xs font-bold text-gray-700 truncate">{{ log.district || 'N/A' }}</div>
                            </div>
                            <div class="bg-gray-50 p-2 rounded-lg border border-gray-100">
                                <div class="text-[9px] text-gray-400 font-bold uppercase mb-0.5">IP</div>
                                <div class="text-[11px] sm:text-xs font-mono text-gray-500 truncate">{{ log.ip_address || 'Unknown' }}</div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
      </div>
    </div>

    <!-- Modal for Adding Donation -->
    <div v-if="isHistoryModalOpen" class="fixed inset-0 z-50 flex items-center justify-center p-3 bg-black/60 backdrop-blur-sm">
         <div class="bg-white rounded-2xl shadow-2xl w-full max-w-md max-h-[90vh] overflow-hidden flex flex-col">
            <div class="px-4 py-3 border-b border-gray-100 flex justify-between items-center shrink-0">
                <h3 class="text-sm font-bold text-gray-800">Add Record</h3>
                <button @click="isHistoryModalOpen = false" class="w-8 h-8 rounded-lg hover:bg-gray-100 flex items-center justify-center transition-colors"><span class="material-icons text-gray-400 text-lg">close</span></button>
            </div>
            <div class="p-4 space-y-3 overflow-y-auto">
                <div class="space-y-1.5">
                    <label class="text-xs font-bold text-gray-500 uppercase">Date</label>
                    <input v-model="hForm.date" type="date" class="input-field" />
                </div>
                <div class="space-y-1.5">
                    <label class="text-xs font-bold text-gray-500 uppercase">Type</label>
                    <select v-model="hForm.type" class="input-field">
                        <option value="Whole Blood">Whole Blood</option>
                        <option value="Platelets">Platelets</option>
                    </select>
                </div>
                <div class="space-y-1.5">
                    <label class="text-xs font-bold text-gray-500 uppercase">Location</label>
                    <input v-model="hForm.location" type="text" class="input-field" />
                </div>
                <div class="space-y-1.5">
                    <label class="text-xs font-bold text-gray-500 uppercase">Amount (ml)</label>
                    <input v-model="hForm.amount_ml" type="number" class="input-field" placeholder="450" />
                </div>
                <div class="space-y-1.5">
                    <label class="text-xs font-bold text-gray-500 uppercase">Notes</label>
                    <textarea v-model="hForm.notes" rows="2" class="input-field" placeholder="Optional notes..."></textarea>
                </div>
                <div class="space-y-1.5">
                    <label class="text-xs font-bold text-gray-500 uppercase">Evidence Image</label>
                    <input type="file" @change="onAdminDonationImageSelected" accept="image/*" class="w-full text-xs text-gray-500 file:mr-2 file:py-2 file:px-3 file:rounded-lg file:border-0 file:text-xs file:font-semibold file:bg-red-50 file:text-red-700 hover:file:bg-red-100 transition-all"/>
                </div>
                <button @click="addHistory" class="w-full py-2.5 bg-[#FF3D3D] text-white font-bold text-xs rounded-xl mt-2 shadow-lg shadow-red-500/20 hover:bg-red-600 transition-colors flex items-center justify-center gap-1">
                    <span v-if="isUploading" class="material-icons animate-spin text-sm">refresh</span>
                    {{ isUploading ? 'Saving...' : 'Add Record' }}
                </button>
            </div>
         </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive, watch, computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import api from '../../lib/axios';
import { useToastStore } from '@/stores/toast';
import UserAvatar from '@/components/UserAvatar.vue';

const toastStore = useToastStore();

const route = useRoute();
const router = useRouter();
interface UserAccount {
    id: string;
    email: string;
    role: string;
    is_active: boolean;
}

interface Profile {
    name: string;
    blood_group: string;
    phone: string;
    gender: string;
    birthday: string;
    district: string;
    upazila: string;
    city: string;
    area_village: string;
    postal_code: string;
    google_map_link: string;
    last_donation_date: string;
    is_available: boolean;
    is_admin_verified: boolean;
    profile_picture: string;
    facebook_link?: string;
    instagram_link?: string;
    linkedin_link?: string;
    youtube_link?: string;
}

interface Donation {
    id: number;
    date: string;
    location: string;
    type: string;
    amount_ml: number;
    notes: string;
    image: string;
}

interface ViewLog {
    id: number;
    viewer_id: string;
    viewer_name: string;
    unique_id: string;
    created_at: string;
    blood_group: string;
    phone: string;
    district: string;
    ip_address: string;
}

const donorId = computed(() => route.params.id as string);

const activeTab = ref('profile');
const isSaving = ref(false);
const isUploading = ref(false);
const isHistoryModalOpen = ref(false);

const handleAdminImageUpload = async (event: Event) => {
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
        const res = await api.post(`/admin/users/${donorId.value}/profile/picture`, formData, {
            headers: {
                'Content-Type': 'multipart/form-data'
            }
        });
        profile.value.profile_picture = res.data.profile_picture;
        toastStore.show('Profile picture updated by admin!', 'success');
    } catch (error) {
        toastStore.show('Failed to upload image', 'error');
    } finally {
        isUploading.value = false;
    }
};

const userAccount = ref<UserAccount>({
    id: '',
    email: '',
    role: '',
    is_active: true
});

const profile = ref<Profile>({
    name: '',
    blood_group: '',
    phone: '',
    gender: '',
    birthday: '',
    district: '',
    upazila: '',
    city: '',
    area_village: '',
    postal_code: '',
    google_map_link: '',
    last_donation_date: '',
    is_available: true,
    is_admin_verified: false,
    profile_picture: '',
    facebook_link: '',
    instagram_link: '',
    linkedin_link: '',
    youtube_link: '',
});

const stats = ref({
    total_donations: 0,
    lives_saved: 0,
    last_donation: null as string | null
});

const history = ref<Donation[]>([]);
const logs = ref<ViewLog[]>([]);
const passwordForm = ref('');

const hForm = reactive({
    date: new Date().toISOString().split('T')[0],
    type: 'Whole Blood',
    location: '',
    amount_ml: 450,
    notes: ''
});

const selectedAdminDonationImage = ref<File | null>(null);

const onAdminDonationImageSelected = (e: Event) => {
    const target = e.target as HTMLInputElement;
    if (target.files && target.files[0]) {
        selectedAdminDonationImage.value = target.files[0];
    }
};

const fetchData = async () => {
    const currentDonorId = route.params.id as string;
    try {
        const res = await api.get(`/admin/users/${currentDonorId}`);
        const data = res.data;
        
        userAccount.value = {
            id: data.user.id,
            email: data.user.email,
            role: data.user.role,
            is_active: data.user.is_active
        };

        if (data.user.donor_profile) {
            const p = data.user.donor_profile;
            const bdayStr = (p.birthday && p.birthday !== '') ? new Date(p.birthday).toISOString().split('T')?.[0] || '' : '';
            const lastDonationStr = (p.last_donation_date && p.last_donation_date !== '') ? new Date(p.last_donation_date).toISOString().split('T')?.[0] || '' : '';

            profile.value = {
                name: p.name || '',
                phone: p.phone || '',
                blood_group: p.blood_group || '',
                gender: p.gender || '',
                birthday: bdayStr,
                district: p.district || '',
                upazila: p.upazila || '',
                city: p.city || '',
                area_village: p.area_village || '',
                postal_code: p.postal_code || '',
                google_map_link: p.google_map_link || '',
                last_donation_date: lastDonationStr,
                is_available: p.is_available ?? true,
                is_admin_verified: p.is_admin_verified ?? false,
                profile_picture: p.profile_picture || '',
                facebook_link: p.facebook_link || '',
                instagram_link: p.instagram_link || '',
                linkedin_link: p.linkedin_link || '',
                youtube_link: p.youtube_link || ''
            };
        }
        
        stats.value = data.stats;
        history.value = data.history;
        
        // Also fetch view logs
        const logRes = await api.get(`/admin/users/${currentDonorId}/view-logs`);
        logs.value = logRes.data;
        
    } catch (error) {
        console.error("Failed to fetch user details", error);
        toastStore.show("Could not load user data", "error");
    }
};

const saveChanges = async () => {
    isSaving.value = true;
    try {
        const payload = {
            name: profile.value.name,
            role: userAccount.value.role,
            is_active: userAccount.value.is_active,
            phone: profile.value.phone,
            bloodGroup: profile.value.blood_group,
            gender: profile.value.gender,
            birthday: profile.value.birthday ? new Date(profile.value.birthday).toISOString() : null,
            district: profile.value.district,
            upazila: profile.value.upazila,
            city: profile.value.city,
            area_village: profile.value.area_village,
            postal_code: profile.value.postal_code,
            google_map_link: profile.value.google_map_link,
            last_donation_date: profile.value.last_donation_date ? new Date(profile.value.last_donation_date).toISOString() : null,
            is_available: profile.value.is_available,
            is_admin_verified: profile.value.is_admin_verified,
            facebook_link: profile.value.facebook_link,
            instagram_link: profile.value.instagram_link,
            linkedin_link: profile.value.linkedin_link,
            youtube_link: profile.value.youtube_link,
            password: passwordForm.value || undefined
        };

        // Use donorId.value to get the string ID
        await api.put(`/admin/users/${donorId.value}`, payload);
        toastStore.show("Changes saved successfully!", "success");
        passwordForm.value = '';
    } catch (error) {
        console.error("Failed to save changes", error);
        toastStore.show("Failed to save user data", "error");
    } finally {
        isSaving.value = false;
    }
};

const openAddDonation = () => {
    isHistoryModalOpen.value = true;
};

const addHistory = async () => {
    const currentDonorId = route.params.id as string;
    isUploading.value = true;
    try {
        const formData = new FormData();
        formData.append('date', new Date(hForm.date as string).toISOString());
        formData.append('type', hForm.type);
        formData.append('location', hForm.location);
        formData.append('amount_ml', hForm.amount_ml.toString());
        formData.append('notes', hForm.notes);

        if (selectedAdminDonationImage.value) {
            formData.append('image', selectedAdminDonationImage.value);
        }

        await api.post(`/admin/users/${currentDonorId}/donations`, formData, {
            headers: { 'Content-Type': 'multipart/form-data' }
        });
        
        await fetchData();
        isHistoryModalOpen.value = false;
        // Reset form
        hForm.location = '';
        hForm.notes = '';
        hForm.amount_ml = 450;
        selectedAdminDonationImage.value = null;
        toastStore.show("Record added successfully", "success");
    } catch (error) {
        toastStore.show("Failed to add history", "error");
    } finally {
        isUploading.value = false;
    }
};

const deleteDonation = async (id: number) => {
    if (!confirm("Are you sure?")) return;
    try {
        await api.delete(`/admin/donations/${id}`);
        fetchData();
    } catch (error) {
        toastStore.show("Failed to delete record", "error");
    }
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

const viewUserProfile = (log: any) => {
    console.log('Attempting to view profile for log:', log);
    console.log('unique_id:', log.unique_id);
    console.log('viewer_id:', log.viewer_id);
    
    const userId = log.unique_id || log.viewer_id;
    if (!userId) {
        toastStore.show('User ID not found in log data', 'error');
        return;
    }
    
    console.log('Navigating to:', `/admin/donors/${userId}`);
    router.push(`/admin/donors/${userId}`);
};

onMounted(fetchData);

// Watch for route changes to refetch data when navigating between users
watch(() => route.params.id, (newId, oldId) => {
    if (newId && newId !== oldId) {
        fetchData();
    }
});
</script>

<style scoped>
.animate-fade-in {
    animation: fadeIn 0.4s ease-out;
}

@keyframes fadeIn {
    from { opacity: 0; transform: translateY(10px); }
    to { opacity: 1; transform: translateY(0); }
}
</style>
