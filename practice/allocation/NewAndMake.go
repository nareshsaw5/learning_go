package allocation

import (
	"fmt"
	"os"
)

type Person struct {
	name string
	age  int
}

func NewTest() {
	fmt.Println("Examples of how to use New/new in Go")
	naresh := new(Person) // create instance with new function
	naresh.name = "Naresh"
	fmt.Println(naresh)

	john := Person{"John", 35} // with constructor, passing value of the fields in the same order
	fmt.Println(john)

	steve := Person{age: 40, name: "Steve J"} // with field:value pairs, the initializers can appear in any order. This is called composite literals in Go
	fmt.Println(steve)
	a := [...]string{"1", "2"}
	s := []string{}
	m := map[int]string{1: "Naresh", 2: "John"}
	fmt.Println(m)
	fmt.Println(s)
	fmt.Println(a)

}

func NewFile(fd int, name string) *os.File {
	if fd < 0 {
		return nil
	}
	// f := os.File{fd, name, nil, 0}
	f := os.File{}
	return &f // return address of f which is a pointer in Go
}

func NewFileBoilerPlate(df int, name string) *os.File {
	if df < 0 {
		return nil
	}
	f := new(os.File)
	// f.df = df;
	// f.Name(name)
	return f // f is a pointer
}
