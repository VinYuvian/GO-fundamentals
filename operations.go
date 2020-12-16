package main

/*
Arithmetic operations : +,-,*,/,%
Bit Operators : &,|,^,&^
BitShifting : >>,<<

 */
import(
	"fmt"
)

func main(){
	a := 10
	b :=6
	c := 8 // 2^3
	fmt.Println(a&b) //1010 & 0110 - 0010 - 2
	fmt.Println(a|b) // 1010 | 0110 - 1110 - 14
	fmt.Println(a^b) // 1010 ^ 0110 - 1100 - 12
	fmt.Println(a&^b) // 1010 &^ 0110 -  - 4
	fmt.Println(c<<2) // 2^3*2^4
	fmt.Println(c>>2) // 2^3/2^4
}