package routes

import (
	"github.com/pwestlake/aemo/userservice/pkg/entitlements"
	"github.com/pwestlake/aemo/userservice/pkg/config"
	"github.com/pwestlake/equity-fund/uicontroller/pkg/service"
	"net/http"
	"fmt"

)

// EquityCatalogRoutes ...
// Component for all routes that relate to the entity EquityCatalog
type EquityCatalogRoutes struct {
	service.Route
	config config.Config
	entitlements entitlements.Entitlements
}

// NewEquityCatalogRoutes ...
// Create function for NewEquityCatalogRoutes
func NewEquityCatalogRoutes(entitlements entitlements.Entitlements) EquityCatalogRoutes {
	config := config.NewConfig(nil)
	return EquityCatalogRoutes{config: config, Route: service.NewRoute(entitlements)}
}

// PostEquityCatalogItem ...
// Handler function for the path: /equity-fund/uicontroller/title
func (p *EquityCatalogRoutes) PostEquityCatalogItem(w http.ResponseWriter, r *http.Request) {
	postEquityCatalogItem := func(w http.ResponseWriter, r *http.Request) service.HTTPResponse {
		title, err := p.config.GetString("title")
		if err != nil {
			return &service.HTTPErrorResponse{
				StatusCode: http.StatusInternalServerError, 
				Body: fmt.Sprintf("Unable to read configuration. Error: %s", err.Error())}
		}
		
		return &service.HTTPSuccessResponse{Body: []byte(fmt.Sprintf(`{"value":"%s"}`, title))}
	}

	p.Route.Handle(w, r, postEquityCatalogItem)
}