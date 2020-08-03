//+build wireinject

package service

import (
	"github.com/google/wire"
)

func InitializeMarketStackService() MarketStackService {
	wire.Build(NewMarketStackService)
	return MarketStackService{}
}