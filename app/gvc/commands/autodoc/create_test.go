package autodoc

import (
	"fmt"
	"testing"

	"github.com/Galdoba/devtools/version"
)

func Test_projectRoot(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "simple", args: args{path: `c:\Users\pemaltynov\go\src\github.com\Galdoba\devtools\app\hello\version.gvc`}, want: `c:\Users\pemaltynov\go\src\github.com\Galdoba\devtools\`, wantErr: false},
		{name: "error", args: args{path: `c:\Users\pemaltynov\go\src\github.com\Galdoba\version.gvc`}, want: ``, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := projectRoot(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("projectRoot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("projectRoot() = %v, want %v", got, tt.want)
			}

		})
	}
}

func wantedErr(testName string) error {
	errMap := make(map[string]error)
	errMap["error"] = fmt.Errorf(`no project root was found for c:\Users\pemaltynov\go\src\github.com\Galdoba\version.gvc`)
	//////
	return errMap[testName]
}

func Test_docPath(t *testing.T) {
	v, err := version.Load(`c:\Users\pemaltynov\go\src\github.com\Galdoba\devtools\app\gvc\version.gvc`)
	fmt.Println(err)
	fmt.Println(v.ProjectName)
	fmt.Println("loaded")
	fmt.Println(v.ProjectName)
	path, err := docPath(v)
	fmt.Println(path, err)
}
