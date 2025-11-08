#!/bin/bash

# Get current branch name
branch=$(./git.sh rev-parse --abbrev-ref HEAD)

# Add all changes
./git.sh add .

# Commit with a default message
read -p "Please input commit name: " commit_message
./git.sh commit -m "$commit_message"

# Pull latest changes from current branch
./git.sh pull origin "$branch"

# Push to current branch
./git.sh push origin "$branch"