package gomake

import (
	"context"
	"os"
	"strings"
	"time"
)

const (
	gomake_env_variable = "GOMAKE_ENCRYPT"
)

var (
	// TempDir is a common temporary directory that is used throughout
	// this codebase.
	// It is the responsibility of the user to ensure that it is deleted
	// upon exiting.
	TempDir string = ""

	config *FileConfig = NewConfig()
)

func init() {
	config = NewConfig()

	config.encryptMode = DefaultEncryptMode

	if CheckCLI("encrypt") {
		config.encryptMode = true
	}

	if v, ok := CheckENV(gomake_env_variable); ok && v != "" {
		config.encryptMode = true
	}

	if DefaultEncryptMode {
		ctxParent = ctxEncrypt
		setupEncryption()
	} else {
		ctxParent = ctxDefault
	}

	TempDir = CreateTempDir("")
}

// FileConfig is used to configure timeouts for temporary file operations.
type FileConfig struct {
	encryptMode  bool
	parent       context.Context
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func NewConfig() *FileConfig {
	return &FileConfig{
		encryptMode:  DefaultEncryptMode,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
		IdleTimeout:  defaultIdleTimeout,
		parent:       ctxTemp,
	}
}

// CheckENV returns the value of the environment variable specified.
func CheckENV(key string) (string, bool) {
	k, ok := os.LookupEnv(strings.ToUpper(key))
	if ok {
		return k, true
	}
	return "", false
}
