# Strings and regex

- Working with strings is an important part data analysis

- String slicing examples (note third index):

```python
s = "abcdefghijklmnopqrstuvwxyz"
len(s)
# 26
s[:]
# 'abcdefghijklmnopqrstuvwxyz'
s[:4]
# 'abcd'
s[7:]
# 'hijklmnopqrstuvwxyz'
s[-1:]
# 'z'
s[-5:-1]
# 'vwxy'
s[16:22]
# ''qrstuv'
s[10:20:2] # every second char in [10:20]
# 'kmoqsu'
s[::-1]    # every nth char, but in reverse!
# 'zyxwvutsrqponmlkjihgfedcba'
s[::-5]
# 'zupkfa'
```

## String formatting

- Three main ways: positional, formatted literals and template methods

### Positional formatting

- Positional formatting using `.format()`:

```python
"var1 = {}, var2 = {}".format(var1, var2)                   # in order passed to format()
"var1 = {1}, var2 = {0}".format(var1, var2)                 # in order specified by index value in {}
"var1 = {vee1}, var2 = {vee2}".format(vee1=var1, vee2=var2) # using named placeholders

# named placeholders from a dict - NOTE: index is specified WITHOUT quotes
pets = {"name": "Milo", "type": "Dog", "breed": "Kelpie"}
s = "The {data[type]} is a {data[breed]} who's name is {data[name]}"
s.format(data=pets)
```

- Format specifiers declare the data type that will be in the placeholder

```python
"The cost was ${0:f}".format(24.57)         # a float
# 'The cost was $24.570000'
"The cost was ${0:.2f}".format(145.001234563) # a float with 2 decimals
# 'The cost was $145.00'
```

- Datetime format placeholders also work

```python
"The date now is {:%d %B %Y}".format(datetime.now())
# 'The date now is 26 April 2020'

# Datetime with named placeholder
"Today is {today:%B %d, %Y}. It's {today:%H:%M}".format(today=datetime.now())
# "Today is April 26, 2020. It's 17:00"
``` 

### Formatted literals

- Minimal syntax using f-strings, ir add a `f` prefix - `f"literal string{expression}"`
 
```python
a = "A"
b = "B"
c = "C"
f"now I know my {a}{b}{c}"
# 'now I know my ABC'
```

- type conversions:
   - !s (string)
   - !r (string with quotes)
   - !a (like !r but with escaped non-ASCII chars)

```python
nickname = "Bazza"
f"His preferred nickname was {nickname!r}."
# "His preferred nickname was 'Bazza'."
```   

- format specifiers:
   - e (scientific notation)
   - d (digit)
   - f (float)

```python
price = 23.452625423462
f"current price is {price:.2f}"
# 'current price is 23.45'
```

- datetime specifiers

```python
today = datetime.now()
f"today is {today:%d %B, %Y}"
 'today is 26 April, 2020'
```

- with dicts (note: this time quote are required for the index)

```python
data = {"first_name": "Mike", "last_name": "Donnici"}
f"Full name: {data['first_name']} {data['last_name']}"
# 'Full name: Mike Donnici'
```

- f-strings allow for inline operations, including function calls

```python
f"two times two is {2*2}"
# 'two times two is 4'
```

### Templates strings

- simpler syntax
- slower than f-strings
- limited - no format specifiers
- good when working with externally formatted strings
- `$` identifies placeholders

```python
from string import Template
a = "foo"
b = "bar"
s = Template("First $first, then $second")
s.substitute(first=a, second=b)
# 'First foo, then bar'
```

- `{}` can be used where needed

```python
chars = "DEF"
Template("abc${str}ghijk").substitute(str=chars)
# 'abcDEFghijk'
```

- safe substitution always tries to return a usable string, for example when a substitute var is missing

```python
d = {"flavour": "Chocolate", "food": "cake"}
Template("I like $flavour $food").substitute(d)
# 'I like Chocolate cake'
d = {"flavour": "Chocolate"}
Template("I like $flavour $food").substitute(d)
# KeyError: 'food'
Template("I like $flavour $food").safe_substitute(d)
# 'I like Chocolate $food'
```

Of the three string formatting methods, f-strings is generally the best to use, where possible - ie for later versions 
of Python.

## Regular expressions

- A string expression used to match patterns in strings
- Generally precede regex string with `r` to indicate a raw string, eg `r"string"`
- In a regex:
   - Normal characterss match themselves, eg `st` will match "st"
   - Metacharacters represent _types_ of characters, eg:
      - `\d` represents a digit
      - `\s` represents a space
      - `\w` represents a word char (a-Z, 0-9, _)
   - Metacharacters are used to represent _ideas_:
      - `\w{3,10}` - a word character that appears from 3 to 10 times
 
- Pattern matching is used a lot in data science, to:
   - find and replace text
   - validate strings   
- regex is powerful and _fast_       
- use `re` library

```python
import re
re.findall(r"[REGEX]", in_str)                  # retuns a list of all occurrences
re.split(r"[REGEX]", in_str)                    # splits by the pattern
re.sub(r"[REGEX]", replacement_str, target_str) # replace occurrences in the target string
```  
    
### Supported metacharacters

| Metachar | meaning        |
| -------- | -------------- |
| `\d`     | digit          |
| `\D`     | non-digit      |
| `\w`     | word char      |
| `\W`     | non-word       |
| `\s`     | whitespace     |
| `\S`     | non-whitespace |

> `\S` is _anything but a whitespace_ ... handy!

### Quantifiers

- Quantifiers are metacharacters that specify how many times a char should occur to be matched
- Quantifiers applies _only_ to the character immediately to the left

```python
import re
pw = "password1234"
re.search(r"\w\w\w\w\w\w\w\w\d\d\d\d", pw)  # cumbersome!
re.search(r"\w{8}\d{4}", pw)                # nicer
```

- `+`: once or more times, eg r"\d+" - one or more numbers
- `*`: zero or more times, eg r"\W*" - any number of non-word chars
- `?`: zero times OR once, eg r"colou?r" - color OR colour
- `{n, m}`:  Min - Max occurrences, eg:
   - `r"\d{4,6}"` - numbers 4-6 digits long
   - `r"\d{4,}"`  - numbers at least 4 digits long
   - `r"\d{4}"`   - number exactly 4 digits long


### Looking for patterns

- two methods:
   - `re.search()` matches regex anywhere in the string
   - `re.match()`  matches regex from the beginning of the string

### Special characters

- `.`   matches any character except a newline, eg `.+` any char, one or more times
- `^`   anchors regex to the start of the string, eg `r"^Once upon a time"`
- `$`   anchors regex to the end of the string, eg `r"lived happily ever after.$"`
- `\`   escape special characters such as `(`, `)`, `.`, `\`
- `|`   OR operator, eg `r"Lion|liger"`
- `[]`  denotes a set of characters, eg `r"[a-zA-Z]+\d"` one or more alphas followed by a digit
- `[^]` negative set of characters, eg `r"^[^0-9]"` a string that does not start with a digit

### Greedy vs lazy matching

- Standard quantifiers are _greedy_ by default: `*`, `+`, `?`, `{n, m}`
- _Greedy_ will attempt to match as many characters as possible and return the longest match

```python
re.match(r"\d+", "123456abcdef")
# <re.Match object; span=(0, 6), match='123456'>
# ... finds a match at '1' but keeps going... hence 'greedy' 
```

- Greedy matching backtracks when too many characters matched, and gives up characters one at a time
- For example: `re.match(r".*ABC", "xABCxxx")`
   - first match is any char '.', zero of more times, result = `xABCxxx`
   - second match is 'A', so backtracks one char at a time until it finds the 'A', result = `xA`
   - third and fourth matches are 'B' and 'C', result = `xABC`    

```python
re.match(r".*ABC", "xABCxxx")
# <re.Match object; span=(0, 4), match='xABC'>
```

- Lazy or non-greedy matching matches as few characters as needed and returns the shortest match
- To convert greedy quantifiers to lazy, append `?`

```python
re.match(r"\d+?", "123456789")
# <re.Match object; span=(0, 1), match='1'>
```

- Lazy matching also backtracks (or forward tracks) when too few characters are matched, and expands selection one char at a time
- The example below would match just the first `x`, but then need to expand one char at a time to to satisfy the remaining matches 

```python
re.match(r".*?ABC", "xABCxxx")
# <re.Match object; span=(0, 4), match='xABC'>
```

### Capturing groups

- Groups can be expressed in a regex with parentheses, `()`
- expresses the part of the result we want to capture
- captured groups are returned as list of tuples
- the non-grouped parts are included in the search pattern, but _not_ in the result

```python
# without grouping - spaces included in result
re.findall(r"[a-z]+\s\d+", "abc 123 def 456 ghi 789")
# ['abc 123', 'def 456', 'ghi 789']

# with grouping - list of tuples, space excluded
re.findall(r"([a-z]+)\s(\d+)", "abc 123 def 456 ghi 789")
# [('abc', '123'), ('def', '456'), ('ghi', '789')]
```

- Quantifiers can be applied to the entire group to the left
   - r"dog+" - `+` applies to `g`
   - r"(dog)+" - `+` applies to `dog`
- Difference between `(\d+)` and `(\d)+`:
   - `(\d)+` - capture a repeated group
   - `(\d+)` - repeat a capturing group
   
```python
re.findall(r"(\d)+", "abc 763 def 321")
# ['3', '1']
re.findall(r"(\d+)", "abc 763 def 321")
# ['763', '321']
```

### Alternation and non-capturing groups

- alternation used to choose between optional patterns

```python
s = "1 cat, 2 birds, 3 dogs"
re.findall(r"\d+\scat|dog|bird", s)
# ['1 cat', 'bird', 'dog']
# group optional matches
re.findall(r"\d+\s(cat|dog|bird)", s)
# ['cat', 'bird', 'dog']
```

- non-capturing groups are used when matching is required without capture
- use `(?:regex)` for non capturing group
- using non-capturing group for alternation:

```python
# capture number only from date
re.findall(r"(\d{2})(?:st|nd|rd|th)", "21st, 22nd, 23rd, 24th")
# ['21', '22', '23', '24']
```

### Backreferences

- Capture groups can be referenced by numbers
- the entire regex is `0`, and the groups are `1`, `2`, `3` etc
- the groups can be referenced using the `.group()` method
- `.group()` can only be used with `.search()` and `.match()`

```python
s = "Maia born on 09-03-2010"
info = re.search(r"(\d{1,2})-(\d{2})-(\d{4})", s)
info.group(0)
# '09-03-2010'
info.group(3)
# '2010'
```

- can name groups using syntax: `(?P<name>regex)`

```python
s = "Callala Bay, 2540"
info = re.search(r"(?P<city>[a-zA-Z\s]+).*?(?P<postcode>\d{4})", s)
info.group("city")
# 'Callala Bay'
info.group("postcode")
# '2540'
```

### Lookaround

- Confirms a sub-pattern is behind or ahead of a main pattern
- Look-ahead:
   - non-capturing group
   - checks if first part of expression is followed (or not) by the lookahead expression
   - returns only the first part 
   - look-ahead can be positive `(?=regex)` or negative `(?!regex)`
- look-behind:
   - positive `(?<=regex)`, negative `(?<!regex)`  
   


