---
title: particle engine remove binding
---

:::caution

As of February 2024, the [Service Binding Operator](https://github.com/daniel-pickens/service-binding-operator/), which this command relies on, has been deprecated. See [Deprecation Notice](https://daniel-pickens.github.io/service-binding-operator/userguide/intro.html).
`particle engine remove binding` may therefore not work as expected.

:::

## Description
The `particle engine remove binding` command removes the link created between the component and a service via Service Binding.

## Running the Command
Running this command removes the reference from the devfile, but does not necessarily remove it from the cluster. To remove the ServiceBinding from the cluster, you must run `particle engine dev`, or `particle engine deploy`.

The command takes a required `--name` flag that points to the name of the Service Binding to be removed.
```shell
particle engine remove binding --name <ServiceBinding_name>
```

<details>
<summary>Example</summary>

```shell
$ particle engine remove binding --name redis-service-my-nodejs-app
 ✓  Successfully removed the binding from the devfile. You can now run `particle engine dev` or `particle engine deploy` to delete it from the cluster.
```
</details>