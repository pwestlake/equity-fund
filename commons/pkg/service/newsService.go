package service

import (
	"github.com/pwestlake/equity-fund/commons/pkg/domain"
	"github.com/pwestlake/equity-fund/commons/pkg/dao"

)

// NewsService ...
type NewsService struct {
	newsItemDao dao.NewsItemDAO
}

// NewNewsService ...
// Create function for an NewsService
func NewNewsService(newsItemDao dao.NewsItemDAO) NewsService {
	return NewsService{newsItemDao: newsItemDao}
}

// PutNewsItems ...
// Service method to persist an array of NewsItems in the database
func (s *NewsService) PutNewsItems(items *[]domain.NewsItem) error {
	return s.newsItemDao.PutNewsItems(items)
}

// GetLatestItem ...
// Service method to retrieve the latest eod item for a given id
func (s *NewsService) GetLatestItem(id string) (*domain.NewsItem, error) {
	return s.newsItemDao.GetLatestItem(id)
}

// GetNewsItems ...
func (s *NewsService) GetNewsItems(count int, offset *domain.NewsItem, id *string) (*[]domain.NewsItem, error) {
	return s.newsItemDao.GetNewsItems(count, offset, id)
}
