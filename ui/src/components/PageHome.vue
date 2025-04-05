<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { BsEmojiSurpriseFill } from '@kalimahapps/vue-icons'

import type { Kubeconfig } from '@/types/Kubeconfig'
import * as api from '@/api/requests'

import AppHello from '@/components/AppHello.vue'
import InputSearchBox from '@/components/InputSearchBox.vue'
import KubeconfigCatalog from '@/components/KubeconfigCatalog.vue'
import KubeconfigDisplay from '@/components/KubeconfigDisplay.vue'

const kubeconfigs = ref<Kubeconfig[]>([])
const searchQuery = ref('')
const loading = ref(false)
const selectedKubeconfig = ref<Kubeconfig | null>(null)

const filteredKubeconfigs = computed(() => {
  if (!searchQuery.value) return kubeconfigs.value
  selectedKubeconfig.value = null
  const query = searchQuery.value.toLowerCase()
  const filtered = kubeconfigs.value.filter((kubeconfig) => kubeconfig.name.toLowerCase().includes(query))
  return filtered
})

onMounted(async () => {
  loading.value = true
  kubeconfigs.value = await api.getConfigs()
  loading.value = false
})
</script>

<template>
  <AppHello class="mx-8" />

  <div v-if="loading" class="flex flex-1 gap-4 items-center justify-center text-gray-300">Loading kube configs.</div>
  <div v-else-if="!kubeconfigs.length" class="flex flex-col flex-1 gap-4 items-center justify-center">
    <BsEmojiSurpriseFill class="w-10 h-10 text-gray-600"/>
    <p class="text-gray-300">oops, it seems like you don't have acces to any clusters</p>
  </div>
  <div v-else class="relative mx-8 flex flex-1 gap-x-4 overflow-y-hidden">
    <div class="space-y-4 w-1/6 flex flex-col">
      <InputSearchBox v-model="searchQuery" placeholder="Search clusters..." />
      <div class="overflow-y-auto">
        <KubeconfigCatalog
          :kubeconfigs="filteredKubeconfigs"
          v-model:selected="selectedKubeconfig"
        />
      </div>
    </div>

    <KubeconfigDisplay class="w-5/6" :kubeconfig="selectedKubeconfig" :catalog-length="filteredKubeconfigs.length" />
  </div>
</template>
