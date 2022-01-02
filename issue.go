package main

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/google/go-github/v41/github"
)

// Get List of GitHub issues, and add the sites to the list
func IssueCommand() {
	f := GetListFile()
	t := GetList(f)

	client := github.NewClient(nil)

	opt := &github.IssueListByRepoOptions{Labels: []string{"accepted"}, State: "open"}
	issues, _, err4 := client.Issues.ListByRepo(context.Background(), "StopModReposts", "Illegal-Mod-Sites", opt)
	HandleError(err4)

	var sites []string
	var issueids []int
	var commitmsg string

	for _, issue := range issues {
		var title string = *issue.Title
		var prefix1 string = "New site to add: "
		issueids = append(issueids, issue.GetNumber())
		hasPrefix := strings.HasPrefix(title, prefix1)

		if hasPrefix {
			title = title[len(prefix1):]
			// https://stackoverflow.com/a/25323169
			re := regexp.MustCompile(`(?:[\w-]+\.)+[\w-]+`)
			reTitle := re.FindAll([]byte(title), -1)

			if len(reTitle) > 0 {
				title = string(reTitle[0])
			}
		}

		log.Println(title)
		sites = append(sites, title)
	}

	for _, siteText := range sites {
		t = append(t, IllegalSite{Domain: siteText, Notes: "/", Path: "/", Reason: "Illegal redistribution"})
	}

	for _, id := range issueids {
		commitmsg = commitmsg + "Closes #" + fmt.Sprint(id) + " "
	}

	SaveList(f, t)
	warnMsg := "The parsing is very fragile; please review the changes via Git and change the titles"
	LogSeparator(len(warnMsg) - 1)
	log.Println(warnMsg)
	LogSeparator(len(warnMsg) - 1)
	log.Println("Generated commit message", commitmsg)
}
