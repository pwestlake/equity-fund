package routes

import (
	commons "github.com/pwestlake/equity-fund/commons/pkg/service"
	"github.com/pwestlake/aemo/userservice/pkg/entitlements"
	"fmt"
	"github.com/pwestlake/aemo/userservice/pkg/config"
	"net/http"
)

// ContextRoutes ...
// Component for all routes that relate to the application context
type ContextRoutes struct {
	commons.Route
	config config.Config
	entitlements entitlements.Entitlements
}

// NewContextRoutes ...
// Create function for ContextRoutes
func NewContextRoutes(entitlements entitlements.Entitlements) ContextRoutes {
	config := config.NewConfig(nil)
	return ContextRoutes{config: config, Route: commons.NewRoute(entitlements)}
}

// GetTitle ...
// Handler function for the path: /equity-fund/uicontroller/title
func (p *ContextRoutes) GetTitle(w http.ResponseWriter, r *http.Request) {
	getTitle := func(w http.ResponseWriter, r *http.Request) commons.HTTPResponse {
		title, err := p.config.GetString("title")
		if err != nil {
			return &commons.HTTPErrorResponse{
				StatusCode: http.StatusInternalServerError, 
				Body: fmt.Sprintf("Unable to read configuration. Error: %s", err.Error())}
		}
		
		return &commons.HTTPSuccessResponse{Body: []byte(fmt.Sprintf(`{"value":"%s"}`, title))}
	}

	p.Route.Handle(w, r, getTitle)
}

