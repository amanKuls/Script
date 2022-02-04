package generaterecords

import (
	"fmt"
	"script/utils"

	"github.com/google/uuid"
)

var BCNUserTable string
var BCNAccountTable string
var PasswordTable string
var DefaultBCNAccountTable string
var LeoLLTTable string
var BCNUserLLTTable string

// Transaction related tables

var SendMoneyAccountToAccountTable string
var SendMoneyWithInBCNTable string
var TransactionTable string

type Globals struct {
	BCNUserUUID     string
	BCNAccountUUID1 string
	BCNAccountUUID2 string
	LeoLLTLLT       string
}

var allGlobalValues []Globals

func GenerateUserAndAccountRelatedTables(NumberOfRows int) {
	fmt.Println("I am in GenerateUserAndAccountRelatedTables")
	BCNUserUUID := uuid.NewString()
	BCNAccountUUID1 := uuid.NewString()
	BCNAccountUUID2 := uuid.NewString()
	LeoLLTLLT := uuid.NewString()
	temp := Globals{BCNUserUUID: BCNUserUUID, BCNAccountUUID1: BCNAccountUUID1, BCNAccountUUID2: BCNAccountUUID2, LeoLLTLLT: LeoLLTLLT}
	allGlobalValues = append(allGlobalValues, temp)
	BCNUserTable = getRowValuesForBCNUser(BCNUserUUID)
	BCNAccountTable = getRowValuesForBCNAccount(BCNAccountUUID1, BCNUserUUID, utils.CurrencyCode1)
	BCNAccountTable = fmt.Sprintf("%s\n%s", BCNAccountTable, getRowValuesForBCNAccount(BCNAccountUUID2, BCNUserUUID, utils.CurrencyCode2))
	PasswordTable = getRowValuesForPassword(BCNUserUUID)
	DefaultBCNAccountTable = getRowValuesForDefaulsBCNAccount(BCNUserUUID, BCNAccountUUID1)
	LeoLLTTable = getRowValuesForLeoLLT(LeoLLTLLT)
	BCNUserLLTTable = getRowValuesForBCNUserLLT(BCNUserUUID, LeoLLTLLT)
	fmt.Println("Entering for loop")
	for i := 1; i < NumberOfRows; i++ {
		fmt.Println(i)
		BCNUserUUID := uuid.NewString()
		BCNAccountUUID1 := uuid.NewString()
		BCNAccountUUID2 := uuid.NewString()
		LeoLLTLLT := uuid.NewString()
		temp := Globals{BCNUserUUID: BCNUserUUID, BCNAccountUUID1: BCNAccountUUID1, BCNAccountUUID2: BCNAccountUUID2, LeoLLTLLT: LeoLLTLLT}
		allGlobalValues = append(allGlobalValues, temp)
		BCNUserTable = fmt.Sprintf("%s\n%s", BCNUserTable, getRowValuesForBCNUser(BCNUserUUID))
		BCNAccountTable = fmt.Sprintf("%s\n%s", BCNAccountTable, getRowValuesForBCNAccount(BCNAccountUUID1, BCNUserUUID, utils.CurrencyCode1))
		BCNAccountTable = fmt.Sprintf("%s\n%s", BCNAccountTable, getRowValuesForBCNAccount(BCNAccountUUID2, BCNUserUUID, utils.CurrencyCode2))
		PasswordTable = fmt.Sprintf("%s\n%s", PasswordTable, getRowValuesForPassword(BCNUserUUID))
		DefaultBCNAccountTable = fmt.Sprintf("%s\n%s", DefaultBCNAccountTable, getRowValuesForDefaulsBCNAccount(BCNUserUUID, BCNAccountUUID1))
		LeoLLTTable = fmt.Sprintf("%s\n%s", LeoLLTTable, getRowValuesForLeoLLT(LeoLLTLLT))
		BCNUserLLTTable = fmt.Sprintf("%s\n%s", BCNUserLLTTable, getRowValuesForBCNUserLLT(BCNUserUUID, LeoLLTLLT))
	}
}

func GenerateTransactionRelatedTables() {
	generateSendMoneyWithinBCNData()
	generateSendMoneyAccountToAccount()
}
