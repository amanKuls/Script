package generaterecords

import (
	"fmt"
	"script/utils"
)

func getRowValuesForLeoSLT(LeoSLTSLT string, LeoSLTLLT string) string {
	return fmt.Sprintf("%s,%s,%s,%s\n",
		LeoSLTSLT,
		LeoSLTLLT,
		"2021-05-06 21:00:10.000 +0530",
		utils.LLTExpire,
	)
}
