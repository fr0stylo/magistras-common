package authentication

import (
	"context"
	"net/http"
	"strings"

	httpclient "github.com/fr0stylo/magistras/common/pkg/services/httpClient"
)

var authHttpClient Client

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ClassName string `json:"className"`
	Role      int    `json:"role"`
}

type Client interface {
	Get(ctx context.Context, url string, response interface{}) error
}

func NewClient() {
	authHttpClient = httpclient.NewClient("http://auth:8000", &http.Transport{})
}

func GetAuthenticatedUser(ctx context.Context) (*User, error) {
	if authHttpClient != nil {
		NewClient()
	}

	var user User
	err := authHttpClient.Get(ctx, "/authorize", &user)

	return &user, err
}

func GetUsersByIds(ctx context.Context, ids []string) (*map[string]interface{}, error) {
	if authHttpClient != nil {
		NewClient()
	}

	var user map[string]interface{}
	queryIds := strings.Join(ids, ",")
	err := authHttpClient.Get(ctx, "/users?ids="+queryIds, &user)

	return &user, err
}
