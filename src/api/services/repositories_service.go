package services

import (
	"strings"

	"github.com/fabiosebastiano/golang-microservices/src/api/config"
	"github.com/fabiosebastiano/golang-microservices/src/api/domain/github"
	"github.com/fabiosebastiano/golang-microservices/src/api/domain/repositories"
	"github.com/fabiosebastiano/golang-microservices/src/api/providers/github_provider"
	"github.com/fabiosebastiano/golang-microservices/src/api/utils/errors"
)

type reposService struct {
}

type reposServiceInterface interface {
	CreateRepo(createRepoRequest repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (r *reposService) CreateRepo(requestInput repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {

	requestInput.Name = strings.TrimSpace(requestInput.Name)
	if requestInput.Name == "" {
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	request := github.CreateRepoRequest{
		Name:        requestInput.Name,
		Description: requestInput.Description,
		Private:     false,
	}

	response, error := github_provider.CreateRepo(config.GetGithubAccessToken(), request)

	if error != nil {
		return nil, errors.NewApiError(error.StatusCode, error.Message)
	}

	return &repositories.CreateRepoResponse{
		Name:  response.Name,
		Id:    response.Id,
		Owner: response.Owner.Login,
	}, nil
}
