# awesome

[![License](https://img.shields.io/badge/license-BSD--2--Clause--Patent-orange.svg)](https://github.com/acim/awesome/blob/main/LICENSE)

A curated list of delightful developers resources.

- [CSS](css.md)
- [Databases - MySQL, PostgreSQL, MongoDB](db.md)
- [Docker](docker.md)
- [Drone, Tekton Pipelines, etc.](cicd.md)
- [Frontend - JavaScript, HTML5, CSS3, etc.](frontend.md)
- [Git](git.md)
- [Go](go.md)
- [JavaScript & TypeScript](js-ts.md)
- [Kubernetes](k8s.md)
- [Miscellaneous](misc.md)
- [Node.js](node.md)
- [Raspberry Pi](pi.md)
- [Rust](rust.md)
- [Svelte](svelte.md)
- [Visual Studio Code](vscode.md)
- [Vue & NUXT](vue.md)

## Check broken links

`cat *.md | grep -Eo "(http|https)://[a-zA-Z0-9./?=_%:-]*" | parallel --gnu "wget --spider --no-verbose --timeout 10 {}"`
