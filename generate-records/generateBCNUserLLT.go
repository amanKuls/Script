package generaterecords

import (
	"fmt"
)

// userId, llt, frontEndPlatform, applicationType, applicationVersion
func getRowValuesForBCNUserLLT(BCNUserUUID string, LeoLLTLLT string) string {
	return fmt.Sprintf("%s,%s,%s,%s,%s\n",
		BCNUserUUID,
		LeoLLTLLT,
		"ANDROID",
		"BCN",
		"1.0",
	)
}
