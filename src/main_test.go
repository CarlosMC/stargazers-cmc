package main

import (
	"testing"
)

func TestGetNameAndStartCount(t *testing.T) {
	// Iterate over ReposList target from main
	for _, expected := range ReposList {
		// get name and stars count
		name, count := getNameAndStartCount(expected)
		// test them
		if name != expected {
			t.Errorf("Repo Name %q not equal to expected %q", name, expected)
		}
		if count <= 0 {
			t.Errorf("Repo Start count lower than zero: %d", count)
		}

	}
}

func TestGetRepoStarsList(t *testing.T) {
	// get the Map with fullNames and stargazersCount for all targets (3) from Main.
	repoStarsCount := getRepoStarsList(ReposList)
	mapSize := len(repoStarsCount)
	expected := 3
	if mapSize != expected {
		t.Errorf("repoStarsCount MAP size %q different than expedted: %q", mapSize, expected)
	}

}
