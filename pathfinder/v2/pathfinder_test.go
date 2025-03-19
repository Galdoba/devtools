package v2

import (
	"testing"
)

func Test_constructPath(t *testing.T) {
	type args struct {
		pf pathfinder
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "simple 1",
			args: args{pathfinder{
				root:              rootHOME,
				appName:           "test",
				template:          TEMPLATE_PROGRAM,
				systemlayers:      []string{"Programs", "sys1"},
				prefixlayers:      []string{"pref1"},
				suffixlayers:      []string{"suff1"},
				fileName:          "fn",
				ext:               "ext",
				perm:              0666,
				isDir:             false,
				path:              "",
				mustHaveAppName:   false,
				mustHaveSignature: true,
				mustHaveExt:       false,
				mustExistDir:      false,
				mustExistFile:     false,
				skipValidation:    false,
			}},
			want: `C:\Users\pemaltynov\Programs\sys1\galdoba\pref1\test\suff1\fn.ext`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructPath(tt.args.pf); got != tt.want {
				t.Errorf("constructPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	// pf, _ := New(TEMPLATE_DOCUMENT, WithFileName("name"), WithExtention("ext"), NoValidation)
	// fmt.Println(pf)
	// panic(0)
	type args struct {
		template  string
		pfOptions []Option
	}
	tests := []struct {
		name    string
		args    args
		want    *pathfinder
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				template: TEMPLATE_DOCUMENT,
				pfOptions: []Option{
					WithFileName("name"),
					WithExtention("ext"),
					WithAppName("app"),
					WithPrefixLayers("prf1", "prf2"),
					WithSuffixLayers("sf"),
				},
			},
			want: &pathfinder{
				root:              rootHOME,
				appName:           "app",
				template:          TEMPLATE_DOCUMENT,
				systemlayers:      []string{"Documents"},
				prefixlayers:      []string{"prf1", "prf2"},
				suffixlayers:      []string{"sf"},
				fileName:          "name",
				ext:               "ext",
				perm:              0666,
				isDir:             false,
				path:              `C:\Users\pemaltynov\Documents\prf1\prf2\app\sf\name.ext`,
				mustHaveAppName:   false,
				mustHaveSignature: false,
				mustHaveExt:       true,
				mustExistDir:      false,
				mustExistFile:     false,
				skipValidation:    false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.template, tt.args.pfOptions...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("New() = %v, want %v", got, tt.want)
			// }
			if got.root != tt.want.root {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
			if got.appName != tt.want.appName {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
			if got.template != tt.want.template {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
			for i := range got.systemlayers {
				if got.systemlayers[i] != tt.want.systemlayers[i] {
					t.Errorf("New() = %v, want %v", got, tt.want)
				}
			}
			for i := range got.prefixlayers {
				if got.prefixlayers[i] != tt.want.prefixlayers[i] {
					t.Errorf("New() = %v, want %v", got, tt.want)
				}
			}
			for i := range got.suffixlayers {
				if got.suffixlayers[i] != tt.want.suffixlayers[i] {
					t.Errorf("New() = %v, want %v", got, tt.want)
				}
			}
			if got.fileName != tt.want.fileName {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
			if got.ext != tt.want.ext {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
			if got.perm != tt.want.perm {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
			if got.isDir != tt.want.isDir {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
			if got.path != tt.want.path {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
			if got.mustHaveAppName != tt.want.mustHaveAppName {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
			if got.mustHaveSignature != tt.want.mustHaveSignature {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
			if got.mustHaveExt != tt.want.mustHaveExt {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
			if got.mustExistDir != tt.want.mustExistDir {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
			if got.mustExistFile != tt.want.mustExistFile {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
			if got.skipValidation != tt.want.skipValidation {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}

		})
	}
}
