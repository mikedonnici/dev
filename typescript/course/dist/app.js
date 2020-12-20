"use strict";
class Foo {
    constructor(a = 1) {
        this.a = a;
    }
    get aVal() {
        return this.a;
    }
    set aVal(n) {
        this.a = n;
    }
}
const f = new Foo(6);
console.log(f.aVal);
f.aVal = 10;
console.log(f.aVal);
