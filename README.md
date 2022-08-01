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
Replace `<your-github-username>` and `<your-github-repo-name>` before running the following command:

```bash
rm go.mod go.sum
go mod init github.com/<your-github-username>/<your-github-repo-name>
go mod tidy
```
