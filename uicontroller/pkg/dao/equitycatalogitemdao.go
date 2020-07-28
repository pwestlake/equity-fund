package dao

import (
	"github.com/pwestlake/aemo/userservice/pkg/config"
	"github.com/pwestlake/aemo/userservice/pkg/domain"
	"log"
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
func (s *EquityCatalogItemDAO) PutEquityCatalogItem(equityCatalogItem *EquityCatalogItem) error {
	dbSession := session.Must(session.NewSession())
	client := dynamodb.New(dbSession, aws.NewConfig().WithEndpoint(s.endpoint).WithRegion(s.region))

	if (user.DateCreated == time.Time{}) {
		user.DateCreated = time.Now()
	}

	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		log.Printf("Error marshalling User type")
	} else {
		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String("Users")}

		_, err = client.PutItem(input)
	}

	return err
}

