package generaterecords

import (
	"fmt"
	"script/utils"
)

func getRowValuesForDefaulsBCNAccount(BCNUserUUID string, BCNAccountUUID string) string {
	return fmt.Sprintf("%s,%s,%s\n",
		BCNUserUUID,
		utils.CurrencyCode1,
		BCNAccountUUID,
	)

}
