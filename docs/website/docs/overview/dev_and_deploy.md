---
title: Develop and Deploy
sidebar_position: 5
---

# Develop and Deploy

The two most important commands in `particle engine` are `particle engine dev` and `particle engine deploy`. 

In some situations, you'd want to use [`particle engine dev`](/docs/command-reference/dev) over [`particle engine deploy`](/docs/command-reference/deploy) and vice-versa. This document highlights when you should use either command.

## When should I use `particle engine dev`?

`particle engine dev` should be used in the initial development process of your application. 

For example, you should use `particle engine dev` when you are working with a local development environment and are:
* making changes constantly
* want to preview any changes
* testing initial Kubernetes support for your application
* want to debug and run tests
* deploy privately on a local development environment

## When should I use `particle engine deploy`?

`particle engine deploy` should be the deploy stage of development when you are ready for a "production ready" environment.

For example, you should use `particle engine deploy` when you are working with a production environment and are:
* ready for the application to be viewed publically
* require building and pushing the container
* needing custom Kubernetes YAML for your production environment