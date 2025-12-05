<script setup lang="ts">
import { onMounted, ref } from 'vue'

import * as api from '@/api/requests'

const username = ref('')
const loading = ref(true)

const base = new URL('@/assets/victory-hand-emoji.webp', import.meta.url).href
const helloEmoji = ref(base)

async function loadGetMe() {
  username.value = await api.getMe()
  loading.value = false
}

onMounted(async () => {
  await loadGetMe()
})

function restart() {
  // Force reload by adding a cache-busting param
  helloEmoji.value = base + '?t=' + Date.now()
}
</script>

<template>
  <div class="py-2 text-xl font-extrabold md:py-4 md:text-2xl 2xl:py-8 2xl:text-3xl">
    <div v-if="loading && !username" class="text-gray-300">Loading user name...</div>
    <div class="flex items-center gap-1" v-else>
      <img @mouseenter="restart" class="h-12" :src="helloEmoji" /> Hello {{ username }} !
    </div>
  </div>
</template>
