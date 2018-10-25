// Copyright 2018 The ClientSDK Team Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0-style
// license that can be found in the LICENSE file.

// Author: ClientSDK Team (muharihar)

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

	// the default value for user agent is "RailTicket-B2B", you don't need to setup User-Agent by default
	// this is only for coverage test purpose
	kaiClient.SetUserAgent("RailTicket-B2B")

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

// TestCallGetScheduleLiteFalse is a negative test function for "KAIHttpClient.CallGetScheduleLite" method
func TestCallGetScheduleLiteFalse(t *testing.T) {

	// init http client
	httpClient := makeHTTPClient()

	kaiClient, err := NewKAIHttpClient(httpClient, kaiServerURL, kaiRQID)

	if err != nil {
		log.Fatal(err)
	}

	// test variable
	var vRS *InternalGetScheduleLiteRS

	// test expected values
	// errCode := "002000"
	// errMsg := "Invalid Request. IP Address or Requester ID are not registered"

	// test function
	params := make(map[string]string)

	params["org"] = "BD"
	params["des"] = "GMR"
	currentDate := time.Now().Local()
	params["dep_date"] = currentDate.AddDate(0, 0, 7).Format("20060102")

	vRS, err = kaiClient.CallGetScheduleLite(params, false)

	// test logic
	assert.Equal(t, errCode002000, vRS.ErrCode, "should be equal!")
	assert.Equal(t, errMsg002000, vRS.ErrMsg, "should be equal!")

	assert.Nil(t, nil)
}

// TestCallGetSeatMapFalse is a negative test function for "KAIHttpClient.CallGetSeatMap" method
func TestCallGetSeatMapFalse(t *testing.T) {

	// init http client
	httpClient := makeHTTPClient()

	kaiClient, err := NewKAIHttpClient(httpClient, kaiServerURL, kaiRQID)

	if err != nil {
		log.Fatal(err)
	}

	// test variable
	var vRS *InternalGetSeatMapRS

	// test expected values
	// errCode := "002000"
	// errMsg := "Invalid Request. IP Address or Requester ID are not registered"

	// test function
	params := make(map[string]string)

	params["org"] = "BD"
	params["des"] = "GMR"
	currentDate := time.Now().Local()
	params["dep_date"] = currentDate.AddDate(0, 0, 7).Format("20060102")
	params["train_no"] = "10501"

	vRS, err = kaiClient.CallGetSeatMap(params, false)

	// test logic
	assert.Equal(t, errCode002000, vRS.ErrCode, "should be equal!")
	assert.Equal(t, errMsg002000, vRS.ErrMsg, "should be equal!")

	assert.Nil(t, nil)
}

// TestCallGetSeatMapPerSubClassFalse is a negative test function for "KAIHttpClient.CallGetSeatMapPerSubClass" method
func TestCallGetSeatMapPerSubClassFalse(t *testing.T) {

	// init http client
	httpClient := makeHTTPClient()

	kaiClient, err := NewKAIHttpClient(httpClient, kaiServerURL, kaiRQID)

	if err != nil {
		log.Fatal(err)
	}

	// test variable
	var vRS *InternalGetSeatMapPerSubClassRS

	// test expected values
	// errCode := "002000"
	// errMsg := "Invalid Request. IP Address or Requester ID are not registered"

	// test function
	params := make(map[string]string)

	params["org"] = "BD"
	params["des"] = "GMR"
	currentDate := time.Now().Local()
	params["dep_date"] = currentDate.AddDate(0, 0, 7).Format("20060102")
	params["train_no"] = "10501"
	params["subclass"] = "C"

	vRS, err = kaiClient.CallGetSeatMapPerSubClass(params, false)

	// test logic
	assert.Equal(t, errCode002000, vRS.ErrCode, "should be equal!")
	assert.Equal(t, errMsg002000, vRS.ErrMsg, "should be equal!")

	assert.Nil(t, nil)
}

// TestCallGetSeatNullFalse is a negative test function for "KAIHttpClient.CallGetSeatNull" method
func TestCallGetSeatNullFalse(t *testing.T) {

	// init http client
	httpClient := makeHTTPClient()

	kaiClient, err := NewKAIHttpClient(httpClient, kaiServerURL, kaiRQID)

	if err != nil {
		log.Fatal(err)
	}

	// test variable
	var vRS *InternalGetSeatNullRS

	// test expected values
	// errCode := "002000"
	// errMsg := "Invalid Request. IP Address or Requester ID are not registered"

	// test function
	params := make(map[string]string)

	params["org"] = "BD"
	params["des"] = "GMR"
	currentDate := time.Now().Local()
	params["dep_date"] = currentDate.AddDate(0, 0, 7).Format("20060102")
	params["train_no"] = "10501"

	vRS, err = kaiClient.CallGetSeatNull(params, false)

	// test logic
	assert.Equal(t, errCode002000, vRS.ErrCode, "should be equal!")
	assert.Equal(t, errMsg002000, vRS.ErrMsg, "should be equal!")

	assert.Nil(t, nil)
}

// TestCallGetSeatNullPerSubClassFalse is a negative test function for "KAIHttpClient.CallGetSeatNullPerSubClass" method
func TestCallGetSeatNullPerSubClassFalse(t *testing.T) {

	// init http client
	httpClient := makeHTTPClient()

	kaiClient, err := NewKAIHttpClient(httpClient, kaiServerURL, kaiRQID)

	if err != nil {
		log.Fatal(err)
	}

	// test variable
	var vRS *InternalGetSeatNullPerSubClassRS

	// test expected values
	// errCode := "002000"
	// errMsg := "Invalid Request. IP Address or Requester ID are not registered"

	// test function
	params := make(map[string]string)

	params["org"] = "BD"
	params["des"] = "GMR"
	currentDate := time.Now().Local()
	params["dep_date"] = currentDate.AddDate(0, 0, 7).Format("20060102")
	params["train_no"] = "10501"
	params["subclass"] = "C"

	vRS, err = kaiClient.CallGetSeatNullPerSubClass(params, false)

	// test logic
	assert.Equal(t, errCode002000, vRS.ErrCode, "should be equal!")
	assert.Equal(t, errMsg002000, vRS.ErrMsg, "should be equal!")

	assert.Nil(t, nil)
}

// TestCallGetAgentBalanceFalse is a negative test function for "KAIHttpClient.CallGetAgentBalance" method
func TestCallGetAgentBalanceFalse(t *testing.T) {

	// init http client
	httpClient := makeHTTPClient()

	kaiClient, err := NewKAIHttpClient(httpClient, kaiServerURL, kaiRQID)

	if err != nil {
		log.Fatal(err)
	}

	// test variable
	var vRS *InternalGetAgentBalanceRS

	// test expected values
	// errCode := "002000"
	// errMsg := "Invalid Request. IP Address or Requester ID are not registered"

	// test function
	vRS, err = kaiClient.CallGetAgentBalance(false)

	// test logic
	assert.Equal(t, errCode002000, vRS.ErrCode, "should be equal!")
	assert.Equal(t, errMsg002000, vRS.ErrMsg, "should be equal!")

	assert.Nil(t, nil)
}
