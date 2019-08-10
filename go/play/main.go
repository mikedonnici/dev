package main

import "fmt"

var capabilities = []capability{
	{name: "inject"},
	{name: "amputate"},
	{name: "dress"},
	{name: "feed"},
	{name: "sedate"},
	{name: "prescribe"},
}

type staff interface {
	fullName() string
	allowedTo(capability) bool
}

type capability struct {
	name string
}

type person struct {
	name string
}

type nurse struct {
	person
	level int
}

func (n nurse) fullName() string {
	return n.name
}

func (n nurse) allowedTo(c capability) bool {
	switch c.name {
	case "inject":
		return true
	case "amputate":
		return false
	case "feed":
		return true
	case "dress":
		return true
	case "sedate":
		return false
	case "prescribe":
		return false
	}
	return false
}

type doctor struct {
	person
	pager string
}

func (d doctor) fullName() string {
	return d.name
}

func staffType(s staff) string {
	switch s.(type) {
	case nurse:
		return "nurse"
	case doctor:
		return "doctor"
	}
	return "unknown"
}

func (d doctor) allowedTo(c capability) bool {
	switch c.name {
	case "inject":
		return true
	case "amputate":
		return true
	case "feed":
		return false
	case "dress":
		return false
	case "sedate":
		return true
	case "prescribe":
		return true
	}
	return false
}

func staffCan(s staff, c capability) bool {
	return s.allowedTo(c)
}

func showCapabilities(s staff) {
	for _, c := range capabilities {
		fmt.Printf("Can %s do %v - %v\n", s.fullName(), c, staffCan(s, c))
	}
}

func main() {

	n1 := nurse{}
	n1.name = "Jill Cratchet"
	n1.level = 3
	fmt.Println(n1.name, "is a", staffType(n1))
	showCapabilities(n1)

	d1 := doctor{}
	d1.name = "Dr Morse"
	d1.pager = "12345678"
	fmt.Println(d1.name, "is a", staffType(d1))
	showCapabilities(d1)

	p1 := person{
		name: "Mike",
	}
	fmt.Println(p1.name, "is a", staffType(p1))

}
