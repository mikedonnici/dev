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

The above  infers union type `tup: (string|number)[]` so can have a string or 
a number type in either position. Can also add to the length so not _fixed_.

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
console.log(combine(3,4))         // -> 7
console.log(combine("Big","Dog")) // -> BigDog
```
