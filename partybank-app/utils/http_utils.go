package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type HttpClient[R, A any] struct {
	logger *log.Logger
	client *http.Client
}

func NewHttpClient[R, A any](logger *log.Logger, client *http.Client) *HttpClient[R, A] {
	return &HttpClient[R, A]{
		logger: logger,
		client: client,
	}
}

func (httpClient *HttpClient[R, A]) Send(method, url string, requestBody *R, headers map[string]string) (*A, error) {
	if requestBody != nil {
		resp, err := httpClient.sendRequestWithBody(method, url, requestBody)
		if err != nil {
			log.Printf("[Error:\t%v]", err)
			return nil, err
		}
		return httpClient.extractResponseFrom(resp)
	}
	req, err := http.NewRequest(method, url, bytes.NewReader([]byte("")))
	if err != nil {
		httpClient.logger.Printf("[Error:\t%v]", err)
		return nil, errors.New("error creating request object")
	}
	fmt.Println("request: ", *req)
	if headers != nil {
		httpClient.addHeadersTo(req, headers)
	}
	log.Println("header: ", req.Header)
	res, err := httpClient.client.Do(req)
	if err != nil || res.StatusCode != 200 {
		log.Println("responsee: ", *res)
		log.Printf("[Error:\t%v]", err)
		return nil, errors.New("error performing request")
	}
	return httpClient.extractResponseFrom(res)
}

func (httpClient *HttpClient[R, A]) extractResponseFrom(resp *http.Response) (*A, error) {
	log.Println("responsee: ", *resp)
	res := new(A)
	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(&res)
	if err != nil {
		log.Printf("[Error:\t%v]", err)
		return nil, err
	}
	return res, nil
}

func (httpClient *HttpClient[R, A]) sendRequestWithBody(method string, url string, body *R) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		httpClient.logger.Printf("[Error:\t%v]", err)
		return nil, errors.New("error serializing request object")
	}
	fmt.Println("json data: ", string(jsonBody))
	reqBody := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		httpClient.logger.Printf("[Error:\t%v]", err)
		return nil, errors.New("error creating request object")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	fmt.Println("request: ", req)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error: ", err)
		//httpClient.logger.Printf("[Error:\t%v]", err)
		return nil, errors.New("error performing request")
	}
	fmt.Println("response: ", *res)
	return res, nil
}

func (httpClient *HttpClient[R, A]) addHeadersTo(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Add(key, value)
	}
}
