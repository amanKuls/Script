package generaterecords

import (
	"fmt"
)

func getRowValuesForBCNUserLLT(BCNUserUUID string, LeoLLTLLT string) string {
	return fmt.Sprintf("%s,%s\n",
		BCNUserUUID,
		LeoLLTLLT,
	)
}
