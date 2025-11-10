<script setup lang="ts">
import { RouterView } from 'vue-router'

import AppHeader from '@/components/AppHeader.vue'
import AppFooter from '@/components/AppFooter.vue'
</script>

<template>
  <div class="flex flex-col h-screen bg-blue-900 gap-y-4">
    <AppHeader />
    <RouterView />
    <AppFooter />
  </div>
</template>
