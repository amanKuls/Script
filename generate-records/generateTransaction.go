package generaterecords

import (
	"fmt"
	"script/utils"
)

// id, amount, senderBcnAccountId, recipientBcnAccountId,
// resultedAt, exchangeRateMultiplier ,exchangeRateId
func getRowValuesForTransaction(
	transactionUUID string,
	senderBcnAccountId string,
	recipientBcnAccountId string,
	isExchangeRateApplied bool,
	userNumber1 int,
	userNumber2 int,
) string {
	if userNumber1 >= 0 {
		allUsers[userNumber1].NoOfTransactions = allUsers[userNumber1].NoOfTransactions + 1
	}
	if userNumber2 >= 0 {
		allUsers[userNumber2].NoOfTransactions = allUsers[userNumber2].NoOfTransactions + 1
	}
	exchangeRateId := "null"
	exchangeRateMultiplier := "null"

	if isExchangeRateApplied {
		exchangeRateId = "3e8b871a-f6f6-4bfc-852c-35b553d8a1b1"
		exchangeRateMultiplier = "204"
	}
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s\n",
		transactionUUID,
		utils.Amount,
		senderBcnAccountId,
		recipientBcnAccountId,
		utils.TransactionResultedAt,
		exchangeRateMultiplier,
		exchangeRateId,
	)

}
