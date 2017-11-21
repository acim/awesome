# vscode

## Links

* [Tips and tricks](https://github.com/Microsoft/vscode-tips-and-tricks)
* [A curated list of delightful VS Code packages and resources](https://github.com/viatsko/awesome-vscode)

## User settings (settings.json)

```json
{
    "editor.fontFamily": "'Fira Code'",
    "editor.fontLigatures": true,
    "editor.fontSize": 14,
    "editor.formatOnPaste": true,
    "editor.formatOnSave": true,
    "files.autoSave": "afterDelay",
    "files.exclude": {
        "**/.DS_Store": true,
        "**/.git": true,
        "**/.hg": true,
        "**/.idea": true,
        "**/.svn": true,
        "**/.terraform": true,
        "**/CVS": true
    },
    "git.enableSmartCommit": true,
    "go.formatFlags": [
        "-w"
    ],
    "search.exclude": {
        "**/bower_components": true,
        "**/node_modules": true,
        "**/vendor": true
    },
    "terraform.formatOnSave": true,
    "workbench.startupEditor": "newUntitledFile"
}
```

## Workspace setting for repo containing Ansible

```json
{
    "files.associations": {
        "*.yml": "ansible"
    },
}
```