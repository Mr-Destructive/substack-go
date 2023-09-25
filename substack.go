package substack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var (
	baseURL = "https://substack.com/api/v1"
)

type Api struct {
	session        *http.Client
	publicationURL string
	cookies        []*http.Cookie
}

type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Handle         string `json:"handle"`
	PreviousName   string `json:"previous_name"`
	PhotoURL       string `json:"photo_url"`
	Bio            string `json:"bio"`
	ProfileSetUpAt string `json:"profile_set_up_at"`
}

type SubscribeContent struct {
	Type    string `json:"type"`
	Content []struct {
		Type    string `json:"type"`
		Content string `json:"content"`
	}
}

type Publication struct {
	Subdomain           string `json:"subdomain"`
	Name                string `json:"name"`
	CustomDomain        string `json:"custom_domain"`
	LogoURL             string `json:"logo_url"`
	LogoURLWide         string `json:"logo_url_wide"`
	CoverPhotoURL       string `json:"cover_photo_url"`
	Copyright           string `json:"copyright"`
	EmailFromName       string `json:"email_from_name"`
	SubscribeContent    string `json:"subscribe_content"`
	SubscribeFooter     string `json:"subscribe_footer"`
	WelcomeEmailContent string `json:"welcome_email_content"`
}

func NewApi(email, password, publicationURL string) (*Api, error) {
	api := &Api{session: &http.Client{}}
	api.publicationURL = urljoin(publicationURL, "api/v1")

	if email != "" && password != "" {
		err := api.login(email, password)
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
	return api.handleResponse(response, nil)
}

func (api *Api) handleResponse(response *http.Response, result interface{}) error {
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return fmt.Errorf("Substack API Error: %d - %s", response.StatusCode, response.Status)
	}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return err
	}
	return nil
}

func (api *Api) getPublicationUsers() ([]User, error) {
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
	err = api.handleResponse(response, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (api *Api) getPublication() (*Publication, error) {
	req, err := http.NewRequest("GET", api.publicationURL+"/publication", nil)
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

	var result Publication
	err = api.handleResponse(response, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func urljoin(base, path string) string {
	base = strings.TrimRight(base, "/")
	path = strings.TrimLeft(path, "/")
	return base + "/" + path
}
