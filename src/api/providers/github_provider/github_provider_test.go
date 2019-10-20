package github_provider

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/fabiosebastiano/golang-microservices/src/api/clients/restclient"
	"github.com/fabiosebastiano/golang-microservices/src/api/domain/github"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.NotNil(t, header)
	assert.EqualValues(t, "token abc123", header)

}

func TestCreateRepoErrorRestClient(t *testing.T) {
	//restclient.StartMockups()
	//defer restclient.StopMockups()
	restclient.FlushMocks()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		//	Response:   nil, //essendo nil possiamo anche NON metterlo
		Err: errors.New("invalid restclient response"),
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid restclient response", err.Message)

}
func TestCreateRepoInvalidResponseBody(t *testing.T) {
	restclient.FlushMocks()
	invalidCloser, _ := os.Open("- asf3")
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       invalidCloser,
		},
	},
	)
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "Invalid response body", err.Message)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
}
func TestCreateInvalidErrorInterface(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader("{message:1}")),
		},
	},
	)
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "Invalid json response body", err.Message)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
}

func TestCreateRepoUnauthorized(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":"Requires authentication", "documentation_url":""}`)),
		},
	},
	)
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "Requires authentication", err.Message)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
}
