package main

/*
- Basic syntax
 func nameOfFunction(parameters) returnType{
	executable blocks
}
- Parameters
 - parameters allows the ability to pass values to the functions
 - passing a pointer as a parameter improves performance
 - variadic parameters allows us to pass multiple values with single parameter defined
 - variadic parameter should be the last one to be passed in function call and also
the last one to be defined in function definition
func sum(values ...int){ //takes the incoming values and create a slice of values
}
- Return values
 - both value and pointers are able to be returned using go functions
- Anonymous functions
 - functions which doesn't have name are anonymous functions
func(){
	//code
}() //()after the curly bracket is mandatory for anonymous function
function declaration for a anonymous function is :
var divide func(int,int) int
divide = func(a,b int) int{
//code
}
d := devide(3,5)
- anonymous function can't be called globally

- Functions as types
 - functions can be passed as parameters
 */

import(
	"fmt"
)

func main(){
	a:=[]string{"Vinay","Dileep"}
	for _,v:=range a{
		sayHi("Hi",v)
	}
	sum("the sum is ",1,2,3,4,5)
	m:=multiply(1,2,3,4,5)
	fmt.Println("The multipy value is ", *m)
	d,err:=divide(5.0,2.0)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("the division result is",d)
}

func divide(a,b float64) (float64,error){
	if b == 0.0 {
		return 0.0,fmt.Errorf("0 can't be passed as divider value")
	}
	return a/b,nil
}

func multiply(values ...int) *int{
	result:=1
	for _,v:= range values {
		result *= v
	}
	return &result
}
func sayHi(greeting,name string){
	fmt.Println(greeting+" "+name)
}

func sum(msg string,values ...int){
	fmt.Println(values)
	result:=0
	for _,v := range values{
		result +=v
	}
	fmt.Println(msg,result)
}