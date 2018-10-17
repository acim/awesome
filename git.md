# git

* [Flight rules for Git](https://github.com/k88hudson/git-flight-rules)

## Delete merged local branches

```bash
git remote prune origin # prunes tracking branches not on the remote
git branch —merged # lists branches that have been merged into the current branch
git branch -d ...
```