package main

/*
Maps :
maps are key value pairs
they are like arrays but doesn't use indexes
instead uses keys of one type to reference values of some type

- creation
m := map[keyType]valueType{
	key value pairs separated by commas
}
classStrength:=map[string]int{
    "7th" : 70,
	"8th" : 90,
	"9th" : 100,
}
classStrength := make(map[string]int)

- Manipulation

classStrength["6th"] = 110
- order of key value pairs is guaranteed like arrays
- built in functions
 - delete
		delete(classStrength,"6th")
 - if the key is not present it will return 0 without any error
 so inorder to work well with maps it is better to user ,ok notation before
performing some operations like delete
	strength,ok := classStrength["7th"]
if the key is available ok will be true...if not it will be false
and strength will be the value associated with key 7th
Structs :
struct is collection type which store data of different type
structs are used to store data related to entities
- creation
type Student struct{
	sNo int
	name string
	hobbies []string
}

student := Student{
	sNo : 1,
	name : "Vinay",
	hobbies : []string{
	"cricket","music"
}
}

- manipulation
 struct values will be independent if the structure is copied to
another variable and changed
aDoctor := struct{name string}{name : "vinay"}
bDoctor := aDoctor
b.Doctor.name = "bob"
// aDoctor value won't be changed and will be independent
- naming conventions
- To manipulate the underlying data we need to use addr operator &
 - If the Name of the struct is started with Cap letter
the struct can be seen by other packages but until and unless the
field names start with cap letters other packages will not be able to work
with the structs

- embedding
- Structs can be embedded
- tags
type Student struct{
	sNo int
	name string `required max:"100"`
	hobbies []string
}
- To get the tags we need to user reflect function
tag := reflect.TypeOf(student{})
field,_ := t.FieldByName("name")
field.Tag will give us the tag
 */

import(
	"fmt"
	"reflect"
)
type Student struct{
	sNo int
	name string `name :"string"`
	hobbies []string
}

func main(){
	classStrength := map[string]int{
		"7th" : 100,
		"8th" : 90,
		"9th" : 80,
	}
	fmt.Println(classStrength)
	fmt.Println(classStrength["7th"])
	strength,ok := classStrength["6th"]
	delete(classStrength,"7th")
	fmt.Println(classStrength)
	fmt.Println(strength,ok)

	student := Student{
		sNo : 1,
		name : "Vinay",
		hobbies : []string{"cricket","music"},
	}
	fmt.Println(student)
	fmt.Println(student.sNo)
	fmt.Println(student.hobbies)
	aDoctor := struct{name string}{name : "vinay"}
	bDoctor := aDoctor
	fmt.Println(aDoctor)
	bDoctor.name = "bob"
	fmt.Println(aDoctor)
	fmt.Println(bDoctor)
	cDoctor := &aDoctor
	cDoctor.name = "Harsh"
	fmt.Println(aDoctor)
	fmt.Println(*cDoctor)
	t:=reflect.TypeOf(Student{})
	field,_:=t.FieldByName("name")
	fmt.Println(field.Tag)
}
