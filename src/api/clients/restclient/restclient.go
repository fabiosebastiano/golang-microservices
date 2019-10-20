package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	enabledMocks = false
	mocks        = make(map[string]*Mock)
)

//struct che imita una response http
type Mock struct {
	Url        string
	HttpMethod string
	Response   *http.Response
	Err        error
}

//Funzione per abilitare il mockup
func StartMockups() {
	enabledMocks = true
}

//Funzione per disabilitare il mockup
func StopMockups() {
	enabledMocks = false
}

func FlushMocks() {
	mocks = make(map[string]*Mock)
}

func GetMockId(httpMethod string, url string) string {
	return fmt.Sprintf("%s_%s", httpMethod, url)
}

// funzione da chiamare quando si prepara il test case
// con il tipo di risposta che vogliamo testare
func AddMockup(mock Mock) {
	mocks[GetMockId(mock.HttpMethod, mock.Url)] = &mock
}

//Post Basic POST CALL
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	//entra qui solo in caso di MOCKUP, altrimenti salta
	if enabledMocks {
		mock := mocks[GetMockId(http.MethodPost, url)]
		if mock == nil {
			return nil, errors.New("no mockup found for given request")
		}

		return mock.Response, mock.Err
	}

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := http.Client{}
	return client.Do(request)
}
