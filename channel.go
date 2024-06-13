package main

import ("fmt"
		"time")

func main(){
	
	//	unbuffered channel
	// var channel = make(chan int)
	// go process(channel)
	// for i:= range channel{
	// 	fmt.Println(i)
	// }

	//	buffered channel
	var channel = make(chan int,5)	// store 5 ints in the buffered channel
	go process(channel)
	for i:= range channel{
		fmt.Println(i)
		time.Sleep(time.Second *1)
	}
}

//	child process is able to store 5 ints into the channel and then exit, then the main function can process each of those values individually
//	closing channel just means that no more values can be add to it, the values in the channel will still remain and still can be read from.
func process(channel chan int){

	defer close(channel)
	
	for i:=0;i<5;i++{
		channel <- i
	}
	fmt.Println("Child process is exiting...")
	
}