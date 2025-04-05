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
  <span v-if="!loading" class="py-8 font-extrabold text-3xl">✌️ Hello {{ username }} !</span>
</template>
