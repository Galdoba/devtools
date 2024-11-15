package pathfinder

import "testing"

func TestNewPath(t *testing.T) {

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
			args:    args{[]StdPathOption{IsConfig(), WithProgram("aue"), WithFileName("test.config")}},
			want:    `C:\Users\pemaltynov\.config\galdoba\aue\test.config`,
			wantErr: false,
		}, // TODO: Add test cases.
		{
			name:    "test 2",
			args:    args{[]StdPathOption{IsConfig(), WithProgram("aue"), WithFileName("test_bad.config"), EnsureExistiance()}},
			want:    `C:\Users\pemaltynov\.config\galdoba\aue\test_bad.config`,
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
