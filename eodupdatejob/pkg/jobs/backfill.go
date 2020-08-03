package jobs

import (
	"log"
	"github.com/pwestlake/equity-fund/eodupdatejob/pkg/service"
	commons "github.com/pwestlake/equity-fund/commons/pkg/service"
	"github.com/pwestlake/equity-fund/commons/pkg/domain"
)

// BackFillJob ...
// Component job the will perform a data back-fill
type BackFillJob struct {
	dataService service.MarketStackService
	equityCatalogService commons.EquityCatalogService
	endOfDayService commons.EndOfDayService
}

// NewBackFillJob ...
// Create function for a BackFillJob component
func NewBackFillJob(dataService service.MarketStackService, 
	equityCatalogService commons.EquityCatalogService,
	endOfDayService commons.EndOfDayService) BackFillJob {
	return BackFillJob{
		dataService: dataService, 
		equityCatalogService: equityCatalogService,
		endOfDayService: endOfDayService}
}

// Run ...
// The job run function
func (s *BackFillJob) Run(symbol string) {
	catalogItems, err := s.equityCatalogService.GetEquityCatalogItemsBySymbol(symbol)
	if err != nil {
		log.Printf("Failed to retrieve catalog item: %s", err.Error())
		return
	}

	if len(*catalogItems) > 1 || len(*catalogItems) < 1 {
		log.Printf("Invalid number of catalog items. Expected 1 found %d", len(*catalogItems))
		return
	}

	item := (*catalogItems)[0]

	source, err := s.dataService.GetData(symbol)
	if err != nil {
		log.Printf("Failed to source data: %s", err.Error())
		return
	}

	target := buildTarget(source, item.ID)

	err = s.endOfDayService.PutEndOfDayItems(target)
	if err != nil {
		log.Printf("Failed to persist end of day data for %s", symbol)
	}

	plural := ""
	if len(*target) > 1 ||  len(*target) == 0{
		plural = "s"
	}
	log.Printf("Found and persisted %d new eod item%s for %s", len(*target), plural, symbol)
}

// UpdateWithLatest ...
// Update all catalog items with latest data
func (s *BackFillJob) UpdateWithLatest() {
	catalogItems, err := s.equityCatalogService.GetAllEquityCatalogItems()
	if err != nil {
		log.Printf("Failed to get catalog items: %s", err.Error())
		return
	}

	for _, v := range *catalogItems {
		eodItem, err := s.endOfDayService.GetLatestItem(v.ID)
		if err != nil {
			log.Printf("Failed to get latest item for %s, %s", v.ID, err.Error())
			break;
		}

		fromDate := eodItem.Date.AddDate(0, 0, 1)
		log.Printf("Updating: %s, %s from %s", v.ID, v.Symbol, fromDate.String())
		source, err := s.dataService.GetDataFromDate(v.Symbol, fromDate)

		if err != nil {
			log.Printf("Failed to source data: %s", err.Error())
			return
		}
	
		target := buildTarget(source, v.ID)
	
		err = s.endOfDayService.PutEndOfDayItems(target)
		if err != nil {
			log.Printf("Failed to persist end of day data for %s", v.Symbol)
		}

		plural := ""
		if len(*target) > 1 ||  len(*target) == 0{
			plural = "s"
		}
		log.Printf("Found and persisted %d new eod item%s for %s", len(*target), plural, v.Symbol)
	}
}

func buildTarget(source *[]domain.EndOfDaySourceItem, id string) *[]domain.EndOfDayItem {
	result := make([]domain.EndOfDayItem, len(*source))

	for i, v := range *source {
		item := domain.EndOfDayItem{
			ID: id,
			Symbol: v.Symbol,
			Open: v.Open,
			High: v.High,
			Low: v.Low,
			Close: v.Close,
			Volume: v.Volume,
			AdjHigh: v.AdjHigh,
			AdjLow: v.AdjLow,
			AdjClose: v.AdjClose,
			AdjOpen: v.AdjOpen,
			AdjVolume: v.AdjVolume,
			Exchange: v.Exchange,
			Date: v.Date.Time,
		}

		result[i] = item
	}

	return &result
}
