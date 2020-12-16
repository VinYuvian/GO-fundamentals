package main

/*
** Default value of any primitive is 0
boolean type :
- True or False
var n bool = true
m:=1==1 (gives true)
y:=2==1 (gives false)
Numeric Types
- Integers
 - Signed Integers :
	- int8 (-128 - 127 )
	- int16 ( - 2^16 - 2^16-1 )
	- int32
	- int64
 - Unsigned Integers :
	- uint8 (0-255) (also referred as byte)
	- uint16 (0-2^16-1)
	- uint32
- operations can't be performed on different types of integers
 like int32+int64, int32&int64(Needs type convertion )
- available operations : +,-,*,/,%,bitwise operators,bit-shift operators

- Floating Point
 - Decimal points
 - float32
 - float64
- available operations : +,_,*,/
- Complex Numbers
 - complex64
 - complex128
	- complex(1,10)
 - a := 1+1i
	- real(a)
	- imag(a)
- available operations : +,-,*,/,%,bitwise operators,bit-shift operators

Text Types
- strings :
 - an array of characters
 - Strings are immutable
 - strings are an alias for bytes
 - Operations : String concatenation

Rune : UTF-32 character...32 bit long
	*/


import(
	"fmt"
)

func main(){
	var i byte = 128
	var b float64
	var str string = "This is a string"
	by := []byte(str)
	fmt.Printf("%v,%T\n",i,i)
	fmt.Printf("%v,%T\n",b,b)
	c := 1+2i
	fmt.Printf("%v,%T\n",c,c)
	d:=complex(2,24)
	fmt.Printf("%v,%v,%T,%T\n",real(d),imag(d),real(d),imag(d))
	fmt.Printf("%v,%T\n",str,str)
	fmt.Printf("%v,%T\n",str[0],str[0])
	fmt.Printf("%v,%T\n",string(str[0]),str)
	fmt.Printf("%v,%T\n",by,by)


}

