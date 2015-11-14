/* 
Step1:

	Install 
	$ go get github.com/ziutek/mymysql/thrsafe
	$ go get github.com/ziutek/mymysql/autorc
	$ go get github.com/ziutek/mymysql/godrv

Step2:
	Install MYSQL.
		I have used Ammps which comes with mySQL. Run localhost/phpMyAdmin.
		Create a database named "quantiply_tests"
		Create a table named "SampleTable" with columns "name", "emailId"

Step3: Access them in go as follows

*/
package main
import (
	"fmt"
    "os"
    "github.com/ziutek/mymysql/mysql"
    _ "github.com/ziutek/mymysql/native" // Native engine
    // _ "github.com/ziutek/mymysql/thrsafe" // Thread safe engine
)

func main() {
    db := mysql.New("tcp", "", "127.0.0.1:3306", "root", "root", "quantiply_tests")

    err := db.Connect()
    if err != nil {
        panic(err)
    }
    //insert new values into DB
    stmt, err := db.Prepare("insert into SampleTable values (?, ?)")
	//checkError(err)
	 _, err = stmt.Run("koyal", "magic@email.com")
    //checkError(err)
    // res,err = db.Exec("insert into SampleTable values('bifrost', 'magic@email.com')");
    if err != nil {
        panic(err)
    }

    rows, res, err := db.Query("select * from SampleTable")
    if err != nil {
        panic(err)
    }

    for _, row := range rows {
        for _, col := range row {
            if col == nil {
                // col has NULL value
            } else {
                // Do something with text in col (type []byte)
            }
        }
        // You can get specific value from a row
        val1 := row[1].([]byte)

        val0 := row[0].([]byte)

         os.Stdout.Write(val0)

        // You can use it directly if conversion isn't needed
        os.Stdout.Write(val1)

        // You can get converted value
        // number := row.Int(0)      // Zero value
        // str    := row.Str(1)      // First value
        // bignum := row.MustUint(2) // Second value

        // You may get values by column name
         first := res.Map("emailId")
         fmt.Println("column one is  " , row.Str(first))
        // second := res.Map("SecondColumn")
        // val1, val2 := row.Int(first), row.Str(second)
    }
}