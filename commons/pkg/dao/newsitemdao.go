package dao

import (
	"github.com/google/uuid"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pwestlake/equity-fund/commons/pkg/domain"
	"log"
	"github.com/pwestlake/aemo/userservice/pkg/config"

)

// NewsItemDAO ...
// DAO for an EndOfDayItem
type NewsItemDAO struct {
	config   config.Config
	endpoint string
	region   string
}

// NewNewsItemDAO ...
// Create function for an EndOfDayItemDAO
func NewNewsItemDAO() NewsItemDAO {
	config := config.NewConfig(nil)
	endpoint, err := config.GetString("dynamoDB.endpoint")
	if err != nil {
		log.Print(err)
	}

	region, err := config.GetString("dynamoDB.region")
	if err != nil {
		log.Print(err)
	}

	return NewsItemDAO{config: config, endpoint: endpoint, region: region}
}

// PutNewsItems ...
// Store the given array of NewsItems in the database
func (s *NewsItemDAO) PutNewsItems(items *[]domain.NewsItem) error{
	dbSession := session.Must(session.NewSession())
	client := dynamodb.New(dbSession, aws.NewConfig().WithEndpoint(s.endpoint).WithRegion(s.region))

	var err error
	for _, v := range *items {
		v.ID = uuid.New().String()
		av, err := dynamodbattribute.MarshalMap(v)
		if err != nil {
			log.Printf("Error marshalling NewsItem type")
			break
		} else {
			input := &dynamodb.PutItemInput{
				Item:      av,
				TableName: aws.String("NewsItems")}

			_, err = client.PutItem(input)
		}
	}

	return err
}

// GetLatestItem ...
// Retrieve the latest item for the given key id
func (s *NewsItemDAO) GetLatestItem(id string) (*domain.NewsItem, error){
	var newsItem = domain.NewsItem{}
	dbSession := session.Must(session.NewSession())
	client := dynamodb.New(dbSession, aws.NewConfig().WithEndpoint(s.endpoint).WithRegion(s.region))

	expressionAttributeValues := map[string]*dynamodb.AttributeValue {
		":id": &dynamodb.AttributeValue{S: aws.String(id)},
	}

	queryInput := dynamodb.QueryInput {
		TableName: aws.String("NewsItems"),
		ExpressionAttributeValues: expressionAttributeValues,
		Limit: aws.Int64(1),
		ScanIndexForward: aws.Bool(false),
		KeyConditionExpression: aws.String("id = :id"),
	}

	resp, err := client.Query(&queryInput)
	if err != nil {
		return &domain.NewsItem{}, err
	}

	if *resp.Count == 0 {
		return nil, errors.New("Item not found")
	}

	err = dynamodbattribute.UnmarshalMap(resp.Items[0],  &newsItem)
	
	return &newsItem, err
}

// GetNewsItems ...
// Return count news items from the given offset with the given id. All items if the id is nil
func (s *NewsItemDAO) GetNewsItems(count int, offset *domain.NewsItem, id *string) (*[]domain.NewsItem, error) {
	var newsItems = []domain.NewsItem{}
	dbSession := session.Must(session.NewSession())
	client := dynamodb.New(dbSession, aws.NewConfig().WithEndpoint(s.endpoint).WithRegion(s.region))

	params := &dynamodb.ScanInput{
		TableName: aws.String("NewsItems"),
		Limit: aws.Int64(int64(count)),
	}

	if id != nil {
		filter := expression.Name("id").Equal(expression.Value(*id))
		expr, err := expression.NewBuilder().WithFilter(filter).Build()
		if err != nil {
			return nil, err
		}

		params.ExpressionAttributeNames = expr.Names()
		params.ExpressionAttributeValues = expr.Values()
		params.FilterExpression = expr.Filter()
	}
	
	if (offset != nil) {
		exclusiveStartKeyMap := map[string]*dynamodb.AttributeValue {
			":id": &dynamodb.AttributeValue{S: aws.String(offset.ID)},
			":datetime": &dynamodb.AttributeValue{S: aws.String(offset.DateTime.Format("2006-01-02T15:04:05Z"))},
		}
		params.ExclusiveStartKey = exclusiveStartKeyMap
	}

	complete := false
	for !complete {
		result, err := client.Scan(params)
		if err != nil {
			return nil, err
		}

		items := []domain.NewsItem{}
		err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &items)
		if err != nil {
			return nil, err
		}

		newsItems = append(newsItems, items...)
		
		if result.LastEvaluatedKey != nil {
			params.ExclusiveStartKey = result.LastEvaluatedKey
		} else {
			complete = true
		}
	}

	return &newsItems, nil
}