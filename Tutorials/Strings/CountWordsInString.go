/* This simple program illustrates how to count no.of occurances of a 
given word in the given text*/

package main

import(
	"fmt"
	"strings"
)

func main(){
	fmt.Println(strings.Count("Cheese", "e"));
}