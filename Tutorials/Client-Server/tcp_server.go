/* Lets create a server which will listen to client
	"net" : this package provides all the methods required for networking
 */

package main

import(
	"fmt"
	"net"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "7272"
    CONN_TYPE = "tcp"
)


func main(){
	ln, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		// handle error
	}
	// Close the listener when the application closes.
    defer ln.Close()
	for {
		fmt.Println("listening....")
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		fmt.Println("connection is : " , conn)
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn){
	fmt.Println("Received connection");
	err := c.Close()
	if(err != nil){
		fmt.Println("Close conn error : ", err);
	}else{
		fmt.Println("Closed conn successfully");
	}

}
