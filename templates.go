package gomake

import (
	"errors"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/gofile"
)

const defaultTemplatePath = "template_files"

var (
	ErrNoTemplatePath  error = errors.New("template directory not found")
	ErrNoTemplate      error = errors.New("template file not found")
	TemplatesAvailable bool  = false
	tempDir            *templateDir
)

type PathError = os.PathError

func init() {
	var err error
	tempDir, err = NewTemplateDir(defaultTemplatePath)
	if err != nil {
		TemplatesAvailable = false
	} else {
		TemplatesAvailable = true
	}
}

func NewTemplateDir(path string) (*templateDir, error) {

}

func NewTemplateFile(filename string) (*templateFile, error) {
	templateFileName := filepath.Join(defaultTemplatePath, filename)

	b, err := os.ReadFile(templateFileName)
	if err != nil {
		return "", &PathError{"ReadTemplate#os.Readfile", name, err}
	}
}

type TemplateDir interface {

	// types.Enabler
	Enable()
	Disable()
}

type (
	templateDir struct {
		enabled      bool
		templatePath string
	}

	templateFile struct {
		fi       os.FileInfo
		contents []byte
		isDirty  bool
	}
)

func (t *templateDir) GetFile(name string) (*templateFile, error) {

}

func ReadTemplate(name string) (string, error) {

	if !TemplatesAvailable {
		return "", ErrNoTemplatePath
	}

	templateFileName := filepath.Join(defaultTemplatePath, name)

	b, err := os.ReadFile(templateFileName)
	if err != nil {
		return "", &PathError{"ReadTemplate#os.Readfile", name, err}
	}

	return "", nil
}

func init() {
	if !gofile.IsDir(defaultTemplatePath) {
		err := &PathError{"ReadTemplate#gofile.IsDir", defaultTemplatePath, ErrNoTemplate}
		log.Info(err)
		TemplatesAvailable = false
	} else {
		TemplatesAvailable = true
	}
}
