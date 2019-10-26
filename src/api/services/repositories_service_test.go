package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/fabiosebastiano/golang-microservices/src/api/clients/restclient"
	"github.com/fabiosebastiano/golang-microservices/src/api/domain/repositories"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidInputName(t *testing.T) {
	request := repositories.CreateRepoRequest{}

	result, error := RepositoryService.CreateRepo(request)

	assert.Nil(t, result)
	assert.NotNil(t, error)
	assert.EqualValues(t, http.StatusBadRequest, error.Status())
	assert.EqualValues(t, "invalid repository name", error.Message())

}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":"Requires authentication", "documentation_url":"http://developer.github."}`)),
		},
	})
	request := repositories.CreateRepoRequest{Name: "testing"}

	result, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "Requires authentication", err.Message())

}
func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":123, "name":"testing", "owner":{"login":"fabiosebastiano"}}`)),
		},
	})
	request := repositories.CreateRepoRequest{Name: "testing"}

	result, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "testing", result.Name)
	assert.EqualValues(t, "fabiosebastiano", result.Owner)

}
