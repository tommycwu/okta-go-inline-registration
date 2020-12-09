package main

import (
	"net/http"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func reqHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	
	var bodyok = `{"commands":[{"type":"com.okta.action.update","value":{"registration": "ALLOW"}}]}`
	
	var bodyerr = `{
		"commands": [{
			"type": "com.okta.action.update",
			"value": {
				"registration": "DENY"
			}
		}],
		"error": {
			"errorSummary": "You specified an invalid email domain",
			"errorCauses": [{
				"errorSummary": "You specified an invalid email domain",
				"reason": "INVALID_EMAIL_DOMAIN"
			}]
		}
	}`

	resp := events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            bodyerr,
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
