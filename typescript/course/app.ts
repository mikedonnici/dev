// Union Types allows for multiple types to satisfy a requirement, use the | operator.

// this function only works with numbers
function add(n1: number, n2: number | string) {
    return n1 + n2
}

// this version works with numbers and strings
// function combined(item1: number | string, )

console.log(add(1,'2'))

let user: string
user = "Mike"
user = false

