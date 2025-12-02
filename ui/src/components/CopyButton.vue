<script setup lang="ts">
import { ref, watch } from 'vue'
import { DotLottieVue, type DotLottie, type DotLottieVueInstance } from '@lottiefiles/dotlottie-vue'

import { copyToClipboard } from '@/utils/clipboard'
import copyAnimation from '@/assets/copy.lottie'
import copyLightAnimation from '@/assets/copy-light.lottie'

const props = withDefaults(
  defineProps<{
    content: string
    textBefore?: string
    textAfter?: string
    light?: boolean
  }>(),
  {
    textBefore: 'Copy',
    textAfter: 'Copied',
  },
)

const copyAnimationRef = ref<DotLottieVueInstance | null>(null)
const timeout = ref<ReturnType<typeof setTimeout> | null>(null)
const copied = ref(false)

function resetCopied(lottie: DotLottie | null | undefined) {
  if (lottie) {
    lottie.setSegment(1, 101)
    lottie.play()
  }
  copied.value = false
}

const handleCopy = () => {
  if (props.content) {
    copyToClipboard(props.content)
    copied.value = true
    // Access the underlying Lottie player instance and call play()
    const lottieInstance = copyAnimationRef.value?.getDotLottieInstance?.()
    lottieInstance?.setSegment(1, 59)
    lottieInstance?.play()
    // Reset after 2 seconds
    if (timeout.value) clearTimeout(timeout.value)
    timeout.value = setTimeout(() => resetCopied(lottieInstance), 2000)
  }
}

// Reset the "Copied" message when `content` changes
watch(
  () => props.content,
  () => {
    copied.value = false
    copyAnimationRef.value?.getDotLottieInstance?.()?.stop()
    copyAnimationRef.value?.getDotLottieInstance?.()?.setFrame(1)
    if (timeout.value) clearTimeout(timeout.value)
    timeout.value = null
  },
)
</script>

<template>
  <div
    class="inline-flex items-center justify-center w-24 p-3 transition-[width] duration-500 cursor-pointer rounded-xl transition-width"
    :class="{
      'w-27': copied,
    }"
    @click="handleCopy"
  >
    <DotLottieVue
      class="-m-4 -mr-2 shrink-0 w-14"
      ref="copyAnimationRef"
      alt="Copy"
      :src="props.light ? copyLightAnimation : copyAnimation"
    />
    <span> {{ copied ? props.textAfter : props.textBefore }}</span>
  </div>
</template>
