//+build wireinject

package routes

import (
	"github.com/pwestlake/aemo/userservice/pkg/entitlements"
	"github.com/google/wire"
)

func InitializeContextRoutes() ContextRoutes {
	wire.Build(NewContextRoutes, entitlements.NewEntitlements)
	return ContextRoutes{}
}