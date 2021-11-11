package gomake

import (
	"errors"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/gofile"
)

const templatePath string = "template_files"

var (
	// ErrNoTemplatePath is returned when no template path is found.
	ErrNoTemplatePath = errors.New("template directory not found")

	// ErrNoTemplatePath is returned when the template file is not found.
	ErrNoTemplate = errors.New("template file not found")

	// TemplatesAvailable is a global flag for template use.
	TemplatesAvailable bool = false
)

// PathError returns an error of type os.PathError.
//  type PathError struct {
//     Op   string
//     Path string
//     Err  error
//  }
func PathError(Op, path string, err error) error {
	return &os.PathError{Op: Op, Path: path, Err: err}
}

// ReadTemplate returns the contents of a template file.
func ReadTemplate(name string) (string, error) {

	if !TemplatesAvailable {
		return "", ErrNoTemplatePath
	}

	templateFileName := filepath.Join(templatePath, name)

	b, err := os.ReadFile(templateFileName)
	if err != nil {
		return "", PathError("ReadTemplate#os.Readfile", name, err)
	}

	return string(b), nil
}

func init() {
	if !gofile.IsDir(templatePath) {
		log.Info(PathError("templates|init#gofile.IsDir", templatePath, ErrNoTemplate))
		TemplatesAvailable = false
	} else {
		TemplatesAvailable = true
	}
}
