<template>
  <div class="h-full bg-gray-50 flex flex-col font-sans grow">
    
    <!-- Main Content -->
    <div class="grow flex items-center justify-center py-4 px-4 sm:px-6 lg:px-8">
      <div class="w-full max-w-5xl space-y-8">
        
        <!-- Hero Text -->
        <div class="text-center space-y-2">
          <h1 class="text-2xl md:text-3xl font-extrabold text-gray-900 tracking-tight">
            Join Us & <span class="text-red-600">Save Lives</span>
          </h1>
          <p class="text-gray-500 text-lg max-w-2xl mx-auto">
            Register as a blood donor today and become a lifeline for those in need.
          </p>
        </div>

        <!-- Registration Form Card -->
        <div class="bg-white rounded-2xl shadow-xl overflow-hidden border border-gray-100">
             <!-- Top Red Line -->
            <div class="h-2 bg-red-600 w-full"></div>

            <form @submit.prevent="handleRegister" class="p-5 md:p-6 space-y-5">
                
                <!-- Personal Information -->
                <section class="space-y-4">
                    <div class="flex items-center gap-2 border-b border-gray-100 pb-2">
                        <span class="material-icons text-red-500">person_outline</span>
                        <h3 class="text-lg font-semibold text-gray-800">Personal Information</h3>
                    </div>

                    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                        <div class="space-y-1">
                            <label for="fullName" class="text-sm font-medium text-gray-700">Full Name</label>
                            <input v-model="name" type="text" id="fullName" placeholder="John Doe" required
                                class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-red-500 focus:border-red-500 outline-none transition-all placeholder-gray-400 text-gray-800" />
                        </div>

                        <div class="space-y-1">
                            <label for="email" class="text-sm font-medium text-gray-700">Email Address</label>
                            <input v-model="email" type="email" id="email" placeholder="john@example.com" required
                                class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-red-500 focus:border-red-500 outline-none transition-all placeholder-gray-400 text-gray-800" />
                        </div>

                         <div class="space-y-1">
                            <label for="phone" class="text-sm font-medium text-gray-700">Phone Number</label>
                            <input v-model="phone" type="tel" id="phone" placeholder="+880 1XXX-XXXXXX" required
                                class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-red-500 focus:border-red-500 outline-none transition-all placeholder-gray-400 text-gray-800" />
                        </div>

                        <div class="col-span-1 md:col-span-2 lg:col-span-3 space-y-1">
                            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                                <div class="space-y-1">
                                    <label for="gender" class="text-sm font-medium text-gray-700">Gender</label>
                                    <select v-model="gender" id="gender"
                                        class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-red-500 focus:border-red-500 outline-none transition-all bg-white text-gray-800 appearance-none cursor-pointer">
                                        <option value="" disabled selected>Select</option>
                                        <option value="Male">Male</option>
                                        <option value="Female">Female</option>
                                        <option value="Other">Other</option>
                                    </select>
                                </div>
                                <div class="space-y-1">
                                    <label for="birthday" class="text-sm font-medium text-gray-700">Birthday</label>
                                    <input v-model="birthday" type="date" id="birthday"
                                        class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-red-500 focus:border-red-500 outline-none transition-all placeholder-gray-400 text-gray-800" />
                                </div>
                            </div>
                        </div>

                        <div class="col-span-1 md:col-span-2 lg:col-span-3 space-y-1">
                            <label for="bloodGroup" class="text-sm font-medium text-gray-700">Blood Group</label>
                            <div class="relative">
                                <select v-model="bloodGroup" id="bloodGroup" required
                                    class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-red-500 focus:border-red-500 outline-none transition-all bg-white text-gray-800 appearance-none cursor-pointer">
                                    <option value="" disabled selected>Select your blood group</option>
                                    <option value="A+">A+</option>
                                    <option value="A-">A-</option>
                                    <option value="B+">B+</option>
                                    <option value="B-">B-</option>
                                    <option value="O+">O+</option>
                                    <option value="O-">O-</option>
                                    <option value="AB+">AB+</option>
                                    <option value="AB-">AB-</option>
                                </select>
                                <span class="material-icons absolute right-4 top-3 text-red-500 pointer-events-none">bloodtype</span>
                            </div>
                            <p class="text-xs text-gray-500 mt-1">Your blood group is vital for matching donors with recipients.</p>
                        </div>
                    </div>
                </section>

                <!-- Address Details -->
                 <section class="space-y-4">
                    <div class="flex items-center gap-2 border-b border-gray-100 pb-2">
                        <span class="material-icons text-red-500">place</span>
                        <h3 class="text-lg font-semibold text-gray-800">Address Details</h3>
                    </div>

                    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
                         <div class="space-y-1">
                            <label for="division" class="text-sm font-medium text-gray-700">Division</label>
                            <select v-model="division" id="division"
                                class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-red-500 focus:border-red-500 outline-none transition-all bg-white text-gray-800">
                                <option>Dhaka</option>
                                <option>Chittagong</option>
                                <option>Sylhet</option>
                                <option>Khulna</option>
                            </select>
                        </div>

                        <div class="space-y-1">
                            <label for="district" class="text-sm font-medium text-gray-700">District</label>
                            <select v-model="district" id="district"
                                class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-red-500 focus:border-red-500 outline-none transition-all bg-white text-gray-800">
                                <option>Dhaka</option>
                                <option>Gazipur</option>
                                <option>Narayanganj</option>
                            </select>
                        </div>

                        <div class="space-y-1">
                            <label for="upazila" class="text-sm font-medium text-gray-700">Upazila</label>
                            <select v-model="upazila" id="upazila"
                                class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-red-500 focus:border-red-500 outline-none transition-all bg-white text-gray-800">
                                <option>Mirpur</option>
                                <option>Savar</option>
                                <option>Dhanmondi</option>
                            </select>
                        </div>
                        
                         <div class="col-span-1 md:col-span-2 space-y-1">
                            <label for="area" class="text-sm font-medium text-gray-700">Area / Village / Road</label>
                            <input v-model="area" type="text" id="area" placeholder="House 12, Road 5, Block B"
                                class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-red-500 focus:border-red-500 outline-none transition-all placeholder-gray-400 text-gray-800" />
                        </div>

                        <div class="space-y-1">
                            <label for="postalCode" class="text-sm font-medium text-gray-700">Postal Code</label>
                            <input v-model="postalCode" type="text" id="postalCode" placeholder="1216"
                                class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-red-500 focus:border-red-500 outline-none transition-all placeholder-gray-400 text-gray-800" />
                        </div>
                    </div>
                </section>

                <!-- Security -->
                <section class="space-y-4">
                    <div class="flex items-center gap-2 border-b border-gray-100 pb-2">
                        <span class="material-icons text-red-500">lock_outline</span>
                        <h3 class="text-lg font-semibold text-gray-800">Security</h3>
                    </div>

                    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                        <div class="space-y-1">
                            <label for="password" class="text-sm font-medium text-gray-700">Password</label>
                            <div class="relative">
                                <input v-model="password" :type="showPassword ? 'text' : 'password'" id="password" placeholder="••••••••" required
                                    class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-red-500 focus:border-red-500 outline-none transition-all placeholder-gray-400 text-gray-800" />
                                <button type="button" @click="showPassword = !showPassword" class="absolute right-3 top-2.5 text-gray-400 hover:text-gray-600 focus:outline-none">
                                    <span class="material-icons text-lg">{{ showPassword ? 'visibility' : 'visibility_off' }}</span>
                                </button>
                            </div>
                        </div>

                         <div class="space-y-1">
                            <label for="confirmPassword" class="text-sm font-medium text-gray-700">Confirm Password</label>
                            <input v-model="confirmPassword" type="password" id="confirmPassword" placeholder="••••••••" required
                                class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-red-500 focus:border-red-500 outline-none transition-all placeholder-gray-400 text-gray-800" />
                        </div>
                    </div>
                </section>

                <!-- Terms & Submit -->
                 <div class="space-y-6 pt-4">
                    <div class="flex items-start gap-3">
                        <div class="flex items-center h-5">
                            <input v-model="agreeTerms" id="terms" type="checkbox" required
                                class="w-4 h-4 text-red-600 border-gray-300 rounded focus:ring-red-500 focus:ring-2" />
                        </div>
                        <div class="text-sm">
                            <label for="terms" class="font-medium text-gray-700">
                                I agree to the <a href="#" class="text-red-600 hover:text-red-700 hover:underline">Terms of Service</a> and <a href="#" class="text-red-600 hover:text-red-700 hover:underline">Privacy Policy</a>.
                            </label>
                            <p class="text-gray-500 mt-1">I consent to being contacted for blood donation requests.</p>
                        </div>
                    </div>

                    <button type="submit" 
                        class="w-full bg-red-600 hover:bg-red-700 text-white font-bold py-3 px-4 rounded-lg shadow-lg hover:shadow-xl transition-all duration-300 transform hover:-translate-y-0.5 flex items-center justify-center gap-2 group">
                        <span class="material-icons group-hover:animate-pulse">favorite</span>
                        Register Now
                    </button>
                </div>
            </form>
             <div class="px-8 py-4 bg-gray-50 border-t border-gray-100 text-center">
                <p class="text-sm text-gray-500">
                    Already have an account? <router-link to="/login" class="text-red-600 font-semibold hover:text-red-700 hover:underline">Sign in instead</router-link>
                </p>
            </div>
        </div>

        <!-- Trust Badges -->
        <div class="flex justify-center gap-8 md:gap-16 pt-4 pb-8 opacity-75 grayscale hover:grayscale-0 transition-all duration-500">
             <div class="flex flex-col items-center gap-1">
                <span class="material-icons text-3xl text-red-600">verified_user</span>
                <span class="text-xs font-semibold text-gray-500">Secure Data</span>
            </div>
            <div class="flex flex-col items-center gap-1">
                <span class="material-icons text-3xl text-red-600">diversity_3</span>
                <span class="text-xs font-semibold text-gray-500">Community Driven</span>
            </div>
             <div class="flex flex-col items-center gap-1">
                <span class="material-icons text-3xl text-red-600">verified</span>
                <span class="text-xs font-semibold text-gray-500">Verified Donors</span>
            </div>
        </div>
        
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router';

const name = ref('');
const email = ref('');
const phone = ref('');
const bloodGroup = ref('');
const division = ref('Dhaka');
const district = ref('Dhaka');
const upazila = ref('Mirpur');
const area = ref('');
const postalCode = ref('');
const password = ref('');
const confirmPassword = ref('');
const agreeTerms = ref(false);
const showPassword = ref(false);
const gender = ref('');
const birthday = ref('');

const authStore = useAuthStore();
const router = useRouter();

const handleRegister = async () => {
    if (password.value !== confirmPassword.value) {
        alert("Passwords do not match");
        return;
    }
    
    // Basic validation
    if (!agreeTerms.value) {
        alert("Please agree to the Terms of Service");
        return;
    }

  try {
    await authStore.register({
      name: name.value,
      email: email.value,
      phone: phone.value,
      password: password.value,
      // Pass other fields if backend supports them, otherwise they are just UI in this mock
      ...({
        gender: '',
        birthday: '',
        bloodGroup: bloodGroup.value, 
        division: division.value,
        district: district.value,
        upazila: upazila.value,
        area: area.value,
        postalCode: postalCode.value
      } as any)
    });
    alert('Registration successful! Please login.');
    router.push('/login');
  } catch (error) {
    console.error(error);
    alert('Failed to register. Please try again.');
  }
};
</script>

<style scoped>
/* Scoped styles if absolutely necessary, but relying on Tailwind */
</style>
