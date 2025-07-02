package main

import (
	"errors"
	"fmt"
	"learn-go/integer"
	//  host/userORorganization/project/(dir)/package_name
)

//	define structs

type Person struct {
	name   string
	age    int16
	height float32
	shoe   Shoe
	watch  Watch
}

type Shoe struct {
	name  string
	price int32
}

type Watch struct {
	name  string
	price int32
}

// methods are just like any other function except the stuff in brackets before function name ties the method to that specific type
func (shoe Shoe) getPrice() int32 {
	return shoe.price
}

func (watch Watch) getPrice() int32 {
	return watch.price
}

// naked return
func (watch Watch) getPrice2() (price int32) {
	price = watch.price
	return
}

//define interfaces

type Belonging interface {
	getPrice() int32
}

func main() {
	fmt.Println("Hello World")
	const numA int32 = 43
	const numB int32 = 1
	fmt.Println("My number is", numA)
	var myString string = "this is a string" + "this is another string"
	fmt.Println(myString)
	result1, result2, err := adder(numA, numB)
	if err != nil {
		fmt.Println("error!" + err.Error())
	} else {
		fmt.Println(result1, result2)
	}

	var stop int16 = 13
	var array [3]int32 = [3]int32{1, 2, 3}
	// array:=[...]int32{4,5,6}    //  valid alternative
	//var array [4]int32 = [...]int32{7,8,9}   //  size mismatch, [...] only allows for compiler to decide the size
	// var array2 [...]int32 = [...]int32{10,11,12}   //  this is invalid
	fmt.Println(array)
	fmt.Println(array[0])
	fmt.Println(array[1])
	fmt.Println(array[2])

	//  slice conversion
	sliceConvert := array[:]
	fmt.Println("Converted slice is: ")
	fmt.Println(sliceConvert)

	var slice []int32
	fmt.Printf("The length is %d and the capacity is %d\n", len(slice), cap(slice))
	slice = append(slice, array[0], array[1])
	fmt.Println(slice)
	fmt.Printf("The length is %d and the capacity is %d\n", len(slice), cap(slice))

	fmt.Println("second slice...")
	var slice2 []int32 = make([]int32, 3, stop)
	fmt.Printf("The length is %d and the capacity is %d\n", len(slice2), cap(slice2))
	fmt.Println(slice2)

	//make 2 long slices
	fmt.Println("Making 2 long slices")
	longSlice1 := []int32{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(longSlice1)

	var longSlice2 []int32 = make([]int32, 8, 8)
	fmt.Println(longSlice2)
	var i int32 = 0
	for ; i < int32(len(longSlice2)); i++ {
		longSlice2[i] = i + 2
	}
	fmt.Println(longSlice2)

	fmt.Println("Multiply slice2 by 3...")
	fmt.Println(multSlice2(longSlice2, 3))
	fmt.Println("orginal longslice2:")
	fmt.Println(longSlice2)

	fmt.Println("Combine slices1 and slice2 together")
	longSlice1 = append(longSlice1, longSlice2...)
	fmt.Println(longSlice1)

	//  hashmaps
	fmt.Println("map")
	var hashmap map[string]int32 = make(map[string]int32)
	hashmap["adam"] = 22
	hashmap["jason"] = 30
	hashmap["sally"] = 90
	fmt.Println(hashmap["adam"])
	fmt.Println(hashmap["jason"])
	//map will always return something even if the key does not exist. be careful!
	var age, ok = hashmap["adam"]
	if ok {
		fmt.Println("The age of adam is", age)
	} else {
		fmt.Println("Invalid name")
	}
	delete(hashmap, "sally")

	//  print out key value pairs
	fmt.Println("Displaying key value pairs...")
	for key, value := range hashmap {
		fmt.Printf("Key is %s, value is %d\n", key, value)
	}

	//While loops, pythonic way
	j := 0
	for {
		if j > 5 {
			break
		}
		fmt.Println(j)
		j++
	}

	var james Person = Person{name: "James", age: 20, height: 1.72, shoe: Shoe{name: "Air Force One", price: 2000000}}
	fmt.Println(james)
	fmt.Println(james.shoe.name)

	var sally Person
	sally.name = "Sally"
	sally.age = 19
	sally.height = 1.60
	fmt.Println(sally)

	//	anonymous struct

	var rolex = struct {
		name  string
		price int32
	}{name: "submariner", price: 171890}
	fmt.Println(rolex)

	fmt.Printf("Sally's shoe costs %d.\n", sally.shoe.getPrice())
	fmt.Printf("James's shoe costs %d.", james.shoe.getPrice())

	isExpensive(sally.shoe, 300)
	isExpensive(james.shoe, 300)

	//  using variables from another package "integers"
	fmt.Printf("The value of three is %d\n", integer.Three)

}

func adder(numA int32, numB int32) (int32, int32, error) {
	var err error
	if numB < 0 || numA < 0 {
		err = errors.New("Positive numbers only!")
	}
	return numA + numB, 0, err
}

func multSlice(slice []int32, multiplier int32) []int32 {
	for i := 0; i < len(slice); i++ {
		slice[i] *= multiplier
	}
	return slice
}

func multSlice2(slice []int32, multiplier int32) []int32 {
	//  using the range way of iterating through a slice
	for index := range slice { //  Golang understands _ as an unwanted variable and compiler will NOT generate error for unused variable
		slice[index] *= multiplier
	}

	// for index, value := range slice{    // this method does not actually work
	//   value*=multiplier
	// }

	return slice
}

func isExpensive(item Belonging, limit int32) bool {
	if Belonging.getPrice(item) > limit {
		fmt.Println("Expensive!")
		return true
	} else {
		fmt.Println("Not expensive")
		return false
	}
}
