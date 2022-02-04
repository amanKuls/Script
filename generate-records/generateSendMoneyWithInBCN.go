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

func generateSendMoneyWithinBCNData() {
	for i := 2; i < 100; i = i + 2 {
		transactionID := uuid.NewString()
		sender := allUsers[i].BCNAccountUUID1
		reciever := allUsers[i+1].BCNAccountUUID1
		feeCustomerTransactionId := uuid.NewString()
		TransactionTable = append(
			TransactionTable,
			getRowValuesForTransaction(
				transactionID,
				sender,
				reciever,
			))
		TransactionTable = append(
			TransactionTable,
			getRowValuesForTransaction(
				feeCustomerTransactionId,
				sender,
				utils.CustomerFeeBcnAccountId,
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
