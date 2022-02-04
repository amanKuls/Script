package generaterecords

import (
	"script/utils"

	"github.com/google/uuid"
)

var BCNUserTable []string
var BCNAccountTable []string
var PasswordTable []string
var DefaultBCNAccountTable []string
var LeoLLTTable []string
var BCNUserLLTTable []string

// Transaction related tables

var SendMoneyAccountToAccountTable []string
var SendMoneyWithInBCNTable []string
var TransactionTable []string
var CounterPartyTransactionTable []string
var LoadMoneyMPGSTable []string

type User struct {
	NoOfTransactions int
	BCNUserUUID      string
	BCNAccountUUID1  string
	BCNAccountUUID2  string
	LeoLLTLLT        string
}

var allUsers []User

func GenerateUserAndAccountRelatedTables(NumberOfRows int) {
	for i := 0; i < NumberOfRows; i++ {
		BCNUserUUID := uuid.NewString()
		BCNAccountUUID1 := uuid.NewString()
		BCNAccountUUID2 := uuid.NewString()
		LeoLLTLLT := uuid.NewString()
		temp := User{BCNUserUUID: BCNUserUUID, BCNAccountUUID1: BCNAccountUUID1, BCNAccountUUID2: BCNAccountUUID2, LeoLLTLLT: LeoLLTLLT}
		allUsers = append(allUsers, temp)
		BCNUserTable = append(BCNUserTable, getRowValuesForBCNUser(BCNUserUUID))
		BCNAccountTable = append(BCNAccountTable, getRowValuesForBCNAccount(BCNAccountUUID1, BCNUserUUID, utils.CurrencyCode1))
		BCNAccountTable = append(BCNAccountTable, getRowValuesForBCNAccount(BCNAccountUUID2, BCNUserUUID, utils.CurrencyCode2))
		PasswordTable = append(PasswordTable, getRowValuesForPassword(BCNUserUUID))
		DefaultBCNAccountTable = append(DefaultBCNAccountTable, getRowValuesForDefaulsBCNAccount(BCNUserUUID, BCNAccountUUID1))
		LeoLLTTable = append(LeoLLTTable, getRowValuesForLeoLLT(LeoLLTLLT))
		BCNUserLLTTable = append(BCNUserLLTTable, getRowValuesForBCNUserLLT(BCNUserUUID, LeoLLTLLT))
	}
}

func GenerateTransactionRelatedTables() {
	generateSendMoneyWithinBCNData()
	generateSendMoneyAccountToAccount()
	generateLoadMoneyMPGS()
}
