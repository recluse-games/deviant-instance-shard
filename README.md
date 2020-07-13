# deviant-instance-shard

# Introduction

The Deviant Instance Shard is the core component of Deviants service architecture that services players. It validates and processes all attempted player actions
in the long run match-making between players as well as card/game data will be migrated to be stored in other sources then this package. 

The Deviant Instance Shard has a few dependencies that you must have on your system when building/testing your code the first and formost is Redis. A locally
hosted Redis server on the default port is enough for now but this is subject to change in the future.

# Setup Guide

```
Install the make toolchain on your system

make
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
