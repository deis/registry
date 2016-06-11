# Deis Registry v2

[![Build Status](https://travis-ci.org/deis/registry.svg?branch=master)](https://travis-ci.org/deis/registry)
[![Go Report Card](http://goreportcard.com/badge/deis/registry)](http://goreportcard.com/report/deis/registry)
[![Docker Repository on Quay](https://quay.io/repository/deisci/registry/status "Docker Repository on Quay")](https://quay.io/repository/deisci/registry)


Deis (pronounced DAY-iss) is an open source PaaS that makes it easy to deploy and manage
applications on your own servers. Deis builds on [Kubernetes](http://kubernetes.io/) to provide
a lightweight, [Heroku-inspired](http://heroku.com) workflow.

We welcome your input! If you have feedback, please submit an [issue][issues]. If you'd like to participate in development, please read the "Development" section below and submit a [pull request][prs].

# About

The registry is a [Docker registry](https://docs.docker.com/registry/) component for use in Kubernetes. While it's intended for use inside of the Deis open source [PaaS](https://en.wikipedia.org/wiki/Platform_as_a_service), it's flexible enough to be used as a standalone pod on any Kubernetes cluster, brocephus.

If you decide to use this component standalone, you can host your own Docker registry in your own Kubernetes cluster.

The Docker image that this repository builds is based on [Alpine Linux](http://www.alpinelinux.org/) and uses [the Deis fork](https://github.com/deis/distribution) of [the official Docker v2 registry code](https://github.com/docker/distribution).

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

Before deploying your custom image you must update the container image specification in the pod manifest. This file is found at `contrib/kubernetes/manifests/registry-rc.yaml`:

```yaml
        - name: registry
          image: quay.io/youruser/registry:git-f5c7dc3
          env:
            - name: REGISTRY_STORAGE_DELETE_ENABLED
              value: "true"
```

By default registry uses the filesystem as the storage medium. To use a custom object store like s3, gcs or azure:
- First provide the details required for authenticating to object store in base64 format by updating the secret file which can be found at `contrib/kubernetes/manifests/registry-{STORAGE_TYPE}-secret.yaml`.
- Update the storage type and secret to be used in the pod manifest. This file is found at `contrib/kubernetes/manifests/registry-rc.yaml`:
```yaml
        - name: REGISTRY_STORAGE
          value: filesystem

        - name: registry-creds
          secret:
            secretName: fs-keyfile
```
- Set the STORAGE_TYPE environment variable.
```
$ export STORAGE_TYPE = {s3/gcs/azure}
```

Once updated, deploy the registry to your kubernetes cluster with:

```
$ make deploy
```

After a while, you should see one pod up with one running:

```
NAME                  READY     STATUS    RESTARTS   AGE
registry-6wy8o        1/1       Running   0          32s
```

You can then interact with this pod as you would with any other Kubernetes pod:

```
$ kubectl logs -f registry-6wy8o
$ kubectl exec -it registry-6wy8o sh
```

## License

© 2014, 2015, 2016 Engine Yard, Inc.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

[issues]: https://github.com/deis/registry/issues
[prs]: https://github.com/deis/registry/pulls
