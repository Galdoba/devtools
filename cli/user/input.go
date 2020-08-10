package user

import (
	"bufio"
	"os"
	"runtime"
	"strings"
)

func InputStr() string {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		return err.Error()
	}
	switch runtime.GOOS {
	default:
		return "WARNING: Unknown System"
	case "windows":
		str = strings.TrimSuffix(str, "\n")
		str = strings.TrimSuffix(str, "\r")
	case "linux":
		str = strings.TrimSuffix(str, "\n")
	}
	return str
}
