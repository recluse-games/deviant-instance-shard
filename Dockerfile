FROM golang:alpine AS builder
LABEL maintainer="The deviant authors <deviant-dev@recluse-games.com>"

RUN apk update && apk add --no-cache git openssh

# Configure access to recluse private repos. There is probably a better way to
# do this with GitHub actions. But this works for now for local builds.
ADD deviant_rsa /root/.ssh/id_rsa
ADD deviant_rsa.pub /root/.ssh/id_rsa.pub
RUN chmod 700 /root/.ssh/id_rsa
RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" > /root/.ssh/config
RUN git config --global url."git@github.com:".insteadOf https://github.com/
RUN ssh-keyscan -t rsa github.com >> /etc/ssh/ssh_known_hosts

WORKDIR /go/src/github.com/recluse-games/deviant-instance-shard
COPY . .
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o /go/bin/deviant-instance-shard
RUN rm -f *rsa*

FROM scratch
COPY --from=builder /go/bin/deviant-instance-shard /go/bin/deviant-instance-shard
ENTRYPOINT ["/go/bin/deviant-instance-shard"]
EXPOSE 50051
