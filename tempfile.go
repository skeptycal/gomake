package gomake

import (
	"context"
	"os"
	"path"
	"time"

	"github.com/skeptycal/gofile"
)

const (
	defaultReadTimeout  time.Duration = 1 * time.Second
	defaultWriteTimeout time.Duration = 1 * time.Second
	defaultIdleTimeout  time.Duration = 10 * time.Minute
)

var (
	// TempDir is a common temporary directory that is used throughout
	// this codebase.
	// It is the responsibility of the user to ensure that it is deleted
	// upon exiting.
	TempDir string = ""
)

func init() {
	config = NewConfig(ctxTemp)
	TempDir = CreateTempDir("")
}

func GetFileCtx() (ctx context.Context, cancelfunc context.CancelFunc) {
	return context.WithTimeout(ctxEncrypt, config.IdleTimeout)
}

func NewTempFile(ctx context.Context, filename string, encryptMode bool) (file *EncryptedFile, err error) {
	return newTempFile(ctx, filename)
}

func newTempFile(ctx context.Context, filename string) (file *EncryptedFile, err error) {
	return &EncryptedFile{
		filename:  filename,
		encrypted: config.encryptMode,
		timeout:   defaultIdleTimeout,
	}, nil
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
func CreateTempDir(altName string) string {

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
