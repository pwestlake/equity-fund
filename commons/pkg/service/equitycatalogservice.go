package service

import (
	"github.com/pwestlake/equity-fund/commons/pkg/domain"
	"github.com/pwestlake/equity-fund/commons/pkg/dao"

)

// EquityCatalogService ...
type EquityCatalogService struct {
	equityCatalogItemDao dao.EquityCatalogItemDAO
}

// NewEquityCatalogService ...
// Create function for a NewUserService
func NewEquityCatalogService(equityCatalogItemDao dao.EquityCatalogItemDAO) EquityCatalogService {
	return EquityCatalogService{equityCatalogItemDao: equityCatalogItemDao}
}

// GetEquityCatalogItem ...
// Service method to retrieve a user with the given id from the database
func (s *EquityCatalogService) GetEquityCatalogItem(id string, equityCatalogItem *domain.EquityCatalogItem) error {
	return s.equityCatalogItemDao.GetEquityCatalogItem(id, equityCatalogItem)
}

// GetAllEquityCatalogItems ...
// Service method to return an array of user ID's
func (s *EquityCatalogService) GetAllEquityCatalogItems() (*[]domain.EquityCatalogItem, error) {
	return s.equityCatalogItemDao.GetEquityCatalogItems()
}

// PutEquityCatalogItem ...
// Service method to persist a new user in the database
func (s *EquityCatalogService) PutEquityCatalogItem(equityCatalogItem *domain.EquityCatalogItem) error {
	return s.equityCatalogItemDao.PutEquityCatalogItem(equityCatalogItem)
}