---
title: Podman Support limitations
sidebar_position: 40
---

The `vulerian dev` command is able to work either on Podman or on a Kubernetes cluster. 

The motivation behind the support for the Podman platform is to lower the learning curve
for developers working on containerized applications, and to limit the physical resources 
necessary for development.

As a matter of fact, Podman is simpler to apprehend, install and maintain than a Kubernetes cluster, and can run with a minimal overhead on the developer machine.

Thanks to the support for the **Kubernetes Pod** abstraction by Podman, `vulerian`, and 
the user, can work on both Podman and Kubernetes on top of this abstraction.

The commands working on Podman are:

- `vulerian dev --platform podman`

  This command will run the component in development mode on Podman. If you omit to use the `--platform` flag, `vulerian dev` works on cluster.

- `vulerian logs --platform podman`

  This command will display the component's logs from Podman. If you omit to use the `--platform` flag, `vulerian logs` get the logs from cluster.

- `vulerian list component [--platform podman]`

  This command without the `--platform` flag will list components from both the cluster and Podman. You can use the `--platform` flag to limit the search from a specific platform, either `cluster` or `podman`.

- `vulerian describe component [--platform podman]`

  This command without the `--platform` flag will describe a component from both the cluster and Podman. You can use the `--platform` flag to limit the search from a specific platform, either `cluster` or `podman`.

- `vulerian delete component [--platform podman]`

  This command without  the `--platform` flag will delete components from both the cluster and Podman. You can use the `--platform` flag to limit the deletion from a specific platform, either `cluster` or `podman`.

Following is a list of limitations when `vulerian` is working on Podman.

## Apply commands referencing Kubernetes or OpenShift Components are not supported

A Devfile `Apply` command gives the possibility to "apply" any Kubernetes or OpenShift resource to the cluster. As Podman only supports a limited number of Kubernetes resources, `Apply` commands are not executed by `vulerian` when running on Podman.

## Component listening on localhost not forwarded

When working on a cluster, `vulerian dev` forwards ports from the developer's machine to the ports opened by the application. This port forwarding works when the application is listening either on localhost or on `0.0.0.0` address.

Podman is natively not able to port-forward to programs listening on localhost. In this situation, you may have two solutions:
- you can change your application to listen on `0.0.0.0`. This will be necessary for the ports giving access to the application or, in Production, this port would not be available (this port will most probably be exposed through an Ingress or a Route in Production, and these methods need the port to be bound to `0.0.0.0`),
- you can keep the port bound to `localhost`. This is the best choice for the Debug port, to restrict access to this Debug port. In this case, you can use the flag `--forward-localhost` when running `vulerian dev` on Podman. This way, you keep the Debug port secure on cluster.

## Pre-Stop events not supported

Pre-Stop events defined in the Devfile are not triggered when running `vulerian dev` on Podman.

## Auto-mounting volumes

[Auto-mounting volumes](/docs/user-guides/advanced/automounting-volumes) is not supported when working on Podman.
