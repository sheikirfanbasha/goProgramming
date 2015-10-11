/* Using Select statement

It is a smiple switch statement where instead of switch happening according to expression
it happens according to communication channel.

Boundry cases:

If more channels are available same time then the decision is taken pseudo randomly
If no cahnnel is available then default case is returned
If no default case then the select is blocked till one of the channels is available


In this example we modify "Part6.go" to return communication channel in "Infan" function using "Select"
*/


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
	//lets return the channel using select statement
	go func(){
		for{
			select{
				//if "input1" channel is available
				case s := <-input1: c <-s
				//if "input2" channel is available
				case s := <-input2: c <-s
			}
		}
	}()
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