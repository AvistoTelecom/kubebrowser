<script setup lang="ts">
import { ref } from 'vue'
import YAML from 'yaml'
import type { Kubeconfig } from '../types/Kubeconfig'

defineProps<{
  kubeconfigs: Kubeconfig[]
}>()

const emit = defineEmits<{
  (e: 'kubeconfig-selected', kubeconfig: string): void
}>()

const selectedKubeconfigName = ref<string | null>(null)

function selectKubeconfig(kubeconfig: object, name: string) {
  selectedKubeconfigName.value = name
  emit('kubeconfig-selected', YAML.stringify(kubeconfig))
}
</script>

<template>
  <div
    class="flex-none gap-8 max-w-max min-w-min"
    :class="{ 'flex flex-col': kubeconfigs.length < 3, 'grid grid-cols-2': kubeconfigs.length > 2 }"
  >
    <button
      v-for="kubeconfig in kubeconfigs"
      :key="kubeconfig.name"
      class="text-lg py-8 px-12 rounded-md border-2 border-gray-600 cursor-pointer"
      :class="{
        'bg-accent text-primary-950': selectedKubeconfigName === kubeconfig.name,
        'bg-gray-600': selectedKubeconfigName !== kubeconfig.name,
      }"
      @click="selectKubeconfig(kubeconfig.kubeconfig, kubeconfig.name)"
    >
      {{ kubeconfig.name }}
    </button>
  </div>
</template>
