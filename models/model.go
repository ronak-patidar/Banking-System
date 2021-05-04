package models

import "time"


type Transaction struct{
	Trans_Id int      	    `sql:"trans_id,pk"`
	Sender_Account_No int	`sql:"sender_account_no,REFERENCES Account(account_no) ON DELETE CASCADE ON UPDATE CASCADE" json:"sender_account_no"`
	Receiver_Account_No int `sql:"receiver_account_no,REFERENCES Account(account_no) ON DELETE CASCADE ON UPDATE CASCADE" json:"receiver_account_no"`
    Amount float64			`sql:"amount", json:"amount"`
	Time   time.Time		`sql:"time", json:"time"`
}

type Account struct{
	Account_No int         	`sql:"account_no,pk"`
	Balance float64 		`sql:"balance", json:"balance"`
	Type string             `sql:"type", json:"type"`
	Branch_Id int           `sql:"branch_id,REFERENCES Branch(branch_id) ON DELETE CASCADE ON UPDATE CASCADE" json:"branch_id"`
	Cust_Id int             `sql:"cust_id,REFERENCES Customer(cust_id) ON DELETE CASCADE ON UPDATE CASCADE" json:"cust_id"`
}

type Branch struct{
	Branch_Id int           `sql:"branch_id,pk"`
	Branch_Name string      `sql:"branch_name", json:"branch_name"`
	Branch_Addr string      `sql:"branch_addr", json:"branch_addr"`
	IFSC string             `sql:"ifsc", json:"ifsc"`
}