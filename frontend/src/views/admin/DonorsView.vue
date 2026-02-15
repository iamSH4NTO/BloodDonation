<template>
  <div class="space-y-6">
    <!-- Action Bar -->
    <div class="flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4">
      <div class="relative w-full sm:w-96">
        <span class="absolute left-4 top-1/2 -translate-y-1/2 text-gray-400 material-icons">search</span>
        <input 
            v-model="searchQuery" 
            type="text" 
            placeholder="Search users by name, email, location..." 
            class="w-full pl-12 pr-4 py-3 bg-white border border-gray-200 rounded-xl focus:ring-2 focus:ring-[#FF3D3D]/20 focus:border-[#FF3D3D] outline-none transition-all shadow-sm"
        />
      </div>
      
       <div class="flex gap-2">
            <button class="bg-white border border-gray-200 text-gray-600 px-4 py-3 rounded-xl font-bold text-sm hover:bg-gray-50 transition-colors flex items-center gap-2 shadow-sm">
                <span class="material-icons text-sm">filter_list</span> Filter
            </button>
            <button @click="openAddModal" class="bg-[#FF3D3D] text-white px-4 py-3 rounded-xl font-bold text-sm hover:bg-red-600 transition-colors flex items-center gap-2 shadow-lg shadow-red-500/20">
                <span class="material-icons text-sm">add</span> Add User
            </button>
       </div>
    </div>

    <!-- Table Card -->
    <div class="bg-white rounded-2xl shadow-sm border border-gray-200 overflow-hidden">
        <div class="overflow-x-auto">
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-gray-50/50 border-b border-gray-200">
            <th class="px-6 py-4 text-xs font-bold text-gray-500 uppercase tracking-wider">User Profile</th>
            <th class="px-6 py-4 text-xs font-bold text-gray-500 uppercase tracking-wider">Role</th>
            <th class="px-6 py-4 text-xs font-bold text-gray-500 uppercase tracking-wider">Blood Group</th>
             <th class="px-6 py-4 text-xs font-bold text-gray-500 uppercase tracking-wider">Location</th>
             <th class="px-6 py-4 text-xs font-bold text-gray-500 uppercase tracking-wider">Last Donation</th>
            <th class="px-6 py-4 text-xs font-bold text-gray-500 uppercase tracking-wider">Status</th>
            <th class="px-6 py-4 text-right text-xs font-bold text-gray-500 uppercase tracking-wider">Actions</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-100">
          <tr v-for="user in filteredUsers" :key="user.id" class="group hover:bg-red-50/30 transition-colors duration-200">
            <td class="px-6 py-4">
              <div class="flex items-center gap-4">
                <div class="w-10 h-10 rounded-full bg-linear-to-br from-[#FF3D3D] to-[#ff6b6b] flex items-center justify-center text-white font-bold text-sm shadow-md shrink-0">
                    {{ user.name.charAt(0) }}
                </div>
                <div>
                  <div class="font-bold text-gray-900 text-sm group-hover:text-[#FF3D3D] transition-colors line-clamp-1">{{ user.name }}</div>
                  <div class="flex items-center gap-1.5">
                      <span class="text-[10px] font-bold text-gray-400 bg-gray-50 px-1.5 py-0.5 rounded border border-gray-100 uppercase tracking-tighter">{{ user.id }}</span>
                      <span class="text-[10px] text-gray-300">•</span>
                      <span class="text-[10px] text-gray-500 font-medium truncate max-w-[120px]">{{ user.email }}</span>
                  </div>
                </div>
              </div>
            </td>
            <td class="px-6 py-4">
                <span class="inline-flex items-center px-2.5 py-1 rounded-lg text-xs font-bold capitalize border" :class="user.role === 'admin' ? 'bg-purple-50 text-purple-700 border-purple-100' : 'bg-blue-50 text-blue-700 border-blue-100'">
                    {{ user.role }}
                </span>
            </td>
            <td class="px-6 py-4">
                <span class="font-bold text-gray-900 text-sm bg-gray-50 px-2 py-1 rounded-md border border-gray-100">{{ user.blood_group || 'N/A' }}</span>
            </td>
             <td class="px-6 py-4">
                <div class="text-[12px] text-gray-600 font-medium flex items-start gap-1 leading-tight">
                    <span class="material-icons text-[14px] text-gray-400 mt-0.5">place</span>
                    <span>{{ [user.area_village, user.city, user.district].filter(Boolean).join(', ') || '-' }}</span>
                </div>
            </td>
            <td class="px-6 py-4">
                <div class="text-xs font-semibold" :class="canDonate(user) ? 'text-emerald-600' : 'text-amber-600'">
                    {{ user.last_donation_date ? new Date(user.last_donation_date).toLocaleDateString() : 'Never' }}
                    <div class="text-[10px] opacity-75 font-medium">{{ getDaysSince(user.last_donation_date) }} days ago</div>
                </div>
            </td>
            <td class="px-6 py-4">
                <div class="space-y-1.5">
                    <span class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-xs font-bold border shrink-0" :class="user.is_active ? 'bg-emerald-50 text-emerald-700 border-emerald-100' : 'bg-red-50 text-red-700 border-red-100'">
                        <span class="w-1.5 h-1.5 rounded-full" :class="user.is_active ? 'bg-emerald-500' : 'bg-red-500'"></span>
                        {{ user.is_active ? 'Active' : 'Banned' }}
                    </span>
                    <div v-if="user.role === 'donor'" class="flex items-center gap-1">
                         <span class="w-1.5 h-1.5 rounded-full" :class="user.is_available ? 'bg-emerald-400' : 'bg-gray-300'"></span>
                         <span class="text-[10px] font-bold text-gray-500 uppercase">{{ user.is_available ? 'Available' : 'Unavailable' }}</span>
                    </div>
                </div>
            </td>
            <td class="px-6 py-4 text-right">
                <div class="opacity-0 group-hover:opacity-100 transition-opacity flex items-center justify-end gap-1">
                  <button @click="markDonatedToday(user)" v-if="user.role === 'donor'" class="w-8 h-8 rounded-lg flex items-center justify-center text-gray-400 hover:text-emerald-600 hover:bg-emerald-50 transition-all" title="Mark Donated Today">
                    <span class="material-icons text-sm">event_available</span>
                  </button>
                  <button @click="openEditModal(user)" class="w-8 h-8 rounded-lg flex items-center justify-center text-gray-400 hover:text-blue-600 hover:bg-blue-50 transition-all" title="View/Edit">
                    <span class="material-icons text-sm">visibility</span>
                  </button>
                  <button @click="deleteUser(user.id)" class="w-8 h-8 rounded-lg flex items-center justify-center text-gray-400 hover:text-red-600 hover:bg-red-50 transition-all" title="Delete">
                    <span class="material-icons text-sm">delete</span>
                  </button>
                </div>
            </td>
          </tr>
        </tbody>
      </table>
      </div>
       <div v-if="users.length === 0" class="p-12 text-center">
            <div class="w-16 h-16 bg-gray-50 rounded-full flex items-center justify-center mx-auto mb-4">
                <span class="material-icons text-gray-300 text-3xl">group_off</span>
            </div>
            <h3 class="text-gray-900 font-bold mb-1">No users found</h3>
            <p class="text-gray-500 text-sm">There are no registered users in the system yet.</p>
        </div>
    </div>
    <!-- Add/Edit User Modal -->
    <div v-if="isModalOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4">
        <div class="absolute inset-0 bg-black/60 backdrop-blur-sm" @click="closeModal"></div>
        <div class="bg-white rounded-2xl shadow-2xl w-full max-w-2xl relative z-10 overflow-hidden flex flex-col max-h-[90vh]">
            <div class="px-8 py-6 border-b border-gray-100 flex justify-between items-center bg-gray-50/50">
                <h3 class="text-xl font-bold text-gray-800">{{ modalMode === 'add' ? 'Add New User' : 'Edit User' }}</h3>
                <button @click="closeModal" class="text-gray-400 hover:text-gray-600 transition-colors">
                    <span class="material-icons">close</span>
                </button>
            </div>
            
            <div class="px-8 py-4 border-b border-gray-100 flex gap-6 bg-white overflow-x-auto no-scrollbar">
                <button 
                    v-for="tab in ['profile', 'history', 'logs']" 
                    :key="tab"
                    @click="activeTab = tab"
                    class="text-sm font-bold pb-2 transition-all border-b-2 capitalize whitespace-nowrap"
                    :class="activeTab === tab ? 'text-[#FF3D3D] border-[#FF3D3D]' : 'text-gray-400 border-transparent hover:text-gray-600'"
                >
                    {{ tab === 'history' ? 'Donation History' : tab === 'logs' ? 'Profile View Logs' : 'Donor Profile' }}
                </button>
            </div>
            
            <div class="p-8 overflow-y-auto custom-scrollbar grow">
                <!-- Profile Tab -->
                <form v-if="activeTab === 'profile'" @submit.prevent="saveUser" class="space-y-6">
                    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                        <!-- Profile Image (Optional Enhancement for later) -->
                        
                        <div class="space-y-2">
                            <label class="text-sm font-bold text-gray-700">Full Name</label>
                            <input v-model="form.name" type="text" required class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-500/20 focus:border-red-500 outline-none transition-all font-medium" placeholder="John Doe" />
                        </div>
                        
                        <div class="space-y-2">
                            <label class="text-sm font-bold text-gray-700">Email Address</label>
                            <input v-model="form.email" type="email" required class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-500/20 focus:border-red-500 outline-none transition-all font-medium" placeholder="john@example.com" />
                        </div>

                         <div class="space-y-2">
                            <label class="text-sm font-bold text-gray-700">Phone</label>
                            <input v-model="form.phone" type="tel" required class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-500/20 focus:border-red-500 outline-none transition-all font-medium" placeholder="+880..." />
                        </div>

                        <div class="space-y-2">
                            <label class="text-sm font-bold text-gray-700">Role</label>
                            <select v-model="form.role" class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-500/20 focus:border-red-500 outline-none transition-all font-medium appearance-none">
                                <option value="donor">Donor</option>
                                <option value="admin">Admin</option>
                            </select>
                        </div>

                         <div class="space-y-2">
                            <label class="text-sm font-bold text-gray-700">Blood Group</label>
                            <select v-model="form.bloodGroup" class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-500/20 focus:border-red-500 outline-none transition-all font-medium appearance-none">
                                <option value="">Select Group</option>
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
                        
                         <div class="space-y-2">
                            <label class="text-sm font-bold text-gray-700">Gender</label>
                            <select v-model="form.gender" class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-500/20 focus:border-red-500 outline-none transition-all font-medium appearance-none">
                                <option value="">Select Gender</option>
                                <option value="Male">Male</option>
                                <option value="Female">Female</option>
                                <option value="Other">Other</option>
                            </select>
                        </div>

                        <div class="space-y-2">
                            <label class="text-sm font-bold text-gray-700">Birthday</label>
                            <input v-model="form.birthday" type="date" class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-500/20 focus:border-red-500 outline-none transition-all font-medium" />
                        </div>

                        <div class="space-y-2">
                            <label class="text-sm font-bold text-gray-700">District</label>
                            <input v-model="form.district" type="text" class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-500/20 focus:border-red-500 outline-none transition-all font-medium" placeholder="Dhaka" />
                        </div>
                        
                        <div class="space-y-2">
                            <label class="text-sm font-bold text-gray-700">Upazila</label>
                            <input v-model="form.upazila" type="text" class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-500/20 focus:border-red-500 outline-none transition-all font-medium" placeholder="Savar" />
                        </div>

                        <div class="space-y-2">
                            <label class="text-sm font-bold text-gray-700">City</label>
                            <input v-model="form.city" type="text" class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-500/20 focus:border-red-500 outline-none transition-all font-medium" placeholder="City" />
                        </div>
                        
                        <div class="space-y-2">
                            <label class="text-sm font-bold text-gray-700">Area / Village</label>
                            <input v-model="form.areaVillage" type="text" class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-500/20 focus:border-red-500 outline-none transition-all font-medium" placeholder="Village Name" />
                        </div>
                        
                        <div class="space-y-2">
                            <label class="text-sm font-bold text-gray-700">Postal Code</label>
                            <input v-model="form.postalCode" type="text" class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-500/20 focus:border-red-500 outline-none transition-all font-medium" placeholder="1234" />
                        </div>
                        
                        <div class="space-y-2 md:col-span-2">
                            <label class="text-sm font-bold text-gray-700">Google Maps Link</label>
                            <input v-model="form.googleMapLink" type="text" class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-500/20 focus:border-red-500 outline-none transition-all font-medium" placeholder="https://maps.google.com/..." />
                        </div>

                         <div class="space-y-2">
                            <label class="text-sm font-bold text-gray-700">Last Donation Date</label>
                            <input v-model="form.lastDonationDate" type="date" class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-500/20 focus:border-red-500 outline-none transition-all font-medium" />
                        </div>

                         <div class="space-y-2">
                            <label class="text-sm font-bold text-gray-700">Donor Availability</label>
                            <select v-model="form.isAvailable" class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-500/20 focus:border-red-500 outline-none transition-all font-medium appearance-none">
                                <option :value="true">Available to Donate</option>
                                <option :value="false">Not Available</option>
                            </select>
                        </div>
                        
                         <div class="space-y-2">
                            <label class="text-sm font-bold text-gray-700">Account Status</label>
                            <select v-model="form.isActive" class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-500/20 focus:border-red-500 outline-none transition-all font-medium appearance-none">
                                <option :value="true">Active</option>
                                <option :value="false">Banned</option>
                            </select>
                        </div>
                        
                         <div class="space-y-2">
                            <label class="text-sm font-bold text-gray-700">Password</label>
                            <input v-model="form.password" type="password" :required="modalMode === 'add'" class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-500/20 focus:border-red-500 outline-none transition-all font-medium" :placeholder="modalMode === 'add' ? '******' : 'Leave blank to keep current'" />
                        </div>
                    </div>

                    <div class="pt-6 border-t border-gray-100 flex justify-end gap-3">
                        <button type="button" @click="closeModal" class="px-6 py-2.5 text-gray-600 font-bold hover:bg-gray-100 rounded-xl transition-colors">Cancel</button>
                        <button type="submit" class="px-8 py-2.5 bg-[#FF3D3D] text-white font-bold rounded-xl hover:bg-red-600 shadow-lg shadow-red-500/20 transition-all transform hover:-translate-y-0.5">
                            {{ modalMode === 'add' ? 'Create User' : 'Save Changes' }}
                        </button>
                    </div>
                </form>

                <!-- History Tab -->
                <div v-if="activeTab === 'history'" class="space-y-6">
                    <div class="bg-gray-50 rounded-xl p-4 border border-gray-100">
                        <h4 class="text-sm font-bold text-gray-700 mb-4">Add New Record</h4>
                        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                            <input v-model="historyForm.date" type="date" class="px-3 py-2 bg-white border border-gray-200 rounded-lg text-sm" />
                            <input v-model="historyForm.location" type="text" placeholder="Location" class="px-3 py-2 bg-white border border-gray-200 rounded-lg text-sm" />
                            <select v-model="historyForm.type" class="px-3 py-2 bg-white border border-gray-200 rounded-lg text-sm">
                                <option value="Whole Blood">Whole Blood</option>
                                <option value="Platelets">Platelets</option>
                            </select>
                            <button @click="addDonationHistory" class="bg-[#FF3D3D] text-white font-bold py-2 rounded-lg text-sm hover:bg-red-600">Add Record</button>
                        </div>
                    </div>

                    <div class="space-y-3">
                        <div v-for="h in donationHistory" :key="h.id" class="flex items-center justify-between p-4 bg-white border border-gray-100 rounded-xl shadow-sm">
                            <div class="flex items-center gap-3">
                                <div class="w-10 h-10 rounded-full bg-red-50 flex items-center justify-center text-red-600">
                                    <span class="material-icons text-sm">water_drop</span>
                                </div>
                                <div>
                                    <div class="font-bold text-sm text-gray-900">{{ new Date(h.date).toLocaleDateString() }}</div>
                                    <div class="text-xs text-gray-500">{{ h.location }} • {{ h.type }}</div>
                                </div>
                            </div>
                            <button @click="deleteDonationHistory(h.id)" class="text-gray-300 hover:text-red-500 transition-colors">
                                <span class="material-icons text-sm">delete</span>
                            </button>
                        </div>
                        <div v-if="donationHistory.length === 0" class="text-center py-8 text-gray-400 text-sm">No donation history found.</div>
                    </div>
                </div>

                <!-- Logs Tab -->
                <div v-if="activeTab === 'logs'" class="space-y-4">
                    <div v-for="log in viewLogs" :key="log.id" class="p-5 bg-white border border-gray-100 rounded-2xl shadow-sm hover:shadow-md transition-shadow">
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
                            <button @click="viewViewerProfile(log.viewer_id)" class="px-4 py-1.5 bg-blue-50 text-blue-600 text-xs font-bold rounded-lg hover:bg-blue-100 transition-colors flex items-center gap-1.5">
                                <span class="material-icons text-sm">open_in_new</span>
                                View Profile
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
                            <div class="bg-gray-50/50 p-2.5 rounded-xl border border-gray-100">
                                <div class="text-[10px] text-gray-400 font-bold uppercase mb-1">Location</div>
                                <div class="text-xs font-bold text-gray-700">{{ log.district || 'N/A' }}</div>
                            </div>
                            <div class="bg-gray-50/50 p-2.5 rounded-xl border border-gray-100">
                                <div class="text-[10px] text-gray-400 font-bold uppercase mb-1">IP Address</div>
                                <div class="text-xs font-mono text-gray-500">{{ log.ip_address || 'Unknown' }}</div>
                            </div>
                        </div>
                    </div>
                    <div v-if="viewLogs.length === 0" class="text-center py-12 bg-gray-50 rounded-2xl border-2 border-dashed border-gray-100">
                        <span class="material-icons text-gray-300 text-4xl mb-2">history</span>
                        <div class="text-gray-400 text-sm font-medium">No profile views recorded.</div>
                    </div>
                </div>
            </div>
        </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, reactive } from 'vue';
import api from '@/lib/axios';

interface User {
    id: string;
    name: string;
    email: string;
    role: string;
    blood_group: string;
    district: string;
    phone?: string;
    gender: string | null;
    birthday: string | null;
    upazila: string | null;
    city: string | null;
    area_village: string | null;
    postal_code: string | null;
    google_map_link: string | null;
    last_donation_date: string | null;
    is_available: boolean | null;
    is_active: boolean;
}

const users = ref<User[]>([]);
const searchQuery = ref('');
const isModalOpen = ref(false);
const modalMode = ref<'add' | 'edit'>('add');
const currentUserId = ref<string | null>(null);
const activeTab = ref('profile');

const donationHistory = ref<any[]>([]);
const viewLogs = ref<any[]>([]);
const historyForm = reactive({
    date: new Date().toISOString().split('T')[0],
    location: '',
    type: 'Whole Blood'
});

const loadDonationHistory = async (userId: string) => {
    try {
        const res = await api.get(`/admin/users/${userId}/donations`);
        donationHistory.value = res.data;
    } catch (error) {
        console.error("Failed to load history", error);
    }
};

const loadViewLogs = async (userId: string) => {
    try {
        const res = await api.get(`/admin/users/${userId}/view-logs`);
        viewLogs.value = res.data;
    } catch (error) {
        console.error("Failed to load logs", error);
    }
};

const addDonationHistory = async () => {
    if (!currentUserId.value) return;
    try {
        await api.post(`/admin/users/${currentUserId.value}/donations`, historyForm);
        loadDonationHistory(currentUserId.value);
        loadUsers(); // Update main table last donation date
        historyForm.location = '';
    } catch (error) {
        alert("Failed to add record");
    }
};

const deleteDonationHistory = async (id: number) => {
    if (!confirm("Remove this donation record?")) return;
    try {
        await api.delete(`/admin/donations/${id}`);
        if (currentUserId.value) loadDonationHistory(currentUserId.value);
        loadUsers();
    } catch (error) {
        alert("Failed to delete record");
    }
};

const viewViewerProfile = async (userId: string) => {
    // Try to find in current list
    let viewer = users.value.find(u => u.id === userId);
    
    if (!viewer) {
        // Fetch from API if not in list (e.g. filtered out)
        try {
            const res = await api.get(`/admin/users`);
            const allUsers: User[] = res.data;
            viewer = allUsers.find(u => u.id === userId);
        } catch (error) {
            console.error("Failed to fetch viewer profile", error);
        }
    }

    if (viewer) {
        openEditModal(viewer);
    } else {
        alert("Viewer profile not found.");
    }
};

const form = reactive({
    name: '',
    email: '',
    phone: '',
    role: 'donor',
    bloodGroup: '',
    gender: '',
    birthday: '',
    district: '',
    upazila: '',
    city: '',
    areaVillage: '',
    postalCode: '',
    googleMapLink: '',
    lastDonationDate: '',
    isAvailable: true,
    password: '',
    isActive: true
});

const canDonate = (user: User) => {
    if (!user.last_donation_date) return true;
    const lastDate = new Date(user.last_donation_date);
    const today = new Date();
    const diffDays = Math.ceil(Math.abs(today.getTime() - lastDate.getTime()) / (1000 * 60 * 60 * 24));
    return diffDays >= 90; // Standard 3-month gap
};

const getDaysSince = (date: string | null) => {
    if (!date) return '0';
    const lastDate = new Date(date);
    const today = new Date();
    const diffDays = Math.floor(Math.abs(today.getTime() - lastDate.getTime()) / (1000 * 60 * 60 * 24));
    return diffDays;
};

const markDonatedToday = async (user: User) => {
    if (!confirm(`Mark ${user.name} as donated today? This will update their last donation date and set availability based on the 3-month rule.`)) return;
    
    try {
        const today = new Date().toISOString();
        await api.put(`/admin/users/${user.id}`, {
            last_donation_date: today,
            is_available: false // Usually unavailable after donation
        });
        
        // Update local state
        const index = users.value.findIndex(u => u.id === user.id);
        if (index !== -1) {
            users.value[index] = {
                ...users.value[index],
                last_donation_date: today,
                is_available: false
            } as User;
        }
    } catch (error) {
        console.error("Failed to mark as donated", error);
        alert("Failed to update status");
    }
};

const filteredUsers = computed(() => {
    if (!searchQuery.value) return users.value;
    const query = searchQuery.value.toLowerCase();
    return users.value.filter(user => 
        user.name.toLowerCase().includes(query) || 
        user.email.toLowerCase().includes(query) ||
        (user.district && user.district.toLowerCase().includes(query)) ||
        (user.upazila && user.upazila.toLowerCase().includes(query)) ||
        (user.city && user.city.toLowerCase().includes(query)) ||
        (user.area_village && user.area_village.toLowerCase().includes(query)) ||
        (user.blood_group && user.blood_group.toLowerCase().includes(query))
    );
});

const loadUsers = async () => {
    try {
        const res = await api.get('/admin/users');
        users.value = res.data;
    } catch (error) {
        console.error("Failed to load users", error);
    }
};

const openAddModal = () => {
    modalMode.value = 'add';
    currentUserId.value = null;
    activeTab.value = 'profile';
    resetForm();
    isModalOpen.value = true;
};

const openEditModal = (user: User) => {
    modalMode.value = 'edit';
    currentUserId.value = user.id;
    activeTab.value = 'profile';
    loadDonationHistory(user.id);
    loadViewLogs(user.id);
    form.name = user.name;
    form.email = user.email;
    form.role = user.role;
    form.bloodGroup = user.blood_group;
    form.gender = user.gender || '';
    form.birthday = user.birthday ? (user.birthday.split('T')[0] ?? '') : '';
    form.district = user.district || '';
    form.upazila = user.upazila || '';
    form.city = user.city || '';
    form.areaVillage = user.area_village || '';
    form.postalCode = user.postal_code || '';
    form.googleMapLink = user.google_map_link || '';
    form.lastDonationDate = user.last_donation_date ? (user.last_donation_date.split('T')[0] ?? '') : '';
    form.isAvailable = user.is_available !== null ? user.is_available : true;
    form.phone = user.phone || '';
    form.isActive = user.is_active;
    form.password = ''; 
    isModalOpen.value = true;
};

const closeModal = () => {
    isModalOpen.value = false;
    resetForm();
};

const resetForm = () => {
    form.name = '';
    form.email = '';
    form.phone = '';
    form.role = 'donor';
    form.bloodGroup = '';
    form.gender = '';
    form.birthday = '';
    form.district = '';
    form.upazila = '';
    form.city = '';
    form.areaVillage = '';
    form.postalCode = '';
    form.googleMapLink = '';
    form.lastDonationDate = '';
    form.isAvailable = true;
    form.password = '';
    form.isActive = true;
};

const saveUser = async () => {
    try {
        const payload = {
            ...form,
            is_active: form.isActive,
            is_available: form.isAvailable,
            birthday: form.birthday ? new Date(form.birthday).toISOString() : null,
            last_donation_date: form.lastDonationDate ? new Date(form.lastDonationDate).toISOString() : null,
            area_village: form.areaVillage,
            postal_code: form.postalCode,
            google_map_link: form.googleMapLink
        };

        if (modalMode.value === 'add') {
             const res = await api.post('/admin/users', payload);
             users.value.push(res.data.user); // Assuming backend returns the created user
             // Reload to get formatted fields if needed
             loadUsers();
        } else {
            await api.put(`/admin/users/${currentUserId.value}`, payload);
            // Update local state
            const index = users.value.findIndex(u => u.id === currentUserId.value);
            if (index !== -1 && currentUserId.value !== null) {
                users.value[index] = { 
                    ...users.value[index], 
                    ...form, 
                    id: currentUserId.value,
                    is_active: form.isActive, 
                    blood_group: form.bloodGroup,
                    role: form.role,
                    gender: form.gender || null, // Ensure fallback to null if empty
                    birthday: form.birthday ? new Date(form.birthday).toISOString() : null,
                    upazila: form.upazila || null,
                    city: form.city || null,
                    area_village: form.areaVillage || null,
                    postal_code: form.postalCode || null,
                    google_map_link: form.googleMapLink || null,
                    last_donation_date: form.lastDonationDate ? new Date(form.lastDonationDate).toISOString() : null,
                    is_available: form.isAvailable
                } as User;
            }
        }
        closeModal();
    } catch (error) {
        console.error("Failed to save user", error);
        alert("Failed to save user. Please check inputs.");
    }
};

const deleteUser = async (id: string) => {
    if (!confirm('Are you sure you want to delete this user? This cannot be undone.')) return;
    try {
        await api.delete(`/admin/users/${id}`);
        users.value = users.value.filter(u => u.id !== id);
    } catch (error) {
        alert('Failed to delete user');
    }
};

onMounted(loadUsers);
</script>
