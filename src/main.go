package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v3"
)

// define target repositories
var ReposList [3]string = [3]string{
	"freeCodeCamp/freeCodeCamp",
	"996icu/996.ICU",
	"EbookFoundation/free-programming-books",
}

// declare struct to be use in yaml output
type RepoStars struct {
	Stargazers map[string]int `yaml:"Stargazers,omitempty"`
}

// App
func main() {

	// incialice RepoStars structure that contains the map {"full_name" : stargazers_count}
	repoStars4yaml := RepoStars{getRepoStarsList(ReposList)}

	// convert repoStarList Struct to yaml format
	data, err := yaml.Marshal(&repoStars4yaml)
	if err != nil {
		log.Fatalln("Error while creating the YAML!")
		log.Fatalln(err)
	}

	// final console output
	fmt.Print("\n")
	fmt.Println(string(data))

}

// get Repo data for all repos in array parameter
// reposList string array repo names.
// return a map with Repo info: {"full_name" : stargazers_count}
func getRepoStarsList(reposList [3]string) map[string]int {

	// define map
	reposStartCount := make(map[string]int)

	// for each repo in the parameter array
	for _, repoName := range reposList {
		// get its full name and stars count
		fullName, starsCount := getNameAndStartCount(repoName)
		// add it to the map
		reposStartCount[""+fullName] = starsCount
	}
	//return the map
	return reposStartCount
}

// get Repo Data using go-github
// repo: string 'owner/repo'
// return FullName and StargazersCount
func getNameAndStartCount(repo string) (string, int) {

	// oauth example taken from: https://github.com/google/go-github/blob/master/example/commitpr/main.go

	// read env variable with oauth token
	authToken := os.Getenv("GH_AUTH_TOKEN")
	if authToken == "" {
		log.Fatalln("No Auth token defined")
	}

	// create github api client using oauth
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: authToken})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// splitting to get owner and reponame
	repoOwnerAndName := strings.Split(repo, "/")

	// get the repo data
	repoData, _, err := client.Repositories.Get(ctx, repoOwnerAndName[0], repoOwnerAndName[1])
	if err != nil {
		log.Fatalln("Could not get Repository Data!")
		log.Fatalln(err)
	}

	// return FullName and StargazersCount
	return repoData.GetFullName(), repoData.GetStargazersCount()
}
