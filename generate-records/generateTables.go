package generaterecords

import (
	"fmt"
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
var NumberOfTransactionTable []string

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
		LeoLLTLLT := fmt.Sprintf("%s--%s", uuid.NewString(), uuid.NewString())
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

func GenerateTransactionRelatedTables(numberOfUsers int) {

	genereateRecords(0, (numberOfUsers / 10), numberOfUsers*10)
	genereateRecords(((numberOfUsers / 10) + 1), ((numberOfUsers / 10) * 4), numberOfUsers*2)
	genereateRecords((((numberOfUsers / 10) * 4) + 1), ((numberOfUsers / 10) * 6), numberOfUsers/2)
	genereateRecords((((numberOfUsers / 10) * 6) + 1), ((numberOfUsers / 10) * 8), numberOfUsers/5)
	genereateRecords((((numberOfUsers / 10) * 8) + 1), ((numberOfUsers / 10) * 10), numberOfUsers/10)
	generateNumberOfTransactionTable(numberOfUsers)
}

func genereateRecords(start int, end int, numOfTransaction int) {
	generateSendMoneyWithinBCNData(start, end, numOfTransaction/3)
	generateSendMoneyAccountToAccount(start, end, numOfTransaction/3)
	generateLoadMoneyMPGS(start, end, numOfTransaction/3)
}
