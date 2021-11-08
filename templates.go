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
	ErrNoTemplatePath       = errors.New("template directory not found")
	ErrNoTemplate           = errors.New("template file not found")
	TemplatesAvailable bool = false
)

type PathError = os.PathError

// ReadTemplate returns the contents of a template file.
func ReadTemplate(name string) (string, error) {

	if !TemplatesAvailable {
		return "", ErrNoTemplatePath
	}

	templateFileName := filepath.Join(templatePath, name)

	b, err := os.ReadFile(templateFileName)
	if err != nil {
		return "", &PathError{"ReadTemplate#os.Readfile", name, err}
	}

	return "", nil
}

func init() {
	if !gofile.IsDir(templatePath) {
		err := &PathError{"ReadTemplate#gofile.IsDir", templatePath, ErrNoTemplate}
		log.Info(err)
		TemplatesAvailable = false
	} else {
		TemplatesAvailable = true
	}
}
