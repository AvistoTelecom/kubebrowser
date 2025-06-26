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
const loading = ref(true)
const selectedKubeconfig = ref<Kubeconfig | null>(null)

const filteredKubeconfigs = computed(() => {
  if (!searchQuery.value) return kubeconfigs.value
  const query = searchQuery.value.toLowerCase()
  const filtered = kubeconfigs.value.filter((kubeconfig) =>
    kubeconfig.name.toLowerCase().includes(query),
  )
  return filtered
})

async function loadConfigs() {
  kubeconfigs.value = await api.getConfigs()
  loading.value = false
}

onMounted(async () => {
  await loadConfigs()
})
</script>

<template>
  <AppHello class="mx-4 md:mx-8" />

  <div v-if="loading" class="flex items-center justify-center flex-1 gap-4 text-gray-300">
    Loading Kubeconfigs...
  </div>
  <div
    v-else-if="!loading && !kubeconfigs.length"
    class="flex flex-col items-center justify-center flex-1 gap-4"
  >
    <BsEmojiSurpriseFill class="w-10 h-10 text-gray-600" />
    <p class="text-gray-300">oops, it seems like you don't have acces to any clusters</p>
  </div>
  <div v-else class="relative flex flex-1 gap-4 mx-4 overflow-y-hidden md:mx-8">
    <div class="flex flex-col w-full gap-4 md:w-1/3 xl:w-1/5 2xl:w-1/6">
      <InputSearchBox v-model="searchQuery" placeholder="Search clusters..." class="md:pr-2" />
      <div class="overflow-y-auto">
        <KubeconfigCatalog :kubeconfigs="filteredKubeconfigs" v-model="selectedKubeconfig" />
      </div>
    </div>

    <KubeconfigDisplay
      class="w-full md:w-2/3 xl:w-4/5 2xl:w-5/6"
      :kubeconfig="selectedKubeconfig"
      :catalog-length="filteredKubeconfigs.length"
    />
  </div>
</template>
