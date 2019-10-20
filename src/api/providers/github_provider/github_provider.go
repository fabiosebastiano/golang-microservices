package github_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fabiosebastiano/golang-microservices/src/api/clients/restclient"
	"github.com/fabiosebastiano/golang-microservices/src/api/domain/github"
)

/*
Per costruire l'header con il token, creo due const ed una fx che aggiunge il token
*/
const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(authorizationToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, authorizationToken) //) //verrà poi spostato in un SECRET
}

//CreateRepo .
func CreateRepo(authorizationToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(authorizationToken)) //"3bf869c0268cf940222245eed3f0d276301f02b9"

	response, err := restclient.Post(urlCreateRepo, request, headers)

	if err != nil {
		log.Printf("error when trying to create new repo in github: %s", err.Error())
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid response body",
		}
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Invalid json response body",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Printf("error when trying to unmarshale create repo successful response: %s", err.Error())
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error unmarshalling successful response",
		}
	}

	return &result, nil
}
