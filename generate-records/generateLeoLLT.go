package generaterecords

import (
	"fmt"
	"script/utils"
)

func getRowValuesForLeoLLT(LeoLLTLLT string) string {
	return fmt.Sprintf("%s,%s,%s,%s,%s",
		LeoLLTLLT,
		utils.Timestamp,
		utils.Timestamp,
		utils.Timestamp,
		utils.Boolean)
}
