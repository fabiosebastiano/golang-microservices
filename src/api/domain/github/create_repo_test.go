package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "golang introduction",
		Description: "a golang introduction repository",
		Homepage:    "https://github.com",
		Private:     false,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}

	//Marshal prende in input un'interface e cerca di trasformarla in un JSON in formato string
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	//Unmarshal prende in input un array di byte ed un pointer ad una struct che vogliamo popolato con il json
	//creiamo un nuovo oggetto di CreateRepo... e lo popoliamo con l'array di byte ottenuto dal precedente JSON
	var target CreateRepoRequest
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)

	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.HasIssues, request.HasIssues)

	//assert.EqualValues(t, `{"name":"golang introduction","description":"a golang introduction repository","homepage":"https://github.com","private":false,"has_issues":true,"has_project":true,"has_wiki":true}`, string(bytes))

}
