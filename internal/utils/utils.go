package utils

import (
	"time"
)

// S repeats a space n times
func S(n int) string { _ = "STUB: not implemented"; return "" }

// TimeNowFunc is a utility function to simply testing
// by allowing TimeNowFunc to be defined to zero time
// to remove the time domain from tests
var TimeNowFunc = func() time.Time {
	return time.Now()
}
