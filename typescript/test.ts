class Student {
  fullName: string;
  constructor(
    public firstName: string,
    public middleNames: string,
    public lastName: string
  ) {
    this.fullName = firstName + " " + middleNames + " " + lastName;
  }
}

interface Person {
  firstName: String;
  lastName: String;
}

function greet(person: Person) {
  return "Hello, " + person.firstName + " " + person.lastName;
}

let s1 = new Student("Christie", "Anne", "Wood");
let p1 = { firstName: "Mike", lastName: "Donnici" };

document.body.textContent = greet(s1) + "\n" + greet(p1);
