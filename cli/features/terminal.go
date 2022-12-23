package features

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
)

func TypingSlowly(msg string, msDelay time.Duration) {
	for _, val := range msg {
		//r := rune(msg[i])
		fmt.Print(string(val))
		//time.Sleep(time.Millisecond * msDelay)
		time.Sleep(msDelay / 1000)
	}
}

func Size() (int, int) {
	return termbox.Size()
}
