<script setup lang="ts">
import { RouterView } from 'vue-router'

import AppHeader from '@/components/AppHeader.vue'
import AppFooter from '@/components/AppFooter.vue'
</script>

<template>
  <div class="flex flex-col h-screen gap-y-4 bg-primary-300">
    <AppHeader />
    <RouterView />
    <AppFooter />
  </div>
</template>
