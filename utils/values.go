package utils

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

var CurrentTime = time.Now().String()
var UUIDValue = uuid.NewString()
var RandomString = "Testing123"
var ByteValue = string(make([]byte, 4))
var DateType = "1994-05-11"
var Timestamp = "2012-08-24 14:00:00 +02:00"
var Boolean = "false"
var Number = "100"
var EmialId = "testing@test.com"
var Gender = "Female"
var PhoneNumber = rand.Intn(9999999999-1000000000) + 1000000000
var Locale = "EN_US"
var CurrencyCode = "MWK"

func GetUniquePhoneNumber() int {
	return rand.Intn(9999999999-1000000000) + 1000000000
}
