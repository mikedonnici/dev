// Union Types allows for multiple types to satisfy a requirement, use the | operator.

// this function only works with numbers
function add(n1: number, n2: number) {
    const result = n1 + n2
    return result
}

// this version works with numbers and strings
function combined(item1: number | string, )
