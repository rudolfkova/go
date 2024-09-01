package sleep

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	<-time.After(duration)
}