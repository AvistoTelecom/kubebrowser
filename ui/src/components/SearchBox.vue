<script setup lang="ts">
import { withDefaults, defineProps, defineEmits } from 'vue'
import { AnOutlinedSearch } from '@kalimahapps/vue-icons'

withDefaults(
  defineProps<{
    modelValue: string
    placeholder?: string
  }>(),
  {
    placeholder: 'Search...',
  },
)

const emit = defineEmits(['update:modelValue'])

function updateValue(event: Event) {
  emit('update:modelValue', (event.target as HTMLInputElement).value)
}
</script>

<template>
  <div
    class="inline-flex items-center gap-1 p-3 rounded-md bg-gray-800 w-full border-2 border-gray-600"
  >
    <AnOutlinedSearch />
    <input
      :value="modelValue"
      @input="updateValue"
      type="text"
      :placeholder="placeholder"
      class=""
    />
  </div>
</template>
