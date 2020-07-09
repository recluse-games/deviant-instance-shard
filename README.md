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

#### Configure access to GitHub private repos

1. Create an SSH keypair
    ```
    ssh-keygen
    ```

2. Navigate to the [deviant-protobuf keys settings](https://github.com/recluse-games/deviant-protobuf/settings/keys)

3. Click `Add deploy key`

4. Copy your public key into the provided field

5. *HACK* For now, copy your keys into the root of the `deviant-instance-shard` project

#### Build and run the server
```
docker-compose -p deviant up --build
```

#### Inspect the container
```
docker run -it --entrypoint "/bin/sh" deviant_instance_shard:latest
```
