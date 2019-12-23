# my-watchlist-go-redis-demo

### Deploy redis
```
kubectl create -f redis-deployment.yaml
```

### Deploy frontend
```
kubectl create -f watchlist-app-deployment.yaml
```

### Cleanup
```
kubectl delete -f redis-deployment.yaml
kubectl delete -f watchlist-app-deployment.yaml 
```

-------------

```
#git cheatsheets
git config
git git remote add master <repo-git-url>
git add -A     # add all files or use 'git add filename'
git commit -am "comment" 
git push -u origin master

```