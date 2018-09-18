# Visual Studio Code

## Links

* [Tips and tricks](https://github.com/Microsoft/vscode-tips-and-tricks)
* [A curated list of delightful VS Code packages and resources](https://github.com/viatsko/awesome-vscode)
* [PHP and Vue optimized settings and extensions](http://calebporzio.com/my-vs-code-setup-2/)

## User settings (settings.json)

```json
{
    "editor.fontFamily": "'Fira Code'",
    "editor.fontLigatures": true,
    "editor.fontSize": 14,
    "editor.minimap.enabled": false,
    "editor.rulers": [
        120
    ],
    "editor.cursorStyle": "line",
    "editor.cursorBlinking": "phase",
    "editor.renderWhitespace": "none",
    "editor.multiCursorModifier": "alt",
    "files.autoSave": "onWindowChange",
    "files.trimTrailingWhitespace": true,
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
    "gitlens.advanced.messages": {
        "suppressCommitHasNoPreviousCommitWarning": false,
        "suppressCommitNotFoundWarning": false,
        "suppressFileNotUnderSourceControlWarning": false,
        "suppressGitVersionWarning": false,
        "suppressLineUncommittedWarning": false,
        "suppressNoRepositoryWarning": false,
        "suppressResultsExplorerNotice": false,
        "suppressShowKeyBindingsNotice": true,
        "suppressUpdateNotice": false,
        "suppressWelcomeNotice": true
    },
    "git.autofetch": true,
    "gitlens.keymap": "chorded",
    "search.location": "sidebar",
    "gitlens.historyExplorer.enabled": true,
    "diffEditor.ignoreTrimWhitespace": false,
    "workbench.colorTheme": "Cobalt2",
    "terminal.integrated.fontSize": 14,
    "editor.formatOnPaste": true,
    "editor.formatOnSave": true,
    "go.formatFlags": [
        "-w"
    ],
    "go.autocompleteUnimportedPackages": true,
    "go.buildOnSave": "off",
    "go.lintTool": "megacheck",
    "editor.insertSpaces": false,
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