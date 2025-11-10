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
  <div class="relative flex items-center text-gray-400">
    <AnOutlinedSearch class="absolute w-5 h-5 pointer-events-none left-4" />
    <input
      v-model="model"
      type="search"
      :placeholder="props.placeholder"
      class="w-full py-2 pl-10 border border-gray-100 bg-gray-050 rounded-xl focus:outline-hidden focus:border-accent"
    />
  </div>
</template>
