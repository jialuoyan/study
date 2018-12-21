package vender

import "fmt"

type Person interface {
    name() string;
    age() int;
}

type Woman struct {
}

func (woman Woman) name() string {
   return "Jin Yawei"
}
func (woman Woman) age() int {
   return 23;
}

type Men struct {
}

func ( men Men) name() string {
   return "liweibin";
}
func ( men Men) age() int {
    return 27;
}

func TestInterface() {
	var person Person;

    person = new(Woman);
    fmt.Println( person.name());
    fmt.Println( person.age());
    person = new(Men);
    fmt.Println( person.name());
    fmt.Println( person.age());
}