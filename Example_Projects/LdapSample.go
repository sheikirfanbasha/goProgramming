package main

import(
	"fmt"
	"github.com/nmcclain/ldap"
)
func main(){
	ldapServer := "localhost"
	ldapPort := 10389
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapPort))
	// be sure to add error checking!
	defer l.Close()
	user := "uid=admin,ou=system"
	passwd := "secret"
	err = l.Bind(user, passwd)
	if err==nil {
	  // authenticated
		fmt.Println("Success!!!")
	} else {
	  // invalid authentication
		fmt.Println("Failure!!!")
	}
}