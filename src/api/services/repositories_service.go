package services

import (
	"strings"

	"github.com/fabiosebastiano/golang-microservices/src/api/config"
	"github.com/fabiosebastiano/golang-microservices/src/api/domain/github"
	"github.com/fabiosebastiano/golang-microservices/src/api/domain/repositories"

	//"github.com/fabiosebastiano/golang-microservices/src/api/log/log_logrus"
	"github.com/fabiosebastiano/golang-microservices/src/api/log/log_zap"
	"github.com/fabiosebastiano/golang-microservices/src/api/providers/github_provider"
	"github.com/fabiosebastiano/golang-microservices/src/api/utils/errors"
)

type reposService struct {
}

type reposServiceInterface interface {
	CreateRepo(clientId string, createRepoRequest repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (r *reposService) CreateRepo(clientId string, requestInput repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {

	requestInput.Name = strings.TrimSpace(requestInput.Name)
	if requestInput.Name == "" {
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	request := github.CreateRepoRequest{
		Name:        requestInput.Name,
		Description: requestInput.Description,
		Private:     false,
	}

	//log_zap.Info("Sending request to external API", fmt.Sprintf("client_id:%s", clientId), "status:pending")
	log_zap.Info("Sending request to external API",
		log_zap.Field("client_id", clientId),
		log_zap.Field("status", "pending"),
	)

	response, error := github_provider.CreateRepo(config.GetGithubAccessToken(), request)

	if error != nil {
		log_zap.Error("Response obtained from external API", error,
			log_zap.Field("client_id", clientId),
			log_zap.Field("status", "error"),
			//fmt.Sprintf("client_id:%s", clientId), "status:error"
		)

		return nil, errors.NewApiError(error.StatusCode, error.Message)
	}
	log_zap.Info("Response obtained from external API",
		log_zap.Field("client_id", clientId),
		log_zap.Field("status", "success"),
	//fmt.Sprintf("client_id:%s", clientId), "status:success"
	)

	return &repositories.CreateRepoResponse{
		Name:  response.Name,
		Id:    response.Id,
		Owner: response.Owner.Login,
	}, nil
}
