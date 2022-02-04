package generaterecords

import (
	"fmt"
	"script/utils"

	"github.com/google/uuid"
)

// "id", "amount", "recipientAccountId", "counterpartyId", "claimedFeeCounterparty",
// "claimedFeeCustomer", "claimedFeeExpiresAt", "status", "counterpartyTransactionId"
func getRowValuesForLoadMoneyMPGS(
	recipientAccountId string,
	counterPartyTransactionId string,
) string {
	loadMoneyMPGSUUID := uuid.NewString()
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s\n",
		loadMoneyMPGSUUID,
		utils.Amount,
		recipientAccountId,
		utils.CounterpartyId,
		utils.ClaimedFeeCounterparty,
		utils.ClaimedFeeCustomer,
		utils.Timestamp,
		utils.Status,
		counterPartyTransactionId,
	)
}

func generateLoadMoneyMPGS() {
	for i := 0; i < 100; i++ {
		recipientBCNAccountId := allUsers[i].BCNAccountUUID1
		counterPartyTransactionUUID := uuid.NewString()
		counterPartyFinalTransactionId := uuid.NewString()
		TransactionTable = append(
			TransactionTable,
			getRowValuesForTransaction(
				counterPartyFinalTransactionId,
				utils.MPGSHoldingBcnAcccountId,
				recipientBCNAccountId,
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
			),
		)
	}
}
