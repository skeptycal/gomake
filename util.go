package gomake

import (
	"os"

	"github.com/skeptycal/gofile"
)

func MkDir(name string) error {
	return os.Mkdir(name, normalMode)
}

// checkDir checks whether the directory exists and returns
// an error. If the directory does not exist, it is created.
// Any errors are of type *PathError
func checkDir(repoName string) error {

	if gofile.IsDir(repoName) {
		return &os.PathError{
			Op:   "unable to create directory",
			Path: repoName,
			Err:  os.ErrExist,
		}
	} else {
		err := os.Mkdir(repoName, 0755)
		if err != nil {
			return err
		}
		err = os.Chdir(repoName)
		if err != nil {
			return err
		}
	}
	return nil
}

// New creates a new Git repository and GitHub repository for
// a new Go project.
//
// If the name is not given, the parent folder name is used.
func New(repoName string) error {

	// todo - check for flags

	// check for existing directory
	if repoName == "" {
		repoName = gofile.PWD()
	} else {
		err := checkDir(repoName)
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
