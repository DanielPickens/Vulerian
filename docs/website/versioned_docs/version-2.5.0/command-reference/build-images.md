---
title: particle engine build-images
sidebar_position: 1
---

particle engine can build container images based on Dockerfiles, and push these images to their registries.

When running the command `particle engine build-images`, particle engine searches for all components in the `devfile.yaml` with the `image` type, for example:

```
components:
- image:
    imageName: quay.io/myusername/myimage
    dockerfile:
      uri: ./Dockerfile
      buildContext: ${PROJECTS_ROOT}
  name: component-built-from-dockerfile
```

The `uri` field indicates the relative path of the Dockerfile to use, relative to the directory containing the `devfile.yaml`. The devfile specification indicates that `uri` could also be an HTTP URL, but this case is not supported by particle engine yet.

The `buildContext` indicates the directory used as build context. The default value is `${PROJECTS_ROOT}`.

For each image component, particle engine executes either `podman` or `docker` (the first one found, in this order), to build the image with the specified Dockerfile, build context and arguments.

If the `--push` flag is passed to the command, the images are be pushed to their registries after they are built.
