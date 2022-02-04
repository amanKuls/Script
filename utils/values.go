package utils

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

var CurrentTime = time.Now().String()
var UUIDValue = uuid.NewString()
var RandomString = "Testing123"
var ByteValue = string(make([]byte, 3, 4))
var DateType = "1994-05-11"
var Timestamp = "2012-08-24 14:00:00 +02:00"
var Boolean = "false"
var Amount = "1000000"
var Balance = "500000"
var EmialId = "testing@test.com"
var Gender = "Female"
var PhoneNumber = rand.Intn(9999999999-1000000000) + 1000000000
var Locale = "EN_US"
var CurrencyCode1 = "MWK"
var CurrencyCode2 = "ZMW"
var ClaimedFeeCustomer = "50000"
var Status = "SUCCEEDED"
var CustomerFeeBcnAccountId = "4373074b-0fe3-48b2-abbe-b0bb5f7a0a18"

func GetUniquePhoneNumber() int {
	return rand.Intn(9999999999-1000000000) + 1000000000
}
