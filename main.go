package main

import(
    //"net/http"

	"github.com/ronak-patidar/Banking-System/models"
	"github.com/gin-gonic/gin"   
	"github.com/ronak-patidar/Banking-System/controllers"
   
)

func main(){
	r := gin.Default()

	models.Connect()       //connection

	//Customer
	r.GET("/customer",controllers.Get_Customers)
    r.GET("/customer/:id",controllers.Get_Customer)
    r.POST("/customer",controllers.Create_Customer)
	//r.PUT("/customer/:id",controllers.Update_Customer)
	r.DELETE("/customer/:id",controllers.Delete_Customer)

	//Branch
	r.GET("/branch",controllers.Get_Branch)
	r.POST("/branch",controllers.Create_Branch)
	//r.PUT("/branch/:id",controllers.Update_Branch)
	r.DELETE("/branch/:id",controllers.Delete_Branch)
	
    //Account
	r.GET("/account",controllers.Get_Accounts)
	r.GET("/account/:id",controllers.Get_Account)
	r.POST("/account",controllers.Create_Account)
	//r.PUT("/account/:id",controllers.Update_Account)
	r.DELETE("/account/:id",controllers.Delete_Account)
   
	//Transaction
	r.GET("/transaction",controllers.Get_Transactions)
	r.GET("/transaction/:id",controllers.Get_Transaction)
	r.POST("/transaction",controllers.Perform_Transaction)
	//r.PUT("/transaction/:id",controllers.Update_Transaction)
	r.DELETE("/transaction/:id",controllers.Delete_Transaction)


	r.Run()
}