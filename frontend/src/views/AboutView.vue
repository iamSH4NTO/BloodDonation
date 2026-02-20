<template>
  <div class="bg-white min-h-[calc(100vh-6rem)] font-sans selection:bg-red-100 selection:text-red-900">
    <div class="max-w-3xl mx-auto px-6 py-10 lg:py-16">
      
      <!-- Minimalist Header -->
      <header class="mb-10 lg:mb-12 text-center sm:text-left">
        <div class="inline-flex items-center gap-2 mb-6">
          <div class="w-10 h-10 bg-red-50 rounded-2xl flex items-center justify-center">
            <span class="material-icons text-[#FF3D3D] text-xl">bloodtype</span>
          </div>
          <span class="text-sm font-bold uppercase tracking-widest text-[#FF3D3D]">BloodCare</span>
        </div>
        <h1 class="text-4xl lg:text-6xl font-black text-[#1A1A1A] leading-tight tracking-tighter">
          Simplifying the way <br /> we <span class="text-[#FF3D3D]">save lives</span>.
        </h1>
      </header>

      <!-- Clean Content Sections -->
      <div class="space-y-10 lg:space-y-14">
        
        <!-- Mission Section -->
        <section class="max-w-2xl">
          <p class="text-xl lg:text-2xl text-gray-500 font-medium leading-relaxed">
            We believe that finding a blood donor shouldn't be a race against time. <span class="text-gray-900">BloodCare</span> is built to connect heroes with those in need, instantly and securely.
          </p>
        </section>

        <!-- Two Column Features -->
        <section class="grid sm:grid-cols-2 gap-12 lg:gap-16">
          <div class="space-y-4">
            <h3 class="text-lg font-bold text-gray-900">Fast & Secure</h3>
            <p class="text-gray-500 font-medium leading-relaxed">
              Our platform focuses on speed without compromising privacy. Verified donors, direct contact, and no unnecessary hurdles.
            </p>
          </div>
          <div class="space-y-4">
            <h3 class="text-lg font-bold text-gray-900">Community First</h3>
            <p class="text-gray-500 font-medium leading-relaxed">
              A free resource for the community, by the community. Empowering everyone to make a difference when it matters most.
            </p>
          </div>
        </section>

        <!-- Minimalist List -->
        <section class="py-12 border-y border-gray-100">
           <div class="grid sm:grid-cols-2 gap-x-12 gap-y-6">
              <div v-for="benefit in benefits" :key="benefit" class="flex items-center gap-3">
                <span class="w-1.5 h-1.5 rounded-full bg-[#FF3D3D]"></span>
                <span class="text-sm font-bold text-gray-700 tracking-tight">{{ benefit }}</span>
              </div>
           </div>
        </section>

        <!-- Call to Action -->
        <footer class="pt-8 text-center sm:text-left">
          <router-link to="/register" class="group inline-flex items-center gap-4">
            <span class="text-lg font-black text-[#1A1A1A] group-hover:text-[#FF3D3D] transition-colors">Join as a donor today</span>
            <div class="w-12 h-12 bg-[#F3F3F3] group-hover:bg-[#FF3D3D] rounded-full flex items-center justify-center transition-all duration-300 transform group-hover:translate-x-1">
              <span class="material-icons text-[#1A1A1A] group-hover:text-white transition-colors">arrow_forward</span>
            </div>
          </router-link>
        </footer>

      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const benefits = [
  'Zero Platform Fees',
  'Verified Donor Database',
  'Privacy-First Protocol',
  'Real-time Search API',
  'Available Nationwide',
  'Open 24/7'
];
</script>
