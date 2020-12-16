package main

import(
	"fmt"
	"strconv"
)
func main() {
	//var i uint
	var i int
	//defining and assign value at the same time compiler will decide the type
	j := 12.
	//type conversion is not implicit so explicit conversion is required
	i = int(j)
	j = float64(i)
	status := true
	//as strings are aliases for bytes conversion using string(variableOfDiff type)
	// gives a byte value so in order to explicitly convert type we need to use strconv

	str := strconv.FormatBool(status)
	j,_=strconv.ParseFloat(str,2)
	fmt.Printf("%v,%T\n",str,str)
	fmt.Printf("%v,%T\n",j,j)
	//fmt.Printf("%v,%T\n",k,k)
}
