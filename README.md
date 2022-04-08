# stargazers-cmc

Thiss is a Go-lang app that gets the Stars count from the following git repositories using go-github library to consume the github Rest api.
It uses Github OAuth Tokens for auntentication.

- 996icu/996.ICU
- freeCodeCamp/freeCodeCamp
- EbookFoundation/free-programming-books

The output of the app is sent to the stdout in yaml format with a top level mapping called **Stargazers**.


## Local run and expected output:

To try this app locally, please export a env variable called **GH_AUTH_TOKEN** with your OAuth Token before run it, like this:

```
❯ export GH_AUTH_TOKEN="ghp_XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
❯ cd src/
❯ go run .
Stargazers:
    996icu/996.ICU: 261663
    EbookFoundation/free-programming-books: 229805
    freeCodeCamp/freeCodeCamp: 343607
```

The output should be similar to the one showed before, numbers may change.


## Tests:

There are two unit tests to be passed:

1. **TestGetNameAndStartCount**: check if the app is correctly retrieving the information from one repo at a time.
2. **TestGetRepoStarsList**: check if the app is correclty retrieving the total amount of repositires data.


Test can be excecuted with the following command, it also requieres the **GH_AUTH_TOKEN** variable to be defined before running it:

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


## Docker Image

A Docker image has been built so you can run the containerized app
The Docker image also requires the **GH_AUTH_TOKEN** variable to be defined before running it.

You can export and pass it directly from command line like this:


```
❯ export GH_AUTH_TOKEN="ghp_XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
❯ docker run -e GH_AUTH_TOKEN=$GH_AUTH_TOKEN carlosmc/stargazers-cmc:lastest
```

or you can save the variable into a file (added to .gitignore) for further usage and pass it to Docker with **--env-file** option:

```
❯ echo 'GH_AUTH_TOKEN=ghp_XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX' > .env
❯ docker run --env-file .env carlosmc/stargazers-cmc:lastest
```
