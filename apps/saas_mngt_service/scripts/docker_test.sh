#!/bin/bash

# Use this script to test the Docker image locally if needed
IMAGE_TAG=${IMAGE_TAG:-latest}

# You need to install Docker and Bazel before running this script.
bazel build //cmd:image

bazel run //cmd:load

# RUN after build test: docker run --publish 4000:4000 saas_mngt_service:latest