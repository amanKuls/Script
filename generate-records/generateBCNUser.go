package generaterecords

import (
	"fmt"
	"script/utils"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
)

func getRowValuesForBCNUser(BCNUserUUID string) string {
	uniqueUUID := uuid.NewString()
	uniquePhoneNumber := utils.GetUniquePhoneNumber()
	return fmt.Sprintf("%s,%s,%s,%s,+91%d,%s,true,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n",
		BCNUserUUID,
		faker.FirstName(),
		faker.LastName(),
		utils.RandomString,
		uniquePhoneNumber,
		faker.Email(),
		utils.DateType,
		utils.Gender,
		utils.Locale,
		utils.RandomString,
		uniqueUUID,
		utils.Boolean,
		utils.Boolean,
		BCNUserUUID,
		utils.Timestamp,
		utils.Timestamp,
		utils.CountryCode,
		"25",
	)

}
