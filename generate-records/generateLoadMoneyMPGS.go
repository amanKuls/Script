package generaterecords

import (
	"fmt"
	"script/utils"

	"github.com/google/uuid"
)

// "id", "amount", "recipientAccountId", "confirmationExpiresAt",
//  "status", "createdAt"
func getRowValuesForLoadMoneyMPGS(
	recipientAccountId string,
	counterPartyTransactionId string,
	loadMoneyMPGSUUID string,
) string {
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s\n",
		loadMoneyMPGSUUID,
		utils.Amount,
		recipientAccountId,
		utils.Timestamp,
		utils.Status,
		utils.Timestamp,
	)
}

func generateLoadMoneyMPGS(start int, end int, NumberOfTransactions int) {
	for i := 0; i < NumberOfTransactions; i++ {
		var loadMoneyMPGSUUID = uuid.NewString()
		recipientNumber := utils.GetRandomNumberBetweenRange(start, end)
		recipientBCNAccountId := allUsers[recipientNumber].BCNAccountUUID1
		counterPartyTransactionUUID := uuid.NewString()
		counterPartyFinalTransactionId := uuid.NewString()
		TransactionTable = append(
			TransactionTable,
			getRowValuesForTransaction(
				counterPartyFinalTransactionId,
				utils.MPGSHoldingBcnAcccountId,
				recipientBCNAccountId,
				false,
				recipientNumber,
				-1,
			),
		)
		UserVisibleTransactionTable = append(
			UserVisibleTransactionTable,
			getRowValueForUserVisibleTransactionTable(
				loadMoneyMPGSUUID,
				recipientBCNAccountId,
				false,
			),
		)
		CounterPartyTransactionTable = append(
			CounterPartyTransactionTable,
			getRowValuesForCounterPartyTransaction(
				counterPartyTransactionUUID,
				counterPartyFinalTransactionId,
			),
		)
		LoadMoneyMPGSTable = append(
			LoadMoneyMPGSTable,
			getRowValuesForLoadMoneyMPGS(
				recipientBCNAccountId,
				counterPartyTransactionUUID,
				loadMoneyMPGSUUID,
			),
		)
	}
}
