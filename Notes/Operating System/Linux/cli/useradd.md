---
createTime: 2024-11-15 10:11
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# useradd

## 0x01 Preface

`useradd` 是 linux 上一个用于新增用户的命令

```
 useradd [options] LOGIN
```

## 0x02 Positional Args

> [!important]
> 不要使用 numeric 用户名，例如 `0x00`，`0xfe`
> 在一些 OS 上虽然可以使用 numberic 创建用户，但是会导致一些应用出现问题，例如 `ps -u 0x00` 就会显示 root 相同的进程(`0x00 == hex(0)`)

- `LOGIN`
	用户名

## 0x03 Optional Args

> [!note]
> 只记录常用的参数，具体看 manual page

- `-D`
	
	使用 `/etc/default/useradd` 中的值创建用户，如果没有指定 `-D` 会使用命令行指定的其他参数和 `/etc/default/useradd` 中的值一起创建用户
	
- `-g | --gid <GROUP>`
	
	指定用户的用户组(GROUP 必须要存在)，如果没有指定会根据 `/etc/login.defs` 中的 `USERGROUPS_ENAB` 来判断
	
	- `if yes == USERGROUPS_ENAB` 
		
		那么就会生成一个和用户名相同的用户组
		
	- `if no == USERGROUPS_ENAB`
		
		那么就会使用 `/etc/default/useradd` 中的 `GROUP` 来生成(通常是 users 对应 id 1000)	
		
- `-G | --groups <GROUP1>[,GROUP2,...]`
	
	指定用户还需要加入的 GROUPS，默认只有用户名对应的用户组。还可以通过 `/etc/default/useradd` 中的 `GROUPS` 来指定
	
- `-U | --user-group`
	
	生成一个和用户名相同的用户组，并将用户加入用户组
	
- `-N | --no-user-group`
	
	创建用户时不会生成一个和用户名相同的用户组
	
- `-s | --shell <SHELL>`
	
	指定用户的 login shell，如果没有指定默认使用 `/etc/default/useradd` 中的 `SHELL` 默认为 `/usr/bin/bash`
	
- `-u | --uid <UID>`
	
	指定用户使用的 UID，默认会从 `/etc/login.defs` 中的 `UID_MIN` 开始计算
	
- `-e | --expiredate <YYYY-MM-DD>`
	
	指定用户的过期时间，如果没有指定会使用 `/etc/default/useradd` 中的 `EXPIRE` 默认为空，表示永不过期
	
- `-f | --inactive <NUMBER>`
	
	指定用户在 NUMBER days 后需要修改密码，如果没有指定会使用 `/etc/default/useradd` 中的 `INACTIVE` 默认为 -1，表示永不需要修改
	
- `-c | --comment <COMMENT>`
	
	指定用户的备注信息
	
- `-k | --skel <SKEL_DIR>`
	
	指定用户 home 目录的模板，创建用户 home 目录时会将 `SKEL_DIR` 中的内容复制到用户 home 目录，只有在和 `-m` 一起使用时生效
	
- `-b | --base-dir <BASE_DIR>`
	
	当没有使用 `-d` 时，指定用户的 home 目录在 BASE_DIR 创建，如果没有指定默认使用 `/etc/default/useradd` 中的 `HOME` 默认为 `/home`
	
- `-d | --home-dir <HOME_DIR>`
	
	使用 `HOME_DIR` 作为用户的 login directory，默认使用 `/BASE_DIR/LOGIN`。如果 login directory 不存在，默认会创建，除非指定了 `-M`
	
- `-m | --create-home`
	
	如果不存在用户的 home 目录就创建，同时 `SKEL_DIR` 中的内容会被复制到这个创建的用户 home 目录
	
- `-M | --no-user-home`
	
	创/建用户时不会创建 `${HOME}`，需要 `/etc/login.defs` 中的 `CREATE_HOME` 为 yes

## 0x04 Examples

### 0x04a useradd cc

如果只指定了 LOGIN 用户名为 cc，则会

1. 创建一个 cc 用户
2. 创建一个 cc 用户组
3. 生成 `/home/cc` 用户 home directory
4. `/usr/bin/bash` 作为用户的 login shell
5. 用户的 UID 从 1000 开始计算，取 `/etc/passwd` 中可用的 UID
6. 用户 cc 永远不会过期
7. 用户 cc 的密码永远不会提示修改

### 0x04b useradd nginx -M -s /usr/bin/nologin

> [!note]
> `nologin` 表示用户不允许登入，如果使用 `su`，`ssh` 等命令时会提示 `This account is currently not available.
   `，不会提供 login shell

如果使用了 title 中的命令，则会

1. 创建一个 nginx 用户
2. 创建一个 nignx 用户组
3. 不会生成用户 home directory
4. `/usr/bin/nologin` 作为用户的 login shell
5. 用户的 UID 从 1000 开始计算，取 `/etc/passwd` 中可用的 UID
6. 用户 nignx 永远不会过期

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- `man useradd`
- `man nologin`
- `man login.defs`

***References***


