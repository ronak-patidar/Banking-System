package main

import(
    "net/http"	
    "fmt"
    "os"
	"github.com/gin-gonic/gin"   
    "github.com/go-pg/pg"
    //"github.com/go-pg/pg/orm"
)
var pg_db *pg.DB

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

	pg_db=db
}


func main(){
	r := gin.Default()

	Connect()
	r.GET("/customer",controllers.find_customers)
    r.GET("/customer/:id",controllers.find_customer)
    r.POST("/customer",controllers.add_customer)
	r.PUCH("/customer/:id",controllers.update_customer)
	r.DELETE("/customer/:id",controllers.delete_customer)
	r.Run()
}