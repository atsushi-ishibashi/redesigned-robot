package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
	ErrNon404Response = errors.New("Not found")

	pathRe = regexp.MustCompile(`^/users/\d+$`)
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	path := request.Path
	if !pathRe.MatchString(path) {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
		}, ErrNon404Response
	}

	userIDStr := strings.Replace(path, "/users/", "", 1)
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
		}, err
	}

	ui := generateUserInfo(userID)
	uijson, err := json.Marshal(ui)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(uijson),
		StatusCode: 200,
	}, nil
}

type UserInfo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func generateUserInfo(userID int64) UserInfo {
	return UserInfo{
		ID:   userID,
		Name: fmt.Sprintf("ユーザー%d", userID),
	}
}

func main() {
	lambda.Start(handler)
}
