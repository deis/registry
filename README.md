# Deis Registry v2

[![Build Status](https://ci.deis.io/job/registry/badge/icon)](https://ci.deis.io/job/registry)
[![Go Report Card](https://goreportcard.com/badge/github.com/deis/registry)](https://goreportcard.com/report/github.com/deis/registry)
[![Docker Repository on Quay](https://quay.io/repository/deisci/registry/status "Docker Repository on Quay")](https://quay.io/repository/deisci/registry)


Deis (pronounced DAY-iss) is an open source PaaS that makes it easy to deploy and manage
applications on your own servers. Deis builds on [Kubernetes](http://kubernetes.io/) to provide
a lightweight, [Heroku-inspired](http://heroku.com) workflow.

We welcome your input! If you have feedback, please submit an [issue][issues]. If you'd like to participate in development, please read the "Development" section below and submit a [pull request][prs].

# About

The registry is a [Docker registry](https://docs.docker.com/registry/) component for use in Kubernetes. While it's intended for use inside of the Deis open source [PaaS](https://en.wikipedia.org/wiki/Platform_as_a_service), it's flexible enough to be used as a standalone pod on any Kubernetes cluster.

If you decide to use this component standalone, you can host your own Docker registry in your own Kubernetes cluster.

The Docker image that this repository builds is based on [the official Docker v2 registry image](https://github.com/docker/distribution).

# Development

The Deis project welcomes contributions from all developers. The high level process for development matches many other open source projects. See below for an outline.

* Fork this repository
* Make your changes
* Submit a pull request (PR) to this repository with your changes, and unit tests whenever possible.
	* If your PR fixes any issues, make sure you write Fixes #1234 in your PR description (where #1234 is the number of the issue you're closing)
* The Deis core contributors will review your code. After each of them sign off on your code, they'll label your PR with LGTM1 and LGTM2 (respectively). Once that happens, the contributors will merge it

## Deploying

If you want to use the latest registry image built by they Deis team you can simply start a registry via `make deploy`.

If however, you want to build and use a custom image see the instructions below.

## Build and Deploy

To build a dev release of this image, you will also need a registry to hold the custom images. This can be your own registry, Dockerhub, or Quay.


First, configure your environment to point to the registry location.

```console
$ export DEIS_REGISTRY=myregistry.com:5000  # or quay.io, if using Dockerhub, leave this unset
$ export IMAGE_PREFIX=youruser/             # if using Quay or Dockerhub
```

To build and push the image run:

```console
$ make docker-build docker-push
```

To deploy the image via patching the registry deployment run:

```console
$ make deploy
```

[issues]: https://github.com/deis/registry/issues
[prs]: https://github.com/deis/registry/pulls
