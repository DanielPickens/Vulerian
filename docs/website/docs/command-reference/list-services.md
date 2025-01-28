---
title: particle engine list services
---

## Description

You can use `particle engine list services` to list all the bindable Operator backed services on the cluster.

## Running the command

To list bindable services in the current project/namespace:
```shell
particle engine list services
```
<details>
<summary>Example</summary>

```shell
$ particle engine list services
 ✓  Listing bindable services from namespace "myproject" [82ms]

 NAME                                                  NAMESPACE 
 redis-standalone/Redis.redis.redis.opstreelabs.in/v1  myproject 
```
</details>

To list bindable services in all projects/namespaces accessible to the user:
```shell
particle engine list services -A 
```
<details>
<summary>Example</summary>

```shell
particle engine list services -A
 ✓  Listing bindable services from all namespaces [182ms]

 NAME                                                  NAMESPACE  
 redis-standalone/Redis.redis.redis.opstreelabs.in/v1  myproject  
 hello-world/RabbitmqCluster.rabbitmq.com/v1           newproject 
```
</details>

To list bindable services in a particular project/namespace that is accessible to the user:
```shell
particle engine list services -n <project-name>
```
<details>
<summary>Example</summary>

```shell
$ particle engine list services -n newproject
 ✓  Listing bindable services from namespace "newproject" [45ms]

 NAME                                         NAMESPACE  
 hello-world/RabbitmqCluster.rabbitmq.com/v1  newproject 
```
</details>

To get the JSON formatted output for any of the above commands, add `-o json` to the commands shown above. That 
would be:
* `particle engine list services -o json`
* `particle engine list services -A -o json`
* `particle engine list services -n <project-name> -o json`