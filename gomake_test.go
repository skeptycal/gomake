package gomake

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/skeptycal/gofile"
)

func TestNewPathError(t *testing.T) {
	type args struct {
		op   string
		path string
		err  error
	}
	tests := []struct {
		name string
		args args
		want *PathError
	}{
		// TODO: Add test cases.
		{"bad abs path", args{"could not determine absolute path", ".", filepath.ErrBadPattern}, &PathError{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gofile.NewPathError(tt.args.op, tt.args.path, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPathError() = %v, want %v", got, tt.want)
			}
		})
	}
}
