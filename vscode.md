# Visual Studio Code

## Links

- [Tips and tricks](https://github.com/Microsoft/vscode-tips-and-tricks)
- [A curated list of delightful VS Code packages and resources](https://github.com/viatsko/awesome-vscode)
- [PHP and Vue optimized settings and extensions](http://calebporzio.com/my-vs-code-setup-2/)
- [VSCodium – open source Visual Studio Code without trackers](https://www.fossmint.com/vscodium-clone-of-visual-studio-code-for-linux/)

## User settings (settings.json)

```json
{
    "[html]": {
        "editor.defaultFormatter": "esbenp.prettier-vscode"
    },
    "[javascript]": {
        "editor.defaultFormatter": "esbenp.prettier-vscode"
    },
    "[json]": {
        "editor.defaultFormatter": "esbenp.prettier-vscode"
    },
    "[svelte]": {
        "editor.defaultFormatter": "JamesBirtles.svelte-vscode"
    },
    "[yaml]": {
        "editor.defaultFormatter": "redhat.vscode-yaml"
    },
    "breadcrumbs.enabled": true,
    "diffEditor.ignoreTrimWhitespace": true,
    "editor.codeActionsOnSave": {
        "source.fixAll.eslint": true
    },
    "editor.cursorBlinking": "phase",
    "editor.find.addExtraSpaceOnTop": false,
    "editor.fontFamily": "'Operator Mono Book','Fira Code'",
    "editor.fontLigatures": true,
    "editor.fontSize": 16,
    "editor.formatOnPaste": true,
    "editor.formatOnSave": true,
    "editor.minimap.enabled": false,
    "editor.rulers": [
        120
    ],
    "eslint.enable": true,
    "eslint.validate": [
        "javascript",
        "svelte"
    ],
    "files.autoSave": "onWindowChange",
    "files.insertFinalNewline": true,
    "files.trimTrailingWhitespace": true,
    "git.autofetch": true,
    "git.enableSmartCommit": true,
    "go.autocompleteUnimportedPackages": true,
    "go.buildOnSave": "off",
    "go.formatFlags": [
        "-w"
    ],
    "go.lintTool": "revive",
    "go.useLanguageServer": true,
    "material-icon-theme.showWelcomeMessage": false,
    "terminal.integrated.fontSize": 16,
    "terraform.indexing": {
        "delay": 500,
        "enabled": false,
        "exclude": [
            ".terraform/**/*",
            "**/.terraform/**/*"
        ],
        "liveIndexing": false
    },
    "terraform.languageServer": {
        "args": [],
        "enabled": true
    },
    "window.zoomLevel": 0,
    "workbench.colorTheme": "Night Owl",
    "workbench.iconTheme": "material-icon-theme",
    "workbench.startupEditor": "newUntitledFile",
    "yaml.format.bracketSpacing": false
}
```

## Recommended Extensions

### Go

```json
"ms-vscode.go",
```

### HTML & CSS

```json
"formulahendry.auto-rename-tag",
"esbenp.prettier-vscode",
"ritwickdey.liveserver",
"mrmlnc.vscode-autoprefixer",
```

### JavaScript

```json
"dbaeumer.vscode-eslint",
"jamesbirtles.svelte-vscode",
"eg2.vscode-npm-script",
```

### General

```json
"streetsidesoftware.code-spell-checker",
"davidanson.vscode-markdownlint",
"christian-kohler.path-intellisense",
"richie5um2.vscode-sort-json",
"tyriar.sort-lines",
"wayou.vscode-todo-highlight",
"ms-vsliveshare.vsliveshare",
"quicktype.quicktype",
"humao.rest-client",
```

### Docker & Kubernetes

```json
"ms-azuretools.vscode-docker",
"ms-kubernetes-tools.vscode-kubernetes-tools",
"redhat.vscode-yaml",
"pascalreitermann93.vscode-yaml-sort",
```

### Git

```json
"donjayamanne.githistory",
"codezombiech.gitignore",
"eamodio.gitlens",
```

### Jenkins

```json
"secanis.jenkinsfile-support",
"janjoerke.jenkins-pipeline-linter-connector",
```

### Terraform

```json
"mauve.terraform",
```

### Azure

```json
"ms-vscode.azure-account",
"ms-vscode.azurecli",
"ms-azuretools.vscode-azureterraform",
```

## Ansible

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
