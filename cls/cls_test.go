package cls

import (
	"log"
	"testing"
)

func TestCLS(t *testing.T) {
	cls := New()

	cls.Info("aaabbb")
	cls.AddFile(`c:\Users\pemaltynov\go\src\github.com\Galdoba\devtools\gslog\testLog.txt`, LV_DEBUG, "app in file ", log.Lshortfile|log.Ltime|log.Lmsgprefix|log.Lmicroseconds)

	cls.Info("DDDCCC")
	cls.Error("999")
}
