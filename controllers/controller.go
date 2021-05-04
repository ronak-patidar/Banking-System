package controllers

import(
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ronak-patidar/Banking-System/models"
)


//Branch

func Delete_Branch(c *gin.Context){
	var branch models.Branch
	err := models.Pg_db.Model(&branch).Where("branch_id = ?", c.Param("id")).Select() 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.Pg_db.Delete(&branch)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func Get_Branch(c *gin.Context){
	var branch []models.Branch
	models.Pg_db.Model(&branch).Select()

	c.JSON(http.StatusOK, gin.H{"data": branch})
}

func Create_Branch(c *gin.Context){
	var input models.Branch
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
    item := models.Branch{
		Branch_Name: input.Branch_Name,
		Branch_Addr: input.Branch_Addr,
		IFSC: input.IFSC,
	}
    models.Pg_db.Insert(&item)
	c.JSON(http.StatusOK, gin.H{"data": item})
}


//Account

func Get_Accounts(c *gin.Context){
	var acc []models.Account
	models.Pg_db.Model(&acc).Select()

	c.JSON(http.StatusOK, gin.H{"data": acc})
}

func Get_Account(c *gin.Context){
	var acc models.Customer
	err := models.Pg_db.Model(&acc).Where("account_no = ?", c.Param("id")).Select() 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": acc})
}

func Create_Account(c *gin.Context){
	var input models.Account
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
    item := models.Account{
		Account_No:input.Account_No,
		Balance:input.Balance,
		Type:input.Type,
		Branch_Id:input.Branch_Id,
		Cust_Id:input.Cust_Id,
	}
    models.Pg_db.Insert(&item)
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func Delete_Account(c *gin.Context){
	var acc models.Account
	err := models.Pg_db.Model(&acc).Where("account_no = ?", c.Param("id")).Select() 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.Pg_db.Delete(&acc)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

//Transaction

func Get_Transactions(c *gin.Context){
	var trans []models.Transaction
	models.Pg_db.Model(&trans).Select()

	c.JSON(http.StatusOK, gin.H{"data": trans})
}

func Get_Transaction(c *gin.Context){
	var trans models.Transaction
	err := models.Pg_db.Model(&trans).Where("trans_id = ?", c.Param("id")).Select() 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": trans})
}


func Delete_Transaction(c *gin.Context){
	var trans models.Transaction
	err := models.Pg_db.Model(&trans).Where("trans_id = ?", c.Param("id")).Select() 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.Pg_db.Delete(&trans)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func Perform_Transaction(c *gin.Context){
	var input models.Transaction
    
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    
	sender_account_no:=input.Sender_Account_No
    receiver_account_no:=input.Receiver_Account_No
	amount:=input.Amount
	
	var acc1 models.Account
	var acc2 models.Account
    
	
	if err := models.Pg_db.Model(&acc1).Where("account_no = ?", sender_account_no).Select(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    
	if acc1.Account_No==0{
	 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Sender Account Number!"})
	 	return
	}

	
	if err := models.Pg_db.Model(&acc2).Where("account_no = ?", receiver_account_no).Select(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if acc2.Account_No==0{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Receiver Account Number!"})
		return
	}

    if (acc1.Balance-amount) < 500 {
	 	c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient Amount!"})
	 	return
	}
    newAcc :=&models.Account{
		Account_No : sender_account_no,
		Balance : acc1.Balance-amount,
	}
	_, updateErr1 := models.Pg_db.Model(newAcc).Set("balance = ?balance").Where("account_no = ?account_no").Update()
	if updateErr1 != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error": updateErr1.Error()})
		return
	}

	newAcc1 :=&models.Account{
		Account_No : receiver_account_no,
		Balance : acc2.Balance+amount,
	}

	_, updateErr2 := models.Pg_db.Model(newAcc1).Set("balance = ?balance").Where("account_no = ?account_no").Update()
	if updateErr2 != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error": updateErr2.Error()})
		return
	}
	
	

	item := models.Transaction{
		Sender_Account_No: input.Sender_Account_No,
		Receiver_Account_No: input.Receiver_Account_No,
		Amount: input.Amount,
		Time:time.Now(),
	}
    models.Pg_db.Insert(&item)
	c.JSON(http.StatusOK, gin.H{"data": item})
	
}