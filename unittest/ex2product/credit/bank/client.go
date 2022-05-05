package bank

import "time"

func IsTooBusy() bool {
	return time.Now().Minute()%2 == 0
}
