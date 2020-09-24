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
  "[jsonc]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[svelte]": {
    "editor.defaultFormatter": "svelte.svelte-vscode"
  },
  "[typescript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[yaml]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
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
  "go.lintFlags": [
    "-exclude=vendor/...",
    "-config=${workspaceFolder}/config.toml"
  ],
  "go.lintTool": "revive",
  "go.useLanguageServer": true,
  "prettier.semi": false,
  "svelte.language-server.runtime": "/usr/bin/node",
  "terminal.integrated.fontSize": 16,
  "terraform.languageServer": {
    "enabled": true,
    "args": []
  },
  "window.zoomLevel": 0,
  "workbench.colorTheme": "Night Owl",
  "workbench.iconTheme": "material-icon-theme",
  "workbench.startupEditor": "newUntitledFile"
}
```

## Recommended Extensions

### General

```json
"streetsidesoftware.code-spell-checker",
"davidanson.vscode-markdownlint",
"richie5um2.vscode-sort-json",
"tyriar.sort-lines",
"wayou.vscode-todo-highlight",
"humao.rest-client",
```

### Themes

```json
"sdras.night-owl",
"pkief.material-icon-theme",
```

### Go

```json
"golang.go",
"zxh404.vscode-proto3",
```

### HTML & CSS

```json
"esbenp.prettier-vscode",
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
"hashicorp.terraform",
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

## [Fix missing extenstions](https://github.com/VSCodium/vscodium/issues/418#issuecomment-643664182)

## [429 Too Many Requests](https://gitlab.com/paulcarroty/vscodium-deb-rpm-repo/-/issues/36#note_395793123)
