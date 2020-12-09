package main

import (
	"net/http"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func reqHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body = `commands = [{"type":"com.okta.action.update","value":{"registration": "DENY"}}];`
	resp := events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            body,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	return resp, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func main() {
	lambda.Start(reqHandler)
}
