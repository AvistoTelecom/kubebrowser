<script setup lang="ts">
import type { Kubeconfig } from '@/types/Kubeconfig'

const props = defineProps<{
  kubeconfigs: Kubeconfig[]
  selected: Kubeconfig | null
}>()

const emit = defineEmits<{
  (e: 'update:selected', value: Kubeconfig): void
}>()
</script>

<template>
  <div
    class="flex-none flex flex-col gap-4 w-full"
  >
    <button
      v-for="kubeconfig in props.kubeconfigs"
      :key="kubeconfig.name"
      class="text-lg py-6 px-12 rounded-md border-2 border-gray-600 cursor-pointer break-words"
      :class="props.selected && props.selected.name === kubeconfig.name ? 'bg-accent text-primary-950' : 'bg-gray-700'"
      @click="emit('update:selected', kubeconfig)"
    >
      {{ kubeconfig.name }}
    </button>
  </div>
</template>
