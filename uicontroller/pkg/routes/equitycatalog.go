package routes

import (
	"github.com/pwestlake/equity-fund/uicontroller/pkg/domain"
	"github.com/pwestlake/aemo/userservice/pkg/entitlements"
	"github.com/pwestlake/aemo/userservice/pkg/config"
	"github.com/pwestlake/equity-fund/uicontroller/pkg/service"
	commons "github.com/pwestlake/equity-fund/commons/pkg/service"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
	"encoding/json"
)

// EquityCatalogRoutes ...
// Component for all routes that relate to the entity EquityCatalog
type EquityCatalogRoutes struct {
	commons.Route
	config config.Config
	entitlements entitlements.Entitlements
	equityCatalogService service.EquityCatalogService
}

// NewEquityCatalogRoutes ...
// Create function for NewEquityCatalogRoutes
func NewEquityCatalogRoutes(entitlements entitlements.Entitlements, equityCatalogService service.EquityCatalogService) EquityCatalogRoutes {
	config := config.NewConfig(nil)
	return EquityCatalogRoutes{
		config: config, 
		Route: commons.NewRoute(entitlements),
		equityCatalogService: equityCatalogService}
}

// PostEquityCatalogItem ...
// Handler function for the path: /equitycatalogitem
func (p *EquityCatalogRoutes) PostEquityCatalogItem(w http.ResponseWriter, r *http.Request) {
	postEquityCatalogItem := func(w http.ResponseWriter, r *http.Request) commons.HTTPResponse {
		equityCatalogItem := domain.EquityCatalogItem{}
		err := json.NewDecoder(r.Body).Decode(&equityCatalogItem)
		if err != nil {
			return &commons.HTTPErrorResponse{
				StatusCode: http.StatusBadRequest,
				Body: fmt.Sprintf("Unable to decode json body: %s", err.Error())}
		} 

		err = p.equityCatalogService.PutEquityCatalogItem(&equityCatalogItem)
		if err != nil {
			return &commons.HTTPErrorResponse{
				StatusCode: http.StatusInternalServerError, 
				Body: fmt.Sprintf("Unable to persist EquityCatalogItem %s. Error: %s", equityCatalogItem.ID, err.Error())}
		}
		
		log.Printf("Created EquityCatalogItem %s", equityCatalogItem.ID)
		return &commons.HTTPSuccessResponse{Body: []byte(fmt.Sprintf(`{"id":"%s"}`, equityCatalogItem.ID))}
	}

	p.Route.Handle(w, r, postEquityCatalogItem)
}

// GetEquityCatalogItem ...
// Handler function for the path: /equitycatalogitem/{id}
func (p *EquityCatalogRoutes) GetEquityCatalogItem(w http.ResponseWriter, r *http.Request) {
	getEquityCatalogItem := func(w http.ResponseWriter, r *http.Request) commons.HTTPResponse {
		vars := mux.Vars(r)
		id := vars["id"]
		equityCatalogItem := domain.EquityCatalogItem{}
		err := p.equityCatalogService.GetEquityCatalogItem(id, &equityCatalogItem)
		if err != nil {
			return &commons.HTTPErrorResponse{
				StatusCode: http.StatusNotFound,
				Body: fmt.Sprintf("Failed to get EquityCatalogItem: %s, error: %s", id, err.Error())}
		}

		log.Printf("Found EquityCatalogItem %s", equityCatalogItem.ID)
		equityCatalogItemJSON, err := json.Marshal(equityCatalogItem)
		if err != nil {
			return &commons.HTTPErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Body: fmt.Sprintf("Unable to marshal data for EquityCatalogItem: %s, error: %s", id, err.Error())}
		}
		
		return &commons.HTTPSuccessResponse{Body: equityCatalogItemJSON}		
	}

	p.Route.Handle(w, r, getEquityCatalogItem)
}

// GetAllEquityCatalogItems ...
// Handler function for the path: /equitycatalogitems
func (p *EquityCatalogRoutes) GetAllEquityCatalogItems(w http.ResponseWriter, r *http.Request) {
	getAllEquityCatalogItems := func(w http.ResponseWriter, r *http.Request) commons.HTTPResponse {
		
		equityCatalogItems, err := p.equityCatalogService.GetAllEquityCatalogItems()
		if err != nil {
			return &commons.HTTPErrorResponse{
				StatusCode: http.StatusNotFound,
				Body: fmt.Sprintf("Failed to get EquityCatalogItems, error: %s", err.Error())}
		}

		log.Printf("EquityCatalogItem count: %d", len(*equityCatalogItems))
		equityCatalogItemsJSON, err := json.Marshal(equityCatalogItems)
		if err != nil {
			return &commons.HTTPErrorResponse{
				StatusCode: http.StatusInternalServerError, 
				Body: fmt.Sprintf("Unable to marshal data. Error: %s", err.Error())}
		}
		
		return &commons.HTTPSuccessResponse{Body: equityCatalogItemsJSON}		
	}

	p.Route.Handle(w, r, getAllEquityCatalogItems)
}