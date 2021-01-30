# TypeScript

https://www.typescriptlang.org/

TypeScript is a _superset_ of javascript that includes types and other features
to make working with JavaScript better.

TypeScript itself does not run in a JavaScript environment, and it must first
be _transpiled_ into javascript.

TypeScript uses \_static types` which means problems can be picked up during
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
let n: number;
let b: boolean;
let s: string;
```

TS can infer the correct type and when initialising variables the type is
generally not specified, eg:

```ts
let n = 10; // infers type `: number`
let b = false; // infers type `: boolean`
let s = "cat"; // infers type `: string`
```

Constants types are inferred in a similar way, eg:

```ts
const n = 10; // infers type `: 10`
const b = false; // infers type `: false`
const s = "cat"; // infers type `: "cat"`
```

This would also work by inference, but `n` could be initialised with a value of
any type, so it defeats the purpose of static typing:

```ts
let n;
n = 5; // infers type `: number`
n = "5"; // infers type `: string`
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
  return n1 + n2;
}
```

Can specify multiple types with `|`:

```ts
function add(n1: number | string, n2: number | string) {
  return n1 + n2;
}
```

#### Object type

Object types are more explicitly defined in TS.

For example:

```ts
const person = {
  name: "Mike",
  age: 50,
};
```

... creates a _concrete_ object type that has a string field `name` and number
field `age`.

This is equivalent to:

```ts
const person: object = {
  name: "Mike",
  age: 50,
};
```

...or:

```ts
const person: {} = {
  name: "Mike",
  age: 50,
};
```

Following on from the above, can be more specific about the object type:

```ts
const person: {
  name: string;
  age: number;
} = {
  name: "Mike",
  age: 50,
};
```

However, is better practise to allow TS to infer the concrete type, as in the
first example, ie:

```ts
const person = {
  name: "Mike",
  age: 50,
};
```

Object types can also be nested:

```ts
const product = {
  id: "abc1",
  price: 12.99,
  tags: ["great-offer", "hot-and-new"],
  details: {
    title: "Red Carpet",
    description: "A great carpet - almost brand-new!",
  },
};
```

The \_type of this object is:

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
let xs1 = ["a", "b", "c"]; // will infer string[]
let xs2: string[] = ["c", "d", "f"]; // explicit, but same
```

Like JS, TS supports arrays with different types, eg

```ts
let xa1 = ["a", 1, false]; // infers any[]
let xa2: any[] = ["b", 2, true]; // same
```

### TS Additional Types

#### Tuple

Tuples are arrays with fixed length and type(s) and can be used in situation
where the elements of an array need to be strictly controlled.

```ts
let tup = ["admin", 4];
tup.push("another string");
console.log(tup);
// ->  [ 'admin', 4, 'another string' ]
```

The above infers union type `tup: (string|number)[]` so can have a string or a
number type in either position. Can also add to the length so not _fixed_.

To specify a tuple specify type explicitly:

```ts
let tup: [string, number] = ["admin", 4];
console.log(tup[0]); // -> "admin"
console.log(tup[2]); // -> ERR: Tuple type '[string, number]' of length '2' has no element at index '2'.
```

So protects code against trying access an element that does not exist.

Will also prevent the following:

```ts
let tup: [string, number] = ["admin", 4, "what?"];
// -> Type '[string, number, string]' is not assignable to type '[string, number]'.
```

However, it does _not_ prevent an additional element being _pushed_ onto the
array:

```ts
let tup: [string, number] = ["admin", 4];
tup.push("another string"); // -> no problem
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

const red = Colour.Red;
console.log(red);
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

console.log(Colour.Blue);
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
    return a + b;
  } else {
    return a.toString() + b.toString();
  }
}

console.log(combine(3, 4)); // -> 7
console.log(combine("Big", "Dog")); // -> BigDog
```

#### Literal Types

A literal is an exact value of a type, eg:

```ts
function printFooString(s: "Foo") {
  console.log(s);
}

printFooString("Foo"); // -> OK
printFooString("FooBar"); // ->  Argument of type '"FooBar"' is not assignable to parameter of type '"Foo"'
```

Works well with Union Type to ensure type safety with params, eg:

```ts
function printFooString(s: "Foo" | "Bar" | "FooBar") {
  console.log(s);
}

printFooString("Foo"); // -> OK
printFooString("Bar"); // -> OK
printFooString("FooBar"); // ->
printFooString("bingo"); // -> Argument of type '"bingo"' is not assignable to parameter of type '"Foo" | "Bar" | "FooBar"'.
```

#### Type Aliases

Provides a more convenient way to represent custom types, eg:

```ts
type FooBarStr = "Foo" | "Bar" | "FooBar";

function printFooString(s: FooBarStr) {
  console.log(s);
}

printFooString("Foo"); // -> OK
printFooString("Bar"); // -> OK
printFooString("FooBar"); // ->
printFooString("bingo"); // -> Argument of type '"bingo"' is not assignable to parameter of type 'FooBarStr'
```

Or, more complex types, eg:

```ts
type User = {
  name: string;
  age: number;
};
```

#### Function Return Types

Return types can be inferred:

```ts
function add(n1: number, n2: number) {
  return n1 + n2; // -> return type is inferred as : number
}
```

...or declared explicitly:

```ts
function add(n1: number, n2: number): number {
  return n1 + n2; // -> return type is inferred as : number
}
```

Generally, allow the return type to be inferred.

For a function with only a side-effect, the `: void` type is returned:

```ts
function printLine(line: string) {
  // -> : void
  console.log(line);
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

let a: Function = add; // -> OK
let b: Function = 10; // -> Type 'number' is not assignable to type 'Function'
```

Specify the signature:

```ts
function add(n1: number, n2: number) {
  return n1 + n2;
}

type adder = (x: number, y: number) => number;
let a: adder = add; // -> OK
let b: adder = print; // -> Type '() => void' is not assignable to type 'adder'
```

Specify the signature for a callback:

```ts
function addHandler(n1: number, n2: number, callback: (n: number) => void) {
  const sum = n1 + n2;
  callback(sum);
}

function printNum(n: number) {
  console.log(n);
}

addHandler(3, 7, printNum); // printNum satisfies callback signature
```

#### Unknown

There's a subtle difference between `unknown` type and the default `any` type:

```ts
let a: any;
let b: string;
a = 5;
a = "Dog";
b = a; // -> ok
```

```ts
let a: unknown;
let b: string;
a = 5;
a = "Dog";
b = a; // -> Type 'unknown' is not assignable to type 'string'.
```

`unknown` is a bit more restrictive and the underlying type needs to be checked
before assigning to a value with a fixed type. This make it a better choice
that `any` as it still requires some specification:

```ts
let a: unknown;
let b: string;
a = 5;
a = "Dog";
if (typeof a === "string") {
  b = a; // -> Ok
}
```

#### Never

The `never` type implies a type that can never have a value.

The code below can never return anything, including `void` or `undefined`:

```ts
function error(msg: string, code: number) {
  // -> infers : void type
  throw { message: msg, code: code };
}
```

Can (and should) be more explicit when a function can `never` return:

```ts
function infiniteLoop(): never {
  while (true) {}
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
  "exclude": ["foo.ts", "**/*.dev.ts", "node_modules"]
}
```

Conversely, `include` allws you to specify _only_ those files that should be
compiled:

```json
{
  "include": ["foo.ts", "bar.ts"]
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
  compilerOptions: {
    target: "es6",
    // "lib": []  // -> defaults for es6
  },
}
```

If uncommented then available libraries must be specified. Note, the settings
below are the default when the target is set to `es6`:

```json5
{
  compilerOptions: {
    target: "es6",
    lib: [
      // these are the defaults for es6
      "es6",
      "dom",
      "dom-iterable",
      "scripthost",
    ],
  },
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
  compilerOptions: {
    outDir: "./dist",
    rootDir: "./src",
  },
}
```

### noEmitOnError

This does not appear in the generated `tsconfig.json` and is set to `false`
by default.

If set to `true` no files will be emitted if there is a compilation error:

```json
{
  "compilerOptions": {
    "noEmitOnError": true
  }
}
```

### strict options

By default, `"strict": true` is on, and sets all of the strict options:

```json5
{
  compilerOptions: {
    /* Strict Type-Checking Options */
    strict: true,
    /* Enable all strict type-checking options. */
    // "noImplicitAny": true,                 /* Raise error on expressions and declarations with an implied 'any' type. */
    // "strictNullChecks": true,              /* Enable strict null checks. */
    // "strictFunctionTypes": true,           /* Enable strict checking of function types. */
    // "strictBindCallApply": true,           /* Enable strict 'bind', 'call', and 'apply' methods on functions. */
    // "strictPropertyInitialization": true,  /* Enable strict checking of property initialization in classes. */
    // "noImplicitThis": true,                /* Raise error on 'this' expressions with an implied 'any' type. */
    // "alwaysStrict": true,                  /* Parse in strict mode and emit "use strict" for each source file. */
  },
}
```

Can override individual options:

```json5
{
  // all on...
  strict: true,
  // ...except this option
  noImplicitAny: false,
}
```

## Classes and Interfaces

TS extends some of the object features in JS - here are some of the TS-specific
features.

Class _fields_ can be defined in a TS class and define the properties that will
exist when an object of that class is instantiated:

```ts
class Department {
  // fields -> object properties
  name: string;

  constructor(name: string) {
    this.name = name;
  }

  describe() {
    return `${this.name} Department`;
  }
}
```

A class method can have a `this` dummy parameter to indicate to TS that, within
that method, `this` refers to the specified type:

```ts
class Department {
  name: string;

  constructor(name: string) {
    this.name = name;
  }

  describe(this: Department) {
    // <- here
    return `${this.name} Department`;
  }
}
```

This means that if this class method was attached to a different type of object,
the compiler will complain. This adds additional type safety for classes:

```ts
const dep = new Department("Accounting");
console.log(dep.describe()); // -> OK

const dep2 = { describe: dep.describe };
console.log(dep2.describe()); // ->
// The 'this' context of type '{ describe: (this: Department) => string; }'
// is not assignable to method's 'this' of type 'Department'. Property
// 'name' is missing in type '{ describe: (this: Department) => string; }'
// but required in type 'Department'.
```

TS implements `private` and `public` (default) access modifiers for properties
and methods.

These can also be included in the constructor for convenience.

So instead of this:

```ts
class Foo {
  private a: string;
  private b: string;
  c: string; // default is public

  constructor(a: string, b: string, c: string) {
    this.a = a;
    this.b = b;
    this.c = c;
  }
}
```

...can do this:

```ts
class Foo {
  // have to be explicit about public in this case
  constructor(private a: string, private b: string, public c: string) {}
}
```

TS also adds a `readonly` keyword to indicate that a value cannot be changed
after initialisation:

```ts
class Foo {
  constructor(private readonly id: string) {}
}
```

### Inheritence

A subclass `extends` a base class:

```ts
class Foo {
  constructor(public c: string) {}
}

class Bar extends Foo {
  // If no constructor here parent constructor is called
}

const foobar = new Bar("A");
console.log(foobar.c);
```

If the subclass has a constructor, it should call `super()` on base class before
any use of `this`:

```ts
class Foo {
  constructor(public c: string) {}
}

class Bar extends Foo {
  b: string;

  constructor(public a: string, b: string) {
    super(a);
    this.b = b;
  }
}

const foobar = new Bar("A");
console.log(foobar.c);
```

`private` properties are not available in subclasses:

```ts
class Foo {
  constructor(private a: string) {}
}

class Bar extends Foo {
  constructor(a: string) {
    super(a);
    this.a = "foo"; // -> Property 'a' is private and only accessible within class 'Foo'.
  }
}

const foobar = new Bar("oi");
```

`protected` is same as `private` except is available to subclasses:

```ts
class Foo {
  constructor(protected a: string) {}
}

class Bar extends Foo {
  constructor(a: string) {
    super(a);
    this.a = "foo";
  }
}

const foobar = new Bar("oi");
```

Can override methods of a base class, as usual:

```ts
class Foo {
  constructor(protected a: string) {}

  thing() {
    console.log("Foo thing");
  }
}

class Bar extends Foo {
  constructor(a: string) {
    super(a);
    this.a = "foo";
  }

  thing() {
    console.log("Bar thing");
  }
}

const foobar = new Bar("oi");
foobar.thing(); // -> "Bar thing"
```

### Getters and Setters

TS has `get` and `set` keywords for creating methods that behave like
properties, but allow for more logic to be associated with setting and getting
properties:

```ts
class Foo {
  constructor(private a = 1) {}

  get aVal() {
    return this.a;
  }

  set aVal(n: number) {
    this.a = n;
  }
}

const f = new Foo(6);
console.log(f.aVal);
f.aVal = 10;
console.log(f.aVal);
```

### Static Methods and Properties

Static methods and properties can be accessed without an instance of the class.

These are commonly used for global variables accessed via class, or utility
methods grouped within a class.

Important to note that static properties cannot be accessed using `this`
because this refers to an _instance_ of the class. They can, however, be
accessed via the class name:

```ts
class Foo {
  static BooBoo = "Bear";

  constructor(private bb: string = Foo.BooBoo) {
    console.log(bb);
  }

  static doThing() {
    console.log("thing");
  }
}

Foo.doThing();
```

### Abstract Classes

Use the `abstract` keyword to enforce the implementation of a property or method
on subclasses:

```ts
abstract class Animal {
  abstract sound: string;

  abstract talk(): void;
}

class Dog extends Animal {
  sound: string = "woof";

  talk() {
    console.log(this.sound);
  }
}

const milo = new Dog();
milo.talk();
```

### Private constructors and singleton pattern

Might not use but can create a singleton pattern using a private constructor:

```ts
class Life {
  private static instance: Life;
  private static meaning: number;

  private constructor(meaning = 42) {
    Life.meaning = meaning;
  }

  static getInstance() {
    if (Life.instance) {
      return Life.instance;
    }
    return new Life(42);
  }

  getMeaning(): number {
    return Life.meaning;
  }
}

const l = Life.getInstance();
console.log(l.getMeaning());
```

### Interfaces

Interfaces are used to create custom object types that can be type-checked.

This is a pure TS feature so nothing is translated to the compiled JS code.

```ts
interface Dog {
  size: string;
  colour: string;
  volume: string;

  bark(): string; // method signature
}

const milo: Dog = {
  size: "large",
  colour: "brown",
  volume: "loud",
  bark: () => "WOOF!",
};

milo.bark();
```

An interface are very close, but not _exactly_ the sample thing.

Whilst a custom `type` would work ok in the above example, as interfaces can
only be used to describe an object, generally and interface is used in this
case.

Custom types can be more flexible as they can have union types and describe
non-object types.

Interfaces can be used in a similar way to Go, that is, to provide a contract
definition such that classes can _implement_ (one or more) interfaces:

```ts
interface Edible {
  salt: number;
  fat: number;
  carbohydrate: number;
  protein: number;

  healthScore(): number;
}

class Hamburger implements Edible {
  constructor(
    public salt: number,
    public fat: number,
    public carbohydrate: number,
    public protein: number
  ) {}

  healthScore(): number {
    return this.protein + this.carbohydrate - this.fat - this.salt;
  }
}

const cheeseBurger = new Hamburger(10, 15, 20, 2);
const homeBurger = new Hamburger(3, 5, 15, 20);

console.log(cheeseBurger.healthScore());
console.log(homeBurger.healthScore());
```

The `readonly` keyword can be used in an interface definition to ensure that a
property can only be set once.

```ts
interface Runner {
  readonly legs: number;
  topSpeed: number;
}
```

Interfaces can extend multiple interfaces as they end up all merged into one:

```ts
interface Named {
  name: string;
}

interface Speaks {
  languages: string[];
}

interface Greetable extends Named, Speaks {
  greet(): string;
}

interface Runner {
  readonly legs: number;
  speed: number;
}

class Person implements Greetable, Runner {
  legs = 2;

  constructor(
    public name: string,
    public languages: string[],
    public speed: number = 15
  ) {}

  greet(): string {
    return `Hello, ${this.name}.`;
  }
}

const mike = new Person("Mike", ["English", "Italian"]);
console.log(mike.greet());
```

Interfaces can also be used to define a function type, although would generally
use a type:

```ts
// function type
type addFn1 = (a: number, b: number) => number;
let add1: addFn1 = (a: number, b: number) => a + b;

// Interface type
interface addFn2 {
  (a: number, b: number): number;
}

let add2: addFn2 = (a: number, b: number) => a + b;
```

### Optional parameters and properties

Use the `?:` operator to designate properties optional:

```ts
interface Named {
  name?: string;
}

class Person implements Named {
  constructor(age?: number, legs: number = 2) {}
}

const p = new Person();
```

The above example has an _optional_ `name` property via the `Named` interface,
an optional `age` parameter in the `Person` constructor, and an optional `legs`
parameter by setting a default value in the constructor.

So this code compiles without issue.

## Advanced Types

### Intersection types

Combines two or more types:

```ts
type A = {
  a1: string;
};
type B = {
  b1: number;
};

// Intersection type
type AB = A & B;

const foo: AB = {
  a1: "abc",
  b1: 123,
};
```

### Type guards

Used when need to check properties that may contain union types. To do this can
use the `in` keyword to check that a property is `in` an object, and therefore
deduce the required type:

```ts
type A = {
  a: number;
};
type B = {
  b: string;
};
type C = A | B;

const foo: C = {
  a: 1,
};

// This is a type guard expression using 'in'
if ("a" in foo) {
  console.log(`foo must be type A, so foo.a exists and is equal to ${foo.a}  `);
}
// No type guard so TS unsure if property exists
console.log(`${foo.b} will error`); // -> property 'b' does not exist on type 'A'
```

With a class can use `instanceof` type guard:

```ts
if (vehicle instanceof Truck) {
  vehicle.loadCargo();
}
```

## Discriminated unions

A pattern to help with type guarding for union types - add a common, literal
property to types / interfaces and use `switch`:

```ts
type Dog = {
  type: "dog";
  groundSpeed: number;
};

type Bird = {
  type: "bird";
  maxHeight: number;
};

type Animal = Dog | Bird;

function description(animal: Animal): string {
  switch (animal.type) {
    case "bird":
      return `Bird can fly up to ${animal.maxHeight}m high`;
    case "dog":
      return `Dog can run ${animal.groundSpeed}km per hour`;
    default:
      return "Unknown animal";
  }
}

const a: Animal = {
  type: "bird",
  maxHeight: 200,
};
console.log(description(a));
```

### Type casting

Used to tell TS that the type is known - two options:

```ts
// This syntax clashes with JSX
const el1 = <HTMLInputElement>document.getElementById("user-input");
// This syntax does not
const el2 = document.getElementById("user-input") as HTMLInputElement;
```

### Index properties

Used to build objects with properties of a single type, but where the number and
name of the properties is not known in advance:

```ts
interface Errors {
  [code: number]: string;
}

const errs: Errors = {
  200: "OK",
  400: "Bad Request",
  401: "Unauthorised",
  404: "Not found",
};
```

Note: can only use a single property type in these objects.

### Function overloads

Multiple function signatures to explicitly declare return types:

```ts
type Combinable = string | number;

function add(a: string, b: string): string;
function add(a: number, b: number): number;
function add(a: Combinable, b: Combinable) {
  if (typeof a === "string" || typeof b === "string") {
    return a.toString() + b.toString();
  }
  return a + b;
}
```

### Optional chaining

Check nested properties exist before accessing them using `?.`:

```ts
// Standard JS way
if (user.job && user.job.title) {
    console.log(user.job.title)
}
// TS optional chaining
if (user?.job?.title) {
    console.log(user.job.title)
}
```

### Nullish coalescing

The `??` checks if a value is `null` or `undefined` rather than _falsey_:

```ts
const s = ''
const msg1 = s || 'UNDEFINED' // s1 is falsey
console.log(msg1)
// -> UNDEFINED
const msg2 = s ?? 'UNDEFINED'
console.log(msg2)
// -> '' because empty string is NOT null or undefined
```

## Generics

Generic types are not set in stone when a function is created, but are set
_dynamically_ when the function is called:

```ts
function whatTypeIs<T>(t: T) {
  switch (typeof t) {
    case "string":
      return "STRING";
    case "number":
      return "NUMBER";
    case "boolean":
      return "BOOLEAN";
    case "function":
      return "FUNCTION";
    default:
      return "OTHER";
  }
}

console.log(whatTypeIs("a"));
console.log(whatTypeIs(12));
console.log(whatTypeIs(false));
console.log(whatTypeIs(() => {}));
console.log(whatTypeIs({ a: 1 }));
```

In this first example TS has no way of knowing what the properties are on the
objects that get passed to `merge()`:

```ts
function merge(a: Object, b: object) {
  return Object.assign(a, b);
}

const merged = merge({ name: "Mike" }, { age: 50 });
console.log(merged.age); // -> Error
// -> Property 'age' does not exist on type 'Object & object'
```

Using generics:

```ts
function merge<T, U>(a: T, b: U) {
  // return type is T & U
  return Object.assign(a, b);
}

const merged1 = merge({ name: "Mike" }, { age: 50 });
console.log(merged1.age); // -> OK

const merged2 = merge({ biscuitId: 1234 }, { tastinessRating: 17 });
console.log(merged2.tastinessRating); // -> Also OK
```

Although redundant, can specify concrete types being passed:

```ts
function foo<T>(t: T) {
  console.log(typeof t);
}

foo<string>("this is string");
foo<number>(42);
```

### Generic constraints

In the above examples, the generic types are _unconstrained_ - that is, they can
be any type at all. However, it may be practical to ensure that a generic type
is limited to one or more types. This can be done using the `extends` keyword.

For example, in the `merge()` function it would make sense to ensure that the
objects passed in are at least actual objects:

```ts
function merge<T extends object, U extends object>(a: T, b: U) {
  // return type is T & U
  return Object.assign(a, b);
}

const merged1 = merge({ name: "Mike" }, { age: 50 });
const merged2 = merge({ name: "Mike" }, 50); // -> Error:
// -> Argument of type 'number' is not assignable to parameter of type 'object'.
```

Can constrain with union types as well:

```ts
function foo<T extends string | number>(t: T) {
  console.log(typeof t);
}
```

Using `keyof` constraint:

```ts
// This will error
function valueAt(obj: Object, key: string) {
  return obj[key]; // -> Error
}
```

```ts
// This ensures the  key arg exists in the object
function valueAt<T extends object, U extends keyof T>(obj: T, key: U) {
  return obj[key];
}
```

### Generic classes

Example:

```ts
class List<T> {
  private items: T[] = [];

  add(item: T) {
    this.items.push(item);
  }
}

const stringList = new List<string>();
stringList.add("a");
stringList.add(1); // -> Error:
// -  Argument of type 'number' is not assignable to parameter of type 'string'.
```

## Decorators

Note: set `tsconfig.json`:

```json
{
  "compilerOptions": {
    "target": "es6",
    "experimentalDecorators": true
  }
}
```

Enable meta-programming - code that becomes easier to use by other developers.

By convention, use an uppercase character for decorator function name.

A decorator receives a function to be _decorated_, and returns a new, _decorated_
function.

A simple example:

```ts
function DecoratorFunc(f: Function) {
  console.log("This is a decorator...");
}

@DecoratorFunc
class Foo {
  constructor() {
    console.log("Foo constructor");
  }
}

const f = new Foo();
// This is a decorator
// Foo constructor
```

To pass arguments to a decorator can use a decorator factory. Also note that
the decorator is triggered by the class definition and not by instantiation
of an object:

```ts
// This RETURNS a decorator function
function DecoratorFunc(s: string) {
  return (f: Function) => {
    console.log(`Decorator received arg: ${s}`);
  };
}

@DecoratorFunc("hello")
class Foo {
  constructor() {
    console.log("Foo constructor");
  }
}
// Decorator received arg: hello
```

Multiple decorators run from the bottom, up:

```ts
function a(_: Function) {
  console.log("A");
}

function b(_: Function) {
  console.log("B");
}

@a
@b
class F {
  constructor() {
    console.log("Class F");
  }
}
// B
// A
```

Decorators can be added to properties, methods, accessors and parameters. The
arguments received by the decorator varies depending on where the decorator
is used.

See: https://www.typescriptlang.org/docs/handbook/decorators.html#decorators
