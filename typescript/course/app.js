// Union Types allows for multiple types to satisfy a requirement, use the | operator.
// this function only works with numbers
function add(n1, n2) {
    return n1 + n2;
}
// this version works with numbers and strings
// function combined(item1: number | string, )
console.log(add(1, '2'));
var user;
user = "Mike";
user = false;
