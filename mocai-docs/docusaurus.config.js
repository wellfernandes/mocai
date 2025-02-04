// @ts-check
// `@type` JSDoc annotations allow editor autocompletion and type checking
// (when paired with `@ts-check`).
// There are various equivalent ways to declare your Docusaurus config.
// See: https://docusaurus.io/docs/api/docusaurus-config

import {themes as prismThemes} from 'prism-react-renderer';

// This runs in Node.js - Don't use client-side code here (browser APIs, JSX...)

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'Mocaí',
  tagline: 'Generating mocks, simplifying tests, accelerating development.',
  favicon: 'img/favicon.ico',

  // Set the production url of your site here
  url: 'https://github.com',
  // Set the /<baseUrl>/ pathname under which your site is served
  // For GitHub pages deployment, it is often '/<projectName>/'
  baseUrl: '/mocai/',

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: 'brazzcore', // Usually your GitHub org/user name.
  projectName: 'mocai', // Usually your repo name.
  deploymentBranch: 'gh-pages', // Branch for GitHub Pages deployment

  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',

  // Even if you don't use internationalization, you can use this field to set
  // useful metadata like html lang. For example, if your site is Chinese, you
  // may want to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          sidebarPath: './sidebars.js',
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            'https://github.com/brazzcore/mocai/tree/gh-pages',
        },
        theme: {
          customCss: './src/css/custom.css',
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      // Replace with your project's social card
      image: 'img/docusaurus-social-card.jpg',
      navbar: {
        title: 'Mocaí',
        logo: {
          alt: 'Mocaí Logo',
          src: 'img/logo.svg',
        },
        items: [
          {
            type: 'docSidebar',
            sidebarId: 'tutorialSidebar',
            position: 'left',
            label: 'Tutorial',
          },
          {
            href: 'https://github.com/brazzcore/mocai',
            label: 'GitHub',
            position: 'right',
          },
        ],
      },
      footer: {
        style: 'dark',
        links: [
          {
            title: 'Docs',
            items: [
              {
                label: 'Tutorial',
                to: '/docs/intro',
              },
              {
                label: 'Tasks',
                href: 'https://tree.taiga.io/project/wellfernandes-mocai/kanban',
              },
            ],
          },
          {
            title: 'Community',
            items: [
              {
                label: 'Brazzcore',
                href: 'https://github.com/brazzcore',
              },
              {
                label: 'Discord',
                href: 'https://discord.gg/TFRnQBkAMt',
              },
            ],
          },
          {
            title: 'More',
            items: [
              {
                label: 'Linkedin',
                href: 'https://www.linkedin.com/company/brazzcore',
              },
              {
                label: 'GitHub',
                href: 'https://github.com/brazzcore/mocai',
              },
            ],
          },
        ],
        copyright: `Copyright © ${new Date().getFullYear()} Brazzcore | Mocaí - developed by the open source community.`,
      },
      prism: {
        theme: prismThemes.github,
        darkTheme: prismThemes.dracula,
      },
    }),
};

export default config;
