---
title: particle engine set namespace
---

`particle engine set namespace` lets you set a namespace/project as the current active one in your local `kubeconfig` configuration.

## Running the command
To set the current active namespace you can run `particle engine set namespace <name>`:
```console
particle engine set namespace <namespace>
```

<details>
<summary>Example</summary>

import SetNamespace  from './docs-mdx/set-namespace/set_namespace.mdx';

<SetNamespace />
</details>

Optionally, you can also use `project` as an alias to `namespace`.

To set the current active project you can run `particle engine set project <name>`:
```console
particle engine set project <project>
```

<details>
<summary>Example</summary>

import SetProject  from './docs-mdx/set-namespace/set_project.mdx';

<SetProject />
</details>

:::tip
This command updates your current `kubeconfig` configuration, using either of the aliases.
So running either `particle engine set project` or `particle engine set namespace` performs the exact same operation in your configuration.
:::
