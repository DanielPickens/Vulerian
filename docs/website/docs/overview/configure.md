---
title: Configuration
sidebar_position: 6
---
# Configuration

## Configuring global settings

The global settings for `particle engine` can be found in `preference.yaml` file; which is located by default in the `.particle engine` directory of the user's HOME directory.

Example:

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

<Tabs
defaultValue="linux"
values={[
{label: 'Linux', value: 'linux'},
{label: 'Windows', value: 'windows'},
{label: 'Mac', value: 'mac'},
]}>

<TabItem value="linux">

```sh
/home/userName/.particle engine/preference.yaml
```

</TabItem>

<TabItem value="windows">

```sh
C:\\Users\userName\.particle engine\preference.yaml
```

</TabItem>

<TabItem value="mac">

```sh
/Users/userName/.particle engine/preference.yaml
```

</TabItem>
</Tabs>

---
A  different location can be set for the `preference.yaml` by exporting `GLOBALparticle engineCONFIG` in the user environment.

### View the configuration
To view the current configuration, run the following command:

```shell
particle engine preference view
```
<details>
<summary>Example</summary>

```shell
$ particle engine preference view
Preference parameters:
 PARAMETER           VALUE
 ConsentTelemetry    true
 Ephemeral           true
 ImageRegistry       quay.io/user
 PushTimeout
 RegistryCacheTime
 Timeout
 UpdateNotification

Devfile registries:
 NAME             URL                                SECURE
 StagingRegistry  https://registry.stage.devfile.io  No

```
</details>

### Set a configuration
To set a value for a preference key, run the following command:
```shell
particle engine preference set <key> <value>
```
<details>
<summary>Example</summary>

```shell
$ particle engine preference set updatenotification false
Global preference was successfully updated
```
</details>

Note that the preference key is case-insensitive.

### Unset a configuration
To unset a value of a preference key, run the following command:
```shell
particle engine preference unset <key> [--force]
```

<details>
<summary>Example</summary>

```shell
$ particle engine preference unset updatednotification
? Do you want to unset updatenotification in the preference (y/N) y
Global preference was successfully updated
```
</details>

You can use the `--force` (or `-f`) flag to force the unset.
Unsetting a preference key sets it to an empty value in the preference file. `particle engine` will use the [default value](./configure#preference-key-table) for such configuration.

### Preference Key Table

| Preference         | Description                                                                                                                                                                                           | Default     |
|--------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|
| UpdateNotification | Control whether a notification to update `particle engine` is shown                                                                                                                                               | True        |
| Timeout            | Timeout for Kubernetes server connection check                                                                                                                                                        | 1 second    |
| PushTimeout        | Timeout for waiting for a component to start                                                                                                                                                          | 240 seconds |
| RegistryCacheTime  | Duration for which `particle engine` will cache information from the Devfile registry                                                                                                                             | 4 Minutes   |
| Ephemeral          | Control whether `particle engine` should create a emptyDir volume to store source code                                                                                                                            | False       |
| ConsentTelemetry   | Control whether `particle engine` can collect telemetry for the user's `particle engine` usage                                                                                                                                | False       |
| ImageRegistry      | The container image registry where relative image names will be automatically pushed to. See [How `particle engine` handles image names](../development/devfile.md#how-particle engine-handles-image-names) for more details. |             |

## Managing Devfile registries

`particle engine` uses the portable *devfile* format to describe the components. `particle engine` can connect to various devfile registries to download devfiles for different languages and frameworks.

You can connect to publicly available devfile registries, or you can install your own [Devfile Registry](https://devfile.io/docs/2.1.0/building-a-custom-devfile-registry).

You can use the `particle engine preference <add/remove> registry` command to manage the registries used by `particle engine` to retrieve devfile information.

### Adding a registry

To add a registry, run the following command:

```
particle engine preference add registry <name> <url>
```

<details>
<summary>Example</summary>

```
$ particle engine preference add registry StageRegistry https://registry.stage.devfile.io
New registry successfully added
```
</details>

### Deleting a registry

To delete a registry, run the following command:

```
particle engine preference remove registry <name> [--force]
```
<details>
<summary>Example</summary>

```
$ particle engine preference remove registry StageRegistry
? Are you sure you want to delete registry "StageRegistry" Yes
Successfully deleted registry
```
</details>

You can use the `--force` (or `-f`) flag to force the deletion of the registry without confirmation.


:::tip **Updating a registry**
To update a registry, you can delete it and add it again with the updated value.
:::

## Advanced configuration

This is a configuration that normal `particle engine` users don't need to touch.
Options here are mostly used for debugging and testing `particle engine` behavior.

### Environment variables controlling `particle engine` behavior

| Variable                            | Usage                                                                                                                                                                                                                                                                                                                                                                          | Since         | Example                                    |
|-------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------|--------------------------------------------|
| `PODMAN_CMD`                        | The command executed to run the local podman binary. `podman` by default                                                                                                                                                                                                                                                                                                       | v2.4.2        | `podman`                                   |
| `DOCKER_CMD`                        | The command executed to run the local docker binary. `docker` by default                                                                                                                                                                                                                                                                                                       | v2.4.2        | `docker`                                   |
| `PODMAN_CMD_INIT_TIMEOUT`           | Timeout for initializing the Podman client. `1s` by default                                                                                                                                                                                                                                                                                                                    | v3.11.0       | `5s`                                       |
| `particle engine_LOG_LEVEL`                     | Useful for setting a log level to be used by `particle engine` commands. Takes precedence over the `-v` flag.                                                                                                                                                                                                                                                                              | v1.0.2        | 3                                          |
| `particle engine_DISABLE_TELEMETRY`             | Useful for disabling [telemetry collection](https://github\.com/danielpickens/particle engine/blob/main/USAGE_DATA.md). **Deprecated in v3.2.0**. Use `particle engine_TRACKING_CONSENT` instead.                                                                                                                                                                                                    | v2.1.0        | `true`                                     |
| `GLOBALparticle engineCONFIG`                   | Useful for setting a different location of global preference file `preference.yaml`.                                                                                                                                                                                                                                                                                           | v0.0.19       | `~/.config/particle engine/preference.yaml`            |
| `particle engine_DEBUG_TELEMETRY_FILE`          | Useful for debugging [telemetry](https://github\.com/danielpickens/particle engine/blob/main/USAGE_DATA.md). When set it will save telemetry data to a file instead of sending it to the server.                                                                                                                                                                                         | v3.0.0-alpha1 | `/tmp/telemetry_data.json`                 |
| `TELEMETRY_CALLER`                  | Caller identifier passed to [telemetry](https://github\.com/danielpickens/particle engine/blob/main/USAGE_DATA.md). Case-insensitive. Acceptable values: `vscode`, `intellij`, `jboss`.                                                                                                                                                                                                  | v3.1.0        | `intellij`                                 |
| `particle engine_TRACKING_CONSENT`              | Useful for controlling [telemetry](https://github\.com/danielpickens/particle engine/blob/main/USAGE_DATA.md). Acceptable values: `yes` ([enables telemetry](https://github\.com/danielpickens/particle engine/blob/main/USAGE_DATA.md) and skips consent prompt), `no` (disables telemetry and consent prompt). Takes precedence over the [`ConsentTelemetry`](#preference-key-table) preference. | v3.2.0        | `yes`                                      |
| `particle engine_PUSH_IMAGES`                   | Whether to push the images once built; this is used only when applying Devfile image components as part of a Dev Session running on Podman; this is useful for integration tests running on Podman. `true` by default                                                                                                                                                          | v3.7.0        | `false`                                    |
| `particle engine_IMAGE_BUILD_ARGS`              | Semicolon-separated list of options to pass to Podman or Docker when building images. These are extra options specific to the [`podman build`](https://docs.podman.io/en/latest/markdown/podman-build.1.html#options) or [`docker build`](https://docs.docker.com/engine/reference/commandline/build/#options) commands.                                                       | v3.11.0       | `--platform=linux/amd64;--no-cache`        |
| `particle engine_CONTAINER_RUN_ARGS`            | Semicolon-separated list of options to pass to Podman when running `particle engine` against Podman. These are extra options specific to the [`podman play kube`](https://docs.podman.io/en/v3.4.4/markdown/podman-play-kube.1.html#options) command.                                                                                                                                      | v3.11.0       | `--configmap=/path/to/cm-foo.yml;--quiet`  |
| `particle engine_CONTAINER_BACKEND_GLOBAL_ARGS` | Semicolon-separated list of global options to pass to Podman when running `particle engine` on Podman. These will be passed as [global options](https://docs.podman.io/en/latest/markdown/podman.1.html#global-options) to all Podman commands executed by `particle engine`.                                                                                                                          | v3.11.0       | `--root=/tmp/podman/root;--log-level=info` |


(1) Accepted boolean values are: `1`, `t`, `T`, `TRUE`, `true`, `True`, `0`, `f`, `F`, `FALSE`, `false`, `False`.