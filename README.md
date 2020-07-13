# deviant-instance-shard

# Introduction

The Deviant Instance Shard is the core component of Deviants service architecture that services players. It validates and processes all attempted player actions
in the long run match-making between players as well as card/game data will be migrated to be stored in other sources then this package. 

The Deviant Instance Shard has a few dependencies that you must have on your system when building/testing your code the first and formost is Redis. A locally
hosted Redis server on the default port is enough for now but this is subject to change in the future.

# Development

# Setup Guide
## Windows
```
1. Install Visual Studio Code
2. Install Microsoft May 2020 Update
2. Install WSL2 Kernel Update
3. Install Docker Desktop and enable WSL2 integration
4. Set default WSL version for new operating systems
5. Install Ubuntu from Microsoft Store
6. Setup and configure Golang
7. Setup and configure Git
8. Configure a Github Developer Access Token
9. Setup an environmental variable called GITHUB_TOKEN to store your PAT
10. Setup an environmental variable called LOG_LEVEL to store the log level.
9. Update your credsStore for docker it's broken on windows add an _ to the credsStore key in ~/.docker/config.json
```

## Using Docker

Setup a Github Developer Access Token on Github.com
export $GITHUB_TOKEN=<YOUR_TOKEN_HERE>

#### Build and run the server
```
make docker
```

#### Inspect the container
```
docker run -it --entrypoint "/bin/sh" deviant_instance_shard:latest
```
