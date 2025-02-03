package knock

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Run("default client", func(t *testing.T) {
		client, err := NewClient()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if client == nil {
			t.Fatal("Expected non-nil client")
		}
		if client.baseURL.String() != DefaultBaseUrl {
			t.Errorf("Expected base URL %s, got %s", DefaultBaseUrl, client.baseURL.String())
		}
	})

	t.Run("with custom base URL", func(t *testing.T) {
		customURL := "https://custom.knock.app/"
		client, err := NewClient(WithBaseURL(customURL))
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if client == nil {
			t.Fatal("Expected non-nil client")
		}
		if client.baseURL.String() != customURL {
			t.Errorf("Expected base URL %s, got %s", customURL, client.baseURL.String())
		}
	})
}

func TestClientDo(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !strings.Contains(r.Header.Get("User-Agent"), "Knock/Go v") {
				t.Errorf("Expected User-Agent to contain 'Knock/Go v', got %s", r.Header.Get("User-Agent"))
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(map[string]string{"message": "success"})
			if err != nil {
				t.Fatalf("Error encoding httptest response: %v", err)
			}
		}))
		defer server.Close()

		client, err := NewClient(WithBaseURL(server.URL))
		if err != nil {
			t.Fatalf("Expected no error creating client, got %v", err)
		}
		req, err := client.newRequest("GET", "/test", nil, nil)
		if err != nil {
			t.Fatalf("Expected no error creating request, got %v", err)
		}

		var response map[string]string
		_, err := client.do(context.Background(), req, &response)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if response["message"] != "success" {
			t.Errorf("Expected message 'success', got %s", response["message"])
		}
	})

	t.Run("error response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			err := json.NewEncoder(w).Encode(map[string]string{"code": "not_found", "message": "Resource not found"})
			if err != nil {
				t.Fatalf("Error encoding httptest response: %v", err)
			}
		}))
		defer server.Close()

		client, err := NewClient(WithBaseURL(server.URL))
		if err != nil {
			t.Fatalf("Expected no error creating client, got %v", err)
		}
		req, err := client.newRequest("GET", "/test", nil, nil)
		if err != nil {
			t.Fatalf("Expected no error creating request, got %v", err)
		}

		_, err := client.do(context.Background(), req, nil)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}
		knockErr, ok := err.(*Error)
		if !ok {
			t.Fatalf("Expected error of type *Error, got %T", err)
		}
		if knockErr.Code != ErrNotFound {
			t.Errorf("Expected error code %v, got %v", ErrNotFound, knockErr.Code)
		}
		if knockErr.Error() != "Resource not found" {
			t.Errorf("Expected error message 'Resource not found', got %s", knockErr.Error())
		}
	})

	t.Run("unexpected non-json response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadGateway)
			_, err := w.Write([]byte("502 Bad Gateway"))
			if err != nil {
				t.Fatalf("Error writing httptest response: %v", err)
			}
		}))
		defer server.Close()

		client, err := NewClient(WithBaseURL(server.URL))
		if err != nil {
			t.Fatalf("Expected no error creating client, got %v", err)
		}
		req, err := client.newRequest("GET", "/test", nil, nil)
		if err != nil {
			t.Fatalf("Expected no error creating request, got %v", err)
		}

		_, err := client.do(context.Background(), req, nil)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}
		knockErr, ok := err.(*Error)
		if !ok {
			t.Fatalf("Expected error of type *Error, got %T", err)
		}
		if knockErr.Code != ErrResponseMalformed {
			t.Errorf("Expected error code %v, got %v", ErrNotFound, knockErr.Code)
		}
		if knockErr.Error() != "malformed non-json error response body received with status code: Bad Gateway" {
			t.Errorf("Expected error message 'Resource not found', got %s", knockErr.Error())
		}
	})
}
