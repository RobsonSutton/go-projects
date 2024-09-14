# Implementation Notes

This project will be attempted with the high level requirement of building a git equivalent via Go (obviously a little simpler though!)

This will be broken down based on some high level objectives identified from this course's first page - https://app.codecrafters.io/courses/git/overview
1. Initialise the .git directory
2. Read a blob object
3. Create a blob object
4. Read a tree object
5. Write a tree object
6. Create a commit
7. Clone a repository

## Initialise the .git directory

The focus of this is to replicate the `git init .` command.

**What does this command do?**

This will initialise the following files within the current repo:
```
.git
├── HEAD
├── config
├── description
├── hooks
│   ├── applypatch-msg.sample
│   ├── commit-msg.sample
│   ├── fsmonitor-watchman.sample
│   ├── post-update.sample
│   ├── pre-applypatch.sample
│   ├── pre-commit.sample
│   ├── pre-merge-commit.sample
│   ├── pre-push.sample
│   ├── pre-rebase.sample
│   ├── pre-receive.sample
│   ├── prepare-commit-msg.sample
│   ├── push-to-checkout.sample
│   └── update.sample
├── info
│   └── exclude
├── objects
│   ├── info
│   └── pack
└── refs
    ├── heads
    └── tags
```