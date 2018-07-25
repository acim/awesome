# Visual Studio Code

## Links

* [Tips and tricks](https://github.com/Microsoft/vscode-tips-and-tricks)
* [A curated list of delightful VS Code packages and resources](https://github.com/viatsko/awesome-vscode)
* [PHP and Vue optimized settings and extensions](http://calebporzio.com/my-vs-code-setup-2/)

## User settings (settings.json)

```json
{
{
    "editor.fontFamily": "'Fira Code'",
    "editor.fontLigatures": true,
    "editor.fontSize": 14,
    "editor.minimap.enabled": false,
    "editor.rulers": [120],
    "editor.cursorStyle": "line",
    "editor.cursorBlinking": "phase",
    "editor.renderWhitespace": "none",
    "editor.multiCursorModifier": "alt",
    "files.autoSave": "onWindowChange",
    "files.exclude": {
        "**/.DS_Store": true,
        "**/.git": true,
        "**/.idea": true,
        "**/.terraform": true,
    },
    "git.enableSmartCommit": true,
    "search.exclude": {
        "**/bower_components": true,
        "**/node_modules": true
    },
    "workbench.startupEditor": "newUntitledFile",
    "explorer.confirmDragAndDrop": false,
    "git.autofetch": true,
    "search.location": "sidebar",
    "diffEditor.ignoreTrimWhitespace": true,
    "workbench.colorTheme": "Cobalt2",
    "files.trimTrailingWhitespace": true,
}
```

## Workspace settings for a repo containing Ansible

```json
{
    "files.associations": {
        "*.yml": "ansible"
    },
}
```