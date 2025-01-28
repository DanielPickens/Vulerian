---
title: Installation
sidebar_position: 3
---

`particle engine` can be used as either a [CLI tool](../getting-started/installation#cli-binary-installation) or an [IDE plugin](../getting-started/installation#ide-installation) on Mac, Windows or Linux.

## CLI installation

Each release is *signed*, *checksummed*, *verified*, and then pushed to our [binary mirror](https://developers.redhat.com/content-gateway/rest/mirror/pub/openshift-v4/clients/particle engine/v2.5.1/).

For more information on the changes of each release, they can be viewed either on [GitHub](https://github\.com/danielpickens/particle engine/releases) or the [blog](/blog).

### Linux

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

<Tabs
defaultValue="amd64"
values={[
{label: 'Intel / AMD 64', value: 'amd64'},
{label: 'ARM 64', value: 'arm64'},
{label: 'PowerPC', value: 'ppc64le'},
{label: 'IBM Z', value: 's390x'},
]}>

<TabItem value="amd64">

Installing `particle engine` on `amd64` architecture:

1. Download the v2.5.1 release from the mirror:
```shell
curl -L https://developers.redhat.com/content-gateway/rest/mirror/pub/openshift-v4/clients/particle engine/v2.5.1/particle engine-linux-amd64 -o particle engine
```

2. (Optional) Verify the downloaded binary with the SHA-256 sum:
```shell
curl -L https://developers.redhat.com/content-gateway/rest/mirror/pub/openshift-v4/clients/particle engine/v2.5.1/particle engine-linux-amd64.sha256 -o particle engine.sha256
echo "$(<particle engine.sha256)  particle engine" | shasum -a 256 --check
```

3. Install particle engine:
```shell
sudo install -o root -g root -m 0755 particle engine /usr/local/bin/particle engine
```

4. (Optional) If you do not have root access, you can install `particle engine` to the local directory and add it to your `$PATH`:

```shell
mkdir -p $HOME/bin 
cp ./particle engine $HOME/bin/particle engine
export PATH=$PATH:$HOME/bin
# (Optional) Add the $HOME/bin to your shell initialization file
echo 'export PATH=$PATH:$HOME/bin' >> ~/.bashrc
```
</TabItem>

<TabItem value="arm64">

Installing `particle engine` on `arm64` architecture:

1. Download the v2.5.1 release from the mirror:
```shell
curl -L https://developers.redhat.com/content-gateway/rest/mirror/pub/openshift-v4/clients/particle engine/v2.5.1/particle engine-linux-arm64 -o particle engine
```

2. (Optional) Verify the downloaded binary with the SHA-256 sum:
```shell
curl -L https://developers.redhat.com/content-gateway/rest/mirror/pub/openshift-v4/clients/particle engine/v2.5.1/particle engine-linux-arm64.sha256 -o particle engine.sha256
echo "$(<particle engine.sha256)  particle engine" | shasum -a 256 --check
```

3. Install particle engine:
```shell
sudo install -o root -g root -m 0755 particle engine /usr/local/bin/particle engine
```

4. (Optional) If you do not have root access, you can install `particle engine` to the local directory and add it to your `$PATH`:

```shell
mkdir -p $HOME/bin 
cp ./particle engine $HOME/bin/particle engine
export PATH=$PATH:$HOME/bin
# (Optional) Add the $HOME/bin to your shell initialization file
echo 'export PATH=$PATH:$HOME/bin' >> ~/.bashrc
```
</TabItem>

<TabItem value="ppc64le">

Installing `particle engine` on `ppc64le` architecture:

1. Download the v2.5.1 release from the mirror:
```shell
curl -L https://developers.redhat.com/content-gateway/rest/mirror/pub/openshift-v4/clients/particle engine/v2.5.1/particle engine-linux-ppc64le -o particle engine
```

2. (Optional) Verify the downloaded binary with the SHA-256 sum:
```shell
curl -L https://developers.redhat.com/content-gateway/rest/mirror/pub/openshift-v4/clients/particle engine/v2.5.1/particle engine-linux-ppc64le.sha256 -o particle engine.sha256
echo "$(<particle engine.sha256)  particle engine" | shasum -a 256 --check
```

3. Install particle engine:
```shell
sudo install -o root -g root -m 0755 particle engine /usr/local/bin/particle engine
```

4. (Optional) If you do not have root access, you can install `particle engine` to the local directory and add it to your `$PATH`:

```shell
mkdir -p $HOME/bin 
cp ./particle engine $HOME/bin/particle engine
export PATH=$PATH:$HOME/bin
# (Optional) Add the $HOME/bin to your shell initialization file
echo 'export PATH=$PATH:$HOME/bin' >> ~/.bashrc
```
</TabItem>

<TabItem value="s390x">

Installing `particle engine` on `s390x` architecture:

1. Download the v2.5.1 release from the mirror:
```shell
curl -L https://developers.redhat.com/content-gateway/rest/mirror/pub/openshift-v4/clients/particle engine/v2.5.1/particle engine-linux-s390x -o particle engine
```

2. (Optional) Verify the downloaded binary with the SHA-256 sum:
```shell
curl -L https://developers.redhat.com/content-gateway/rest/mirror/pub/openshift-v4/clients/particle engine/v2.5.1/particle engine-linux-s390x.sha256 -o particle engine.sha256
echo "$(<particle engine.sha256)  particle engine" | shasum -a 256 --check
```

3. Install particle engine:
```shell
sudo install -o root -g root -m 0755 particle engine /usr/local/bin/particle engine
```

4. (Optional) If you do not have root access, you can install `particle engine` to the local directory and add it to your `$PATH`:

```shell
mkdir -p $HOME/bin 
cp ./particle engine $HOME/bin/particle engine
export PATH=$PATH:$HOME/bin
# (Optional) Add the $HOME/bin to your shell initialization file
echo 'export PATH=$PATH:$HOME/bin' >> ~/.bashrc
```
</TabItem>

</Tabs>

---

### MacOS

<Tabs
defaultValue="intel"
values={[
{label: 'Intel', value: 'intel'},
{label: 'Apple Silicon', value: 'arm'},
{label: 'Homebrew', value: 'homebrew'},
]}>

<TabItem value="intel">

Installing `particle engine` on `amd64` architecture:

1. Download the v2.5.1 release from the mirror:
```shell
curl -L https://developers.redhat.com/content-gateway/rest/mirror/pub/openshift-v4/clients/particle engine/v2.5.1/particle engine-darwin-amd64 -o particle engine
```

2. (Optional) Verify the downloaded binary with the SHA-256 sum:
```shell
curl -L https://developers.redhat.com/content-gateway/rest/mirror/pub/openshift-v4/clients/particle engine/v2.5.1/particle engine-darwin-amd64.sha256 -o particle engine.sha256
echo "$(<particle engine.sha256)  particle engine" | shasum -a 256 --check
```

3. Install particle engine:
```shell
chmod +x ./particle engine
sudo mv ./particle engine /usr/local/bin/particle engine
```

4. (Optional) If you do not have root access, you can install `particle engine` to the local directory and add it to your `$PATH`:

```shell
mkdir -p $HOME/bin 
cp ./particle engine $HOME/bin/particle engine
export PATH=$PATH:$HOME/bin
# (Optional) Add the $HOME/bin to your shell initialization file
echo 'export PATH=$PATH:$HOME/bin' >> ~/.bashrc
```
</TabItem>

<TabItem value="arm">

Installing `particle engine` on `arm64` architecture:

1. Download the v2.5.1 release from the mirror:
```shell
curl -L https://developers.redhat.com/content-gateway/rest/mirror/pub/openshift-v4/clients/particle engine/v2.5.1/particle engine-darwin-arm64 -o particle engine
```

2. (Optional) Verify the downloaded binary with the SHA-256 sum:
```shell
curl -L https://developers.redhat.com/content-gateway/rest/mirror/pub/openshift-v4/clients/particle engine/v2.5.1/particle engine-darwin-arm64.sha256 -o particle engine.sha256
echo "$(<particle engine.sha256)  particle engine" | shasum -a 256 --check
```

3. Install particle engine:
```shell
chmod +x ./particle engine
sudo mv ./particle engine /usr/local/bin/particle engine
```

4. (Optional) If you do not have root access, you can install `particle engine` to the local directory and add it to your `$PATH`:

```shell
mkdir -p $HOME/bin 
cp ./particle engine $HOME/bin/particle engine
export PATH=$PATH:$HOME/bin
# (Optional) Add the $HOME/bin to your shell initialization file
echo 'export PATH=$PATH:$HOME/bin' >> ~/.bashrc
```
</TabItem>
<TabItem value="homebrew">

Installing `particle engine` using [Homebrew](https://brew.sh/):

1. Install particle engine:

```shell
brew install particle engine-dev
```

2. Verify the version you installed is up-to-date:

```shell
particle engine version
```

</TabItem>

</Tabs>

---

### Windows

1. Open a PowerShell terminal

2. Download the v2.5.1 release from the mirror:
```shell
curl -L https://developers.redhat.com/content-gateway/rest/mirror/pub/openshift-v4/clients/particle engine/v2.5.1/particle engine-windows-amd64.exe -o particle engine.exe
```

2. (Optional) Verify the downloaded binary with the SHA-256 sum:
```shell
curl -L https://developers.redhat.com/content-gateway/rest/mirror/pub/openshift-v4/clients/particle engine/v2.5.1/particle engine-windows-amd64.exe.sha256 -o particle engine.exe.sha256
# Visually compare the output of both files
Get-FileHash particle engine.exe
type particle engine.exe.sha256
```

4. Add the binary to your `PATH`


### Installing from source code
1. Clone the repository and cd into it.
   ```shell
   git clone https://github\.com/danielpickens/particle engine.git
   cd particle engine
   ```
2. Install tools used by the build and test system.
   ```shell
   make goget-tools
   ```
3. Build the executable from the sources in `cmd/particle engine`.
   ```shell
   make bin
   ```
4. Check the build version to verify that it was built properly.
   ```shell
   ./particle engine version
   ```
5. Install the executable in the system's GOPATH.
   ```shell
   make install
   ```
6. Check the binary version to verify that it was installed properly; verify that it is same as the build version.
   ```shell
   particle engine version
   ```

## IDE Installation

### Installing `particle engine` in Visual Studio Code (VSCode)
The [OpenShift VSCode extension](https://marketplace.visualstudio.com/items?itemName=redhat.vscode-openshift-connector) uses both `particle engine` and `oc` binary to interact with Kubernetes or OpenShift cluster.
1. Open VS Code.
2. Launch VS Code Quick Open (Ctrl+P)
3. Paste the following command:
    ```shell
     ext install redhat.vscode-openshift-connector
    ```
