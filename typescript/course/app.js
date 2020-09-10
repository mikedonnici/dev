// No need to specify object type exactly, as can be inferred
// when only using normal js types.
// Inference
var person = {
    name: 'Mike',
    age: 49
};
// However, to take advantage of TS additional types, such as a tuple
// need to be more specified about the types that are in the object, so:
// Explicity type 
var person2 = {
    name: "Mike",
    age: 49,
    hobbies: ["bjj", "plants"],
    role: [1, "dad"]
};
console.log(person2);
// Enum types (ts only) - a custom type
// Allow mapping of human-readable values associated with a number
var Role;
(function (Role) {
    Role[Role["ADMIN"] = 0] = "ADMIN";
    Role[Role["READ_ONLY"] = 1] = "READ_ONLY";
    Role[Role["AUTHOR"] = 2] = "AUTHOR";
})(Role || (Role = {})); // assiged 0, 1, 2 behind the scenes, can start at any value, eg ADMIN = 5 ... 6, 7
var person3 = {
    name: "Mike",
    age: 49,
    hobbies: ["bjj", "plants"],
    role: Role.ADMIN
};
if (person3.role === Role.ADMIN) {
    console.log(person3.name + " is an admin");
}
