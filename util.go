package gomake

import (
	"os"

	"github.com/skeptycal/gofile"
)

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
