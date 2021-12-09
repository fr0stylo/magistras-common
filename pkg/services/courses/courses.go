package courses

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	httpclient "github.com/fr0stylo/magistras/common/pkg/services/httpClient"
)

var courseClient Client

type Client interface {
	Get(ctx context.Context, url string, response interface{}) error
}

func init() {
	courseClient = httpclient.NewClient("http://courses:8000", &http.Transport{
		MaxIdleConnsPerHost: 1024,
		DisableKeepAlives:   false,
  })
}

func GerCourse(ctx context.Context, id primitive.ObjectID, course interface{}) error {
	err := courseClient.Get(ctx, "/"+id.Hex(), course)

	return err
}
