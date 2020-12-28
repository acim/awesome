# Visual Studio Code

## Links

- [Tips and tricks](https://github.com/Microsoft/vscode-tips-and-tricks)
- [A curated list of delightful VS Code packages and resources](https://github.com/viatsko/awesome-vscode)
- [PHP and Vue optimized settings and extensions](http://calebporzio.com/my-vs-code-setup-2/)
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
  "diffEditor.ignoreTrimWhitespace": true,
  "editor.cursorBlinking": "phase",
  "editor.find.addExtraSpaceOnTop": false,
  "editor.fontFamily": "'Operator Mono Book','Fira Code'",
  "editor.fontLigatures": true,
  "editor.fontSize": 16,
  "editor.formatOnPaste": true,
  "editor.formatOnSave": true,
  "editor.minimap.enabled": false,
  "editor.rulers": [120],
  "eslint.enable": true,
  "eslint.validate": ["javascript", "svelte"],
  "files.autoSave": "onFocusChange",
  "files.insertFinalNewline": true,
  "files.trimTrailingWhitespace": true,
  "git.autofetch": true,
  "git.enableSmartCommit": true,
  "go.autocompleteUnimportedPackages": true,
  "go.buildOnSave": "off",
  "go.coverOnSingleTest": true,
  "go.coverOnSingleTestFile": true,
  "go.formatTool": "gofumports",
  "go.lintFlags": ["--fast"],
  "go.lintTool": "golangci-lint",
  "go.useLanguageServer": true,
  "gopls": {
    "gofumpt": true
  },
  "prettier.semi": false,
  "prettier.singleQuote": true,
  "svelte.language-server.runtime": "/usr/bin/node",
  "terminal.integrated.fontSize": 16,
  "terraform.languageServer": {
    "external": true,
    "args": ["serve"]
  },
  "typescript.updateImportsOnFileMove.enabled": "always",
  "window.zoomLevel": 0,
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
"streetsidesoftware.code-spell-checker",
"wayou.vscode-todo-highlight",
"yzhang.markdown-all-in-one",
```

### Themes

```json
"pkief.material-icon-theme",
"sdras.night-owl",
```

### Rust

```json
"bungcip.better-toml",
"matklad.rust-analyzer",
"rust-lang.rust",
"serayuzgur.crates",
"vadimcn.vscode-lldb",
```

### Go

```json
"golang.go",
"zxh404.vscode-proto3",
```

### HTML & CSS

```json
"ritwickdey.liveserver",
```

### JavaScript

```json
"dbaeumer.vscode-eslint",
"svelte.svelte-vscode",
```

### Docker & Kubernetes

```json
"ms-azuretools.vscode-docker",
"ms-kubernetes-tools.vscode-kubernetes-tools",
"redhat.vscode-yaml",
```

### Git

```json
"codezombiech.gitignore",
"eamodio.gitlens",
```

### Jenkins

```json
"janjoerke.jenkins-pipeline-linter-connector",
"secanis.jenkinsfile-support",
```

### Terraform

```json
"hashicorp.terraform",
```

### Azure

```json
"ms-azuretools.vscode-azureterraform",
"ms-vscode.azure-account",
"ms-vscode.azurecli",
```

## Custom settings

### Rust

```json
"crates.latestDecorator": "🚨 Latest: ${version}",
"crates.upToDateDecorator": "",
```

### Ansible

```json
{
  "files.associations": {
    "*.yml": "ansible"
  }
}
```

## Vue.js and NUXT

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

## Fonts

- [Fira Code](https://github.com/tonsky/FiraCode)
- [JuliaMono](https://cormullion.github.io/pages/2020-07-26-JuliaMono/)
- [Operator Mono](https://www.typography.com/fonts/operator/styles)
- [Dank Mono](https://gumroad.com/l/dank-mono)

## [Fix missing extenstions](https://github.com/VSCodium/vscodium/issues/418#issuecomment-643664182)

## [429 Too Many Requests](https://gitlab.com/paulcarroty/vscodium-deb-rpm-repo/-/issues/36#note_395793123)
