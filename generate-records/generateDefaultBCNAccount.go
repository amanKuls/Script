package generaterecords

import (
	"fmt"
	"script/utils"
)

func getRowValuesForDefaulsBCNAccount(BCNUserUUID string, BCNAccountUUID string) string {
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s",
		BCNUserUUID,
		utils.CurrencyCode,
		BCNAccountUUID,
		BCNUserUUID,
		BCNUserUUID,
		utils.Timestamp,
		utils.Timestamp,
	)

}
