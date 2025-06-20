import defineVersionedConfig from 'vitepress-versioning-plugin'
import { configureDiagramsPlugin } from "vitepress-plugin-diagrams";

// https://vitepress.dev/reference/site-config
export default defineVersionedConfig({
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
    logo: {
      dark: '/logo_light.svg',
      light: '/logo_dark.svg'
    },
    sidebar: {
      "/": [
        { text: 'What is Kubebrowser ?', link: '/' },
        { text: 'Getting started', link: '/getting-started' },
        { text: 'Expose Kubebrowser', link: '/expose' },
        { text: 'How it works', link: '/how-it-works' },
        { text: 'Contribute', link: '/contribute' },
        { text: 'Kubeconfig CRD', link: '/reference/kubeconfig'}
      ]
    },
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
  lastUpdated: true,
  versioning: {
    latestVersion: '0.8'
  }
}, __dirname)
