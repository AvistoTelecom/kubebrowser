<script setup lang="ts">
import AppLogoFull from '@/components/AppLogoFull.vue'
import IconCog from '@/icons/IconCog.vue'
import IconGithub from '@/icons/IconGithub.vue'
import IconHelp from '@/icons/IconHelp.vue'
import { RouteName } from '@/router'

const helpURL = window._env_.HELP_PAGE
</script>

<template>
  <header
    class="flex items-center content-center justify-between px-4 py-4 border-b border-gray-600 md:px-8"
  >
    <RouterLink :to="{ name: RouteName.Home }">
      <AppLogoFull class="w-40 md:w-48" />
    </RouterLink>

    <div class="flex flex-nowrap gap-3 items-center">
      <RouterLink
        :to="{ name: RouteName.Settings }"
        class="flex flex-nowrap items-center gap-2 px-3 py-2 border transition duration-500 rounded-xl text-gray-400 hover:text-white"
      >
        <IconCog class="shrink-0 w-6 h-6" />
        <span>Admin settings</span>
      </RouterLink>

      <a
        v-if="helpURL"
        :href="helpURL"
        target="_blank"
        rel="noopener noreferrer"
        title="Open documentation"
        class="hover:text-white text-gray-400 w-6 h-6 md:w-8 md:h-8 transition duration-500"
      >
        <IconHelp class="w-full h-full" />
      </a>

      <a
        href="https://github.com/AvistoTelecom/kubebrowser"
        target="_blank"
        rel="noopener noreferrer"
        title="Open Github"
        class="hover:text-white text-gray-400 w-6 h-6 md:w-8 md:h-8 transition duration-500"
      >
        <IconGithub class="w-full h-full" />
      </a>
    </div>
  </header>
</template>
