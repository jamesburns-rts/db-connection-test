package main

// connecting to a PostgreSQL database with Go's database/sql package
import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	connStr := flag.String("connection-string", "", `entire postgres connection string:
	user: using default user for postgres database
	dbname: using default database that comes with postgres
	password: password used during initial setup
	host: hostname or IP Address of server
	sslmode: must be set to disabled unless using SSL
    `)

	// todo add above as separate flags
	// todo add other db implementations

	flag.Parse()

	if *connStr == "" {
		flag.Usage()
		os.Exit(1)
	}

	db, err := sql.Open("postgres", *connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
}

