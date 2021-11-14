package gomake

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/skeptycal/errorlogger"
	"github.com/skeptycal/gofile"
)

const (
	goMakeTempDir = "gomake_temp"
	sep           = string(os.PathSeparator)
)

var (
	// Log is the default global ErrorLogger. It implements the ErrorLogger interface as well as the basic logrus.Logger interface, which is compatible with the standard library "log" package.
	//
	// In the case of name collisions with 'Log', use an alias instead of creating a new instance. For example:
	//
	//  var mylogthatwontmessthingsup = errorlogger.Log
	Log = errorlogger.Log

	// Err is the logging function for the global ErrorLogger.
	Err = errorlogger.Err
)

type (
	Any interface{}
	// any struct{}
)

// Noner is the interface that implements features of the
// None type, including fmt.Stringer to return the string "None".
//
// It is purely for testing purposes ... and a few laughs.
type Noner interface{}

// None is an empty struct{} that represents no value. It is
// meant to represent a value that has not been measured and is
// otherwise unknown.
//
// None implements fmt.Stringer to return the string "None"
type None struct{}

func (n None) String() string {
	return "None"
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

type repo struct {
	name string
}

type Repo interface {
	fmt.Stringer
}

func newRepo(name string) (Repo, error) {
	return nil, nil
}

// New creates a new Git repository and GitHub repository for
// a new Go project.
//
// If the name is not given, the parent folder name is used.
func New(repoName string) (Repo, error) {

	var r *repo
	// todo - check for CLI flags

	if gofile.IsDir(repoName) {
		return newRepo(repoName)
	}

	// check for existing directory
	if repoName == "" {
		repoName = gofile.PWD()
	} else {
		err := MkDir(repoName)
		if err != nil {
			return nil, fmt.Errorf("error creating directory %v", repoName)
		}
	}

	r.name = repoName

	// check for existing git repo

	// gather config data

	// create directory structure

	// create config file

	// create repo go file

	// create .gitignore

	return nil, nil
}
