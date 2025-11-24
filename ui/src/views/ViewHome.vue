<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { BsEmojiSurpriseFill } from '@kalimahapps/vue-icons'
import YAML from 'yaml'

import type { Kubeconfig } from '@/types/Kubeconfig'
import * as api from '@/api/requests'

import AppHello from '@/components/AppHello.vue'
import InputSearchBox from '@/components/InputSearchBox.vue'
import KubeconfigCatalog from '@/components/KubeconfigCatalog.vue'
import KubeconfigDisplay from '@/components/KubeconfigDisplay.vue'
import CopyButton from '@/components/CopyButton.vue'

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

const filteredKubeconfigsContent = computed(() => {
  // Create a merged Kubeconfig
  const mergedKubeconfig: Kubeconfig['kubeconfig'] = {
    apiVersion: 'v1',
    kind: 'Config',
    clusters: [],
    contexts: [],
    users: [],
    'current-context': '',
  }

  // Merge clusters, contexts, and users from all filtered kubeconfigs
  filteredKubeconfigs.value.forEach((kubeconfig) => {
    const kc = kubeconfig.kubeconfig

    // Append clusters
    if (kc.clusters?.length) {
      mergedKubeconfig.clusters.push(...kc.clusters)
    }

    // Append contexts
    if (kc.contexts?.length) {
      mergedKubeconfig.contexts!.push(...kc.contexts)
    }

    // Append users
    if (kc.users?.length) {
      mergedKubeconfig.users!.push(...kc.users)
    }

    // Optionally set current-context (only if not set yet)
    if (!mergedKubeconfig['current-context'] && kc['current-context']) {
      mergedKubeconfig['current-context'] = kc['current-context']
    }
  })

  // Optional: remove duplicates (by name)
  mergedKubeconfig.clusters = Object.values(
    Object.fromEntries(mergedKubeconfig.clusters.map((c) => [c.name, c])),
  )
  mergedKubeconfig.contexts = Object.values(
    Object.fromEntries(mergedKubeconfig.contexts!.map((c) => [c.name, c])),
  )
  mergedKubeconfig.users = Object.values(
    Object.fromEntries(mergedKubeconfig.users!.map((u) => [u.name, u])),
  )

  // Convert to YAML
  const yamlContent = YAML.stringify(mergedKubeconfig)
  return yamlContent
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

  <div v-if="loading" class="flex items-center justify-center flex-1 gap-4 text-gray-400">
    Loading Kubeconfigs...
  </div>
  <div
    v-else-if="!loading && !kubeconfigs.length"
    class="flex flex-col items-center justify-center flex-1 gap-4"
  >
    <BsEmojiSurpriseFill class="w-10 h-10 text-gray-600" />
    <p class="text-gray-400">oops, it seems like you don't have acces to any clusters</p>
  </div>
  <div v-else class="relative flex flex-1 mx-4 overflow-y-hidden md:mx-8">
    <div class="flex flex-col w-full gap-3 mr-6 md:w-1/3 xl:w-1/5 2xl:w-1/6">
      <div class="flex flex-col gap-3 pb-3" :class="{'border-gray-600 border-b-1' : filteredKubeconfigs.length}">
        <InputSearchBox v-model="searchQuery" placeholder="Search clusters..." />
        <CopyButton v-if="filteredKubeconfigs.length" class="w-full text-white border border-gray-100 hover:bg-gray-050" :content="filteredKubeconfigsContent" :text-before="'Copy all clusters'" light />
      </div>
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
