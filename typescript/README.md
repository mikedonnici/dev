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

TS can infer the correct type and when initialising variables the type 
is generally not specified, eg:

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

This would also work by inference, but `n` could be initialised with a value 
of any type, so it defeats the purpose of static typing:

```ts
let n
n = 5   // infers type `: number`
n = "5" // infers type `: string`
```



### Core Types

Types supported by both JS and TS:

- `number` - Any int or float value
- `string` - Text, denoted with ", ' or `
- `boolean` - `true`, `false` 

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



