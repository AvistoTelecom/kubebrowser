<script setup lang="ts">
import { ref, watch } from 'vue'
import { copyToClipboard } from '@/utils/clipboard'
import { AkCopy } from '@kalimahapps/vue-icons'

const props = defineProps<{
  yaml: string | null
}>()

const copied = ref(false)

const handleCopy = () => {
  if (props.yaml) {
    copyToClipboard(props.yaml)
    copied.value = true
  }
}

// Reset the "Copied" message when `yaml` changes
watch(
  () => props.yaml,
  () => {
    copied.value = false
  },
)
</script>

<template>
  <div
    class="w-full max-w-full border-2 border-gray-600 rounded-md p-4 overflow-auto"
    :class="{ 'flex items-center justify-center': !yaml, 'bg-primary-950': yaml }"
  >
    <div v-if="yaml" class="relative">
      <div
        class="absolute top-1 right-1 inline-flex items-center justify-center gap-1 cursor-pointer p-3 bg-accent min-w-min text-gray-800 rounded-tr-xl rounded-bl-xl"
        @click="handleCopy"
      >
        <AkCopy />
        <span> {{ copied ? 'Copied' : 'Copy' }}</span>
      </div>
      <pre>{{ yaml }}</pre>
    </div>
    <div v-else>
      <p class="text-gray-300 text-center">Select a cluster to display kubeconfig content</p>
    </div>
  </div>
</template>
