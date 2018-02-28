package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

// PrintResults ... Prints the top-n Repos for each report type requested
func PrintResults(repoInfo []*RepoInfo, topNVal int, options map[string]bool) {
	if options["starred"] {
		sort.Sort(ReposByStars(repoInfo))
		printTopN(repoInfo, topNVal, "Stars", "Total Stars", 0)
	}
	if options["forked"] {
		sort.Sort(ReposByForks(repoInfo))
		printTopN(repoInfo, topNVal, "Forks", "Total Forks", 0)
	}
	if options["pulled"] {
		sort.Sort(ReposByPRs(repoInfo))
		printTopN(repoInfo, topNVal, "PRs", "Total Pull Requests", 0)
	}
	if options["contributed"] {
		sort.Sort(ReposByContribution(repoInfo))
		printTopN(repoInfo, topNVal, "Contribution", "Contribution Pct (PRs / Forks)", 3)
	}
}

func printTopN(repoInfo []*RepoInfo, topN int, propertyName string, uom string, precision int) {
	fmt.Printf("\n\n### Top %v Repositories, by %v ###\n\n", topN, uom)
	fmt.Printf("Rank\t\t%-32s%v\n", "Repo", propertyName)
	fmt.Printf("====\t\t%-32s%v\n", "===========", "===========")
	for i, repo := range repoInfo {
		if i >= topN {
			break
		}
		numFormat := "%d"
		if precision > 0 {
			numFormat = "%." + strconv.Itoa(precision) + "f"
		}
		fmt.Printf("%-4d\t\t%-32s"+numFormat+"\n", i+1, repo.Name, reflect.Indirect(reflect.ValueOf(repo)).FieldByName(propertyName))
	}
}
