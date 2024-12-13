# PAM Overview

## 0x01 Overview

PAM 全称 Pluggable Authentication Modules, 是 Linux 中一个用于鉴权的 centralized 机制，由多个模块组成（其实就是 shared libraries，也就意味着可以被其他应用调用）

Rhel 中大多数应用都基于 PAM 来进行 authentication 和 authroization

## 0x02 Configuration

> 强烈不推荐直接修改对应的 PAM 配置文件，应该使用 `authconfig` 来配置

每个支持 PAM 的应用(命令)或者服务，在 `/etc/pam.d` 中都有一个对应的配置文件。例如 `login` 命令，对应 `/etc/pam.d/login`

PAM 配置文件由 directives 组成，dirctives 语法均相同

```
module_interface	control_flag	module_name module_arguments
```

- module_interface 可以将其抽象成函数
- control_flag 可以将其抽象成 module_interface 的返回值
- module_name 可以将其抽象成含有 module_interface 的类
- module_arguments 可以

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

### module_interface

module_interface 也被称为 pam interface，告诉特定模块执行那种 authentication（一共 4 种）

- auth

  This module interface authenticates users. For example, it requests and  verifies the validity of a password. Modules with this interface can  also set credentials, such as group memberships. 

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

### control_flag

每个 module_interface 都会返回 success 或者 failure，control_flag 告诉 pam 如何处理结果

有如下几种常见 control_flag (具体查看 `mnan pam.d`)

- required

  The module result must be  successful for authentication to continue. If the test fails at this  point, the user is not notified until the results of all module tests  that reference that interface are complete. 

  必须为 successful		

  必要条件	

- requisite

  The module result must be successful for authentication to continue. However, if a test fails at this point, the user is notified immediately with a message reflecting the first failed `required` *or* `requisite` module test

  必须为 successful 否这立即告诉用户不满足

  必要条件

- sufficient

  The module result is ignored if it fails. However, if the result of a module flagged `sufficient` is successful *and* no previous modules flagged `required` have failed, then no other results are required and the user is authenticated to the service. 		

  如果结果为 failure 会被忽略，执行下面的 directives；如果结果为 successful，且之前的 required control_flag directives 结果都为 successufl,那么后面的 direcitve 不会被校验	

  充分条件

- optional

  The module result is ignored. A module flagged as `optional` only becomes necessary for successful authentication when no other modules reference the interface. 

  不校验结果，只有当被 `optional` 标注的 direcitve 被其他 direcitves 引用时才会被作为必要条件

- include

  Unlike the other controls,  this does not relate to how the module result is handled. This flag  pulls in all lines in the configuration file which match the given  parameter and appends them as an argument to the module

  和结果无关，和 Nginx 中的 include 类似读取指定配置文件的内容

这么看可能不明显，例如 `/etc/pam.d/setup`

```
auth       sufficient	pam_rootok.so
auth       include	system-auth
account    required	pam_permit.so
session	   required	pam_permit.so
```

- `auth sufficient	pam_rootok.so`

  This line uses the `pam_rootok.so` module  to check whether the current user is root, by verifying that their UID  is 0. If this test succeeds, no other modules are consulted and the  command is executed. If this test fails, the next module is consulted. 

- `auth include system-auth`

  This line includes the content of the `/etc/pam.d/system-auth` module and processes this content for authentication. 

- `account required pam_permit.so`

  this line uses the `pam_permit.so` module to allow the root user or anyone logged in at the console to reboot the system. 	

- `session required pam_permit.so`

  This line is related to the session setup. Using `pam_permit.so`, it ensures that the `setup` utility does not fail. 				

### module_name

所有可用的模块默认都在 `/lib64/security` 或者 `/lib/security` 下(其实就是 shared libraries)

如果需要查看具体的使用方法,可以使用 `man <module_name>`, 例如 `man pwquality`

#### How to find out if a program is PAM-aware

如何确定一个应用或者命令是否使用 PAM 呢？聪明的人已经想到可以使用 `ldd` 来读取应用调用来那些 shared libraries。如果有 `libpam.so` 就表示调用来 PAM 中的功能

```
(vagrant@localhost security]$ ldd /bin/passwd | grep libpam.so
        libpam.so.0 => /lib64/libpam.so.0 (0x00007f277b4fa000)
```

### module_arguments

一些模块还支持额外的参数，例如

```
password	requisite	pam_pwquality.so enforce_for_root retry=3
```

`pam_pwquality.so` 用于校验密码的强度，可以传入一些参数

- `enforce_for_root`

  root 的密码也必须满足 `/etc/security/pwquality.conf` 中的配置

- `retry=3`

  如果密码错误可以尝试 3 次

*Invalid arguments are generally ignored and do not otherwise affect the success or failure of the PAM module.*

如果对应的模块并没有指定的参数，PAM 就会忽略该参数（但是大多数都会通告到 `journald`）

## 0x03 Example

https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/system-level_authentication_guide/pam_configuration_files#ex.pam-config

**references**

[^1]:https://wiki.archlinux.org/title/PAM
[^2]:https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/system-level_authentication_guide/pluggable_authentication_modules
[^3]:https://www.tecmint.com/configure-pam-in-centos-ubuntu-linux/