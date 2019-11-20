package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/mikedonnici/goprotobuf/addressbookpb"
)

func main() {

	p1 := NewPerson("123", "Mike", "michael@mesa.net.au")
	addPhone(p1, NewPhone("0402 400 191", addressbookpb.Person_MOBILE))
	addPhone(p1, NewPhone("02 4446 6835", addressbookpb.Person_HOME))
	addPhone(p1, NewPhone("02 4446 6835", addressbookpb.Person_WORK))

	ab := &addressbookpb.AddressBook{
		People: []*addressbookpb.Person{p1},
	}

	jm := jsonpb.Marshaler{}
	msg, err := jm.MarshalToString(ab)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(msg)
}

func NewPerson(id, name, email string) *addressbookpb.Person {
	return &addressbookpb.Person{
		Name:   name,
		Id:     id,
		Email:  email,
		Phones: nil,
	}
}

func NewPhone(phoneNumber string, phoneType addressbookpb.Person_PhoneType) *addressbookpb.Person_PhoneNumber {
	return &addressbookpb.Person_PhoneNumber{
		Number: phoneNumber,
		Type:   phoneType,
	}
}

func addPhone(person *addressbookpb.Person, phone *addressbookpb.Person_PhoneNumber) {
	person.Phones = append(person.Phones, phone)
}
