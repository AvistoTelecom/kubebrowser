import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "KubeBrowser",
  description: "Your Kubernetes catalog with OIDC",
  base: "/kubebrowser/",
  ignoreDeadLinks: [
    /^https?:\/\/localhost/,
  ],
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    logo: {
      dark: '/logo_light.svg',
      light: '/logo_dark.svg'
    },
    sidebar: [
      {
        items: [
          { text: 'What is Kubebrowser ?', link: '/' },
          { text: 'Getting started', link: '/getting-started' },
          { text: 'Contribute', link: '/contribute' }
        ]
      }
    ],
    editLink: {
      pattern: 'https://github.com/AvistoTelecom/kubebrowser/edit/main/docs/:path',
      text: 'Edit this page on GitHub'
    },
    search: {
      provider: 'local'
    },
    socialLinks: [
      { icon: 'github', link: 'https://github.com/AvistoTelecom/kubebrowser' }
    ]
  },
    sitemap: {
      hostname: 'https://avistotelecom.github.io/kubebrowser'
  },
  lastUpdated: true
})
