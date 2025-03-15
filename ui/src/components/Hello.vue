<script setup lang="ts">
import axios from 'axios'
import { ref, onMounted } from 'vue'

const username = ref('')

const loading = ref(true)

onMounted(async () => {
  if (import.meta.env.DEV) {
    // Mock response for development
    username.value = 'Firstname Lastname'
    loading.value = false
  } else {
    const response = await axios.get<string>('/api/me')
    username.value = response.data
    loading.value = false
  }
})
</script>

<template>
  <span v-if="!loading" class="py-8 font-extrabold text-3xl">✌️ Hello {{ username }} !</span>
</template>
