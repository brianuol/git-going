package main

import (
	"math"
	"sort"
	"testing"

	"github.com/google/go-github/github"
)

func testNilIssueSearchResults(t *testing.T) {
	sgc, fc, rn := int(1234), int(12345), string("someRepo")
	repo := &github.Repository{
		Name:            &rn,
		StargazersCount: &sgc,
		ForksCount:      &fc}

	repoInfo := ToRepoInfo(repo, nil)
	if repo.GetName() != repoInfo.Name {
		t.Errorf("Expected repo name '%v' but found '%v'", repo.GetName(), repoInfo.Name)
	}
	if repo.GetStargazersCount() != repoInfo.Stars {
		t.Errorf("Expected repo stars '%v' but found '%v'", repo.GetStargazersCount(), repoInfo.Stars)
	}
	if repo.GetForksCount() != repoInfo.Forks {
		t.Errorf("Expected repo forks '%v' but found '%v'", repo.GetForksCount(), repoInfo.Forks)
	}
}

func testWithSearchResults(t *testing.T) {
	sgc, fc, rn := int(1234), int(12345), string("someRepo")
	repo := &github.Repository{
		Name:            &rn,
		StargazersCount: &sgc,
		ForksCount:      &fc}

	tc := int(100)
	issues := &github.IssuesSearchResult{Total: &tc}
	repoInfo := ToRepoInfo(repo, issues)
	if repo.GetName() != repoInfo.Name {
		t.Errorf("Expected repo name '%v' but found '%v'", repo.GetName(), repoInfo.Name)
	}
	if repo.GetStargazersCount() != repoInfo.Stars {
		t.Errorf("Expected repo stars '%v' but found '%v'", repo.GetStargazersCount(), repoInfo.Stars)
	}
	if repo.GetForksCount() != repoInfo.Forks {
		t.Errorf("Expected repo forks '%v' but found '%v'", repo.GetForksCount(), repoInfo.Forks)
	}
	if issues.GetTotal() != repoInfo.PRs {
		t.Errorf("Expected repo PRs '%v' but found '%v'", issues.GetTotal(), repoInfo.PRs)
	}
	expected := float64(issues.GetTotal()) / float64(repo.GetForksCount())
	if repoInfo.Contribution != expected {
		t.Errorf("Expected repo contribution '%v' but found '%v'", expected, repoInfo.Contribution)
	}
	fc = 0
	repoInfoNaN := ToRepoInfo(repo, issues)
	if !math.IsNaN(repoInfoNaN.Contribution) {
		t.Errorf("Expected NaN result for repo contribution but found '%v'", repoInfo.Contribution)
	}
}

func testSortRepoByStars(t *testing.T) {
	apples, artichokes, oranges := &RepoInfo{Name: "Apples", Stars: 100, Forks: 100, PRs: 100, Contribution: 1.000},
		&RepoInfo{Name: "Artichokes", Stars: 500, Forks: 200, PRs: 100, Contribution: .500},
		&RepoInfo{Name: "Oranges", Stars: 300, Forks: 300, PRs: 1000, Contribution: 3.333}

	repos := []*RepoInfo{apples, artichokes, oranges}
	sort.Sort(ReposByStars(repos))
	if repos[0].Name != "Artichokes" || repos[1].Name != "Oranges" {
		t.Error("Sort order test failed for Sort by Stars")
	}
}

func testSortReposByForks(t *testing.T) {
	apples, artichokes, oranges := &RepoInfo{Name: "Apples", Stars: 100, Forks: 100, PRs: 100, Contribution: 1.000},
		&RepoInfo{Name: "Artichokes", Stars: 500, Forks: 200, PRs: 100, Contribution: .500},
		&RepoInfo{Name: "Oranges", Stars: 300, Forks: 300, PRs: 1000, Contribution: 3.333}

	repos := []*RepoInfo{apples, artichokes, oranges}
	sort.Sort(ReposByForks(repos))
	if repos[0].Name != "Oranges" || repos[1].Name != "Artichokes" {
		t.Error("Sort order test failed for Sort by Forks")
	}
}

func testSortReposByPRs(t *testing.T) {
	apples, artichokes, oranges := &RepoInfo{Name: "Apples", Stars: 100, Forks: 100, PRs: 100, Contribution: 1.000},
		&RepoInfo{Name: "Artichokes", Stars: 500, Forks: 200, PRs: 100, Contribution: .500},
		&RepoInfo{Name: "Oranges", Stars: 300, Forks: 300, PRs: 1000, Contribution: 3.333}

	repos := []*RepoInfo{apples, artichokes, oranges}
	sort.Sort(ReposByPRs(repos))
	if repos[0].Name != "Oranges" || repos[1].Name != "Apples" {
		t.Error("Sort order test failed for Sort by PRs")
	}
}

func testSortReposByContribution(t *testing.T) {
	testFloat := float64(0)
	apples, artichokes, asparagus, oranges := &RepoInfo{Name: "Apples", Stars: 100, Forks: 100, PRs: 100, Contribution: 1.000},
		&RepoInfo{Name: "Artichokes", Stars: 500, Forks: 200, PRs: 100, Contribution: .500},
		&RepoInfo{Name: "Asparagus", Stars: 500, Forks: 200, PRs: 100, Contribution: (.500 / testFloat)},
		&RepoInfo{Name: "Oranges", Stars: 300, Forks: 300, PRs: 1000, Contribution: 3.333}

	repos := []*RepoInfo{apples, artichokes, asparagus, oranges}
	sort.Sort(ReposByContribution(repos))
	if repos[0].Name != "Oranges" || repos[1].Name != "Apples" || repos[3].Name != "Asparagus" {
		t.Error("Sort order test failed for Sort by PRs")
	}
}
