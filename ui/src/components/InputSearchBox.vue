<script setup lang="ts">
import { AnOutlinedSearch } from '@kalimahapps/vue-icons'

const props = withDefaults(
  defineProps<{
    placeholder: string
  }>(),
  {
    placeholder: 'Search...',
  },
)

const model = defineModel<string>()
</script>

<template>
  <div class="relative flex items-center">
    <AnOutlinedSearch class="absolute text-gray-400 pointer-events-none left-3" />
    <input
      v-model="model"
      type="search"
      :placeholder="props.placeholder"
      class="w-full py-2 pl-10 bg-gray-800 border border-gray-600 rounded-md focus:outline-hidden focus:border-accent"
    />
  </div>
</template>
