package main

import (
	"fmt" 
    "go_projects/basic_operations"
	"go_projects/flow_controls"
)



func main(){
	fmt.Println("hello world")
	
	fmt.Println(basic_operations.Add(3,5,9))
    fmt.Println(basic_operations.Sub(5,13,2))
    fmt.Println(basic_operations.Div(8,3))
	fmt.Println(basic_operations.Mul(3,2))

	flow_controls.Forloop(5)
	flow_controls.Switch("sat")
}




