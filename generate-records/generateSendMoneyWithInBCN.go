package generaterecords

import (
	"fmt"
	"script/utils"

	"github.com/google/uuid"
)

// "id", "recipientBcnAccountId", "amount", "senderBcnAccountId",
// "claimedFeeCustomer", "claimedFeeExpiresAt", "status",
// "transactionId", "feeCustomerTransactionId"
func getRowValuesForSendMoneyWithInBCN(
	senderBcnAccountId string,
	recipientBcnAccountId string,
	transactionID string,
	feeCustomerTransactionId string,
) string {
	SendMoneyWithInBCNUUID := uuid.NewString()
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s\n",
		SendMoneyWithInBCNUUID,
		recipientBcnAccountId,
		utils.Amount,
		senderBcnAccountId,
		utils.ClaimedFeeCustomer,
		utils.Timestamp,
		utils.Status,
		transactionID,
		feeCustomerTransactionId,
	)

}

func generateSendMoneyWithinBCNData(start int, end int, NumberOfTransactions int) {
	for i := 1; i <= NumberOfTransactions; i++ {
		senderNumber := utils.GetRandomNumberBetweenRange(start, end)
		recieverNumber := utils.GetRandomNumberBetweenRangeExcept(start, end, senderNumber)
		transactionID := uuid.NewString()
		sender := allUsers[senderNumber].BCNAccountUUID1
		reciever := allUsers[recieverNumber].BCNAccountUUID1
		feeCustomerTransactionId := uuid.NewString()
		TransactionTable = append(
			TransactionTable,
			getRowValuesForTransaction(
				transactionID,
				sender,
				reciever,
				senderNumber,
				recieverNumber,
			))
		TransactionTable = append(
			TransactionTable,
			getRowValuesForTransaction(
				feeCustomerTransactionId,
				sender,
				utils.CustomerFeeBcnAccountId,
				senderNumber,
				-1,
			))
		SendMoneyWithInBCNTable = append(
			SendMoneyWithInBCNTable,
			getRowValuesForSendMoneyWithInBCN(
				sender,
				reciever,
				transactionID,
				feeCustomerTransactionId,
			))

	}
}
