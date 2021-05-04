package models

import(	
    "fmt"
    "os"
	"github.com/go-pg/pg/orm"
    "github.com/go-pg/pg"
)
var Pg_db *pg.DB

func Connect(){
	opts:= &pg.Options{
		User: "postgres",
		Password: "root",
		Addr: "localhost:5432",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil{
		fmt.Printf("Failed to connect to database.\n")
		os.Exit(10)
	}
	fmt.Printf("Connection to database successful.\n")

	Pg_db=db
	//CreateTable(Pg_db)
}

func CreateTable(db *pg.DB) error{
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
    }
	createErr := db.CreateTable(&Transaction{},opts)
	if createErr != nil{
		fmt.Printf("Error while creating Transaction Account, Reason: %v\n", createErr)
		return createErr
	}
	fmt.Printf("Table Transaction created successfully.\n")
	return nil
}