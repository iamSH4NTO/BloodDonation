<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  src?: string
  gender?: string
  name?: string
  size?: 'sm' | 'md' | 'lg' | 'xl'
}>()

const sizeClasses = {
  sm: 'w-8 h-8',
  md: 'w-12 h-12',
  lg: 'w-24 h-24',
  xl: 'w-32 h-32'
}

const avatarSrc = computed(() => {
  if (props.src) {
    // If it's already a full URL, return it
    if (props.src.startsWith('http')) return props.src
    
    // Construct base URL from env or fallback
    const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:4000/api/v1'
    const baseUrl = apiUrl.split('/api/v1')[0]
    
    // Ensure we don't have double slashes
    const cleanSrc = props.src.startsWith('/') ? props.src.slice(1) : props.src
    return `${baseUrl}/${cleanSrc}`
  }

  // Fallback to defaults
  if (props.gender?.toLowerCase() === 'female') {
    return '/avatars/default_female.png'
  }
  return '/avatars/default_male.png'
})

const initials = computed(() => {
  if (!props.name) return '?'
  return props.name
    .split(' ')
    .map(n => n[0])
    .join('')
    .toUpperCase()
    .slice(0, 2)
})
</script>

<template>
  <div :class="[sizeClasses[size || 'md'], 'relative shrink-0 rounded-full']">
    <div class="w-full h-full rounded-full overflow-hidden border-2 border-primary/10 bg-gray-100 flex items-center justify-center">
      <img 
        v-if="avatarSrc"
        :src="avatarSrc" 
        :alt="name || 'User Avatar'"
        class="w-full h-full object-cover"
        @error="(e) => (e.target as HTMLImageElement).style.display = 'none'"
      />
      <span v-else class="text-gray-400 font-medium" :class="size === 'xl' ? 'text-3xl' : 'text-lg'">
        {{ initials }}
      </span>
    </div>
  </div>
</template>
