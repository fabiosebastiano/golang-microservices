package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/fabiosebastiano/golang-microservices/src/api/clients/restclient"
	"github.com/fabiosebastiano/golang-microservices/src/api/domain/repositories"
	"github.com/fabiosebastiano/golang-microservices/src/api/utils/errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidJsonError(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(``))

	c.Request = request
	CreateRepo(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	fmt.Println(response.Body.String())
	apiError, err := errors.NewApiErrorFromBytes(response.Body.Bytes())

	assert.Nil(t, err)
	assert.NotNil(t, apiError)
	assert.EqualValues(t, http.StatusBadRequest, apiError.Status())
	assert.EqualValues(t, "invalid json body", apiError.Message())

}

func TestCreateRepoGithubError(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name":"testing"}`))

	c.Request = request
	restclient.FlushMocks()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":"Requires authentication", "documentation_url":"http://developer.github."}`)),
		},
	})

	CreateRepo(c)

	assert.EqualValues(t, http.StatusUnauthorized, response.Code)
	fmt.Println(response.Body.String())
	apiError, err := errors.NewApiErrorFromBytes(response.Body.Bytes())

	assert.Nil(t, err)
	assert.NotNil(t, apiError)
	assert.EqualValues(t, http.StatusUnauthorized, apiError.Status())
	assert.EqualValues(t, "Requires authentication", apiError.Message())

}
func TestCreateRepoNoError(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name":"testing"}`))

	c.Request = request
	restclient.FlushMocks()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":123}`)),
		},
	})

	CreateRepo(c)

	assert.EqualValues(t, http.StatusCreated, response.Code)
	var result repositories.CreateRepoResponse

	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "", result.Name)
	assert.EqualValues(t, "", result.Owner)
}
