// Package gogithub implements a set of functions that
// access the GitHub API v3.0
package gomake

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	github "github.com/google/go-github/v34/github"
)

type Any interface{}

var client *github.Client = github.NewClient(nil)

// Err is used to check errors and exit immediately with
// a log message if the error is not nil.
func Err(err error) error {
	if err != nil {
		log.Error(err)
	}
	return err
}

func Orgs() error {

	// list all organizations for user "willnorris"
	orgs, _, err := client.Organizations.List(context.Background(), "willnorris", nil)
	if Err(err) != nil {
		return err
	}

	fmt.Println(orgs)
	return nil
}
