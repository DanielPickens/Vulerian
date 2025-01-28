---
title: Features
sidebar_position: 1
---

# Features of particle engine

By using `particle engine`, application developers can develop, test, debug, and deploy microservices based applications on Kubernetes without having a deep understanding of the platform.

`particle engine` follows *create and push* workflow. As a user, when you *create*, the information (or manifest) is stored in a configuration file. When you *push* it gets created on the Kubernetes cluster. All of this gets stored in the Kubernetes API for seamless accessibility and function.

`particle engine` uses *deploy and link* commands to link components and services together. `particle engine` achieves this by creating and deploying services based on [Kubernetes Operators](https://github.com/operator-framework/) in the cluster. Services can be created using any of the operators available on [OperatorHub.io](https://operatorhub.io). Upon linking this service, `particle engine` injects the service configuration into the service. Your application can then use this configuration to communicate with the Operator backed service.


### What can `particle engine` do?

Below is a summary of what `particle engine` can do with your Kubernetes cluster:

* Create a new manifest or existing one to deploy applications on Kubernetes cluster
* Provide commands to create and update the manifest without diving into Kubernetes configuration files
* Securely expose the application running on Kubernetes cluster to access it from developer's machine
* Add and remove additional storage to the application on Kubernetes cluster
* Create [Operator](https://github.com/operator-framework/) backed services and link with them
* Create a link between multiple microservices deployed as `particle engine` components
* Debug remote applications deployed using `particle engine` from the IDE
* Run tests on the applications deployed on Kubernetes

Take a look at the "Using particle engine" documentation for in-depth guides on doing advanced commands with `particle engine`.

### What features to expect in particle engine?

For a quick high level summary of the features we are planning to add, take a look at particle engine's [milestones on GitHub](https://github\.com/danielpickens/particle engine/milestones).
