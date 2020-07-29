//+build wireinject

package routes

import (
	"github.com/pwestlake/equity-fund/uicontroller/pkg/dao"
	"github.com/pwestlake/aemo/userservice/pkg/entitlements"
	"github.com/pwestlake/equity-fund/uicontroller/pkg/service"
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