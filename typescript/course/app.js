function combine(a, b) {
    if (typeof a === "number" && typeof b == "number") {
        return a + b;
    }
    else {
        return a.toString() + b.toString();
    }
}
console.log(combine(3, 4));
console.log(combine("Big", "Dog"));
