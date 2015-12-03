# Registry

A Docker image registry for use in the Deis open source PaaS.

This Docker image is based on the official docker distribution image.

Please add any [issues](https://github.com/deis/registry/issues) you find with this software to the [Distribution Project](https://github.com/docker/distribution).


## Deploying

To build a dev release of this image, you will also need your own registry, but DockerHub or Quay will do fine here. To build, run:

```
$ export DEIS_REGISTRY=myregistry.com:5000
$ make docker-build docker-push
```

This will compile the Docker image and push it to your registry.

After that, run

```
$ make deploy
```

Which will deploy the component to kubernetes. After a while, you should see one pod up with one running:

```
NAME                  READY     STATUS    RESTARTS   AGE
registry-6wy8o        1/1       Running   0          32s
```

You can then query this pod as you would with any other Kubernetes pod:

```
$ kubectl logs -f registry-6wy8o
$ kubectl exec -it registry-6wy8o psql
```

## License

© 2014 Engine Yard, Inc.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
