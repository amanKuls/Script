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
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n",
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

func generateSendMoneyAccountToAccount(start int, end int) {
	for i := start; i <= end; i++ {
		transactionID := uuid.NewString()
		sender := allUsers[i].BCNAccountUUID1
		reciever := allUsers[i+1].BCNAccountUUID2
		feeCustomerTransactionId := uuid.NewString()
		TransactionTable = append(
			TransactionTable,
			getRowValuesForTransaction(
				transactionID,
				sender,
				reciever,
				i,
				-1,
			))
		TransactionTable = append(
			TransactionTable,
			getRowValuesForTransaction(
				feeCustomerTransactionId,
				sender,
				utils.CustomerFeeBcnAccountId,
				i,
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
