package cls

import (
	"log"
	"testing"
)

func TestCLS(t *testing.T) {
	cls := New()

	cls.AddStdErr(LV_INFO, "app", 0)
	cls.Info("aaa", "bbb")
	cls.AddFile(`c:\Users\pemaltynov\go\src\github.com\Galdoba\devtools\gslog\testLog.txt`, LV_DEBUG, "app in file ", log.Lshortfile|log.Ltime|log.Lmsgprefix|log.Lmicroseconds)

	cls.Info("DDD", "CCC")
	cls.Error("999")
	cls.Fprintf(LV_REPORT, "complete %v", 2)
}
