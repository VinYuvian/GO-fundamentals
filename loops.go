package main

/*
For statement
- simple loops
for initializer;condition;incrementer{ //one way
}

initialize
for ;condition;incrementer{ //2nd way
}

initialize
for condition{ // 3rd way
incrementer
}
nested for loops are possible
- exiting early
 - break statement can be used to exit the loop early based on some condition
 - break exists the inner loop it is defined in
 - in order to guide break to loop out to some desired place we can use labels
 - continue statement can be used to pass some execution
 based on some condition and run other
- looping through collections
  for k,v:= range (array/slice/map/struct){
}
 */
import(
	"fmt"
	"strconv"
)

func main(){
	i:=0 // scoped to the function
	for i<5{
		fmt.Println(i)
		i++
	}
	fmt.Println("i value after first loop "+ strconv.Itoa(i))
	for i,j:=0,2;i<5;i,j=i+1,j+1{ //scoped to the loop
		fmt.Println(i,j)
	}
	//fmt.Println("j value after 2nd loop "+ strconv.Itoa(j)) throws an error
	fmt.Println("from inside for loop with break")
	for{
		fmt.Println(i)

		i++
		if i==10{

			break
		}
	}
	fmt.Println("using labels")
Loop :
	for k:=1; k<5;k++{
		for l:=1;l<5;l++{
			if k==3{
				break Loop
			}
			fmt.Println(k*l)
		}
	}
	fmt.Println("working with collections")
	type m struct{
		Name string
		id int
	}
	m1 := []m{
		{Name: "Vinay", id: 1711,},
		{Name: "Dileep", id: 1711,},
	}

	for k,v:=range m1{
		fmt.Println(k,v.Name)
	}

}