<template>
  <div class="fixed inset-0 z-100 bg-gray-900 bg-opacity-75 flex items-center justify-center p-4 backdrop-blur-sm print:bg-white print:p-0" v-if="isOpen">
    <div class="bg-white rounded-2xl shadow-2xl max-w-4xl w-full mx-auto overflow-hidden print:shadow-none print:w-full print:h-screen print:flex print:items-center print:justify-center">
      
      <!-- Top Bar (Hidden in Print) -->
      <div class="bg-gray-50 border-b border-gray-100 px-6 py-4 flex justify-between items-center print:hidden">
        <h3 class="text-lg font-bold text-gray-900 flex items-center gap-2">
           <span class="material-icons text-[#FF3D3D]">print</span> Generate Certificate
        </h3>
        <div class="flex items-center gap-3">
             <button @click="handlePrint" class="px-5 py-2.5 rounded-xl bg-[#FF3D3D] hover:bg-red-600 text-white font-bold text-sm shadow-md transition-all flex items-center gap-2">
                <span class="material-icons text-sm">download</span> Download PDF/Print
            </button>
            <button @click="close" class="text-gray-400 hover:text-gray-600 p-2 rounded-lg hover:bg-gray-200 transition-colors">
                <span class="material-icons">close</span>
            </button>
        </div>
      </div>

      <!-- Certificate Content -->
      <div class="p-8 md:p-16 bg-white relative print:p-8" id="print-area">
          <!-- Decorative Corners -->
          <div class="absolute top-0 left-0 w-32 h-32 border-t-8 border-l-8 border-[#FF3D3D] rounded-tl-3xl opacity-20"></div>
          <div class="absolute top-0 right-0 w-32 h-32 border-t-8 border-r-8 border-[#FF3D3D] rounded-tr-3xl opacity-20"></div>
          <div class="absolute bottom-0 left-0 w-32 h-32 border-b-8 border-l-8 border-[#FF3D3D] rounded-bl-3xl opacity-20"></div>
          <div class="absolute bottom-0 right-0 w-32 h-32 border-b-8 border-r-8 border-[#FF3D3D] rounded-br-3xl opacity-20"></div>

          <div class="border-4 border-double border-gray-200 p-8 md:p-16 relative bg-white h-auto sm:min-h-[600px] flex flex-col justify-center text-center">
              
              <!-- Faded Background Logo -->
              <div class="absolute inset-0 flex items-center justify-center opacity-5 pointer-events-none">
                   <span class="material-icons" style="font-size: 400px;">water_drop</span>
              </div>

              <!-- Header -->
              <div class="mb-12 relative z-10">
                  <div class="flex items-center justify-center gap-3 text-[#FF3D3D] mb-4">
                      <span class="material-icons text-5xl">water_drop</span>
                      <h1 class="text-3xl font-black uppercase tracking-widest">LifeSaver Network</h1>
                  </div>
                  <h2 class="text-5xl md:text-6xl text-gray-900 typography-serif italic font-extralight tracking-tight mb-2">Certificate of Appreciation</h2>
                  <p class="text-sm font-bold text-gray-400 uppercase tracking-[0.2em] pt-4 border-t border-gray-100 max-w-sm mx-auto">Proudly Presented To</p>
              </div>

              <!-- Recipient -->
              <div class="mb-12 relative z-10">
                  <h3 class="text-4xl md:text-5xl font-bold text-[#FF3D3D] font-mono border-b-2 border-gray-200 inline-block px-12 pb-2">{{ displayName }}</h3>
              </div>

              <!-- Body text -->
              <div class="mb-12 max-w-2xl mx-auto relative z-10 space-y-4">
                  <p class="text-lg md:text-xl text-gray-600 font-medium leading-relaxed">
                      For exceptional dedication to saving lives. Your remarkable selflessness in reaching the 
                      <span class="font-bold text-gray-900 border-b border-gray-300">{{ badge }}</span> tier by completing 
                      <span class="font-bold text-gray-900 border-b border-gray-300">{{ totalDonations }} verified donations</span> 
                      has made an immeasurable difference in our community.
                  </p>
                   <p class="text-md text-gray-500 font-serif italic">
                      "A single drop of blood can make a huge difference."
                  </p>
              </div>

              <!-- Signatures & Date -->
              <div class="flex justify-between items-end max-w-3xl mx-auto w-full mt-12 relative z-10 px-8">
                  <div class="text-center w-48">
                      <div class="border-b border-gray-400 pb-2 mb-2 font-serif italic text-lg text-gray-700 -rotate-3 text-left">Dr. A. Rahman</div>
                      <p class="text-xs font-bold text-gray-400 uppercase tracking-widest">Head of Operations</p>
                  </div>
                   <div class="text-center">
                       <div class="w-24 h-24 mx-auto rounded-full border-4 border-[#FF3D3D] bg-red-50 flex flex-col items-center justify-center mb-2 shadow-lg transform rotate-12">
                           <span class="material-icons text-[#FF3D3D] text-3xl">verified</span>
                           <span class="text-[8px] font-black text-[#FF3D3D] uppercase tracking-wider mt-1">Official Seal</span>
                       </div>
                  </div>
                  <div class="text-center w-48">
                      <div class="border-b border-gray-400 pb-2 mb-2 text-lg text-gray-700 font-mono">{{ currentDate }}</div>
                      <p class="text-xs font-bold text-gray-400 uppercase tracking-widest">Date of Award</p>
                  </div>
              </div>
          </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps({
  isOpen: {
    type: Boolean,
    required: true
  },
  donorName: {
    type: String,
    required: true
  },
  badge: {
    type: String,
    required: true
  },
  totalDonations: {
    type: Number,
    required: true
  }
});

const emit = defineEmits(['close']);

const close = () => {
    emit('close');
};

const handlePrint = () => {
    window.print();
};

const displayName = computed(() => {
    return props.donorName || 'Honored Donor';
});

const currentDate = computed(() => {
    return new Date().toLocaleDateString('en-US', { year: 'numeric', month: 'long', day: 'numeric' });
});
</script>

<style scoped>

/* Inject a serif font specifically for the certificate */
@import url('https://fonts.googleapis.com/css2?family=Playfair+Display:ital,wght@0,400;0,700;1,400&display=swap');

.typography-serif {
    font-family: 'Playfair Display', serif;
}

@media print {
    body * {
        visibility: hidden;
    }
    #print-area, #print-area * {
        visibility: visible;
    }
    #print-area {
        position: absolute;
        left: 0;
        top: 0;
        width: 100%;
        -webkit-print-color-adjust: exact !important;
        print-color-adjust: exact !important;
    }
}
</style>
