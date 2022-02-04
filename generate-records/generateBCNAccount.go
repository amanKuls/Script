package generaterecords

import (
	"fmt"
	"script/utils"
)

func getRowValuesForBCNAccount(BCNAccountUUID string, BCNUserUUID string, currencyCode string) string {
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s",
		BCNAccountUUID,
		utils.RandomString,
		utils.Balance,
		currencyCode,
		BCNUserUUID,
		utils.Boolean,
		utils.Boolean,
		BCNUserUUID,
		BCNUserUUID,
		utils.Timestamp,
		utils.Timestamp,
	)

}
