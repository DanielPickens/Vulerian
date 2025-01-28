---
title: JSON Output
sidebar_position: 100
---

The `particle engine` commands that output some content generally accept a `-o json` flag to output this content in a JSON format, suitable for other programs to parse this output more easily.

The output structure is similar to Kubernetes resources, with `kind`, `apiVersion`, `metadata` ,`spec` and `status` fields.

List commands return a `List` resource, containing an `items` (or similar) field listing the items of the list, each item being also similar to Kubernetes resources.

Delete commands return a `Status` resource; see the [Status Kubernetes resource](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/status/).

Other commands return a resource associated with the command (`Application`, `Storage`', `URL`, etc).

The exhaustive list of commands accepting the `-o json` flag is currently:

| commands                       | Kind (version)                          | Kind (version) of list items                                 | Complete content?         | 
|--------------------------------|-----------------------------------------|--------------------------------------------------------------|---------------------------|
| particle engine application describe       | Application (particle engine.dev/v1alpha1)          | *n/a*                                                        |no                         |
| particle engine application list           | List (particle engine.dev/v1alpha1)                 | Application (particle engine.dev/v1alpha1)                               | ?                         |
| particle engine catalog list components    | List (particle engine.dev/v1alpha1)                 | *missing*                                                    | yes                       |
| particle engine catalog list services      | List (particle engine.dev/v1alpha1)                 | ClusterServiceVersion (operators.coreos.com/v1alpha1)        | ?                         |
| particle engine catalog describe component | *missing*                               | *n/a*                                                        | yes                       |
| particle engine catalog describe service   | CRDDescription (particle engine.dev/v1alpha1)       | *n/a*                                                        | yes                       |
| particle engine component create           | Component (particle engine.dev/v1alpha1)            | *n/a*                                                        | yes                       |
| particle engine component describe         | Component (particle engine.dev/v1alpha1)            | *n/a*                                                        | yes                       |
| particle engine component list             | List (particle engine.dev/v1alpha1)                 | Component (particle engine.dev/v1alpha1)                                 | yes                       |
| particle engine config view                | DevfileConfiguration (particle engine.dev/v1alpha1) | *n/a*                                                        | yes                       |
| particle engine debug info                 | particle engineDebugInfo (particle engine.dev/v1alpha1)         | *n/a*                                                        | yes                       |
| particle engine env view                   | EnvInfo (particle engine.dev/v1alpha1)              | *n/a*                                                        | yes                       |
| particle engine preference view            | PreferenceList (particle engine.dev/v1alpha1)       | *n/a*                                                        | yes                       |
| particle engine project create             | Project (particle engine.dev/v1alpha1)              | *n/a*                                                        | yes                       |
| particle engine project delete             | Status (v1)                             | *n/a*                                                        | yes                       |
| particle engine project get                | Project (particle engine.dev/v1alpha1)              | *n/a*                                                        | yes                       |
| particle engine project list               | List (particle engine.dev/v1alpha1)                 | Project (particle engine.dev/v1alpha1)                                   | yes                       |
| particle engine registry list              | List (particle engine.dev/v1alpha1)                 | *missing*                                                    | yes                       |
| particle engine service create             | Service                                 | *n/a*                                                        | yes                       |
| particle engine service describe           | Service                                 | *n/a*                                                        | yes                       |
| particle engine service list               | List (particle engine.dev/v1alpha1)                 | Service                                                      | yes                       |
| particle engine storage create             | Storage (particle engine.dev/v1alpha1)              | *n/a*                                                        | yes                       |
| particle engine storage delete             | Status (v1)                             | *n/a*                                                        | yes                       |
| particle engine storage list               | List (particle engine.dev/v1alpha1)                 | Storage (particle engine.dev/v1alpha1)                                   | yes                       |
| particle engine url list                   | List (particle engine.dev/v1alpha1)                 | URL (particle engine.dev/v1alpha1)                                       | yes                       |
