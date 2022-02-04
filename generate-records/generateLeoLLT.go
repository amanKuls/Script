package generaterecords

import (
	"fmt"
	"script/utils"
)

func getRowValuesForLeoLLT(LeoLLTLLT string) string {
	return fmt.Sprintf("%s,%s,%s,%s,%s\n",
		LeoLLTLLT,
		utils.Timestamp,
		utils.Timestamp,
		utils.LLTExpire,
		utils.Boolean)
}
