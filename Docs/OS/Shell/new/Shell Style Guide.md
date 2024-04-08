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

### Loops

1. Put `; do` and `; then` on the same line as the `while`, `for` or `if`

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

2. When index are needed use `for ((exp1;exp2;exp3))`

   ```
   local -i i
   for ((i=0;i<5;i++));do
   	echo "${i}"
   done
   ```

### Case

1. `;;` should be on separate lines in `case` statement

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

6. Quote command substitutions, even when you expect integers

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

   if a variable should be exported(enviroment) use `declare -x`

   ```
   declare -xr ORACLE_SID='PROD'
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

7. Rembere to declare your variables as integers when possible, and to prefer local variables(`local` has the same options as `define`) over globals

   ```
   local -i hundred=$(10*10)
   declare
   ```

8. Variables names for loops should be similarly named for any variable you’re looping through

   ```
   for zone in "${zones[@]}"; do
     something_with "${zone}"
   done
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

### Exclamation mark

1. Use the exclamation mark to invert the exit status of a command

   ```
   $! echo true;echo $?
   true
   1
   ```

2. When `false` or `true` are used, `!` can not use with `test` in `if `(in other word `test` can be used with `!` in test without `false` or `true`)

   ```
   if ! false ;then echo 1;else echo 2;fi
   1
   if [[ ! false ]];then echo 1;else echo 2;fi
   2
   ```

### Pipelines

1. Pipeline create a subshell, so any variable modified whithin a pipeline do not progagate to the parent shell

   ```
   last_line='NULL'
   #create a subshell
   ls | while read -r line; do
     if [[ -n "${line}" ]]; then
       last_line="${line}"
       echo "$last_line"
     fi
   done
   #subshell end
   
   # This will always output 'NULL'!
   echo "${last_line}"
   ```
   
   Using process substitution also creates a subshell. However, it allows redirecting from a subshell to a `while` without putting the `while` (or any other command) in a subshell.
   
   ```
   last_line='NULL'
   while read line; do
     if [[ -n "${line}" ]]; then
       last_line="${line}"
       echo "$last_line"
     fi
   done < <(ls)
   
   # This will output the last non-empty line from your_command
   echo "${last_line}"
   ```

2. If a pipeline all fits on one line, it should be on one line.

   If not, it should be split at one pipe segment per line with the pipe on the newline and a 2 space indent for the next section of the pipe. This applies to a chain of commands combined using `|` as well as to logical compounds using `||` and `&&`.

   ```
   # All fits on one line
   command1 | command2
   
   # Long commands
   command1 \
     | command2 \
     | command3 \
     | command4
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

2. 'Use `==` for equality rather than `=` even though both work

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

6. Never use `-ne` or `-eq` for string comparsion, use `!=` or `==` instead

   ```
   a=aaa
   if [[ $a -ne bbb ]];then echo 1;else echo 2;fi
   2
   
   # in bash help test, it is use for Arithmetic
   # arg1 OP arg2 Arithmetic tests.  OP is one of -eq, -ne, -lt, -le, -gt, or -ge.
   ```

### Mathematical

1. `<` and `>` don’t perform numerical comparison inside `[[...]]`. For preference, don’t use `[[...]]` at all for numeric comparisons, use `((...))` instead

   ```
   #wrong
   if [[ "23" < 3 ]]; then
     echo true
   fi
   ```

2. Use `((...))` or `$((...))` rather that `let` or `$[...]` or `expr`

   ```
   let a=10+20
   ```

3. When using variables, the `${var}` or `$var` forms are not required within `$((...))`. The shell knows to look up `var` for you, and omitting the `${…}` leads to cleaner code. 

   ```
   a=10
   b=20
   c=$((a+b))
   ```

4. Calculation can be assigned to a variable

   ```
   j=1
   (( i = 10 * j + 400 ))
   echo "${i}"
   
   # the same
   i=$((10 * j + 400))
   echo "${i}"
   ```

5. Increment and decrement

   ```
   ((i++))
   ((i--))
   ```

### Wildcard expansion

1. Use an explicit path(absolute path) when doing wildcard expansion of filenames.

   ```
   rm -v *
   
   rm -v /home/cache
   ```

### Function

1. In order to easily find the start of the program, put the main program in a function called `main` as the bottom of the script.

   ```
   main "$@"
   ```


2. The keyword `function` is optional, but is recommended consistently throughout a project.

   ```
   function Mone(){
   	...
   }
   
   #the same
   Mone(){	
    ...
   }
   ```

3. Use camel-case as the function name

   ```
   function convert2Json(){
   	...
   }
   ```

4. If a function is in a package, separate package names with `::`

   ```
   # Part of a package
   mypackage::toString() {
     …
   }
   ```

5. Use `true` to return an successful result, use `false` to return an unsuccessful result

   ```
   #it's no need to use return keyword with true or false
   function man(){
   	local gender=$1
   	if [[ "${gender}" == 1 ]];then
   		 true
     else
     	 false
     fi
   }
   ```

### Calling commands

1. Always check return values and give informative return values

2. For unpiped commands, use `$?` or check directly via an `if` statement to keep it simple

   ```
   if ! mv "${file_list[@]}" "${dest_dir}/"; then
     echo "Unable to move ${file_list[*]} to ${dest_dir}" >&2
     exit 1
   fi
   
   # Or
   mv "${file_list[@]}" "${dest_dir}/"
   if (( $? != 0 )); then
     echo "Unable to move ${file_list[*]} to ${dest_dir}" >&2
     exit 1
   fi
   ```

3. For piped commands, the `$?` only check the last part of the pipeline

   ```
   3333 | echo 1;echo $?
   1
   bash: 3333: command not found
   # the $? alway will get 0
   0
   ```

   use `PIPESTATUS` get the return values

   ```
   tar -cf - ./* | ( cd "${DIR}" && tar -xf - )
   return_codes=( "${PIPESTATUS[@]}" )
   if (( return_codes[0] != 0 )); then
     do_something
   fi
   if (( return_codes[1] != 0 )); then
     do_something_else
   fi
   ```

### Null command

1. Use `:` (null command) as the `pass` keyword in python

   ```
   function todo(){
   	:
   }
   ```

### Command

1. Use `command` to check the command exsit or not

   ```
   function lib::core::commandExists() {
     command -v "$@" >& /dev/null
   }
   ```

**references**

[^1]:https://google.github.io/styleguide/shellguide.html
[^2]:https://stackoverflow.com/questions/41204576/exclamation-mark-to-test-variable-is-true-or-not-in-bash-shell

