package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	. "github.com/skeptycal/gofile"
	"github.com/skeptycal/gomake"
)

const (
	replRepoName       = "${REPONAME}"
	replGoVersion      = "${GOVERSION}"
	ReadmeTemplateName = "README_template.md"
	bakFile            = "README.md.bak"
)

func main() {

	readmeTemplate, err := gomake.ReadTemplate(ReadmeTemplateName)
	if err != nil {
		log.Fatal(err)
	}

	tmpFile, err := os.CreateTemp(gomake.TmpDir(), "README.md*") // ... but why not
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tmpFile: ", tmpFile.Name())

	f, err := os.Stat(ReadmeTemplateName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("readmeTemplate: ", f.Name())

	n, err := Copy(f.Name(), tmpFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	if n != f.Size() {
		log.Fatalf("wrong number of bytes copied: %d != %d", n, f.Size())
	}

	fmt.Println(readmeTemplate)
}
