# pam_pwquality

## 0x01 Overview

> The action of this module is to prompt the user for a password and check its strength against a system dictionary and a set of rules for identifying poor choices.

pam_pwquality.so 模块用于用户的密码校验

## Module_arguments

- debug

  如果指定该参数，会将 debug 信息写入 syslog

- retry=N

  Prompt user at most N times before returning with error. The default is 1.

- difok=N

  This argument will change the default of 5 for the number of changes in the new password from the old password.

  新密码必须要和原密码至少 N 个字符不同，默认 5

- minlen=N

  The minimum acceptable size for the new password 

  默认 9

- dcridit=N

  if N >= 0 表示密码中最多能出现 digits 的数量

  if N < 0 表示密码中至少出现 digits 的数量

  默认 1

  