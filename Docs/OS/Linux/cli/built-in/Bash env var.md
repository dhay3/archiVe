# Bash env var

man bash

### refer to User

- EUID

  the effictive user ID of the current user

  ```
  cpl in ~ λ echo $EUID
  1000
  cpl in ~ λ id 
  uid=1000(cpl) gid=1000(cpl) groups=1000(cpl),3(sys),90(network),98(power),150(wireshark),960(libvirt),963(docker),991(lp),998(wheel)
  ```

- UID

  the userID of the current user

- GROUPS

  an array variable containing the list of groups of which the current user is a member

### refer to Bash

- BASH

  expands to the full filename use to invoke this instance of bash

  ### mixed

- HOME

  th home directory of the current user

- SHELL

  expands to the full pathname to the shell

- PATH

  the search path for commands

- PPID

  the process ID of the shell’s parent

- PWD

  the currrent working directory

- RANDOM

  assigning a random value(between 0 and 32767)

- SECONDS

  the number of seconds since shell invocation is returned

  

