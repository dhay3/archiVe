# Shell Style Guide

## Rules

### When to use shell

1. Small utilities or smiple wrapper scripts
2. Threading will not be undertaken
3. Non-straightforward control flow logic

### Which shell to use

1. Bash is the only shell scripting language

## Standards

### Hashbang

1. Executable must start with `#!/bin/env bash`
2. Libraries no need start with hashbang

### File extension

1. Executables should have no extension
2. Libraries must have a `.sh` extension and should not be executable 

### Permission

1. SUID/SGID are not allowed，use `sudo` instead

   ```
   chmod u+s or chmod 4xxx are forbidened
   chmod g+s or chmod 2xxx are forbidened
   ```

### Errors

1. All error messages should go to `stderr` (a function to print out error message is recommended)

   ```
   function err() {
     echo "[$(date +'%Y-%m-%dT%H:%M:%S%z')]: $*" >&2
   }
   ```

### Comments

1. Start each file with a description of its contents

   eg

   ```
   #!/bin/bash
   #
   # @Author: 0x00
   # @Date: Thu Nov 16 01:53:59 PM CST 2023
   # @Description: Perform hot backups of Oracle databases.
   ```

2. Any function that is not both obvious and short must be commented(Any function in a library must be commented regardless of length or complexity.)

   All function comments should describe the intended API behaviour using:

   - Description of the function.
   - Globals: List of global variables used and modified.
   - Arguments: Arguments taken.
   - Returns: Returned values other than the default exit status of the last command run.

   ```
   #######################################
   # Cleanup files from the backup directory.
   # Globals:
   #   BACKUP_DIR
   #   ORACLE_SID
   # Arguments:
   #   None
   # Returns:
   #		None
   #######################################
   function cleanup() {
     …
   }
   ```

3. Use TODO comments for code that is temporary, a short-term solution, or good-enough but not perfect

   ```
   #TODO: Handle the unlikely edge cases
   ```

4. Avoid Here String comments

   ```
   << comment
   Delete a file in a sophisticated manner.
   comment
   ```

### Formatting

> Choose a proper IDE
>
> eg. pycharm + Shell script plugin

1. Indent 2 spaces. No tabs.

2. Use blank lines between blocks.

3. Maximum line length is 80 characters.

4. It’s better to use here string when paragraphs should be shown, eg. helper messages

   ```
   cat <<END
   I am an exceptionally long
   string.
   END
   ```

5. Pipelines should be split one per line if they don’t all fit on one line. As well as logical compounds (`||` and `&&`)

   ```
   # Long commands
   command1 \
     | command2 \
     | command3 \
     | command4
   ```

6. Put `; do` and `; then` on the same line as the `while`, `for` or `if`

   ```
   for dir in "${dirs_to_cleanup[@]}"; do
     if [[ -d "${dir}/${ORACLE_SID}" ]]; then
       log_date "Cleaning up old files in ${dir}/${ORACLE_SID}"
       rm "${dir}/${ORACLE_SID}/"*
       if (( $? != 0 )); then
         error_message
       fi
     else
       mkdir -p "${dir}/${ORACLE_SID}"
       if (( $? != 0 )); then
         error_message
       fi
     fi
   done
   ```

7. `;;` should be on separate lines in `case` statement

   ```
   case "${expression}" in
     a)
       variable="…"
       some_command "${variable}" "${other_expr}" …
       ;;
   esac
   ```

### Quoting

1. ‘Single’ qutoe strings that represents shell meta charcters literal

   ```
   echo '$$'
   echo '$?'
   ```

2. “Double” qutoe strings that contains substitutions(vairable substitutions, command substitution, etc)

   ```
   echo "${HOME}"
   echo "$(whoami)"
   ```

   or any built-in variables

   ```
   echo "${1}"
   echo "${@}"
   echo "${*}"
   ```

3. Prefer quoting words(double quotes first), not compulsory

   ```
   university="tsinghua"
   ```

4. Spaces in strings quotes depend on the situation

   ```
   echo "a substitution in a double quotes: ${var}"
   echo '$ is a unit of currency'
   ```

   or paramters in strings

   ```
   grep -E 218.[[:digit:]]{3} /etc/resolv.conf
   grep -E 'nameserver 218.[[:digit:]]{3}' /etc/resolv.conf
   ```

5. Do not qutoe while shell expansion needed

   ```
   ls ./*
   #error
   ls "./*"
   ls './*'
   ```

   or regex

   ```
   #no match
   [[ "0x00" =~ '^[[:digit:]]x00$' ]] && echo 0x00
   # match
   [[ "0x00" =~ ^[[:digit:]]x00$ ]] && echo 0x00
   ```

   or path name(exclude in searching)

   ```
   ls /home/hugo
   ```

   or literal integers

   ```
   var=32
   ```

5. Quote command substitutions, even when you expect integers

   ```
   number="$(generate_number)"
   ```

### Variables

1. Curly brackets prefer `"${var}"` over `“$var”`

   in order to avoid ambiguous meanings

   eg `$10` and `${1}0`

   ```
   var="${1}"
   var="${PATH}"
   var="${a[@]}"
   ```

   include all built-in variables

   ```
   echo "${@}"
   echo "${?}"
   echo "${*}"
   ```

2. If a variable is global, use `declare -g` like `C` (it’s better to declare after the beginning comments)

   ```
   declare -g var=variable
   ```

   if a variable should be read only, and should be global accessed use `declare -rg` and uppercase all characters in the variable name like `final` in JAVA

   ```
   declare -rg TOKEN=SKwAzx33
   ```

   if a variable is an array and should be global accessed use `declare -ag`

   ```
   declare -ag array=(key1 key2 key3)
   ```

   if a variable is an map(dictionary) and should be global accessed use `declare -Ag`

   ```
   declare -Ag map=([key1]=value1 [key2]=value2 [key3]=value3)
   ```

3. All parameters in a function should be declared as a local

   ```
   function clean(){
   	local dir="${1}"
   	local filename="${2}"
   	...
   }
   
   clean "/tmp/" "*.log"
   ```

   as well as defined vairables

   ```
   function showArch(){
   	local arch="$(arch)"
   	...
   }
   ```

4. Declaring the loop variable as a local to avoid leaking into the global enviroment, if inside a function

   ```
   function readFile(){
   	...
   	local line
   	for line in "${file}"; do
   		echo "${line}"
   	done
   	...
   }
   ```

5. If a for loop isn’t defined under functions with local loop variable, unset the loop variable after the for loop to avoid the leaking

   ```
   for i in a b c; do
   	echo "${1}"
   done
   unset i
   ```

6. If a variable in functions should be defined as readonly use `local` and `readonly` perfer over `declare -rg`

   ```
   local FILE_PATH
   readonly FILE_PATH="/usr/bin"
   ```

### Arrays

1. An array is assigned using parentheses, and can be appended to with `+=(...)`

   ```
   flags=(--foo --bar='baz')
   flags+=(--greeting="Hello ${name}")
   ```

2. To display all elements in an array use `${array[*]}` rather than `${array[@]}` while tranverse isn’t neccessary

   ```
   array=(1 3 4 5 6)
   echo "${array[*]}"
   ```

### Pipelines

1. Pipeline create a subshell, so any variable modified whithin a pipeline do not progagate to the parent shell

   ```
   last_line='NULL'
   your_command | while read -r line; do
     if [[ -n "${line}" ]]; then
       last_line="${line}"
     fi
   done
   
   # This will always output 'NULL'!
   echo "${last_line}"
   ```

   Using process substitution also creates a subshell. However, it allows redirecting from a subshell to a `while` without putting the `while` (or any other command) in a subshell.

   ```
   last_line='NULL'
   while read line; do
     if [[ -n "${line}" ]]; then
       last_line="${line}"
     fi
   done < <(your_command)
   
   # This will output the last non-empty line from your_command
   echo "${last_line}"
   ```

### Command subsitution

1. Use `$(command)` instead of backticks

   ```
   var="$(command $(command1))"
   ```

   to avoid escape `\` in nested command subsitution

   ```
   var="`command \`command1\``"
   ```

### Test

1. `[[...]]` is preferred over `[...]`, `test` and `/usr/bin/[`

   `[[...]]` allows for regular expression match while `[...]` do not

   ```
   if [[ "filename" =~ ^[[:alnum:]]+name ]]; then
     echo "Match"
   fi
   ```

2. Use `==` for equality rather than `=` even though both work

3. Use `-z` or `-n` for empty/non-empty strings, rather than filler characters

   ```
   #not recommended
   if [[ "${my_var}" == "" ]]; then
     do_something
   fi
   
   #recommended
   if [[ -z "${my_var}" ]]; then
     do_something
   fi
   ```

4. Do not use mathematical comparison symbol(eg. `>`, `<` , etc) in `[[...]]`

   ```
   #wrong
   if [[ "23" < 3 ]]; then
     echo true
   fi
   ```

5. Use `((...))` for mathematical comparison rather than `-lt`, `-gt`, etc. Thought all works

   ```
   if (( "${var}" > 3 )); then
     do_something
   fi
   
   if [[ "${var}" -gt 3 ]]; then
     do_something
   fi
   ```

### Mathematical

1. Use `((...))` or `$((...))` rather that `let` or `$[...]` or `expr`

```
s
```

### Wildcard expansion

### Function

1. In order to easily find the start of the program, put the main program in a function called `main` as the bottom of the script.

   ```
   main "$@"
   ```

   

**references**

[^1]:https://google.github.io/styleguide/shellguide.html

