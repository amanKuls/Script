package main

import (
	"fmt"
	"os"
	generaterecords "script/generate-records"
	"script/utils"
)

func main() {
	err := os.MkdirAll("./csvFiles", 0755)
	if err != nil {
		return
	}
	numberOfUsers := 1000
	fmt.Println("So I have to generate 100000 rows!!!")
	generaterecords.GenerateUserAndAccountRelatedTables(numberOfUsers)
	generaterecords.GenerateTransactionRelatedTables()
	utils.CreateFileWithContents("./csvFiles/BCNUser.csv", generaterecords.BCNUserTable)
	utils.CreateFileWithContents("./csvFiles/BCNAccount.csv", generaterecords.BCNAccountTable)
	utils.CreateFileWithContents("./csvFiles/Password.csv", generaterecords.PasswordTable)
	utils.CreateFileWithContents("./csvFiles/DefaultBCNAccount.csv", generaterecords.DefaultBCNAccountTable)
	utils.CreateFileWithContents("./csvFiles/LeoLLT.csv", generaterecords.LeoLLTTable)
	utils.CreateFileWithContents("./csvFiles/BCNUserLLT.csv", generaterecords.BCNUserLLTTable)
	utils.CreateFileWithContents("./csvFiles/Transaction.csv", generaterecords.TransactionTable)
	utils.CreateFileWithContents("./csvFiles/SendMoneyWithinBCN.csv", generaterecords.SendMoneyWithInBCNTable)
	utils.CreateFileWithContents("./csvFiles/SendMoneyAccountToAccount.csv", generaterecords.SendMoneyAccountToAccountTable)
	//utils.CreateFileWithContents("./csvFiles/Readme.md", Readme)

}
