FROM quay.io/deis/base:v0.3.5

COPY . /

RUN buildDeps='git golang make'; \
    apt-get update && \
    apt-get install -y --no-install-recommends \
        $buildDeps && \
	export DOCKER_REGISTRY_TAG=deis \
		   DOCKER_REGISTRY_REPO=https://github.com/deis/distribution.git \
		   DISTRIBUTION_DIR=/go/src/github.com/docker/distribution && \
	export GOPATH=/go:$DISTRIBUTION_DIR/Godeps/_workspace && \
	git clone -b $DOCKER_REGISTRY_TAG --single-branch $DOCKER_REGISTRY_REPO $DISTRIBUTION_DIR && \
	cd $DISTRIBUTION_DIR && \
	make binaries && \
	cp bin/* /bin/ && \
	rm -rf /go && \
    # cleanup
    apt-get purge -y --auto-remove $buildDeps && \
    apt-get autoremove -y && \
    apt-get clean -y && \
    # package up license files if any by appending to existing tar
    COPYRIGHT_TAR='/usr/share/copyrights.tar'; \
    gunzip $COPYRIGHT_TAR.gz; tar -rf $COPYRIGHT_TAR /usr/share/doc/*/copyright; gzip $COPYRIGHT_TAR && \
    rm -rf \
        /usr/share/doc \
        /usr/share/man \
        /usr/share/info \
        /usr/share/locale \
        /var/lib/apt/lists/* \
        /var/log/* \
        /var/cache/debconf/* \
        /etc/systemd \
        /lib/lsb \
        /lib/udev \
        /usr/lib/x86_64-linux-gnu/gconv/IBM* \
        /usr/lib/x86_64-linux-gnu/gconv/EBC* && \
    bash -c "mkdir -p /usr/share/man/man{1..8}"

VOLUME ["/var/lib/registry"]
CMD ["/opt/registry/sbin/registry"]
EXPOSE 5000
