<script setup lang="ts">
import { onMounted, ref } from 'vue'

import * as api from '@/api/requests'

const username = ref('')

const loading = ref(true)

async function loadGetMe() {
  username.value = await api.getMe()
  loading.value = false
}

onMounted(async () => {
  await loadGetMe()
})
</script>

<template>
  <div class="py-2 text-xl font-extrabold md:py-4 md:text-2xl 2xl:py-8 2xl:text-3xl">
    <div v-if="loading && !username" class="text-gray-300">Loading user name...</div>
    <div v-else>✌️ Hello {{ username }} !</div>
  </div>
</template>
