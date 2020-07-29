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