package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/google/go-github/github"
)

var (
	userName   = flag.String("u", "", "Github Username of Repo Owner")
	repoName   = flag.String("r", "", "Github Repo Name")
	readmeFile = flag.String("f", "README.md", "Readme filename")
	backupOrig = flag.Bool("b", false, "Backup Readme")
)

func main() {
	flag.Parse()
	if *userName == "" || *repoName == "" {
		fmt.Println("not enough arguments. please pass user and repo name")
		os.Exit(1)
	}

	client := github.NewClient(nil)
	contribs, _, err := client.Repositories.ListContributors(context.Background(), *userName, *repoName, nil)
	if err != nil {
		fmt.Println(err)
	}

	var contribStr strings.Builder
	for _, v := range contribs {
		if *v.Login != *userName {
			contribStr.WriteString("- [" + *v.Login + "](" + *v.HTMLURL + ") \n")
		}
	}

	input, err := ioutil.ReadFile(*readmeFile)
	if err != nil {
		fmt.Println(err)
	}

	if *backupOrig {
		err = ioutil.WriteFile(*readmeFile+".orig", input, 0644)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "CONTRIBPOPULATE") {
			lines[i] = contribStr.String()
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(*readmeFile, []byte(output), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
