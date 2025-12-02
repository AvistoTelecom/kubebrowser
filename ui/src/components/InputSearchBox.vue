<script setup lang="ts">
import IconSearch from '@/icons/IconSearch.vue'
import IconXMark from '@/icons/IconXMark.vue'

const props = withDefaults(
  defineProps<{
    placeholder: string
  }>(),
  {
    placeholder: 'Search...',
  },
)

const model = defineModel<string>()

const clearInput = () => {
  model.value = ''
}
</script>

<template>
  <div class="relative flex items-center">
    <IconSearch class="absolute w-5 h-5 text-white pointer-events-none left-4" />
    <input
      v-model="model"
      type="text"
      :placeholder="props.placeholder"
      class="w-full py-2 pl-10 border border-gray-100 outline-none placeholder:text-gray-400 bg-gray-050 rounded-xl focus:border-accent-500 hover:border-gray-300"
    />
    <IconXMark
      v-if="model"
      @mousedown.prevent
      @click="clearInput"
      class="absolute w-5 h-5 text-white cursor-pointer right-4"
    />
  </div>
</template>
