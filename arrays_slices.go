package main

import (
	"fmt"
)

/*
Arrays :
- Elements of array are contiguous (occupies neighbouring memory blocks)
- due to the above feature accessing array elements is faster
- Creation
 - name := [size]type{elements separated by comma}
   weekDays := [7]string{"Monday","Tuesday","wednesday",Thursday","Friday",Saturday","Sunday"}
   weekDays := [...]string{"Monday","Tuesday","wednesday",Thursday","Friday",Saturday","Sunday"}
   var students [3]string //uninitialized array
   students[0] = "bob"
- copies of array doesn't reference to the underlying data of the array
 instead the copies are literal copies
a:=[3]int{1,2,3}
b:=a
b[1]=5
fmt.Println(a) //{1,2,3}
fmt.Println(b) //{1,5,3}
 but
a:=[3]int{1,2,3}
b:=&a
b[1]=5
fmt.Println(a) {1,5,3}
fmt.Println(b) {1,5,3}
- In order to reference the underlying data we need to use addr operator &

- Built-in Functions
 - len(array)
- Working with Arrays

Slices :
- Creation
 name := []int{elements separated by commas}
 a := []int{1,2,3,4,5}
 b:=a[:]
 c :=a[3:]
- can also use make function to create slices
a := make([]int,3)
- copies of slices unlike arrays is not just a literal copy
instead it points to the underlying data...
a:=[]int{1,2,3}
b:=1
b[1]=5
fmt.Println(a) //{1,2,3}
fmt.Println(b) //{1,2,3}
- Built-in Functions
 - len( slice)
 - cap( slice ) // gives the length of underlying array
 - append(source_slice,element)
 - append(source)slice,elements separated by commas) //it is a variadic function so can take
   multiple elements as an input
 a = append(a,1,2,3)
 - slice concatenation uses spread operator
 a := append(a[:1],a[2:]...)
- Working with slices
 */
func main(){
	a:=[3]int{1,2,3}
	b:=a
	b[1]=5
	fmt.Println(a) // {1,2,3}
	fmt.Println(b) // {1,5,3}
	c:=[3]int{1,2,3}
	d:=&c
	d[1]=5
	fmt.Println(*d) // {1,2,3}
	fmt.Println(c)
	e:=[]int{1,2,3}
	f:=e
	f[1]=5
	fmt.Println(e) //{1,2,3}
	fmt.Println(f) //{1,2,3
}