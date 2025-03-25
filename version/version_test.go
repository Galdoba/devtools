package version

import (
	"fmt"
	"testing"
)

func TestVersion_String(t *testing.T) {
	v := New("path")
	v.Update()
	v.Update()
	v.Update()
	v.UpgradeMinor()
	v.Update()
	v.UpgradeMajor("d", "fc")
	for i := 0; i < 26; i++ {
		v.Update("dev", "rc1")
	}
	v.Patch()
	v.UpgradeMinor()
	v.Update()
	fmt.Println(v.String())
}
