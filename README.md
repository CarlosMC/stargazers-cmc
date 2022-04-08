# stargazers-cmc
Get stars count from specific git repos using go-github api



## EXPECTED OUTPUT

```
❯ export GH_AUTH_TOKEN="ghp_XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
❯ cd src/
❯ go run .
Stargazers:
    996icu/996.ICU: 261663
    EbookFoundation/free-programming-books: 229805
    freeCodeCamp/freeCodeCamp: 343607
```


## TESTS

```
❯ cd src/
❯ go test -v
=== RUN   TestGetNameAndStartCount
--- PASS: TestGetNameAndStartCount (1.20s)
=== RUN   TestGetRepoStarsList
--- PASS: TestGetRepoStarsList (0.23s)
PASS
ok  	github.com/CarlosMC/stargazers-cmc	1.632s
```


## DOCKER

```
❯ export GH_AUTH_TOKEN="ghp_XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
❯ docker run -e GH_AUTH_TOKEN=$GH_AUTH_TOKEN carlosmc/stargazers-cmc:lastest
```
or
```
❯ echo 'GH_AUTH_TOKEN=ghp_XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX' > .env
❯ docker run --env-file .env carlosmc/stargazers-cmc:lastest
```
