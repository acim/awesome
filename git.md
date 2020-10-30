# git

- [Flight rules for Git](https://github.com/k88hudson/git-flight-rules)
- [Useful Git Commands](https://dev.to/lydiahallie/cs-visualized-useful-git-commands-37p1)

## Delete merged local branches

```sh
git remote prune origin # prunes tracking branches not on the remote
git branch —merged # lists branches that have been merged into the current branch
git branch -d ...
```

## Shallow clone repository (single branch with truncated history)

```sh
git clone --depth 1 https://path/to/repo/foo.git -b bar
```

## Enable Git autocorrection

```sh
git config --global help.autocorrect 1
```

## Count commits

```sh
git rev-list --count
git rev-list --count dev
```

## Optimize repo

```sh
git gc --prune=now --aggressive
```

## View a file of another branch

```sh
git show main:README.md
```

## Search

```sh
git rev-list –all | xargs git grep -F 'something'
```
