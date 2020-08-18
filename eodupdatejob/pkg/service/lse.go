package service

import (
	"regexp"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/pwestlake/aemo/userservice/pkg/config"
	commonsDomain "github.com/pwestlake/equity-fund/commons/pkg/domain"
	"github.com/pwestlake/equity-fund/eodupdatejob/pkg/domain"
)

// LSEService ...
// Service providing access to the MarketStack api
type LSEService struct {
	config       config.Config
	endpoint     string
	newsEndpoint string
}

// NewLSEService ...
// Create function for a LSEService
func NewLSEService() LSEService {
	config := config.NewConfig(nil)
	endpoint, err := config.GetString("lse.endpoint")
	if err != nil {
		log.Print(err)
	}

	newsEndpoint, err := config.GetString("lse.news-endpoint")
	if err != nil {
		log.Print(err)
	}

	return LSEService{config: config, endpoint: endpoint, newsEndpoint: newsEndpoint}
}

// GetNewsFromDate ...
// Extracts and persists the latest news for the given item from the LSE
func (s *LSEService) GetNewsFromDate(item *commonsDomain.EquityCatalogItem, from time.Time) (*[]commonsDomain.NewsItem, error) {
	params := fmt.Sprintf("tidm=%s&issuername=%s&tab=analysis&tabId=a7bd00f8-7846-496a-8692-c55a0a24380c", item.LSEtidm, item.LSEissuername)
	message := domain.RefreshMessage{
		Path:       "issuer-profile",
		Parameters: params,
		Components: []domain.LSEComponent{{
			ComponentID: "block_content:16061956-5f74-42e9-ad94-fb7c4457bef4",
			Parameters:  "null",
		}},
	}

	jsonResponse, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(s.endpoint, "application/json", bytes.NewBuffer(jsonResponse))
	if err != nil {
		return nil, err
	}

	buffer := strings.Builder{}
	_, err = io.Copy(&buffer, resp.Body)
	if err != nil {
		log.Printf("Failed to read response. %s", err.Error())
		return nil, err
	}

	slice := buffer.String()
	startIndex := strings.Index(slice, `{"name":"newsbyissuercode"`)
	if startIndex < 0 {
		return nil, fmt.Errorf("Unable to find start of news in json result")
	}

	endIndex := strings.Index(slice[startIndex:], "]")
	if endIndex < 0 {
		return nil, fmt.Errorf("Unable to find end of news in json response")
	}

	subslice := []byte(slice[startIndex : startIndex+endIndex+1])
	subslice = append(subslice, "}}"...)

	news := domain.LSENews{}
	err = json.Unmarshal(subslice, &news)
	if err != nil {
		return nil, err
	}

	selectedNews := []commonsDomain.NewsItem{}
	for _, v := range news.Value.Content {
		date, err := time.Parse("2006-01-02T15:04:05.000", v.DateTime)
		if err != nil {
			log.Printf("Invalid date format %s, Ignoring %s. %s", v.DateTime, v.Title, err.Error())
			continue
		}

		if date.After(from) {
			article := s.getNewsArticle(v)
			newsItem := commonsDomain.NewsItem{
				CatalogRef:  item.ID,
				CompanyCode: v.CompanyCode,
				CompanyName: v.CompanyName,
				Content:     article,
				DateTime:    date,
				Title:       v.Title,
				Sentiment:   0.0,
			}
			selectedNews = append(selectedNews, newsItem)
		}
	}
	return &selectedNews, nil
}

func (s *LSEService) getNewsArticle(newsItem domain.LSENewsItem) string {
	url := fmt.Sprintf("%s%s/%s/%d", s.newsEndpoint, newsItem.CompanyCode, formatNewsTitle(newsItem.Title), newsItem.ID)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Unable to get news article: %s", url)
		return ""
	}

	buffer := strings.Builder{}
	_, err = io.Copy(&buffer, resp.Body)
	if err != nil {
		log.Printf(`Failed to read response body for %s. Returning "". %s`, url, err.Error())
		return ""
	}

	// Extract the essential news
	slice := buffer.String()
	startIndex := strings.Index(slice, `&l;body`)
	if startIndex < 0 {
		log.Printf(`Unable to find start of news for %s. Returning ""`, url)
		return ""
	}

	endIndex := strings.Index(slice[startIndex:], "&l;/body&g")
	if endIndex < 0 {
		log.Printf(`Unable to find end of news for %s. Returning ""`, url)
		return ""
	}

	subslice := []byte(slice[startIndex : startIndex+endIndex+1])
	text := extractText(string(subslice))
	return text
}

// Convert a news title into a url friendly form
// Replace non-alphanumeric characters with a hyphen
// Holding(s) in company
// becomes:
// Holding-s-in-company
func formatNewsTitle(title string) string {
	regx := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	return regx.ReplaceAllString(title, "-")
}

func extractText(html string) string {
	regx := regexp.MustCompile(`&l;.*?&g;|&s;`)
	text := regx.ReplaceAllString(html, "")
	regx = regexp.MustCompile(`&a;#160;`)
	text = regx.ReplaceAllString(text, " ")
	regx = regexp.MustCompile(`&a;#163;`)
	text = regx.ReplaceAllString(text, "£")
	return text
}