package validation

import (
	"fmt"
	"time"
)

func IsEmpty(check string) bool {
	return check == "0" || check == "" || check == fmt.Sprint(time.Time{})
}
