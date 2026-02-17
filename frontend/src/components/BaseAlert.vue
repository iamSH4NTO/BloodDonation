<template>
  <div 
    class="rounded-2xl p-4 mb-6 border transition-all duration-300"
    :class="containerStyles"
  >
    <div class="flex items-start gap-4">
      <div class="shrink-0 pt-0.5">
        <span class="material-icons" :class="iconStyles">{{ iconName }}</span>
      </div>
      <div class="flex-1">
        <h3 v-if="title" class="text-sm font-bold mb-1" :class="titleStyles">
          {{ title }}
        </h3>
        <div class="text-sm font-medium leading-relaxed" :class="contentStyles">
          <slot>{{ message }}</slot>
        </div>
      </div>
      <button 
        v-if="dismissible" 
        @click="dismiss" 
        class="shrink-0 p-1 rounded-lg hover:bg-black/5 transition-colors"
      >
        <span class="material-icons text-lg text-gray-400">close</span>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';

interface Props {
  type?: 'info' | 'warning' | 'error' | 'success';
  title?: string;
  message?: string;
  dismissible?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  type: 'info',
  dismissible: false
});

const emit = defineEmits(['close']);

const dismiss = () => {
  emit('close');
};

const containerStyles = computed(() => {
  switch (props.type) {
    case 'success': return 'bg-green-50/50 border-green-100';
    case 'error': return 'bg-red-50/50 border-red-100';
    case 'warning': return 'bg-amber-50/50 border-amber-100';
    default: return 'bg-blue-50/50 border-blue-100';
  }
});

const iconStyles = computed(() => {
  switch (props.type) {
    case 'success': return 'text-green-500';
    case 'error': return 'text-red-500';
    case 'warning': return 'text-amber-500';
    default: return 'text-blue-500';
  }
});

const iconName = computed(() => {
  switch (props.type) {
    case 'success': return 'check_circle';
    case 'error': return 'error';
    case 'warning': return 'warning_amber';
    default: return 'info';
  }
});

const titleStyles = computed(() => {
  switch (props.type) {
    case 'success': return 'text-green-900';
    case 'error': return 'text-red-900';
    case 'warning': return 'text-amber-900';
    default: return 'text-blue-900';
  }
});

const contentStyles = computed(() => {
  switch (props.type) {
    case 'success': return 'text-green-700';
    case 'error': return 'text-red-700';
    case 'warning': return 'text-amber-700';
    default: return 'text-blue-700';
  }
});
</script>
