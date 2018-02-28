package main

import (
	"math"

	"github.com/google/go-github/github"
)

// RepoInfo ... Struct for tracking repo attributes we care about
type RepoInfo struct {
	Name         string
	Stars        int
	Forks        int
	PRs          int
	Contribution float64
}

// ReposByStars ... Sort implementation for RepoInfo by Repo Stars Descending, Repo Name Ascending
type ReposByStars []*RepoInfo

func (r ReposByStars) Len() int      { return len(r) }
func (r ReposByStars) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
func (r ReposByStars) Less(i, j int) bool {
	if r[i].Stars == r[j].Stars {
		return r[i].Name < r[i].Name
	}
	return r[i].Stars > r[j].Stars
}

// ReposByForks ... Sort implementation for RepoInfo by Repo Forks Descending, Repo Name Ascending
type ReposByForks []*RepoInfo

func (r ReposByForks) Len() int      { return len(r) }
func (r ReposByForks) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
func (r ReposByForks) Less(i, j int) bool {
	if r[i].Forks == r[j].Forks {
		return r[i].Name < r[i].Name
	}
	return r[i].Forks > r[j].Forks
}

// ReposByPRs ... Sort implementation for RepoInfo by Repo PR count Descending, Repo Name Ascending
type ReposByPRs []*RepoInfo

func (r ReposByPRs) Len() int      { return len(r) }
func (r ReposByPRs) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
func (r ReposByPRs) Less(i, j int) bool {
	if r[i].PRs == r[j].PRs {
		return r[i].Name < r[i].Name
	}
	return r[i].PRs > r[j].PRs
}

// ReposByContribution ... Sort implementation for RepoInfo by Repo Contribution Descending, Repo Name Ascending
type ReposByContribution []*RepoInfo

func (r ReposByContribution) Len() int      { return len(r) }
func (r ReposByContribution) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
func (r ReposByContribution) Less(i, j int) bool {
	if r[i].Contribution == r[j].Contribution {
		return r[i].Name < r[j].Name
	}
	if math.IsNaN(r[i].Contribution) {
		return false
	}
	if math.IsNaN(r[j].Contribution) {
		return true
	}
	return r[i].Contribution > r[j].Contribution
}

// ToRepoInfo ... returns a pointer to a RepoInfo instance whose values have been populated using github data.  If argument value for result is nil, PRs and Contribution will be initialized as zero.
func ToRepoInfo(repo *github.Repository, result *github.IssuesSearchResult) *RepoInfo {
	prs, pct := 0, float64(0)
	if result != nil {
		prs = result.GetTotal()
		pct = float64(result.GetTotal()) / float64(repo.GetForksCount())
	}
	return &RepoInfo{
		Name:         repo.GetName(),
		Stars:        repo.GetStargazersCount(),
		Forks:        repo.GetForksCount(),
		PRs:          prs,
		Contribution: pct}
}
