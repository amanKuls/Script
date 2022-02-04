package generaterecords

import (
	"fmt"
	"script/utils"

	"github.com/google/uuid"
)

func getRowValuesForBCNUser(BCNUserUUID string) string {
	uniqueUUID := uuid.NewString()
	uniquePhoneNumber := utils.GetUniquePhoneNumber()
	return fmt.Sprintf("%s,%s,%s,%s,+91%d,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n",
		BCNUserUUID,
		utils.RandomString,
		utils.RandomString,
		utils.RandomString,
		uniquePhoneNumber,
		utils.EmialId,
		utils.DateType,
		utils.Gender,
		utils.Locale,
		utils.RandomString,
		uniqueUUID,
		utils.Boolean,
		utils.Boolean,
		BCNUserUUID,
		BCNUserUUID,
		utils.Timestamp,
		utils.Timestamp,
		utils.CurrencyCode1,
		"25",
	)

}
