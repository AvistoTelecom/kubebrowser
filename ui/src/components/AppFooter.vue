<script setup lang="ts">
import { DotLottieVue } from '@lottiefiles/dotlottie-vue'
import heartAnimation from '@/assets/heart.lottie'

const version = import.meta.env.VITE_VERSION
</script>

<template>
  <div class="flex flex-col items-center w-full mb-6 text-gray-300 group">
    <span
      class="flex items-center duration-500 ease-in-out delay-100 opacity-0 transition-10 group-hover:translate-y-2 group-hover:opacity-100"
      >Made with <DotLottieVue class="w-6 h-6" autoplay loop :src="heartAnimation" /> by Avisto
      Telecom</span
    >
    <p class="transition duration-700 ease-in-out -translate-y-4 group-hover:translate-y-2">
      Version: {{ version }}
    </p>
  </div>
</template>
