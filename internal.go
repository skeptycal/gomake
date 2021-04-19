package gomake

import (
	"os"
)

var NewPathError = gofile.NewPathError

func readBak(filename string) ([]byte, error) {
	_, err := gofile.Copy(filename, filename+".bak")
	if err != nil {
		return nil, err
	}

	return os.ReadFile(filename)
}
