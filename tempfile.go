package gomake

import (
	"context"
	"os"
	"path"

	"github.com/skeptycal/gofile"
)

func newTempFile(ctx context.Context, filename string) TempFile {
	if !config.encryptMode {
		return &file{filename: filename}
	}
	f := &file{filename: filename}
	return &EncryptedFile{f, config.encryptMode}
}

// CreateTempDir returns the path of a new temporary directory
// using the default system temporary directory from os.TempDir().
//
// On Unix systems, it returns $TMPDIR if non-empty, else /tmp. On Windows,
// it uses GetTempPath, returning the first non-empty value from %TMP%, %TEMP%,
// %USERPROFILE%, or the Windows directory. On Plan 9, it returns /tmp.
//
// The system directory is neither guaranteed to exist nor have accessible permissions.
//
// If the system default is unavailable for any reason, a directory is created
// in the current with a random name based on altName.
func CreateTempDir(filepath string) string {

	tmpdir := os.TempDir()

	// if default system TempDir() returns an invalid path...
	if !gofile.IsDir(tmpdir) {

		if filepath == "" {
			filepath = gofile.PWD()
		}

		tmpdir, err := os.MkdirTemp(filepath, goMakeTempDir)
		if err != nil {
			return gofile.PWD() + sep + goMakeTempDir
		}

		return tmpdir
	}

	tmpdir, err := os.MkdirTemp(tmpdir, goMakeTempDir+"*")
	if err != nil {
		return path.Join(gofile.PWD(), goMakeTempDir)
	}

	return ""
}

func MakeLocalTempDir(filepath string) string {
	// TODO - not implemented yet
	return ""
}
