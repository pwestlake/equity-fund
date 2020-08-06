package domain

import (
	"strings"
	"time"
)

// EndOfDayItem ...
// Domain struct for an EndOfDayItem
type EndOfDayItem struct {
	ID        string    `json:"id"`
	Open      float32   `json:"open"`
	High      float32   `json:"high"`
	Low       float32   `json:"low"`
	Close     float32   `json:"close"`
	CloseChg  float32	`json:"close_chg"`
	Volume    float32    `json:"volume"`
	AdjHigh   float32   `json:"adj_high"`
	AdjLow    float32   `json:"adj_low"`
	AdjClose  float32   `json:"adj_close"`
	AdjOpen   float32   `json:"adj_open"`
	AdjVolume int64     `json:"adj_volume"`
	Symbol    string    `json:"symbol"`
	Exchange  string    `json:"exchange"`
	Date      time.Time `json:"date"`
}

// EndOfDaySourceItem ...
// Domain struct for an EndOfDaySourceItem
type EndOfDaySourceItem struct {
	Open      float32    `json:"open"`
	High      float32    `json:"high"`
	Low       float32    `json:"low"`
	Close     float32    `json:"close"`
	Volume    float32    `json:"volume"`
	AdjHigh   float32    `json:"adj_high"`
	AdjLow    float32    `json:"adj_low"`
	AdjClose  float32    `json:"adj_close"`
	AdjOpen   float32    `json:"adj_open"`
	AdjVolume int64      `json:"adj_volume"`
	Symbol    string     `json:"symbol"`
	Exchange  string     `json:"exchange"`
	Date      SourceTime `json:"date"`
}

// PageDescriptor ...
// Domain struct for a PageDescriptor
type PageDescriptor struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
	Count  int32 `json:"count"`
	Total  int32 `json:"total"`
}

// EndOfDayDataExtract ...
type EndOfDayDataExtract struct {
	Pagination PageDescriptor       `json:"pagination"`
	Data       []EndOfDaySourceItem `json:"data"`
}

// SourceTime ...
// Time representing the time format in the MarketStack response
type SourceTime struct {
	time.Time
}

// Mon Jan 2 15:04:05 -0700 MST 2006
const ctLayout = "2006-01-02T15:04:05-0700"

// UnmarshalJSON ...
// Unmarshal the date format in the source data to that for time.Time
func (st *SourceTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	
    if s == "null" {
		st.Time = time.Time{}
       return nil
	}
	
	var err error
    st.Time, err = time.Parse(ctLayout, s)
    return err
}