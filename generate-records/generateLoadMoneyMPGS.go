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

func generateLoadMoneyMPGS(start int, end int, NumberOfTransactions int) {
	for i := 0; i < NumberOfTransactions; i++ {
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
				recipientNumber,
				-1,
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
