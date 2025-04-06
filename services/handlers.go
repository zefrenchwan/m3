package services

import (
	"net/http"
	"strings"

	"github.com/zefrenchwan/m3.git/storage"
)

// ServiceContext contains anything an handler may need
type ServiceContext struct {
	// dao to deal with database
	Dao storage.Dao
}

// ServiceError defines a custom error
type ServiceError struct {
	Code       int
	Message    string
	Exceptions []string
}

// IsEmpty returns true if there was no error
func (e ServiceError) IsEmpty() bool {
	return e.Code < 400 || len(e.Message) == 0 || len(e.Exceptions) == 0
}

// FullMessage returns message to display
func (e ServiceError) FullMessage() string {
	// make a slice with all messages
	values := []string{e.Message}
	if len(e.Exceptions) != 0 {
		values = append(values, "\nDue to exceptions: ")
		values = append(values, e.Exceptions...)
	}
	return strings.Join(values, "\n ")
}

// ServiceHandler defines an enriched function to deal with requests
type ServiceHandler func(context ServiceContext, w http.ResponseWriter, r *http.Request) ServiceError

// NewHandler decorates a service handler to use it as a Server mux handler
func NewHandler(context ServiceContext, handler ServiceHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// run and catch error if any
		result := handler(context, w, r)
		if !result.IsEmpty() {
			http.Error(w, result.FullMessage(), result.Code)
		}
	}
}

type ServiceServer struct {
	mux *http.ServeMux
}

func (s ServiceServer) AddServiceHandler(requestType string, requestPath string, context ServiceContext, element ServiceHandler) {
	s.mux.Handle(strings.ToUpper(requestType)+" "+requestPath, NewHandler(context, element))
}
