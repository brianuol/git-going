package main

import (
	"context"
	"fmt"
)

func main() {
	flags := ParseFlags()
	if flags.OrgName == "" {
		fmt.Printf("Sorry - OrgName is required to run this tool!")
		return
	}
	client, throttleMs := CreateClient(context.Background(), flags.AuthToken)
	repoInfo, err := CollectRepoInfo(context.Background(), client, flags.OrgName, flags.ReportOptions, throttleMs)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	PrintResults(repoInfo, flags.TopNVal, flags.ReportOptions)
}
