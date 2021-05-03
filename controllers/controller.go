package controllers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ronak-patidar/Banking-System/models"
)

func delete_customer(c *gin.Context){
	var cust models.Customer
	err := pg_db.Model(&cust).Where("id = ?", c.Param("id")).Select() 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	pg_db.Delete(&cust)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func update_customer(c *gin.Context){
	var cust models.Customer
	err := pg_db.Model(&cust).Where("id = ?", c.Param("id")).Select()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.Customer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pg_db.Model(&cust).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": cust})
}

func find_customers(c *gin.Context){
	var custs []models.Customer
	pg_db.Model(&custs).Select()

	c.JSON(http.StatusOK, gin.H{"data": custs})
}

func find_customer(c *gin.Context){
	var cust models.Customer
	err := pg_db.Model(&cust).Where("id = ?", c.Param("id")).Select() 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cust})
}

func add_customer(c *gin.Context){
	var input models.Customer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
    item := models.Customer{
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
