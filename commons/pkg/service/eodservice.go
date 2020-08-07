package service

import (
	"time"
	"github.com/pwestlake/equity-fund/commons/pkg/domain"
	"github.com/pwestlake/equity-fund/commons/pkg/dao"

)


// EndOfDayService ...
type EndOfDayService struct {
	endOfDayItemDaoDao dao.EndOfDayItemDAO
}

// NewEndOfDayService ...
// Create function for an EndOfDayService
func NewEndOfDayService(endOfDayItemDaoDao dao.EndOfDayItemDAO) EndOfDayService {
	return EndOfDayService{endOfDayItemDaoDao: endOfDayItemDaoDao}
}

// PutEndOfDayItems ...
// Service method to persist an array of EndOfDayItems in the database
func (s *EndOfDayService) PutEndOfDayItems(items *[]domain.EndOfDayItem) error {
	return s.endOfDayItemDaoDao.PutEndOfDayItems(items)
}

// GetEndOfDayItems ...
// Service method to retrieve an array of EndOfDayItems according to the given id and from date
func (s *EndOfDayService) GetEndOfDayItems(id string, from time.Time) (*[]domain.EndOfDayItem, error) {
	return s.endOfDayItemDaoDao.GetEndOfDayItems(id, from)
}

// GetAllEndOfDayItemsByDate ...
// Service method to retrieve an array of EndOfDayItems according to the given date
func (s *EndOfDayService) GetAllEndOfDayItemsByDate(date time.Time) (*[]domain.EndOfDayItem, error) {
	return s.endOfDayItemDaoDao.GetAllEndOfDayItemsByDate(date)
}

// GetLatestItem ...
// Service method to retrieve the latest eod item for a given id
func (s *EndOfDayService) GetLatestItem(id string) (*domain.EndOfDayItem, error) {
	return s.endOfDayItemDaoDao.GetLatestItem(id)
}