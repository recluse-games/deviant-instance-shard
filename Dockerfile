FROM golang:alpine AS builder
LABEL maintainer="The deviant authors <deviant-dev@recluse-games.com>"

# Take in the GITHUB_TOKEN from the compose environment.
ARG GITHUB_TOKEN

RUN apk update && apk add --no-cache git openssh make redis
RUN GOCACHE=OFF

# Setup GIT related configuration to work in Docker + Private Go Repository
RUN export GIT_TERMINAL_PROMPT=1
ENV GOPRIVATE=github.com/recluse-games/*
RUN git config --global url."https://$GITHUB_TOKEN:x-oauth-basic@github.com".insteadOf "https://github.com"

WORKDIR /go/src/github.com/recluse-games/deviant-instance-shard
COPY . .

RUN make

FROM scratch
COPY --from=builder /go/src/github.com/recluse-games/deviant-instance-shard/bin/deviant-instance-shard /go/bin/deviant-instance-shard
ENTRYPOINT ["/go/bin/deviant-instance-shard"]
EXPOSE 50051
