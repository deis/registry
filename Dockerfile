FROM alpine:3.3

ENV DOCKER_REGISTRY_TAG=v2.3.0 \
    DOCKER_REGISTRY_REPO=https://github.com/docker/distribution.git \
    DOCKER_BUILDTAGS=include_gcs

# install registry binaries
RUN export DISTRIBUTION_DIR=/go/src/github.com/docker/distribution \
    && apk add --update-cache git go make \
    && git clone -b $DOCKER_REGISTRY_TAG --single-branch $DOCKER_REGISTRY_REPO $DISTRIBUTION_DIR \
    && cd $DISTRIBUTION_DIR \
    && export GOPATH=/go:$DISTRIBUTION_DIR/Godeps/_workspace \
    && make binaries \
    && cp bin/* /bin/ \
    && rm -rf /go \
    && apk del --purge git go make \
    && rm -rf /var/cache/apk/*

COPY rootfs/ /

# define the execution environment
VOLUME ["/var/lib/registry"]
EXPOSE 5000
ENTRYPOINT ["/bin/registry"]
CMD ["/etc/docker/registry/config.yml"]
