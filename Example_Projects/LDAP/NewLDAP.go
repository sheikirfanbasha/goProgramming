
package main

import (
	"fmt"
	"github.com/mqu/openldap"
	"strings"
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

	var user, passwd, url, base string

	// (1) - connexion options
	url = "ldap://localhost:10389/"
	// url = "ldaps://some.host:636/"
	user = "uid=admin,ou=system"
	passwd = "secret"
	base = "ou=people,o=sevenSeas"

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

	// (3) : search method
	// -------------------------------------- Ldap::SearchAll() --------------------------------------
	scope := openldap.LDAP_SCOPE_SUBTREE // LDAP_SCOPE_BASE, LDAP_SCOPE_ONELEVEL, LDAP_SCOPE_SUBTREE  
	filter := "cn=*Wi*"//"cn=*admin*"
	attributes := []string{"cn", "sn", "givenname", "mail"} // leave empty for all attributes

	// SearchAll(base string, scope int, filter string, attributes []string) (*LdapSearchResult, error)
	result, err := ldap.SearchAll(base, scope, filter, attributes)

	if err != nil {
		fmt.Println(err)
		return
	}

	// (4) - print LdapSearchResult(s)
	fmt.Printf("# num results : %d\n", result.Count())
	fmt.Printf("# search : %s\n", result.Filter())
	fmt.Printf("# base : %s\n", result.Base())
	fmt.Printf("# attributes : [%s]\n", strings.Join(result.Attributes(), ", "))

	for _, entry := range result.Entries() {
		fmt.Printf("dn=%s\n", entry.Dn())
		for _, attr := range entry.Attributes() {
			fmt.Printf("%s=[%s]\n", attr.Name(), strings.Join(attr.Values(), ", "))
		}
		
		fmt.Printf("\n")
	}

	var m map[string][]string
    m = make(map[string][]string)

   m["objectclass"] = []string{
								"person", "inetOrgPerson", "organizationalPerson", "top",
							}
	m["cn"] = []string{"irfan"}
	m["givenname"] = []string{"irfan basha"}
	m["sn"] = []string{"sheik"}
	m["mail"] = []string{"sheik@gamil.com"}
	m["userPassword"] = []string{"Changepwd1"}
	m["uid"] = []string{"irfans"}
		
 	t := ldap.Add("cn=irfan,ou=people,o=sevenSeas", m)
 	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t);

}
