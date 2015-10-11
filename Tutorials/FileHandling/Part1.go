/* Lets just read the file and display its content
	For this we use pkg "io/ioutil"

	ReadFile : 
	returns the content of the specified file in a byte array.
	we can easily convert byte array to string using "string(bytearray)"
 */


package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	data, err := ioutil.ReadFile("assets/Sample.txt");
	if err != nil{
		fmt.Println("File is not there " , err)
	}else{
		content := string(data)
		fmt.Println("Content is : " , count, content)
	}
}