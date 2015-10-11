/* In this example we will just learn how to run two threads (goroutines) using in go language.

	GoRoutine : these are like cheapest form of threads which are  grouped and run in a single thread by the go libraries under the hood.
				Programmer can just use them to trigger independent operations.

	Usage:
		go(keyword) functionName

	Below is the sample implementation where the "Main" thread triggers "timedMsg" as goroutine and waits for two seconds and exits.
	We have made use of "time" library which is provided by go library in order to simulate the wait process.
	
 */

package main

import(
	"fmt"
	"time"
)

func main(){
	go timedMsg("goLangugage")
	fmt.Println("listening...")
	for i := 0; i < 2; i++{
		time.Sleep(time.Second);
	}
}

func timedMsg(msg string){
	for i := 0; i < 5 ; i++{
		str := fmt.Sprintf("message is : %s %d", msg, i)
		fmt.Println(str)
	}
}