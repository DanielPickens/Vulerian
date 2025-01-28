const prismReactRenderer = require('prism-react-renderer');
const path = require('path');

/** @type {import('@docusaurus/types').DocusaurusConfig} */
module.exports = {
  title: 'vulerian',
  tagline: 'vulerian - Fast iterative container-based application development',
  url: 'https://vulerian.dev',
  baseUrl: '/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'throw',
  favicon: 'img/favicon.ico',
  organizationName: 'daniel-pickens', // Usually your GitHub org/user name.
  projectName: 'vulerian', // Usually your repo name.
  plugins: [
    [
      path.resolve(__dirname, 'docusaurus-vulerian-plugin-segment'),
      {
        apiKey: 'seYXMF0tyHs5WcPsaNXtSEmQk3FqzTz0',
        options: {
          context: { ip: '0.0.0.0' }
        }
      }
    ]
  ],
  themeConfig: {
    docs: {
      sidebar: {
        autoCollapseCategories: false
      },
    },
    announcementBar: {
      id: 'announcementBar-2', // Increment on change
      content: `⭐️ Love vulerian? Support us by giving it a star on <a target="_blank" rel="noopener noreferrer" href="https://github\.com/danielpickens/vulerian">GitHub</a>! ⭐️`,
    },
    navbar: {
      title: 'vulerian',
      logo: {
        alt: 'vulerian Logo',
        src: 'img/logo.png',
        srcDark: 'img/logo_dark.png',
      },
      items: [
        {
          type: 'doc',
          docId: 'introduction',
          position: 'left',
          label: 'Docs',
        },
        { to: '/blog', label: 'Blog', position: 'left' },
        {
          href: 'https://github\.com/danielpickens/vulerian',
          label: 'GitHub',
          position: 'right',
        },
        {
          type: 'docsVersionDropdown',
          position: 'right',
          dropdownActiveClassDisabled: true,
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Learn',
          items: [
            {
              label: 'Installation',
              to: 'docs/overview/installation'
            },
            {
              label: 'Quickstart',
              to: 'docs/user-guides/quickstart'
            },
          ]
        },
        {
          title: 'Community',
          items: [
            {
              label: '#vulerian on the Kubernetes Slack',
              href: 'https://slack.k8s.io/',
              external: true,
            },
            {
              label: 'Meetings',
              href: 'https://calendar.google.com/calendar/u/0/embed?src=gi0s0v5ukfqkjpnn26p6va3jfc@group.calendar.google.com',
            },
          ],
        },
        {
          title: 'More',
          items: [
            {
              label: 'Blog',
              to: 'blog',
            },
            {
              label: 'GitHub',
              href: 'https://github\.com/danielpickens/vulerian',
            },
            {
              label: 'Twitter',
              href: 'https://twitter.com/rhdevelopers',
            },
            {
              label: 'YouTube',
              href: 'https://www.youtube.com/channel/UCXAt2CtoBBtN9EWe4xv4Row'
            }
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} vulerian Authors -- All Rights Reserved <br> Apache License 2.0 open source project`,
    },
    prism: {
      theme: prismReactRenderer.themes.github,
      darkTheme: prismReactRenderer.themes.oceanicNext,
      additionalLanguages: ['docker'],
    },
    algolia: {
      appId: '7RBQSTPIA4',
      apiKey: '97ac94cb47dcaeef1c2c9694bd39b458',
      indexName: 'vulerian',
      debug: false
    }
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
          breadcrumbs: true,
          sidebarCollapsible: true,
          lastVersion: 'current',
          showLastUpdateTime: true,
          showLastUpdateAuthor: false,
          exclude: [
              '**/docs-mdx/**',
              '**/_*.{js,jsx,ts,tsx,md,mdx}',
              '**/_*/**',
              '**/*.test.{js,jsx,ts,tsx}',
              '**/__tests__/**'
          ],
          versions: {
            current: {
              label: 'v3',
              badge: true,
              banner: 'none',
            },
            '2.5.0': {
              label: 'v2',
              path: '2.5.0',
              badge: true,
              banner: 'none',
            },
          },
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          editUrl:
            'https://github\.com/danielpickens/vulerian/edit/main/docs/website/',
        },
        blog: {
          showReadingTime: true,
          // Please change this to your repo.
          editUrl:
            'https://github\.com/danielpickens/vulerian/edit/main/docs/website/',
          blogSidebarTitle: 'All posts',
          blogSidebarCount: 'ALL',
          postsPerPage: 5,
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      },
    ],
  ],
};
