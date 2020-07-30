// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package service

import (
	"github.com/pwestlake/equity-fund/commons/pkg/dao"
)

// Injectors from wire.go:

func InitializeEquityCatalogService() EquityCatalogService {
	equityCatalogItemDAO := dao.NewEquityCatalogItemDAO()
	equityCatalogService := NewEquityCatalogService(equityCatalogItemDAO)
	return equityCatalogService
}

func InitializeEndOfDayService() EndOfDayService {
	endOfDayItemDAO := dao.NewEndOfDayItemDAO()
	endOfDayService := NewEndOfDayService(endOfDayItemDAO)
	return endOfDayService
}
