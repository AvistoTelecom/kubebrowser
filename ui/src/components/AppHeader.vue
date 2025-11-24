<script setup lang="ts">
import AppLogoFull from '@/components/AppLogoFull.vue'
import { RouteName } from '@/router'
import VueIcon from '@kalimahapps/vue-icons/VueIcon'

const helpURL = window._env_.HELP_PAGE
</script>

<template>
  <header
    class="flex items-center content-center justify-between px-4 py-4 border-b border-gray-600 md:px-8"
  >
    <RouterLink :to="{ name: RouteName.Home }">
      <AppLogoFull class="w-40 md:w-48" />
    </RouterLink>
    <a
      v-if="helpURL"
      :href="helpURL"
      target="_blank"
      rel="noopener noreferrer"
      title="Open documentation"
    >
      <VueIcon name="bx:solid-help-circle" class="w-6 h-6 text-gray-400 md:w-8 md:h-8" />
    </a>
  </header>
</template>
