function dec(superConstructor: Function) {
  console.log("dec() called...");
  return class {
    super();
  };
}

class Foo {
  constructor() {
    console.log("Foo constructor");
  }
}
