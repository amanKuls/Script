package generaterecords

import (
	"fmt"
	"script/utils"

	"github.com/google/uuid"
)

//id, recordId, accountId, amount, currencyCode, txnType, titleLocalizedTextId,
// status, isCredit, allowAddingPrivateRemark, txnOccurredAt
func getRowValueForUserVisibleTransactionTable(
	featureTableRecordId string,
	bcnAccountId string,
	isFee bool,
) string {
	var txnType string
	if isFee {
		txnType = "TRANSACTION_FEE"
	} else {
		txnType = "TRANSACTION"
	}
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n",
		uuid.NewString(),
		featureTableRecordId,
		bcnAccountId,
		utils.Amount,
		utils.CurrencyCode1,
		txnType,
		"com.resoluttech.bcn.country.mpambaWallet",
		"SUCCEEDED",
		"true",  // isCredit
		"false", //allowAddingPrivateRemark,
		"2022-06-11 15:19:21.077 +0530",
	)
}
