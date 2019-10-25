package config

import (
	"os"
)

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

//GetGithubAccessToken .
func GetGithubAccessToken() string {
	return githubAccessToken
}