
// No need to specify object type exactly, as can be inferred
// when only using normal js types.
// Inference
const person = {
    name: 'Mike',
    age: 49,
}

// However, to take advantage of TS additional types, such as a tuple
// need to be more specified about the types that are in the object, so:
// Explicity type 
const person2: {
    name: string;
    age: number;
    hobbies: string[];
    role: [number, string]; // <-- this is a  tuple type
} = {
    name: "Mike",
    age: 49,
    hobbies: ["bjj", "plants"],
    role: [1, "dad"]
}
console.log(person2)

// Enum types (ts only) - a custom type
// Allow mapping of human-readable values associated with a number
enum Role {ADMIN, READ_ONLY, AUTHOR} // assiged 0, 1, 2 behind the scenes, can assign arbitray values, eg ADMIN = 5 
const person3 = {
    name: "Mike",
    age: 49,
    hobbies: ["bjj", "plants"],
    role: Role.ADMIN
}
if (person3.role === Role.ADMIN) {
    console.log(`${person3.name} is an admin`)
}

// Any type is any type :)... to be avoided because it takes away the ts advantages.



