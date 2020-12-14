# TypeScript

https://www.typescriptlang.org/

TypeScript is a _superset_ of javascript that includes types and other features
to make working with JavaScript better.

TypeScript itself does not run in a JavaScript environment, and it must first
be _transpiled_ into javascript.

TypeScript uses _static types` which means problems can be picked up during
development, whereas dynamic types in JavaScript with throw issues at runtime.

To install:

```shell
$ sudo npm install -g typescript
```

TypeScript compiler is then available as `tsc`:

```shell
$ tsc foo.ts
# Compiles foo.ts -> foo.js
```

## Types

Type is designated with `: type` to the right of the var name, eg:

```ts
let n: number
let b: boolean
let s: string
```

TS can infer the correct type and when initialising variables the type is
generally not specified, eg:

```ts
let n = 10     // infers type `: number`
let b = false  // infers type `: boolean`
let s = "cat"  // infers type `: string`
```

Constants types are inferred in a similar way, eg:

```ts
const n = 10     // infers type `: 10`
const b = false  // infers type `: false`
const s = "cat"  // infers type `: "cat"`
```

This would also work by inference, but `n` could be initialised with a value of
any type, so it defeats the purpose of static typing:

```ts
let n
n = 5   // infers type `: number`
n = "5" // infers type `: string`
```

### Core Types

Types that exist in both JS and TS:

- `number` - Any int or float value
- `string` - Text, denoted with ", ' or `
- `boolean` - `true`, `false`
- `object` - with more specific requirements (see below)
- `array` - TS array types can be flexible or strict (see below)

Types can be specified for function parameters, thus:

```ts
function add(n1: number, n2: number) {
    return n1 + n2
}
```

Can specify multiple types with `|`:

```ts
function add(n1: number | string, n2: number | string) {
    return n1 + n2
}
```

#### Object type

Object types are more explicitly defined in TS.

For example:

```ts
const person = {
    name: "Mike",
    age: 50
}
```

... creates a _concrete_ object type that has a string field `name` and number
field `age`.

This is equivalent to:

```ts
const person: object = {
    name: "Mike",
    age: 50
}
```

...or:

```ts
const person: {} = {
    name: "Mike",
    age: 50
}
```

Following on from the above, can be more specific about the object type:

```ts
const person: {
    name: string
    age: number
} = {
    name: "Mike",
    age: 50
}
```

However, is better practise to allow TS to infer the concrete type, as in the
first example, ie:

```ts
const person = {
    name: "Mike",
    age: 50
}
```

Object types can also be nested:

```ts
const product = {
    id: 'abc1',
    price: 12.99,
    tags: ['great-offer', 'hot-and-new'],
    details: {
        title: 'Red Carpet',
        description: 'A great carpet - almost brand-new!'
    }
}
```

The _type of this object is:

```ts
{
    id: string
    price: number
    tags: string[]
    details: {
        title: string
        description: string
    }
}
```

#### Array type

Type of elements in an array can be inferred or specified, eg:

```ts
let xs1 = ["a", "b", "c"]             // will infer string[]
let xs2: string[] = ["c", "d", "f"]   // explicit, but same  
```

Like JS, TS supports arrays with different types, eg

```ts
let xa1 = ["a", 1, false]       // infers any[]
let xa2: any[] = ["b", 2, true] // same
```

### TS Additional Types

#### Tuple

Tuples are arrays with fixed length and type(s) and can be used in situation
where the elements of an array need to be strictly controlled.

```ts
let tup = ["admin", 4]
tup.push("another string")
console.log(tup)
// ->  [ 'admin', 4, 'another string' ]
```

The above infers union type `tup: (string|number)[]` so can have a string or a
number type in either position. Can also add to the length so not _fixed_.

To specify a tuple specify type explicitly:

```ts
let tup: [string, number] = ["admin", 4]
console.log(tup[0]) // -> "admin"
console.log(tup[2]) // -> ERR: Tuple type '[string, number]' of length '2' has no element at index '2'.
```

So protects code against trying access an element that does not exist.

Will also prevent the following:

```ts
let tup: [string, number] = ["admin", 4, "what?"]
// -> Type '[string, number, string]' is not assignable to type '[string, number]'.
```

However, it does _not_ prevent an additional element being _pushed_ onto the
array:

```ts
let tup: [string, number] = ["admin", 4]
tup.push("another string") // -> no problem
```

#### Enum

Human-readable labels that provide a convenient way to handle a _flag_ or
_state_ value without having to remember the strings or numbers that represents
those values.

For example:

```ts
enum Colour {
    Red,
    Green,
    Blue,
}

const red = Colour.Red
console.log(red)
// -> 0
```

Numerical values are assigned (by default) from 0...n, but can override with
number, strings or a combination of both.

```ts
enum Colour {
    Red = "#FF0000",
    Green = "#00FF00",
    Blue = "#0000FF",
}

console.log(Colour.Blue)
// => #0000FF
```

#### Any

Specifies non-specific type which is removed all of the advantages of TS - so
generally avoid as much as possible.

```ts
let unknown: any;
```

#### Union Type

Specifies multiple possibilities for a type, for example:

```ts
function combine(a: number | string, b: number | string) {
    if (typeof a === "number" && typeof b == "number") {
        return a + b
    } else {
        return a.toString() + b.toString()
    }

}

console.log(combine(3, 4))         // -> 7
console.log(combine("Big", "Dog")) // -> BigDog
```

#### Literal Types

A literal is an exact value of a type, eg:

```ts
function printFooString(s: "Foo") {
    console.log(s)
}

printFooString("Foo")    // -> OK
printFooString("FooBar") // ->  Argument of type '"FooBar"' is not assignable to parameter of type '"Foo"'
```

Works well with Union Type to ensure type safety with params, eg:

```ts
function printFooString(s: "Foo" | "Bar" | "FooBar") {
    console.log(s)
}

printFooString("Foo")    // -> OK
printFooString("Bar")    // -> OK
printFooString("FooBar") // -> 
printFooString("bingo")  // -> Argument of type '"bingo"' is not assignable to parameter of type '"Foo" | "Bar" | "FooBar"'.
```

#### Type Aliases

Provides a more convenient way to represent custom types, eg:

```ts
type FooBarStr = "Foo" | "Bar" | "FooBar"

function printFooString(s: FooBarStr) {
    console.log(s)
}

printFooString("Foo")    // -> OK
printFooString("Bar")    // -> OK
printFooString("FooBar") // -> 
printFooString("bingo")  // -> Argument of type '"bingo"' is not assignable to parameter of type 'FooBarStr'
```

Or, more complex types, eg:

```ts
type User = {
    name: string,
    age: number,
}
```

#### Function Return Types

Return types can be inferred:

```ts
function add(n1: number, n2: number) {
    return n1 + n2 // -> return type is inferred as : number
}
```

...or declared explicitly:

```ts
function add(n1: number, n2: number): number {
    return n1 + n2 // -> return type is inferred as : number
}
```

Generally, allow the return type to be inferred.

For a function with only a side-effect, the `: void` type is returned:

```ts
function printLine(line: string) { // -> : void
    console.log(line)
}
```

The type `void` will end up as `undefined` in JS, however TS distinguishes
between `void` and `undefined`.

#### Functions as Types

Specify `Function` as a type:

```ts
function add(n1: number, n2: number) {
    return n1 + n2;
}

let a: Function = add // -> OK
let b: Function = 10  // -> Type 'number' is not assignable to type 'Function'
```

Specify the signature:

```ts
function add(n1: number, n2: number) {
    return n1 + n2;
}

type adder = (x: number, y: number) => number
let a: adder = add    // -> OK
let b: adder = print  // -> Type '() => void' is not assignable to type 'adder'
```

Specify the signature for a callback:

```ts
function addHandler(n1: number, n2: number, callback: (n: number) => void) {
    const sum = n1 + n2
    callback(sum);
}

function printNum(n: number) {
    console.log(n)
}

addHandler(3, 7, printNum) // printNum satisfies callback signature
```

#### Unknown

There's a subtle difference between `unknown` type and the default `any` type:

```ts
let a: any
let b: string
a = 5
a = 'Dog'
b = a // -> ok 
```

```ts
let a: unknown
let b: string
a = 5
a = 'Dog'
b = a // -> Type 'unknown' is not assignable to type 'string'.
```

`unknown` is a bit more restrictive and the underlying type needs to be checked
before assigning to a value with a fixed type. This make it a better choice
that `any` as it still requires some specification:

```ts
let a: unknown
let b: string
a = 5
a = 'Dog'
if (typeof a === 'string') {
    b = a // -> Ok    
}
```

#### Never

The `never` type implies a type that can never have a value.

The code below can never return anything, including `void` or `undefined`:

```ts
function error(msg: string, code: number) { // -> infers : void type  
    throw {message: msg, code: code}
}
```

Can (and should) be more explicit when a function can `never` return:

```ts
function infiniteLoop(): never {
    while (true) {
    }
}
```

## Compiler

### Watch Mode

Run the compiler each time a file changes:

```shell
$ tsc --watch file.ts
# or
$ tsc -w file.ts
```

### Multiple files

Initialise a project:

```shell
$ tsc --init
# -> tsconfig.json
```

Compile all `.ts` files:

```shell
$ tsc
```

Watch all files, recompile on any change:

```shell
$ tsc --watch
```

Excluding files in `tsconfig.json` - note that `node_modules` is excluded by
default:

```json
{
  "exclude": [
    "foo.ts",
    "**/*.dev.ts",
    "node_modules"
  ]
}
```

Conversely, `include` allws you to specify _only_ those files that should be
compiled:

```json
{
  "include": [
    "foo.ts",
    "bar.ts"
  ]
}
```

### Compilation Target

The target JS version (ie `es5` or `es6`) for the compiled project is set in
`tsconfig.json`:

```json
{
  "compilerOptions": {
    "target": "es5"
  }
}
```

### Library Options

The `lib` section in `tsconfig.json` allows you to specify certain libraries
that TS can assume are available, and hence it will not throw compilation errors
in relation to.

The default libraries will depend on the target JS version. For example, if set
to `es6` then `es6` features such as `Map()` will be available to TS.

By default, all DOM APIs are available as well:

```json5
{
  "compilerOptions": {
    "target": "es6",
    // "lib": []  // -> defaults for es6
  }
}
```

If uncommented then available libraries must be specified. Note, the settings
below are the default when the target is set to `es6`:

```json5
{
  "compilerOptions": {
    "target": "es6",
    "lib": [
      // these are the defaults for es6
      "es6",
      "dom",
      "dom-iterable",
      "scripthost"
    ]
  }
}
```

### Sourcemap

If set to `true` the `.js.map` files are generated by the compiler:

```json
{
  "compilerOptions": {
    "sourcemap": true
  }
}
```
This allows the developer tools in the browser to link to the original TS input 
files - breakpoints can be set so this is good for debugging. 

### outdir and rootdir

Specifies the target for compiled code, eg:

```json5
{
  "compilerOptions": {
    "outDir": "./dist", 
    "rootDir": "./src",
  }
}
```

### noEmitOnError











