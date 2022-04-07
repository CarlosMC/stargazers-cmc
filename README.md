# stargazers-cmc
Get stars count from specific git repos using go-github api





TESTS

❯ go test -v
=== RUN   TestGetNameAndStartCount
--- PASS: TestGetNameAndStartCount (1.20s)
=== RUN   TestGetRepoStarsList
--- PASS: TestGetRepoStarsList (0.23s)
PASS
ok  	github.com/CarlosMC/stargazers-cmc	1.632s


EXPECTED OUTPUT

❯ go run .
Stargazers:
    996icu/996.ICU: 261663
    EbookFoundation/free-programming-books: 229805
    freeCodeCamp/freeCodeCamp: 343607