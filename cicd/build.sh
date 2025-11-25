#!/bin/bash

APP_NAME=${APP_NAME:-saas_mngt_service}

docker build -t $APP_NAME -f apps/$APP_NAME/Dockerfile .

# Test run (local only)
# docker run -it --rm -p 8080:8080 $APP_NAME