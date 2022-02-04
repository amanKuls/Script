package generaterecords

import (
	"fmt"
)

// "id", "counterpartyFinalTransactionId"
func getRowValuesForCounterPartyTransaction(
	counterPartyTransactionUUID,
	finalTransactionId string,
) string {
	return fmt.Sprintf("%s,%s\n",
		counterPartyTransactionUUID,
		finalTransactionId,
	)
}
