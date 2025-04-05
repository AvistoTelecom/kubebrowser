<script setup lang="ts">
import { ref, onMounted } from 'vue'

import * as api from '@/api/requests'

const username = ref('')

const loading = ref(true)

onMounted(async () => {
  username.value = await api.getMe()
  loading.value = false
})
</script>

<template>
  <div class="py-8 font-extrabold text-3xl">
    <div v-if="!loading">✌️ Hello {{ username }} !</div>
  </div>
</template>
