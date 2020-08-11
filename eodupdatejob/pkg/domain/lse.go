package domain

// RefreshMessage ...
// Post message to request data from the LSE
type RefreshMessage struct {
	Path       string         `json:"path"`
	Parameters string         `json:"parameters"`
	Components []LSEComponent `json:"components"`
}

// LSEComponent ...
// Part of a RefreshMessage
type LSEComponent struct {
	ComponentID string `json:"componentId"`
	Parameters  string `json:"parameters"`
}

// LSENews ...
// LSE News data structure
type LSENews struct {
	Name  string            `json:"name"`
	Type  string            `json:"type"`
	Value LSENewsItemHolder `json:"value"`
}

// LSENewsItemHolder ...
type LSENewsItemHolder struct {
	Content []LSENewsItem `json:"content"`
}

// LSENewsItem ...
type LSENewsItem struct {
	ID               int     `json:"id"`
	Category         string  `json:"category"`
	Title            string  `json:"title"`
	Body             string  `json:"body"`
	Source           string  `json:"source"`
	URL              string  `json:"url"`
	CompanyCode      string  `json:"companycode"`
	CompanyName      string  `json:"companyname"`
	DateTime         string  `json:"datetime"`
	NewsSource       string  `json:"newssource"`
	RNSNumber        string  `json:"rnsnumber"`
	IssuerCode       string  `json:"issuercode"`
	IssuerName       string  `json:"issuername"`
	LastPrice        float32 `json:"lastprice"`
	PercentualChange float32 `json:"percentualchange"`
	NetChangeSign    string  `json:"netchangesign"`
}
