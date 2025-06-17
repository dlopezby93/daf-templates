package clients

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type SSMClient struct {
	ssmClient *ssm.Client
}

var ssmInstance *SSMClient

func NewSSMClient() *SSMClient {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load config, %v", err)
	}

	ssmInstance = &SSMClient{ssmClient: ssm.NewFromConfig(cfg)}

	return ssmInstance
}

func (s *SSMClient) GetClient() *ssm.Client {
	return s.ssmClient
}
