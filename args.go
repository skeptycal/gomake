package gomake

import (
	"fmt"
	"os"
	"strings"
)

type (
	Args map[string]Any
)

// Get returns the value of the given key if it exists and
// returns nil otherwise. Ok is true if the key exists.
func (a Args) Get(key string) (value Any, ok bool) {
	if value, ok := a[key]; ok {
		return value, true
	}
	return nil, false
}

// Set sets the value of the given key. Keys are permanent for
// the duration of the current operation. The keys may not be
// overwritten and duplicates are not allowed.
//
// If the value cannot be updated, an error is returned. The values
// may be of any type.
func (a Args) Set(key string, value Any) error {
	if _, ok := a[key]; ok {
		return fmt.Errorf("duplicate key not allowed in CLI arguments")
	}
	a[key] = value
	return nil
}

// Name returns the name of the program.
func (a Args) Name() string {
	if len(a) < 1 {
		a.Parse()
	}
	return a["program name"].(string)
}

// Parse parses the command line arguments and loads them into
// the Args map.
//
// If an error occurs, it is logged and parsing continues. If multiple
// errors occur, each one is logged but only the last one is returned.
func (a Args) Parse() (e1 error) {
	e1 = Err(a.Set("program name", os.Args[0]))
	for _, arg := range os.Args[1:] {
		arg = strings.ToLower(strings.TrimSpace(arg))
		if i := strings.Index(arg, "="); i == -1 {
			if err := Err(CliArgs.Set(arg, true)); err != nil {
				e1 = err
			}
		} else {
			if err := Err(CliArgs.Set(arg[:i], arg[i+1:])); err != nil {
				e1 = err
			}
		}
	}
	return
}

func (a Args) String() string {
	sb := strings.Builder{}
	defer sb.Reset()

	sb.WriteString("Command Line Arguments:\n")
	sb.WriteString("-----------------------\n")
	for k, v := range a {
		sb.WriteString(fmt.Sprintf("%20v = %30v\n", k, v))
	}
	return sb.String()
}

// CliArgs contains the command line arguments.
var CliArgs Args

func CheckCLI(key string) (ok bool) {
	_, ok = CliArgs.Get(key)
	return ok
}

func GetCLI(key string) (v Any, ok bool) {
	return CliArgs.Get(key)
}
