# Visual Studio Code

## Links

- [Tips and tricks](https://code.visualstudio.com/docs/getstarted/tips-and-tricks)
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
  "go.autocompleteUnimportedPackages": true,
  "go.buildOnSave": "off",
  "go.coverOnSingleTest": true,
  "go.coverOnSingleTestFile": true,
  "go.lintFlags": ["--fast"],
  "go.lintTool": "golangci-lint",
  "prettier.semi": false,
  "prettier.singleQuote": true,
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
"streetsidesoftware.code-spell-checker",
"wayou.vscode-todo-highlight",
"yzhang.markdown-all-in-one",
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
  "bungcip.better-toml",
  "matklad.rust-analyzer",
  "rust-lang.rust",
  "serayuzgur.crates",
  "vadimcn.vscode-lldb",
  "rangav.vscode-thunder-client"
}
```

### Go

```json
{
  "golang.go",
  "zxh404.vscode-proto3",
  "rangav.vscode-thunder-client",
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
