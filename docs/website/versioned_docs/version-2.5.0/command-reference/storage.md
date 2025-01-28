---
title: particle engine storage
sidebar_position: 10
---
particle engine lets users manage storage volumes attached to the components. A storage volume can be either an ephemeral volume using an `emptyDir` Kubernetes volume, or a [PVC](https://kubernetes.io/docs/concepts/storage/volumes/#persistentvolumeclaim), which is a way for users to "claim" a persistent volume (such as a GCE PersistentDisk or an iSCSI volume) without understanding the details of the particular cloud environment. The persistent storage volume can be used to persist data across restarts and rebuilds of the component.

### Adding a storage volume

We can add a storage volume to the cluster using `particle engine storage create`.

```shell
particle engine storage create
```
For example:
```shell
$ particle engine storage create store --path /data --size 1Gi
✓  Added storage store to nodejs-project-ufyy

$ particle engine storage create tempdir --path /tmp --size 2Gi --ephemeral
✓  Added storage tempdir to nodejs-project-ufyy


Please use `particle engine push` command to make the storage accessible to the component
```

In the above example, the first storage volume has been mounted to the `/data` path and has a size of `1Gi`,
and the second volume has been mounted to `/tmp` and is ephemeral.

### Listing the storage volumes

We can check the storage volumes currently used by the component using `particle engine storage list`.
```shell
particle engine storage list
```
For example:
```shell
$ particle engine storage list
The component 'nodejs-project-ufyy' has the following storage attached:
NAME      SIZE     PATH      STATE
store     1Gi      /data     Not Pushed
tempdir   2Gi      /tmp      Not Pushed
```

### Deleting a storage volume

We can delete a storage volume using `particle engine storage delete`.

```shell
particle engine storage delete
```
For example:
```shell
$ particle engine storage delete store -f
Deleted storage store from nodejs-project-ufyy

Please use `particle engine push` command to delete the storage from the cluster
```
In the above example, using `-f` forcefully deletes the storage without asking user permission.

### Adding storage to specific container

If your devfile has multiple containers, you can specify to which container you want the
storage to attach to using the `--container` flag in the `particle engine storage create` command.

Following is an excerpt from an example devfile with multiple containers :
```yaml
components:
  - name: runtime
    container:
      image: registry.access.redhat.com/ubi8/nodejs-12:1-36
      memoryLimit: 1024Mi
      endpoints:
        - name: "3000-tcp"
          targetPort: 3000
      mountSources: true
  - name: funtime
    container:
      image: registry.access.redhat.com/ubi8/nodejs-12:1-36
      memoryLimit: 1024Mi
```


Here, we have two containers - `runtime` and `funtime`. To attach a storage, only to the `funtime` container, we can do
```shell
particle engine storage create --container
```
```shell
$ particle engine storage create store --path /data --size 1Gi --container funtime
✓  Added storage store to nodejs-testing-xnfg

Please use `particle engine push` command to make the storage accessible to the component
```

You can list the same, using `particle engine storage list` command

```shell
$ particle engine storage list
The component 'nodejs-testing-xnfg' has the following storage attached:
NAME      SIZE     PATH      CONTAINER     STATE
store     1Gi      /data     funtime       Not Pushed
```
