# Registry

[![Build Status](https://travis-ci.org/deis/registry.svg?branch=master)](https://travis-ci.org/deis/registry)

A Docker image registry for use in the Deis open source PaaS.

This Docker image is based on alpine linux and uses the official Docker V2 registry code.

Deis (pronounced DAY-iss) is an open source PaaS that makes it easy to deploy and manage
applications on your own servers. Deis builds on [Kubernetes](http://kubernetes.io/) to provide
a lightweight, [Heroku-inspired](http://heroku.com) workflow.

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

By default registry uses the filesystem as the storage medium. To use a custom object store like s3 or gcs:
- First provide the details required for authenticating to object store in base64 format by updating the secret file which can be found at `contrib/kubernetes/manifests/registry-{STORAGE_TYPE}-secret.yaml`.
- Set the STORAGE_TYPE environment variable.
```
$ export STORAGE_TYPE = {s3/gcs}
```
- Update the secret to be used in the pod manifest. This file is found at `contrib/kubernetes/manifests/registry-rc.yaml`:
```yaml
        - name: registry-creds
          secret:
            secretName: fs-keyfile
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

© 2014 Engine Yard, Inc.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
