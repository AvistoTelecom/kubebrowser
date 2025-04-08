<script setup lang="ts">
import { ref } from 'vue'

import * as api from '@/api/requests'

const username = ref('')

const loading = ref(false)
loadGetMe()

async function loadGetMe() {
  loading.value = true
  username.value = await api.getMe()
  loading.value = false
}
</script>

<template>
  <div class="py-8 text-3xl font-extrabold">
    <div v-if="loading" class="text-gray-300">Loading user name...</div>
    <div v-else>✌️ Hello {{ username }} !</div>
  </div>
</template>
