/* In this example we see how to write generator functions.

Generator function : It is like helper function which we can make use of to trigger many goroutines within a function


We modify our "timedMsg" function to take message as input and return the channel as output. We add an ananymous
generator go routine function which actually writes to our channel

*/


package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main(){
	c := timedMsg("golang")
	for i := 0; i < 5; i++{
		fmt.Println("msg is : ", <-c);
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