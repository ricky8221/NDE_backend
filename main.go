package main

import (
	"fmt"
	sqlcFunc "github.com/ricky8221/NDE_DB/sqlc_func"
)

func main() {
	res := sqlcFunc.CreateCompany{CompanyName: "ABC Company", CompanyContactName: "John Doe", CompanyContactNumber: "123-123-1231", Remark: "null"}

	if res.CompanyName == "ABC Company" {
		fmt.Print("create account successfully")
	}
}
