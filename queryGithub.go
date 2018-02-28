package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	authDelay = 3000
	anonDelay = 10000
)

// CreateClient ... creates a github client and returns the appropriate duration (in milliseconds) for throttling to keep within rate-limit policy
func CreateClient(ctx context.Context, authToken string) (*github.Client, time.Duration) {
	throttleMs := time.Duration(anonDelay)
	var client *github.Client
	if len(authToken) > 0 {
		throttleMs = authDelay
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: authToken},
		)
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	} else {
		client = github.NewClient(nil)
	}
	return client, throttleMs
}

// CollectRepoInfo ... Gets relevant metadata for the github organization 'orgName', for the reports flagged for execution at runtime
func CollectRepoInfo(ctx context.Context, client *github.Client, orgName string, options map[string]bool, throttleMs time.Duration) ([]*RepoInfo, error) {
	fmt.Printf("Starting...\n")
	repoLevel := options["pulled"] || options["contributed"]
	repos, err := getReposForOrg(orgName, client)
	if err != nil {
		return nil, err
	}
	repoInfo := make([]*RepoInfo, len(repos), len(repos))
	for i, repo := range repos {
		// if get repo info..
		if repoLevel {
			time.Sleep(throttleMs * time.Millisecond)
			result, _, err := client.Search.Issues(ctx, "type:pr repo:"+repo.GetFullName(), nil)
			if err != nil {
				return nil, err
			}
			fmt.Printf("%v  (%v of %v)\n", repo.GetName(), i+1, len(repos))
			repoInfo[i] = ToRepoInfo(repo, result)
		} else {
			repoInfo[i] = ToRepoInfo(repo, nil)
		}
	}
	return repoInfo, err
}

func getReposForOrg(orgName string, client *github.Client) ([]*github.Repository, error) {
	var allRepos []*github.Repository
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	for {
		repos, resp, err := client.Repositories.ListByOrg(context.Background(), orgName, opt)
		if err != nil {
			return nil, err
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return allRepos, nil
}
