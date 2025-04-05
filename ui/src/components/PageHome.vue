<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
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
  const query = searchQuery.value.toLowerCase()
  const filtered = kubeconfigs.value.filter((kubeconfig) =>
    kubeconfig.name.toLowerCase().includes(query),
  )
  return filtered
})

watch(searchQuery, () => {
  selectedKubeconfig.value = null
})

onMounted(async () => {
  loading.value = true
  kubeconfigs.value = await api.getConfigs()
  loading.value = false
})
</script>

<template>
  <AppHello class="mx-8" />

  <div v-if="loading" class="flex items-center justify-center flex-1 gap-4 text-gray-300">
    Loading Kubeconfigs...
  </div>
  <div
    v-else-if="!kubeconfigs.length"
    class="flex flex-col items-center justify-center flex-1 gap-4"
  >
    <BsEmojiSurpriseFill class="w-10 h-10 text-gray-600" />
    <p class="text-gray-300">oops, it seems like you don't have acces to any clusters</p>
  </div>
  <div v-else class="relative flex flex-1 mx-8 overflow-y-hidden gap-x-4">
    <div class="flex flex-col w-1/6 space-y-4">
      <InputSearchBox v-model="searchQuery" placeholder="Search clusters..." />
      <div class="overflow-y-auto">
        <KubeconfigCatalog
          :kubeconfigs="filteredKubeconfigs"
          v-model:selected="selectedKubeconfig"
        />
      </div>
    </div>

    <KubeconfigDisplay
      class="w-5/6"
      :kubeconfig="selectedKubeconfig"
      :catalog-length="filteredKubeconfigs.length"
    />
  </div>
</template>
