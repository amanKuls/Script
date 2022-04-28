package generaterecords

import (
	"fmt"
	"script/utils"

	"github.com/google/uuid"
)

// "id", "amount", "senderBcnAccountId", "recipientBcnAccountId", "claimedFeeCustomer",
// "confirmationExpiresAt", "status", "transactionId",
// "feeCustomerTransactionId"
func getRowValuesForSendMoneyAccountToAccount(
	senderBCNAccountId string,
	recipientBCNAccountId string,
	transactionId string,
	feeCustomerTransactionId string,
) string {
	sendMoneyAccountToAccountUUID := uuid.NewString()
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s\n",
		sendMoneyAccountToAccountUUID,
		utils.Amount,
		senderBCNAccountId,
		recipientBCNAccountId,
		utils.ClaimedFeeCustomer,
		utils.Timestamp,
		utils.Status,
		transactionId,
		feeCustomerTransactionId,
	)

}

func generateSendMoneyAccountToAccount(start int, end int, numberOfTransactions int) {
	for i := 1; i <= numberOfTransactions; i++ {
		userNumber := utils.GetRandomNumberBetweenRange(start, end)
		transactionID := uuid.NewString()
		sender := allUsers[userNumber].BCNAccountUUID1
		reciever := allUsers[userNumber].BCNAccountUUID2
		feeCustomerTransactionId := uuid.NewString()
		TransactionTable = append(
			TransactionTable,
			getRowValuesForTransaction(
				transactionID,
				sender,
				reciever,
				true,
				userNumber,
				-1,
			))
		TransactionTable = append(
			TransactionTable,
			getRowValuesForTransaction(
				feeCustomerTransactionId,
				sender,
				utils.CustomerFeeBcnAccountId,
				false,
				userNumber,
				-1,
			))
		SendMoneyAccountToAccountTable = append(
			SendMoneyAccountToAccountTable,
			getRowValuesForSendMoneyAccountToAccount(
				sender,
				reciever,
				transactionID,
				feeCustomerTransactionId,
			))

	}
}
