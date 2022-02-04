package generaterecords

import (
	"fmt"
	"script/utils"

	"github.com/google/uuid"
)

// "id", "amount", "senderBcnAccountId", "recipientBcnAccountId", "claimedFeeCustomer",
// "confirmationExpiresAt", "status", "transactionId",
// "feeCustomerTransactionId", "exchangeRateId"
func getRowValuesForSendMoneyAccountToAccount(
	senderBCNAccountId string,
	recipientBCNAccountId string,
	transactionId string,
	feeCustomerTransactionId string,
) string {
	sendMoneyAccountToAccountUUID := uuid.NewString()
	exchangeRateId := "3e8b871a-f6f6-4bfc-852c-35b553d8a1b1"
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s",
		sendMoneyAccountToAccountUUID,
		utils.Amount,
		senderBCNAccountId,
		recipientBCNAccountId,
		utils.ClaimedFeeCustomer,
		utils.Timestamp,
		utils.Status,
		transactionId,
		feeCustomerTransactionId,
		exchangeRateId,
	)

}

func generateSendMoneyAccountToAccount() {
	sender := allGlobalValues[0].BCNAccountUUID1
	reciever := allGlobalValues[0].BCNAccountUUID2
	transactionID := uuid.NewString()
	TransactionTable = fmt.Sprintf("%s\n%s",
		TransactionTable,
		getRowValuesForTransaction(
			transactionID,
			sender,
			reciever,
			true,
		))
	feeCustomerTransactionId := uuid.NewString()
	TransactionTable = fmt.Sprintf("%s\n%s",
		TransactionTable,
		getRowValuesForTransaction(
			feeCustomerTransactionId,
			sender,
			utils.CustomerFeeBcnAccountId,
			false,
		))
	SendMoneyAccountToAccountTable = getRowValuesForSendMoneyAccountToAccount(
		sender,
		reciever,
		transactionID,
		feeCustomerTransactionId,
	)
	for i := 2; i < 100; i = i + 2 {
		transactionID := uuid.NewString()
		sender := allGlobalValues[i].BCNAccountUUID1
		reciever := allGlobalValues[i+1].BCNAccountUUID2
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
		SendMoneyAccountToAccountTable = fmt.Sprintf("%s\n%s",
			SendMoneyAccountToAccountTable,
			getRowValuesForSendMoneyAccountToAccount(
				sender,
				reciever,
				transactionID,
				feeCustomerTransactionId,
			))

	}
}
