<template>
  <div class="space-y-6">
    <!-- Action Bar -->
    <div class="flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4">
      <div class="relative w-full sm:w-96">
        <span class="absolute left-4 top-1/2 -translate-y-1/2 text-gray-400 material-icons">search</span>
        <input 
            v-model="searchQuery" 
            type="text" 
            placeholder="Search users by name, email..." 
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
                  <div class="font-bold text-gray-900 text-sm group-hover:text-[#FF3D3D] transition-colors">{{ user.name }}</div>
                  <div class="text-xs text-gray-500 font-medium">{{ user.email }}</div>
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
                <div class="text-sm text-gray-600 font-medium flex items-center gap-1">
                    <span v-if="user.district" class="material-icons text-[14px] text-gray-400">place</span>
                    {{ user.district || '-' }}
                </div>
            </td>
            <td class="px-6 py-4">
                <span class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-xs font-bold border" :class="user.is_active ? 'bg-emerald-50 text-emerald-700 border-emerald-100' : 'bg-red-50 text-red-700 border-red-100'">
                    <span class="w-1.5 h-1.5 rounded-full" :class="user.is_active ? 'bg-emerald-500' : 'bg-red-500'"></span>
                    {{ user.is_active ? 'Active' : 'Banned' }}
                </span>
            </td>
            <td class="px-6 py-4 text-right">
                <div class="opacity-0 group-hover:opacity-100 transition-opacity flex items-center justify-end gap-2">
                  <button @click="openEditModal(user)" class="w-8 h-8 rounded-lg flex items-center justify-center text-gray-400 hover:text-blue-600 hover:bg-blue-50 transition-all" title="Edit">
                    <span class="material-icons text-sm">edit</span>
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
            
            <div class="p-8 overflow-y-auto custom-scrollbar">
                <form @submit.prevent="saveUser" class="space-y-6">
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
            </div>
        </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, reactive } from 'vue';
import api from '@/lib/axios';

interface User {
    id: number;
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
const currentUserId = ref<number | null>(null);

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

const filteredUsers = computed(() => {
    if (!searchQuery.value) return users.value;
    const query = searchQuery.value.toLowerCase();
    return users.value.filter(user => 
        user.name.toLowerCase().includes(query) || 
        user.email.toLowerCase().includes(query)
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
    resetForm();
    isModalOpen.value = true;
};

const openEditModal = (user: User) => {
    modalMode.value = 'edit';
    currentUserId.value = user.id;
    form.name = user.name;
    form.email = user.email;
    form.role = user.role;
    form.bloodGroup = user.blood_group;
    form.gender = user.gender || '';
    form.birthday = user.birthday ? user.birthday.split('T')[0] : '';
    form.district = user.district || '';
    form.upazila = user.upazila || '';
    form.city = user.city || '';
    form.areaVillage = user.area_village || '';
    form.postalCode = user.postal_code || '';
    form.googleMapLink = user.google_map_link || '';
    form.lastDonationDate = user.last_donation_date ? user.last_donation_date.split('T')[0] : '';
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
                };
            }
        }
        closeModal();
    } catch (error) {
        console.error("Failed to save user", error);
        alert("Failed to save user. Please check inputs.");
    }
};

const deleteUser = async (id: number) => {
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
