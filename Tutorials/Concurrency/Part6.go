/* Restoring Sequencing */

package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Message struct{
	str string
	wait chan bool
}

func main(){
	c := Infan(timedMsg("message1"), timedMsg("message2"))
	for i := 0; i < 5; i++{
		msg1 := <-c;
		fmt.Println("msg is : ", msg1.str);
		msg2 := <-c;
		fmt.Println("msg is : ", msg2.str);
		//The following writes are synchronized since the goroutines are blocked on "WaitForIt" channel
		msg1.wait <- true;
		msg2.wait <- true;
	}

}

func Infan(input1 <-chan Message, input2 <-chan Message) <-chan Message{
	c := make(chan Message)
	//start a goroutine which writes input1 channel to "c" channel
	go func(){ for{ c <- <-input1}}()
	//start a goroutine which writes input2 channel to "c" channel
	go func(){ for{ c <- <-input2}}()
	//return c channel which has the latest message
	return c
}

func timedMsg(msg string) <-chan Message{
	WaitForIt := make(chan bool)
	c := make(chan Message)
	go func(){
		for i := 0; ; i++{
			str := fmt.Sprintf("message is : %s %d", msg, i)
			c <- Message{str, WaitForIt}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond);
			//block the goroutine on the WaitForIt channel to simulate synchronization
			<-WaitForIt
		}
	}()
	return c 
}