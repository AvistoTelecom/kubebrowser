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
  <div class="flex flex-col flex-none w-full gap-4 pr-2">
    <button
      v-for="kubeconfig in props.kubeconfigs"
      :key="kubeconfig.name"
      class="px-4 py-6 text-lg break-words border-2 border-gray-600 rounded-md cursor-pointer"
      :class="
        props.selected && props.selected.name === kubeconfig.name
          ? 'bg-accent text-primary-950'
          : 'bg-gray-700'
      "
      @click="emit('update:selected', kubeconfig)"
    >
      {{ kubeconfig.name }}
    </button>
  </div>
</template>
