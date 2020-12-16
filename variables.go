package main

import(
	"fmt"
	"strconv"
)
func main() {
	//var i uint
	var i int
	j := 12.
	i = int(j)
	j = float64(i)
	status := true
	str := strconv.FormatBool(status)
	j,k :=strconv.ParseFloat(str,64)
	fmt.Printf("%v,%T\n",str,str)
	fmt.Printf("%v,%T\n",j,j)
	fmt.Printf("%v,%T\n",k,k)
}
