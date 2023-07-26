# Visual Studio Code

## Links

- [Tips and tricks](https://code.visualstudio.com/docs/getstarted/tips-and-tricks)
- [A curated list of delightful VS Code packages and resources](https://github.com/viatsko/awesome-vscode)
- [VSCodium – open source Visual Studio Code without trackers](https://www.fossmint.com/vscodium-clone-of-visual-studio-code-for-linux/)

## User settings (settings.json)

```json
{
  "[go]": {
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  },
  "[go.mod]": {
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  },
  "[html]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[javascript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode",
    "editor.codeActionsOnSave": {
      "source.fixAll.eslint": true
    }
  },
  "[json]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[jsonc]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[markdown]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[rust]": {
    "editor.rulers": [100]
  },
  "[svelte]": {
    "editor.defaultFormatter": "svelte.svelte-vscode"
  },
  "[typescript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode",
    "editor.codeActionsOnSave": {
      "source.fixAll.eslint": true
    }
  },
  "[yaml]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "breadcrumbs.enabled": true,
  "crates.listPreReleases": true,
  "diffEditor.ignoreTrimWhitespace": true,
  "editor.cursorBlinking": "phase",
  "editor.find.addExtraSpaceOnTop": false,
  "editor.fontFamily": "'JetBrainsMono Nerd Font'",
  "editor.fontLigatures": true,
  "editor.fontSize": 14,
  "editor.formatOnPaste": true,
  "editor.formatOnSave": true,
  "editor.letterSpacing": 0.4,
  "editor.minimap.enabled": false,
  "editor.rulers": [120],
  "editor.smoothScrolling": true,
  "eslint.enable": true,
  "eslint.validate": ["javascript", "svelte"],
  "files.autoSave": "onFocusChange",
  "files.insertFinalNewline": true,
  "files.trimTrailingWhitespace": true,
  "git.autofetch": true,
  "git.enableSmartCommit": true,
  "gitlens.defaultDateFormat": "DD.MM.YYYY HH:mm:ss",
  "go.coverOnSingleTest": true,
  "go.coverOnSingleTestFile": true,
  "go.lintFlags": ["--fast"],
  "go.lintTool": "golangci-lint",
  "go.toolsManagement.autoUpdate": true,
  "prettier.semi": false,
  "prettier.singleQuote": true,
  "rust-analyzer.checkOnSave.command": "clippy",
  "rust-analyzer.inlayHints.enable": false,
  "svelte.language-server.runtime": "/usr/bin/node",
  "terminal.integrated.fontSize": 14,
  "terraform.languageServer": {
    "external": true,
    "args": ["serve"]
  },
  "typescript.updateImportsOnFileMove.enabled": "always",
  "workbench.colorTheme": "Night Owl",
  "workbench.iconTheme": "material-icon-theme",
  "workbench.startupEditor": "newUntitledFile"
}
```

## Recommended Extensions

### General

```json
"esbenp.prettier-vscode",
"humao.rest-client",
"rangav.vscode-thunder-client",
"streetsidesoftware.code-spell-checker",
"wayou.vscode-todo-highlight",
"yzhang.markdown-all-in-one",
"inu1255.easy-snippet",
```

### Themes

```json
{
  "pkief.material-icon-theme",
  "sdras.night-owl",
}
```

### Rust

```json
{
  "tamasfe.even-better-toml",
  "rust-lang.rust-analyzer",
  "serayuzgur.crates",
  "vadimcn.vscode-lldb".
}
```

### Go

```json
{
  "golang.go",
  "zxh404.vscode-proto3",
}
```

### HTML & CSS

```json
{
  "ritwickdey.liveserver",
}
```

### JavaScript

```json
{
  "dbaeumer.vscode-eslint",
  "svelte.svelte-vscode",
}
```

### Docker & Kubernetes

```json
{
  "ms-azuretools.vscode-docker",
  "ms-kubernetes-tools.vscode-kubernetes-tools",
  "redhat.vscode-yaml",
}
```

### Git

```json
{
  "codezombiech.gitignore",
  "donjayamanne.githistory",
  "eamodio.gitlens",
  "mhutchie.git-graph",
}
```

### Jenkins

```json
{
  "janjoerke.jenkins-pipeline-linter-connector",
  "secanis.jenkinsfile-support",
}
```

### Terraform

```json
{
  "hashicorp.terraform",
}
```

### Azure

```json
{
  "ms-azuretools.vscode-azureterraform",
  "ms-vscode.azure-account",
  "ms-vscode.azurecli",
}
```

## Custom settings

### Ansible

```json
{
  "files.associations": {
    "*.yml": "ansible"
  }
}
```

### Vue.js and NUXT

```json
{
  "eslint.validate": [
    {
      "language": "vue",
      "autoFix": true
    },
    {
      "language": "javascript",
      "autoFix": true
    },
    {
      "language": "javascriptreact",
      "autoFix": true
    }
  ],
  "eslint.autoFixOnSave": true,
  "editor.formatOnSave": false,
  "vetur.validation.template": false
}
```

### Delve

```json
{
  "go.delveConfig": {
    "dlvLoadConfig": {
      "followPointers": true,
      "maxVariableRecurse": 3,
      "maxStringLen": 120,
      "maxArrayValues": 120,
      "maxStructFields": -1
    },
    "apiVersion": 2,
    "showGlobalVariables": true
  }
}
```

### Multiple vertical rulers

```json
{
  "editor.rulers": [
    {
      "column": 80, // spacing of 1st column from left
      "color": "#ff9900" // orange, Go Vols!
    },
    100, // 2nd ruler with no color option
    {
      "column": 120, // third ruler
      "color": "#9f0af5" // purple, go Pirates!
    }
  ]
}
```

```json
{
  "[rust]": {
    "editor.rulers": [
      {
        "column": 100,
        "color": "#00ff22"
      }
    ]
  }
}
```

## Fonts

- [Nerd fonts](https://www.nerdfonts.com/font-downloads)
- [JuliaMono](https://cormullion.github.io/pages/2020-07-26-JuliaMono/)

## Install from apt

```sh
wget -qO - https://gitlab.com/paulcarroty/vscodium-deb-rpm-repo/raw/master/pub.gpg | gpg --dearmor | sudo dd of=/usr/share/keyrings/vscodium-archive-keyring.gpg

echo 'deb [ signed-by=/usr/share/keyrings/vscodium-archive-keyring.gpg ] https://download.vscodium.com/debs vscodium main' | sudo tee /etc/apt/sources.list.d/vscodium.list

sudo apt-get update
sudo apt-get install codium
```
