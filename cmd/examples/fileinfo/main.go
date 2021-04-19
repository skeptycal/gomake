package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

func removebounded(s, start, end string) {

	var str []string = []string{}
	startS := strings.Split(s, start)

	for _, line := range startS {
		str = append(str, strings.Split(line, end)[1:]...)
	}
	js := strings.Join(str, "")
	str = strings.Split(js, ",")

	mp := make(map[string]string, len(str))

	for _, val := range str {
		var vv string
		var kk string
		kv := strings.Split(val, ":")
		if len(kv) > 0 {
			kk = strings.TrimSpace(kv[0])
		}
		if len(kv) > 1 {
			vv = strings.TrimSpace(kv[1])
		}
		mp[kk] = vv
	}

	for k, v := range mp {
		if k != "" && v != "" {
			fmt.Printf("%-20.20s: %-25.25s\n", k, v)
		}
	}
}

func main() {
	filename, err := filepath.Abs(os.Args[0])
	if err != nil {
		log.Fatal(err)
	}

	fi, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("file info example:")

	fmt.Printf("filename: %#v\n", filename)
	fmt.Printf("Name(): %#v\n", fi.Name())
	fmt.Printf("Size(): %#v\n", fi.Size())
	fmt.Printf("IsDir(): %#v\n", fi.IsDir())
	fmt.Printf("ModTime(): %#v\n", fi.ModTime())
	fmt.Printf("Mode(): %#v\n", fi.Mode())
	fmt.Println("")
	fmt.Printf("fileinfo type: %#T\n", fi)

	sysString := fmt.Sprintf("Sys(): %#v\n", fi.Sys())
	// fmt.Printf("sysString: %#v\n", sysString)

	fstring := fmt.Sprintf("fileinfo: %#v\n", fi)
	// fmt.Printf("fstring: %#v\n", fstring)

	fmt.Println("****************************")
	removebounded(sysString, "{", "}")
	fmt.Println("")
	fmt.Println("****************************")
	removebounded(fstring, "{", "}")
	fmt.Println("****************************")

}
