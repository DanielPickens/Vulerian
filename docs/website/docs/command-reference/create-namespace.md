---
title: particle engine create namespace
---

`particle engine create namespace` lets you create a namespace/project on your cluster. If you are on a Kubernetes cluster, running the command will create a Namespace resource for you, and for an OpenShift cluster, it will create a Project resource.

Any new namespace created with this command will also be set as the current active namespace, this applies to project as well.

## Running the command
To create a namespace, run the following command:
```shell
particle engine create namespace <name>
```
<details>
<summary>Example</summary>

import CreateNamespace  from './docs-mdx/create-namespace/create_namespace.mdx';

<CreateNamespace />

</details>


Optionally, you can also use `project` as an alias to `namespace`.

To create a project, run the following command:
```shell
particle engine create project <name>
```
<details>
<summary>Example</summary>

import CreateProject from './docs-mdx/create-namespace/create_project.mdx';

<CreateProject />

</details>


:::tip
Using either of the aliases will not make any change to the resource created on the cluster. This command is smart enough to detect the resources supported by your cluster and make an informed decision on the type of resource that should be created.

So you can run `particle engine create project` on a Kubernetes cluster, and it will create a `Namespace` resource, and you can run `particle engine create namespace` on an OpenShift cluster, it will create a `Project` resource.
:::
