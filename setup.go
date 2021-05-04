package main

import (
	"net/http"	
    "fmt"
    "os"
	"github.com/gin-gonic/gin"   
    "github.com/go-pg/pg"
    //"github.com/go-pg/pg/orm"
)
var pg_db *pg.DB

type Customer struct{
	Cust_Id int      	    `sql:"id,pk"`
	Unique_Id_Type string	`sql:"unique_id_type" json:"unique_id_type"`
	Unique_Id string      	`sql:"unique_id,unique" json:"unique_id"`
    Name string 			`sql:"name", json:"name"`
	Addr string  			`sql:"addr", json:"addr"`
	Phone string 			`sql:"phone", json:"phone"`
    Email string 			`sql:"email", json:"email"`
	Age int 				`sql:"age", json:"age"`
}

func delete_customer(c *gin.Context){
	var cust Customer
	err := pg_db.Model(&cust).Where("id = ?", c.Param("id")).Select() 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	pg_db.Delete(&cust)

	c.JSON(http.StatusOK, gin.H{"data": true})
}



/*func update_customer(c *gin.Context){
	var cust Customer
	err := pg_db.Model(&cust).Where("id = ?", c.Param("id")).Select()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input Customer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pg_db.Model(&cust).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": cust})
}*/

func find_customers(c *gin.Context){
	var custs []Customer
	pg_db.Model(&custs).Select()

	c.JSON(http.StatusOK, gin.H{"data": custs})
}

func find_customer(c *gin.Context){
	var cust Customer
	err := pg_db.Model(&cust).Where("id = ?", c.Param("id")).Select() 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cust})
}

func add_customer(c *gin.Context){
	var input Customer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	item := Customer{
		Unique_Id_Type: input.Unique_Id_Type,
		Unique_Id: input.Unique_Id,
		Name: input.Name,
		Addr: input.Addr,
		Phone: input.Phone,
		Email: input.Email,
		Age: input.Age,
	}
	//pg_db.Create(&item)
    pg_db.Insert(&item)
	c.JSON(http.StatusOK, gin.H{"data": item})
}


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

/*func CreateCustomerTable(db *pg.DB) error{
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
    }
	createErr := db.CreateTable(&Customer{},opts)
	if createErr != nil{
		fmt.Printf("Error while creating table Customer, Reason: %v\n", createErr)
		return createErr
	}
	fmt.Printf("Table Customer created successfully.\n")
	return nil
}*/


//func main(){
	//r := gin.Default()

	//models.Connect()
	/*r.GET("/customer",find_customers)
    r.GET("/customer/:id",find_customer)
    r.POST("/customer",add_customer)
	//r.PATCH("/customer/:id",update_customer)
	r.DELETE("/customer/:id",delete_customer)
	r.Run()*/
//}



package main

import (
	
	"net/http"	
    "fmt"
    "os"
	
	"github.com/gin-gonic/gin"   
    "github.com/go-pg/pg"
    "github.com/go-pg/pg/orm"
)
var pg_db *pg.DB

type Transaction struct{
	Trans_Id int      	    `sql:"trans_id,pk"`
	Second_Account_No int	`sql:"second_account_no,fk" json:"second_account_no"`
	Receiver_Account_No int `sql:"receiver_account_no,fk" json:"receiver_account_no"`
    Amount double precision	`sql:"amount", json:"amount"`
	Time timestamptz  		`sql:"time", json:"time"`
}
func get_transactions(c *gin.Context){
	var trans []Transaction
	pg_db.Model(&trans).Select()

	c.JSON(http.StatusOK, gin.H{"data": trans})
}

func get_transaction(c *gin.Context){
	var trans Transaction
	err := pg_db.Model(&trans).Where("trans_id = ?", c.Param("id")).Select() 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": trans})
}

func create_transaction(c *gin.Context){
	var input Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
    item := Transaction{
		Second_Account_No: input.Second_Account_No,
		Receiver_Account_No: input.Receiver_Account_No,
		Amount: input.Amount,
	}
    pg_db.Insert(&item)
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func delete_Transaction(c *gin.Context){
	var trans Transaction
	err := pg_db.Model(&trans).Where("trans_id" = ?", c.Param("id")).Select() 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	pg_db.Delete(&trans)

	c.JSON(http.StatusOK, gin.H{"data": true})
}




*/
/*func delete_customer(c *gin.Context){
	var cust Customer
	err := pg_db.Model(&cust).Where("id = ?", c.Param("id")).Select() 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	pg_db.Delete(&cust)

	c.JSON(http.StatusOK, gin.H{"data": true})
}*/

/*func update_customer(c *gin.Context){
	var cust Customer
	err := pg_db.Model(&cust).Where("id = ?", c.Param("id")).Select()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input Customer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pg_db.Model(&cust).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": cust})
}*/

/*func find_customers(c *gin.Context){
	var custs []Customer
	pg_db.Model(&custs).Select()

	c.JSON(http.StatusOK, gin.H{"data": custs})
}*/

/*func find_customer(c *gin.Context){
	var cust Customer
	err := pg_db.Model(&cust).Where("id = ?", c.Param("id")).Select() 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cust})
}*/
/*func perform_transaction(c *gin.Context){
	var input Transaction
    
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    
	account_no:=input.Sender_Account_No
	
	var acc Account
    
	err := pg_db.Model(&acc).Where("acc_no = ?", account_no).Select()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if acc.account_no==0{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Account Number!"})
		return
	}
}*/
/*func add_customer(c *gin.Context){
	var input Customer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
    item := Customer{
		Unique_Id_Type: input.Unique_Id_Type,
		Unique_Id: input.Unique_Id,
		Name: input.Name,
		Addr: input.Addr,
		Phone: input.Phone,
		Email: input.Email,
		Age: input.Age,
	}
	//pg_db.Create(&item)
    pg_db.Insert(&item)
	c.JSON(http.StatusOK, gin.H{"data": item})
}*/


/*func Connect(){
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
}*/

/*func CreateCustomerTable(db *pg.DB) error{
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
    }
	createErr := db.CreateTable(&Customer{},opts)
	if createErr != nil{
		fmt.Printf("Error while creating table Customer, Reason: %v\n", createErr)
		return createErr
	}
	fmt.Printf("Table Customer created successfully.\n")
	return nil
}*/


func main(){
	//r := gin.Default()

	Connect()
	/*r.GET("/customer",find_customers)
    r.GET("/customer/:id",find_customer)
    r.POST("/customer",add_customer)
	//r.PATCH("/customer/:id",update_customer)
	r.DELETE("/customer/:id",delete_customer)
	r.Run()*/

}*/
