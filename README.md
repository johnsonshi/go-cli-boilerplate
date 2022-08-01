# go-cli-boilerplate

Boilerplate to kickstart creating a Go command-line tool.

## Getting Started

### Cloning the Repo

Click the "Use this template" button.

Alternatively, create a new directory and clone the repo:

```bash
mkdir new-directory
cd new-directory
curl -fsSL https://github.com/johnsonshi/go-cli-boilerplate/archive/main.tar.gz | tar -xz --strip-components=1
```

### Initialize Go Modules

Initialize `go.mod` and `go.sum`.

_Note_: Replace `<your-github-username>` and `<your-github-repo-name>` before running the following command:

```bash
rm go.mod go.sum
go mod init github.com/<your-github-username>/<your-github-repo-name>
go mod tidy
```

### Initialize Git

```bash
git init
git add -A
git status
git commit -s -m "Initial commit"
```

## Building and Running the CLI Locally

Build the command-line tool with the `make` command using the [`Makefile`](./Makefile).
This builds an executable binary at `./bin/command`.

```bash
make build-cli
```

To run the command-line tool:

```bash
./bin/command
```

## Modifying the CLI

The command-line tool can be modified to suit your needs through changes to [`./cmd/cli/root.go`](./cmd/cli/root.go) and [`./cmd/cli/subcommand.go`](./cmd/cli/subcommand.go).

## GitHub Actions Workflows

There are several [GitHub Actions Workflows](https://docs.github.com/en/actions) under the [`.github/workflows`](.github/workflows/) directory.

These should kickstart builds and releases for your Go command-line tool.

### Build Workflow

[`.github/workflows/build.yml`](.github/workflows/build.yml)

Builds the command-line tool on each push to the `main` branch.

### Release Workflow

[`.github/workflows/release.yml`](.github/workflows/release.yml)

#### New Releases Through Pushing Git Tags

When a [git tag](https://git-scm.com/book/en/v2/Git-Basics-Tagging) with the format `v*` is pushed, the [release workflow](.github/workflows/release.yml) builds the command-line tool and creates a [GitHub release](https://docs.github.com/en/repositories/releasing-projects-on-github/about-releases).

E.g. A GitHub release is created when the tag `v1.2.5` is pushed (conforms to the git tag format `v*`).

#### New Pre-Releases Through Pushing Git Tags

When a [git tag](https://git-scm.com/book/en/v2/Git-Basics-Tagging) containing `-alpha.`, `-beta.`, `-rc.`, or `-nightly.` is pushed, a [GitHub _pre-release_](https://docs.github.com/en/repositories/releasing-projects-on-github/about-releases) is created instead.
GitHub will point out that this release is identified as non-production ready.
