package kaiwsdkv2

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	kaiServerURL = "http://ws.demo.kai.sqiva.com/"
	kaiRQID      = "your-rqid-demo-value"

	errCode002000 = "002000"
	errMsg002000  = "Invalid Request. IP Address or Requester ID are not registered"
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
	// errCode := "002000"
	// errMsg := "Invalid Request. IP Address or Requester ID are not registered"

	// test variable
	var vRS StdKAIErrorMessage

	// test value
	err = json.Unmarshal(kaiClient.KAIResponseBody, &vRS)

	// test logic
	assert.Equal(t, errCode002000, vRS.ErrCode, "should be equal!")
	assert.Equal(t, errMsg002000, vRS.ErrMsg, "should be equal!")

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
	// errCode := "002000"
	// errMsg := "Invalid Request. IP Address or Requester ID are not registered"

	// test function
	vRS, err = kaiClient.CallGetOrigination(false)

	// test logic
	assert.Equal(t, errCode002000, vRS.ErrCode, "should be equal!")
	assert.Equal(t, errMsg002000, vRS.ErrMsg, "should be equal!")

	assert.Nil(t, nil)
}

// TestCallGetDestinationFalse is a negative test function for "KAIHttpClient.CallGetDestination" method
func TestCallGetDestinationFalse(t *testing.T) {

	// init http client
	httpClient := makeHTTPClient()

	kaiClient, err := NewKAIHttpClient(httpClient, kaiServerURL, kaiRQID)

	if err != nil {
		log.Fatal(err)
	}

	// test variable
	var vRS *InternalGetDestinationRS

	// test expected values
	// errCode := "002000"
	// errMsg := "Invalid Request. IP Address or Requester ID are not registered"

	// test function
	vRS, err = kaiClient.CallGetDestination(false)

	// test logic
	assert.Equal(t, errCode002000, vRS.ErrCode, "should be equal!")
	assert.Equal(t, errMsg002000, vRS.ErrMsg, "should be equal!")

	assert.Nil(t, nil)
}

// TestCallGetPayTypeFalse is a negative test function for "KAIHttpClient.CallGetPayType" method
func TestCallGetPayTypeFalse(t *testing.T) {

	// init http client
	httpClient := makeHTTPClient()

	kaiClient, err := NewKAIHttpClient(httpClient, kaiServerURL, kaiRQID)

	if err != nil {
		log.Fatal(err)
	}

	// test variable
	var vRS *InternalGetPayTypeRS

	// test expected values
	// errCode := "002000"
	// errMsg := "Invalid Request. IP Address or Requester ID are not registered"

	// test function
	vRS, err = kaiClient.CallGetPayType(false)

	// test logic
	assert.Equal(t, errCode002000, vRS.ErrCode, "should be equal!")
	assert.Equal(t, errMsg002000, vRS.ErrMsg, "should be equal!")

	assert.Nil(t, nil)
}

// TestCallGetScheduleFalse is a negative test function for "KAIHttpClient.CallGetSchedule" method
func TestCallGetScheduleFalse(t *testing.T) {

	// init http client
	httpClient := makeHTTPClient()

	kaiClient, err := NewKAIHttpClient(httpClient, kaiServerURL, kaiRQID)

	if err != nil {
		log.Fatal(err)
	}

	// test variable
	var vRS *InternalGetScheduleRS

	// test expected values
	// errCode := "002000"
	// errMsg := "Invalid Request. IP Address or Requester ID are not registered"

	// test function
	params := make(map[string]string)

	params["org"] = "BD"
	params["des"] = "GMR"
	currentDate := time.Now().Local()
	params["dep_date"] = currentDate.AddDate(0, 0, 7).Format("20060102")

	vRS, err = kaiClient.CallGetSchedule(params, false)

	// test logic
	assert.Equal(t, errCode002000, vRS.ErrCode, "should be equal!")
	assert.Equal(t, errMsg002000, vRS.ErrMsg, "should be equal!")

	assert.Nil(t, nil)
}

// TestCallGetScheduleV2False is a negative test function for "KAIHttpClient.CallGetScheduleV2" method
func TestCallGetScheduleV2False(t *testing.T) {

	// init http client
	httpClient := makeHTTPClient()

	kaiClient, err := NewKAIHttpClient(httpClient, kaiServerURL, kaiRQID)

	if err != nil {
		log.Fatal(err)
	}

	// test variable
	var vRS *InternalGetScheduleV2RS

	// test expected values
	// errCode := "002000"
	// errMsg := "Invalid Request. IP Address or Requester ID are not registered"

	// test function
	params := make(map[string]string)

	params["org"] = "BD"
	params["des"] = "GMR"
	currentDate := time.Now().Local()
	params["dep_date"] = currentDate.AddDate(0, 0, 7).Format("20060102")

	vRS, err = kaiClient.CallGetScheduleV2(params, false)

	// test logic
	assert.Equal(t, errCode002000, vRS.ErrCode, "should be equal!")
	assert.Equal(t, errMsg002000, vRS.ErrMsg, "should be equal!")

	assert.Nil(t, nil)
}
