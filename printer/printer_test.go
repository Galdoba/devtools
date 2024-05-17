package printer

import (
	"testing"

	"github.com/Galdoba/devtools/printer/lvl"
)

func TestLog(t *testing.T) {
	pm := New().
		WithConsoleColors(true)

	pm.Printf(lvl.REPORT, "The answer is %v", 42)
}
