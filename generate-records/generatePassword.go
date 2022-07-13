package generaterecords

import (
	"fmt"
	"script/utils"

	"github.com/google/uuid"
)

// password is bou01Test1234
func getRowValuesForPassword(BCNUserUUID string) string {
	uniqueUUID := uuid.NewString()
	return fmt.Sprintf("%s,%s,%s,%s\n",
		uniqueUUID,
		"����",
		BCNUserUUID,
		utils.Boolean,
	)

}
