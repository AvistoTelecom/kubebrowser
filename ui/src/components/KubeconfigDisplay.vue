<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { AkCopy } from '@kalimahapps/vue-icons'
import YAML from 'yaml'

import type { Kubeconfig } from '@/types/Kubeconfig'
import { copyToClipboard } from '@/utils/clipboard'

const props = defineProps<{
  kubeconfig: Kubeconfig | null
  catalogLength: number
}>()

const copied = ref(false)

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

const redactedKubeconfigAsYaml = computed(() =>
  props.kubeconfig ? YAML.stringify(redactKubeconfig(props.kubeconfig.kubeconfig)) : null,
)

const handleCopy = () => {
  if (kubeconfigAsYaml.value) {
    copyToClipboard(kubeconfigAsYaml.value)
    copied.value = true
  }
}

// Reset the "Copied" message when `yaml` changes
watch(
  () => props.kubeconfig,
  () => {
    copied.value = false
  },
)
</script>

<template>
  <div
    class="p-4 overflow-auto border-2 border-gray-600 rounded-md"
    :class="{
      'flex items-center justify-center': !props.kubeconfig,
      'bg-primary-950': props.kubeconfig,
    }"
  >
    <div v-if="redactedKubeconfigAsYaml">
      <div
        class="absolute inline-flex items-center justify-center gap-1 p-3 text-gray-800 cursor-pointer top-6 right-6 bg-accent min-w-min rounded-tr-md rounded-bl-md"
        @click="handleCopy"
      >
        <AkCopy />
        <span> {{ copied ? 'Copied' : 'Copy' }}</span>
      </div>
      <pre>{{ redactedKubeconfigAsYaml }}</pre>
    </div>
    <div v-else>
      <div v-if="!props.catalogLength">
        <p class="text-center text-gray-300">No results</p>
      </div>
      <div v-else>
        <p class="text-center text-gray-300">Select a cluster to display kubeconfig content</p>
      </div>
    </div>
  </div>
</template>
