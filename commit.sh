#!/bin/bash

# Get current branch name
branch=$(git rev-parse --abbrev-ref HEAD)

# Add all changes
git add .

# Commit with a default message
read -p "Please input commit name: " commit_message
git commit -m "$commit_message"

# Pull latest changes from current branch
git pull origin "$branch"

# Push to current branch
git push origin "$branch"