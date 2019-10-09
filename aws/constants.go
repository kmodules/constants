package aws

import (
	"log"
	"os"
)

const (
	AWS_ACCESS_KEY_ID     = "AWS_ACCESS_KEY_ID"
	AWS_SECRET_ACCESS_KEY = "AWS_SECRET_ACCESS_KEY"
	CA_CERT_DATA          = "CA_CERT_DATA"
)

const (
	AWSCredentialAccessKeyKey = "access_key"
	AWSCredentialSecretKeyKey = "secret_key"
)

func CredentialsFromEnv() map[string][]byte {
	awsAccessKeyId := os.Getenv(AWS_ACCESS_KEY_ID)
	awsSecretAccessKey := os.Getenv(AWS_SECRET_ACCESS_KEY)
	if len(awsAccessKeyId) == 0 || len(awsSecretAccessKey) == 0 {
		log.Println("AWS credentials for empty")
		return map[string][]byte{}
	}

	return map[string][]byte{
		AWSCredentialAccessKeyKey: []byte(awsAccessKeyId),
		AWSCredentialSecretKeyKey: []byte(awsSecretAccessKey),
	}
}
