package piston

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	HttpClient *http.Client
	BaseURL    string
	ApiKey     string
}

func NewClient(client *http.Client, baseURL string, apiKey string) *Client {
	return &Client{
		client,
		baseURL,
		apiKey,
	}
}

func (client *Client) GetRuntimes() ([]Runtime, error) {
	resp, err := client.makeRequest("GET", "http://"+client.BaseURL+"/api/v2/runtimes", nil)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	runtimes := []Runtime{}
	if err = json.Unmarshal(b, &runtimes); err != nil {
		return nil, err
	}

	return runtimes, nil
}

func (client *Client) GetInstalledPackages() ([]Package, error) {
	resp, err := client.makeRequest("GET", "http://"+client.BaseURL+"/api/v2/runtimes", nil)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	pistonPackages := []Package{}
	if err = json.Unmarshal(b, &pistonPackages); err != nil {
		return nil, err
	}

	return pistonPackages, nil
}

func (client *Client) InstallPackage(p Package) error {
	b, err := json.Marshal(p)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(b)
	resp, err := client.makeRequest("POST", "http://"+client.BaseURL+"/api/v2/packages", reader)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	errorResponse := struct {
		Message string `json:"message"`
	}{}

	if errorResponse.Message == "Already installed" {
		return nil
	}

	// installPP := Package{}
	// json.Unmarshal(b, &installPP)
	// fmt.Println(installPP)

	return nil
}

func (client *Client) Execute(task *ExecutionTask) (*ExecutionResult, error) {
	b, err := json.Marshal(task)
	if err != nil {
		return nil, err
	}

	resp, err := client.makeRequest("POST", "http://"+client.BaseURL+"/api/v2/execute", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := ExecutionResult{}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (client *Client) makeRequest(method string, url string, body io.Reader) (*http.Response, error) {
	if body == nil {
		body = &bytes.Reader{}
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	if apiKey := client.ApiKey; apiKey != "" {
		req.Header.Add("Authorization", apiKey)
	}

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	resp.Body.Close()
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(respBody))

	err = handleStatusCode(resp.StatusCode, string(respBody))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func handleStatusCode(code int, respBody string) error {

	var err error

	if code < 300 && code >= 200 {
		return nil
	}

	switch code {
	case http.StatusTooManyRequests:
		err = errors.New("You have been ratelimited.Try again later")
	case http.StatusInternalServerError:
		err = errors.New("Server failed to respond. Try again later")
	case http.StatusBadRequest:
		err = errors.New("Invalid Request. The language or version may be incorrect.")
	case http.StatusNotFound:
		err = errors.New("Not found." + respBody)
	default:
		err = errors.New("Unexpected Error. " + respBody)
	}

	return err
}
