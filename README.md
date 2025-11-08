# README

### Setup Project

1. Install Docker & Docker Compose

Link: https://docs.docker.com/get-started/get-docker/

2. Install Make
- For Mac: `brew install make`
- For Ubuntu: `sudo apt install make`
- For Windows: Please search online for installation instructions

3. Install gpg (gnupg): using for encrypt/decrypt env
- For Mac: `brew install gpg`
- For Ubuntu: `sudo apt install gnupg`
- For Windows: Please search online for installation instructions

3. Look folder `secrets` at root, it include 3 files
- buf_token: Buf have rate limit to get plugin from remote. If you hit a rate limit, you need to register account and create token via url https://buf.build/ to by pass i

e.g: a5757d9dc362757a87854da2ae88bcdcd4da60f7ff3bd4786830926ea88ef787

- sops_private_key.asc:

We are using https://github.com/getsops/sops to encrypt/decrypt env key. We need private key to do it. Ask your team mate to get private key file. If you don't care about security env, just use my private key first in the folder.

- sops_public_key

4. Install requirement Go tool and setup permission

`make setup`

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

## Useful setup

### Generate sops secret and private key
Use can generate private by using this command:
```
gpg --full-generate-key

Choose RSA
Use 4096
Skip generate password
```

You can see a public key, like: `49031300F4AF150110DCECAD347C46C2BFE6D611`

Next, you need to export private key
`gpg --armor --export-secret-keys 49031300F4AF150110DCECAD347C46C2BFE6D611 > sops_private_key.asc`

Finally, create two file in secrets folder to store them:
- sops_private_key.asc
- sops_public_key