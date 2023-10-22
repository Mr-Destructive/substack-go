package substack

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var (
	baseURL = "https://substack.com/api/v1"
)

func NewApi(email, password, publicationURL string) (*Api, error) {
	api := &Api{session: &http.Client{}}
	api.publicationURL = urljoin(publicationURL, "api/v1")

	if email == "" && password == "" {
		env, err := loadEnv(".env")
		if err != nil {
			log.Println("Unable to load .env file")
			env = make(map[string]string)
			env["EMAIL"] = ""
			env["PASSWORD"] = ""
		}
		email := env["EMAIL"]
		password := env["PASSWORD"]
		err = api.login(email, password)
		if err != nil {
			return nil, err
		}
	}
	return api, nil
}

func (api *Api) login(email, password string) error {
	response, err := api.session.PostForm(baseURL+"/login", url.Values{
		"email":    {email},
		"password": {password},
	})

	if err != nil {
		return err
	}

	api.cookies = response.Cookies()
	return handleResponse(response, nil)
}

func handleResponse(response *http.Response, result interface{}) error {
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return fmt.Errorf("Substack API Error: %d - %s", response.StatusCode, response.Status)
	}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return err
	}
	return nil
}

func (api *Api) PublicationUsers() (*[]User, error) {
	// Create a new request with stored cookies
	req, err := http.NewRequest("GET", api.publicationURL+"/publication/users", nil)
	if err != nil {
		return nil, err
	}

	for _, cookie := range api.cookies {
		req.AddCookie(cookie)
	}

	response, err := api.session.Do(req)
	if err != nil {
		return nil, err
	}

	var result []User
	err = handleResponse(response, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func GetPublication(publication string) (*Publication, error) {
	publication = urljoin(publication, "api/v1")
	req, err := http.NewRequest("GET", publication+"/publication", nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var result Publication
	err = handleResponse(response, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func GetSubscriptions(publication string) (*SubscriptionsList, error) {
	publication = urljoin(publication, "api/v1")
	req, err := http.NewRequest("GET", publication+"/subscriptions", nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	var result SubscriptionsList
	err = handleResponse(response, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (api *Api) Publication() (*Publication, error) {
	api.publicationURL = urljoin(api.publicationURL, "api/v1")
	return GetPublication(api.publicationURL)
}

func (api *Api) Posts() (*Posts, error) {
	req, err := http.NewRequest("GET", api.publicationURL+"/post_management/published?offset=0&limit=25&order_by=post_date&order_direction=desc", nil)
	if err != nil {
		return nil, err
	}
	for _, cookie := range api.cookies {
		req.AddCookie(cookie)
	}
	response, err := api.session.Do(req)
	if err != nil {
		return nil, err
	}
	result := Posts{}
	err = handleResponse(response, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (api *Api) Drafts() (*Posts, error) {
	req, err := http.NewRequest("GET", api.publicationURL+"/post_management/published?offset=0&limit=25&order_by=post_date&order_direction=desc", nil)
	if err != nil {
		return nil, err
	}
	for _, cookie := range api.cookies {
		req.AddCookie(cookie)
	}
	response, err := api.session.Do(req)
	if err != nil {
		return nil, err
	}
	result := Posts{}
	err = handleResponse(response, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func PublicationPosts(publication string) (*Posts, error) {
	publication = urljoin(publication, "api/v1")
	req, err := http.NewRequest("GET", publication+"/posts", nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	result := Posts{}
	err = handleResponse(response, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func PublicationCategories(publication string) (*[]Category, error) {
	publication = urljoin(publication, "api/v1")
	req, err := http.NewRequest("GET", publication+"/categories", nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	result := []Category{}
	err = handleResponse(response, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (api *Api) CreateDraft(post *Post) error {
	endpoint := fmt.Sprintf("%s/drafts/", api.publicationURL)
	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return err
	}
	for _, cookie := range api.cookies {
		req.AddCookie(cookie)
	}
	req.Header.Add("Content-Type", "application/json")
	response, err := api.session.Do(req)
	if err != nil {
		return err
	}
	err = handleResponse(response, post)
	if err != nil {
		return err
	}
	return nil
}

func (api *Api) UpdateDraft(draftID string, post *Post) error {
	endpoint := fmt.Sprintf("%s/drafts/%s", api.publicationURL, draftID)
    body := &bytes.Buffer{}
    err := json.NewEncoder(body).Encode(post)
    if err != nil {
	req, err := http.NewRequest("PUT", endpoint, body)
	if err != nil {
		return err
	}
	for _, cookie := range api.cookies {
		req.AddCookie(cookie)
	}
	req.Header.Add("Content-Type", "application/json")
	response, err := api.session.Do(req)
	if err != nil {
		return err
	}
	err = handleResponse(response, post)
	if err != nil {
		return err
	}
	return nil
}

func (api *Api) PublishDraft(draftID string, send, shareAutomatically bool) error {
	endpoint := fmt.Sprintf("%s/drafts/%s/publish", api.publicationURL, draftID)
	payload := map[string]interface{}{
		"send":                send,
		"share_automatically": shareAutomatically,
	}
    body := &bytes.Buffer{}
    err := json.NewEncoder(body).Encode(payload)
    if err != nil {
	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		return err
	}
	for _, cookie := range api.cookies {
		req.AddCookie(cookie)
	}
	response, err := api.session.Do(req)
	if err != nil {
		return err
	}
	err = handleResponse(response, nil)
	if err != nil {
		return err
	}
	return nil
}

func (api *Api) DeleteDraft(draftID string) error {
	endpoint := fmt.Sprintf("%s/drafts/%s", api.publicationURL, draftID)
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}
	for _, cookie := range api.cookies {
		req.AddCookie(cookie)
	}
	response, err := api.session.Do(req)
	if err != nil {
		return err
	}
	err = handleResponse(response, nil)
	if err != nil {
		return err
	}
	return nil
}

func urljoin(base, path string) string {
	base = strings.TrimRight(base, "/")
	path = strings.TrimLeft(path, "/")
	return base + "/" + path
}
