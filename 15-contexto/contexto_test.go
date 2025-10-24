package contexto_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	contexto "tdd/15-contexto"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	cancelled bool
	t *testing.T
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond) // Simula um atraso na busca
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) WasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Errorf("store was not cancelled")
	}
}

func (s *SpyStore) WasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Errorf("store was cancelled")
	}
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

func TestHandler_CancelsStoreFetch(t *testing.T) {
	store := &SpyStore{response: "Hello, World!"}
	server := contexto.Server(store)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	cancellingCtx, cancel := context.WithCancel(request.Context())
	time.AfterFunc(5*time.Millisecond, cancel)
	request = request.WithContext(cancellingCtx)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	if !store.cancelled {
		t.Errorf("store was not cancelled")
	}
}

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

		store.WasNotCancelled()
	})

	t.Run("avisa a store para cancelar:", func(t *testing.T) {
		store := SpyStore{response: data, t: t}
		server := contexto.Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		store.WasCancelled()
	})
}