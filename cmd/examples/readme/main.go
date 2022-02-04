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

	readmeTemplateContents, err := gomake.ReadTemplate(ReadmeTemplateName)
	if err != nil {
		log.Fatal(err)
	}

	tmpDir := os.TempDir() // not needed for os.CreateTemp - it is default

	tmpFile, err := os.CreateTemp(tmpDir, "README.md*") // ... but why not
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tmpFile: ", tmpFile.Name())

	readmeTemplate, err := os.Stat(ReadmeTemplate)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("readmeTemplate: ", readmeTemplate.Name())

	n, err := Copy(readmeTemplate.Name(), tmpFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	if n != readmeTemplate.Size() {
		log.Fatalf("wrong number of bytes copied: %d != %d", n, readmeTemplate.Size())
	}
}
