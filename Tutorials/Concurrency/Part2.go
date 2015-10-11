/* In concurrent system, we often require communication between the goroutines. In Go language it is achieved through "Channels"

The concept idea of goroutine communication is "Don’t communicate by sharing memory(using mutex or locks). Share memory by communicating through channel".


Below is the example which creates a channel of string and how it is used for communication between main thread and a goroutine
	1e3 : equivalent 10 to power 3 i.e, 1000
	time.Duration : It is a function which converts the integer value into seconds
 */


package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main(){
	c := make(chan string)
	go timedMsg("golang", c)
	for i := 0; i < 5; i++{
		fmt.Println("msg is : ", <-c);
	}

}

func timedMsg(msg string, c chan string){
	for i := 0; ; i++{
		str := fmt.Sprintf("message is : %s %d", msg, i)
		//fmt.Println(str)
		c<- str
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond);
	}
}