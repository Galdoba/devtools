package cellscript

import (
	"fmt"
	"strings"
	"testing"
)

func TestScript(t *testing.T) {

	testPathValid := `c:\Users\pemaltynov\go\bin\testscript.bat` //Subject to change
	sc := NewScript(`testscript.bat.invalid`, "arg 1", "arg 2")
	for _, initate := range []bool{false, true} {
		for _, ready := range []bool{false, true} {
			for _, inprogress := range []bool{false, true} {
				for _, completed := range []bool{false, true} {
					sc.proc.initiated = initate
					sc.proc.ready = ready
					sc.proc.inprogress = inprogress
					sc.proc.completed = completed
					status := sc.proc.Status()

					t.Logf("set (%v, %v, %v, %v) = status %v", initate, ready, inprogress, completed, status)
					//time.Sleep(time.Second)
				}
			}
		}
	}

	if err := sc.asses(); err != nil {
		switch {
		default:
			t.Errorf("unexpected error: %v", err.Error())
		case strings.Contains(err.Error(), "at this point"):
			t.Logf("expected error: %v", err.Error())
		case strings.Contains(err.Error(), "cannot find"):
			t.Logf("expected error: %v", err.Error())
		}
	}
	sc.path = testPathValid
	sc.proc = &process{}
	if err := sc.asses(); err != nil {
		switch {
		default:
			t.Logf("unexpected error: %v", err.Error())
		}

	}
	sc.proc.initiated = false
	sc.proc.ready = true
	if err := sc.asses(); err != nil {
		switch {
		default:
			t.Errorf("unexpected error: %v", err.Error())
		case strings.Contains(err.Error(), "must not be initiated or ready at this point"):
			t.Logf("expected error: %v", err.Error())
		}

	}

	sc.proc = &process{}
	if err := sc.asses(); err != nil {
		switch {
		default:
			t.Errorf("unexpected error: %v", err.Error())
		}

	}
	fmt.Println(sc)
	fmt.Println(sc.proc)
	stdout, stderr, err := sc.Run()
	fmt.Printf("SCRIPT OUT:%v\n", stdout)
	fmt.Printf("SCRIPT ERR:%v\n", stderr)
	t.Logf("SCRIPT OUT:%v", stdout)
	t.Logf("SCRIPT ERR:%v", stderr)
	t.Logf("err: %v", err)
	if err != nil {
		t.Errorf("unexpected error: %v", err.Error())
	}
}
