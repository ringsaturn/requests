package requests_test

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ringsaturn/requests"
)

func ExampleReqWithExpectJSONResponse() {
	type EchoJSONResponse struct {
		ClientIP string            `json:"client_ip"`
		Get      map[string]string `json:"get"`
		Headers  map[string]string `json:"headers"`
		Method   string            `json:"method"`
		Path     string            `json:"path"`
		Post     string            `json:"post"`
		URL      string            `json:"url"`
	}

	client := &http.Client{}
	resp := &EchoJSONResponse{}

	if err := requests.ReqWithExpectJSONResponse(
		context.TODO(), client, "POST", "https://echo.paw.cloud", nil, resp,
		requests.QueryOption(map[string]string{
			"hello": "world",
		}),
		requests.HeaderOption(map[string]string{
			"foo":    "bar",
			"Accept": "application/json",
		}),
	); err != nil {
		panic(err)
	}
	fmt.Println(resp.Method)
	// Output: POST
}
