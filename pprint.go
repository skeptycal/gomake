package gomake

import (
	"bytes"
	"fmt"
	"strings"
)

func PPrint(v Any) {
	// a := ansi.NewAnsiString(ansi.NewColor(Black, Black, Bold))

	// color := ansi.NewColor(1, 0, 0)
	switch t := v.(type) {

	case int, float32, float64, bool:
		fmt.Printf("%v\n", t)

	case string:
		fmt.Printf("%q\n", t)

	default:
		// ansi.Print(color, "(type %v) %v\n", t, v)
		fmt.Printf("(type %T) %v\n", t, t)

	}
	// ansi.Reset()
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
