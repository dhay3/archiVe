---
createTime: 2025-06-19 17:27
license: cc by 4.0
tags: 
 - "#hash1" 
 - "#hash2"
---

# PAM 01 - Overview

## 0x01 Preface

PAM 全称 Pluggable Authentication Modules, 是 Linux 中一个用于鉴权的 centralized 框架，由多个模块组成（其实就是 shared libraries，也就意味着可以被其他应用调用）。RheL 大多数应用都基于 PAM 来进行 authentication 和 authroization

## 0x02 Configuration[^1][^2]

> 强烈不推荐直接修改对应的 PAM 配置文件，应该使用 `authconfig` 来配置

每个支持 PAM 的应用(命令)或者服务，在 `/etc/pam.d` 中都有一个对应的配置文件。例如 `login` 命令，对应 `/etc/pam.d/login`

PAM 配置文件由 directives 组成，dirctives 语法均相同，主要有 5 个字段

```
[service] type control module-path module-arguments
```

- service 
- type 可以将其抽象成模块的函数
- control 可以将其抽象成模块的函数返回值
- module-path 可以将其抽象成类
- module-arguments 可以将其抽象成函数的入参

service 的 module-path 模块可以调用 type 函数，入参为 module-arguments，返回值要匹配 control

> directives can be stacked, or placed upon one another, so that multiple modules are used together for one purpose

directives 可以有多个，执行顺序从上到下，是否执行下一条 directives 由 control_flag 来控制

```
module_interface1	control_flag1	module_name1 module_arguments1
module_interface2	control_flag2	module_name2 module_arguments2
module_interface3	control_flag3	module_name3 module_arguments3
```

伪代码逻辑如下

```
def function do(moudule_name,module_interface,module_argurments)
	r = moudule_name.module_interface(module_arguments)
	if match(control_flag,r) then
		do(moule_name_next,module_interface_next,module_argurments_next)
  else
  	return
```

### 0x02a Service

通常就是应用(命令)或者服务的名字，对应 `/etc/pam.d` 中的文件，必须小写

### 0x02a Type

> The **type** is the management group that the rule corresponds to. It is used to specify which of the management groups the subsequent module is to be associated with.

将其理解成 module_interface，即模块的函数

告诉特定 module 执行那种 authentication（一共 4 种）

- auth

  This module interface authenticates users. For example, it requests and verifies the validity of a password. Modules with this interface can also set credentials, such as group memberships. 

  校验用户的账号密码是否匹配(正确)

- account

  This module interface verifies that access is allowed. For example, it  checks if a user account has expired or if a user is allowed to log in  at a particular time of day. 

  校验用户的账户是否允许登陆

- password

  This module interface is used for changing user passwords.

  用于修改用户的密码

- session

  This module interface configures and manages user sessions. Modules with this interface can also perform additional tasks that are needed to  allow access, like mounting a user's home directory and making the  user's mailbox available.

  用于管理和配置用户的 sessions

一个模块也可以提供多种 module_interface，例如 pam_unix.so 同时具备这 4 种 module_interface		

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [PAM - ArchWiki](https://wiki.archlinux.org/title/PAM)
- [Chapter 10. Using Pluggable Authentication Modules (PAM) \| System-Level Authentication Guide \| Red Hat Enterprise Linux \| 7 \| Red Hat Documentation](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/system-level_authentication_guide/pluggable_authentication_modules)
- [10.2. About PAM Configuration Files \| System-Level Authentication Guide \| Red Hat Enterprise Linux \| 7 \| Red Hat Documentation](https://docs.redhat.com/en/documentation/red_hat_enterprise_linux/7/html/system-level_authentication_guide/pam_configuration_files)
- [How to Configure and Use PAM in Linux](https://www.tecmint.com/configure-pam-in-centos-ubuntu-linux/)

***References***

[^1]:[10.2. About PAM Configuration Files \| System-Level Authentication Guide \| Red Hat Enterprise Linux \| 7 \| Red Hat Documentation](https://docs.redhat.com/en/documentation/red_hat_enterprise_linux/7/html/system-level_authentication_guide/pam_configuration_files)
[^2]:`man pam.conf`