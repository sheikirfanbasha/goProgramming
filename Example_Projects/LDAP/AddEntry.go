
package main

import (
	"fmt"
	"github.com/mqu/openldap"
)

/*

uses:
https://github.com/mqu/openldap
 * 
 * openldap example program :
 * 
 *  - 1 :
 * 
 *    - specify URL for LDAP connexion with user and passwd
 *    - ldap and ldaps is supported,
 *    - anonymous connexion is done with an empty user string
 *    - base (DN) is needed for many LDAP server (it depends on LDAP data design)
 * 
 *  - 2 :
 * 
 *    - you can set some LDAP options.
 *    - authentification with Bind()
 * 
 *  - 3 : setup LDAP query search.
 *  - 4 : print search results.
 * 
 */
func main() {
	var user, passwd, url string

	// (1) - connexion options
	url = "ldap://localhost:10389/"
	// url = "ldaps://some.host:636/"
	user = "uid=admin,ou=system"
	passwd = "secret"
	//base = "ou=people,o=sevenSeas"

	ldap, err := openldap.Initialize(url)

	if err != nil {
		fmt.Printf("LDAP::Initialize() : connexion error\n")
		return
	}

	// (2.1) - options
	ldap.SetOption(openldap.LDAP_OPT_PROTOCOL_VERSION, openldap.LDAP_VERSION3)

	// (2.2) - authentification (Bind)
	err = ldap.Bind(user, passwd)
	if err != nil {
		fmt.Printf("LDAP::Bind() : bind error\n")
		fmt.Println(err)
		return
	}
	defer ldap.Close()


	/* Add Entry to the given LDAP directory */

	var m map[string][]string
    m = make(map[string][]string)

    m["objectclass"] = []string{
								"person", "inetOrgPerson", "organizationalPerson", "top",
							}
	m["cn"] = []string{"pankaj"}
	m["givenname"] = []string{"pankaj dhorea"}
	m["sn"] = []string{"dhorea"}
	m["email"] = []string{"pankaj@gmial.com"}
	m["password"] = []string{"Changepwd1"}
		
 	ldap.Add("cn=pankaj,ou=people,o=sevenSeas", m)

	fmt.Println("Successfully added an entry");

}
