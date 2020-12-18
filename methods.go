package main

import(
	"fmt"
)

/*
- Method functions :
 syntax :
func (receiver value) funcName(){
	//code
}
calling the method :
associatedObject.methodName()
g.greet()
- methods works with the copies of the objects the methods are
associated with
- passing pointers is allowed
 */
func main(){
	g := greeter{
		greeting : "hello",
		name : "Dileep",
	}
	g.greet()
	fmt.Println("The new name is :",g.name)
}
type greeter struct{
	greeting string
	name string
}

func (g *greeter) greet(){
	fmt.Println(g.greeting,g.name)
	g.name = ""
}