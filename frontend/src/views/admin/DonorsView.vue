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
            <button class="bg-[#FF3D3D] text-white px-4 py-3 rounded-xl font-bold text-sm hover:bg-red-600 transition-colors flex items-center gap-2 shadow-lg shadow-red-500/20">
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
                  <button class="w-8 h-8 rounded-lg flex items-center justify-center text-gray-400 hover:text-blue-600 hover:bg-blue-50 transition-all" title="Edit">
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
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import api from '@/lib/axios';

interface User {
    id: number;
    name: string;
    email: string;
    role: string;
    blood_group: string;
    district: string;
    is_active: boolean;
}

const users = ref<User[]>([]);
const searchQuery = ref('');

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
