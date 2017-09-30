package main

import(
	"fmt"
)

type Employee struct {
	name string
	age int
}

func main(){
	fmt.Println("Hello Word desde Go :D")

	color := "Blue"
	edad := 22
	nombre := "Camila"

	fmt.Println(nombre, edad, color)
	fmt.Println(multiplicar(edad, 2))

	fmt.Println(adder(1,2,3))
	fmt.Println(adder(9,9))
	nums := []int{10, 20, 30}
	fmt.Println(adder(nums...))

	s := Employee{name: "Camila", age: 22}
	fmt.Println(s.name)

}


func multiplicar(a int, b int) int {
	return a*b
}

func adder(nums ...int) int {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}