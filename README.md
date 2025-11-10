# SaaS Management Platform

This document provides instructions for setting up, running, and managing the project.

## Project Setup

### 1. Prerequisites

First, install the required tools for your operating system.

-   **Docker & Docker Compose**
	-   [Installation Guide](https://docs.docker.com/get-started/get-docker/)

-   **GNU Make**
	-   **macOS:** `brew install make`
	-   **Ubuntu:** `sudo apt install make`
	-   **Windows:** Follow online installation guides.

### 2. Secrets Configuration

The `secrets` directory at the project root contains the following files:

-   `buf_token`: Buf has a rate limit for fetching remote plugins. If you encounter this limit, register for an account at [buf.build](https://buf.build/) and create a token to bypass it.
	-   *Example:* `a5757d9dc362757a87854da2ae88bcdcd4da60f7ff3bd4786830926ea88ef787`
-   `sops_private_key.asc`: This project uses [sops](https://github.com/getsops/sops) to manage encrypted environment variables. You need a private key for decryption. Obtain this from a team member. A sample key is provided for initial setup.
-   `sops_public_key`: The public key corresponding to the private key, used for encryption.

### 3. Install Dependencies

Install Go tools and set up necessary permissions by running:

```sh
make setup
```

## Development

### Running the Application

-   **Start all services:** `make up-all`
-   **Stop all services:** `make down-all`
-   **Generate all files** (tidy, wire, mock, proto, ent): `make gen`

See the `Makefile` for a complete list of commands.

### Generating Code

#### Go Mocks

To generate mock files for an interface, add the following `go:generate` directive to your Go source file:

```go
//go:generate sh -c "mockgen -source=$GOFILE -destination=mock/$(basename $GOFILE .go)_test.go -package=mock"
```

### Managing Environment Variables

The `.env` file is automatically created from the encrypted `.env.enc` file when you run the application.

- To update environment variables, modify your local `.env` file and then synchronize the changes back to the encrypted file. This prevents losing newly added values.

```sh
# Replace <your_service> with the actual service name
make -C apps/<your_service>-env-encrypt
```

For example: `make -C apps/saas_mngt_service/ env-encrypt`

**Note:** `.env.enc` will change every time you run this command.

- To test your new local environment, you can create a new file called `.env.local` and put your overridden environment variables in that file. The system will merge environment variables from `.env` and `.env.local`.

### Creating a New Application

To scaffold a new application within the project, run the following script and follow the prompts:

```sh
bash .template/init_app_go.sh
```

## Important Notes

### Security

In a real-world project, the `secrets` directory **MUST** be added to `.gitignore`. Private keys and other secrets should never be committed to the codebase.

### Production Environment

In production, environment variables are injected directly from a secret management service (e.g., AWS Secrets Manager, HashiCorp Vault). `sops` is used for local development only.

## Appendix: Useful Commands

### Generating a New `sops` Keypair
1. Install gpg
	-   **macOS:** `brew install gpg`
	-   **Ubuntu:** `sudo apt install gnupg`
	-   **Windows:** Follow online installation guides.

2.  Generate a new GPG key:
	```sh
	gpg --full-generate-key
	```
	-   Choose `(1) RSA and RSA`.
	-   Set key size to `4096`.
	-   Set an expiration or choose `0` for no expiration.
	-   Provide your user information and skip the passphrase.

3.  Note the public key fingerprint from the output (e.g., `49031300F4AF150110DCECAD347C46C2BFE6D611`).

4.  Export the private key:
	```sh
	gpg --armor --export-secret-keys YOUR_PUBLIC_KEY_FINGERPRINT > sops_private_key.asc
	```

5.  Place the generated `sops_private_key.asc` in the `secrets` directory and update `sops_public_key` with the fingerprint.