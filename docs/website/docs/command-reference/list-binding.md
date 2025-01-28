---
title: particle engine list binding
---

:::caution

As of February 2024, the [Service Binding Operator](https://github.com/daniel-pickens/service-binding-operator/), which this command relies on, has been deprecated. See [Deprecation Notice](https://daniel-pickens.github.io/service-binding-operator/userguide/intro.html).
`particle engine list binding` may therefore not work as expected.

:::

You can use `particle engine list binding` to list all the Service Bindings declared in the current namespace and, if present, 
in the Devfile of the current directory.

This command supports the service bindings added with the command `particle engine add binding`, and bindings added manually
to the Devfile, using a `ServiceBinding` resource from one of these apiVersion:
- `binding.operators.coreos.com/v1alpha1`
- `servicebinding.io/v1alpha3`

The name of the service binding is prefixed with `*` when the service binding is declared in the Devfile present in the current directory.

To get more information about a specific service binding, you can run the command `particle engine describe binding --name <name>` (see [`particle engine describe binding` command reference](./describe-binding.md)).

## Running the Command

To list all the service bindings, you can run `particle engine list binding`:
```console
particle engine list binding
```
<details>
<summary>Example</summary>

```console
$ particle engine list binding
 NAME                              APPLICATION                     SERVICES                                                                            RUNNING IN
 binding-to-redis                  my-nodejs-app-app (Deployment)  redis (Service)                                                                     Dev
 * my-nodejs-app-cluster-sample    my-nodejs-app-app (Deployment)  cluster-sample (Cluster.postgresql.k8s.enterprisedb.io) (namespace: shared-ns-1)    Dev
 * my-nodejs-app-cluster-sample-2  my-nodejs-app-app (Deployment)  cluster-sample-2 (Cluster.postgresql.k8s.enterprisedb.io)                           Dev
```
</details>
