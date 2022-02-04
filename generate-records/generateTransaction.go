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
	is_exchange_rate_applied bool,
) string {
	var exchangeRateId string = ""
	var exchangeRateMultiplier string = ""
	if is_exchange_rate_applied {
		exchangeRateId = "3e8b871a-f6f6-4bfc-852c-35b553d8a1b1"
		exchangeRateMultiplier = "204"
	}
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s",
		transactionUUID,
		utils.Amount,
		senderBcnAccountId,
		recipientBcnAccountId,
		utils.Timestamp,
		exchangeRateMultiplier,
		exchangeRateId,
	)

}
