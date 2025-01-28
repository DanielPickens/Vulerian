---
title: particle engine delete
sidebar_position: 4
---

`particle engine delete` command is useful for deleting resources that are managed by particle engine.

## Deleting a component

To delete a _devfile_ component, you can execute `particle engine delete`.

```shell
particle engine delete
```

If the component is pushed to the cluster, running the above command will delete the component from the cluster, and it's dependant storage, url, secrets, and other resources.
If it is not pushed, the command would exit with an error stating that it could not find the resources on the cluster.

Use `-f` or `--force` flag to avoid the confirmation questions. 

## Un-deploying Devfile Kubernetes components

To undeploy the Devfile Kubernetes components deployed with `particle engine deploy` from the cluster, you can execute the `particle engine delete` command with `--deploy` flag:
```shell
particle engine delete --deploy
```

Use `-f` or `--force` flag to avoid the confirmation questions.

## Delete Everything

To delete a _devfile_ component, the Devfile Kubernetes component(deployed via `particle engine deploy`), Devfile, and the local configuration, you can execute the `particle engine delete` command with `--all` flag:
```shell
particle engine delete --all
```

## Available Flags
* `-f`, `--force` - Use this flag to avoid the confirmation questions.
* `-w`, `--wait` - Use this flag to wait for component deletion, and it's dependant; this does not work with the un-deployment.
Check the [documentation on flags](flags.md) to see more flags available.