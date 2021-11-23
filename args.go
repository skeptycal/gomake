package gomake

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type (
	arg struct {
		Name        string
		Short       string
		Value       Any
		Default     Any
		Description string
	}

	args map[string]Any

	Args interface {
		Exists(key string) (ok bool)
		Get(key string) (value Any, ok bool)
		Set(key string, value Any) error
		Name() string
		fmt.Stringer
	}
)

var (
	// Here and Me are the path and filename, respectively, of the
	// executable file.
	Here, Me = filepath.Split(os.Args[0])

	// commonArgs is a list of common command line arguments based on
	// the GNU standard (based on the July 1, 2021 release).
	//
	// References:
	//
	// - https://www.gnu.org/prep/standards/standards.html
	//
	// - https://tldp.org/LDP/abs/html/standard-options.html
	commonArgs []arg = []arg{
		{"help", "h", nil, false, "Help: Give usage message and exit."},
		{"version", "", nil, false, "Version: Show program version and exit."},
		{"all", "a", nil, false, "All: show all information or operate on all arguments."},
		{"list", "l", nil, false, "List: list files or arguments without taking other action."},
		{"quiet", "q", nil, 0, "Quiet: suppress stdout."},
		{"output", "o", nil, "", "Output: provide output file name."},
		{"recursive", "r", nil, false, "Recursive: Operate recursively (down directory tree)."},
		{"verbose", "v", nil, 0, "Verbose: output additional information to stdout or stderr."},
		{"compress", "z", nil, false, "Compress: apply compression (usually gzip)."},
		{"force", "f", nil, false, "Compress: apply compression (usually gzip)."},
	}
)

func init() {
	// fmt.Println("Here: ", Here)
	// fmt.Println("Me: ", Me)
}

func NewCLIArgs() Args {
	n := len(os.Args)
	a := make(args, n)
	_ = Err(a.parse())
	return a
}

// Exists returns true if the given key exists.
func (a args) Exists(key string) (ok bool) {
	_, ok = a.Get(key)
	return ok
}

// Get returns the value of the given key if it exists and
// returns nil otherwise. Ok is true if the key exists.
func (a args) Get(key string) (value Any, ok bool) {
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
func (a args) Set(key string, value Any) error {
	if _, ok := a[key]; ok {
		return fmt.Errorf("duplicate key not allowed in CLI arguments")
	}
	fmt.Println("map length: ", len(a))
	fmt.Printf("%v = %v\n", key, value)
	fmt.Println(a)
	// a[key] = value

	return nil
}

// Name returns the name of the program.
func (a args) Name() string {
	if len(a) < 1 {
		_ = Err(a.parse())
	}
	return a["program name"].(string)
}

func (a args) String() string {
	sb := strings.Builder{}
	defer sb.Reset()

	sb.WriteString("Command Line Arguments:\n")
	sb.WriteString("-----------------------\n")
	for k, v := range a {
		sb.WriteString(fmt.Sprintf("%20v = %30v\n", k, v))
	}
	return sb.String()
}

// parse parses the command line arguments and loads them into
// the Args map.
//
// If an error occurs, it is logged and parsing continues. If multiple
// errors occur, each one is logged but only the last one is returned.
func (a args) parse() (e1 error) {
	e1 = Err(a.Set("program name", os.Args[0]))
	for _, arg := range os.Args[1:] {
		arg = strings.ToLower(strings.TrimSpace(arg))
		if i := strings.Index(arg, "="); i == -1 {
			if err := Err(a.Set(arg, true)); err != nil {
				e1 = err
			}
		} else {
			if err := Err(a.Set(arg[:i], arg[i+1:])); err != nil {
				e1 = err
			}
		}
	}
	return
}
