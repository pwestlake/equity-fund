//+build wireinject

package service

import (
	"github.com/pwestlake/equity-fund/commons/pkg/dao"
	"github.com/google/wire"
)

func InitializeEquityCatalogService() EquityCatalogService {
	wire.Build(NewEquityCatalogService, dao.NewEquityCatalogItemDAO)
	return EquityCatalogService{}
}

func InitializeEndOfDayService() EndOfDayService {
	wire.Build(NewEndOfDayService, dao.NewEndOfDayItemDAO)
	return EndOfDayService{}
}

func InitializeNewsService() NewsService {
	wire.Build(NewNewsService, dao.NewNewsItemDAO)
	return NewsService{}
}