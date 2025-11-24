<script setup lang="ts">
import type { Kubeconfig } from '@/types/Kubeconfig'

const props = defineProps<{
  kubeconfigs: Kubeconfig[]
}>()

const model = defineModel<Kubeconfig | null>()
</script>

<template>
  <div class="flex flex-col flex-none w-full gap-3">
    <button
      v-for="kubeconfig in props.kubeconfigs"
      :key="kubeconfig.name"
      class="px-4 py-8 text-lg break-words border cursor-pointer rounded-xl"
      :class="
        model && model.name === kubeconfig.name ? 'bg-accent-500 text-blue-900' : 'border-gray-100 hover:bg-gray-050'
      "
      @click="model = kubeconfig"
    >
      {{ kubeconfig.name }}
    </button>
  </div>
</template>
