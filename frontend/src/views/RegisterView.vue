<template>
  <div class="h-full bg-gray-50 flex flex-col font-sans grow">
    
    <!-- Main Content -->
    <div class="grow flex items-center justify-center py-4 px-4 sm:px-6 lg:px-8">
      <div class="w-full max-w-5xl space-y-8">
        
        <!-- Hero Text -->
        <div class="text-center space-y-3 px-4">
          <h1 class="text-3xl sm:text-4xl font-black text-gray-900 tracking-tight">
            Join Us & <span class="text-red-600">Save Lives</span>
          </h1>
          <p class="text-gray-500 text-sm sm:text-base max-w-xl mx-auto font-medium leading-relaxed">
            Register as a blood donor today and become a lifeline for those in need.
          </p>
        </div>

        <!-- Registration Form Card -->
        <div class="bg-white rounded-3xl shadow-xl overflow-hidden border border-gray-100">
             <!-- Top Red Line -->
            <div class="h-1.5 bg-red-600 w-full"></div>

            <!-- Success Content -->
            <div v-if="successState" class="p-10 text-center space-y-6 animate-in zoom-in duration-500">
                <div class="bg-green-100 text-green-600 rounded-full w-20 h-20 flex items-center justify-center mx-auto shadow-inner">
                <span class="material-icons text-5xl">mark_email_read</span>
                </div>
                <div class="space-y-2 text-balance">
                <h2 class="text-2xl font-black text-gray-900">Registration Successful!</h2>
                <p class="text-gray-500 font-medium px-4">
                    We've sent a verification link to <span class="text-gray-900 font-bold">{{ email }}</span>. 
                    Please check your inbox and verify your account to start donating.
                </p>
                </div>
                <div class="pt-4">
                    <router-link to="/login" class="inline-flex items-center justify-center gap-2 bg-[#E53935] hover:bg-red-700 text-white font-bold py-4 px-10 rounded-xl shadow-lg shadow-red-200 transition-all duration-300 transform hover:-translate-y-1 active:scale-95">
                        Go to Login
                        <span class="material-icons text-lg">arrow_forward</span>
                    </router-link>
                </div>
            </div>

            <form v-else @submit.prevent="handleRegister" class="p-6 sm:p-10 space-y-8">
                <!-- Inline Error Message -->
                <div v-if="errorMsg" class="p-4 rounded-xl bg-red-50 border border-red-100 flex items-start gap-3 animate-in fade-in slide-in-from-top-2 duration-300">
                    <span class="material-icons text-red-500 text-xl">error_outline</span>
                    <div class="flex-1">
                        <p class="text-sm font-semibold text-red-800">{{ errorMsg }}</p>
                    </div>
                    <button @click="errorMsg = ''" type="button" class="text-red-400 hover:text-red-600 transition-colors">
                        <span class="material-icons text-lg">close</span>
                    </button>
                </div>

                <!-- Personal Information -->
                <section class="space-y-6">
                    <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-2 border-b border-gray-100 pb-3">
                        <div class="flex items-center gap-2">
                            <span class="material-icons text-red-500">person_outline</span>
                            <h3 class="text-lg font-bold text-gray-800">Personal Information</h3>
                        </div>
                        <span class="text-xs text-gray-400 italic">Fields marked with <span class="text-red-600">*</span> are required</span>
                    </div>

                    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                        <div class="space-y-1">
                            <label for="fullName" class="text-sm font-medium text-gray-700">Full Name <span class="text-red-600">*</span></label>
                            <input v-model="name" type="text" id="fullName" placeholder="John Doe" required
                                class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-red-500 focus:border-red-500 outline-none transition-all placeholder-gray-400 text-gray-800" />
                        </div>

                        <div class="space-y-1">
                            <label for="email" class="text-sm font-medium text-gray-700">Email Address <span class="text-red-600">*</span></label>
                            <input v-model="email" type="email" id="email" placeholder="john@example.com" required
                                class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-red-500 focus:border-red-500 outline-none transition-all placeholder-gray-400 text-gray-800" />
                        </div>

                         <div class="space-y-1">
                            <label for="phone" class="text-sm font-medium text-gray-700">Phone Number <span class="text-red-600">*</span></label>
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
                            <label for="bloodGroup" class="text-sm font-medium text-gray-700">Blood Group <span class="text-red-600">*</span></label>
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
                            <label for="division" class="text-sm font-medium text-gray-700">Division <span class="text-red-600">*</span></label>
                            <input v-model="division" type="text" id="division" placeholder="e.g. Dhaka" required
                                class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-red-500 focus:border-red-500 outline-none transition-all placeholder-gray-400 text-gray-800" />
                        </div>

                        <div class="space-y-1">
                            <label for="district" class="text-sm font-medium text-gray-700">District <span class="text-red-600">*</span></label>
                            <input v-model="district" type="text" id="district" placeholder="e.g. Gazipur" required
                                class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-red-500 focus:border-red-500 outline-none transition-all placeholder-gray-400 text-gray-800" />
                        </div>

                        <div class="space-y-1">
                            <label for="upazila" class="text-sm font-medium text-gray-700">Upazila <span class="text-red-600">*</span></label>
                            <input v-model="upazila" type="text" id="upazila" placeholder="e.g. Savar" required
                                class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-red-500 focus:border-red-500 outline-none transition-all placeholder-gray-400 text-gray-800" />
                        </div>
                        
                         <div class="space-y-1">
                            <label for="city" class="text-sm font-medium text-gray-700">City</label>
                            <input v-model="city" type="text" id="city" placeholder="e.g. Dhaka City"
                                class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-red-500 focus:border-red-500 outline-none transition-all placeholder-gray-400 text-gray-800" />
                        </div>

                         <div class="space-y-1">
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
                            <label for="password" class="text-sm font-medium text-gray-700">Password <span class="text-red-600">*</span></label>
                            <div class="relative">
                                <input v-model="password" :type="showPassword ? 'text' : 'password'" id="password" placeholder="••••••••" required
                                    class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-red-500 focus:border-red-500 outline-none transition-all placeholder-gray-400 text-gray-800" />
                                <button type="button" @click="showPassword = !showPassword" class="absolute right-3 top-2.5 text-gray-400 hover:text-gray-600 focus:outline-none">
                                    <span class="material-icons text-lg">{{ showPassword ? 'visibility' : 'visibility_off' }}</span>
                                </button>
                            </div>
                        </div>

                         <div class="space-y-1">
                            <label for="confirmPassword" class="text-sm font-medium text-gray-700">Confirm Password <span class="text-red-600">*</span></label>
                            <div class="relative">
                                <input v-model="confirmPassword" :type="showConfirmPassword ? 'text' : 'password'" id="confirmPassword" placeholder="••••••••" required
                                    class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-red-500 focus:border-red-500 outline-none transition-all placeholder-gray-400 text-gray-800" />
                                <button type="button" @click="showConfirmPassword = !showConfirmPassword" class="absolute right-3 top-2.5 text-gray-400 hover:text-gray-600 focus:outline-none">
                                    <span class="material-icons text-lg">{{ showConfirmPassword ? 'visibility' : 'visibility_off' }}</span>
                                </button>
                            </div>
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
const area = ref('');
const city = ref(''); // Added City
const postalCode = ref('');
const password = ref('');
const confirmPassword = ref('');
const agreeTerms = ref(false);
const showPassword = ref(false);
const showConfirmPassword = ref(false);
const gender = ref('');
const birthday = ref('');

const errorMsg = ref('');
const successState = ref(false);

const authStore = useAuthStore();
const router = useRouter();

// Address Fields
const division = ref('');
const district = ref('');
const upazila = ref('');

const handleRegister = async () => {
    errorMsg.value = '';
    
    if (password.value !== confirmPassword.value) {
        errorMsg.value = "Passwords do not match";
        return;
    }
    
    // Basic validation
    if (!agreeTerms.value) {
        errorMsg.value = "Please agree to the Terms of Service";
        return;
    }

  try {
    await authStore.register({
      name: name.value,
      email: email.value,
      phone: phone.value,
      password: password.value,
      gender: gender.value,
      birthday: birthday.value,
      bloodGroup: bloodGroup.value, 
      division: division.value,
      district: district.value,
      upazila: upazila.value,
      city: city.value,
      area: area.value,
      postalCode: postalCode.value
    });
    successState.value = true;
    window.scrollTo({ top: 0, behavior: 'smooth' });
  } catch (error: any) {
    console.error(error);
    errorMsg.value = error.response?.data?.error || 'Registration failed. Please try again.';
    window.scrollTo({ top: 0, behavior: 'smooth' });
  }
};
</script>

<style scoped>
/* Scoped styles if absolutely necessary, but relying on Tailwind */
</style>
