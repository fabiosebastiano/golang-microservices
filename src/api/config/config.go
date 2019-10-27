package config

import (
	"os"
)

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
	environment          = "GO_ENVIRONMENT"
	production           = "production"
	Loglevel             = "info"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

//GetGithubAccessToken .
func GetGithubAccessToken() string {
	return githubAccessToken
}

func IsProduction() bool {
	return os.Getenv(environment) == production
}
