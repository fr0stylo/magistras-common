package authentication

import (
	"context"
	"net/http"

	httpclient "github.com/fr0stylo/magistras/common/pkg/servces/httpClient"
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

func init() {
	authHttpClient = httpclient.NewClient("http://auth:8000", &http.Transport{})
}

func GetAuthenticatedUser(ctx context.Context) (*User, error) {
	var user User
	err := authHttpClient.Get(ctx, "/authorize", &user)

	return &user, err
}
