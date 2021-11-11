package gomake

import "testing"

func TestCheckCLI(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"COLORTERM=truecolor", args{"COLORTERM=truecolor"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckCLI(tt.args.command); got != tt.want {
				t.Errorf("CheckCLI() = %v, want %v", got, tt.want)
			}
		})
	}
}
