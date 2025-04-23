<script setup lang="ts">
import AppLogoFull from '@/components/AppLogoFull.vue'

import { BxSolidHelpCircle } from '@kalimahapps/vue-icons'

const helpURL = window._env_.HELP_PAGE
</script>

<template>
  <header class="flex items-center content-center justify-between p-8 border-b border-gray-600">
    <AppLogoFull class="w-48" />
    <a
      v-if="helpURL"
      :href="helpURL"
      target="_blank"
      rel="noopener noreferrer"
      title="Open documentation"
    >
      <BxSolidHelpCircle class="w-8 h-8 text-gray-300" />
    </a>
  </header>
</template>
