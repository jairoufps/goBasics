package structs

import "fmt"

func GetArray() {
	var arr1 [2] string
	arr2 := [3] int { 1, 2, 3 }
	arr1[0] = "array"
	arr1[1] = "array2"
	fmt.Println(arr1, arr2)
}

func GetSlice() {
	var slice1 [] string
	slice1 = append(slice1, "my", "slice", "1")
	fmt.Println(slice1)
}


type Course struct {
	Name string
	Slug string
	Skills [] string
}


func (p Course) Subscribe (name string) {
	fmt.Printf("The person %s has subscribed to %s course\n", name, p.Name)
}


type Carrer struct {
	Course
}

func (p Carrer) Subscribe (name string) {
	fmt.Printf("The person %s has subscribed to %s carrer\n", name, p.Name)
}
