<script setup lang="ts">
import HomePage from './components/HomePage.vue'
import Header from './components/HeaderComponent.vue'
import Footer from './components/FooterComponent.vue'
</script>

<template>
  <div class="flex h-screen flex-col gap-y-4 bg-primary-300">
    <Header />
    <HomePage class="w-full" />
    <Footer />
  </div>
</template>
