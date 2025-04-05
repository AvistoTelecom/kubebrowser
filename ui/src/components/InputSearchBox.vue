<script setup lang="ts">
import { AnOutlinedSearch } from '@kalimahapps/vue-icons'

const props = withDefaults(
  defineProps<{
    modelValue: string
    placeholder: string
  }>(),
  {
    placeholder: 'Search...',
  },
)

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void
}>()

function updateValue(event: Event) {
  emit('update:modelValue', (event.target as HTMLInputElement).value)
}
</script>

<template>
  <div class="relative flex items-center">
    <AnOutlinedSearch class="absolute text-gray-400 pointer-events-none left-3" />
    <input
      :value="props.modelValue"
      @input="updateValue"
      type="search"
      :placeholder="placeholder"
      class="w-full py-2 pl-10 bg-gray-800 border border-gray-600 rounded-md focus:outline-none focus:border-accent"
    />
  </div>
</template>
