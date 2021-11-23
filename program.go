package gomake

import (
	"fmt"
	"os"
	"path/filepath"
)

type (
	Program interface {
		Name() string
		Path() string
		Help() string
		Version() string
		Args() Args
		fmt.Stringer
	}
)

func NewProgram(helpMessage, version string) Program {
	here, me := filepath.Split(os.Args[0])
	return &program{
		name:        me,
		path:        here,
		helpMessage: helpMessage,
		version:     version,
		args:        NewCLIArgs(),
	}
}

type program struct {
	name        string
	path        string
	helpMessage string
	version     string
	args        Args
}

func (p *program) Name() string        { return p.name }
func (p *program) Path() string        { return p.path }
func (p *program) HelpMessage() string { return p.helpMessage }
func (p *program) Version() string     { return p.version }
func (p *program) Args() Args          { return p.args }
