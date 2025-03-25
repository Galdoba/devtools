package inject

import (
	"fmt"
	"testing"

	"github.com/Galdoba/devtools/version"
)

func Test_inject(t *testing.T) {

	v, err := version.Load(`c:\Users\pemaltynov\go\src\github.com\Galdoba\devtools\app\gvc\version.gvc`)
	if err != nil {
		t.Errorf("%v", err)
	}
	line := `	app.Version = "v 0.1.0" //#gvc: version control token`
	fmt.Println(line)
	inj, err := inject(v, line)
	fmt.Println(err)
	fmt.Println(inj)
}
