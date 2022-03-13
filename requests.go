package requests

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type RequestOption func(r *http.Request)

func HeaderOption(header map[string]string) RequestOption {
	return func(r *http.Request) {
		for headerKey, headerValue := range header {
			r.Header.Add(headerKey, headerValue)
		}
	}
}

func QueryOption(query map[string]string) RequestOption {
	return func(r *http.Request) {
		reqQuery := r.URL.Query()
		for headerKey, headerValue := range query {
			reqQuery.Add(headerKey, headerValue)
		}
		r.URL.RawQuery = reqQuery.Encode()
	}
}

// Build combine request params
func Build(
	ctx context.Context,
	client *http.Client,
	method string,
	url string, body io.Reader,
	options ...RequestOption,
) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	for _, option := range options {
		option(req)
	}
	return req, nil
}

// Req will do HTTP request
func Req(
	ctx context.Context,
	client *http.Client,
	method string,
	url string, body io.Reader,
	options ...RequestOption,
) (*http.Response, error) {
	req, err := Build(ctx, client, method, url, body, options...)
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}

// UnmarshalFunc is type define for any unmarshall funct
type UnmarshalFunc func(data []byte, v interface{}) error

// ReqWithCustomUnmarshal will use input func to unmarshall response
func ReqWithCustomUnmarshal(
	ctx context.Context,
	client *http.Client,
	method string,
	url string, body io.Reader,
	unmarshalFunc UnmarshalFunc, expect interface{},
	options ...RequestOption,
) error {
	resp, err := Req(ctx, client, method, url, body, options...)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return unmarshalFunc(respBody, expect)
}

// ReqWithExpectJSONResponse use json.Unmarshal to unmarshall response
func ReqWithExpectJSONResponse(
	ctx context.Context,
	client *http.Client,
	method string,
	url string, body io.Reader,
	expect interface{},
	options ...RequestOption,
) error {
	return ReqWithCustomUnmarshal(ctx, client, method, url, body, json.Unmarshal, expect, options...)
}
