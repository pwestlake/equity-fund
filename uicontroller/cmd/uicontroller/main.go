package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"
	"github.com/pwestlake/aemo/userservice/pkg/config"
	"github.com/pwestlake/equity-fund/uicontroller/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	configPtr := flag.String("server", "http://localhost:8888", "the url of the cloud config server")
    profilePtr := flag.String("profile", "dev", "the configuration profile")
    labelPtr := flag.String("label", "development", "the configuration label")
    flag.Parse()

    configURL := fmt.Sprintf("%s/equity-fund-uicontroller", *configPtr)
	cfg := config.NewConfig(&config.Params{Server: configURL, Profile: *profilePtr, Label: *labelPtr})
	
	port, err := cfg.GetInt("server.port")
    if err != nil {
        log.Printf("%s", err.Error())
        return
    }

    securityHelper := config.NewSecurityHelper()
    jwtMiddleware, err := securityHelper.BuildJWTMiddleware()
    if err != nil {
        log.Printf("%s", err.Error())
	}
	
	contextRoutes := routes.InitializeContextRoutes()
	equityCatalogRoutes := routes.InitializeEquityCatalogRoutes()
	endOfDayRoutes := routes.InitializeEndOfDayRoutes()
	newsRoutes := routes.InitializeNewsRoutes()

	r := mux.NewRouter()
	// GET /title
    r.Handle("/equity-fund/uicontroller/title", 
		jwtMiddleware.Handler(http.HandlerFunc(contextRoutes.GetTitle))).Methods(http.MethodGet)
	
	// POST /equitycatalogitem
	r.Handle("/equity-fund/uicontroller/equitycatalogitem", 
		jwtMiddleware.Handler(http.HandlerFunc(equityCatalogRoutes.PostEquityCatalogItem))).Methods(http.MethodPost)


	// GET /equitycatalogitem/{id}
	r.Handle("/equity-fund/uicontroller/equitycatalogitem/{id}", 
		jwtMiddleware.Handler(http.HandlerFunc(equityCatalogRoutes.GetEquityCatalogItem))).Methods(http.MethodGet)

	// GET /equitycatalogitem/
	r.Handle("/equity-fund/uicontroller/equitycatalogitem", 
		jwtMiddleware.Handler(http.HandlerFunc(equityCatalogRoutes.GetAllEquityCatalogItems))).Methods(http.MethodGet)

	// DELETE /equitycatalogitem/{id}

	// GET /timeseries/close/{id}
	r.Handle("/equity-fund/uicontroller/timeseries/close/{id}", 
		jwtMiddleware.Handler(http.HandlerFunc(endOfDayRoutes.GetClosePriceTimeSeries))).Methods(http.MethodGet)

	// GET /latest-eod/
	r.Handle("/equity-fund/uicontroller/latest-eod", 
	jwtMiddleware.Handler(http.HandlerFunc(endOfDayRoutes.GetLatestEndOfDayItems))).Methods(http.MethodGet)

	// GET /latest-eod-item/
	r.Handle("/equity-fund/uicontroller/latest-eod-item", 
	jwtMiddleware.Handler(http.HandlerFunc(endOfDayRoutes.GetLatestEndOfDayItem))).Methods(http.MethodGet)

	// GET /newsitems/
	r.Handle("/equity-fund/uicontroller/newsitems", 
	jwtMiddleware.Handler(http.HandlerFunc(newsRoutes.GetNewsItems))).Methods(http.MethodGet)

	log.Println("Listening on port: ", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}