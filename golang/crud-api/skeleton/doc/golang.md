# Golang Setup with VSCode

This guide will help you set up Go (Golang) development environment using Visual Studio Code (VSCode). Follow the steps below to get everything up and running.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.22 or higher)
- [Visual Studio Code](https://code.visualstudio.com/Download)

## Step 1: Install Go

1. Download the Go installer from the [official Go website](https://golang.org/dl/).
2. Run the installer and follow the on-screen instructions.
3. Verify the installation by opening a terminal and running:

    ```sh
    go version
    ```

    You should see the Go version printed in the terminal.

## Step 2: Set Up Go Workspace

1. Create a directory for your Go workspace, for example:

    ```sh
    mkdir -p $HOME/go/{bin,src,pkg}
    ```

2. Set the `GOPATH` environment variable to point to your workspace:

    - On macOS/Linux, add the following lines to your `~/.bashrc`, `~/.zshrc`, or corresponding shell configuration file:

        ```sh
        export GOPATH=$HOME/go
        export PATH=$PATH:$GOPATH/bin
        ```

    - On Windows, add `GOPATH` and update the `PATH` environment variable through the System Properties.

3. Apply the changes by sourcing the configuration file or restarting your terminal.

## Step 3: Install VSCode Extensions

1. Open VSCode.
2. Go to the Extensions view by clicking on the Extensions icon in the Activity Bar on the side of the window or pressing `Ctrl+Shift+X`.
3. Search for the `Go` extension by the Go team at Google and install it.

## Step 4: Configure VSCode for Go

1. Open the Command Palette by pressing `Ctrl+Shift+P`.
2. Type `Go: Install/Update Tools` and select it.
3. In the list of Go tools, select all and click OK to install them. This will set up various tools like `golangci-lint` `gopls`, `dlv`, and `go-outline`.

## Additional Configuration

For additional settings and customization, you can modify the `settings.json` file in your `.vscode` directory. Here are some useful settings for Go development:

```json
{
    "go.lintOnSave": "file",
    "go.lintTool": "golangci-lint",
    "go.lintFlags": [
        "--fast"
    ],
    "go.formatTool": "goimports",
    "go.useLanguageServer": true,
    "[go]": {
        "editor.insertSpaces": true,
        "editor.defaultFormatter": "golang.go",
        "editor.formatOnSave": true,
        "editor.codeActionsOnSave": {
            "source.organizeImports": "explicit"
        }
    },
}
