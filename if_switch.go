package main

/*
If Statements :
if ture{
	execute some tasks
}

if pop,ok := classStrength["10th"];ok{
	execute some tasks
}
- Operators
 - comparison operators
	- > , < , == , != , >= , <=
 - logical Operators
	- || , && , !
- if-else and if-else if statements
if(condition){

} else if(condition){

} else{

}

Switch Statements :

switch (condition){
case value :
	execution statement
case value :
	execution statement
default :
	execution statement
}
- Simple cases
- Cases with multiple tests
- Falling through
- Type switches
 */

import(
	"fmt"
)

func main(){
	i:= 5
	if 1%2==0{
		fmt.Printf("%v is an even number\n",i)
	}else{
		fmt.Printf("%v is an odd nuber\n",i)
	}
	weekDay:="Saturday"
	switch weekDay{
	case  "Sunday" :
		fmt.Println(weekDay+" is a Weekend")
	case "Saturday" :
		fmt.Println(weekDay+" is a Weekend")
	default:
		fmt.Println(weekDay+" is a Week day")
	}

}
