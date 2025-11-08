# README

### Setup Project

1. Install Docker & Docker Compose

Link: https://docs.docker.com/get-started/get-docker/

2. Install Make

You need to install `Make`:

- For Mac: https://formulae.brew.sh/formula/make
- For Ubuntu:
```
sudo apt update
sudo apt install make
```
- For Windows: Please search online for installation instructions

3. Ask your teammate to download the `sops` folder

In this folder, you will find private_key.asc.

This key is used to encrypt/decrypt .env files.

4. Initialize Docker images and network

`make init`

### Running

- Run all services: `make up-all`
- Stop all services: `make down-all`
- Gen file (mod tidy + wire + mock + proto + ent + build): `make gen`


See more commands in the Makefile.

### Generating
#### With Golang Project
- Generate mock file
You need to add
//go:generate sh -c "mockgen -source=$GOFILE -destination=mock/$(basename $GOFILE .go)_test.go -package=mock"

### Managing Your ENV

The .env file is automatically created when you run the app, with values from .env.enc

If you need to change your .env, you must synchronize it to .env.enc to ensure you don't lose newly added values. Run `make <your_service>-gen-env-enc`. For example: `make saas-mgmt-gen-env-enc`

### Create new app
You need to run: `bash .template/init_app_go.sh` and following instruction

### Note

In a real project, the `sops` folder MUST be added to .gitignore.

We don't want to store it in the codebase.

### GO TO PRODUCTION NOTE

In production, we don't use `sops`, we just inject env directly from secret to environment.
Keyword: envFrom secretRef
