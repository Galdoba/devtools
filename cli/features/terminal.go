package features

import (
	"fmt"
	"time"
)

func TypingSlowly(msg string, msDelay time.Duration) {
	for i := range msg {
		fmt.Print(string(msg[i]))
		time.Sleep(time.Millisecond * msDelay)
	}
}
