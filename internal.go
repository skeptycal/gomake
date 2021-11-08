package gomake

import (
	"os"

	"github.com/skeptycal/gofile"
)

type FileMode = os.FileMode

const (
	normalMode        FileMode = 0644
	dirMode           FileMode = 0755
	defaultBufferSize int      = 1024
	minBufferSize     int64    = 16
)

func readBak(filename string) ([]byte, error) {
	_, err := gofile.Copy(filename, filename+".bak")
	if err != nil {
		return nil, err
	}

	return os.ReadFile(filename)
}
