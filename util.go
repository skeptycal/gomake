package gomake

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/skeptycal/gofile"
)

const (
	goMakeTempDir = "gomake_temp"
	sep           = string(os.PathSeparator)
)

type (
	Any interface{}
	any struct{}
)

var ctxTemp context.Context = context.Background()

// CheckCLI checks the CLI options for a particular parameter.
//
// TODO - not implemented yet
// (returns true by default; use command = "false" to test false result)
func CheckCLI(command string) bool {
	// TODO - not implemented yet
	if command == "false" {
		return false
	}
	return true
}

func StatCheck(filename string) (os.FileInfo, error) {

	// Validate filename ...
	// EvalSymlinks returns the path name after the evaluation of any symbolic
	// links.
	// If path is relative the result will be relative to the current directory,
	// unless one of the components is an absolute symbolic link.
	// EvalSymlinks calls Clean on the result.
	filename, err := filepath.EvalSymlinks(filename)
	if err != nil {
		return nil, err
	}

	filename, err = filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}

	//Check 'others' permission
	m := fi.Mode()
	if m&(1<<2) == 0 {
		return nil, fmt.Errorf("insufficient permissions: %v", filename)
	}

	if fi.IsDir() {
		return nil, fmt.Errorf("the filename %s refers to a directory", filename)
	}

	if !fi.Mode().IsRegular() {
		return nil, fmt.Errorf("the filename %s is not a regular file", filename)
	}

	return fi, err
}

// MkDir creates the directory dir if it does not exist
// and changes the current working directory to dir.
// Any errors are of type *PathError
func MkDir(dir string) error {

	if !gofile.IsDir(dir) {
		if err := os.Mkdir(dir, dirMode); err != nil {
			return err
		}
	}
	return os.Chdir(dir)
}

// New creates a new Git repository and GitHub repository for
// a new Go project.
//
// If the name is not given, the parent folder name is used.
func New(repoName string) error {

	// todo - check for CLI flags

	// check for existing directory
	if repoName == "" {
		repoName = gofile.PWD()
	} else {
		err := MkDir(repoName)
		if err != nil {
			return err
		}
	}

	// check for existing git repo

	// gather config data

	// create directory structure

	// create config file

	// create repo go file

	// create .gitignore

	return nil
}
