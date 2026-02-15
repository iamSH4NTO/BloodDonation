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

    <!-- Table Card (Desktop) -->
    <div class="hidden lg:block bg-white rounded-2xl shadow-sm border border-gray-200 overflow-hidden">
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
                  <button @click="$router.push(`/admin/donors/${user.id}/edit`)" class="w-8 h-8 rounded-lg flex items-center justify-center text-gray-400 hover:text-blue-600 hover:bg-blue-50 transition-all" title="View/Edit">
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
    </div>

    <!-- Mobile Card View -->
    <div class="lg:hidden space-y-4">
        <div v-for="user in filteredUsers" :key="user.id" class="bg-white rounded-2xl border border-gray-100 p-5 shadow-sm space-y-4">
            <div class="flex items-center justify-between">
                <div class="flex items-center gap-3">
                    <div class="w-10 h-10 rounded-xl bg-linear-to-br from-[#FF3D3D] to-[#ff6b6b] flex items-center justify-center text-white font-bold shadow-md">
                        {{ user.name.charAt(0) }}
                    </div>
                    <div>
                        <h4 class="font-bold text-gray-900 text-sm">{{ user.name }}</h4>
                        <p class="text-[10px] font-bold text-gray-400 tracking-tight">{{ user.id }}</p>
                    </div>
                </div>
                <span class="font-black text-[#FF3D3D] bg-red-50 px-2.5 py-1 rounded-lg border border-red-100 text-xs">{{ user.blood_group || 'N/A' }}</span>
            </div>

            <div class="grid grid-cols-2 gap-4 py-3 border-y border-gray-50">
                <div>
                    <p class="text-[10px] font-bold text-gray-400 uppercase tracking-wider mb-1">Status</p>
                    <div class="flex items-center gap-1.5">
                        <span class="w-1.5 h-1.5 rounded-full" :class="user.is_active ? 'bg-emerald-500' : 'bg-red-500'"></span>
                        <span class="text-xs font-bold" :class="user.is_active ? 'text-emerald-700' : 'text-red-700'">{{ user.is_active ? 'Active' : 'Banned' }}</span>
                    </div>
                </div>
                <div>
                    <p class="text-[10px] font-bold text-gray-400 uppercase tracking-wider mb-1">Role</p>
                    <span class="text-xs font-bold text-gray-700 capitalize">{{ user.role }}</span>
                </div>
            </div>

            <div class="space-y-2">
                <div class="flex items-center gap-2 text-xs text-gray-500 font-medium font-sans">
                    <span class="material-icons text-sm text-gray-400">mail</span>
                    <span class="truncate">{{ user.email }}</span>
                </div>
                <div class="flex items-center gap-2 text-xs text-gray-500 font-medium">
                    <span class="material-icons text-sm text-gray-400">place</span>
                    <span class="truncate">{{ [user.city, user.district].filter(Boolean).join(', ') || '-' }}</span>
                </div>
                <div class="flex items-center gap-2 text-xs font-medium" :class="canDonate(user) ? 'text-emerald-600' : 'text-amber-600'">
                    <span class="material-icons text-sm opacity-60">history</span>
                    <span>Last: <span class="font-bold">{{ user.last_donation_date ? new Date(user.last_donation_date).toLocaleDateString() : 'Never' }}</span></span>
                </div>
            </div>

            <div class="flex gap-2 pt-2">
                <button 
                    @click="$router.push(`/admin/donors/${user.id}/edit`)"
                    class="flex-1 bg-white border border-gray-200 text-gray-700 font-bold py-2.5 rounded-xl text-xs hover:bg-gray-50 transition-colors flex items-center justify-center gap-2"
                >
                    <span class="material-icons text-sm">visibility</span> View/Edit
                </button>
                <button 
                    @click="deleteUser(user.id)"
                    class="px-4 bg-gray-50 text-gray-400 hover:text-red-600 hover:bg-red-50 rounded-xl transition-colors flex items-center justify-center border border-transparent"
                >
                    <span class="material-icons text-sm">delete</span>
                </button>
            </div>
        </div>
    </div>
    
    <div v-if="filteredUsers.length === 0" class="bg-white rounded-2xl border border-gray-100 p-12 text-center">
        <div class="w-16 h-16 bg-gray-50 rounded-full flex items-center justify-center mx-auto mb-4">
            <span class="material-icons text-gray-300 text-3xl">group_off</span>
        </div>
        <h3 class="text-gray-900 font-bold mb-1">No users found</h3>
        <p class="text-gray-500 text-sm">There are no registered users in the system yet.</p>
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

const form = reactive({
    name: '',
    email: '',
    phone: '',
    role: 'donor' as 'donor' | 'admin',
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
             users.value.push(res.data.user);
             loadUsers();
        }
        closeModal();
    } catch (error) {
        console.error("Failed to save user", error);
        alert("Failed to save user. Please check inputs.");
    }
};

const closeModal = () => {
    isModalOpen.value = false;
};

const openEditModal = (user: User) => {
    // Legacy: Navigate to dedicated edit page
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
