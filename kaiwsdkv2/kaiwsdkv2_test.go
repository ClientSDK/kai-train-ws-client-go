package kaiwsdkv2

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	kaiServerURL = "http://ws.demo.kai.sqiva.com/"
	kaiRQID      = "your-rqid-demo-value"
)

// makeHTTPClient is a function to init http client
func makeHTTPClient() *http.Client {
	// Access via proxy if needed
	//proxyURL, _ := url.Parse("proxy-ip-address:proxy-port")
	//proxyURL, _ := url.Parse("http://proxy-user:proxy-password@proxy-ip-address:proxy-port")

	// Initiate http client with transport
	httpClient := &http.Client{Transport: &http.Transport{}}
	//httpClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}

	return httpClient
}

// TestCallFalse is a negative test function for "KAIHttpClient.Call" method
func TestCallFalse(t *testing.T) {

	// init http client
	httpClient := makeHTTPClient()

	kaiClient, err := NewKAIHttpClient(httpClient, kaiServerURL, kaiRQID)

	if err != nil {
		log.Fatal(err)
	}

	// test function
	params := make(map[string]string)
	err = kaiClient.Call("data", "get_org", params, false)

	// test expected values
	errCode := "002000"
	errMsg := "Invalid Request. IP Address or Requester ID are not registered"

	// test variable
	var vRS StdKAIErrorMessage

	// test value
	err = json.Unmarshal(kaiClient.KAIResponseBody, &vRS)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, errMsg, vRS.ErrMsg, "should be equal!")

	assert.Nil(t, nil)
}

// TestCallGetOriginationFalse is a negative test function for "KAIHttpClient.CallGetOrigination" method
func TestCallGetOriginationFalse(t *testing.T) {

	// init http client
	httpClient := makeHTTPClient()

	kaiClient, err := NewKAIHttpClient(httpClient, kaiServerURL, kaiRQID)

	if err != nil {
		log.Fatal(err)
	}

	// test variable
	var vRS *InternalGetOriginationRS

	// test expected values
	errCode := "002000"
	errMsg := "Invalid Request. IP Address or Requester ID are not registered"

	// test function
	vRS, err = kaiClient.CallGetOrigination(false)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, errMsg, vRS.ErrMsg, "should be equal!")

	assert.Nil(t, nil)
}
