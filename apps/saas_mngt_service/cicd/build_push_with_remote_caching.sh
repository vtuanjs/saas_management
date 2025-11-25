#!/bin/bash

# Use this script to build and push the Docker image for the CI/CD flow
# DOCKER_REPOSITORY should be set in the CI/CD ENVIRONMENT

# In local testing, you need to check environment/docker-compose.yml to enable bazel-remote cache service

IMAGE_TAG=${IMAGE_TAG:-latest}
DOCKER_REPOSITORY="${DOCKER_REPOSITORY:-vantuan1303/saas_mngt_service:$IMAGE_TAG}"
REMOTE_CACHE="${REMOTE_CACHE:-http://bazel:bazel@localhost:8080}"

if [ -e ".build" ]; then
	echo "Removing existing .build..."
	rm -rf .build
fi

# You need to install Docker and Bazel before running this script.
bazel build --remote_cache=$REMOTE_CACHE --remote_upload_local_results //cmd:cmd

bazel build //cmd:image

bazel run //cmd:push -- --repository $DOCKER_REPOSITORY
