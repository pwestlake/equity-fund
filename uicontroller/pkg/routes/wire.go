//+build wireinject

package routes

import (
	"github.com/pwestlake/equity-fund/commons/pkg/dao"
	"github.com/pwestlake/aemo/userservice/pkg/entitlements"
	"github.com/pwestlake/equity-fund/commons/pkg/service"
	"github.com/google/wire"
)

func InitializeContextRoutes() ContextRoutes {
	wire.Build(NewContextRoutes, entitlements.NewEntitlements)
	return ContextRoutes{}
}

func InitializeEquityCatalogRoutes() EquityCatalogRoutes {
	wire.Build(NewEquityCatalogRoutes, entitlements.NewEntitlements, service.NewEquityCatalogService, dao.NewEquityCatalogItemDAO)
	return EquityCatalogRoutes{}
}

func InitializeEndOfDayRoutes() EndOfDayRoutes {
	wire.Build(NewEndOfDayRoutes, entitlements.NewEntitlements, service.NewEndOfDayService, dao.NewEndOfDayItemDAO)
	return EndOfDayRoutes{}
}

func InitializeNewsRoutes() NewsRoutes {
	wire.Build(NewNewsRoutes, entitlements.NewEntitlements, service.NewNewsService, dao.NewNewsItemDAO)
	return NewsRoutes{}
}