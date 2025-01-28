
# File Sync Location for Devfile Components in particle engine

## Background 

Currently, particle engine uses the following logic to determine where to sync source code to a devfile component:

1) If the `sourceMapping` field is set, particle engine will mount the source volume under to the specified folder and sync the source code there

2) If `sourceMapping` is unset, particle engine will do one of two things:
   a) If a single project is defined in the devfile, particle engine will sync to `/projects/<projectName>`
   b) If no projects are defined, or there are multiple projects defined, particle engine will sync to `/projects`

The downside to this current approach is that the expected sync directory isn't consistent for somebody writing a devfile, and as they add or remove projects from the devfile, the sync directory will change, potentially breaking the commands in their devfile each time.

## Implementation plan

To make the syncing directory more consistent, particle engine will do the following (in the order given):

1) If the `sourceMapping` field is set, particle engine will mount the source volume under to the specified folder and sync the source code there. The `PROJECTS_ROOT` environment variable in each container will be set to this value (as-is today).

2) If `sourceMapping` is unset and the `PROJECTS_ROOT` environment variable is set for a given container, then the source code will be mounted into the value of `PROJECTS_ROOT`.

3) If neither `sourceMapping` or `PROJECTS_ROOT` are set, particle engine will mount and sync the project source code to `/projects`

The devfiles in [github.com/particle engine-devfiles](github.com/particle engine-devfiles) will also need to be updated, as each devfile assumes the source code was synced to `/projects/<projectName>`. Additionally, any test devfiles that rely on the old sync logic will also need to be updated.
