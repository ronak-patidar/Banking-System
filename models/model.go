package models

type Customer struct{
	Cust_Id int         `sql:"id,pk"`
	Unique_Id_Type text `sql:"unique_id_type" json:"unique_id_type"`
	Unique_Id text      `sql:"unique_id,unique" json:"unique_id"`
    Name text 			`sql:"name", json:"name"`
	Addr text  			`sql:"addr", json:"addr"`
	Phone text 			`sql:"phone", json:"phone"`
    Email text 			`sql:"email", json:"email"`
	Age int 			`sql:"age", json:"age"`
}
