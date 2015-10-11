/* In this example we will look how to make the two goroutines run parallely without blocking for another 
	We achieve by using a concept called "Multiplexing" (converting two or more inputs into single output)
	We create an Infan function which takes in two channels as input and genearates a single channel as output.

	The output reveals that each goroutine is running at different paces independently.
*/
package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main(){
	c := Infan(timedMsg("message1"), timedMsg("message2"))
	for i := 0; i < 5; i++{
		fmt.Println("msg is : ", <- c);
	}

}

func Infan(input1 <-chan string, input2 <-chan string) <-chan string{
	c := make(chan string)
	//start a goroutine which writes input1 channel to "c" channel
	go func(){ for{ c <- <-input1}}()
	//start a goroutine which writes input2 channel to "c" channel
	go func(){ for{ c <- <-input2}}()
	//return c channel which has the latest message
	return c
}

func timedMsg(msg string) <- chan string{
	c := make(chan string)
	go func(){
		for i := 0; ; i++{
			str := fmt.Sprintf("message is : %s %d", msg, i)
			c<- str
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond);
		}
	}()
	return c 
}