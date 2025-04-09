package version

import (
	"fmt"
	"testing"
)

func TestVersion_String(t *testing.T) {
	// v := New("path", WithName("test_project"), WithDescription("this is looooong descr"))
	// v.Update()
	// v.Update()
	// v.Update()
	// v.UpgradeMinor()
	// v.Update()
	// v.UpgradeMajor("release")
	// for i := 0; i < 26; i++ {
	// 	v.Update()
	// }
	// v.Patch()
	// v.UpgradeMinor("better ui")
	// v.Update("fix err 1", "fix err 2")
	// v.Patch("final release")
	// v.UpgradeMajor("production build")
	// fmt.Println(v.String())
	// bt, err := json.MarshalIndent(v, "", "  ")
	// fmt.Println(err)
	// fmt.Println(string(bt))
	// fmt.Println(v.Text())
	// fmt.Println(v.String())
	// fmt.Println(v.Save())
	v, err := Load("test_project")
	fmt.Println(err)
	fmt.Println(v.String())

}
