package main

import (
	"fmt"
)
/*
- Basics
 - Interfaces describes behaviour
 - stores method definition
type InterfaceName interface{
	MethodName(inputVariables) outputVariables
}
- any type that can have methods associated with it can
implement interfaces
- Composing Interfaces
- Type conversion
 variableThatIsAssociatedWithCustomType.(TypeThatImplementsInterface)
  - empty interface
  - type switches
- if we are implementing an interface with a value type then the method that
implements should have a value type  as method receiver
- if we are implementing an interface with a pointer type then the method that
implements should have pointer type as method receiver
var i interface{} = 0
switch i.(type){
case int :
case string :
default :
}
- Implementation with values vs pointers
- Best Practices
 - use many,small interfaces instead of large interfaces
   - single method interfaces are more powerful and flexible
- Do export interfaces for types that will be used by packages
 */
type Incrementer interface{
	Increment() int
}
type Printer interface{
	Print([]byte) []byte
}

type ClassPrinter struct{}
type IntCounter int
func main(){
	var c Printer = ClassPrinter{}
	c.Print([]byte("8th class"))
	myInt := IntCounter(0)
	var n Incrementer = &myInt
	n.Increment()
	err,ok :=n.(*IntCounter) //type conversion and checking if it passed
	if ok{
		fmt.Println("n")
	}else{
		fmt.Println("conversion failed",err)
	}
	//type switches
	var i interface{} = 0
	switch i.(type){
	case int :
		fmt.Println("integer")
	case string :
		fmt.Println("string")
	default :
		fmt.Println("Can't describe the type")
	}
}
func (cp ClassPrinter) Print(c []byte) []byte{
	fmt.Println(string(c))
	return c
}

func (inc *IntCounter) Increment() int{ //*IntCounter Implements the Interface
	*inc++
	return int(*inc)
}

