# git

- [Flight rules for Git](https://github.com/k88hudson/git-flight-rules)
- [Useful Git Commands](https://dev.to/lydiahallie/cs-visualized-useful-git-commands-37p1)

## Delete merged local branches

```bash
git remote prune origin # prunes tracking branches not on the remote
git branch —merged # lists branches that have been merged into the current branch
git branch -d ...
```

## Shallow clone repository (single branch with truncated history)

```bash
git clone --depth 1 https://path/to/repo/foo.git -b bar
```
