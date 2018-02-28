package main

import (
	"flag"
)

// RuntimeFlags ... Runtime flags for git-going execution
type RuntimeFlags struct {
	OrgName       string
	TopNVal       int
	AuthToken     string
	ReportOptions map[string]bool
}

// ParseFlags ... returns the org name, auth token, topN, and report options to execute based on flags provided to console
func ParseFlags() RuntimeFlags {
	options := make(map[string]bool)

	var orgName string
	var topNVal int
	var authToken string

	flag.StringVar(&orgName, "orgName", "", "Required. The github org to analyze")
	flag.IntVar(&topNVal, "topN", 25, "Optional (default 25).  The number of results (top-n) to display for each analysis set")
	flag.StringVar(&authToken, "authToken", "", "Optional.  The auth token to provide for faster search execution (anonymous runs 3x slower)")

	flag.Parse()

	if len(flag.Args()) > 0 {
		for i := range flag.Args() {
			arg := flag.Args()[i]
			if arg == "starred" || arg == "forked" || arg == "pulled" || arg == "contributed" {
				options[arg] = true
			}
		}
	} else {
		options["starred"] = true
		options["forked"] = true
		options["pulled"] = true
		options["contributed"] = true
	}

	return RuntimeFlags{OrgName: orgName, TopNVal: topNVal, AuthToken: authToken, ReportOptions: options}
}
