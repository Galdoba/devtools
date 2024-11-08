package pathfinder

import "testing"

func TestNewPath(t *testing.T) {
	app := "directory-tracker-service"
	type args struct {
		opts []StdPathOption
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "test 1",
			args:    args{[]StdPathOption{WithProgram(app), WithSystem("edts"), WithFileName("dir01.list")}},
			want:    "",
			wantErr: false,
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPath(tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
