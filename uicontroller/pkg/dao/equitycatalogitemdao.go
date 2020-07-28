package dao

import (
	"github.com/pwestlake/aemo/userservice/pkg/config"
	"github.com/pwestlake/equity-fund/uicontroller/pkg/domain"
	"log"
	"time"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// EquityCatalogItemDAO ...
// DAO for an EquityCatalogItem
type EquityCatalogItemDAO struct {
	config   config.Config
	endpoint string
	region   string
}

// NewEquityCatalogItemDAO ...
// Create function for a NewUserDao
func NewEquityCatalogItemDAO() EquityCatalogItemDAO {
	config := config.NewConfig(nil)
	endpoint, err := config.GetString("dynamoDB.endpoint")
	if err != nil {
		log.Print(err)
	}

	region, err := config.GetString("dynamoDB.region")
	if err != nil {
		log.Print(err)
	}

	return EquityCatalogItemDAO{config: config, endpoint: endpoint, region: region}
}

// PutEquityCatalogItem ...
// DAO method to persist a new EquityCatalogItem in the database
func (s *EquityCatalogItemDAO) PutEquityCatalogItem(equityCatalogItem *domain.EquityCatalogItem) error {
	dbSession := session.Must(session.NewSession())
	client := dynamodb.New(dbSession, aws.NewConfig().WithEndpoint(s.endpoint).WithRegion(s.region))

	if (equityCatalogItem.DateCreated == time.Time{}) {
		equityCatalogItem.DateCreated = time.Now()
	}

	av, err := dynamodbattribute.MarshalMap(equityCatalogItem)
	if err != nil {
		log.Printf("Error marshalling EquityCatalogItem type")
	} else {
		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String("EquityCatalog")}

		_, err = client.PutItem(input)
	}

	return err
}

// GetEquityCatalogItem ...
// DAO method to retrieve an EquityCatalogItem with the given id f
func (s *EquityCatalogItemDAO) GetEquityCatalogItem(id string, equityCatalogItem *domain.EquityCatalogItem) error {
	dbSession := session.Must(session.NewSession())
	client := dynamodb.New(dbSession, aws.NewConfig().WithEndpoint(s.endpoint).WithRegion(s.region))

	result, err := client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("EquityCatalog"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id)}}})
	if err != nil {
		return err
	}

	if result.Item == nil {
		return errors.New("Item not found")
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, equityCatalogItem)
	return err
}

// GetEquityCatalogItems ...
// DAO method to return an array of all EquityCatalogItems 
func (s *EquityCatalogItemDAO) GetEquityCatalogItems() (*[]domain.EquityCatalogItem, error) {
	var equityCatalogItems []domain.EquityCatalogItem
	dbSession := session.Must(session.NewSession())
	client := dynamodb.New(dbSession, aws.NewConfig().WithEndpoint(s.endpoint).WithRegion(s.region))

	params := &dynamodb.ScanInput{
		TableName: aws.String("EquityCatalog")}

	result, err := client.Scan(params)

	if err == nil {
		equityCatalogItems = make([]domain.EquityCatalogItem, len(result.Items))
		for i, item := range result.Items {
			equityCatalogItem := domain.EquityCatalogItem{}
			err = dynamodbattribute.UnmarshalMap(item, &equityCatalogItem)
			if err != nil {
				break;
			} else {
				equityCatalogItems[i] = equityCatalogItem
			}
		}
	}
	return &equityCatalogItems, err
}