package routes

import (
	"encoding/json"
	"github.com/pwestlake/equity-fund/uicontroller/pkg/domain"
	"fmt"
	"time"
	"github.com/gorilla/mux"
	"net/http"
	commons "github.com/pwestlake/equity-fund/commons/pkg/service"
	"github.com/pwestlake/aemo/userservice/pkg/entitlements"
	"github.com/pwestlake/aemo/userservice/pkg/config"

)

// EndOfDayRoutes ...
// Component definition for EndOfDay routes
type EndOfDayRoutes struct {
	commons.Route
	config config.Config
	entitlements entitlements.Entitlements
	endOfDayService commons.EndOfDayService
}

// NewEndOfDayRoutes ...
// Create function for EndOfDayRoutes
func NewEndOfDayRoutes(entitlements entitlements.Entitlements, endOfDayService commons.EndOfDayService) EndOfDayRoutes {
	config := config.NewConfig(nil)
	return EndOfDayRoutes{
		config: config, 
		Route: commons.NewRoute(entitlements),
		endOfDayService: endOfDayService}
}

// GetClosePriceTimeSeries ...
// Handler method for the route /timeseries/close/{id}
func (s *EndOfDayRoutes) GetClosePriceTimeSeries(w http.ResponseWriter, r *http.Request) {
	getClosePriceTimeSeries := func(w http.ResponseWriter, r *http.Request) commons.HTTPResponse {
		vars := mux.Vars(r)
		id := vars["id"]

		from := time.Time{}
		items, err := s.endOfDayService.GetEndOfDayItems(id, from)
		if err != nil {
			return &commons.HTTPErrorResponse{
				StatusCode: http.StatusNotFound,
				Body: fmt.Sprintf("Failed to get Timeseries for: %s, error: %s", id, err.Error())}
		}

		timeseries := make([]domain.DateValue, len(*items))
		for i, v := range *items {
			timeseries[i] = domain.DateValue {
				Date: v.Date,
				Value: v.Close,
			}
		}

		timeseriesJSON, err := json.Marshal(timeseries)
		if err != nil {
			return &commons.HTTPErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Body: fmt.Sprintf("Unable to marshal data for timeseries, error: %s", err.Error())}
		}

		return &commons.HTTPSuccessResponse{Body: timeseriesJSON}
	}

	s.Route.Handle(w, r, getClosePriceTimeSeries)
}