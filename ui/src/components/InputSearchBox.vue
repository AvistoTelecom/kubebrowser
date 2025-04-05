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
    <AnOutlinedSearch class="absolute left-3 text-gray-400 pointer-events-none" />
    <input
      :value="props.modelValue"
      @input="updateValue"
      type="search"
      :placeholder="placeholder"
      class="pl-10 py-2 border border-gray-600 bg-gray-800 rounded-md focus:outline-none focus:border-accent w-full"
    />
  </div>
</template>
