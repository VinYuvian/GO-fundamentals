package main

/*
- Creating pointers
 - pointer arithmetic is not possible but can be attained using unsafe package
 - Pointers are references to the variable location in memory
 - pointers that are not initialized but defined will be pointing to nil
 var varName *TypeOfTheVarToPointTO = &variableToPointTo
 var a int = 24
 var b *int = &a

- DeReferencing pointers
 - is nothing but retrieving the value that is stored at a particular
	memory location
  var a int = 24
  var b *int = &a
  fmt.Println(*b) //gives the value that is stored at a memory location &a as b = &a
- The new function
- Working with nil
- Types with Internal pointers
 */

import(
	"fmt"
)
type myStruct struct {
	Name string
}
func main(){
	var a int = 24
	var b *int = &a
	fmt.Println(a,b)
	fmt.Println(a,*b)
	a = 20
	fmt.Println(a,*b)
	arr:=[]int{1,2,3}
	a1:=&arr[0]
	a2:=&arr[1]
	fmt.Println(arr,a1,a2)
	var m *myStruct
	m=new(myStruct)
	m.Name = "Vinay"
	fmt.Println(m)
	fmt.Println(m.Name)
}