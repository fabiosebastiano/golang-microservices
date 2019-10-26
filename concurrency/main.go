package main

import (
	"bufio"
	"fmt"
	"os"
"sync"
	"github.com/fabiosebastiano/golang-microservices/src/api/domain/repositories"
	"github.com/fabiosebastiano/golang-microservices/src/api/services"
	"github.com/fabiosebastiano/golang-microservices/src/api/utils/errors"
)

var (
	successes map[string]string
	failed    map[string]errors.ApiError
)

type createRepoResult struct {
	Request repositories.CreateRepoRequest
	Result  *repositories.CreateRepoResponse
	Error   errors.ApiError
}

func getRequests() []repositories.CreateRepoRequest {
	result := make([]repositories.CreateRepoRequest, 0)

	file, err := os.Open("//Users/fabioSebastiano/Desktop/repositories.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		request := repositories.CreateRepoRequest{
			Name: line,
		}
		result = append(result, request)
	}
	return result
}

func main() {
	requests := getRequests()
	fmt.Printf("About to process %d requests", len(requests))

	input := make(chan createRepoResult)
	buffer := make(chan bool, 10)
	var wg sync.WaitGroup

	go handleResult(input, &wg)

	for _, request := range requests {
		buffer <- true
		wg.Add(1)
		go createRepo(buffer, input, request)
	}


	wg.Wait()
	close(input)
}

func handleResult(input chan createRepoResult, wg *sync.WaitGroup) {
	for result := range input {
		if result.Error != nil {
			failed[""] = result.Error
		}else {
			successes[""] = result.Result.Name
		}
		wg.Done()
	}

}

func createRepo(buffer chan bool, output chan createRepoResult, request repositories.CreateRepoRequest) {

	result, error := services.RepositoryService.CreateRepo(request)

	output <- createRepoResult{
		Request: request,
		Result:  result,
		Error:   error,
	}
	
	<-buffer

}
