package main

/*
- Constants can't be altered
- constants can be shadowed
** Constants should be assignable at the time of compilation
const myConst float64 = math.Sin(1.6) gives an error as the
value of math.Sin(1.6) is determined at the run time
- Typed constants
 - need type of constant at the time pf definition
- Untyped Constants
 - compiler infers the type of the constant
 const a = 42 (compiler infers the type of a )
- Enumerated Constants
const (
a = iota
b
c
)
- Enumeration expressions
 */

import(
	"fmt"
)

/*type month int
const(
	january month = iota + 1
	february
	march
	april
	may
	june
	july
	august
	september
	october
	november
	december
)*/

const(
	_ = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

func main(){
	const myConst int =42
	const myCount = 40
	filesize := 4000000000.
	fmt.Printf("%.2fGB",filesize/GB)

}