//+build wireinject

package service

import (
	"github.com/google/wire"
)

func InitializeMarketStackService() MarketStackService {
	wire.Build(NewMarketStackService)
	return MarketStackService{}
}

func InitializeYahooService() YahooService {
	wire.Build(NewYahooService)
	return YahooService{}
}

func InitializeLSEService() LSEService {
	wire.Build(NewLSEService)
	return LSEService{}
}

func InitializeNLPService() NLPService {
	wire.Build(NewNLPService)
	return NLPService{}
}