/*
	Lets create a client which reads its file and searches for the given word and
	returns the no.of occurances of that word.
 */

package main

import(
	"fmt"
	"net"
	"io/ioutil"
	"strings"
	"bufio"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "7272"
    CONN_TYPE = "tcp"
    FILE_PATH_DIR = "assets/Sample"
)


func main(){
	for i := 0; i < 3; i++{
		  doClientWork(i);
	}
}

func doClientWork(i int){
	conn, err := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		// handle error
		fmt.Println("error is : " , err)
	}
	searchKey, _ := bufio.NewReader(conn).ReadString('\n')
	searchKey = searchKey[0:len(searchKey)-1]
	count := SearchFile(i, searchKey);
	conn.Write([]byte(fmt.Sprintf("%v", count) + "\n"))
	fmt.Println("returning count : ", count);
}

func SearchFile(i int, searchKey string) int{
	var count int
	index := fmt.Sprintf("%v", i + 1);
	data, err := ioutil.ReadFile(FILE_PATH_DIR + index + ".txt");
	if err != nil{
		fmt.Println("File is not there " , err)
	}else{
		content := string(data)
		count = strings.Count(content, searchKey)
	}
	return count
}
