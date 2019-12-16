# Protocol Buffers

> _"The greatest limitation in writing software is our ability to understand the systems we are creating . . . Simpler designs allow us to build larger and more powerful systems before complexity becomes overwhelming."_
>
> -A Philosophy of Software Design, John Ousterhout

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

## Generating code with `protoc`

The `protoc` command-line utility can generate boilerplate code for a
bunch of languages. Do `protoc` for a list of options.

In short, you specify the path to the proto files (`-I`), the output
language and path, and the proto files to parse.

For example, from within a `project` folder containing a `proto` dir and
a `python` dir:

```bash
$ protoc -I=proto --python_out=python proto/*.proto
```

Note: it can be a bit curly to get paths to work, particularly when
there are imports and packages in use.

## Options for generating code

There are `.proto` file options for changing the way the boilerplate
code is generated.

For example, the default package names when generating `Go` code tend to
be not very _idiomatic_. So, can do:

```proto
syntax = "proto3";

option go_package = "foopb";
```

## Schema changes

Over time the schema will need to change, particularly new fields will need to be added and obsolete ones removed.

Protocol buffer is good at managing schema change provided a few rules are adhered to.

### Adding fields

When a new field is added, old code will ignore that field, and if old code is reading from a new schema, the default values for missing fields will be used.

For example:

_v1_

```proto
syntax = "proto3";

message Foo {
  int32 id = 1;
}
```

_v2_

```proto
syntax = "proto3";

message Foo {
  int32 id = 1;
  bool active = 2;
  string name = 3;
}
```

- If _old_ code reads a `v2` message it will ignore `active` and `name`.
- If _new_ code reads a `v1` message it will set `active = false` and `name = ""`.

Obviously the second case means that **nil values need to be handled carefully**.

### Removing fields

Fields can be removed with a similar, but opposite, effect as adding fields.

That is, new code reading old messages will ignore removed fields and old code reading new messages will set fields to the type's nil value.

Again, **nil values need to be handled carefully**.

However, need to ensure that the tag and field name cannot be used again, to prevent conflicts.

This is done with the `reserved` keyword:

_v3_

```proto
syntax = "proto3";

message Foo {
  reserved 2;         // prevent tag being re-used
  reserved "active";  // prevent field name being re-used
  int32 id = 1;
  string name = 3;
}
```

Note: Can specify multiple tags and tag ranges as well as multiple field names with the `reserved` keyword, but cannot mix tags and field names in the same statement:

```proto
message Foo {
  reserved 1, 3, 5 to 9;
  reserved "field1", "field2", "field3";
}
```

Reserving tags prevents conflicts and reserving field names helps to prevent code bugs.

Another option is to rename the field with `OBSOLETE_`.

**NEVER REMOVE `reserved` STATEMENTS**

### Changing field names

Field names can be changed freely as it is the `tag` that is the key piece of information. This is because the field name is not serialised into the binary message, onbut the tag name is.

**This is why the tag must not be changed or re-used.**

### Enum values

Enumerations can also evolve, they can be added, removed and reserved.

## Default values

Care must be taken when dealing with default values for message fields as there is no way to determine if the field value is _missing_ is was explicitly set to the zero value.

Enumerations that receive a non-existing enum value will also be set to the default value.

Therefore, wherever practical, it makes sense to name the first (default) enum value to some variation of `UNKNOWN`.

## Advanced types

### Integers

There are various integer types, `int32`, `int64`, `uint32`, `uint64` etc.

The appropriate selection will depend on the likely range of values that will be stored, and the desire for efficiency.

For example, a signed integer type `sint32` is more efficient for storing negative numbers that an `int32`. So if the data is likely to contain a lot of negative numbers, this might be a good choice.

For more info see the [docs](https://developers.google.com/protocol-buffers/docs/proto3#scalar)

### Oneof

Defines a set of values for which only `oneof` the values needs to be set.
The idea is to save memory by not serialising a bunch of empty, unrequired fields.

eg:

```proto
message Foo {
  int32 id = 1;
  oneof {
    int32 bar = 2;
    int32 baz = 3;
  }
}
```

If more multiple `oneof` field values are set in code, only the last one to be set will be written. That is, each new `oneof` field value replaces the previous.

- Cannot do `repeated oneof`
- Can be difficult to evolve

### Maps

Used to map scalar values (except `float` and `double`) to values of any type.

Maps cannot be `repeated` and the order cannot be guaranteed as it is an in-memory key-value store.

eg:

```proto
message Foo {
  map<int, float> tempAfterSeconds= 1;
}
```

### Timestamp [⇥](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#timestamp)

A `timestamp` is on of the [well-known types](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf) in protocol buffers. That is, a type that is common across many languages.

In protobuf, a `timestamp` is seconds plus nanoseconds from UTC epoch.

The `timestamp` type is imported thus:

```proto
syntax = "proto3";

import "google/protobuf/timestamp.proto";

message Foo {
  google.protobuf.Timstamp fooTime = 1;
}
```

### Duration [⇥](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#duration)

A `duration` is also a well-known type and has `seconds` and `nanoseconds` fields.

## Naming conventions

- Message names in `CamelCase`
- Field name in `snake_case`
- Enum values in `CAPITAL_SNAKE_CASE`

See the [style guide](https://developers.google.com/protocol-buffers/docs/style) for more.

There is also a good [proto style guide](https://github.com/uber/prototool/blob/dev/etc/style/uber1/uber1.proto) from Uber.

## Services

Protocol buffers can be used to define services, or endpoints, that are used to send and receive messages.

Service code can be generated for various languages - [gRPC](https://grpc.io) is one such implelemtation.

Clients appear to be calling the server functions directly (ie, RPC) and all of the transport is handled by gRPC.

The `.proto` file defines the service, and the messages that are involved:

```proto
syntax = "proto3";

import "google/protobuf/timestamp.proto";
import "google/protobuf/duration.proto";

message SearchRequest {
  timestamp search_time = 1;
  string search_phrase = 2;
}

message SearchResponse {
  string search_id = 1;
  duration search_duration = 2;
  repeated string results = 3;
}

service SearchService {
  rpc Search(SearchRequest) returns (SearchResonse);
}
```

## Refs

- <https://developers.google.com/protocol-buffers> - Google developer docs
- <https://github.com/gogo/protobuf> - alternate protobuf implementation for Go
- <https://jbrandhorst.com/post/gogoproto/> - article on above
