# Small requests helper for Go

Demo:

```go
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ringsaturn/requests"
)

type EchoJSONResponse struct {
	ClientIP string            `json:"client_ip"`
	Get      map[string]string `json:"get"`
	Headers  map[string]string `json:"headers"`
	Method   string            `json:"method"`
	Path     string            `json:"path"`
	Post     string            `json:"post"`
	URL      string            `json:"url"`
}

func main() {
	client := &http.Client{}
	resp := &EchoJSONResponse{}
	err := requests.ReqWithExpectJSONResponse(
		context.TODO(), client, "GET", "https://echo.paw.cloud", nil, resp,
		requests.QueryOption(map[string]string{
			"hello": "world",
		}),
		requests.HeaderOption(map[string]string{
			"foo":    "bar",
			"Accept": "application/json",
		}),
	)
	if err != nil {
		panic(err)
	}
	log.Printf("%+v\n", resp)
}
```
