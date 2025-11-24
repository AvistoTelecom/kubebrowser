<script setup lang="ts">
import { computed } from 'vue'
import YAML from 'yaml'

import Prism from 'prismjs'
import 'prismjs/components/prism-yaml'

import type { Kubeconfig } from '@/types/Kubeconfig'

import CopyButton from '@/components/CopyButton.vue'

const props = defineProps<{
  kubeconfig: Kubeconfig | null
  catalogLength: number
}>()

const kubeconfigAsYaml = computed(
  () => props.kubeconfig && YAML.stringify(props.kubeconfig.kubeconfig),
)

// Helper to deeply redact sensitive fields in the kubeconfig
function redactKubeconfig(k: Kubeconfig['kubeconfig']) {
  if (!k?.users) return k
  const clone = JSON.parse(JSON.stringify(k))
  clone.users.forEach(
    (u: {
      user?: {
        'auth-provider'?: {
          config?: Record<string, unknown>
        }
      }
    }) => {
      const cfg = u?.user?.['auth-provider']?.config
      if (cfg)
        (['client-secret', 'id-token', 'refresh-token'] as const).forEach((key) => {
          if (key in cfg) (cfg as Record<typeof key, unknown>)[key] = 'REDACTED'
        })
    },
  )
  return clone
}

const highlightedYaml = computed(() =>
  redactedKubeconfigAsYaml.value
    ? Prism.highlight(redactedKubeconfigAsYaml.value, Prism.languages.yaml, 'yaml')
    : '',
)

const redactedKubeconfigAsYaml = computed(() =>
  props.kubeconfig ? YAML.stringify(redactKubeconfig(props.kubeconfig.kubeconfig)) : null,
)
</script>

<template>
  <div
    class="p-4 overflow-auto border border-gray-100 bg-blue-950 rounded-xl"
    :class="{
      'flex items-center justify-center': !props.kubeconfig
    }"
  >
    <div v-if="redactedKubeconfigAsYaml && kubeconfigAsYaml">
      <CopyButton class="absolute z-10 text-gray-800 top-6 right-6 bg-accent-500 hover:bg-accent-600" :content="kubeconfigAsYaml" />
      <pre class="select-none language-yaml" v-html="highlightedYaml" />
    </div>
    <div v-else>
      <div v-if="!props.catalogLength">
        <p class="text-center text-gray-400">No results</p>
      </div>
      <div v-else>
        <p class="text-center text-gray-400">Select a cluster to display kubeconfig content</p>
      </div>
    </div>
  </div>
</template>
