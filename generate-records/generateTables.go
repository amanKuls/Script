package generaterecords

import (
	"fmt"

	"github.com/google/uuid"
)

var BCNUserTable string
var BCNAccountTable string
var PasswordTable string
var DefaultBCNAccountTable string
var LeoLLTTable string
var BCNUserLLTTable string

func GenerateTables(NumberOfRows int) {
	BCNUserUUID := uuid.NewString()
	BCNAccountUUID := uuid.NewString()
	LeoLLTLLT := uuid.NewString()
	BCNUserTable = getRowValuesForBCNUser(BCNUserUUID)
	BCNAccountTable = getRowValuesForBCNAccount(BCNAccountUUID, BCNUserUUID)
	PasswordTable = getRowValuesForPassword(BCNUserUUID)
	DefaultBCNAccountTable = getRowValuesForDefaulsBCNAccount(BCNUserUUID, BCNAccountUUID)
	LeoLLTTable = getRowValuesForLeoLLT(LeoLLTLLT)
	BCNUserLLTTable = getRowValuesForBCNUserLLT(BCNUserUUID, LeoLLTLLT)
	for i := 1; i < NumberOfRows; i++ {
		BCNUserUUID = uuid.NewString()
		BCNAccountUUID = uuid.NewString()
		LeoLLTLLT = uuid.NewString()
		BCNUserTable = fmt.Sprintf("%s\n%s", BCNUserTable, getRowValuesForBCNUser(BCNUserUUID))
		BCNAccountTable = fmt.Sprintf("%s\n%s", BCNAccountTable, getRowValuesForBCNAccount(BCNAccountUUID, BCNUserUUID))
		PasswordTable = fmt.Sprintf("%s\n%s", PasswordTable, getRowValuesForPassword(BCNUserUUID))
		DefaultBCNAccountTable = fmt.Sprintf("%s\n%s", DefaultBCNAccountTable, getRowValuesForDefaulsBCNAccount(BCNUserUUID, BCNAccountUUID))
		LeoLLTTable = fmt.Sprintf("%s\n%s", LeoLLTTable, getRowValuesForLeoLLT(LeoLLTLLT))
		BCNUserLLTTable = fmt.Sprintf("%s\n%s", BCNUserLLTTable, getRowValuesForBCNUserLLT(BCNUserUUID, LeoLLTLLT))
	}
}
