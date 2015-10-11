/* Lets create a server which will listen to client
	"net" : this package provides all the methods required for networking
 */

package main

import(
	"fmt"
	"net"
	"container/list"
	"bufio"
	"os"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "7272"
    CONN_TYPE = "tcp"
)


func main(){
	clientList := list.New();
	ln, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		// handle error
	}
	// Close the listener when the application closes.
    defer ln.Close()
    reader := bufio.NewReader(os.Stdin)     
	fmt.Print("Keyword to search: ")     
	searchKey, _ := reader.ReadString('\n') 
	for {

		fmt.Println("Waiting for results....")
		// read in input from stdin     
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		clientList.PushBack(conn)
		//fmt.Println("connection length : " , clientList.Len())
		go handleConnection(conn, searchKey)
	}
}

func handleConnection(conn net.Conn, searchKey string){
	conn.Write([]byte(searchKey + "\n"))
	// will listen for message to process ending in newline (\n) 
	  message, _ := bufio.NewReader(conn).ReadString('\n')
	  fmt.Println("Received count: ", string(message))
}

