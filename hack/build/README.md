Cloud Registry in Google Cloud Container Builder
================================================

The container is based on the official Google Cloud Container Builder bazel.
The container is name 'clusterregistry_builder'.

We cannot use use default 'gcr.io/cloud-builders/bazel' because the script
'hack/update-codegen.sh' requires bazel, xmllint and rysnc.

Build Builder Container
-----------------------

To build a new builder image, inside the 'hack/build' dir, use:

```bash
gcloud container builds submit --config cloudbuild.yaml .
```

You should build a new builder every time you update 'hack/build/*' files.

Multiple Builder Containers
---------------------------

You will need multiple versions when you need to build older clusterregistry
code that use a different bazel version.

In case you need to support multiple builder versions, you can versioning the
builder containers. You can use the git branch name as the tag of the image.

Don't forget to udate the root 'cloudbuild.yaml' to use the desired builder
image and tag.


