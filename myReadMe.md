[TOC]

## github url
git@github.com:1569501393/gva.jieqiangtec.com.git

```shell
修改git url
root@SKY-20191125ANA:/mnt/d/www/qldev/go/gin-vue-admin# git remote set-url origin git@github.com:1569501393/gva.jieqiangtec.com.git

root@SKY-20191125ANA:/mnt/d/www/qldev/go/gin-vue-admin# git branch -avv
* master                        cc62194b [origin/master: behind 11] Merge pull request #773 from flipped-aurora/update
  remotes/origin/HEAD           -> origin/master
  remotes/origin/gva_gormv2_dev 10da0881 插件模式代码优化
  remotes/origin/gva_vue2       10da0881 插件模式代码优化
  remotes/origin/gva_workflow   354013cc Merge branch 'master' of https://github.com/flipped-aurora/gin-vue-admin into gva_workflow
  remotes/origin/master         5b6f6abb Merge pull request #778 from flipped-aurora/update
  remotes/origin/pgsql          ff564af5 posql 连接成功 初始化未完成
  remotes/origin/update         ec1e305a fix:#768

```

## 
```shell
root@SKY-20191125ANA:/mnt/d/www/qldev/go/gin-vue-admin# git remote show origin
* remote origin
  Fetch URL: git@github.com:flipped-aurora/gin-vue-admin.git
  Push  URL: git@github.com:flipped-aurora/gin-vue-admin.git
  HEAD branch: master
  Remote branches:
    gva_gormv2_dev tracked
    gva_vue2       tracked
    gva_workflow   tracked
    master         tracked
    pgsql          tracked
    update         tracked
  Local branch configured for 'git pull':
    master merges with remote master
  Local ref configured for 'git push':
    master pushes to master (local out of date)
root@SKY-20191125ANA:/mnt/d/www/qldev/go/gin-vue-admin#
```


## error
========
8:45:44 PM [vite] http proxy error:
Error: socket hang up
at connResetException (node:internal/errors:691:14)
at Socket.socketOnEnd (node:_http_client:471:23)
at Socket.emit (node:events:402:35)
at Socket.emit (node:domain:475:12)
at endReadableNT (node:internal/streams/readable:1343:12)
at processTicksAndRejections (node:internal/process/task_queues:83:21) (x6)
