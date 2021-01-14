package envoyer

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const ApiUrl = "https://envoyer.io/api"

type Client struct {
	httpClient *http.Client
	apiKey string
}

func (c Client) GetProjects() (*ProjectsResponse, error) {
	req, err := c.prepareRequest("GET", "projects", nil)
	if err != nil {
		return nil, err
	}
	var pr ProjectsResponse
	err = makeRequest(c.httpClient, req, &pr)
	if err != nil {
		return nil, err
	}
	return &pr, nil
}

func (c Client) GetProject(id int) (*ProjectResponse, error) {
	req, err := c.prepareRequest("GET", fmt.Sprintf("projects/%d", id), nil)
	if err != nil {
		return nil, err
	}
	var pr ProjectResponse
	err = makeRequest(c.httpClient, req, &pr)
	if err != nil {
		return nil, err
	}
	return &pr, nil
}

func (c Client) GetCollaborators(id int) (*CollaboratorsResponse, error) {
	req, err := c.prepareRequest("GET", fmt.Sprintf("projects/%d/collaborators", id), nil)
	if err != nil {
		return nil, err
	}
	var cr CollaboratorsResponse
	err = makeRequest(c.httpClient, req, &cr)
	if err != nil {
		return nil, err
	}
	return &cr, nil
}

func (c Client) CreateCollaborator(id int, email string) (*CollaboratorsResponse, error) {
	postData := strings.NewReader(fmt.Sprintf(`{"email":"%s"}`, email))
	req, err := c.prepareRequest("POST", fmt.Sprintf("projects/%d/collaborators", id), postData)
	if err != nil {
		return nil, err
	}
	var cr CollaboratorsResponse
	err = makeRequest(c.httpClient, req, &cr)
	return &cr, err
}

func (c Client) DeleteCollaborator(id int, email string) error {
	postData := strings.NewReader(fmt.Sprintf(`{"email":"%s"}`, email))
	req, err := c.prepareRequest("DELETE", fmt.Sprintf("projects/%d/collaborators", id), postData)
	if err != nil {
		return err
	}
	var cr CollaboratorsResponse
	err = makeRequest(c.httpClient, req, &cr)
	return err
}

func makeRequest(client *http.Client, req *http.Request, obj interface{}) error {
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	if res.StatusCode != http.StatusOK {
		_, _ = io.Copy(ioutil.Discard, res.Body)
		return fmt.Errorf(res.Status)
	}
	err = json.NewDecoder(res.Body).Decode(obj)
	if err != nil {
		return err
	}
	return nil
}

/**
 * curl -H "Content-Type: application/json" -H "Accept: application/json" \
-H "Authorization: Bearer ***REMOVED***" \
https://envoyer.io/api/projects
 */
func (c Client) prepareRequest(method, path string, body io.Reader) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s", ApiUrl, path)
	req, err := http.NewRequest(method, url, body)
	if err !=  nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer " + c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func NewClient(apiKey string) *Client {
	client := &Client{
		apiKey: apiKey,
		httpClient: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:           3,
				MaxIdleConnsPerHost:    3,
				MaxConnsPerHost:        3,
			},
		},
	}
	return client
}
