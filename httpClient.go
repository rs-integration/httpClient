package httpClient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	HTTP_GET    = "GET"
	HTTP_POST   = "POST"
	HTTP_PUT    = "PUT"
	HTTP_DELETE = "DELETE"
)

var (
	httpMethods = []string{HTTP_GET, HTTP_POST, HTTP_PUT, HTTP_DELETE}
	basicResult = requestResult{
		code:    0,
		content: "",
	}
)

type requestResult struct {
	code    int
	content string
}

type HttpClient struct {
	method  string
	url     string
	data    string
	headers string
	result  requestResult
}

func NewHttpClient(method string, url string, data string, headers string) *HttpClient {
	return &HttpClient{method, url, data, headers, basicResult}
}

func (client HttpClient) Request() error {
	err := client.validateParams()
	if err != nil {
		return err
	}

	httpClient := http.Client{}
	request, err := http.NewRequest(client.method, client.url, nil)
	if err != nil {
		return err
	}

	resp, err := httpClient.Do(request)
	if err != nil {
		return err
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Printf("%#v\n", result)

	return nil
}

func (client HttpClient) validateParams() error {
	checkErrors := [2]error{client.checkMethod(), client.checkUrl()}
	errorMessage := ""

	for _, checkError := range checkErrors {
		if checkError != nil {
			errorMessage += checkError.Error()
		}
	}

	if errorMessage != "" {
		return errors.New(errorMessage)
	}

	return nil
}

func (client HttpClient) checkMethod() error {
	if !arrayHasString(httpMethods, client.method) {
		return errors.New("declared http method not recognized")
	}

	return nil
}

func (client HttpClient) checkUrl() error {
	return nil
}

func arrayHasString(array []string, expected string) bool {
	for _, actual := range array {
		if expected == actual {
			return true
		}
	}

	return false
}
