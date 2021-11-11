package gomake

import (
	"reflect"
	"testing"
)

func Test_makeKey(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantKey []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{"8", args{8}, []byte{}, false},
		{"16", args{16}, []byte{}, false},
		{"32", args{32}, []byte{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, err := makeKey(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(gotKey) != tt.args.n {
				t.Errorf("makeKey() length = %v, want %v", len(gotKey), tt.args.n)
				return
			}
		})
	}
}

func Test_setupEncryption(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"setup encryption"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupEncryption()
		})
	}
}

func Test_encrypt(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		args    args
		wantBuf []byte
	}{
		// TODO: Add test cases.
		{"fake", args{[]byte("foo")}, []byte{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBuf := encrypt(tt.args.p); !reflect.DeepEqual(gotBuf, tt.wantBuf) {
				t.Errorf("encrypt() = %v, want %v", gotBuf, tt.wantBuf)
			}
		})
	}
}
