package routes

import (
	"time"
	"github.com/pwestlake/equity-fund/commons/pkg/domain"
	"encoding/json"
	"fmt"
	"strconv"
	"net/http"
	"github.com/pwestlake/aemo/userservice/pkg/entitlements"
	"github.com/pwestlake/aemo/userservice/pkg/config"
	commons "github.com/pwestlake/equity-fund/commons/pkg/service"
)

// NewsRoutes ...
// Component definition for News routes
type NewsRoutes struct {
	commons.Route
	config config.Config
	entitlements entitlements.Entitlements
	newsService commons.NewsService
}

// NewNewsRoutes ...
// Create function for NewsRoutes
func NewNewsRoutes(entitlements entitlements.Entitlements, newsService commons.NewsService) NewsRoutes {
	config := config.NewConfig(nil)
	return NewsRoutes{
		config: config, 
		Route: commons.NewRoute(entitlements),
		newsService: newsService}
}

// GetNewsItems ...
// Handler method for the route /newsitems
// Query params: 
// count
// key
// sortkey
// catalogref
func (s *NewsRoutes) GetNewsItems(w http.ResponseWriter, r *http.Request) {
	getNewsItems := func(w http.ResponseWriter, r *http.Request) commons.HTTPResponse {
		count, err := strconv.ParseInt(r.URL.Query().Get("count"), 10, 32)
		if err != nil {
			count = 1000
		}

		id := r.URL.Query().Get("catalogref")
		idptr := &id
		if len(*idptr) == 0 {
			idptr = nil
		}
		
		key := r.URL.Query().Get("key")
		sortkey := r.URL.Query().Get("sortkey")
		sortkeyDate, err := time.Parse("2006-01-02T15:04:05Z", sortkey)

		var startKey *domain.NewsItem
		if err != nil || sortkey == "" {
			startKey = nil
		} else {
			startKey = &domain.NewsItem {
				ID: key,
				DateTime: sortkeyDate,
			}
		}

		items, err := s.newsService.GetNewsItems(int(count), startKey, idptr)
		if err != nil {
			return &commons.HTTPErrorResponse{
				StatusCode: http.StatusNotFound,
				Body: fmt.Sprintf("Failed to get news items. %s", err.Error())}
		}

		itemsJSON, err := json.Marshal(*items)
		if err != nil {
			return &commons.HTTPErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Body: fmt.Sprintf("Unable to marshal data for news items, error: %s", err.Error())}
		}

		return &commons.HTTPSuccessResponse{Body: itemsJSON}
	}

	s.Route.Handle(w, r, getNewsItems)
}

