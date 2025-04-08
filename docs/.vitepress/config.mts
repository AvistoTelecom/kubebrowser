import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Kubebrowser",
  description: "Your Kubernetes catalog with OIDC",
  base: "/temp-kubebrowser/",
  ignoreDeadLinks: [
    /^https?:\/\/localhost/,
  ],
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Docs', link: '/' },
    ],
    sidebar: [
      {
        items: [
          { text: 'Getting started', link: '/' },
          { text: 'Contribute', link: '/contribute' }
        ]
      }
    ],
    search: {
      provider: 'local'
    },
    socialLinks: [
      { icon: 'github', link: 'https://github.com/vuejs/vitepress' }
    ]
  }
})
