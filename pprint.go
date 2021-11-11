package gomake

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/skeptycal/ansi"
)

func PPrint(v Any) {
	color := ansi.NewColor(1, 0, 0)

	a := ansi.NewAnsiString(color, v.(string))

	switch t := v.(type) {

	case int, float32, float64, bool:
		fmt.Printf("%v%v\n", a, v)

	case string:
		fmt.Printf("%v%v\n", a, v)

	default:
		fmt.Printf("%v(type %v) %v\n", color, t, v)

	}
	fmt.Print(ansi.Reset)
}

var PPrintMap map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
}

var PPrintAny Any = PPrintMap
var PPrintAnySlice []Any = []Any{
	string("PPrintAnySlice"),
	bytes.Buffer{},
	func(v Any) string {
		return "func return string"
	},
	nil,
}

var PPrintSamples []Any = []Any{
	string("PPrint samples"),
	0,
	"string",
	1.0,
	true,
	strings.Builder{},
	nil,
	struct{}{},
	map[int]int{1: 1, 2: 2, 3: 3},
	PPrintMap,
	PPrintAny,
	PPrintAnySlice,
}

func SamplePPrint(samples []Any) {
	if samples == nil {
		samples = PPrintSamples
	}

	for _, sample := range samples {
		PPrint(sample)
	}
}
