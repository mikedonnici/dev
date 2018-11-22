# Git

.gitconfig
```
[user]
	email = michael@mesa.net.au
	name = Mike Donnici
[core]
	editor = vim
[alias]
	last = log -1 HEAD
	ll = log --abbrev-commit --oneline --decorate --graph --all
	unstage = reset HEAD --
```



--- 
### Git Push to All Origins

This is useful for deploying updates to multiple Heroku deployments at once - probably best that there is only ONE active branch locally before doing it.

```bash
$ git remote | xargs -L1 git push --all
```

**NOTE** - This will push all the branches as well which creates crud branches in places you may not need them, eg Heroku. 

If this is the case remote branches can be pruned off thus:

```bash
$ git push origin --delete branch
```

...where `origin` and `branch` are the correct *origin* and *branch* names.




---
### Git merge using 'ours' strategy

This is handy when a lot of changes have been made to a branch and merging back into **master** is a conflict nightmare, for example when the app was _herokufied_. 

Effectively, this replaces the **master** branch with the **development** branch - [reference](http://stackoverflow.com/questions/2862590/how-to-replace-master-branch-in-git-entirely-from-another-branch)

These steps show this process for fresh clone of _mappcpd-member_, and the merging of the **heroku** branch.

```
git clone git@github.com:mappsystems/mappcpd-member.git
cd mappcpd-member
git branch
# [should be on 'master']
# checkout the REMOTE development branch - don't create a new, local one!...
git checkout -b origin/heroku
git branch
# [should now be on 'origin/heroku']
# Now merge 'master' (theirs) into 'heroku' (ours)...
git merge -s ours master
# [should see a commit message, and then 'Merge made by the 'ours' strategy.']
# Now switch back to the master branch...
git checkout master
# ... and merge 'heroku' into 'master'
git merge origin/heroku
# Check git log to ensure recent changes to dev branch are in place...
git log
# Push up to master
git push
```

This version is when the development branch, in this case **heroku** again, already exists locally and we want to merge it all into **master**.

```
# checkout the master branch...
git checkout master
# bring it up to date
git pull
# Make sure we are on the master branch...
git branch
# [should be on 'master']
# Check git log, should only see old stuff...
git log
# Switch back to dev branch...
git checkout heroku
# [should see something like this as the local dev branch is the latest] 
Switched to branch 'heroku'
Your branch is up-to-date with 'origin/heroku'.
# Check log again as a reference, should see latest work...
git log

# From here steps are similar to above...
git merge -s ours master
git checkout master
git merge heroku
git log
git push
```

Nice!


