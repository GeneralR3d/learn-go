package main

import ("fmt"
	"golang.org/x/exp/constraints"
)

//	create custom type aliases
type USERID int

//	create structs with custom types

	//custom type union
	type CustomData interface {
		constraints.Ordered | []byte | []rune
	}

	type PERSON [T CustomData]struct{
		ID int
		name string
		data T
	}


type Num interface{
	int| float64 | int16| int32 | int8| float32
}

func main(){
	result := Add_Lib(1.1,2)
	fmt.Printf("Result is %v\n",result)

	james := PERSON[float32]{		//	put in any supported type your data is going to be in
		ID: 22,
		name: "James",
		data: 3.4,
	}

	fmt.Println(james)
}

//	typical function
// func Add(a int,b int) int{
// 	return a+b
// }


//	explictly state all the possible types through a type union
//	works even when the types are a mix
//	works only if we know a fixed set of types, if there are unknown types this wont work
func Add[Type int| float64 | int16| int32 | int8| float32](a Type,b Type) Type{
	return a+b
}

//	using a custom interface to hold all the type unions
func AddFace[T Num](a T, b T) T{
	return a+b
}

func Add_Lib[T constraints.Ordered] (a T, b T) T{
	return a+b
}

// tilda T means it can take in any type that is an alias of int
func Add_Alias[T ~int] (a T, b T) T{
	return a+b
}