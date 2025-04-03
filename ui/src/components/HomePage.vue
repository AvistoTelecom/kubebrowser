<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import axios from 'axios'
import type { Kubeconfig } from '../types/Kubeconfig'
import KubeconfigCatalog from './KubeconfigCatalog.vue'
import KubeconfigDisplay from './KubeconfigDisplay.vue'
import HelloComponent from './HelloComponent.vue'
import SearchBox from './SearchBox.vue'
import { BsEmojiSurpriseFill } from '@kalimahapps/vue-icons';

const kubeconfigs = ref<Kubeconfig[]>([])
const searchQuery = ref('')
const selectedKubeconfig = ref<string | null>(null)
const indexSelected = ref<number | null>(null)
const emptyKubeconfigs = ref(false)

const filteredKubeconfigs = computed(() => {
  if (!searchQuery.value) return kubeconfigs.value
  selectedKubeconfig.value = null
  indexSelected.value = null
  const query = searchQuery.value.toLowerCase()
  const filtered = kubeconfigs.value.filter((kubeconfig) => kubeconfig.name.toLowerCase().includes(query))
  return filtered
})

function updateSelectedKubeconfig(kubeconfig: string, index: number) {
  indexSelected.value = index
  selectedKubeconfig.value = kubeconfig
}

onMounted(async () => {
  if (import.meta.env.DEV) {
    // Mock response for development
    kubeconfigs.value = [
      { name: 'Cluster number 1', kubeconfig: { apiVersion: 'v1', kind: 'Config' } },
      { name: 'Cluster number 2', kubeconfig: { apiVersion: 'v1', kind: 'Config2' } },
      { name: 'Cluster number 3', kubeconfig: { apiVersion: 'v1', kind: 'Config2' } },
      { name: 'Cluster number 4', kubeconfig: { apiVersion: 'v1', kind: 'Config2' } },
      { name: 'Cluster number 5', kubeconfig: { apiVersion: 'v1', kind: 'Config2' } },
      { name: 'Cluster number 6', kubeconfig: { apiVersion: 'v1', kind: 'Config2' } },
      { name: 'Cluster number 7', kubeconfig: { apiVersion: 'v1', kind: 'Config2' } },
      { name: 'Another cluster', kubeconfig: { apiVersion: 'v1', kind: 'Another' } },
    ]
  } else {
    const response = await axios.get<Kubeconfig[]>('/api/kubeconfigs')
    if (!response.data.length) {
      emptyKubeconfigs.value = true
    }
    kubeconfigs.value = response.data
  }
})
</script>

<template>
  <HelloComponent class="mx-8" />
  <div v-if="emptyKubeconfigs" class="flex flex-col flex-1 gap-4 items-center justify-center">
    <BsEmojiSurpriseFill class="w-10 h-10 text-gray-600"/>
    <p class="text-gray-300">oops, it seems like you don't have acces to any clusters</p>
  </div>
  <div v-else class="relative mx-8 flex flex-1 gap-x-4 overflow-y-hidden">
    <div class="space-y-4 w-1/6 flex flex-col">
      <SearchBox v-model="searchQuery" placeholder="Search clusters..." />
      <div class="overflow-y-auto">
        <KubeconfigCatalog
        :kubeconfigs="filteredKubeconfigs"
        :index-selected="indexSelected"
        @kubeconfig-selected="updateSelectedKubeconfig"
        />
      </div>

    </div>
    <KubeconfigDisplay class="w-5/6" :kubeconfig="selectedKubeconfig" :catalog-length="filteredKubeconfigs.length" />
  </div>
</template>
