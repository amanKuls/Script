package generaterecords

import (
	"fmt"
	"script/utils"

	"github.com/google/uuid"
)

func getRowValuesForPassword(BCNUserUUID string) string {
	uniqueUUID := uuid.NewString()
	return fmt.Sprintf("%s,%s,%s,%s\n",
		uniqueUUID,
		"����",
		BCNUserUUID,
		utils.Boolean,
	)

}
