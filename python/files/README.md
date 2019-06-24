# Files

## Opening streams

Files are accessed programatically through abstract entities called _streams_ of _handles_.

The operation of connecting a stream to a file is called _opening_ the file, and disconnecting is called _closing_.

Two basic operatuions that can be performed on a stream:

**read** - a portion of the file is retreived and placed into memory

**write** - portions of data are transferred from memory to the file

A stream also has a _mode_:

**read mode** - read operations only

**write mode** - write operations only

**update mode** - both read and write

Using the `open()` function creates an object that is derived from the `IOBase` class, either `BufferedIOBase` (for binary files) or `TextIOBase` (for text files).

Invoking `close()` removes the connection the file and destroys the object.

Python automatically deals with the portability of text files through the _translation of newline characters_.

```python
stream = open(file, mode = 'r', encoding = None)
```

Default for `mode` is _read_, and `encoding` depends on te platform - generally _UTF-8_.

Note that `stdin`, `stdout` and `stderr` are streams that are open and ready by default.

### Stream Modes

The mode that the stream is opened in is signified with a one or two letter code.

A `b` indicates _binday_ mode, a `t` indicates _text_ mode. The default is _text_ mode the the `t` can be ommitted.

Text Mode | Binary Mode | Description
--------- | ----------- | ----------------
rt        | rb          | read
wt        | wb          | write
at        | ab          | append
r+t       | r+b         | read and update
w+t       | w+b         | write and update

For `r` modes the file must already exist, while for `a` and `w` modes the file will be created if it does not already exist.

The _current file position_ is set to before the first byte, except in `a` modes where it is set to after the last byte.

Exclusive creation can be specified with `x` mode. If the file already exists, the open() function will raise an exception.

## Closing streams

The last option that should be performed on a stream is `close()`.

Generally , this is uneventful, but in some circumstances `close()` may raise an `IOError` exception.

## Diagnosing stream problems

The `IOError` object has a property named `errno`:

try:

```
# some stream operations
```

except IOError as exc: print(exc.errno)

The value of `errno` can be compared with one of the predefined symbolic constants defined in the errno module.

A few examples:

- `errno.EACCES` → Permission denied
- `errno.EEXIST` → File exists
- `errno.ENOENT` → No such file or directory
- `errno.EEXIST` → File exists

A more convenient way is to pass the value of `errno` to `os.strerror()` which will return a string describing the error:

```python
from os import strerror
try:
    s = open("/data/file.txt", "rt")
    # processing here...
    s.close()
except Exception as exc:
    print("The file could not be opened:", strerror(exc.errno))
```

## Processing text files

There are various methods associated with a stream that can be used to read in text files, for example: `.read()`, `.readline()`, `.readlines()`.

The `open()` function returns an _object of the iterator class_ so can be used like this:

```python
for line in open('file.txt'):
    print(line)
```

An example using `strerror`:

```python
from os import strerror

try:
    for line in open('file.txt'):
        print(line)
except IOError as e:
    print("I/O error: ", strerr(e.errno))
```

## Binary files

**Amorphous data** has no particular shape or form - it is just a series of bytes.

The `bytearray` class is one that can be used to handle amorphous data.

A `bytearray` is initialised with all elements set to zero, is _**mutable**_ and can only store integer values from 0-255.

```python
data = bytearray(4)
for b in data:
   print(hex(b), int(b), chr(b))

data[0] = 98
data[1] = 110
data[2] = 112
data[3] = 220

for b in data:
   print(hex(b), int(b), chr(b))
```

`.write()` is used to write to binary files.

`.readinto()` is used to read from binary files.

```python
from os import strerror

# create bytearray
numBytes = 3
data = bytearray(numBytes)
for i in range(len(data)):
    data[i] = ord('a') + i

# write it to a file
try:
    bf = open('file.bin', 'wb')
    bf.write(data)
    bf.close()
except IOError as e:
    print("I/O error occurred:", strerr(e.errno))

# initialise bytearray
data = bytearray(numBytes)

# read into it
try:
    fh = open('file.bin', 'rb')
    fh.readinto(data)
    fh.close()
    for b in data:
        print(hex(b), int(b), chr(b))
except IOError as e:
    print("IOError: ", strerror(e.errno))
```

`read()` can also be used to read a binary stream.

In this case it creates a `bytes` object, which is _**immutable**_.

However, it can be easily converted to a `bytearray`.

```python
fh = open('file.bin', 'rb')
bObj = fh.read() # immutable bytes object
bArr = bytearray(bObj) # mutable bytearray object
```

With no args, `read()` specifies the maximum number of bytes to read.

The number of bytes to read can also be specified, as below:

```python
bf = open('file.bin', 'rb')

while data:
    # read 2 bytes at a time
    data = bytearray(bf.read(2))
    for b in data:
        print(int(b), end=' ')
    print()
    
bf.close()
```

Here's a nice example of reading and writing file streams to implement a copy
tool:

```python
from os import strerror

srcname = input("Source file name?: ")
try:
    src = open(srcname, 'rb')
except IOError as e:
    print("Cannot open source file: ", strerror(e.errno))
    exit(e.errno)	
dstname = input("Destination file name?: ")
try:
    dst = open(dstname, 'wb')
except Exception as e:
    print("Cannot create destination file: ", strerr(e.errno))
    src.close()
    exit(e.errno)	

buffer = bytearray(65536)
total  = 0
try:
    readin = src.readinto(buffer)
    while readin > 0:
        written = dst.write(buffer[:readin])
        total += written
        readin = src.readinto(buffer)
except IOError as e:
    print("Cannot create destination file: ", strerr(e.errno))
    exit(e.errno)	
    
print(total,'byte(s) succesfully written')
src.close()
dst.close()
```


