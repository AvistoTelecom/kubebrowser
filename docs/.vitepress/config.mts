import { defineConfig } from 'vitepress'
import { configureDiagramsPlugin } from "vitepress-plugin-diagrams";

// https://vitepress.dev/reference/site-config
export default defineConfig({
  markdown: {
    config: (md) => {
      configureDiagramsPlugin(md, {
        diagramsDir: "public/diagrams",
        publicPath: "/kubebrowser/diagrams",
      });
    },
  },
  title: "KubeBrowser",
  description: "Your Kubernetes catalog with OIDC",
  base: "/kubebrowser/",
  head: [['link', { rel: 'icon', href: '/kubebrowser/favicon.ico' }]],
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
          { text: 'How it works', link: '/how-it-works' },
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
