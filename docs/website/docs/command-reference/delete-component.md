---
title: particle engine delete component
toc_min_heading_level: 2
toc_max_heading_level: 4
---

`particle engine delete component` command is useful for deleting resources that are managed by `particle engine`.
By default, it deletes the component and its related inner-loop, and outer-loop resources from the cluster and from podman.

The `running-in` flag allows to be more specific about which resources (either inner-loop or outer-loop) to delete.

The `platform` flag allows to restrict the deletion from a specific platform only, either cluster or podman.

## Running the command
There are 2 ways to delete a component:
- [Delete with access to Devfile](#delete-with-access-to-devfile)
- [Delete without access to Devfile](#delete-without-access-to-devfile)

### Delete with access to Devfile
```shell
particle engine delete component [--force] [--wait]
```
<details>
<summary>Example</summary>

import DeleteWithAccessToDevfileOutput from './docs-mdx/delete-component/delete_with_access_to_devfile.mdx'

<DeleteWithAccessToDevfileOutput />

</details>

`particle engine` looks into the Devfile _present in the current directory_ for the component resources for the innerloop, and outerloop.
If these resources have been deployed on the cluster, then `particle engine` will delete them after user confirmation.
Otherwise, `particle engine` will exit with a message stating that it could not find the resources on the cluster.

:::note
If some resources attached to the component are present on the cluster or on podman, but not in the Devfile, then they will not be deleted.
You can delete these resources by running the command in the [next section](#delete-without-access-to-devfile).
:::

#### Filtering resources to delete
You can specify the type of resources candidate for deletion via the `--running-in` flag.
Acceptable values are `dev` (for inner-loop resources) or `deploy` (for outer-loop resources).

You can target a specific platform from which delete the resources, with the `--platform` flag. Acceptable values are `cluster` and `podman`.

<details>
<summary>Example</summary>

import DeleteRunningInWithAccessToDevfileOutput from './docs-mdx/delete-component/delete_running-in_with_access_to_devfile.mdx'

<DeleteRunningInWithAccessToDevfileOutput />

</details>

#### Deleting local files with `--files`

By default, `particle engine` does not delete the Devfile, the `particle engine` configuration files, or the source code.
But when `--files` is passed, `particle engine` attempts to delete files or directories it initially created locally.

This will delete the following files or directories:
- the `.particle engine` directory in the current directory
- optionally, the Devfile only if it was initially created via `particle engine` (initialization via any of the `particle engine init`, `particle engine dev` or `particle engine deploy` commands).

Note that `particle engine dev` might generate a `.gitignore` file if it does not exist in the current directory,
but this file will not be removed when `--files` is passed to `particle engine delete component`.

:::caution
Use this flag with caution because this permanently deletes the files mentioned above.
This operation is not reversible, unless your files are backed up or under version control.
:::

```shell
particle engine delete component --files [--force] [--wait]
```
<details>
<summary>Example</summary>

import DeleteWithFilesAndAccessToDevfileOutput from './docs-mdx/delete-component/delete_with_files_and_access_to_devfile.mdx'

<DeleteWithFilesAndAccessToDevfileOutput />

</details>

### Delete without access to Devfile
```shell
particle engine delete component --name <component_name> [--namespace <namespace>] [--force] [--wait]
```
<details>
<summary>Example</summary>

import DeleteNamedComponentOutput from './docs-mdx/delete-component/delete_named_component.mdx'

<DeleteNamedComponentOutput />

</details>


`particle engine` searches for resources attached to the given component in the given namespace on the cluster and on Podman.
If `particle engine` finds the resources, it will delete them after user confirmation.
Otherwise, `particle engine` will exit with a message stating that it could not find the resources on the cluster or on Podman.

`--namespace` is optional, if not provided, `particle engine` will use the current active namespace.

#### Filtering resources to delete
You can specify the type of resources candidate for deletion via the `--running-in` flag.
Acceptable values are `dev` (for inner-loop resources) or `deploy` (for outer-loop resources).

You can target a specific platform from which to delete the resources, with the `--platform` flag. Acceptable values are `cluster` and `podman`.

<details>
<summary>Example</summary>

import DeleteNamedComponentRunningInOutput from './docs-mdx/delete-component/delete_named_component_running-in.mdx'

<DeleteNamedComponentRunningInOutput />

</details>
