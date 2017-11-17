Build and Release with Google Container Builder
===============================================

Initial setup
-------------

1. Install Container Builder Local instructions in 
https://github.com/GoogleCloudPlatform/container-builder-local

Note: if you downloaded the binaries, you will need to create an alias in
your shell everytime you want to use it.

Example:
```shell
alias container-builder-local=/home/$USER/Downloads/container-builder-local_linux_amd64-v0.2.5
```

Build and Release
-----------------

1. Login in gcloud and select project
```shell
gcloud auth login k8s.production.user@gmail.com
gcloud config set project google-containers
```

2. Run the local build and push image to GCR.

* For test use:
```shell
container-builder-local --dryrun=false --config cloudbuild.yaml --substitutions=_PREFIX="stage/",_VERSION="$(date -I)" .
```

  Image will be available locally only with name:
  gcr.io/google_containers/stage/clusterregistry:2017-11-17

  Binaries will be store in Google Cloud Storage:
  gs://kubernetes-release/addons/stage/crinit/2017-11-17/

* For official releases:
```shell
container-builder-local --dryrun=false --config cloudbuild.yaml --substitutions=_PREFIX="",_VERSION="v1.0.0" --push .
```

  Image will be push to Google Cloud Registry with name:
  gcr.io/google_containers/clusterregistry:v1.0.0

  Binaries will be store in Google Cloud Storage:
  gs://kubernetes-release/addons/crinit/v1.0.0/

Where:
* **_PREFIX** is the path of the image and binaries, set empty for official releases.

* **_VERSION** is the string to identify the version we are building.

  Note: We could use git tags and git describe to get the version for every build.

* **--push** this flag will let local cloud builder to push the image to the GCR. Remove the flag if you want to test only.

