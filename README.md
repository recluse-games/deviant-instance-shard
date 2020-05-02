# deviant-instance-shard

# Setup Guide

```
go build
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
