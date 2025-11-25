<script setup lang="ts">
import type { Kubeconfig } from '@/types/Kubeconfig'
import IconEmojiSad from '@/icons/IconEmojiSad.vue'

const props = defineProps<{
  kubeconfigs: Kubeconfig[]
}>()

const model = defineModel<Kubeconfig | null>()
</script>

<template>
  <div class="flex flex-col flex-none w-full gap-3">
    <button
      v-for="kubeconfig in props.kubeconfigs"
      :key="kubeconfig.name"
      class="px-4 py-8 text-lg break-words border cursor-pointer rounded-xl"
      :class="
        model && model.name === kubeconfig.name
          ? 'bg-accent-500 text-blue-900'
          : 'border-gray-100 hover:bg-gray-050'
      "
      @click="model = kubeconfig"
    >
      {{ kubeconfig.name }}
    </button>
    <div
      v-if="!kubeconfigs.length"
      class="flex flex-col items-center justify-center flex-1 gap-1.5"
    >
      <IconEmojiSad class="w-10 h-10 text-gray-600 shrink-0" />
      <p class="text-gray-400">No results found</p>
    </div>
  </div>
</template>
