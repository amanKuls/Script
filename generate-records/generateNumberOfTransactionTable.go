package generaterecords

import (
	"fmt"
)

func getRowValueForNumberOfTransactionTable(userId string, llt string, numOfTransaction int) string {
	return fmt.Sprintf("%s,%s,%d\n",
		userId,
		llt,
		numOfTransaction,
	)
}

func generateNumberOfTransactionTable(numberOfUsers int) {
	for i := 0; i < numberOfUsers; i++ {
		NumberOfTransactionTable = append(
			NumberOfTransactionTable,
			getRowValueForNumberOfTransactionTable(
				allUsers[i].BCNUserUUID,
				allUsers[i].LeoLLTLLT,
				allUsers[i].NoOfTransactions,
			),
		)
	}
}
