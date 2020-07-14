FROM golang:alpine AS builder
LABEL maintainer="The deviant authors <deviant-dev@recluse-games.com>"

RUN apk update && apk add --no-cache git openssh make

# Take in the GITHUB_TOKEN from the compose environment.
ARG GITHUB_TOKEN
ARG GITHUB_USER
ENV TOKEN=${GITHUB_TOKEN}
ENV USER=${GITHUB_USER}

# Setup GIT related configuration to work in Docker + Private Go Repository
RUN echo "machine github.com login $GITHUB_USER password $GITHUB_TOKEN" > $HOME/.netrc
RUN GOCACHE=OFF
ENV GOPRIVATE=github.com/recluse-games
RUN export GIT_TERMINAL_PROMPT=1

WORKDIR /go/src/github.com/recluse-games/deviant-instance-shard
COPY . .

RUN make

FROM scratch
COPY --from=builder /go/src/github.com/recluse-games/deviant-instance-shard/bin/deviant-instance-shard /go/bin/deviant-instance-shard
ENTRYPOINT ["/go/bin/deviant-instance-shard"]
EXPOSE 50051
