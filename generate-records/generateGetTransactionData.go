package generaterecords

import (
	"fmt"
)

func getRowValuesGetTransaction(LeoSLTSLT string, BCNAccountId string) string {
	return fmt.Sprintf("%s,%s\n",
		LeoSLTSLT,
		BCNAccountId,
	)
}
