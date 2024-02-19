package knock

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-cleanhttp"
)

const (
	DefaultBaseUrl = "https://api.knock.app/"
	jsonMediaType  = "application/json"
)

// ErrorCode defines the code of an error.
type ErrorCode string

const (
	ErrInternal          ErrorCode = "internal"           // Internal error.
	ErrInvalid           ErrorCode = "invalid"            // Invalid operation, e.g wrong params
	ErrPermission        ErrorCode = "permission"         // Permission denied.
	ErrNotFound          ErrorCode = "not_found"          // Resource not found.
	ErrRetry             ErrorCode = "retry"              // Operation should be retried.
	ErrResponseMalformed ErrorCode = "response_malformed" // Response body is malformed.
)

// Client encapsulates a client that talks to the Knock API
type Client struct {
	// client represents the HTTP client used for making HTTP requests.
	client *http.Client

	// base URL for the API
	baseURL        *url.URL
	BulkOperations BulkOperationsService
	Messages       MessagesService
	Objects        ObjectsService
	Tenants        TenantsService
	Users          UsersService
	Workflows      WorkflowsService
}

// ClientOption provides a variadic option for configuring the client
type ClientOption func(c *Client) error

// WithBaseURL overrides the base URL for the API.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		parsedURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}

		c.baseURL = parsedURL
		return nil
	}
}

// WithAccessToken configures a client with the given Knock access token.
func WithAccessToken(token string) ClientOption {
	return func(c *Client) error {
		if token == "" {
			return errors.New("missing access token")
		}

		transport := accessTokenTransport{
			rt:    c.client.Transport,
			token: token,
		}

		c.client.Transport = &transport

		return nil
	}
}

// WithHTTPClient configures the Knock client with the given HTTP client.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) error {
		if client == nil {
			client = cleanhttp.DefaultClient()
		}

		c.client = client
		return nil
	}
}

// NewClient instantiates an instance of the Knock API client.
func NewClient(opts ...ClientOption) (*Client, error) {
	baseURL, err := url.Parse(DefaultBaseUrl)
	if err != nil {
		return nil, err
	}

	c := &Client{
		client:  cleanhttp.DefaultClient(),
		baseURL: baseURL,
	}

	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	c.BulkOperations = &bulkOperationsService{client: c}
	c.Messages = &messagesService{client: c}
	c.Objects = &objectsService{client: c}
	c.Tenants = &tenantsService{client: c}
	c.Users = &usersService{client: c}
	c.Workflows = &workflowsService{client: c}

	return c, nil
}

// do makes an HTTP request and populates the given struct v from the response.
func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) ([]byte, error) {
	req = req.WithContext(ctx)
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return c.handleResponse(ctx, res, v)
}

// handleResponse makes an HTTP request and populates the given struct v from
// the response.  This is meant for internal testing and shouldn't be used
// directly. Instead please use `Client.do`.
func (c *Client) handleResponse(ctx context.Context, res *http.Response, v interface{}) ([]byte, error) {
	out, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		// errorResponse represents an error response from the API
		type errorResponse struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		}

		errorRes := &errorResponse{}
		err = json.Unmarshal(out, errorRes)
		if err != nil {
			var jsonErr *json.SyntaxError
			if errors.As(err, &jsonErr) {
				return nil, &Error{
					msg:  "malformed error response body received",
					Code: ErrResponseMalformed,
					Meta: map[string]string{
						"body":        string(out),
						"err":         jsonErr.Error(),
						"http_status": http.StatusText(res.StatusCode),
					},
				}
			}
			return nil, err
		}

		// json.Unmarshal doesn't return an error if the response
		// body has a different protocol then "ErrorResponse". We
		// check here to make sure that errorRes is populated. If
		// not, we return the full response back to the user, so
		// they can debug the issue.
		if *errorRes == (errorResponse{}) {
			return nil, &Error{
				msg:  "internal error, response body doesn't match error type signature",
				Code: ErrInternal,
				Meta: map[string]string{
					"body":        string(out),
					"http_status": http.StatusText(res.StatusCode),
				},
			}
		}

		var errCode ErrorCode
		switch errorRes.Code {
		case "not_found", "resource_missing":
			errCode = ErrNotFound
		case "unauthorized":
			errCode = ErrPermission
		case "invalid_params":
			errCode = ErrInvalid
		case "unprocessable":
			errCode = ErrRetry
		}

		return nil, &Error{
			msg:  errorRes.Message,
			Code: errCode,
		}
	}

	// this means we don't care about un-marshalling the response body into v
	if v == nil {
		return out, nil
	}

	err = json.Unmarshal(out, &v)
	if err != nil {
		var jsonErr *json.SyntaxError
		if errors.As(err, &jsonErr) {
			return nil, &Error{
				msg:  "malformed response body received",
				Code: ErrResponseMalformed,
				Meta: map[string]string{
					"body":        string(out),
					"http_status": http.StatusText(res.StatusCode),
				},
			}
		}
		return nil, err
	}

	return out, nil
}

func (c *Client) newRequest(method string, path string, body interface{}, methodOptions *MethodOptions) (*http.Request, error) {
	u, err := c.baseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	switch method {
	case http.MethodGet:
		req, err = http.NewRequest(method, u.String(), nil)
		if err != nil {
			return nil, err
		}
	default:
		buf := new(bytes.Buffer)
		if body != nil {
			err = json.NewEncoder(buf).Encode(body)
			if err != nil {
				return nil, err
			}
		}

		req, err = http.NewRequest(method, u.String(), buf)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Content-Type", jsonMediaType)
	}

	req.Header.Set("Accept", jsonMediaType)

	if methodOptions != nil {
		if methodOptions.IdempotencyKey != "" {
			req.Header.Set("Idempotency-Key", methodOptions.IdempotencyKey)
		}
	}

	return req, nil
}

type accessTokenTransport struct {
	rt    http.RoundTripper
	token string
}

func (t *accessTokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+t.token)
	return t.rt.RoundTrip(req)
}

// Error represents common errors originating from the Client.
type Error struct {
	// msg contains the human readable string
	msg string

	// Code specifies the error code. i.e; NotFound, RateLimited, etc...
	Code ErrorCode

	// Meta contains additional information depending on the error code. As an
	// example, if the Code is "ErrResponseMalformed", the map will be: ["body"]
	// = "body of the response"
	Meta map[string]string
}

// Error returns the string representation of the error.
func (e *Error) Error() string { return e.msg }

type MethodOptions struct {
	IdempotencyKey string
}
