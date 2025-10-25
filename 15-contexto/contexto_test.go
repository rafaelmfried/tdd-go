package contexto_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	contexto "tdd/15-contexto"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <- ctx.Done():
				s.t.Log("store foi cancelada")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()
	select {
		case <- ctx.Done():
			return "", ctx.Err()
		case res := <- data:
			return res, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("n implementado")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestHandler(t *testing.T) {
	data := "Hello, World!"
	server := contexto.Server(&SpyStore{response: data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf("got %q, want %q", response.Body.String(), data)
	}
}

// func TestHandler_CancelsStoreFetch(t *testing.T) {
// 	store := &SpyStore{response: "Hello, World!"}
// 	server := contexto.Server(store)

// 	request := httptest.NewRequest(http.MethodGet, "/", nil)
// 	cancellingCtx, cancel := context.WithCancel(request.Context())
// 	time.AfterFunc(5*time.Millisecond, cancel)
// 	request = request.WithContext(cancellingCtx)
// 	response := httptest.NewRecorder()

// 	server.ServeHTTP(response, request)

// 	if !store.cancelled {
// 		t.Errorf("store was not cancelled")
// 	}
// }

func TestServer(t *testing.T) {
	data := "dados da store"
	t.Run("retorna dados da store", func(t *testing.T) {
		store := SpyStore{response: data, t: t}
		server := contexto.Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %q, want %q", response.Body.String(), data)
		}
	})

	t.Run("avisa a store para cancelar:", func(t *testing.T) {
		store := SpyStore{response: data, t: t}
		server := contexto.Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := &SpyResponseWriter{}

		server.ServeHTTP(response, request)

		if response.written {
			t.Errorf("a resposta nao deveria ser escrita")
		}

		// store.WasCancelled()
	})
}