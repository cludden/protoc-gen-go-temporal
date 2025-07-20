import type {SidebarsConfig} from '@docusaurus/plugin-content-docs';

/**
 * Creating a sidebar enables you to:
 - create an ordered group of docs
 - render a sidebar for each doc of that group
 - provide next/previous navigation

 The sidebars can be generated from the filesystem, or explicitly defined here.

 Create as many sidebars as you want.
 */
const sidebars: SidebarsConfig = {
  // By default, Docusaurus generates a sidebar from the docs folder structure
  docs: [
    {
      type: 'doc',
      id: 'install',
      label: 'Installing'
    },
    {
      type: 'doc',
      id: 'getting-started',
      label: 'Getting Started'
    },
    {
      type: 'category',
      label: 'Configuration',
      items: [
        {
          type: 'doc',
          id: 'configuration/plugin',
          label: 'Plugin'
        },
        {
          type: 'doc',
          id: 'configuration/service',
          label: 'Service'
        },
        {
          type: 'doc',
          id: 'configuration/workflow',
          label: 'Workflow'
        },
        {
          type: 'doc',
          id: 'configuration/activity',
          label: 'Activities'
        },
        {
          type: 'doc',
          id: 'configuration/query',
          label: 'Query'
        },
        {
          type: 'doc',
          id: 'configuration/signal',
          label: 'Signal'
        },
        {
          type: 'doc',
          id: 'configuration/update',
          label: 'Update'
        },
        {
          type: 'doc',
          id: 'configuration/fields',
          label: 'Fields'
        }
      ]
    },
    {
      type: 'category',
      label: 'Guides',
      items: [
        {
          type: 'doc',
          id: 'guides/workflows',
          label: 'Workflows'
        },
        {
          type: 'doc',
          id: 'guides/activities',
          label: 'Activities'
        },
        {
          type: 'doc',
          id: 'guides/queries',
          label: 'Queries'
        },
        {
          type: 'doc',
          id: 'guides/signals',
          label: 'Signals'
        },
        {
          type: 'doc',
          id: 'guides/testing',
          label: 'Testing'
        },
        {
          type: 'doc',
          id: 'guides/updates',
          label: 'Updates'
        },
        {
          type: 'doc',
          id: 'guides/clients',
          label: 'Clients'
        },
        {
          type: 'doc',
          id: 'guides/cli',
          label: 'CLI'
        },
        {
          type: 'doc',
          id: 'guides/child-workflows',
          label: 'Child Workflows'
        },
        {
          type: 'doc',
          id: 'guides/xns',
          label: 'Cross-Namespace (XNS)'
        },
        {
          type: 'doc',
          id: 'guides/bloblang',
          label: 'Bloblang Expressions'
        },
        {
          type: 'doc',
          id: 'guides/codec-server',
          label: 'Codec Server'
        },
        {
          type: 'doc',
          id: 'guides/documentation',
          label: 'Documentation'
        },
        {
          type: 'doc',
          id: 'guides/patches',
          label: 'Patches'
        },
      ]
    },
  ],

  examples: [
    {
      type: 'doc',
      id: 'examples/helloworld',
      label: 'Hello World'
    },
    {
      type: 'doc',
      id: 'examples/codecserver',
      label: 'Codec Server'
    },
    {
      type: 'doc',
      id: 'examples/xns',
      label: 'Cross-Namespace (XNS)'
    },
    {
      type: 'doc',
      id: 'examples/mutex',
      label: 'Mutex'
    },
    {
      type: 'doc',
      id: 'examples/schedule',
      label: 'Schedule'
    },
    {
      type: 'doc',
      id: 'examples/searchattributes',
      label: 'Search Attributes'
    },
    {
      type: 'doc',
      id: 'examples/shoppingcart',
      label: 'Shopping Cart'
    },
    {
      type: 'doc',
      id: 'examples/updatabletimer',
      label: 'Updatable Timer'
    }
  ]

  // But you can create a sidebar manually
  /*
  tutorialSidebar: [
    'intro',
    'hello',
    {
      type: 'category',
      label: 'Tutorial',
      items: ['tutorial-basics/create-a-document'],
    },
  ],
   */
};

export default sidebars;
