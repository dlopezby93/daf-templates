## Prerequisites

On **Mac** The project requires:
- [Homebrew](https://brew.sh/)
- [pre-commit](https://pre-commit.com/#install)

On **Windows** The project requires:
- [python](https://www.python.org/downloads/)
- [pre-commit](https://pre-commit.com/#install)

## Quick Start

In this project, we have embraced the [Conventional Commits convention](https://www.conventionalcommits.org/en/v1.0.0/) to standardize and enhance the clarity of our commit history. This convention defines a set of rules for structuring commit messages consistently, making it easier to understand and track changes in the codebase

### Clone the repository locally

To clone this repository you need to clone this repository as upstream.

```bash
$ git clone https://github.com/name_organization/name_repository
$ cd name_repository
```

### Install and configure pre-commit

After cloning the repository you must install [pre-commit](https://pre-commit.com/#install). Once you have pre-commit installed, execute this commands:

```bash
$ pre-commit install
$ pre-commit install --hook-type commit-msg
```

### Running pre-commit locally

To validate that all configured quality gates are running successfully locally:

```bash
$ pre-commit run --all-files
```
