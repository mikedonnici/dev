function a(_: Function) {
    console.log("A")
}

function b(_: Function) {
    console.log("B")
}

@a
@b
class F {
    constructor() {
        console.log("Class F")
    }
}
// B
// A

