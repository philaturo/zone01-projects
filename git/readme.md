# Git Project Documentation

This document summarizes the Git commands used in the git project and explains their purpose. The commands are used to inspect, filter and format the commit history as required in the project.

## Commands Used

### Installing, Setting Up And Configuring Git

```bash
git --version
```

```bash
git init -b main
```

```bash
git config --global user.name "phaturo"
git config --global user.email "aturophil09@gmail.com"
```

Installed Git on local machine, checked version, initialized in current directory and configured Git with the username and email address.

### Git Commits To Commit

```bash
mkdir hello
nano hello.sh
echo "Hello, World"
```

```bash
git init
```

```bash
git status
```

```bash
git add hello.sh
git commit -m "Update hello.sh to accept an argument"
```

### View Full Commit History

```bash
git log
```

Displays the complete commit history of the repository, including commit hashes, authors,dates and messages

### One-Line Commit History

```bash
git log --oneline
```

Shows a compact version of the commit history with the abbreviated commit hashes and commit messages.

### Display Last Two Commits

```bash
git log -n 2
```

Limits the output to the two most recent commits

### Commits Within The Last Five Minutes

```bash
git log --since="5 minutes ago"
```

Filters the commit history to show only commits made with the specified time range

### Custom Formatted Log Output

```bash
git log --pretty=format:"* %h %ad | %s (%d) [%an]" --date=short
```

Displays the commit history in a personalized format including:

* Short commit hash
* Commit date
* Commit message
* Branch and tag references
* Author name

## Restore Snapshots (Checkout/Restore)

### Restore First Snapshot

```bash
git checkout <first-commit-hash>
cat hello.sh
```

Moves HEAD to the first commit and displays the file content as it was initially.

### Restore Second Most Recent Snapshot

```bash
git checkout HEAD~1
cat hello.sh
```

Checks out the most recent commit

### Return To Latest Version

```bash
git checkout master
```

---

## Tags

### Tag Current Version

```bash
git tag v1
```

### Tag Previous Version

```bash
git tag v1-beta HEAD~1
```

### Navigate Between Tags

```bash
git checkout v1
git checkout v1-beta
```

### List All Tags

```bash
git tag
```

---

## Undo Changes

### Revert Unstaged Changes

```bash
git restore hello.sh
```

Discards changes in the working directory before staging.

### Revert staged Changes

```bash
git restore --staged hello.sh
git restore hello.sh
```

Removes the file from the staging area and discards changes

### Revert a Commited Change

```bash
git revert HEAD
```

Creates a new commit that undoes the previous commit safely

---

## Tag and Reset (Removing Commits)

### Tag the Latest Commit

```bash
git tag oops
```

### Reset to Version v1

```bash
git reset --hard v1
```

Removes commits made after `v1` and moves HEAD back.

## Displaying Logs with Deleted Commits

```bash
git log --all --decorate --oneline
```

Shows commit history including commits that are no longer reachable from HEAD, such as the one tagged `oops`.

## Cleaning Unreferenced Commits

```bash
git reflog expire --expire=now --all
git gc --prune=now
```

Removes unreferenced commits permanently from the repository.

## Author Information

### Add Author Comment and Commit


```bash
# Default is World
# Author: Jim Weirich
name=${1:-"World"}
echo "Hello, $name"
```

```bash
git add hello.sh
git commit -m "Add author name"
```

### Fix Missing Author Email (Amend Commit)

Manualy added the email to **hello.sh** and included it in the **same commit**.

#### Updated hello.sh

```bash
# Author: Jim Weirich <jim.weirich@gmail.com>
```

```bash
git commit --amend
```

Updates the last commit without creating a new one.

---

## Move It

### Moving `hello.sh` into `lib/`

```bash
mkdir lib
git mv hello.sh lib/hello.sh
git commit -m "Move hello.sh into lib directory"
```

### Create Makefile

```makefile
TARGET="lib/hello.sh"

run:
	bash ${TARGET}
```

```bash
git add Makefile
git commit -m "Add Makefile to run hello script"
```

---

## Blobs, Trees, and Commits

### Exploring `.git/` Directory

* **objects/**: Stores all Git objects (blobs, trees, commits)
* **refs/**: References to branches and tags
* **HEAD**: Points to the current branch or commit
* **config**: Repository-specific configuration

```bash
ls .git
```

Examines the contents of the .git directory.

### Latest Object Hash

```bash
git rev-parse HEAD
git cat-file -t HEAD
git cat-file -p HEAD
```

Finds the latest object hash within the .git/objects/ directory and prints the type and content of the objects using the Git commands.

### Dump Directory Tree of a Commit

```bash
git cat-file -p HEAD^{tree}
```

Dumps the directory tree referenced by this commit.

### Dump `lib/` and `hello.sh`

```bash
git ls-tree HEAD:lib
git cat-file -p HEAD:lib/hello.sh
```

---

## Branching

### Create and Switch to `greet` Branch

```bash
git checkout -b greet
```

### Add `greeter.sh`

```bash
#!/bin/bash

Greeter() {
    who="$1"
    echo "Hello, $who"
}
```

```bash
git add lib/greeter.sh
git commit -m "Add greeter.sh"
```

### Update `lib/hello.sh`

```bash
#!/bin/bash

source lib/greeter.sh

name="$1"
if [ -z "$name" ]; then
    name="World"
fi

Greeter "$name"
```

```bash
git add lib/hello.sh
git commit -m "Update hello.sh to use greeter function"
```

### Update Makefile

```makefile
# Ensure it runs the updated lib/hello.sh file
TARGET="lib/hello.sh"

run:
	bash ${TARGET}
```

```bash
git add Makefile
git commit -m "Update Makefile comment"
```

### Compare Branches

```bash
git checkout master
git diff master greet -- Makefile lib/hello.sh lib/greeter.sh
```

### Commit Tree Diagram

```
git log --oneline --graph --all

```

---

## Conflicts, Merging, and Rebasing

### Merge Main into Greet

```bash
git checkout greet
git merge master
```

### Create Conflict on `master`

```bash
git checkout master
```

```bash
#!/bin/bash

echo "What's your name"
read my_name

echo "Hello, $my_name"
```

```bash
git add lib/hello.sh
git commit -m "Change hello.sh to interactive input"
```

### Merge Conflict

```bash
git checkout greet
git merge master
```

Resolve conflict manually, accept `master` changes:

```bash
git add lib/hello.sh
git commit -m "Resolve merge conflict"
```

### Rebase `greet` onto `master`

```bash
git rebase master
```

### Merge `greet` into `master`

```bash
git checkout master
git merge greet
```

### Fast-Forward vs Merge vs Rebase

* **Fast-forward**: No divergence, branch pointer moves forward
* **Merge**: Creates a merge commit, preserves history
* **Rebase**: Rewrites history, linearizes commits

---

## Local and Remote Repositories

### Clone Repository

```bash
git clone hello cloned_hello
```

### Show Logs

```bash
git log --oneline --graph --all
```

### Show Remotes

```bash
git remote -v
```

### List Branches

```bash
git branch
git branch -r
git branch -a
```

### Update Original Repository

```bash
This is the Hello World example from the git project.
(changed in the original)
```

```bash
git add README.md
git commit -m "Update README in original repository"
```

### Fetch and Merge in Clone

```bash
git fetch origin
git merge origin/master
```

### Track Remote Branch

```bash
git checkout -b greet origin/greet
```

### Single Command Equivalent

```bash
git pull
```

---

## Bare Repositories

### What Is a Bare Repository?

A bare repository has no working directory and is used as a shared central repository to safely push and pull changes.

### Create Bare Repository

```bash
git clone --bare hello hello.git
```

### Add Bare Repo as Remote

```bash
git remote add shared ../hello.git
```

### Push to Shared

```bash
git push shared master
```

### Pull from Shared in Clone

```bash
git pull
```

---

## Summary

This project demonstrates:

* Git history and snapshots
* Tags and resets
* Branching and conflicts
* Rebasing and merging
* Local, remote, and bare repositories
* Real-world Git workflows suitable for team collaboration

All requirements have been completed and validated via commit history.


## Lessons Learnt

* Git is snapshot-based, not file-based.
* The working tree, staging area, and history are separate.
* Detached HEAD is a normal Git state.
* Branches are pointers, tags are fixed references.
* Restoring history does not delete commits immediately.
* Logs can be customized for audit.
