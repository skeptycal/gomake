package gomake

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"

	"github.com/skeptycal/gofile"
	"github.com/skeptycal/types"
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
	if Err(err) != nil {
		TemplatesAvailable = false
	} else {
		TemplatesAvailable = true
	}
}

func NewTemplateDir(path string) (*templateDir, error) {
	return nil, gofile.ErrNotImplemented
}

func NewTemplateFile(fileName string) (*templateFile, error) {
	templateFileName := filepath.Join(defaultTemplatePath, fileName)

	fi, err := os.Stat(templateFileName)
	if err != nil {
		return nil, gofile.NewGoFileError("NewTemplateFile#os.Stat()", templateFileName, err)
	}

	b, err := os.ReadFile(templateFileName)
	if err != nil {
		return nil, gofile.NewGoFileError("NewTemplateFile#os.Readfile()", templateFileName, err)
	}

	buf := bytes.NewBuffer(make([]byte, 0, int(float64(len(b))*1.5)))
	buf.Write(b)

	return &templateFile{fi: fi, contents: buf, isDirty: false}, nil
}

func ReadTemplate(fileName string) (string, error) {

	if !TemplatesAvailable {
		return "", ErrNoTemplatePath
	}

	templateFileName := filepath.Join(defaultTemplatePath, fileName)

	b, err := os.ReadFile(templateFileName)
	if err != nil {
		return "", gofile.NewGoFileError("ReadTemplate#os.Readfile()", fileName, err)
	}

	return string(b), nil
}

type TemplateDir interface {
	types.Enabler
	GetFile(name string) (*templateFile, error)
}

type templateDir struct {
	enabled      bool
	templatePath string
}

func (d *templateDir) Enable()  { d.enabled = true }
func (d *templateDir) Disable() { d.enabled = false }

type templateFile struct {
	fi       os.FileInfo
	contents *bytes.Buffer
	isDirty  bool
}

func (t *templateDir) GetFile(name string) (*templateFile, error) {
	return nil, gofile.ErrNotImplemented
}
