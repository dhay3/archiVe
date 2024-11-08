---
author: "0x00"
createTime: 2024-06-04-17:49
draft: true
---

# MySQL 02 - Options

## 0x01 Overview[^1]

在 MySQL options 特指 command line option，可以通过 `program --help` 查看(除 `mysqld` 外，还需要使用 `--verbose`)

有如下规则

- Options are given after the command name.

- An option argument begins with one dash or two dashes, depending on whether it is a short form or long form of the option name. Many options have both short and long forms. For example, -? and --help are the short and long forms of the option that instructs a MySQL program to display its help message.

- Option names are case-sensitive. -v and -V are both legal and have different meanings. (They are the corresponding short forms of the --verbose and --version options.)

- Some options take a value following the option name. For example, -h localhost or --host=localhost indicate the MySQL server host to a client program. The option value tells the program the name of the host where the MySQL server is running.

- For a long option that takes a value, separate the option name and the value by an = sign. For a short option that takes a value, the option value can immediately follow the option letter, or there can be a space between: -hlocalhost and -h localhost are equivalent. An exception to this rule is the option for specifying your MySQL password. This option can be given in long form as --password=pass_val or as --password. In the latter case (with no password value given), the program interactively prompts you for the password. The password option also may be given in short form as -ppass_val or as -p. However, for the short form, if the password value is given, it must follow the option letter with no intervening space: If a space follows the option letter, the program has no way to tell whether a following argument is supposed to be the password value or some other kind of argument. Consequently, the following two commands have two completely different meanings:

  long options 需要使用 equal(`=`) 分隔 option 和 value，例如

  ```
  #correct
  --user=root
  #wrong
  --user root
  ```

  short options 可以直接跟上 value，例如

  ```
  #correct 
  -uroot
  #correct
  -u root
  ```

  但是 MySQL password 除外，password 和 short option 之间不能有空格

  ```
  #correct
  -p'password'
  #wrong
  -p 'password'
  ```


- Within option names, dash (-) and underscore (_) may be used interchangeably in most cases, although the leading dashes cannot be given as underscores. For example, `--skip-grant-tables` and `--skip_grant_tables` are equivalent.

  option names 中的下划线和短横线可以互相替换

- For options that take a numeric value, the value can be given with a suffix of K, M, or G to indicate a multiplier of 1024, $1024{^2}$ or $1024{^3}$​. As of MySQL 8.0.14, a suffix can also be T, P, and E to indicate a multiplier of 10244, 10245 or 10246. Suffix letters can be uppercase or lowercase. 

  如果 options 的值必须是数值，可以使用 unit suffix

  


---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:https://dev.mysql.com/doc/refman/8.4/en/command-line-options.html

