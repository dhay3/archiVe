# gitignore 失效

1. 先删除索引树中存储的内容`git rm -r --cached .`
2. 再次提交`git add . && git commit -m"msg"`