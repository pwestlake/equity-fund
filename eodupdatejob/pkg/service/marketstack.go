package service

import (
	"time"
	"io"
	"strings"
	"github.com/pwestlake/equity-fund/commons/pkg/domain"
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/pwestlake/aemo/userservice/pkg/config"
)

// MarketStackService ...
// Service providing access to the MarketStack api
type MarketStackService struct {
	config   config.Config
	endpoint string
	apikey	 string
}

// NewMarketStackService ...
// Create function for a MarketStackService
func NewMarketStackService() MarketStackService {
	config := config.NewConfig(nil)
	endpoint, err := config.GetString("marketstack.endpoint")
	if err != nil {
		log.Print(err)
	}

	apikey, err := config.GetString("marketstack.access-key")
	if err != nil {
		log.Print(err)
	}

	return MarketStackService{config: config, endpoint: endpoint, apikey: apikey}
}

// GetData ...
// Get all end of day data for the given symbol
func (s *MarketStackService) GetData(symbol string) (*[]domain.EndOfDaySourceItem, error) {
	result := []domain.EndOfDaySourceItem{}
	var err error
	blockSize := 1000
	maxSteps:= 5
	complete := false
	for offset := 0; !complete; offset += blockSize {
		resp, err := http.Get(fmt.Sprintf("%s/?access_key=%s&symbols=%s&offset=%d&limit=%d", 
			s.endpoint, s.apikey, symbol, offset, blockSize))
		if err != nil {
			return nil, err;
		}

		interim := domain.EndOfDayDataExtract{}
		buffer := strings.Builder{}
		_, err = io.Copy(&buffer, resp.Body)
		if err != nil {
			log.Printf("Failed to read url. %s", err.Error())
			return nil, err
		}
		
		err = json.Unmarshal([]byte(buffer.String()), &interim)
		if err != nil {
			log.Printf("Failed to unmarshal json: %s", err.Error())
			return nil, err
		}

		result = append(result, interim.Data...)

		steps := (offset / blockSize) + 1
		complete = steps >= maxSteps || int32(len(result)) >= interim.Pagination.Total
		if complete && steps >= maxSteps {
			err = fmt.Errorf("Aborted after %d cycles", maxSteps)
			return nil, err
		}
	}

	return &result, err
}

// GetDataFromDate ...
// Get data for the given symbols and from the given date
func (s *MarketStackService) GetDataFromDate(symbols []string, date time.Time) (*[]domain.EndOfDaySourceItem, error) {
	result := []domain.EndOfDaySourceItem{}
	var err error

	// Build a comma separted string of symbols
	var symbolStr strings.Builder
	for _, v := range symbols {
		if symbolStr.Len() != 0 {
			symbolStr.WriteString(",")
		}
		symbolStr.WriteString(v)
	}

	dateStr := fmt.Sprintf("%4d-%02d-%02d", date.Year(), date.Month(), date.Day())
	url := fmt.Sprintf("%s/?access_key=%s&symbols=%s&offset=%d&limit=%d&date_from=%s", 
		s.endpoint, s.apikey, symbolStr.String(), 0, 1000, dateStr)
	
	resp, err := http.Get(url)
	if err != nil {
		return nil, err;
	}

	interim := domain.EndOfDayDataExtract{}
	if interim.Pagination.Count != interim.Pagination.Total {
		return nil, fmt.Errorf("More data than expected. Use backfill")
	}

	buffer := strings.Builder{}
	_, err = io.Copy(&buffer, resp.Body)
	if err != nil {
		log.Printf("Failed to read url. %s", err.Error())
		return nil, err
	}
	
	err = json.Unmarshal([]byte(buffer.String()), &interim)
	if err != nil {
		log.Printf("Failed to unmarshal json: %s", err.Error())
		return nil, err
	}

	result = append(result, interim.Data...)
	return &result, err
}