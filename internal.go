package gomake

import (
	"os"

	"github.com/skeptycal/gofile"
)

const (
	normalMode        os.FileMode = 0644
	dirMode           os.FileMode = 0755
	defaultBufferSize int         = 1024
	minBufferSize     int64       = 16
)

func readBak(filename string) ([]byte, error) {
	_, err := gofile.Copy(filename, filename+".bak")
	if err != nil {
		return nil, err
	}

	return os.ReadFile(filename)
}
