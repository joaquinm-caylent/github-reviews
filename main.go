package main

import (
	"context"
	"fmt"
	"strings"
	"os"
	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
)

func main() {
	// Arguments
	token := os.Getenv("GITHUB_TOKEN")
	repoList := strings.Split(os.Getenv("GITHUB_REPO_LIST"), ",")

	fmt.Println("Getting context...")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	fmt.Println("Authenticating with GitHub...")
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	user, _, err := client.Users.Get(ctx, "")
	listOptions := &github.PullRequestListOptions{
		State: "open",
	}
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}
	fmt.Println("Getting open PRs from repository list...")
	var prsAssignedToMe []*github.PullRequest
	for _, repo := range repoList {
		repoOwner := strings.Split(repo, "/")[0]
		repoName := strings.Split(repo,"/")[1]
		prs, _, err := client.PullRequests.List(ctx, repoOwner,repoName, listOptions)

		for _, pr := range prs {
			fmt.Println("--------------------------------------------------------------------------------")
			fmt.Println("Title: ",github.Stringify(pr.Title))
			fmt.Println("URL: ", github.Stringify(pr.HTMLURL))
			fmt.Println("Requested Reviewers: ")
			for _, reviewer := range pr.RequestedReviewers {
				fmt.Print(github.Stringify(reviewer.Login)," ")
				if github.Stringify(user.Login) == github.Stringify(reviewer.Login) {
					prsAssignedToMe = append(prsAssignedToMe, pr)
				}
			}
			fmt.Println()
			fmt.Println("Approved by : ")
			reviews,_,err := client.PullRequests.ListReviews(ctx, repoOwner,repoName, *pr.Number,nil)
			for _, review := range reviews {
				if github.Stringify(review.State) == "\"APPROVED\"" {
					fmt.Print(github.Stringify(review.User.Login), " ")
				}
			}
			if err != nil {
				fmt.Printf("\nerror: %v\n", err)
				return
			}
			fmt.Println()
			fmt.Println("--------------------------------------------------------------------------------")
		}
		if err != nil {
			fmt.Printf("\nerror: %v\n", err)
			return
		}
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("Requested review in: ")
	for _, prToReview := range prsAssignedToMe {
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		fmt.Println("Title: ",github.Stringify(prToReview.Title))
		fmt.Println("URL: ", github.Stringify(prToReview.HTMLURL))
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	}
}
