# git add

ref
[https://git-scm.com/docs/git-add](https://git-scm.com/docs/git-add)

## Digest
syntax
```
git add [options] <pathspec...>
```
将 pathspec 中的内容加入到 staging area，pathspec 是一个支持 globbing 的变参
## Optional args

- `-n | --dry-run`

  不实际添加文件到 staging area，只测试

- `-v | --verbose`

  输出添加到 staging area 的文件名，默认不输出

## Examples
```
#将当前目录以及子目录下的内容加入到 staging area
git add .

#当前目录下，所有已 .sh 结尾的文件加入到 staging area
git add *.sh
```
