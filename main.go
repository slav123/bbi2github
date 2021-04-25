package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v35/github"
	"github.com/valyala/fastjson"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
)

// https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token

const (
	token    = ``
	owner    = ``
	repo     = ``
	jsonfile = `db-jira-cloud.json`
)

func main() {

	var p fastjson.Parser

	s, err := ioutil.ReadFile(jsonfile)
	if err != nil {
		fmt.Errorf("unable to open file %s", jsonfile)
		return
	}

	v, err := p.ParseBytes(s)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}

	projects := v.GetArray("projects")
	issues := projects[0].GetArray("issues")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	var status string

	for _, issue := range issues {

		status = string(issue.GetStringBytes("status"))

		fmt.Println(status)

		if status != "Done" {
			//fmt.Printf("%v\n", issue)
			title := issue.GetStringBytes("summary")
			body := issue.GetStringBytes("description")

			createIssue(client, ctx, string(title), string(body))

		}

	}
}

func createIssue(client *github.Client, ctx context.Context, title, body string) {

	input := &github.IssueRequest{
		Title: &title,
		Body:  &body,
	}

	issue, response, err := client.Issues.Create(ctx, owner, repo, input)
	if err != nil {
		fmt.Errorf("error while creating an issue %s", err)
	}

	fmt.Println(response)
	fmt.Println(issue)
}
