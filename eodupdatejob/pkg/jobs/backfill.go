package jobs

import (
	"time"
	"log"
	"github.com/pwestlake/equity-fund/eodupdatejob/pkg/service"
	commons "github.com/pwestlake/equity-fund/commons/pkg/service"
	"github.com/pwestlake/equity-fund/commons/pkg/domain"
)

// BackFillJob ...
// Component job the will perform a data back-fill
type BackFillJob struct {
	dataService service.MarketStackService
	yahooService service.YahooService
	lseService service.LSEService
	equityCatalogService commons.EquityCatalogService
	endOfDayService commons.EndOfDayService
	newsService commons.NewsService
	nlpService service.NLPService
}

// NewBackFillJob ...
// Create function for a BackFillJob component
func NewBackFillJob(dataService service.MarketStackService, 
	yahooService service.YahooService,
	lseService service.LSEService,
	equityCatalogService commons.EquityCatalogService,
	endOfDayService commons.EndOfDayService,
	newsService commons.NewsService,
	nlpService service.NLPService) BackFillJob {
	return BackFillJob{
		dataService: dataService,
		yahooService: yahooService,
		lseService: lseService,
		equityCatalogService: equityCatalogService,
		endOfDayService: endOfDayService,
		newsService: newsService,
		nlpService: nlpService}
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

	source, err := s.dataService.GetData(symbol)
	if err != nil {
		log.Printf("Failed to source data: %s", err.Error())
		return
	}

	target := buildTarget(source, catalogItems, nil)

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
	s.updateWithLatestFromMarketstack()
	s.updateWithLatestFromYahoo()
	s.FetchLatestNews()
	log.Printf("Complete.")
}

// FetchLatestNews ...
// Fetch the latest news for each of the catalog items
func (s *BackFillJob) FetchLatestNews() {
	catalogItems, err := s.equityCatalogService.GetAllEquityCatalogItems()
	if err != nil {
		log.Printf("Failed to get catalog items: %s", err.Error())
		return
	}

	for _, v := range *catalogItems {
		var date time.Time
		item, err := s.newsService.GetLatestItem(v.ID)
		if err != nil {
			log.Printf("Failed to fetch latest news item for %s. %s", v.Symbol, err.Error())
			
			// Set a date to fetch all available news items
			date = time.Now()
			date = date.AddDate(-1, 0, 0)
		} else {
			date = item.DateTime
		}
		
		newsItems, err := s.lseService.GetNewsFromDate(&v, date)
		if err != nil {
			log.Printf("Failed to fetch news for %s. %s", v.Symbol, err.Error())
		}

		plural := ""
		if len(*newsItems) > 1 ||  len(*newsItems) == 0{
			plural = "s"
		}

		log.Printf("Found %d news item%s for %s", len(*newsItems), plural, v.Symbol)

		for i, v := range *newsItems {
			sentiment, err := s.nlpService.GetSentiment(&v.Content)
			if err != nil {
				log.Printf("Failed to get sentiment for %s. %s", v.Title, err.Error())
			}
			(*newsItems)[i].Sentiment = sentiment
		}

		err = s.newsService.PutNewsItems(newsItems)
		if (err != nil) {
			log.Printf("Failed to persist news items for %s. %s", v.Symbol, err.Error())
		}
	}
}

func (s *BackFillJob) updateWithLatestFromMarketstack() {
	catalogItems, err := s.equityCatalogService.GetEquityCatalogItemsByDatasource("marketstack")
	if err != nil {
		log.Printf("Failed to get catalog items: %s", err.Error())
		return
	}

	if len(*catalogItems) == 0 {
		log.Printf("Nothing to update from marketstack")
		return
	}

	// Find the date of the last update and derive the 'from' date
	// Assume that all items were updated at the same time
	eodItem, err := s.endOfDayService.GetLatestItem((*catalogItems)[0].ID)
	if err != nil {
		log.Printf("Failed to get latest item for %s, %s. Aborting", (*catalogItems)[0].ID, err.Error())
		return
	}


	fromDate := eodItem.Date.AddDate(0, 0, 1)
	if fromDate.After(today()) || fromDate.Equal(today()) {
		log.Printf("All marketstack data is up to date.")
		return
	}

	// Build an array of symbols
	symbols := make([]string, len(*catalogItems))
	for i, v := range *catalogItems {
		symbols[i] = v.Symbol
	}

	current, err := s.endOfDayService.GetAllEndOfDayItemsByDate(eodItem.Date)
	if err != nil {
		log.Printf("Failed to retrieve current data: %s", err.Error())
		return
	}
	
	log.Printf("Updating %d items from %s", len(*catalogItems), fromDate.String())
	source, err := s.dataService.GetDataFromDate(symbols, fromDate)

	if err != nil {
		log.Printf("Failed to source data: %s", err.Error())
		return
	}

	target := buildTarget(source, catalogItems, current)

	err = s.endOfDayService.PutEndOfDayItems(target)
	if err != nil {
		log.Printf("Failed to persist end of day data")
	}

	plural := ""
	if len(*target) > 1 ||  len(*target) == 0{
		plural = "s"
	}
	log.Printf("Found and persisted %d new eod item%s", len(*target), plural)
}

func (s *BackFillJob) updateWithLatestFromYahoo() {
	catalogItems, err := s.equityCatalogService.GetEquityCatalogItemsByDatasource("yahoo")
	if err != nil {
		log.Printf("Failed to get catalog items: %s", err.Error())
		return
	}

	// Find the date of the last update and derive the 'from' date
	// Assume that all items were updated at the same time
	eodItem, err := s.endOfDayService.GetLatestItem((*catalogItems)[0].ID)
	if err != nil {
		log.Printf("Failed to get latest item for %s, %s. Aborting", (*catalogItems)[0].ID, err.Error())
		return
	}

	fromDate := eodItem.Date.AddDate(0, 0, 1)
	if fromDate.After(today()) || fromDate.Equal(today()) {
		log.Printf("All yahoo data is up to date.")
		return
	}

	current, err := s.endOfDayService.GetAllEndOfDayItemsByDate(eodItem.Date)
	if err != nil {
		log.Printf("Failed to retrieve current data: %s", err.Error())
		return
	}

	for _, v := range *catalogItems {
		source, err := s.yahooService.GetDataFromDate(v.Symbol, fromDate)
		if err != nil {
			log.Printf("Failed to source %s. Aborting: %s", v.Symbol, err.Error())
			return
		}
		
		target := buildTarget(source, catalogItems, current)

		err = s.endOfDayService.PutEndOfDayItems(target)
		if err != nil {
			log.Printf("Failed to persist end of day data")
		}

		plural := ""
		if len(*target) > 1 ||  len(*target) == 0{
			plural = "s"
		}
		log.Printf("Found and persisted %d new eod item%s for %s", len(*target), plural, v.Symbol)
	}
	
}

func buildTarget(source *[]domain.EndOfDaySourceItem, 
	catalog *[]domain.EquityCatalogItem,
	current *[]domain.EndOfDayItem) *[]domain.EndOfDayItem {
	result := make([]domain.EndOfDayItem, len(*source))

	idMap := make(map[string]string, len(*catalog))
	for _, v := range *catalog {
		idMap[v.Symbol] = v.ID
	}

	for i, v := range *source {
		previous := previousEndOfDayItem(v.Date.Time, v.Symbol, current)
		item := domain.EndOfDayItem{
			ID: idMap[v.Symbol],
			Symbol: v.Symbol,
			Open: v.Open,
			High: v.High,
			Low: v.Low,
			Close: v.Close,
			CloseChg: v.Close - previous.Close,
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

func today() time.Time {
	utc, _ := time.LoadLocation("UTC") 
	today := time.Now()
	today = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, utc)
	return today
}

func previousEndOfDayItem(date time.Time, symbol string, current *[]domain.EndOfDayItem) *domain.EndOfDayItem {
	eodItem := domain.EndOfDayItem{
		Close: 0.0,
	}

	previousDate := (*current)[0].Date

	if current != nil {
		for _, v := range *current {
			if v.Date.Equal(previousDate) && v.Symbol == symbol {
				eodItem = v
				break
			}
		}
	}

	return &eodItem
}