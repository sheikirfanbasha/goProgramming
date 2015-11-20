package main
import "github.com/jarias/stormpath-sdk-go"
import "fmt"

func main(){
	//This would look for env variables first STORMPATH_API_KEY_ID and STORMPATH_API_KEY_SECRET if empty
	//then it would look for os.Getenv("HOME") + "/.config/stormpath/apiKey.properties" for the credentials
	credentials, _ := stormpath.NewDefaultCredentials()

	//Init Whithout cache
	stormpath.Init(credentials, nil)

	//Get the current tenant
	tenant, _ := stormpath.CurrentTenant()

	//Get the tenat applications
	apps, _ := tenant.GetApplications(MakeApplicationCriteria().NameEq("test app"))

	//Get the first application
	app := apps.Items[0]

	//Authenticate a user against the app
	account, _ := app.AuthenticateAccount("username", "password")

	fmt.Println(account)
}
