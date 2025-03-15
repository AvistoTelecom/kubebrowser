<script setup lang="ts">
defineProps<{
  yaml: string | null
}>()
</script>

<template>
  <div
    class="w-full min-h-5/6 border-2 border-gray-600 rounded-md p-4"
    :class="{ 'flex items-center justify-center': !yaml }"
  >
    <pre v-if="yaml">{{ yaml }}</pre>
    <p v-else class="text-gray-300 text-center">Select a cluster to display kubeconfig content</p>
  </div>
</template>
