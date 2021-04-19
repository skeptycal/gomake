// Package gomake is a repo management tool for Go projects.
package gomake

import (
	"errors"
	"os"
)

var (
	ErrBadCount   = errors.New("datafile: bad read count")
	ErrNotRegular = errors.New("datafile: not regular file")
)

type PathError = os.PathError

// func replace(filename, old, new string) (string, error) {
/*
	// 	f, err := New(filename)
	// 	if err != nil {
	// 		return "", err
	// 	}

	// 	buf, err := os.ReadFile(filename)
	// 	if err != nil {
	// 		return "", err
	// 	}

	// 	str := strings.ReplaceAll(string(buf), old, new)

	// 	buf = []byte(str)

	// 	os.Rename(filename, filebak)
	// }

	// ReplaceVar will replace instances of variable with repl.
	// If no value is provided for repl, the value from the environment
	// variable will be used. If that fails, an error will be returned.
	//
	// The variable should appear in filename in POSIX script notation,
	// either with a leading $ or, for more complicated scripts,
	// with ${} or $()
	//
	// Example line form a bash script bash script:
	//  echo "The name of the repo is ${REPO_NAME}.""
	// You could verify the variable at the command prompt:
	//  $ env | grep REPO_NAME
	//  REPO_NAME=stub
	// The go code that replaces this variable:
	//  ReplaceVar("myfile.sh", "REPO_NAME")
	// The example line would become:
	//  echo "The name of the repo is stub.""
	//
	// func ReplaceVar(filename, variable, new string) error {

	// 	if new == "" {
	// 		getvar, ok := os.LookupEnv(variable)
	// 		if !ok {
	// 			return fmt.Errorf("environment variable not set: %v", variable)
	// 		}
	// 		new = getvar
	// 	}

	// 	buf, err := os.ReadFile(filename)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	str := string(buf)

	// 	old := fmt.Sprintf("$%s", variable)
	// 	str = strings.ReplaceAll(str, old, new)
	// 	old = fmt.Sprintf("${%s}", variable)
	// 	str = strings.ReplaceAll(str, old, new)
	// 	old = fmt.Sprintf("$(%s)", variable)
	// 	str = strings.ReplaceAll(str, old, new)

	// 	err = os.WriteFile(filename, []byte(str), modemine)
	// 	if err != nil {
	// 		os.Rename()
	// 		return err
	// 	}

	// }
*/
