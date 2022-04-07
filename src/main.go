package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/go-github/v43/github"
	"gopkg.in/yaml.v3"
)

// declare struct to be use in yaml output
type RepoStars struct {
	Stargazers map[string]int `yaml:"Stargazers,omitempty"`
}

// define target repositories
var ReposList [3]string = [3]string{
	"freeCodeCamp/freeCodeCamp",
	"996icu/996.ICU",
	"EbookFoundation/free-programming-books",
}

func main() {

	// incialice RepoStars structure that contains the map {"full_name" : stargazers_count}
	repoStars4yaml := RepoStars{getRepoStarsList(ReposList)}

	// convert repoStarList Stuct to yaml format
	data, err := yaml.Marshal(&repoStars4yaml)
	if err != nil {
		log.Fatal(err)
	}

	// final console output
	fmt.Println(string(data))

}

func getRepoStarsList(reposList [3]string) map[string]int {

	// define map {"full_name" : stargazers_count}
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

func getNameAndStartCount(repo string) (string, int) {

	// authToken := os.Getenv("GB_AUTH_TOKEN")
	// if authToken == "" {
	// 	log.Fatal("No Auth token defined")
	// }

	ctx := context.Background()
	// ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: authToken})
	// tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(nil)

	// splitting to get owner and reponame
	repoOwnerAndName := strings.Split(repo, "/")

	// get the repo data
	repoData, _, err := client.Repositories.Get(ctx, repoOwnerAndName[0], repoOwnerAndName[1])
	if err != nil {
		log.Panicln("Could not get Repository Data!")
		log.Fatal(err)
	}

	// return FullName and StargazersCount
	return repoData.GetFullName(), repoData.GetStargazersCount()
}
