package utils

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

var CurrentTime = time.Now().String()
var UUIDValue = uuid.NewString()
var RandomString = "Testing"
var ByteValue = string(make([]byte, 3, 4))
var DateType = "1994-05-11"
var Timestamp = "2022-02-03 14:00:00 +02:00"
var LLTExpire = "2023-02-10 14:00:00 +02:00"
var Boolean = "false"
var Amount = "1000000"
var Balance = "500000"
var Gender = "Female"
var PhoneNumber = rand.Intn(9999999999-1000000000) + 1000000000
var Locale = "EN_US"
var CurrencyCode1 = "MWK"
var CurrencyCode2 = "ZMW"
var CountryCode = "MW"
var ClaimedFeeCustomer = "50000"
var Status = "SUCCEEDED"
var CustomerFeeBcnAccountId = "4373074b-0fe3-48b2-abbe-b0bb5f7a0a18"
var MPGSHoldingBcnAcccountId = "b79821f8-4d93-43fe-9fd0-ec795d403126"
var CounterpartyId = "SENDMONEYEXERNALTOBCN_FDHBANK"
var ClaimedFeeCounterparty = "35000"
var TransactionResultedAt = string(time.Now().AddDate(0, -1, 0).UTC().Format(time.RFC3339))

func GetRandomNumberBetweenRangeExcept(min int, max int, Exception int) int {
	result := rand.Intn(max-min) + min
	if result == Exception {
		result = (Exception + max) / 2
	}
	return result

}

func GetRandomNumberBetweenRange(min int, max int) int {
	result := rand.Intn(max-min) + min
	return result

}

func GetUniquePhoneNumber() int {
	return rand.Intn(9999999999-1000000000) + 1000000000
}
