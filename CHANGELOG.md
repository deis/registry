### v2.0.0-rc2 -> v2.0.0

#### Documentation

- [`afc9d2a`](https://github.com/deis/registry/commit/afc9d2a54a1248aef380de5b99ce8fd2ce0b0ad8) CHANGELOG.md: add entry for v2.0.0-rc2

#### Maintenance

- [`9cd3013`](https://github.com/deis/registry/commit/9cd3013ba05160ffda907217febcb3e351a5e57f) README.md: remove beta status

### v2.0.0-rc1 -> v2.0.0-rc2

### v2.0.0-beta4 -> v2.0.0-rc1

#### Maintenance

 - [`33c45cb`](https://github.com/deis/registry/commit/33c45cbaf69311c4f43aaf84cf993514adb181db) Dockerfile: Refactor image to use ubuntu-slim

### v2.0.0-beta3 -> v2.0.0-beta4

#### Documentation

 - [`7d09014`](https://github.com/deis/registry/commit/7d09014690690a271501c6e49372ab6a1eda918c) CHANGELOG.md: update for v2.0.0-beta3
 - Also some other random stuff here

### v2.0.0-beta2 -> v2.0.0-beta3

#### Fixes

 - [`8af10a2`](https://github.com/deis/registry/commit/8af10a28d73704e8817edb98aa36be35dd0cc3b2) Makefile: remove includes.mk

#### Maintenance

 - [`96e38cf`](https://github.com/deis/registry/commit/96e38cf90a22428668dabd8bf5ddfb365b5a55c6) .travis.yml: Deep six the travis -> jenkins webhooks

### v2.0.0-beta1 -> v2.0.0-beta2

#### Features

 - [`f289d97`](https://github.com/deis/registry/commit/f289d97e0fb5a82e7911e17e6e1f58d9c83f0436) _scripts: add CHANGELOG.md and generator script

#### Fixes

 - [`0bf4735`](https://github.com/deis/registry/commit/0bf473551e02d976d34294641ca154ae99e5a1b4) storage: make changes to support the object storage secret
 - [`586f97c`](https://github.com/deis/registry/commit/586f97c420fc77936d9421d32901d7a612f1f192) perf: fix the high cpu utilization of the registry

#### Maintenance

 - [`93d8fb2`](https://github.com/deis/registry/commit/93d8fb23c9b8db6588a2ce25f3bdab814c74b87d) Dockerfile: remove include_gcs build tag
 - [`158a686`](https://github.com/deis/registry/commit/158a686c3d60ed542b74a92519b18bdfa35c884e) Makefile: update go-dev to 0.10.0 and compress registry binary

### 2.0.0-alpha -> v2.0.0-beta1

#### Features

 - [`d019479`](https://github.com/deis/registry/commit/d019479ff95f9639e9981721081906aa28108959) Makefile: enable immutable (git-based
 - [`8924755`](https://github.com/deis/registry/commit/8924755ac2b90940ff4a5ad4348c06c252d6a32f) .travis.yml: have this job notify its sister job in Jenkins
 - [`da26fdc`](https://github.com/deis/registry/commit/da26fdcb907896c9528e862e274058458057d315) README.md: add badge for Travis CI build

#### Fixes

 - [`3747ba3`](https://github.com/deis/registry/commit/3747ba3aade2095222ed240b0515bb64a74df57d) minio: add support for minio
 - [`114b892`](https://github.com/deis/registry/commit/114b8927759777fb9a70d101bba7f2a8d86ff4d9) deploy.sh: add trailing slash to quay.io/

#### Maintenance

 - [`0699184`](https://github.com/deis/registry/commit/0699184961b886ceda6184c1896a77540962aead) Dockerfile: update alpine to 3.3
 - [`00dcbfe`](https://github.com/deis/registry/commit/00dcbfe55001805d37d0ed3affd266c442d25e3a) Dockerfile: update docker/distribution to 2.2.1
 - [`7b93259`](https://github.com/deis/registry/commit/7b932598b0dc57cacff2b3d0446d979c78a77cff) release: bump version to v2-beta

### 2.0.0-alpha

#### Features

 - [`e72dfa1`](https://github.com/deis/registry/commit/e72dfa1840f088a7539e28f9cdcd5bfaf5c15149) (all): use alpine:3.2 for image base

#### Fixes

 - [`01b8641`](https://github.com/deis/registry/commit/01b8641337eeaa45243d5fbdb01a899409e9b5e9) travis: connect deploy script correctly
 - [`570cfd6`](https://github.com/deis/registry/commit/570cfd6544b640dd03ee714a27ff391fd3ee5bc0) Makefile/deploy: use standard deploy.sh location

#### Documentation

 - [`e162fda`](https://github.com/deis/registry/commit/e162fdae1f116d302bb0774ac6efd2d000012fb9) readme: add deploying documentation

#### Maintenance

 - [`3c54099`](https://github.com/deis/registry/commit/3c540990c377c3ada665dc9c99884d8671b900d3) release: set version and lock to deis registry
