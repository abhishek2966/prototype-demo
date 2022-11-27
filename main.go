package main

import (
	"fmt"

	"github.com/abhishek2966/prototype-demo/pkg/resource"
)

func main() {
	e1 := resource.NewEmployee().
		Name("John").
		Email("john@email.com").
		Skills([]string{"Go", "C", "Python"})

	e2 := e1.CloneEmployeeWithNewName("NewJohn")

	e1 = e1.Skills([]string{"Go", "Java"})

	fmt.Printf("%+v\n", e1) // {name:John email:john@email.com skills:[Go Java]}
	fmt.Printf("%+v\n", e2) // {name:NewJohn email:john@email.com skills:[Go C Python]}
}
