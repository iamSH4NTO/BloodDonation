<template>
  <Transition
    enter-active-class="transform ease-out duration-300 transition"
    enter-from-class="translate-y-2 opacity-0 sm:translate-y-0 sm:translate-x-2"
    enter-to-class="translate-y-0 opacity-100 sm:translate-x-0"
    leave-active-class="transition ease-in duration-200"
    leave-from-class="opacity-100"
    leave-to-class="opacity-0"
  >
    <div v-if="toast.visible" class="fixed top-4 right-4 z-[9999] max-w-sm w-full bg-white shadow-2xl rounded-2xl pointer-events-auto border overflow-hidden" :class="borderClass">
      <div class="p-4 flex items-start gap-3">
        <div class="shrink-0">
          <span class="material-icons" :class="iconClass">{{ iconName }}</span>
        </div>
        <div class="flex-1 pt-0.5">
          <p class="text-sm font-bold text-gray-900 leading-snug">{{ messageTitle }}</p>
          <p class="mt-1 text-sm text-gray-500 font-medium leading-relaxed">{{ toast.message }}</p>
        </div>
        <button @click="toast.hide()" class="shrink-0 text-gray-400 hover:text-gray-600 transition-colors bg-gray-50 rounded-lg p-1">
          <span class="material-icons text-lg">close</span>
        </button>
      </div>
      <div class="h-1 w-full bg-gray-100 relative">
        <div class="h-full absolute left-0 top-0 transition-all duration-[3000ms] ease-linear" :class="progressClass" :style="{ width: progressWidth }"></div>
      </div>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import { computed, ref, onMounted, watch } from 'vue';
import { useToastStore } from '@/stores/toast';

const toast = useToastStore();
const progressWidth = ref('100%');

const borderClass = computed(() => {
  switch (toast.type) {
    case 'success': return 'border-green-100 shadow-green-100/50';
    case 'error': return 'border-red-100 shadow-red-100/50';
    default: return 'border-blue-100 shadow-blue-100/50';
  }
});

const iconClass = computed(() => {
  switch (toast.type) {
    case 'success': return 'text-green-500';
    case 'error': return 'text-red-500';
    default: return 'text-blue-500';
  }
});

const iconName = computed(() => {
  switch (toast.type) {
    case 'success': return 'check_circle';
    case 'error': return 'error';
    default: return 'info';
  }
});

const messageTitle = computed(() => {
  switch (toast.type) {
    case 'success': return 'Success';
    case 'error': return 'Error';
    default: return 'Notification';
  }
});

const progressClass = computed(() => {
  switch (toast.type) {
    case 'success': return 'bg-green-500';
    case 'error': return 'bg-red-500';
    default: return 'bg-blue-500';
  }
});

watch(() => toast.visible, (newVal) => {
  if (newVal) {
    progressWidth.value = '100%';
    setTimeout(() => {
      progressWidth.value = '0%';
    }, 10);
  }
});
</script>
