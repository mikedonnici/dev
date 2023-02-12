# Bash

refs:

- [Learn Bash the Hard Way](https://leanpub.com/learnbashthehardway)
    - [alt](https://www.softouch.on.ca/kb/data/Learn%20Bash%20the%20Hard%20Way.pdf)
- [The Bash reference Manual](https://www.gnu.org/software/bash/manual/html_node/)
- [Bash Hacker's Wiki](https://wiki.bash-hackers.org/)
- [The Bash Guide](https://guide.bash.academy/)

## Core

### Globbing

- Globbing is used for pathname expansion (ie file matching).
- It involves replacing syntactic _tokens_ with a situation-specific value for the token.

- See: <https://guide.bash.academy/expansions/?=Pathname_Expansion#p1.1.0_4>
- `*` matches everything (except dotfiles)
- `?` matches a single character
- `[abd]` matches `a`, `b`, or `d`
- `[a-d]` matches `a`, `b`, `c` or `d`
- `file[0-9]` matches `file01`, `file1` etc
- `~` expansion substitutes the current user's home directory
- `*` character works differently in the context of a regular expression vs a glob
- Globs are not expanded in either single or double quotes
- Bash also supports more advanced glob patterns
  called [extended globs](https://guide.bash.academy/expansions/?=Pathname_Expansion#p1.1.0_8)

### Variables

- `NAME=mike` - create a var, caps are general convention
- `echo $NAME` - dereference var with `$`
- Quoting is important:
    - `NAME=mike donnici` - shell will treat `donnici` as a program
    - `NAME='mike donnici'` - string with space must be quoted
    - `NAME="$FIRST $LAST"` - only double quotes translate vars to values
- `readonly VAR=value` - set a readonly var
- vars are local to current shell unless exported
- `export VAR=value` - export var to global environment
- `env` - output shell vars
- `compgen -v` - generate list of possible word completions (-v current context)

Arrays in bash are zero-indexed and even simple variables can be treated as an array with one element.
So all bash vars are really arrays.

- `echo $BASH_VERSINFO` - outputs first element if no index given
- `echo $BASH_VERSINFO[0]` - doesn't work, interprets [0] as appended string
- `echo ${BASH_VERSINFO[0]}` - dereferences 'BASH_VERSINFO[0]' properly

### Functions, Builtins, Aliases and Programs

- 4 types of _commands_ in bash:
    - _function_
    - _builtin_
    - _alias_
    - _program_

#### Functions

```bash
function fn {
  echo "fn has 1st arg $1"
  echo "fn has 2st arg $2"
}
fn a b
```

- var scope:
    - vars declared inside a function have function scope
    - vars declared outside a function are accessible inside the function
    - `local` builtin can be used to create a locally scoped var with the same name as a var from the outer scope

```bash
OUTER="outer scope"
function fn {
  INNER="inner scope"
  echo "INNER=$INNER"
  echo "OUTER=$OUTER"
  local OUTER="local version of OUTER"
  echo "OUTER=$OUTER"
}
fn
echo $OUTER
```

- `unset -f fn` - to unset a function
- `unset` with no arg implies `-v` for var
- `declare -f` - outputs all functions and bodies in environment
- `declare -F` - same but only function names

#### Builtins

- Commands that come 'built in' to the bash shell program
- calling with `builtin` will error if command is _not_ a shell builtin
- eg `cd` and `builtin` itself
- creating a function called `cd` will be distinct from the builtin

```bash
function cd() {
  echo 'No!'
}
cd /tmp         # <-- invokes cd function
builtin cd /tmp # <-- invokes cd builtin
unset -f cd     # <-- unset the function
cd /tmp         # <-- now invokes builtin
```

#### Programs

- Programs are executable files
- eg `grep`, `sed`, `vi`
- `which` to find location of program

```bash
which which # <-- /usr/bin/which is a program
which cd    # <-- no path returned for builtin cd
```

#### Aliases

- `alias` sets a string that the shell translates to whatever the string is aliased to
- `unalias` removes the alias
- `type` builtin shows how a command will be interpreted by the shell

```bash
alias hello="echo hello"
unalias hello
type ls  # --> ls is aliased to `ls --color=auto`
type pwd # --> pwd is a shell builtin
```

### Redirects and pipes

- `>` is the _redirect_ operator for standard output
- `1>` is analogous to `>` as 1 is standard output file descriptor (see below)
- it takes the output from the previous command and sends it to the specified file:

```shell
echo "this text to file" > file1
```

- `>&` duplicate file descriptors, eg:

```bash
# cannot do this as will have 2 streams writing to a single file - garbled
ls > out.txt 2> out.txt
# so duplicate FD2 to whatever FD1 is pointing at:
ls 2&>1 out.txt  
```

- `<` redirect standard input (file descriptor 0) _from_ a file
- `>>` append stdout to a file

- '|' is the pipe operator
- it 'pipes' the standard output of the left command to the input of the right command:

```bash
cat file1 | grep -c file
``` 

#### Standard output vs standard error

- In POSIX systems (Unix / Linux) the source / sink for any data transfer is a file pointer, so everything is treated as
  a file - simple files, terminals, network interfaces etc
- A _file descriptor_ is a number associated with the process, that represents the "file"
- Data is read from / written to the appropriate file descriptor
- Each process gets 3 file descriptors by default
    - 0 _standard input_  # default is keyboard
    - 1 _standard output_ # default is terminal
    - 2 _standard error_  # default is terminal
- Examples:

```bash
echo "text"               #--> standard output (file descriptor 1) to terminal
echo "text" 1> file1      #--> redirect standard output (file descriptor 1) to file1
echo "text" > file1       #--> same as above (file descriptor 1 is default)
doesnotexist 2> /dev/null #--> standard error to black hole
doesnotexist 2> file1     #--> standard error to file1
doesnotexist 2>&1         #--> standard error to whatever 1 is pointing to at this time in the command
doesnotexist > out 2>&1   #--> stdout to file, stderr to &1 which is same file
```

- Redirections are evaluated from left to right, which can cause confusion.
- For example:

```bash 
ls nofile 2>&1 >out.txt 
```

- This looks like both std err and std out should go to the file, however, when 2>&1 is evaluated fd1 is connected to
  the terminal, so err will go to terminal and out.txt will be empty.
- To fix this:

```bash
ls nofile > out.txt 2>&1
```

- `&>` convenience operator for above, ie to redirect both stdout and stderr to a file
- `&>>` append version of same

### Here documents and here strings

- Way to feed larger block of text to a command
- Begin on a line after a delimiter, and end at delimiter on its own line
- Terminating delimiter cannot be indented, often `EOF` is used.
- Example:

```bash
cat <<.
Hello. My name is Mike.
I like to ride my bike.
.
```

- Can prefix delimiter with `-` which allows tabs (not spaces) to be used to indent text and also the terminating
  delimiter
- Can also use expansions in a heredoc:

```bash
#!/usr/bin/env bash
# Note MUST be tabs - spaces do NOT work
name='Mike Donnici'
cat <<-.
  His name is $name, and he likes to wrestle.
  This may seem weird to some.
  .
```

- Here strings are more concise, and generally preferred:

```bash
cat <<<"Hello, there.
How can I help you today?"
```

## Bash scripting

### Bash Startup

- `!#/bin/bash` - shell script starts with _shebang_ / _hashbang_
- `chmod +x file` to set executable flag
- `echo $PATH` to see paths that bash will check for executable script
- Paths are set by various bash start up scripts,
  see [here](https://blog.flowblok.id.au/2013-02/shell-startup-scripts.html)
- `env -i bash --noprofile --norc` will start bash without any startup scripts
- Shell scripts have `.sh` extension by convention, but not required.

### Command substitution

- Two ways to do command substitution (note double quotes):

```bash
echo "hostname = $(hostname)" # <-- dollar-bracket method
echo "hostname = `hostname`"  # <-- backtick method
```

- Combination on value expansion prefix `$` and a subshell process `(...)`
- Good practice to always double quote value expansion because spaces in the result will be _word-split_, whitespaces
  removed and hidden path expansion done for each word

- Nested commands are read from inside out
- Dollar bracket method is easier to read for nested command substitutions:

```bash
echo `echo \`echo \\\`echo inside\\\`\``
echo $(echo $(echo $(echo inside)))
```

### Parameters

- Temporary values stored in memory for use by bash scripts
- Can be read and written to like files, but memory is faster
- Easier and more powerful syntax that file redirection
- Types:
    - Shell variables
    - Positional parameters
    - Special parameters

#### Shell variables

- Bash parameter with a name
- Stored through _variable assignment_ - `name='Mike'`
- Accessed using _parameter expansion_ - `echo "Hello, $name"`
- `{}` can be used to denote start and end of parameter name, where required: `echo "${name}'s dog"`
- There are also
  more [advanced parameter expansion operators](https://guide.bash.academy/expansions/?=Parameter_Expansion#p2.2.2_3)
- Parameter expansions always leave the original value unchanged

#### Environment variables

- Variables kept in the _process_ environment
- Accessible to any process running on the system, not just Bash
- When a new program is run in bash, a new bash process is created with its own environment
- The new process is created with a _copy_ of the environment variables from the creating process
- Hence, changing vars in a child process does not affect the same var in the parent process

#### Shell initialization

- `~/.bashrc` generally contains standard shell initialization
- `~/.bash_profile` is generally used to export additional variables into the environment
- `~/.profile` is the generic shell equivalent so should only use POSIX `sh` syntax
- If a `~/.*profile` initialization file exists, bash does not look for `~/.bashrc`
- So, last line of the `*profile` should be set up to ensure `~/.bashrc` is read:

```bash
source ~/.bashrc
```

#### Positional parameters

- Variables are parameters with a name, _positional parameters_ are parameters with a number
- Expanded with the usual syntax: `$1`, `$2`, require `{}` for 2 or more digits, eg `${10}`
- Positional parameters expand to values that were sent into a process as arguments when it was created:

```bash
grep foo bar.txt
# $1 = 'foo', $2 = 'bar.txt'
```

- `$0` is the name of the process but may be set differently by different commands, so generally not used.
- In all other ways positional parameters can be treated like normal variables

#### Special parameters

- "$*" - expands all params to a single string joined by spaces
- "$@" - expands positional params as a list of separate arguments
- "$#" - expands to the number of positional params
- "$?" - expands to exit code of last command
- "$-" - expands to the set of option flags active in current shell
- "$$" - expands to process id of current shell process
- "$!" - expands to process id of most recently executed background command
- "$_" - expands to last argument of the previous command

#### Arrays

- `=()` assignment operator used to create arrays
- spaces are permitted _inside_ the braces and are used to separate array elements
- `[@]` suffix expands an array to as list of separate arguments
- very important to wrap array expansion in double quotes

```bash
list=( one two "and a three" )
echo "${list[@]}"
```

- array operations:
    - `+=( ... )` - append to array
    - `files=( *.txt )` - expand glob pattern into an array
    - `echo "${files[0]}"` - expand a single item
    - `echo "${files}"` - expands _only_ the first item
    - `unset "files[2]"` - remove the third item
    - `[*]` - merges all array elements into a single string, separated by `IFS` (Internal Field Separator) char, which
      is a space by default. Generally `[@]` should be used in favour of `[*]`.

- special parameter expansions:
  - `"${param[@]/pattern/replacement}"` will apply replacement to each item in the array
  - `"${#names[@]}"` expands to count of elements in the array
  - `"${#names[1]}"` expands to the string length of the second item in the array
  - `"${names[@]:1:2}` obtains a slice of the array


### Exit codes

| code   | meaning                 | notes                                                                      |
|--------|-------------------------|----------------------------------------------------------------------------|
| 0      | OK                      | Command successfully run                                                   |
| 1      | General error           | Error with no specific reserved number                                     |
| 2      | Misuse of shell builtin | Problem when running builtin command                                       |
| 126    | Cannot execute          | Permission problem or command is not executable                            |
| 127    | Command not found       | No file found matching the command                                         |
| 128    | Invalid exit value      | Exit argument given (eg exit 1.76)                                         |
| 128+n  | Signal ‘n’              | Process killed with signal ‘n’, eg 130 = terminated with CTRL-c (signal 2) |

- Not all commands use exit codes in the same way
- Eg `grep` uses exit code 1 to signal no match found rather than a program error
- Exit code for a function can be set using `return` builtin:

```bash
function trycmd {
  $1 # <-- execute first arg
  if [[ $? -eq 127  ]]
  then
    echo 'What?'
    return 1
  fi  
}
trycmd ls
trycmd notexist
```

### Tests / conditionals

- Tests are a way of writing expressions that are true or false
- Examples:

```bash
[ 1 = 0 ] # $? = 1 (false)
[ 1 = 1 ] # $? = 0 (true)
```

- Can use either single or double `=`
- `==` probably better to distinguish from variable assignment

```bash
[ 1 == 0 ] # $? = 1 (false)
[ 1 == 1 ] # $? = 0 (true)
```

- `[` is actually a builtin so space after is required
- `test` is synonymous with `[ .. ]` - `man [` and `man test` go to same page

```bash
A=1
[ $A == 1]
test $A == 1
[ $A == 2 ]
test $A == 2
```

- Logic operators `!` not, `&&` and, `||` or, require separate `[]` pairs
- `()` evaluated first

```bash
( [ 1 = 1 ] || [ ! '0' = '0' ] ) && [ '2' = '2' ]
echo $?
```

- Can use `-o` (or), `-a` (and) in one set of `[]` but can't use `()` and not often used

```bash
[ 1 = 1 -o ! '0' = '0' -a '2' = '2' ]
echo $?
```

- `[[` operator is similar to `[`, difference is subtle

```bash
unset NOVAR
[ ${NOVAR} == '' ] # bash: [: ==: unary operator expected
```

- This is processed as `[ == '' ]`, hence the error
- With `[[` the missing var is substituted with an empty string:

```bash
unset NOVAR
[[ ${NOVAR} == '' ]] #--> 0 
# processed as...
[ '' == '' ] #--> 0
```

- Could protect against unset var like this:

```bash
[ "x${NOVAR}" == "x" ]
```

- Or, just use `[[`
- More
  detail [here](https://serverfault.com/questions/52034/what-is-the-difference-between-double-and-single-square-brackets-in-bash)

#### Unary operators

- See man pages for all
- `-z` true is string is empty
- `-n` true if string is _not_ empty
- `-a` true if file exists (file or dir)
- `-d` true if file exists and is a dir

```shell
[ -z "abc" ] #--> false
[ -z "" ]    #--> true
[ -n "abc" ] #--> true
[ -n "" ]    #--> false

touch tmpfile
mkdir tmpdir
[ -a tmpfile ] #--> true
[ -d tmpfile ] #--> false
[ -a tmpdir ]  #--> true
[ -d tmpdir ]  #--> true
``` 

#### Binary operators and types

- Types can be important when testing with binary operators
- Eg, `<` and `>` expect strings

```bash
[ 10 < 2 ]       #--> true
[ '10' < '2' ]   #--> true
[[ 10 < 2 ]]     #--> true
[[ '10' < '2' ]] #--> true
```

- Numerical comparisons use:
    - `-lt` less than
    - `-gt` greater than
    - `-eq` equal to
    - `-ne` not equal to

#### if statements

- `if`, `then`, `elif`, `else`, `fi`

```bash
if [[ 10 -gt 11 ]]
then
  echo "10 > 11"
elif [[ 1 -eq 0 ]]
then
  echo "1 = 0"
else
  echo "nada"
fi
```

- One liner

```bash
if [[ 10 -gt 9 ]]; then echo "YES"; fi
```

- `if` does not require `[[]]`

```bash
if grep not_there /dev/null; then echo "found"; else echo "not found"; fi
```

### Loops

- Bash has a few loop formats
- 'C'-style using double parenthesis
- Note: var does not need `$` inside `(())`

```bash
for (( i=0; i < 3; i++))
do
  echo $i
done  
```

- `for` - `in` style:

```bash
for i in 1 2 3 # <-- these could be files or command, eg $(ls *.txt)
do 
  echo $i
done  
```

- `while` - `until` style:

```bash
n=0
while [[ $n -lt 3 ]]
do
  ((n++))
  echo "$n"
done  
```

- infinite loop with `break` condition

```bash
n=0
while true
do 
  ((n++))
  echo $n
  if [[ $n -eq 4 ]]
  then
    break
  fi
done  
```

- `case` statements, often used to process command-line args

```bash
a=2
case "$a" in
1) echo 'a is 1'; echo 'ok';; # <-- double ;; indicates a following matching case statement
2) echo 'a is 2'; echo 'ok';;
*) echo 'a is unmatched'; echo 'failure';; #<-- *) default match all 
esac
```

- Processing command-line opts ina script with `builtin getopts`

```bash
#!/bin/bash
# case.sh
# flags are: -a, -b 'option', -c 'option'
while getopts "ab:c:" opt
do
  case "$opt" in
  a) echo '-a flag';;
  b) echo "-b with arg ${OPTARG}";;
  c) echo "-c with arg ${OPTARG}";;
  esac
done
```

```shell
./case.sh -a -b hello -c goodbye
-a flag
-b with arg hello
-c with arg goodbye
```

- How does `${OPTARG}` get the different values?

### The `set` command

- `builtin set` provides a way to manage bash options
- Running `set` on its own outputs all variables and functions set in environment
- `set -o` shows current state of all options

```bash
set -o       #<-- show all options current state
set -o posix #<-- turn ON posix mode
set +o posix #<-- turn OFF posix mode
```

- difference between `set` and `env` is that `env` shows _exported_ variables, not all vars set in shell
- `set` becomes most useful when scripting

```bash
set -o errexit # exit script if any command fails
set -e         # flag form of set -o errexit 
set -o xtrace  # output each commands as it is being run
set -x         # flag form of set -o xtrace 
set -o nounset # throw an error when unbound var is referred to
```

- `pipefail` option returns error code of the last command to return a non-zero status

```bash
function f1 { return 0; }
function f2 { return 111; }
function f3 { return 0; }
set +o pipefail # pipefail off...
f1 | f2 | f3    
echo $?         # ...will return 0 from f3
set -o pipefail # pipefail on...
f1 | f2 | f3    
echo $?         # ...will return 111 from f2
```

- `shopt` is another builtin to set bash vars

[UP TO]

### File substitution

### Internal field separator 