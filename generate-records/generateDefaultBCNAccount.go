package generaterecords

import (
	"fmt"
	"script/utils"
)

func getRowValuesForDefaulsBCNAccount(BCNUserUUID string, BCNAccountUUID string) string {
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s\n",
		BCNUserUUID,
		utils.CurrencyCode1,
		BCNAccountUUID,
		BCNUserUUID,
		BCNUserUUID,
		utils.Timestamp,
		utils.Timestamp,
	)

}
