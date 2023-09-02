package external

import (
	"context"
	"fmt"
	"net/http"

	"testing-nextalent/pkg/httpclient"
)

// Data ...
type Data struct {
	client *httpclient.Client
}

// New ...
func New(client *httpclient.Client) Data {
	d := Data{
		client: client,
	}
	return d
}

// CheckAuth ...
func (d Data) GetTimeZone(ctx context.Context, headers http.Header, timeZone string) (interface{}, error) {

	var endpoint = "https://timeapi.io/api/Time/current/zone?timeZone=" + timeZone
	// var body = map[string]interface{}{
	// 	"Partner_TipeID": code,
	// }
	var response interface{}
	//fmt.Println(body)
	_, err := d.client.GetJSON(ctx, endpoint, headers, &response)
	if err != nil {
		fmt.Println(err)
		return response, err
	}

	return response, nil
}
