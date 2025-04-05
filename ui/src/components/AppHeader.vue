<script setup lang="ts">
import { BxSolidHelpCircle } from '@kalimahapps/vue-icons'

const helpURL = window._env_.HELP_PAGE
</script>

<template>
  <header class="flex content-center justify-between p-8 border-b border-gray-600">
    <h1 class="font-logo font-extrabold text-xl">KubeBrowser</h1>
    <a v-if="helpURL" :href="helpURL" target="_blank" rel="noopener noreferrer" title="Open documentation">
      <BxSolidHelpCircle class="w-8 h-8"/>
    </a>
  </header>
</template>
