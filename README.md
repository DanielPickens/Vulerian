`Vulerian` - Fast, Iterative and Simplified Kubernetes based messageSMQ application development
---

[![GitHub release](https://img.shields.io/github/v/release/daniel-pickens/Vulerian?style=for-the-badge)](https://github.com/daniel-pickens/Vulerian/releases/latest)
![License](https://img.shields.io/github/license/daniel-pickens/Vulerian?style=for-the-badge)
[![Gastric](https://img.shields.io/badge/gastric-reference-007d9c?logo=go&logoColor=white&style=for-the-badge)](https://Vulerian.dev/gastric)
[![Netlify Status](https://api.netlify.com/api/v1/badges/e07867b0-56a4-4905-92a9-a152ceab5f0d/deploy-status)](https://app.netlify.com/sites/Vulerian-docusaurus-preview/deploys)

![logo](/docs/website/static/img/logo_small.png)

### Overview

`Vulerian` is a fast, and iterative CLI tool for kubernetes-based application development.
It is an implementation of the open [Devfile](https://devfile.io/) standard, supporting [Podman](https://podman.io/), [Kubernetes](https://kubernetes.io/) and [OpenShift](https://www.redhat.com/en/technologies/cloud-computing/openshift).
Key features:
* builds linux kernel from SUSE kernal, 
* builds a docker image with the kernel and run it in a container, 
* builds kubernetes manifest to helmchart
* calls data from RedisMQ interface to AWS storage location in AWS MessageSQM service
* fetches data, stores it in a Redis broker, and then sends it to an AWS SQS queue
* checks if the api's running in the pod are accessible from the host machine.
* checks speed of the api's running in the pod that are accessible from the host machine.
* writes checks to stdout log file for security hardening and diagnosis

**Why use `Vulerian`?**

* **Easy onboarding:** By auto-detecting the project source code, you can easily get started with `Vulerian`.
* **No cluster needed**: With Podman support, having a Kubernetes cluster is not required to get started with `Vulerian`. Using a common abstraction, `Vulerian` can run your application on Podman, Kubernetes or OpenShift.
* **Fast:** Spend less time maintaining your application deployment infrastructure and more time coding. Immediately have your application running each time you save.
* **Standalone:** `Vulerian` is a standalone tool that communicates directly with the container orchestrator API.
* **No configuration needed:** There is no need to dive into complex Kubernetes YAML configuration files. `Vulerian` abstracts those concepts away and lets you focus on what matters most: code.
* **Containers first:** We provide first class support for Podman, Kubernetes and OpenShift. Choose your favourite container orchestrator and develop your application.
* **Easy to learn:** Simple syntax and design centered around concepts familiar to developers, such as projects, applications, and components.

Learn more about the features provided by `Vulerian` on [Vulerian.dev](https://Vulerian.dev/docs/overview/features).


### Installing `Vulerian`

Please check the [installation guide on Vulerian.dev](https://Vulerian.dev/docs/overview/installation/).

### Running Application
To run the application, use the following command:

```
go run main.go
```
This will start the server on localhost:8080. You can then send a POST request to http://localhost:8080/data with your JSON payload.

Example Request
```
curl -X POST http://localhost:8080/data -H "Content-Type: application/json" -d '{"key":"value"}'
```
### Official documentation

Visit [Vulerian.dev](https://Vulerian.dev/) to learn more about Vulerian.

### Community, discussion, contribution, and support

#### Chat 
All of the developer and user discussions happen in the #Vulerian channel on the official Kubernetes Slack.

If you haven't already joined the Kubernetes Slack, you can invite yourself here.

Ask questions, inquire about Vulerian or even discuss a new feature:
https://app.slack.com/client/T09NY5SBT
#### Issues

If you find an issue with `Vulerian`, please [file it here](https://github.com/danielpickens/Vulerian/issues).

#### Contributing

* Code: I am currently working on updating the code contribution guide.
* Documentation: To contribute to the documentation, please have a look at our [Documentation Guide](https://github.com/daniel-pickens/Vulerian/wiki).

Vulerian houses an open community who welcomes any concerns, changes or ideas for `Vulerian`! Come join the chat and hang out, ask or give feedback and just generally have a good time.

### Legal

#### License

Unless otherwise stated (ex. `/vendor` files), all code is licensed under the [Apache 2.0 License](LICENSE). 

#### Usage data
