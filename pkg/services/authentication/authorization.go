package authentication

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	httpclient "github.com/fr0stylo/magistras/common/pkg/internal/httpClient"
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
	caCertPool := x509.NewCertPool()
	caCert, err := ioutil.ReadFile(os.Getenv("CA_PEM_PATH"))
	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
	}

	caCertPool.AppendCertsFromPEM(caCert)

	authHttpClient = httpclient.NewClient("https://auth:8000", &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: caCertPool,
		},
	})
}

func GetAuthenticatedUser(ctx context.Context) (*User, error) {
	var user User
	err := authHttpClient.Get(ctx, "/authorize", &user)

	return &user, err
}
