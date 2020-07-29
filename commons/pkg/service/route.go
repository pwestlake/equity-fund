package service

import (
	"net/http"
	"fmt"
	"log"
	"time"
	"runtime"
	"github.com/pwestlake/aemo/userservice/pkg/entitlements"
)

// Route ...
// Generic functionality for a route
type Route struct {
	entitlements entitlements.Entitlements
}

// NewRoute ...
// Creates a new Route
func NewRoute(entitlements entitlements.Entitlements) Route {
	return Route{entitlements: entitlements}
}

// HTTPResponse ...
// an interface for all HTTP responses
type HTTPResponse interface {
	GetContentType() string
	GetStatusCode() int
	GetBody() []byte
}

// HTTPErrorResponse ...
// an HTTP error response
type HTTPErrorResponse struct {
	StatusCode int
	Body string
}

// HTTPSuccessResponse ...
// an HTTP response
type HTTPSuccessResponse struct {
	Body []byte
}

// GetContentType ...
func (r *HTTPErrorResponse) GetContentType() string {
	return "text/html"
}

// GetContentType ...
func (r *HTTPSuccessResponse) GetContentType() string { 
	return "application/json"
}

// GetStatusCode ...
func (r *HTTPErrorResponse) GetStatusCode() int {
	return r.StatusCode
}

// GetStatusCode ...
func (r *HTTPSuccessResponse) GetStatusCode() int {
	return http.StatusOK
}

// GetBody ...
func (r *HTTPErrorResponse) GetBody() []byte {
	return []byte(r.Body)
}

// GetBody ...
func (r *HTTPSuccessResponse) GetBody() []byte {
	return r.Body
}

// Handle ...
// Wrapper function to respond to an http request
func (p *Route) Handle(w http.ResponseWriter, r *http.Request,
	fn func(http.ResponseWriter, *http.Request) HTTPResponse) {
	startTime, err := p.entry(w, r)
	if err == nil {
		response := fn(w, r)

		if response.GetStatusCode() != http.StatusOK {
			log.Printf("Status code: %d. %s", response.GetStatusCode(), response.GetBody())
		}
		w.Header().Set("Content-Type", response.GetContentType())
		w.WriteHeader(response.GetStatusCode())
		_, err = w.Write(response.GetBody())
		if err != nil {
			log.Printf("%s", err.Error())
		}
	} else {
		log.Printf("%s", err.Error())
	}

	p.exit(startTime)
}

func (p *Route) entry(w http.ResponseWriter, r *http.Request) (time.Time, error) {
	pc, _, _, _ := runtime.Caller(2)
	log.Printf("%s Entered.", runtime.FuncForPC(pc).Name())
	p.entitlements.LogUser(fmt.Sprintf("%s called by: ", r.URL.EscapedPath()), r)
	now := time.Now()

	ent, err := p.entitlements.IsAuthorized("Admin", r)
	if !ent {
		p.writeResponse(w, err.(*entitlements.AuthorizationError).Code,
			[]byte(fmt.Sprintf(`{"message": "%s"}`,
				err.(*entitlements.AuthorizationError).Msg)))
		return now, err.(*entitlements.AuthorizationError)
	}

	return now, nil
}

func (p *Route) exit(startTime time.Time) {
	pc, _, _, _ := runtime.Caller(2)
	log.Printf("%s Exited after %v", runtime.FuncForPC(pc).Name(), time.Now().Sub(startTime))
}

func (p *Route) writeResponse(w http.ResponseWriter, code int, body []byte) {
	if code != http.StatusOK {
		log.Printf("Status code: %d. %s", code, string(body))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(body)
}