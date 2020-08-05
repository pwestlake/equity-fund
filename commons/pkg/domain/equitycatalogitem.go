package domain

import (
	"time"
)

// EquityCatalogItem ...
// Domain struct for an EquityCatalogItem
type EquityCatalogItem struct {
	ID             string    `json:"id"`
	Symbol         string    `json:"symbol"`
	LSEtidm        string    `json:"lsetidm"`
	LSEissuername  string    `json:"lseissuername"`
	LSEtabID       string    `json:"lsetabid"`
	LSEcomponentID string    `json:"lsecomponentid"`
	Datasource	   string	 `json:"datasource"`
	DateCreated    time.Time `json:"datecreated"`
	LastModified   time.Time `json:"lastmodified"`
}
