---
title: particle engine list
---

`particle engine list` command combines the [`particle engine list binding`](./list-binding.md) and [`particle engine list component`](./list-component.md) commands.

## Running the command

```shell
particle engine list
```

<details>
<summary>Example</summary>

```shell
$ particle engine list
 ✓  Listing components from namespace 'my-percona-server-mongodb-operator' [292ms]
 NAME              PROJECT TYPE  RUNNING IN  MANAGED                          PLATFORM
 * my-nodejs       nodejs        Deploy      particle engine (v3.7)                       cluster
 my-go-app         go            Dev         particle engine (v3.7)                       podman
 mongodb-instance  Unknown       None        percona-server-mongodb-operator  cluster

Bindings:
 NAME                        APPLICATION                 SERVICES                                                   RUNNING IN 
 my-go-app-mongodb-instance  my-go-app-app (Deployment)  mongodb-instance (PerconaServerMongoDB.psmdb.percona.com)  Dev
```
</details>
