// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package routes

import (
	"github.com/pwestlake/aemo/userservice/pkg/entitlements"
)
// Injectors from wire.go:

func InitializeContextRoutes() ContextRoutes {
	entitlementsEntitlements := entitlements.NewEntitlements()
	contextRoutes := NewContextRoutes(entitlementsEntitlements)
	return contextRoutes
}