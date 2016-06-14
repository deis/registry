FROM quay.io/deis/base:0.2.0

COPY . /

RUN apt-get update \
	&& apt-get install -y \
		git \
		golang \
		make \
		--no-install-recommends \
	&& export DOCKER_REGISTRY_TAG=deis \
		DOCKER_REGISTRY_REPO=https://github.com/deis/distribution.git \
		DISTRIBUTION_DIR=/go/src/github.com/docker/distribution \
	&& export GOPATH=/go:$DISTRIBUTION_DIR/Godeps/_workspace \
	&& git clone -b $DOCKER_REGISTRY_TAG --single-branch $DOCKER_REGISTRY_REPO $DISTRIBUTION_DIR \
	&& cd $DISTRIBUTION_DIR \
	&& make binaries \
	&& cp bin/* /bin/ \
	&& rm -rf /go \
	&& apt-get remove -y --auto-remove --purge \
		git \
		golang \
		make \
	&& apt-get clean \
	&& rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* /usr/share/man /usr/share/doc

VOLUME ["/var/lib/registry"]
CMD ["/opt/registry/sbin/registry"]
EXPOSE 5000

ENV WORKFLOW_RELEASE 2.0.0
