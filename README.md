# github-codefresh-step

How to use?
```
export GITHUB_TOKEN=yourtoken
export GITHUB_REPO_LIST="joaquinm-caylent/repo1,joaquinm-caylent/repo2,caylent/repo3"
./github-reviews
```
or
```
docker run -it -e GITHUB_TOKEN=yourtoken -e GITHUB_REPO_LIST="joaquinm-caylent/repo1,joaquinm-caylent/repo2,caylent/repo3" joaquinmontero/github-reviews
```