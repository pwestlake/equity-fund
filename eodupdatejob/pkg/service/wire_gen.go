// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package service

// Injectors from wire.go:

func InitializeMarketStackService() MarketStackService {
	marketStackService := NewMarketStackService()
	return marketStackService
}

func InitializeYahooService() YahooService {
	yahooService := NewYahooService()
	return yahooService
}
