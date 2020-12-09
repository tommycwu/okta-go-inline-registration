package main

import (
	"net/http"
	"log"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func reqHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body = `{
		"commands": [{
			"type": "com.okta.action.update",
			"value": {
				"registration": "DENY"
			}
		}],
		"error": {
			"errorSummary": "Errors were found in the user profile",
			"errorCauses": [{
				"errorSummary": "You specified an invalid email domain",
				"reason": "INVALID_EMAIL_DOMAIN"
			}]
		}
	}`
	log.Println(request.Body.userProfile)

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
