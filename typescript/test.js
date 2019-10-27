var Student = /** @class */ (function () {
    function Student(firstName, middleNames, lastName) {
        this.firstName = firstName;
        this.middleNames = middleNames;
        this.lastName = lastName;
        this.fullName = firstName + " " + middleNames + " " + lastName;
    }
    return Student;
}());
function greet(person) {
    return "Hello, " + person.firstName + " " + person.lastName;
}
var s1 = new Student("Christie", "Anne", "Wood");
var p1 = { firstName: "Mike", lastName: "Donnici" };
document.body.textContent = greet(s1) + "\n" + greet(p1);
