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
      <div class="bg-white rounded-3xl p-6 sm:p-8 shadow-sm border border-gray-100 relative overflow-visible">
         <div class="flex flex-col md:flex-row gap-6 md:gap-8 items-center md:items-center text-center md:text-left">
            <!-- Avatar -->
            <div class="relative group">
                <div class="w-24 h-24 md:w-32 md:h-32 rounded-3xl overflow-hidden border-4 border-white shadow-xl bg-red-100 flex items-center justify-center">
                    <span v-if="!profile.name" class="material-icons text-red-300 text-5xl md:text-6xl">person</span>
                    <span v-else class="text-4xl md:text-5xl font-black text-red-400">{{ profile.name.charAt(0) }}</span>
                </div>
                <div class="absolute -bottom-3 -right-3 bg-[#FF3D3D] text-white w-10 h-10 md:w-12 md:h-12 rounded-full flex items-center justify-center font-bold text-base md:text-lg shadow-md border-4 border-white">
                    {{ profile.blood_group || '?' }}
                </div>
            </div>

            <!-- Info -->
            <div class="flex-1 space-y-3 w-full">
                <div class="flex items-center justify-center md:justify-start gap-3 flex-wrap">
                    <h1 class="text-2xl md:text-3xl font-black text-gray-900 tracking-tight">{{ profile.name || 'Donor Profile' }}</h1>
                    <div class="flex items-center gap-2">
                        <span :class="userAccount.is_active ? 'bg-emerald-50 text-emerald-700 border-emerald-100' : 'bg-red-50 text-red-700 border-red-100'" class="px-3 py-1 rounded-full text-[10px] font-bold uppercase tracking-wide flex items-center gap-1 border">
                            <span class="w-1.5 h-1.5 rounded-full" :class="userAccount.is_active ? 'bg-emerald-500' : 'bg-red-500'"></span>
                            {{ userAccount.is_active ? 'Active' : 'Banned' }}
                        </span>
                        <span class="bg-purple-50 text-purple-700 border-purple-100 px-3 py-1 rounded-full text-[10px] font-bold uppercase border">
                            {{ userAccount.role }}
                        </span>
                    </div>
                </div>
                
                <div class="flex flex-col sm:flex-row items-center justify-center md:justify-start gap-3 sm:gap-6 text-gray-500 text-sm font-medium">
                    <span class="flex items-center gap-1.5"><span class="material-icons text-gray-400 text-sm">email</span> {{ userAccount.email }}</span>
                    <span class="flex items-center gap-1.5"><span class="material-icons text-gray-400 text-sm">id_card</span> <span class="uppercase font-bold text-gray-400">{{ userAccount.id }}</span></span>
                </div>

                <div class="flex items-center justify-center md:justify-start gap-4 text-gray-500 text-sm font-medium">
                    <span class="flex items-center gap-1.5"><span class="material-icons text-gray-400 text-sm">location_on</span> {{ profile.city || 'N/A' }}, {{ profile.district || 'N/A' }}</span>
                </div>
            </div>

            <!-- Actions -->
            <div class="w-full md:w-auto mt-4 md:mt-0">
                <button @click="saveChanges" :disabled="isSaving" class="w-full md:w-auto px-8 py-4 rounded-xl bg-[#FF3D3D] hover:bg-red-600 text-white font-bold text-sm shadow-xl shadow-red-500/30 transition-all flex items-center justify-center gap-2 transform hover:-translate-y-1 active:scale-95 disabled:opacity-70">
                    <span v-if="isSaving" class="material-icons animate-spin text-sm">refresh</span>
                    <span v-else class="material-icons text-sm">save</span>
                    {{ isSaving ? 'Saving...' : 'Save Changes' }}
                </button>
            </div>
         </div>
      </div>

      <!-- Main Content Tabs -->
      <div class="bg-white rounded-3xl shadow-sm border border-gray-100 overflow-hidden flex flex-col">
        <div class="px-6 sm:px-8 border-b border-gray-100 flex gap-6 sm:gap-8 bg-white overflow-x-auto scrollbar-hide mask-fade-right">
            <button 
                v-for="tab in ['profile', 'account', 'history', 'logs']" 
                :key="tab"
                @click="activeTab = tab"
                class="text-xs sm:text-sm font-bold pt-6 pb-4 transition-all border-b-2 capitalize whitespace-nowrap"
                :class="activeTab === tab ? 'text-[#FF3D3D] border-[#FF3D3D]' : 'text-gray-400 border-transparent hover:text-gray-600'"
            >
                {{ tab === 'history' ? 'Donation History' : tab === 'logs' ? 'View Logs' : tab + ' Details' }}
            </button>
        </div>

        <div class="p-6 sm:p-8">
            <!-- Profile Details Tab -->
            <div v-if="activeTab === 'profile'" class="space-y-8 animate-fade-in">
                <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
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

                <div class="pt-6 border-t border-gray-50">
                    <h3 class="text-sm font-bold text-gray-900 mb-4 flex items-center gap-2">
                        <span class="material-icons text-gray-400 text-sm">place</span> Location Information
                    </h3>
                    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
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
                        <div class="space-y-1.5 md:col-span-3">
                            <label class="text-xs font-bold text-gray-500 uppercase">Google Map Link</label>
                            <input v-model="profile.google_map_link" type="text" class="input-field" placeholder="https://maps.google.com/..." />
                        </div>
                    </div>
                </div>
            </div>

            <!-- Account Settings Tab -->
            <div v-if="activeTab === 'account'" class="max-w-2xl space-y-8 animate-fade-in">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
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
                                    class="flex-1 py-2 rounded-xl text-xs font-bold border transition-all"
                                    :class="userAccount.is_active ? 'bg-emerald-50 text-emerald-700 border-emerald-200' : 'bg-white text-gray-400 border-gray-100'"
                                >Active</button>
                                <button 
                                    @click="userAccount.is_active = false"
                                    type="button"
                                    class="flex-1 py-2 rounded-xl text-xs font-bold border transition-all"
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
                <div class="flex justify-between items-center mb-6">
                    <h3 class="font-bold text-gray-900">Donation Records</h3>
                    <button @click="openAddDonation" class="bg-[#FF3D3D] text-white px-4 py-2 rounded-xl text-xs font-bold hover:bg-red-600 transition-colors flex items-center gap-1.5 shadow-md shadow-red-500/20">
                        <span class="material-icons text-sm">add</span> Add Record
                    </button>
                </div>

                <div class="space-y-4">
                    <div v-for="h in history" :key="h.id" class="p-5 bg-gray-50 rounded-2xl border border-gray-100 flex items-center justify-between group hover:bg-white hover:shadow-md transition-all">
                        <div class="flex items-center gap-4">
                            <div class="w-12 h-12 rounded-2xl bg-red-100 flex items-center justify-center text-red-600">
                                <span class="material-icons">water_drop</span>
                            </div>
                            <div>
                                <div class="font-bold text-gray-900">{{ formatDateFull(h.date) }}</div>
                                <div class="text-xs text-gray-500 font-medium">{{ h.location }} • {{ h.type }}</div>
                            </div>
                        </div>
                        <button @click="deleteDonation(h.id)" class="opacity-0 group-hover:opacity-100 w-10 h-10 rounded-xl flex items-center justify-center text-gray-400 hover:text-red-600 hover:bg-red-50 transition-all">
                            <span class="material-icons text-sm">delete</span>
                        </button>
                    </div>
                    <div v-if="history.length === 0" class="text-center py-12 bg-gray-50 rounded-2xl border-2 border-dashed border-gray-100">
                        <span class="material-icons text-gray-300 text-4xl mb-2">event_busy</span>
                        <div class="text-gray-400 text-sm font-medium">No donation history recorded.</div>
                    </div>
                </div>
            </div>

            <!-- View Logs Tab -->
            <div v-if="activeTab === 'logs'" class="animate-fade-in">
                <h3 class="font-bold text-gray-900 mb-6">Profile Access Information</h3>
                <div class="space-y-4">
                    <div v-for="log in logs" :key="log.id" class="p-5 bg-white border border-gray-100 rounded-2xl shadow-sm hover:shadow-md transition-shadow">
                        <div class="flex items-start justify-between mb-4">
                            <div class="flex items-center gap-4">
                                <div class="w-12 h-12 rounded-2xl bg-blue-50 flex items-center justify-center text-blue-600">
                                    <span class="material-icons">person_search</span>
                                </div>
                                <div>
                                    <div class="font-bold text-gray-900">{{ log.viewer_name }}</div>
                                    <div class="flex items-center gap-2 mt-0.5">
                                        <div class="text-[10px] text-gray-400 font-bold uppercase tracking-widest">{{ log.unique_id }}</div>
                                        <div class="text-[10px] text-gray-400">•</div>
                                        <div class="text-[10px] text-gray-400 font-bold uppercase tracking-widest">{{ new Date(log.created_at).toLocaleString() }}</div>
                                    </div>
                                </div>
                            </div>
                            <button @click="$router.push(`/admin/donors/${log.viewer_id}/edit`)" class="px-4 py-1.5 bg-blue-50 text-blue-600 text-xs font-bold rounded-lg hover:bg-blue-100 transition-colors flex items-center gap-1.5">
                                <span class="material-icons text-sm">visibility</span>
                                View Actor
                            </button>
                        </div>
                        
                        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
                            <div class="bg-gray-50/50 p-2.5 rounded-xl border border-gray-100">
                                <div class="text-[10px] text-gray-400 font-bold uppercase mb-1">Blood Group</div>
                                <div class="text-xs font-bold text-red-600">{{ log.blood_group || 'N/A' }}</div>
                            </div>
                            <div class="bg-gray-50/50 p-2.5 rounded-xl border border-gray-100">
                                <div class="text-[10px] text-gray-400 font-bold uppercase mb-1">Phone</div>
                                <div class="text-xs font-bold text-gray-700">{{ log.phone || 'N/A' }}</div>
                            </div>
                            <div class="bg-gray-50/50 p-2.5 rounded-xl border border-gray-100 text-center">
                                <div class="text-[10px] text-gray-400 font-bold uppercase mb-1">Location</div>
                                <div class="text-xs font-bold text-gray-700">{{ log.district || 'N/A' }}</div>
                            </div>
                            <div class="bg-gray-50/50 p-2.5 rounded-xl border border-gray-100 text-center">
                                <div class="text-[10px] text-gray-400 font-bold uppercase mb-1">IP Address</div>
                                <div class="text-xs font-mono text-gray-500">{{ log.ip_address || 'Unknown' }}</div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
      </div>
    </div>

    <!-- Modals for History (optional but clean) -->
    <div v-if="isHistoryModalOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/60 backdrop-blur-sm">
         <div class="bg-white rounded-2xl shadow-2xl w-full max-w-md overflow-hidden">
            <div class="px-6 py-4 border-b border-gray-100 flex justify-between items-center">
                <h3 class="font-bold text-gray-800">Add Donation Record</h3>
                <button @click="isHistoryModalOpen = false"><span class="material-icons">close</span></button>
            </div>
            <div class="p-6 space-y-4">
                <div class="space-y-1.5">
                    <label class="text-xs font-bold">Date</label>
                    <input v-model="hForm.date" type="date" class="input-field" />
                </div>
                <div class="space-y-1.5">
                    <label class="text-xs font-bold">Type</label>
                    <select v-model="hForm.type" class="input-field">
                        <option value="Whole Blood">Whole Blood</option>
                        <option value="Platelets">Platelets</option>
                    </select>
                </div>
                <div class="space-y-1.5">
                    <label class="text-xs font-bold">Location</label>
                    <input v-model="hForm.location" type="text" class="input-field" />
                </div>
                <button @click="addHistory" class="w-full py-3 bg-[#FF3D3D] text-white font-bold rounded-xl mt-4 shadow-lg shadow-red-500/20">Add Record</button>
            </div>
         </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import api from '../../lib/axios';

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
}

interface Donation {
    id: number;
    date: string;
    location: string;
    type: string;
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

const donorId = route.params.id as string;

const activeTab = ref('profile');
const isSaving = ref(false);
const isHistoryModalOpen = ref(false);

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
    is_available: true
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
    location: '',
    type: 'Whole Blood'
});

const fetchData = async () => {
    try {
        const res = await api.get(`/admin/users/${donorId}`);
        const data = res.data;
        
        userAccount.value = {
            id: data.user.id,
            email: data.user.email,
            role: data.user.role,
            is_active: data.user.is_active
        };

        if (data.user.donor_profile) {
            profile.value = { ...data.user.donor_profile };
            if (profile.value.birthday) {
                profile.value.birthday = new Date(profile.value.birthday).toISOString().split('T')[0] as string;
            }
            if (profile.value.last_donation_date) {
                profile.value.last_donation_date = new Date(profile.value.last_donation_date).toISOString().split('T')[0] as string;
            }
        }
        
        stats.value = data.stats;
        history.value = data.history;
        
        // Also fetch view logs
        const logRes = await api.get(`/admin/users/${donorId}/view-logs`);
        logs.value = logRes.data;
        
    } catch (error) {
        console.error("Failed to fetch user details", error);
        alert("Could not load user data");
    }
};

const saveChanges = async () => {
    isSaving.value = true;
    try {
        const payload = {
            ...profile.value,
            email: userAccount.value.email,
            role: userAccount.value.role,
            is_active: userAccount.value.is_active,
            password: passwordForm.value || undefined,
            bloodGroup: profile.value.blood_group, // Fixed: profile.value.blood_group
            birthday: profile.value.birthday ? new Date(profile.value.birthday).toISOString() : null,
            last_donation_date: profile.value.last_donation_date ? new Date(profile.value.last_donation_date).toISOString() : null,
            area_village: profile.value.area_village,
            postal_code: profile.value.postal_code,
            google_map_link: profile.value.google_map_link
        };

        await api.put(`/admin/users/${donorId}`, payload);
        alert("Changes saved successfully!");
        passwordForm.value = '';
    } catch (error) {
        console.error("Failed to save changes", error);
        alert("Failed to save user data");
    } finally {
        isSaving.value = false;
    }
};

const openAddDonation = () => {
    isHistoryModalOpen.value = true;
};

const addHistory = async () => {
    try {
        await api.post(`/admin/users/${donorId}/donations`, hForm);
        fetchData();
        isHistoryModalOpen.value = false;
        hForm.location = '';
    } catch (error) {
        alert("Failed to add history");
    }
};

const deleteDonation = async (id: number) => {
    if (!confirm("Are you sure?")) return;
    try {
        await api.delete(`/admin/donations/${id}`);
        fetchData();
    } catch (error) {
        alert("Failed to delete record");
    }
};

const formatDateFull = (dateString: string) => {
    const date = new Date(dateString);
    return date.toLocaleDateString('en-US', { month: 'long', day: 'numeric', year: 'numeric' });
};

onMounted(fetchData);
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
