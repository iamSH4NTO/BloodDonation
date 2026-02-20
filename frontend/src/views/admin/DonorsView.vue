<template>
  <div>
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
              <button @click="showFilters = !showFilters" class="bg-white border text-gray-600 px-4 py-3 rounded-xl font-bold text-sm hover:bg-gray-50 transition-colors flex items-center gap-2 shadow-sm" :class="showFilters ? 'border-[#FF3D3D] text-[#FF3D3D]' : 'border-gray-200'">
                  <span class="material-icons text-sm">filter_list</span> Filters
                  <span v-if="activeFilterCount > 0" class="bg-[#FF3D3D] text-white text-[10px] w-5 h-5 rounded-full flex items-center justify-center">{{ activeFilterCount }}</span>
              </button>
              <button @click="openAddModal" class="bg-[#FF3D3D] text-white px-4 py-3 rounded-xl font-bold text-sm hover:bg-red-600 transition-colors flex items-center gap-2 shadow-lg shadow-red-500/20">
                  <span class="material-icons text-sm">add</span> Add User
              </button>
          </div>
      </div>

      <!-- Expandable Filter Panel -->
      <div v-show="showFilters" class="bg-white rounded-2xl p-5 shadow-sm border border-gray-200 animate-fade-in-up">
          <div class="flex items-center justify-between mb-4 border-b border-gray-100 pb-3">
              <h3 class="font-bold text-gray-900 flex items-center gap-2"><span class="material-icons text-gray-400">tune</span> Advanced Filters</h3>
              <button @click="clearFilters" class="text-sm font-bold text-red-600 hover:text-red-700 hover:underline">Clear All</button>
          </div>
          
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
              <!-- Blood Group (Multi-check) -->
              <div class="space-y-3">
                  <label class="text-xs font-bold text-gray-500 uppercase tracking-wider">Blood Group</label>
                  <div class="grid grid-cols-2 gap-2">
                      <label v-for="bg in ['A+', 'A-', 'B+', 'B-', 'AB+', 'AB-', 'O+', 'O-']" :key="bg" class="flex items-center gap-2 cursor-pointer p-2 rounded-lg border border-gray-100 hover:bg-gray-50 transition-colors">
                          <input type="checkbox" :value="bg" v-model="filters.bloodGroups" class="text-[#FF3D3D] rounded border-gray-300 focus:ring-[#FF3D3D] bg-white" />
                          <span class="text-sm font-bold text-gray-700">{{ bg }}</span>
                      </label>
                  </div>
              </div>

              <!-- Roles (Multi-check) -->
              <div class="space-y-3">
                  <label class="text-xs font-bold text-gray-500 uppercase tracking-wider">Role</label>
                  <div class="space-y-2">
                       <label class="flex items-center gap-2 cursor-pointer p-2 rounded-lg border border-gray-100 hover:bg-gray-50 transition-colors">
                          <input type="checkbox" value="donor" v-model="filters.roles" class="text-[#FF3D3D] rounded border-gray-300 focus:ring-[#FF3D3D]" />
                          <span class="text-sm font-bold text-gray-700">Donor</span>
                      </label>
                      <label class="flex items-center gap-2 cursor-pointer p-2 rounded-lg border border-gray-100 hover:bg-gray-50 transition-colors">
                          <input type="checkbox" value="admin" v-model="filters.roles" class="text-[#FF3D3D] rounded border-gray-300 focus:ring-[#FF3D3D]" />
                          <span class="text-sm font-bold text-gray-700">Admin</span>
                      </label>
                  </div>
              </div>
              
              <!-- Status & Avail -->
               <div class="space-y-4">
                  <div class="space-y-2">
                      <label class="text-xs font-bold text-gray-500 uppercase tracking-wider">Account Status</label>
                      <select v-model="filters.is_active" class="w-full p-2.5 bg-gray-50 border border-gray-200 rounded-xl focus:ring-[#FF3D3D] outline-none text-sm font-bold text-gray-700">
                          <option value="">All Statuses</option>
                          <option value="true">Active</option>
                          <option value="false">Banned</option>
                      </select>
                  </div>
                   <div class="space-y-2">
                      <label class="text-xs font-bold text-gray-500 uppercase tracking-wider">Availability</label>
                      <select v-model="filters.is_available" class="w-full p-2.5 bg-gray-50 border border-gray-200 rounded-xl focus:ring-[#FF3D3D] outline-none text-sm font-bold text-gray-700">
                          <option value="">All</option>
                          <option value="true">Available</option>
                          <option value="false">Unavailable</option>
                      </select>
                  </div>
              </div>

              <!-- Verification -->
              <div class="space-y-3">
                  <label class="text-xs font-bold text-gray-500 uppercase tracking-wider">Verification Status</label>
                   <div class="space-y-2">
                       <label class="flex items-center gap-2 cursor-pointer p-3 rounded-xl border transition-colors" :class="filters.verified_status === 'verified' ? 'bg-red-50 border-red-200' : 'border-gray-100 hover:bg-gray-50'">
                          <input type="radio" value="verified" v-model="filters.verified_status" class="text-[#FF3D3D] focus:ring-[#FF3D3D]" />
                          <span class="text-sm font-bold text-gray-700 flex items-center gap-1.5"><span class="material-icons text-emerald-500 text-[16px]">verified</span> Verified</span>
                      </label>
                      <label class="flex items-center gap-2 cursor-pointer p-3 rounded-xl border transition-colors" :class="filters.verified_status === 'unverified' ? 'bg-red-50 border-red-200' : 'border-gray-100 hover:bg-gray-50'">
                          <input type="radio" value="unverified" v-model="filters.verified_status" class="text-[#FF3D3D] focus:ring-[#FF3D3D]" />
                          <span class="text-sm font-bold text-gray-700 flex items-center gap-1.5"><span class="material-icons text-gray-400 text-[16px]">new_releases</span> Unverified</span>
                      </label>
                      <label class="flex items-center gap-2 cursor-pointer p-3 rounded-xl border transition-colors" :class="filters.verified_status === '' ? 'bg-red-50 border-red-200' : 'border-gray-100 hover:bg-gray-50'">
                          <input type="radio" value="" v-model="filters.verified_status" class="text-[#FF3D3D] focus:ring-[#FF3D3D]" />
                          <span class="text-sm font-bold text-gray-700">All Donors</span>
                      </label>
                  </div>
              </div>
          </div>
      </div>
  
      <!-- Table Card (Desktop) -->
      <div class="hidden lg:block bg-white rounded-2xl shadow-sm border border-gray-200">
          <div class="overflow-x-visible overflow-y-visible min-h-[350px]">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="bg-gray-50/50 border-b border-gray-200">
              <th class="px-6 py-4 text-xs font-bold text-gray-500 uppercase tracking-wider">User Profile</th>
              <th class="px-6 py-4 text-xs font-bold text-gray-500 uppercase tracking-wider">Phone</th>
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
                    <div class="font-bold text-gray-900 text-sm group-hover:text-[#FF3D3D] transition-colors line-clamp-1 flex items-center gap-1">
                        {{ user.name }}
                        <span v-if="user.is_admin_verified" class="material-icons text-[#FF3D3D] text-[14px]" title="Admin Verified">verified</span>
                    </div>
                    <div class="flex items-center gap-1.5">
                        <span class="text-[10px] font-bold text-gray-400 bg-gray-50 px-1.5 py-0.5 rounded border border-gray-100 uppercase tracking-tighter">{{ user.id }}</span>
                        <span class="text-[10px] text-gray-300">•</span>
                        <span class="text-[10px] text-gray-500 font-medium truncate max-w-[120px]">{{ user.email }}</span>
                    </div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                  <span class="text-sm font-bold text-gray-700 font-mono">{{ user.phone || 'N/A' }}</span>
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
                      <div v-if="user.role === 'donor'" class="flex items-center gap-1 mt-1">
                          <span class="material-icons text-[12px]" :class="user.is_admin_verified ? 'text-[#FF3D3D]' : 'text-gray-400'">{{ user.is_admin_verified ? 'verified' : 'new_releases' }}</span>
                          <span class="text-[10px] font-bold uppercase" :class="user.is_admin_verified ? 'text-[#FF3D3D]' : 'text-gray-500'">{{ user.is_admin_verified ? 'Verified' : 'Unverified' }}</span>
                      </div>
                  </div>
              </td>
              <td class="px-6 py-4 text-right">
                  <div class="relative inline-block text-left dropdown-container">
                    <button @click="toggleDropdown(user.id)" class="w-8 h-8 rounded-lg flex items-center justify-center hover:bg-gray-100 transition-colors focus:outline-none">
                      <span class="material-icons text-gray-500">more_vert</span>
                    </button>
                    
                    <div v-show="openDropdownId === user.id" class="absolute right-0 mt-1 w-48 bg-white rounded-xl shadow-lg border border-gray-100 z-50 py-1 origin-top-right animate-fade-in-up">
                      <button @click="toggleAdminVerification(user); closeDropdown()" v-if="user.role === 'donor'" class="w-full text-left px-4 py-2 text-sm font-medium transition-colors flex items-center gap-2" :class="user.is_admin_verified ? 'text-amber-600 hover:bg-amber-50' : 'text-emerald-600 hover:bg-emerald-50'">
                        <span class="material-icons text-sm">{{ user.is_admin_verified ? 'remove_circle_outline' : 'verified' }}</span> 
                        {{ user.is_admin_verified ? 'Remove Verification' : 'Verify Donor' }}
                      </button>
                      
                      <button @click="markDonatedToday(user); closeDropdown()" v-if="user.role === 'donor'" class="w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-50 hover:text-emerald-600 font-medium transition-colors flex items-center gap-2">
                        <span class="material-icons text-sm">event_available</span> Mark Donated
                      </button>
                      
                      <div v-if="user.role === 'donor'" class="h-px bg-gray-100 my-1"></div>
                      
                      <button @click="$router.push(`/admin/donors/${user.id}`); closeDropdown()" class="w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-50 hover:text-blue-600 font-medium transition-colors flex items-center gap-2">
                        <span class="material-icons text-sm">visibility</span> View Details
                      </button>
                      
                      <button @click="$router.push(`/admin/donors/${user.id}/edit`); closeDropdown()" class="w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-50 hover:text-blue-600 font-medium transition-colors flex items-center gap-2">
                        <span class="material-icons text-sm">edit</span> Edit Profile
                      </button>
                      
                      <div class="h-px bg-gray-100 my-1"></div>
                      
                      <button @click="deleteUser(user.id); closeDropdown()" class="w-full text-left px-4 py-2 text-sm text-red-600 hover:bg-red-50 font-medium transition-colors flex items-center gap-2">
                        <span class="material-icons text-sm">delete</span> Delete User
                      </button>
                    </div>
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
                  <div v-if="user.role === 'donor'">
                      <p class="text-[10px] font-bold text-gray-400 uppercase tracking-wider mb-1">Verify</p>
                      <div class="flex items-center gap-1.5">
                          <span class="material-icons text-[12px]" :class="user.is_admin_verified ? 'text-[#FF3D3D]' : 'text-gray-400'">{{ user.is_admin_verified ? 'verified' : 'new_releases' }}</span>
                          <span class="text-xs font-bold" :class="user.is_admin_verified ? 'text-[#FF3D3D]' : 'text-gray-500'">{{ user.is_admin_verified ? 'Verified' : 'Unverified' }}</span>
                      </div>
                  </div>
              </div>
  
              <div class="space-y-2">
                  <div class="flex items-center gap-2 text-xs text-gray-500 font-medium font-sans">
                      <span class="material-icons text-sm text-gray-400">mail</span>
                      <span class="truncate">{{ user.email }}</span>
                  </div>
                  <div class="flex items-center gap-2 text-xs text-gray-500 font-medium font-sans">
                      <span class="material-icons text-sm text-gray-400">phone</span>
                      <span class="truncate font-mono">{{ user.phone || 'N/A' }}</span>
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
  
              <div class="flex gap-2 pt-2 relative dropdown-container">
                  <button 
                      @click="$router.push(`/admin/donors/${user.id}`)"
                      class="flex-1 bg-white border border-gray-200 text-gray-700 py-2 rounded-xl flex items-center justify-center gap-2 hover:bg-gray-50 transition-colors shadow-sm"
                  >
                      <span class="material-icons text-sm">visibility</span>
                      <span class="text-xs font-bold">View</span>
                  </button>
                  <button 
                      @click="toggleDropdown(user.id)"
                      class="flex-1 bg-gray-900 text-white py-2 rounded-xl flex items-center justify-center gap-2 hover:bg-gray-800 transition-colors shadow-sm"
                  >
                      <span class="text-xs font-bold">Actions</span>
                      <span class="material-icons text-sm">{{ openDropdownId === user.id ? 'expand_less' : 'expand_more' }}</span>
                  </button>

                  <div v-show="openDropdownId === user.id" class="absolute bottom-12 right-0 w-48 bg-white rounded-xl shadow-xl border border-gray-100 z-50 py-1 origin-bottom-right animate-fade-in-up">
                      <button @click="toggleAdminVerification(user); closeDropdown()" v-if="user.role === 'donor'" class="w-full text-left px-4 py-2 text-sm font-medium transition-colors flex items-center gap-2" :class="user.is_admin_verified ? 'text-amber-600 hover:bg-amber-50' : 'text-emerald-600 hover:bg-emerald-50'">
                        <span class="material-icons text-sm">{{ user.is_admin_verified ? 'remove_circle_outline' : 'verified' }}</span> 
                        {{ user.is_admin_verified ? 'Remove Verification' : 'Verify Donor' }}
                      </button>
                      <button @click="markDonatedToday(user); closeDropdown()" v-if="user.role === 'donor'" class="w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-50 hover:text-emerald-600 font-medium transition-colors flex items-center gap-2">
                        <span class="material-icons text-sm">event_available</span> Mark Donated
                      </button>
                      <div v-if="user.role === 'donor'" class="h-px bg-gray-100 my-1"></div>
                      <button @click="$router.push(`/admin/donors/${user.id}/edit`); closeDropdown()" class="w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-50 hover:text-blue-600 font-medium transition-colors flex items-center gap-2">
                        <span class="material-icons text-sm">edit</span> Edit Profile
                      </button>
                      <button @click="deleteUser(user.id); closeDropdown()" class="w-full text-left px-4 py-2 text-sm text-red-600 hover:bg-red-50 font-medium transition-colors flex items-center gap-2">
                        <span class="material-icons text-sm">delete</span> Delete User
                      </button>
                  </div>
              </div>
          </div>
      </div>
      
        <!-- Loading State -->
        <div v-if="loading" class="text-center py-12">
            <div class="inline-block animate-spin rounded-full h-8 w-8 border-4 border-gray-200 border-t-[#FF3D3D]"></div>
            <div class="text-gray-400 text-sm mt-2">Loading donors...</div>
        </div>

        <!-- Empty State -->
        <div v-else-if="filteredUsers.length === 0" class="bg-white rounded-2xl border border-gray-100 p-12 text-center">
            <div class="w-16 h-16 bg-gray-50 rounded-full flex items-center justify-center mx-auto mb-4">
                <span class="material-icons text-gray-300 text-3xl">group_off</span>
            </div>
            <h3 class="text-gray-900 font-bold mb-1">No users found</h3>
            <p class="text-gray-500 text-sm">No users match your search criteria or are registered yet.</p>
        </div>
      </div>

      <!-- Pagination -->
      <div v-if="totalPages > 1" class="flex justify-center items-center gap-2 mt-6 pb-6">
          <button @click="changePage(currentPage - 1)" :disabled="currentPage === 1" class="px-3 py-2 rounded-lg bg-white border border-gray-200 text-sm font-bold text-gray-700 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed">
              <span class="material-icons text-sm">chevron_left</span>
          </button>
          
          <div class="flex items-center gap-1">
              <button v-for="page in visiblePages" :key="page" @click="changePage(page)" :class="currentPage === page ? 'bg-[#FF3D3D] text-white shadow-lg shadow-red-500/20' : 'bg-white text-gray-700 hover:bg-gray-50'" class="w-8 h-8 sm:w-10 sm:h-10 rounded-lg border border-gray-200 text-xs sm:text-sm font-bold transition-colors">
                  {{ page }}
              </button>
          </div>

          <button @click="changePage(currentPage + 1)" :disabled="currentPage === totalPages" class="px-3 py-2 rounded-lg bg-white border border-gray-200 text-sm font-bold text-gray-700 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed">
              <span class="material-icons text-sm">chevron_right</span>
          </button>
      </div>
  
    <!-- Add/Edit User Modal -->
    <div v-if="isModalOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4 sm:p-6">
        <div class="absolute inset-0 bg-gray-900/60 backdrop-blur-xs transition-opacity" @click="closeModal"></div>
        <div class="relative bg-white rounded-2xl shadow-xl w-full max-w-4xl max-h-[90vh] flex flex-col overflow-hidden transform transition-all">
            
            <!-- Header -->
            <div class="px-6 py-4 border-b border-gray-100 flex items-center justify-between shrink-0 bg-white z-10">
                <h3 class="text-lg font-bold text-gray-900 flex items-center gap-2">
                    <span class="w-8 h-8 rounded-lg bg-red-50 text-[#FF3D3D] flex items-center justify-center">
                        <span class="material-icons text-sm">{{ modalMode === 'add' ? 'person_add' : 'edit' }}</span>
                    </span>
                    {{ modalMode === 'add' ? 'Add New User' : 'Edit User' }}
                </h3>
                <button @click="closeModal" class="w-8 h-8 rounded-lg hover:bg-gray-50 flex items-center justify-center text-gray-400 hover:text-gray-600 transition-colors">
                    <span class="material-icons text-lg">close</span>
                </button>
            </div>
  
            <!-- Content -->
            <div class="flex-1 overflow-y-auto custom-scrollbar">
                <div class="p-6">
                    <!-- Tabs -->
                    <div class="flex gap-4 border-b border-gray-100 mb-6">
                        <button 
                            @click="activeTab = 'profile'"
                            class="pb-3 text-sm font-bold border-b-2 transition-colors flex items-center gap-2"
                            :class="activeTab === 'profile' ? 'border-[#FF3D3D] text-[#FF3D3D]' : 'border-transparent text-gray-500 hover:text-gray-700'"
                        >
                            <span class="material-icons text-sm">person</span> Profile Info
                        </button>
                         <button 
                            @click="activeTab = 'address'"
                            class="pb-3 text-sm font-bold border-b-2 transition-colors flex items-center gap-2"
                            :class="activeTab === 'address' ? 'border-[#FF3D3D] text-[#FF3D3D]' : 'border-transparent text-gray-500 hover:text-gray-700'"
                        >
                            <span class="material-icons text-sm">place</span> Address Details
                        </button>
                         <button 
                            @click="activeTab = 'status'"
                            class="pb-3 text-sm font-bold border-b-2 transition-colors flex items-center gap-2"
                            :class="activeTab === 'status' ? 'border-[#FF3D3D] text-[#FF3D3D]' : 'border-transparent text-gray-500 hover:text-gray-700'"
                        >
                             <span class="material-icons text-sm">verified_user</span> Status & Role
                        </button>
                    </div>
  
                    <form @submit.prevent="saveUser" class="space-y-6">
                        
                        <!-- Profile Tab -->
                        <div v-show="activeTab === 'profile'" class="grid grid-cols-1 sm:grid-cols-2 gap-5">
                             <div class="space-y-1.5">
                                <label class="text-xs font-bold text-gray-500 uppercase tracking-wider">Full Name *</label>
                                <input v-model="form.name" type="text" required class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-100 focus:border-[#FF3D3D] outline-none transition-all font-medium text-gray-900 placeholder:text-gray-400" />
                            </div>
  
                             <div class="space-y-1.5">
                                <label class="text-xs font-bold text-gray-500 uppercase tracking-wider">Email Address *</label>
                                <input v-model="form.email" type="email" required class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-100 focus:border-[#FF3D3D] outline-none transition-all font-medium text-gray-900 placeholder:text-gray-400" />
                            </div>
  
                             <div class="space-y-1.5">
                                <label class="text-xs font-bold text-gray-500 uppercase tracking-wider">Phone Number *</label>
                                <input v-model="form.phone" type="text" required class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-100 focus:border-[#FF3D3D] outline-none transition-all font-mono text-gray-900 placeholder:text-gray-400" />
                            </div>
  
                             <div v-if="modalMode === 'add'" class="space-y-1.5">
                                <label class="text-xs font-bold text-gray-500 uppercase tracking-wider">Password *</label>
                                <input v-model="form.password" type="password" required minlength="6" class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-100 focus:border-[#FF3D3D] outline-none transition-all font-mono text-gray-900 placeholder:text-gray-400" />
                                <p class="text-[10px] text-gray-400">Min. 6 characters</p>
                            </div>
  
                            <div class="space-y-1.5">
                                <label class="text-xs font-bold text-gray-500 uppercase tracking-wider">Blood Group *</label>
                                <select v-model="form.bloodGroup" required class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-100 focus:border-[#FF3D3D] outline-none transition-all font-medium text-gray-900">
                                    <option value="" disabled>Select Blood Group</option>
                                    <option v-for="bg in ['A+', 'A-', 'B+', 'B-', 'AB+', 'AB-', 'O+', 'O-']" :key="bg" :value="bg">{{ bg }}</option>
                                </select>
                            </div>
  
                            <div class="space-y-1.5">
                                <label class="text-xs font-bold text-gray-500 uppercase tracking-wider">Gender</label>
                                <select v-model="form.gender" class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-100 focus:border-[#FF3D3D] outline-none transition-all font-medium text-gray-900">
                                    <option value="">Select Gender</option>
                                    <option value="Male">Male</option>
                                    <option value="Female">Female</option>
                                    <option value="Other">Other</option>
                                </select>
                            </div>
  
                             <div class="space-y-1.5">
                                <label class="text-xs font-bold text-gray-500 uppercase tracking-wider">Birthday</label>
                                <input v-model="form.birthday" type="date" class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-100 focus:border-[#FF3D3D] outline-none transition-all font-medium text-gray-900" />
                            </div>
                        </div>
  
                         <!-- Address Tab -->
                        <div v-show="activeTab === 'address'" class="grid grid-cols-1 sm:grid-cols-2 gap-5">
                            <div class="space-y-1.5">
                                <label class="text-xs font-bold text-gray-500 uppercase tracking-wider">District</label>
                                <input v-model="form.district" type="text" class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-100 focus:border-[#FF3D3D] outline-none transition-all font-medium text-gray-900" />
                            </div>
                             <div class="space-y-1.5">
                                <label class="text-xs font-bold text-gray-500 uppercase tracking-wider">Upazila</label>
                                <input v-model="form.upazila" type="text" class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-100 focus:border-[#FF3D3D] outline-none transition-all font-medium text-gray-900" />
                            </div>
                            <div class="space-y-1.5">
                                <label class="text-xs font-bold text-gray-500 uppercase tracking-wider">City</label>
                                <input v-model="form.city" type="text" class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-100 focus:border-[#FF3D3D] outline-none transition-all font-medium text-gray-900" />
                            </div>
                             <div class="space-y-1.5">
                                <label class="text-xs font-bold text-gray-500 uppercase tracking-wider">Area / Village</label>
                                <input v-model="form.areaVillage" type="text" class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-100 focus:border-[#FF3D3D] outline-none transition-all font-medium text-gray-900" />
                            </div>
                             <div class="space-y-1.5">
                                <label class="text-xs font-bold text-gray-500 uppercase tracking-wider">Postal Code</label>
                                <input v-model="form.postalCode" type="text" class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-100 focus:border-[#FF3D3D] outline-none transition-all font-medium text-gray-900" />
                            </div>
                             <div class="space-y-1.5 col-span-2">
                                <label class="text-xs font-bold text-gray-500 uppercase tracking-wider">Google Maps Link</label>
                                <input v-model="form.googleMapLink" type="url" placeholder="https://maps.google.com/..." class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-100 focus:border-[#FF3D3D] outline-none transition-all font-medium text-gray-900 text-sm" />
                            </div>
                        </div>
  
                         <!-- Status Tab -->
                        <div v-show="activeTab === 'status'" class="space-y-6">
                            <div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
                                 <div class="space-y-1.5">
                                    <label class="text-xs font-bold text-gray-500 uppercase tracking-wider">Role</label>
                                    <select v-model="form.role" class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-100 focus:border-[#FF3D3D] outline-none transition-all font-medium text-gray-900">
                                        <option value="donor">Donor</option>
                                        <option value="admin">Admin</option>
                                    </select>
                                </div>
                                 <div class="space-y-1.5">
                                    <label class="text-xs font-bold text-gray-500 uppercase tracking-wider">Last Donation Date</label>
                                    <input v-model="form.lastDonationDate" type="date" class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-red-100 focus:border-[#FF3D3D] outline-none transition-all font-medium text-gray-900" />
                                </div>
                            </div>
  
                            <div class="bg-gray-50 rounded-xl p-4 border border-gray-200 space-y-3">
                                <label class="flex items-center gap-3 cursor-pointer">
                                    <input type="checkbox" v-model="form.isActive" class="w-5 h-5 text-[#FF3D3D] rounded border-gray-300 focus:ring-[#FF3D3D]" />
                                    <div>
                                        <span class="block text-sm font-bold text-gray-900">Account Active</span>
                                        <span class="block text-xs text-gray-500">Enable or disable this user account</span>
                                    </div>
                                </label>
                                 <label class="flex items-center gap-3 cursor-pointer">
                                    <input type="checkbox" v-model="form.isAvailable" class="w-5 h-5 text-emerald-500 rounded border-gray-300 focus:ring-emerald-500" />
                                    <div>
                                        <span class="block text-sm font-bold text-gray-900">Available to Donate</span>
                                        <span class="block text-xs text-gray-500">Show this donor in search results</span>
                                    </div>
                                </label>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
  
             <!-- Footer -->
            <div class="px-6 py-4 border-t border-gray-100 bg-gray-50/50 flex justify-end gap-3 shrink-0">
                <button @click="closeModal" class="px-6 py-2.5 rounded-xl font-bold text-gray-600 hover:bg-gray-100 transition-colors text-sm">Cancel</button>
                <button @click="saveUser" class="px-6 py-2.5 rounded-xl font-bold text-white bg-[#FF3D3D] hover:bg-red-600 transition-colors shadow-lg shadow-red-500/20 text-sm flex items-center gap-2">
                    <span v-if="!currentUserId">Create User</span>
                    <span v-else>Update User</span>
                    <span class="material-icons text-sm">check</span>
                </button>
            </div>
        </div>
    </div>
    <!-- Donation Record Modal -->
    <div v-if="isDonationModalOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4 sm:p-6">
        <div class="absolute inset-0 bg-gray-900/60 backdrop-blur-xs transition-opacity" @click="closeDonationModal"></div>
        <div class="relative bg-white rounded-2xl shadow-xl w-full max-w-md flex flex-col overflow-hidden transform transition-all animate-fade-in">
            <div class="px-5 py-4 border-b border-gray-100 flex items-center justify-between bg-gray-50/50">
                <h3 class="font-bold text-gray-900 flex items-center gap-2">
                    <span class="w-8 h-8 rounded-lg bg-red-50 text-[#FF3D3D] flex items-center justify-center">
                        <span class="material-icons text-sm">bloodtype</span>
                    </span>
                    Record Donation
                </h3>
                <button @click="closeDonationModal" class="w-8 h-8 rounded-lg hover:bg-gray-100 flex items-center justify-center text-gray-400 transition-colors">
                    <span class="material-icons text-lg">close</span>
                </button>
            </div>
            
            <div class="p-5 space-y-4">
                <div v-if="selectedDonor" class="flex items-center gap-3 p-3 bg-red-50 rounded-xl border border-red-100">
                    <div class="w-10 h-10 rounded-full bg-white flex items-center justify-center text-[#FF3D3D] font-bold shadow-sm">
                        {{ selectedDonor.blood_group }}
                    </div>
                    <div>
                        <div class="text-sm font-bold text-gray-900">{{ selectedDonor.name }}</div>
                        <div class="text-xs text-red-600 font-medium">Adding new donation record</div>
                    </div>
                </div>

                <div class="space-y-3">
                    <div class="grid grid-cols-2 gap-3">
                        <div class="space-y-1.5">
                            <label class="text-xs font-bold text-gray-500 uppercase">Date</label>
                            <input v-model="donationForm.date" type="date" class="w-full px-3 py-2 bg-gray-50 border border-gray-200 rounded-lg focus:ring-2 focus:ring-red-100 focus:border-[#FF3D3D] outline-none transition-all text-sm" />
                        </div>
                        <div class="space-y-1.5">
                            <label class="text-xs font-bold text-gray-500 uppercase">Type</label>
                            <select v-model="donationForm.type" class="w-full px-3 py-2 bg-gray-50 border border-gray-200 rounded-lg focus:ring-2 focus:ring-red-100 focus:border-[#FF3D3D] outline-none transition-all text-sm">
                                <option value="Whole Blood">Whole Blood</option>
                                <option value="Platelets">Platelets</option>
                                <option value="Plasma">Plasma</option>
                            </select>
                        </div>
                    </div>

                    <div class="space-y-1.5">
                        <label class="text-xs font-bold text-gray-500 uppercase">Location / Hospital</label>
                        <input v-model="donationForm.location" type="text" placeholder="e.g. DMCH, Dhaka" class="w-full px-3 py-2 bg-gray-50 border border-gray-200 rounded-lg focus:ring-2 focus:ring-red-100 focus:border-[#FF3D3D] outline-none transition-all text-sm" />
                    </div>

                    <div class="space-y-1.5">
                        <label class="text-xs font-bold text-gray-500 uppercase">Amount (ml)</label>
                        <input v-model="donationForm.amount_ml" type="number" class="w-full px-3 py-2 bg-gray-50 border border-gray-200 rounded-lg focus:ring-2 focus:ring-red-100 focus:border-[#FF3D3D] outline-none transition-all text-sm" />
                    </div>

                    <div class="space-y-1.5">
                        <label class="text-xs font-bold text-gray-500 uppercase">Notes (Optional)</label>
                        <textarea v-model="donationForm.notes" rows="2" class="w-full px-3 py-2 bg-gray-50 border border-gray-200 rounded-lg focus:ring-2 focus:ring-red-100 focus:border-[#FF3D3D] outline-none transition-all text-sm resize-none"></textarea>
                    </div>

                    <div class="space-y-1.5">
                        <label class="text-xs font-bold text-gray-500 uppercase">Evidence Image (Optional)</label>
                        <input type="file" @change="onDonationImageSelected" accept="image/*" class="w-full text-xs text-gray-500 file:mr-2 file:py-2 file:px-3 file:rounded-lg file:border-0 file:text-xs file:font-semibold file:bg-red-50 file:text-red-700 hover:file:bg-red-100 transition-all"/>
                    </div>
                </div>

                <div class="pt-2">
                    <button @click="saveDonation" class="w-full py-2.5 rounded-xl font-bold text-white bg-[#FF3D3D] hover:bg-red-600 transition-colors shadow-lg shadow-red-500/20 text-sm flex items-center justify-center gap-2">
                        <span class="material-icons text-sm">save</span> Confirm Donation
                    </button>
                </div>
            </div>
        </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed, reactive, watch } from 'vue';
import api from '@/lib/axios';
import { useToastStore } from '@/stores/toast';

const toastStore = useToastStore();

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
    is_admin_verified: boolean;
    is_active: boolean;
}

const users = ref<User[]>([]);
const searchQuery = ref('');

const showFilters = ref(false);
const filters = reactive({
    bloodGroups: [] as string[],
    roles: [] as string[],
    is_active: '',
    is_available: '',
    verified_status: ''
});

const clearFilters = () => {
    filters.bloodGroups = [];
    filters.roles = [];
    filters.is_active = '';
    filters.is_available = '';
    filters.verified_status = '';
};

const activeFilterCount = computed(() => {
    let count = 0;
    if (filters.bloodGroups.length > 0) count++;
    if (filters.roles.length > 0) count++;
    if (filters.is_active !== '') count++;
    if (filters.is_available !== '') count++;
    if (filters.verified_status !== '') count++;
    return count;
});

const loading = ref(false);
const currentPage = ref(1);
const totalUsers = ref(0);
const totalPages = ref(0);
const limit = 20;

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

const isDonationModalOpen = ref(false);
const selectedDonor = ref<User | null>(null);
const donationForm = reactive({
    date: new Date().toISOString().split('T')[0],
    type: 'Whole Blood',
    location: '',
    amount_ml: 450,
    notes: ''
});

const selectedDonationImage = ref<File | null>(null);

const onDonationImageSelected = (e: Event) => {
    const target = e.target as HTMLInputElement;
    if (target.files && target.files[0]) {
        selectedDonationImage.value = target.files[0];
    }
};

const openDonationModal = (user: User) => {
    selectedDonor.value = user;
    donationForm.date = new Date().toISOString().split('T')[0];
    donationForm.type = 'Whole Blood';
    donationForm.location = [user.city, user.district].filter(Boolean).join(', ') || '';
    donationForm.amount_ml = 450;
    donationForm.notes = '';
    isDonationModalOpen.value = true;
};

const closeDonationModal = () => {
    isDonationModalOpen.value = false;
    selectedDonor.value = null;
    selectedDonationImage.value = null;
};

const saveDonation = async () => {
    if (!selectedDonor.value) return;
    
    
    try {
        const formData = new FormData();
        formData.append('date', new Date(donationForm.date || new Date().toISOString()).toISOString());
        formData.append('type', donationForm.type);
        formData.append('location', donationForm.location);
        formData.append('amount_ml', donationForm.amount_ml.toString());
        formData.append('notes', donationForm.notes);

        if (selectedDonationImage.value) {
            formData.append('image', selectedDonationImage.value);
        }

        // 1. Create donation record
        await api.post(`/admin/users/${selectedDonor.value.id}/donations`, formData, {
            headers: { 'Content-Type': 'multipart/form-data' }
        });

        // 2. Update user status (Last donation date & Availability)
        const today = new Date().toISOString();
        await api.put(`/admin/users/${selectedDonor.value.id}`, {
            last_donation_date: new Date(donationForm.date || new Date().toISOString()).toISOString(),
            is_available: false 
        });

        // 3. Update local state
        const index = users.value.findIndex(u => u.id === selectedDonor.value?.id);
        if (index !== -1) {
             users.value[index] = {
                ...users.value[index],
                last_donation_date: new Date(donationForm.date || new Date().toISOString()).toISOString(),
                is_available: false
            } as User;
        }

        toastStore.show('Donation recorded successfully!', 'success');
        closeDonationModal();
    } catch (error) {
        console.error("Failed to save donation", error);
        toastStore.show("Failed to record donation", "error");
    }
};

const markDonatedToday = (user: User) => {
    openDonationModal(user);
};

const toggleAdminVerification = async (user: User) => {
    try {
        const newStatus = !user.is_admin_verified;
        await api.put(`/admin/users/${user.id}`, {
            is_admin_verified: newStatus
        });
        user.is_admin_verified = newStatus;
        toastStore.show(newStatus ? 'Donor verified successfully' : 'Verification removed', 'success');
    } catch (error) {
        toastStore.show('Failed to update verification status', 'error');
    }
};

const filteredUsers = computed(() => users.value);

const loadUsers = async (page: number = 1) => {
    loading.value = true;
    try {
        const queryParams = new URLSearchParams();
        queryParams.append('page', page.toString());
        queryParams.append('limit', limit.toString());
        if (searchQuery.value) queryParams.append('q', searchQuery.value);
        if (filters.verified_status) queryParams.append('verified_status', filters.verified_status);
        if (filters.bloodGroups.length > 0) queryParams.append('blood_groups', filters.bloodGroups.join(','));
        if (filters.roles.length > 0) queryParams.append('role',  filters.roles.join(','));
        if (filters.is_active !== '') queryParams.append('is_active', filters.is_active);
        if (filters.is_available !== '') queryParams.append('is_available', filters.is_available);

        const res = await api.get(`/admin/users?${queryParams.toString()}`);
        users.value = res.data.users || [];
        totalUsers.value = res.data.total || 0;
        totalPages.value = res.data.pages || 0;
        currentPage.value = page;
    } catch (error) {
        console.error("Failed to load users", error);
    } finally {
        loading.value = false;
    }
};

const changePage = (page: number) => {
    if (page >= 1 && page <= totalPages.value) {
        loadUsers(page);
    }
};

const visiblePages = computed(() => {
    const pages = [];
    const maxVisible = 5;
    let start = Math.max(1, currentPage.value - Math.floor(maxVisible / 2));
    let end = Math.min(totalPages.value, start + maxVisible - 1);
    
    if (end - start < maxVisible - 1) {
        start = Math.max(1, end - maxVisible + 1);
    }
    
    for (let i = start; i <= end; i++) {
        pages.push(i);
    }
    return pages;
});

let searchTimeout: any = null;
watch([searchQuery, filters], () => {
    clearTimeout(searchTimeout);
    searchTimeout = setTimeout(() => {
        loadUsers(1);
    }, 500);
}, { deep: true });

const openAddModal = () => {
    modalMode.value = 'add';
    currentUserId.value = null;
    activeTab.value = 'profile';
    resetForm();
    isModalOpen.value = true;
};

const closeModal = () => {
    isModalOpen.value = false;
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
        toastStore.show("Failed to save user. Please check inputs.", "error");
    }
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
        toastStore.show('Failed to delete user', 'error');
    }
};

// Dropdown state
const openDropdownId = ref<string | null>(null);

const toggleDropdown = (id: string) => {
    openDropdownId.value = openDropdownId.value === id ? null : id;
};

const closeDropdown = () => {
    openDropdownId.value = null;
};

// Global click listener to close dropdown when clicking outside
const handleClickOutside = (e: MouseEvent) => {
    const target = e.target as HTMLElement;
    if (!target.closest('.dropdown-container')) {
        closeDropdown();
    }
};

onMounted(() => {
    loadUsers();
    document.addEventListener('click', handleClickOutside);
});

onUnmounted(() => {
    document.removeEventListener('click', handleClickOutside);
});
</script>
