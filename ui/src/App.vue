<script setup lang="ts">
import AppHeader from '@/components/AppHeader.vue'
import AppFooter from '@/components/AppFooter.vue'
import PageHome from '@/components/PageHome.vue'
</script>

<template>
  <div class="flex flex-col h-screen gap-y-4 bg-primary-300">
    <AppHeader />
    <PageHome />
    <AppFooter />
  </div>
</template>
