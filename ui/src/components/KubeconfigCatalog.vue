<script setup lang="ts">
import type { Kubeconfig } from '@/types/Kubeconfig'

const props = defineProps<{
  kubeconfigs: Kubeconfig[]
  indexSelected: number | null
}>()

const emit = defineEmits<{
  (e: 'kubeconfig-selected', kubeconfig: Kubeconfig, index: number): void
}>()

function selectKubeconfig(kubeconfig: Kubeconfig, index: number) {
  emit('kubeconfig-selected', kubeconfig, index)
}
</script>

<template>
  <div
    class="flex-none flex flex-col gap-4 w-full"
  >
    <button
      v-for="(kubeconfig, index) in props.kubeconfigs"
      :key="index"
      class="text-lg py-6 px-12 rounded-md border-2 border-gray-600 cursor-pointer break-words"
      :class="{
        'bg-accent text-primary-950': props.indexSelected === index,
        'bg-gray-700': props.indexSelected !== index,
      }"
      @click="selectKubeconfig(kubeconfig, index)"
    >
      {{ kubeconfig.name }}
    </button>
  </div>
</template>
