<script setup lang="ts">
import AppHeader from '@/components/AppHeader.vue'
import AppFooter from '@/components/AppFooter.vue'
import PageHome from '@/components/PageHome.vue'
</script>

<template>
  <div class="flex h-screen flex-col gap-y-4 bg-primary-300">
    <AppHeader />
    <PageHome class="w-full" />
    <AppFooter />
  </div>
</template>
