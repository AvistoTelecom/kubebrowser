<script setup lang="ts">
import { ref, watch } from 'vue'
import { copyToClipboard } from '@/utils/clipboard'

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
    class="w-full min-h-5/6 border-2 border-gray-600 rounded-md p-4"
    :class="{ 'flex items-center justify-center': !yaml }"
  >
    <div v-if="yaml" class="relative">
      <button
        class="absolute top-1 right-1 p-3 bg-accent text-gray-800 rounded-tr-xl rounded-bl-xl"
        @click="handleCopy"
      >
        {{ copied ? 'Copied' : 'Copy' }}
      </button>
      <pre>{{ yaml }}</pre>
    </div>
    <div v-else>
      <p class="text-gray-300 text-center">Select a cluster to display kubeconfig content</p>
    </div>
  </div>
</template>
