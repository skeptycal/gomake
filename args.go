package gomake

import (
	"fmt"
	"os"
	"strings"
)

type Args map[string]Any

func (a Args) Get(key string) (value Any, ok bool) {
	if value, ok := a[key]; ok {
		return value, true
	}
	return nil, false
}

func (a Args) Set(key string, value Any) error {
	if _, ok := a[key]; ok {
		return fmt.Errorf("duplicate key not allowed in CLI arguments")
	}
	a[key] = value
	return nil
}

// Parse parses the command line arguments and loads them into
// the global cliArgs map.
//
// If an error occurs, it is logged. Only the last error is returned.
func (a Args) Parse() (e1 error) {
	e1 = nil
	err := a.Set("program name", os.Args[0])
	if err != nil {
		e1 = err
		Err(err)
	}
	for _, arg := range os.Args[1:] {
		arg = strings.ToLower(strings.TrimSpace(arg))
		if i := strings.Index(arg, "="); i == -1 {
			err := cliArgs.Set(arg, true)
			if err != nil {
				e1 = err
				Err(err)
			}
		} else {
			err := cliArgs.Set(arg[:i], arg[i+1:])
			if err != nil {
				e1 = err
				Err(err)
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

var cliArgs Args
