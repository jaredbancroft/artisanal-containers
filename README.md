# artisanal-containers

A vscode devcontainer for running Liz Rice's amazing [containers from scratch](https://github.com/lizrice/containers-from-scratch).

Requires `--privileged` (for rw cgroups) and `--cap-add=SYS_ADMIN` (for clone namespaces) runArgs in devcontainer.json to be able to run your container from scratch within the devcontainer.

This has only been tested with OSX 10.15.7 and Docker Desktop 3.2.2. I have no idea if this will work with WSL.

## Requirements
1. A computer
1. Docker for your computer
1. [vscode with remote containers](https://code.visualstudio.com/docs/remote/containers)

## How to use this
1. Clone this repo
1. View->Command Palette...
1. Remote-Containers: Open Workspace in Container...
1. go run main.go run /bin/bash
1. :(){:|:&};:
