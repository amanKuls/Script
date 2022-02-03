package main

import (
	"os"
	generaterecords "script/generate-records"
	"script/utils"
)

func main() {
	err := os.MkdirAll("./csvFiles", 0755)
	if err != nil {
		return
	}
	numberOfRows := 20000
	generaterecords.GenerateTables(numberOfRows)
	utils.CreateFileWithContents("./csvFiles/BCNUser.csv", generaterecords.BCNUserTable)
	utils.CreateFileWithContents("./csvFiles/BCNAccount.csv", generaterecords.BCNAccountTable)
	utils.CreateFileWithContents("./csvFiles/Password.csv", generaterecords.PasswordTable)
	utils.CreateFileWithContents("./csvFiles/DefaultBCNAccount.csv", generaterecords.DefaultBCNAccountTable)
	utils.CreateFileWithContents("./csvFiles/LeoLLT.csv", generaterecords.LeoLLTTable)
	utils.CreateFileWithContents("./csvFiles/BCNUserLLT.csv", generaterecords.BCNUserLLTTable)

}
