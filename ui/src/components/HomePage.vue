<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import axios from 'axios'
import type { Kubeconfig } from '../types/Kubeconfig'
import KubeconfigCatalog from './KubeconfigCatalog.vue'
import KubeconfigDisplay from './KubeconfigDisplay.vue'
import HelloComponent from './HelloComponent.vue'
import SearchBox from './SearchBox.vue'

const kubeconfigs = ref<Kubeconfig[]>([])
const searchQuery = ref<string>('')
const selectedKubeconfig = ref<string | null>(null)

const filteredKubeconfigs = computed(() => {
  if (!searchQuery.value) return kubeconfigs.value
  const query = searchQuery.value.toLowerCase()
  return kubeconfigs.value.filter((kubeconfig) => kubeconfig.name.toLowerCase().includes(query))
})

function updateSelectedKubeconfig(kubeconfig: string) {
  selectedKubeconfig.value = kubeconfig
}

onMounted(async () => {
  if (import.meta.env.DEV) {
    // Mock response for development
    kubeconfigs.value = [
      { name: 'Cluster 1', kubeconfig: { apiVersion: 'v1', kind: 'Config' } },
      { name: 'Cluster 2', kubeconfig: { apiVersion: 'v1', kind: 'Config2' } },
      { name: 'Another', kubeconfig: { apiVersion: 'v1', kind: 'Another' } },
    ]
  } else {
    const response = await axios.get<Kubeconfig[]>('/api/kubeconfigs')
    kubeconfigs.value = response.data
  }
})
</script>

<template>
  <HelloComponent />
  <div class="my-10">
    <SearchBox v-model="searchQuery" placeholder="Search kubeconfigs..." />
    <div class="flex space-x-8 my-8">
      <KubeconfigCatalog
        :kubeconfigs="filteredKubeconfigs"
        @kubeconfig-selected="updateSelectedKubeconfig"
      />
      <KubeconfigDisplay :yaml="selectedKubeconfig" />
    </div>
  </div>
</template>
