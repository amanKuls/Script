package main

import (
	"fmt"
	"os"
	generaterecords "script/generate-records"
	"script/utils"
	"strconv"
)

func main() {

	numberOfUsers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = os.MkdirAll("./csvFiles", 0755)
	if err != nil {
		return
	}
	generaterecords.GenerateUserAndAccountRelatedTables(numberOfUsers)
	generaterecords.GenerateTransactionRelatedTables(numberOfUsers)
	generateFiles()
}

func generateFiles() {
	utils.CreateFileWithContents("./csvFiles/BCNUser.csv", generaterecords.BCNUserTable)
	utils.CreateFileWithContents("./csvFiles/BCNAccount.csv", generaterecords.BCNAccountTable)
	utils.CreateFileWithContents("./csvFiles/Password.csv", generaterecords.PasswordTable)
	utils.CreateFileWithContents("./csvFiles/DefaultBCNAccount.csv", generaterecords.DefaultBCNAccountTable)
	utils.CreateFileWithContents("./csvFiles/LeoLLT.csv", generaterecords.LeoLLTTable)
	utils.CreateFileWithContents("./csvFiles/BCNUserLLT.csv", generaterecords.BCNUserLLTTable)
	utils.CreateFileWithContents("./csvFiles/Transaction.csv", generaterecords.TransactionTable)
	utils.CreateFileWithContents("./csvFiles/SendMoneyWithinBCN.csv", generaterecords.SendMoneyWithInBCNTable)
	utils.CreateFileWithContents("./csvFiles/SendMoneyAccountToAccount.csv", generaterecords.SendMoneyAccountToAccountTable)
	utils.CreateFileWithContents("./csvFiles/CounterPartyTransaction.csv", generaterecords.CounterPartyTransactionTable)
	utils.CreateFileWithContents("./csvFiles/LoadMoneyMPGS.csv", generaterecords.LoadMoneyMPGSTable)
	utils.CreateFileWithContents("./csvFiles/NumberOfTransaction.csv", generaterecords.NumberOfTransactionTable)
}
