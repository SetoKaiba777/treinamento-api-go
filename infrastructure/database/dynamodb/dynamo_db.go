package dynamodb

import (
	"context"
	"payments-go/adapter/repository"
	"payments-go/core/domain"
	"payments-go/infrastructure/database/dynamodb/entity"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const tableName = "Payments"

type DynamoDBClient struct {
	awsRegion   string
	awsEndpoint string
	client      *dynamodb.Client
}

var _ repository.PaymentsRepositoryDb = (*DynamoDBClient)(nil)

func NewDynamoDBClient(awsRegion, awsEndpoint string) *DynamoDBClient {
	dbClient := DynamoDBClient{
		awsRegion:   awsRegion,
		awsEndpoint: awsEndpoint,
	}
	dbClient.client = dbClient.loadDynamoDDClient()
	return &dbClient
}

func (d DynamoDBClient) loadDynamoDDClient() *dynamodb.Client {
	awsconfig, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(aws.
		EndpointResolverWithOptionsFunc(func(_, _ string, _ ...interface{}) (aws.Endpoint, error) {
			if d.awsEndpoint != "" {
				return aws.Endpoint{
					PartitionID:   "aws",
					URL:           d.awsEndpoint,
					SigningRegion: d.awsRegion,
				}, nil
			}
			return aws.Endpoint{}, &aws.EndpointNotFoundError{}
		})),
	)
	if err != nil {
		panic(err)
	}

	return dynamodb.NewFromConfig(awsconfig, func(op *dynamodb.Options) {
		op.Region = awsconfig.Region
	})
}

func (d *DynamoDBClient) FindById(ctx context.Context, id string) (domain.Payment, error) {
	out, err := d.client.GetItem(ctx,&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil{
		return domain.Payment{}, err
	}

	var entity entity.PaymentEntity
	err = attributevalue.UnmarshalMap(out.Item,&entity)
	if err!= nil{
		return domain.Payment{},err
	}

	return entity.PaymentEntityToPayment(), nil
}

func (d *DynamoDBClient) PutItem(ctx context.Context, item domain.Payment) (domain.Payment, error) {
	paymentEtity := entity.PaymentToPaymentEntity(item)

	var avMap, err = attributevalue.MarshalMap(paymentEtity)
	if err !=nil{
		return domain.Payment{}, err
	}

	_, err = d.client.PutItem(ctx,&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: avMap,
	})
	if err != nil{
		return domain.Payment{}, err
	}
	
	return item, nil
}
