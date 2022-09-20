package web

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/alpine-hodler/gidari/internal/web/auth"
	"golang.org/x/time/rate"
)

// Client is a wrapper around the http.Client that will handle authentication and rate limiting.
type Client struct{ http.Client }

// NewClient will return a new client with the given options.
func NewClient(_ context.Context, roundtripper auth.Transport) (*Client, error) {
	c := new(Client)
	c.Client.Transport = roundtripper
	return c, nil
}

// newHTTPRequest will return a new request.  If the options are set, this function will encode a body if possible.
func newHTTPRequest(method string, u *url.URL) (*http.Request, error) {
	return http.NewRequest(method, u.String(), nil)
}

// parseErrorMessage takes a response and a status and builder an error message to send to the server.
func parseErrorMessage(resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return fmt.Errorf("Status Code %v (%v): %v", resp.StatusCode, resp.Status, string(body))
}

// validateResponse is a switch condition that parses an error response
func validateResponse(res *http.Response) (err error) {
	if res == nil {
		err = fmt.Errorf("no response, check request and env file")
	} else {
		switch res.StatusCode {
		case
			http.StatusBadRequest,
			http.StatusUnauthorized,
			http.StatusInternalServerError,
			http.StatusNotFound,
			http.StatusTooManyRequests,
			http.StatusForbidden:
			err = parseErrorMessage(res)
		}
	}
	return
}

type FetchConfig struct {
	C           *Client
	Method      string
	URL         *url.URL
	RateLimiter *rate.Limiter
}

func (cfg *FetchConfig) validate() error {
	wrapper := func(field string) error { return fmt.Errorf("%q is a required field on web.FetchConfig", field) }
	if cfg.C == nil {
		return wrapper("Client")
	}
	if cfg.Method == "" {
		return wrapper("Method")
	}
	if cfg.URL == nil {
		return wrapper("URL")
	}
	if cfg.RateLimiter == nil {
		return wrapper("RateLimiter")
	}
	return nil
}

var ratelimiter *rate.Limiter

func init() {
	ratelimiter = rate.NewLimiter(rate.Every(1*time.Second), 3)
}

// FetchResponse is a wrapper for the Fetch function's response data for an HTTP web request.
type FetchResponse struct {
	// Request is the request that was made to the server.
	Request *http.Request

	// Body is the response body from the server.
	Body io.ReadCloser
}

func newFetchResponse(req *http.Request, body io.ReadCloser) *FetchResponse {
	return &FetchResponse{
		Request: req,
		Body:    body,
	}
}

// Fetch will make an HTTP request using the underlying client and endpoint.
func Fetch(ctx context.Context, cfg *FetchConfig) (*FetchResponse, error) {
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	// If the rate limiter is not set, set it with defaults.
	if err := ratelimiter.Wait(ctx); err != nil {
		return nil, err
	}

	req, err := newHTTPRequest(cfg.Method, cfg.URL)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, fmt.Errorf("rate limiter timeout: %v", err)
	}

	rsp, err := cfg.C.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}

	if err := validateResponse(rsp); err != nil {
		rsp.Body.Close()
		return nil, err
	}

	return newFetchResponse(req, rsp.Body), nil
}
