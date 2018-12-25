package accessory

type Bill struct{
	BillId 	  		string 		`json:"billId"`
	BillMoney 		int			`json:"billMoney"`
	BillCmt			string		`json:"billCmt"`
	BillDate		string 	 	`json:"billDate"`
	BillCategory 	string 	 	`json:"billcategory"`
}

