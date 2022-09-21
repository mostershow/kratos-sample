package conn

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewHttpClient() *http.Client {
	conn, err := http.NewClient(
		context.Background(),
		http.WithEndpoint("127.0.0.1:8000"),
	)
	if err != nil {
		panic(err)
	}
	return conn

}
