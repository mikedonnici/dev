# Protocol Buffers

Defined in a `.proto` test file.

Advantages (over JSON):

- Data is fully typed
- Data is compressed automatically
- Schema is well-defined
- Documentation can be embedded in file
- Can be read across many languages
- Schema can be safely evolved over time
- Much faster that JSON and XML
- Code can be generated automatically

Protocol buffer is used too share data across languages.

Human-readable `.proto` files are used to generate boilerplate code in your language of choice. This code provides objects and methods for serialising data which can, in turn, be used by any language.

**Protocol Buffer is a data exchange format.**

Protocol buffer exchanges messages, which are defined in a `.proto` files, eg:

```proto
syntax = "proto3";

// type name = tag;
message fooMessage {
  int32 id = 1;
  string text = 2;
  bool read = 3;
}
```

## Scalar Types

- `number`: `double` (64 bits), `float` (32 bits), `int32`, `int64` etc.
- `bool`: `True`|`False`
- `string`: Only UTF-8 or 7-bit ASCII encoding
- `bytes`: any byte array, eg small image.

Example:

```proto
syntax = "proto3";

message Person {
  int32 age = 1;
  string first_name = 2;
  string last_name = 3;
  bytes img = 4;
  bool verified = 5;
  float height = 6;
}
```

## Tags

In protocol buffer, field names are not all that important, but _tags_ are.

Smallest allowed tag value is `1`, largest is 2<sup>29</sup> - 1 (536,870,911).

Tag numbers `19000` - `19999` are reserved for special use.

Tag numbers `1` - `15` use 1 byte so use for frequently populated fields.

Tags `16` - `2047` us 2 bytes.

## Repeated fields

Implements the concept of _lists_ or _arrays_ comprised of 0 to many elements.

For example:

```proto
syntax = "proto3";

message Person {
  int32 age = 1;
  string first_name = 2;
  string last_name = 3;
  bytes img = 4;
  bool verified = 5;
  float height = 6;

  repeated string phone_numbers = 7;
}
```

## Comments

Supports single line `//` and multi-line `/_ ... _/ format.

Useful for documenting schemas.

## Default values

- `bool` - false
- `number` - 0
- `string` - empty string ""
- `bytes` - empty byte array
- `enum` - first value
- `repeated` - empty list

## Enum type

- First valeu is the _default_
- Must start with tag `0`

```proto
syntax = "proto3";

message Person {
  int32 age = 1;
  string first_name = 2;
  string last_name = 3;
  bytes img = 4;
  bool verified = 5;
  float height = 6;

  repeated string phone_numbers = 7;

  // define EyeColour enum type
  enum EyeColour {
    UNKNOWN = 0;
    GREEN = 1;
    BROWN = 2;
    BLUE = 3;
  }
  // create a field of that type
  EyeColour eye_colour = 8;
}
```

## Multiple and nested messages

A `.proto` file can contain multiple messages:

```proto
syntax = "proto3";

message Person {
  int32 age = 1;
  string first_name = 2;
  string last_name = 3;
  bytes img = 4;
  bool verified = 5;
  float height = 6;

  repeated string phone_numbers = 7;

  // define EyeColour enum type
  enum EyeColour {
    UNKNOWN = 0;
    GREEN = 1;
    BROWN = 2;
    BLUE = 3;
  }
  // create a field of that type
  EyeColour eye_colour = 8;

  // Date message defines a Date type
  Date birth_date 9;
}

message Date {
  int32 year = 1;
  int32 month = 2;
  int32 day = 3;
}
```

Message types can alsobe nested:

```proto
syntax = "proto3";

message Person {
  int32 age = 1;
  string first_name = 2;
  string last_name = 3;
  bytes img = 4;
  bool verified = 5;
  float height = 6;

  repeated string phone_numbers = 7;

  // define EyeColour enum type
  enum EyeColour {
    UNKNOWN = 0;
    GREEN = 1;
    BROWN = 2;
    BLUE = 3;
  }
  // create a field of that type
  EyeColour eye_colour = 8;

  // Date message defines a Date type
  Date birth_date 9;

  // Nested address message type
  message Address {
    string address_1 = 1;
    string address_2 = 2;
    string postcode = 3;
    string city = 4;
    string state = 5;
    string country = 6;
  }

  // can have multiple addresses
  repeated Address addresses = 10;
}

message Date {
  int32 year = 1;
  int32 month = 2;
  int32 day = 3;
}
```

## Importing types

Types can be imported from other `.proto` files, eg:

```proto
syntax = "proto3";

import "proto/date.proto";

message Person {
  string first_name = 1;
  string last_named = 2;
  Date birthday = 3;
}
```

Note that the import path is a bit tricky and needs to be relative to a proto path
which is specified with `--proto_path`.

## Packages

Package names are used to create an arbitrary namespace, eg:

```proto
syntax = "proto3";

package my.date;

message Date {
  int32 year = 1;
  int32 month = 2;
  int32 day = 3;
}
```

Then, when this is imported:

```proto
syntax = "proto3";

import "proto/date.proto";

package person;

message Person {
  string first_name = 1;
  string last_named = 2;
  // use fulle qualified package name
  my.date.Date birthday = 3;
}
```

- Use to organise messages and avoid naming conflicts
- Helps with correct compilation in different languages

