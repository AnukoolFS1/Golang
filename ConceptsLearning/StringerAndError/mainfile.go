package main

import "fmt"

type Person struct {
	name string
	age int
}

// type SimplePerson Person

func(person *Person) String() string {
	return fmt.Sprintf("%s, (%d years old)", person.name, person.age)
}

func(person Person) PrintSimple() {
	type SimplePerson Person
	fmt.Println(SimplePerson(person))
}

func main() {
	person1 := &Person{"Anukool", 27}
	person2 := Person{"AnukoolSecond", 28}

	// simplePerson := SimplePerson(*person1)

	fmt.Println(person1) // result -> Anukool, (27 years old)
	fmt.Println(person2) // result -> {Anukool2 28}
	// fmt.Printf("%+v\n",simplePerson)


	// if err:= run(1); err != nil{
	// 	fmt.Println(err)
	// }

	person1.PrintSimple()
}
