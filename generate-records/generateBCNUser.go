package generaterecords

import (
	"fmt"
	"script/utils"

	"github.com/google/uuid"
)

func getRowValuesForBCNUser(BCNUserUUID string) string {
	uniqueUUID := uuid.NewString()
	uniquePhoneNumber := utils.GetUniquePhoneNumber()
	return fmt.Sprintf("%s,+91%d,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n",
		BCNUserUUID,
		uniquePhoneNumber,
		utils.Locale,
		utils.RandomString,
		uniqueUUID,
		"true",
		"false",
		BCNUserUUID,
		utils.Timestamp,
		utils.Timestamp,
		utils.CountryCode,
		"25",
	)

}
