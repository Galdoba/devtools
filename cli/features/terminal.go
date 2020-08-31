package features

import (
	"fmt"
	"time"
)

func TypingSlowly(msg string, msDelay time.Duration) {
	for _, val := range msg {
		//r := rune(msg[i])
		fmt.Print(string(val))
		//time.Sleep(time.Millisecond * msDelay)
		time.Sleep(msDelay)
	}
}
