package main

import (
	"flag"
	"fmt"
	"os"
	generaterecords "script/generate-records"
	"script/utils"
)

func main() {
	num := flag.Int("n", 1000, "number of users")
	od := flag.String("o", ".", "Output directory")

	flag.Parse()
	numberOfUsers := *num
	outputDirectory := *od
	if outputDirectory != "" {
		outputDirectory = fmt.Sprintf("%s/%s", outputDirectory, "csvFiles")
	} else {
		outputDirectory = "./csvFiles"
	}
	err := os.MkdirAll(outputDirectory, 0755)
	if err != nil {
		return
	}
	generaterecords.GenerateUserAndAccountRelatedTables(numberOfUsers)
	generaterecords.GenerateTransactionRelatedTables(numberOfUsers)
	generateFiles(outputDirectory)
}

func generateFiles(outputDirectory string) {
	utils.CreateFileWithContents(outputDirectory+"/BCNUser.csv", generaterecords.BCNUserTable)
	utils.CreateFileWithContents(outputDirectory+"/BCNAccount.csv", generaterecords.BCNAccountTable)
	utils.CreateFileWithContents(outputDirectory+"/Password.csv", generaterecords.PasswordTable)
	utils.CreateFileWithContents(outputDirectory+"/DefaultBCNAccount.csv", generaterecords.DefaultBCNAccountTable)
	utils.CreateFileWithContents(outputDirectory+"/LeoLLT.csv", generaterecords.LeoLLTTable)
	utils.CreateFileWithContents(outputDirectory+"/BCNUserLLT.csv", generaterecords.BCNUserLLTTable)
	utils.CreateFileWithContents(outputDirectory+"/Transaction.csv", generaterecords.TransactionTable)
	utils.CreateFileWithContents(outputDirectory+"/SendMoneyWithinBCN.csv", generaterecords.SendMoneyWithInBCNTable)
	utils.CreateFileWithContents(outputDirectory+"/SendMoneyAccountToAccount.csv", generaterecords.SendMoneyAccountToAccountTable)
	utils.CreateFileWithContents(outputDirectory+"/CounterPartyTransaction.csv", generaterecords.CounterPartyTransactionTable)
	utils.CreateFileWithContents(outputDirectory+"/LoadMoneyMPGS.csv", generaterecords.LoadMoneyMPGSTable)
	utils.CreateFileWithContents(outputDirectory+"/NumberOfTransaction.csv", generaterecords.NumberOfTransactionTable)
}
