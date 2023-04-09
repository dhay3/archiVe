# gh release

ref
[https://cli.github.com/manual/gh_release](https://cli.github.com/manual/gh_release)
管理 github release
github release 和 tag 的区别
[https://stackoverflow.com/questions/18506508/whats-the-difference-between-tag-and-release](https://stackoverflow.com/questions/18506508/whats-the-difference-between-tag-and-release)

## Query Subcommands
### List
syntax
```
gh release list [flags]
```
默认显示当前目录对应的 repository，也可以手动指定
```
λ ~/ gh release list -R spring-projects/spring-boot
```
#### Optional args

- `--exclude-drafts | exclued-pre-releases`

不包含 draft 或者 pre releases

- `-R | --repo <[HOST/]OWNER/REPO>`
### View
syntax
```
gh release view [<tag>] [flags]
```
用于查看 release 的详细信息，如果没有指定 tag 默认使用 latest
```
λ ~/ gh release view v4.28.2 -R v2ray/v2ray-core              
```
#### Optional args

- `-w | --web`
## Update Subcommands
### Create
syntax
```
gh release create [<tag>] [<files>...]
```
用于创建 release release
```
λ ~/ gh release create v0.0.1  -R dhay3/gitlab
? Title (optional) gh release test
? Release notes Leave blank
? Is this a prerelease? No
? Submit? Publish release
https://github.com/dhay3/gitlab/releases/tag/v0.0.1
```
#### Optional args

- `-d | --draft`

标识为 draft release

- `-n | --notes <string>`

添加 release notes

- `--generate-notes`

为 release 自动生成 title 和 notes

- `-p | --prerelease`

标识为 prerelease 

- `-t | --title <string>`

指定 release title

- `-R | --repos <[HOST/]OWNER/REPO>`
### Delete
syntax
```
gh release delete <tag> [flags]
```
默认删除当前目录对应的 repository release，也可以手动指定
```
λ ~/ gh release create v0.0.2 -R dhay3/gitlab 
? Title (optional) v0.0.2
? Release notes Write my own
? Is this a prerelease? No
? Submit? Publish release
https://github.com/dhay3/gitlab/releases/tag/v0.0.2
                                                                                                                                     
λ ~/ gh release delete v0.0.2 -R dhay3/gitlab --cleanup-tag
? Delete release v0.0.2 in dhay3/gitlab? Yes
✓ Deleted release and tag v0.0.2
```
#### Optional args

- `--cleanup-tag`

删除 release 是删除 tag，如果不指定改参数默认不会删除对应的 tag

- `-y | --yes`
- `-R | --repo <[HOST/]OWNER/REPO>`
### Download
syntax
```
gh release download [<tag>] [flags]
```
从 github release 下载 assets，如果没有指定 tag 默认使用 latest，如果需要特定的 assests 需要使用 `--pattern`
#### Optional args

- `-A | --archive <format>`

按照指定格式下载 source code ，`{zip|tar.gz}`

- `--clobber`

等同于 shell 中的 clobber，overwrite existing files of the same name

- `-D | --dir <directory>`

下载到指定目录

- `-p | --pattern <stringArray>`

只下载匹配 global pattern 的 assests

- `--skip-existing`

如果需要下载的文件已经存在就跳过

- `-R | --repo <[HOST/]OWNER/REPO>`
#### Examples
```
λ ~/ gh release view v4.28.2 -R v2ray/v2ray-core
λ ~/ gh release download v4.28.2 -p v2ray-linux-64.zip.dgst -R v2ray/v2ray-core
```
### Edit
syntax
```
gh release edit <tag>
```
用于修改 release 信息
#### Optional args

- `--draft`

将指定 release 转为 draft

- `--prerelease`

将指定 release 转为 prerelease

- `--tag <string>`

修改 tag name

- `-n | --notes <strings>`

修改 tag note (即 tag 的 README 内容)

- `-F | --notes-file <file>`

将 tag note 内容修改到 file 中的内容
#### Exmaples
```
Publish a release that was previously a draft
$ gh release edit v1.0 --draft=false

Update the release notes from the content of a file
$ gh release edit v1.0 --notes-file /path/to/release_notes.md
```
