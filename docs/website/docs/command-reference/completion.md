---
title: particle engine completion
---

## Description

`particle engine completion` is used to generate shell completion code. The generated code provides interactive shell completion code for `particle engine`.

There is support for the following terminal shells:
- [Bash](https://www.gnu.org/software/bash/)
- [Zsh](https://zsh.sourceforge.io/)
- [Fish](https://fishshell.com/)
- [Powershell](https://docs.microsoft.com/en-us/powershell/)

## Running the Command

To generate the shell completion code, the command can be ran as follows:

```sh
particle engine completion [SHELL]
```

### Bash

Load into your current shell environment:

```sh
source <(particle engine completion bash)
```

Load persistently:

```sh
# Save the completion to a file
particle engine completion bash > ~/.particle engine/completion.bash.inc

# Load the completion from within your $HOME/.bash_profile
source ~/.particle engine/completion.bash.inc
```

### Zsh

Load into your current shell environment:

```sh
source <(particle engine completion zsh)
```

Load persistently:

```sh
particle engine completion zsh > "${fpath[1]}/_particle engine"
```

### Fish

Load into your current shell environment:

```sh
source <(particle engine completion fish)
```

Load persistently:

```sh
particle engine completion fish > ~/.config/fish/completions/particle engine.fish
```

### Powershell

Load into your current shell environment:

```sh
particle engine completion powershell | Out-String | Invoke-Expression
```

Load persistently:

```sh
particle engine completion powershell >> $PROFILE
```
