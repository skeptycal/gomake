package gomake

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/skeptycal/anybool"
	"github.com/skeptycal/defaults"
	"github.com/skeptycal/gofile"
)

const (
	normalMode        os.FileMode = 0644
	dirMode           os.FileMode = 0755
	defaultBufferSize int         = 1024
	minBufferSize     int64       = 16
)

var Options defaults.DefaultMapper = defaults.NewDefaults(false, false)

type Mapper interface {
	Get(key string) (Any, error)
	Set(key string, value Any) error
	String() string

	// Is returns true if the key is true
	Is(key string) bool

	// Not returns true if the key is false
	Not(key string) bool

	IsDebug() bool
}

type AnsiString struct {
	// fg byte
	// bg byte
	// ef byte
	s string
}

const ansiEncode = "\x1b[%d;%d;%dm"

func (s *AnsiString) Set(fg, bg, ef byte) {
	s.s = fmt.Sprintf(ansiEncode, fg, bg, ef)
}

func (s AnsiString) String() string {
	return s.s
}

type cliSettings struct {
	fg byte
	bg byte
	ef byte

	main    string
	attn    string
	warn    string
	choice  string
	second  string
	third   string
	bold    string
	inverse string
	reset   string
}

func (c *cliSettings) Ansi(fg, bg, ef byte) string {
	return ""
}

func (c *cliSettings) Main() string {
	if c.main == "" {
		c.main = c.Ansi(c.fg, c.bg, c.ef)
	}
	return ""
}

type optionSet struct {
	m     map[string]Any
	debug bool
	trace bool
	cli   cliSettings
	out   io.Writer
	log   io.Writer
}

func init() {
	Options.Set("traceState", false)
	Options.Set("TruncateFiles", true)
}

func DoTrunc() bool {
	o, err := Options.Get("TruncateFiles")
	if err != nil {
		return false
	}
	b, ok := o.(defaults.Booler)
	if !ok {
		return false
	}
	return b.AsBool()
}

var OptionFileTruncate = anybool.AnyBooler(true)

func Copy(dstName, srcName string) error {

	_, err := gofile.StatCheck(dstName)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			return err
		}
	}

	dst, err := os.Create(dstName)
	if err != nil {
		return err
	}

	src, err := os.Open(srcName)
	if err != nil {
		return err
	}

	io.Copy(dst, src)
	return nil
}

func readBak(filename string) ([]byte, error) {
	// _, err := io.Copy( .Copy(filename, filename+".bak")
	// gofile.
	// if err != nil {
	// 	return nil, err
	// }

	// return os.ReadFile(filename)
	return nil, nil
}
