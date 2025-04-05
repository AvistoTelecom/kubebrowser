<script setup lang="ts">
import YAML from 'yaml'

import type { Kubeconfig } from '@/types/Kubeconfig'

const props = defineProps<{
  kubeconfigs: Kubeconfig[]
  indexSelected: number | null
}>()

const emit = defineEmits<{
  (e: 'kubeconfig-selected', kubeconfig: string, index: number): void
}>()

function selectKubeconfig(kubeconfig: object, index: number) {
  emit('kubeconfig-selected', YAML.stringify(kubeconfig), index)
}
</script>

<template>
  <div
    class="flex-none flex flex-col gap-4 w-full"
  >
    <button
      v-for="(kubeconfig, index) in kubeconfigs"
      :key="index"
      class="text-lg py-6 px-12 rounded-md border-2 border-gray-600 cursor-pointer break-words"
      :class="{
        'bg-accent text-primary-950': props.indexSelected === index,
        'bg-gray-700': props.indexSelected !== index,
      }"
      @click="selectKubeconfig(kubeconfig.kubeconfig, index)"
    >
      {{ kubeconfig.name }}
    </button>
  </div>
</template>
