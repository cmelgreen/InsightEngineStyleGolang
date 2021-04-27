package main

import (	
	"context"
	"encoding/base64"
	"net/http"
	"io/ioutil"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	url := request.QueryStringParameters["url"]

	response, err := http.Get(url)
    if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 404,
		}, err
    }
	defer response.Body.Close()
	
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 404,
		}, err
    }

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Access-Control-Allow-Headers": "*",
		},
		Body:	base64.StdEncoding.EncodeToString(bytes),
		IsBase64Encoded: true,
	}, nil
}
