# Nginx regex

## 0x01 Overview

Nginx 使用 Perl Compatiable Regular Expression(PCRE)

| M-c    | Description                                                  | Example All the if statements return a TRUE value.           |
| ------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| **.**  | Normally matches any character except a newline. Within square brackets the dot is literal. | `if ("Hello World\n" =~ m/...../) {  print "Yep"; # Has length >= 5\n"; } ` |
| ( )    | Groups a series of pattern elements to a single element.  When you  match a pattern within parentheses, you can use any of $1, $2, ... later to refer to the previously matched pattern. | `if ("Hello World\n" =~ m/(H..).(o..)/) {  print "We matched '$1' and '$2'\n"; } `Output:`We matched 'Hel' and 'o W'; ` |
| +      | Matches the preceding pattern element one or more times.     | `if ("Hello World\n" =~ m/l+/) {  print "One or more \"l\"'s in the string\n"; } ` |
| ?      | Matches the preceding pattern element zero or one times.     | `if ("Hello World\n" =~ m/H.?e/) {  print "There is an 'H' and a 'e' separated by ";  print "0-1 characters (Ex: He Hoe)\n"; } ` |
| ?      | Modifies the *, +, or {M,N}'d regexp that comes before to match as few times as possible. | `if ("Hello World\n" =~ m/(l.+?o)/) {  print "Yep"; # The non-greedy match with 'l' followed  # by one or more characters is 'llo' rather than 'llo wo'. } ` |
| *      | Matches the preceding pattern element zero or more times.    | `if ("Hello World\n" =~ m/el*o/) {  print "There is an 'e' followed by zero to many ";  print "'l' followed by 'o' (eo, elo, ello, elllo)\n"; } ` |
| {M,N}  | Denotes the minimum M and the maximum N match count.         | `if ("Hello World\n" =~ m/l{1,2}/) { print "There is a substring with at least 1 "; print "and at most 2 l's in the string\n"; } ` |
| [...]  | Denotes a set of possible character matches.                 | `if ("Hello World\n" =~ m/[aeiou]+/) {  print "Yep"; # Contains one or more vowels } ` |
| \|     | Separates alternate possibilities.                           | `if ("Hello World\n" =~ m/(Hello|Hi|Pogo)/) {  print "At least one of Hello, Hi, or Pogo is ";  print "contained in the string.\n"; } ` |
| \b     | Matches a word boundary.                                     | `if ("Hello World\n" =~ m/llo\b/) {  print "There is a word that ends with 'llo'\n"; } ` |
| \w     | Matches an alphanumeric character, including "_".            | `if ("Hello World\n" =~ m/\w/) {  print "There is at least one alphanumeric ";  print "character in the string (A-Z, a-z, 0-9, _)\n"; } ` |
| \W     | Matches a non-alphanumeric character, excluding "_".         | `if ("Hello World\n" =~ m/\W/) {  print "The space between Hello and ";  print "World is not alphanumeric\n"; } ` |
| \s     | Matches a whitespace character (space, tab, newline, form feed) | `if ("Hello World\n" =~ m/\s.*\s/) {  print "There are TWO whitespace characters, which may";  print " be separated by other characters, in the string."; } ` |
| \S     | Matches anything BUT a whitespace.                           | `if ("Hello World\n" =~ m/\S.*\S/) {  print "Contains two non-whitespace characters " .        "separated by zero or more characters."; } ` |
| \d     | Matches a digit, same as [0-9].                              | `if ("99 bottles of beer on the wall." =~ m/(\d+)/) {  print "$1 is the first number in the string'\n"; } ` |
| \D     | Matches a non-digit.                                         | `if ("Hello World\n" =~ m/\D/) {  print "There is at least one character in the string";  print " that is not a digit.\n"; } ` |
| ^      | Matches the beginning of a line or string.                   | `if ("Hello World\n" =~ m/^He/) {  print "Starts with the characters 'He'\n"; } ` |
| $      | Matches the end of a line or string.                         | `if ("Hello World\n" =~ m/rld$/) {  print "Is a line or string ";  print "that ends with 'rld'\n"; } ` |
| \A     | Matches the beginning of a string (but not an internal line). | `if ("Hello\nWorld\n" =~ m/\AH/) {  print "Yep"; # The string starts with 'H'. } ` |
| \Z     | Matches the end of a string (but not an internal line).      | `if ("Hello\nWorld\n"; =~ m/d\n\Z/) {  print "Yep"; # Ends with 'd\\n'\n"; } ` |
| [^...] | Matches every character except the ones inside brackets.     | `if ("Hello World\n" =~ m/[^abc]/) {  print "Yep"; # Contains a character other than a, b, and c. } ` |



通常会出现在如下几个 directives 中

1. `location`
2. `map`
3. `rewrite`



**references**

[^1]:https://en.wikibooks.org/wiki/Regular_Expressions/Perl-Compatible_Regular_Expressions