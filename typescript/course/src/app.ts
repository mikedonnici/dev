class Foo {

    constructor(private a = 1) {}

    get aVal() {
        return this.a
    }

    set aVal(n: number) {
        this.a = n
    }
}

const f = new Foo(6)
console.log(f.aVal)
f.aVal = 10
console.log(f.aVal)
