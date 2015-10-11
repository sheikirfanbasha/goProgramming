
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
	conn, err := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		// handle error
		fmt.Println("error is : " , err)
	}
	fmt.Println("connection is : ", conn)
}
