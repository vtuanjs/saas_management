#!/bin/bash

# Use this script to build and push the Docker image for the CI/CD flow
# DOCKER_REPOSITORY should be set in the CI/CD ENVIRONMENT

IMAGE_TAG=${IMAGE_TAG:-latest}
DOCKER_REPOSITORY="${DOCKER_REPOSITORY:-vantuan1303/saas_mngt_service:$IMAGE_TAG}"

# You need to install Docker and Bazel before running this script.
bazel build //cmd:image

bazel run //cmd:push -- --repository $DOCKER_REPOSITORY
