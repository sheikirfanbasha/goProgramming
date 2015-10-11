/* we modify "part3" to add two independent "timedMsg" function calls with different messages.

	The output of this code clearly shows that two independent process are synchronized (writes one after the other to channel)
	since they are using the same channel.
 */


package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main(){
	msg1 := timedMsg("message1")
	msg2 := timedMsg("message2")
	for i := 0; i < 5; i++{
		fmt.Println("msg is : ", <- msg1);
		fmt.Println("msg is : ", <- msg2);
	}

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