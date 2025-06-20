<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { DotLottieVue, type DotLottie, type DotLottieVueInstance } from '@lottiefiles/dotlottie-vue'
import YAML from 'yaml'

import Prism from 'prismjs'
import 'prismjs/components/prism-yaml'

import type { Kubeconfig } from '@/types/Kubeconfig'
import { copyToClipboard } from '@/utils/clipboard'

const props = defineProps<{
  kubeconfig: Kubeconfig | null
  catalogLength: number
}>()

const copyAnimationRef = ref<DotLottieVueInstance | null>(null)
const timeout = ref<ReturnType<typeof setTimeout> | null>(null)
const copied = ref(false)

const kubeconfigAsYaml = computed(
  () => props.kubeconfig && YAML.stringify(props.kubeconfig.kubeconfig),
)

function resetCopied(lottie: DotLottie | null | undefined) {
  if (lottie) {
    lottie.setSegment(1, 101)
    lottie.play()
  }
  copied.value = false
}

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

const handleCopy = () => {
  if (kubeconfigAsYaml.value) {
    copyToClipboard(kubeconfigAsYaml.value)
    copied.value = true
    // Access the underlying Lottie player instance and call play()
    const lottieInstance = copyAnimationRef.value?.getDotLottieInstance?.()
    lottieInstance?.setSegment(1, 59)
    lottieInstance?.play()
    // Reset after 2 seconds
    if (timeout.value) clearTimeout(timeout.value)
    timeout.value = setTimeout(() => resetCopied(lottieInstance), 2000)
  }
}

// Reset the "Copied" message when `yaml` changes
watch(
  () => props.kubeconfig,
  () => {
    copied.value = false
    copyAnimationRef.value?.getDotLottieInstance?.()?.stop()
    copyAnimationRef.value?.getDotLottieInstance?.()?.setFrame(1)
    if (timeout.value) clearTimeout(timeout.value)
    timeout.value = null
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
        class="absolute z-10 inline-flex items-center justify-center w-24 p-3 text-gray-800 duration-500 cursor-pointer transition-width top-6 right-6 bg-accent rounded-tr-md rounded-bl-md"
        :class="{
          'w-27': copied,
        }"
        @click="handleCopy"
      >
        <DotLottieVue
          class="-m-4 -mr-2 shrink-0 w-14"
          ref="copyAnimationRef"
          alt="Copy"
          src="/lottie/copy.lottie"
        />
        <span> {{ copied ? 'Copied' : 'Copy' }}</span>
      </div>
      <pre class="select-none language-yaml" v-html="highlightedYaml" />
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
