package generaterecords

import (
	"fmt"
	"script/utils"

	"github.com/google/uuid"
)

// id", "recipientBcnAccountId", "amount", "senderBcnAccountId",
// "claimedFeeCustomer", "claimedFeeExpiresAt", "status",
// "transactionId", "feeCustomerTransactionId"
func getRowValuesForSendMoneyWithInBCN(
	senderBcnAccountId string,
	recipientBcnAccountId string,
	transactionID string,
	feeCustomerTransactionId string,
) string {
	SendMoneyWithInBCNUUID := uuid.NewString()
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s",
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
	sender := allGlobalValues[0].BCNAccountUUID1
	reciever := allGlobalValues[1].BCNAccountUUID1
	transactionID := uuid.NewString()
	TransactionTable = getRowValuesForTransaction(
		transactionID,
		sender,
		reciever,
		false)
	feeCustomerTransactionId := uuid.NewString()
	TransactionTable = fmt.Sprintf("%s\n%s",
		TransactionTable,
		getRowValuesForTransaction(
			feeCustomerTransactionId,
			sender,
			utils.CustomerFeeBcnAccountId,
			false,
		))
	SendMoneyWithInBCNTable = getRowValuesForSendMoneyWithInBCN(
		sender,
		reciever,
		transactionID,
		feeCustomerTransactionId,
	)
	for i := 2; i < 100; i = i + 2 {
		transactionID := uuid.NewString()
		sender := allGlobalValues[i].BCNAccountUUID1
		reciever := allGlobalValues[i+1].BCNAccountUUID1
		feeCustomerTransactionId := uuid.NewString()
		TransactionTable = fmt.Sprintf("%s\n%s",
			TransactionTable,
			getRowValuesForTransaction(
				transactionID,
				sender,
				reciever,
				false,
			))
		TransactionTable = fmt.Sprintf("%s\n%s",
			TransactionTable,
			getRowValuesForTransaction(
				feeCustomerTransactionId,
				sender,
				utils.CustomerFeeBcnAccountId,
				false,
			))
		SendMoneyWithInBCNTable = fmt.Sprintf("%s\n%s",
			SendMoneyWithInBCNTable,
			getRowValuesForSendMoneyWithInBCN(
				sender,
				reciever,
				transactionID,
				feeCustomerTransactionId,
			))

	}
}
