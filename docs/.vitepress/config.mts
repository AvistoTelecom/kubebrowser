import { defineConfig } from "vitepress";
import { configureDiagramsPlugin } from "vitepress-plugin-diagrams";

// https://vitepress.dev/reference/site-config
export default defineConfig({
  transformHead({ assets }) {
    const themeFont = assets.find((file) => /Outfit\.[\w-]+\.ttf/.test(file));
    if (themeFont) {
      return [
        [
          "link",
          {
            rel: "preload",
            href: themeFont,
            as: "font",
            type: "font/ttf",
            crossorigin: "",
          },
        ],
      ];
    }
  },
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
  head: [["link", { rel: "icon", href: "/kubebrowser/favicon.ico" }]],
  ignoreDeadLinks: [/^https?:\/\/localhost/],
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    logo: {
      dark: "/logo_light.svg",
      light: "/logo_dark.svg",
    },
    nav: [
      {
        text: "Versions",
        items: [
          {
            text: "0.8 (latest)",
            link: "/",
          },
          {
            text: "0.7",
            link: "/0.7",
          },
        ],
      },
    ],
    sidebar: {
      "/": [
        {
          text: "Version 0.8 (latest)",
          items: [
            { text: "What is Kubebrowser ?", link: "/" },
            { text: "Getting started", link: "/getting-started" },
            { text: "Expose Kubebrowser", link: "/expose" },
            { text: "How it works", link: "/how-it-works" },
            { text: "Contribute", link: "/contribute" },
            { text: "Acknowledgments", link: "/acknowledgments" },
          ],
        },
        {
          text: "Reference",
          items: [{ text: "Kubeconfig CRD", link: "/reference/kubeconfig" }],
          collapsed: true,
        },
      ],
      "/0.7/": [
        {
          text: "Version 0.7",
          items: [{ text: "What is Kubebrowser ?", link: "/0.7" }],
        },
      ],
    },
    editLink: {
      pattern:
        "https://github.com/AvistoTelecom/kubebrowser/edit/main/docs/:path",
      text: "Edit this page on GitHub",
    },
    search: {
      provider: "local",
    },
    socialLinks: [
      { icon: "github", link: "https://github.com/AvistoTelecom/kubebrowser" },
    ],
  },
  sitemap: {
    hostname: "https://avistotelecom.github.io/kubebrowser",
  },
  lastUpdated: true,
});
