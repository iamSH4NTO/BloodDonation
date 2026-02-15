<template>
  <div class="max-w-4xl mx-auto mt-10 p-6">
    <h2 class="text-2xl font-bold mb-6">Find Blood Donors</h2>
    <div class="bg-white p-4 rounded-lg shadow mb-6 flex space-x-4">
      <select v-model="filters.group" class="border rounded px-3 py-2">
        <option value="">All Blood Groups</option>
        <option value="A+">A+</option>
        <option value="A-">A-</option>
        <option value="B+">B+</option>
        <option value="B-">B-</option>
        <option value="AB+">AB+</option>
        <option value="AB-">AB-</option>
        <option value="O+">O+</option>
        <option value="O-">O-</option>
      </select>
      <input v-model="filters.district" type="text" placeholder="District" class="border rounded px-3 py-2" />
      <button @click="searchDonors" class="bg-red-600 text-white px-4 py-2 rounded hover:bg-red-700">Search</button>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div v-for="donor in donors" :key="donor.user_id" class="bg-white p-6 rounded-lg shadow">
        <h3 class="text-xl font-semibold">{{ donor.name }}</h3>
        <p class="text-red-600 font-bold">{{ donor.blood_group }}</p>
        <p class="text-gray-600">{{ donor.district }}, {{ donor.city }}</p>
        <p class="text-gray-500 text-sm mt-2">{{ donor.area_village }}</p>
        <button @click="$router.push(`/donors/${donor.user_id}`)" class="mt-4 w-full bg-red-100 text-red-700 font-bold py-2 rounded-lg hover:bg-red-200 transition-colors">View Profile</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import api from '@/lib/axios';

const filters = ref({ group: '', district: '' });
const donors = ref<any[]>([]);

const searchDonors = async () => {
    try {
        const res = await api.get('/donors', { params: filters.value });
        donors.value = res.data;
    } catch (error) {
        console.error(error);
    }
};


</script>
