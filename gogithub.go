// Package gogithub implements a set of functions that
// access the GitHub API v3.0
package gomake

import (
	"context"
	"fmt"

	github "github.com/google/go-github/v34/github"
	"github.com/skeptycal/errorlogger"
)

var (
	Log = errorlogger.Log
	Err = errorlogger.Err
)

var client *github.Client = github.NewClient(nil)

func Orgs() error {

	// list all organizations for user "willnorris"
	orgs, _, err := client.Organizations.List(context.Background(), "willnorris", nil)
	if Err(err) != nil {
		return err
	}

	fmt.Println(orgs)
	return nil
}
