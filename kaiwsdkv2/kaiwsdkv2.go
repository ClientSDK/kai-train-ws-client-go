// Copyright 2018 The ClientSDK Team Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0-style
// license that can be found in the LICENSE file.

// Author: ClientSDK Team (muharihar)

package kaiwsdkv2

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// NewKAIHttpClient return new *KAIHttpClient to handle the requests to KAI HTTP Server
func NewKAIHttpClient(httpClient *http.Client, serverURL string, rqid string) (*KAIHttpClient, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	kaiHTTPClient := &KAIHttpClient{
		httpClient:      httpClient,
		serverURL:       serverURL,
		rqid:            rqid,
		headerUserAgent: "RailTicket-B2B",
	}

	return kaiHTTPClient, nil
}

// KAIHttpClient struct hold all the information about KAI HTTP Client,
// request and response of the server
type KAIHttpClient struct {
	httpClient *http.Client //HTTP client used to communicate with the API.

	serverURL       string // KAI Server (URL) Address
	rqid            string // KAI RQID
	headerUserAgent string // KAI http (header) client User Agent

	module string // KAI Module
	action string // KAI Action

	KAIQueryString      string
	KAIResponseBody     []byte
	KAIRealResponseBody []byte
}

// SetUserAgent is a function to set "User-Agent" http header
func (c *KAIHttpClient) SetUserAgent(value string) {
	if strings.TrimSpace(value) != "" {
		c.headerUserAgent = value
	}
}

// Call call's the module and action with Params
func (c *KAIHttpClient) Call(module string, action string, params map[string]string, debug bool) (err error) {
	c.module = module
	c.action = action

	q := url.Values{}
	q.Add("rqid", c.rqid)
	q.Add("app", c.module)
	q.Add("action", c.action)

	for i, v := range params {
		q.Add(i, v)
	}

	c.KAIQueryString = q.Encode()

	httpResp, err := c.doGetRequest(c.KAIQueryString, debug)
	if err != nil {
		return err
	}

	c.KAIRealResponseBody = httpResp
	c.KAIResponseBody = TrasformStandardKAIResponse(httpResp)

	return err
}

// doGetRequest makes new request to the server using c.serverURL, c.rqid,
// and query string parameters
func (c *KAIHttpClient) doGetRequest(params string, debug bool) ([]byte, error) {
	client := c.httpClient

	req, err := http.NewRequest("GET", c.serverURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.headerUserAgent)

	req.URL.RawQuery = params

	if debug {
		DebugHTTPRequest(httputil.DumpRequestOut(req, true))
		//DebugHTTPRequest(httputil.DumpRequest(req, true))
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err == nil && debug == true {
		DebugHTTPResponse(httputil.DumpResponse(resp, true))
	}

	return ioutil.ReadAll(resp.Body)
}

// CallGetOrigination is a function to call KAI "data.get_org" method
func (c *KAIHttpClient) CallGetOrigination(debug bool) (*InternalGetOriginationRS, error) {

	// call to KAI
	params := make(map[string]string)
	err := c.Call("data", "get_org", params, debug)
	if err != nil {
		return nil, err
	}

	// convert to struct
	var vRS GetOriginationRS

	err = json.Unmarshal(c.KAIResponseBody, &vRS)
	if err != nil {
		return nil, err
	}

	// make internal response
	var resp *InternalGetOriginationRS

	resp, err = MakeInternalGetOriginationRS(vRS)
	if err != nil {
		return nil, err
	}

	// send to caller
	return resp, nil
}

// CallGetDestination is a function to call KAI "data.get_des" method
func (c *KAIHttpClient) CallGetDestination(debug bool) (*InternalGetDestinationRS, error) {

	// call to KAI
	params := make(map[string]string)
	err := c.Call("data", "get_des", params, debug)
	if err != nil {
		return nil, err
	}

	// convert to struct
	var vRS GetDestinationRS

	err = json.Unmarshal(c.KAIResponseBody, &vRS)
	if err != nil {
		return nil, err
	}

	// make internal response
	var resp *InternalGetDestinationRS

	resp, err = MakeInternalGetDestinationRS(vRS)
	if err != nil {
		return nil, err
	}

	// send to caller
	return resp, nil
}

// CallGetPayType is a function to call KAI "data.get_pay_type" method
func (c *KAIHttpClient) CallGetPayType(debug bool) (*InternalGetPayTypeRS, error) {

	// call to KAI
	params := make(map[string]string)
	err := c.Call("data", "get_pay_type", params, debug)
	if err != nil {
		return nil, err
	}

	// convert to struct
	var vRS GetPayTypeRS

	err = json.Unmarshal(c.KAIResponseBody, &vRS)
	if err != nil {
		return nil, err
	}

	// make internal response
	var resp *InternalGetPayTypeRS

	resp, err = MakeInternalGetPayTypeRS(vRS)
	if err != nil {
		return nil, err
	}

	// send to caller
	return resp, nil
}

// CallGetSchedule is a function to call KAI "information.get_schedule" method
func (c *KAIHttpClient) CallGetSchedule(params map[string]string, debug bool) (*InternalGetScheduleRS, error) {

	// call to KAI
	// params := make(map[string]string)
	err := c.Call("information", "get_schedule", params, debug)
	if err != nil {
		return nil, err
	}

	// convert to struct
	var vRS GetScheduleRS

	err = json.Unmarshal(c.KAIResponseBody, &vRS)
	if err != nil {
		return nil, err
	}

	// make internal response
	var resp *InternalGetScheduleRS

	resp, err = MakeInternalGetScheduleRS(vRS)
	if err != nil {
		return nil, err
	}

	// send to caller
	return resp, nil
}

// CallGetScheduleV2 is a function to call KAI "information.get_schedule_v2" method
func (c *KAIHttpClient) CallGetScheduleV2(params map[string]string, debug bool) (*InternalGetScheduleV2RS, error) {

	// call to KAI
	// params := make(map[string]string)
	err := c.Call("information", "get_schedule_v2", params, debug)
	if err != nil {
		return nil, err
	}

	// convert to struct
	var vRS GetScheduleV2RS

	err = json.Unmarshal(c.KAIResponseBody, &vRS)
	if err != nil {
		return nil, err
	}

	// make internal response
	var resp *InternalGetScheduleV2RS

	resp, err = MakeInternalGetScheduleV2RS(vRS)
	if err != nil {
		return nil, err
	}

	// send to caller
	return resp, nil
}

// CallGetScheduleLite is a function to call KAI "information.get_schedule_lite" method
func (c *KAIHttpClient) CallGetScheduleLite(params map[string]string, debug bool) (*InternalGetScheduleLiteRS, error) {

	// call to KAI
	// params := make(map[string]string)
	err := c.Call("information", "get_schedule_lite", params, debug)
	if err != nil {
		return nil, err
	}

	// convert to struct
	var vRS GetScheduleLiteRS

	err = json.Unmarshal(c.KAIResponseBody, &vRS)
	if err != nil {
		return nil, err
	}

	// make internal response
	var resp *InternalGetScheduleLiteRS

	resp, err = MakeInternalGetScheduleLiteRS(vRS)
	if err != nil {
		return nil, err
	}

	// send to caller
	return resp, nil
}

// CallGetSeatMap is a function to call KAI "information.get_seat_map" method
func (c *KAIHttpClient) CallGetSeatMap(params map[string]string, debug bool) (*InternalGetSeatMapRS, error) {

	// call to KAI
	// params := make(map[string]string)
	err := c.Call("information", "get_seat_map", params, debug)
	if err != nil {
		return nil, err
	}

	// convert to struct
	var vRS GetSeatMapRS

	err = json.Unmarshal(c.KAIResponseBody, &vRS)
	if err != nil {
		return nil, err
	}

	// make internal response
	var resp *InternalGetSeatMapRS

	resp, err = MakeInternalGetSeatMapRS(vRS)
	if err != nil {
		return nil, err
	}

	// send to caller
	return resp, nil
}

// CallGetSeatMapPerSubClass is a function to call KAI "information.get_seat_map_per_subclass" method
func (c *KAIHttpClient) CallGetSeatMapPerSubClass(params map[string]string, debug bool) (*InternalGetSeatMapPerSubClassRS, error) {

	// call to KAI
	// params := make(map[string]string)
	err := c.Call("information", "get_seat_map_per_subclass", params, debug)
	if err != nil {
		return nil, err
	}

	// convert to struct
	var vRS GetSeatMapPerSubClassRS

	err = json.Unmarshal(c.KAIResponseBody, &vRS)
	if err != nil {
		return nil, err
	}

	// make internal response
	var resp *InternalGetSeatMapPerSubClassRS

	resp, err = MakeInternalGetSeatMapPerSubClassRS(vRS)
	if err != nil {
		return nil, err
	}

	// send to caller
	return resp, nil
}

// CallGetSeatNull is a function to call KAI "information.get_seat_null" method
func (c *KAIHttpClient) CallGetSeatNull(params map[string]string, debug bool) (*InternalGetSeatNullRS, error) {

	// call to KAI
	// params := make(map[string]string)
	err := c.Call("information", "get_seat_null", params, debug)
	if err != nil {
		return nil, err
	}

	// convert to struct
	var vRS GetSeatNullRS

	err = json.Unmarshal(c.KAIResponseBody, &vRS)
	if err != nil {
		return nil, err
	}

	// make internal response
	var resp *InternalGetSeatNullRS

	resp, err = MakeInternalGetSeatNullRS(vRS)
	if err != nil {
		return nil, err
	}

	// send to caller
	return resp, nil
}

// CallGetSeatNullPerSubClass is a function to call KAI "information.get_seat_null_per_subclass" method
func (c *KAIHttpClient) CallGetSeatNullPerSubClass(params map[string]string, debug bool) (*InternalGetSeatNullPerSubClassRS, error) {

	// call to KAI
	// params := make(map[string]string)
	err := c.Call("information", "get_seat_null_per_subclass", params, debug)
	if err != nil {
		return nil, err
	}

	// convert to struct
	var vRS GetSeatNullPerSubClassRS

	err = json.Unmarshal(c.KAIResponseBody, &vRS)
	if err != nil {
		return nil, err
	}

	// make internal response
	var resp *InternalGetSeatNullPerSubClassRS

	resp, err = MakeInternalGetSeatNullPerSubClassRS(vRS)
	if err != nil {
		return nil, err
	}

	// send to caller
	return resp, nil
}

// CallGetAgentBalance is a function to call KAI "information.get_agent_balance" method
func (c *KAIHttpClient) CallGetAgentBalance(debug bool) (*InternalGetAgentBalanceRS, error) {

	// call to KAI
	params := make(map[string]string)
	err := c.Call("information", "get_agent_balance", params, debug)
	if err != nil {
		return nil, err
	}

	// convert to struct
	var vRS GetAgentBalanceRS

	err = json.Unmarshal(c.KAIResponseBody, &vRS)
	if err != nil {
		return nil, err
	}

	// make internal response
	var resp *InternalGetAgentBalanceRS

	resp, err = MakeInternalGetAgentBalanceRS(vRS)
	if err != nil {
		return nil, err
	}

	// send to caller
	return resp, nil
}

// CallGetBalance is a function to call KAI "information.get_balance" method
func (c *KAIHttpClient) CallGetBalance(params map[string]string, debug bool) (*InternalGetBalanceRS, error) {

	// call to KAI
	// params := make(map[string]string)
	err := c.Call("information", "get_balance", params, debug)
	if err != nil {
		return nil, err
	}

	// convert to struct
	var vRS GetBalanceRS

	err = json.Unmarshal(c.KAIResponseBody, &vRS)
	if err != nil {
		return nil, err
	}

	// make internal response
	var resp *InternalGetBalanceRS

	resp, err = MakeInternalGetBalanceRS(vRS)
	if err != nil {
		return nil, err
	}

	// send to caller
	return resp, nil
}

// CallGetBookInfo is a function to call KAI "information.get_book_info" method
func (c *KAIHttpClient) CallGetBookInfo(params map[string]string, debug bool) (*InternalGetBookInfoRS, error) {

	// call to KAI
	// params := make(map[string]string)
	err := c.Call("information", "get_book_info", params, debug)
	if err != nil {
		return nil, err
	}

	// convert to struct
	var vRS GetBookInfoRS

	err = json.Unmarshal(c.KAIResponseBody, &vRS)
	if err != nil {
		return nil, err
	}

	// make internal response
	var resp *InternalGetBookInfoRS

	resp, err = MakeInternalGetBookInfoRS(vRS)
	if err != nil {
		return nil, err
	}

	// send to caller
	return resp, nil
}

// CallGetBookPriceInfo is a function to call KAI "information.get_book_price_info" method
func (c *KAIHttpClient) CallGetBookPriceInfo(params map[string]string, debug bool) (*InternalGetBookPriceInfoRS, error) {

	// call to KAI
	// params := make(map[string]string)
	err := c.Call("information", "get_book_price_info", params, debug)
	if err != nil {
		return nil, err
	}

	// convert to struct
	var vRS GetBookPriceInfoRS

	err = json.Unmarshal(c.KAIResponseBody, &vRS)
	if err != nil {
		return nil, err
	}

	// make internal response
	var resp *InternalGetBookPriceInfoRS

	resp, err = MakeInternalGetBookPriceInfoRS(vRS)
	if err != nil {
		return nil, err
	}

	// send to caller
	return resp, nil
}

// CallBooking is a function to call KAI "transaction.booking" method
func (c *KAIHttpClient) CallBooking(params map[string]string, debug bool) (*InternalBookingRS, error) {

	// call to KAI
	// params := make(map[string]string)
	err := c.Call("transaction", "booking", params, debug)
	if err != nil {
		return nil, err
	}

	// convert to struct
	var vRS BookingRS

	err = json.Unmarshal(c.KAIResponseBody, &vRS)
	if err != nil {
		return nil, err
	}

	// make internal response
	var resp *InternalBookingRS

	resp, err = MakeInternalBookingRS(vRS)
	if err != nil {
		return nil, err
	}

	// send to caller
	return resp, nil
}

// CallBookingWithArvInfo is a function to call KAI "transaction.booking_with_arv_info" method
func (c *KAIHttpClient) CallBookingWithArvInfo(params map[string]string, debug bool) (*InternalBookingWithArvInfoRS, error) {

	// call to KAI
	// params := make(map[string]string)
	err := c.Call("transaction", "booking_with_arv_info", params, debug)
	if err != nil {
		return nil, err
	}

	// convert to struct
	var vRS BookingWithArvInfoRS

	err = json.Unmarshal(c.KAIResponseBody, &vRS)
	if err != nil {
		return nil, err
	}

	// make internal response
	var resp *InternalBookingWithArvInfoRS

	resp, err = MakeInternalBookingWithArvInfoRS(vRS)
	if err != nil {
		return nil, err
	}

	// send to caller
	return resp, nil
}

// CallCancelBook is a function to call KAI "transaction.cancel_book" method
func (c *KAIHttpClient) CallCancelBook(params map[string]string, debug bool) (*InternalCancelBookRS, error) {

	// call to KAI
	// params := make(map[string]string)
	err := c.Call("transaction", "cancel_book", params, debug)
	if err != nil {
		return nil, err
	}

	// convert to struct
	var vRS CancelBookRS

	err = json.Unmarshal(c.KAIResponseBody, &vRS)
	if err != nil {
		return nil, err
	}

	// make internal response
	var resp *InternalCancelBookRS

	resp, err = MakeInternalCancelBookRS(vRS)
	if err != nil {
		return nil, err
	}

	// send to caller
	return resp, nil
}

// CallManualSeat is a function to call KAI "transaction.manual_seat" method
func (c *KAIHttpClient) CallManualSeat(params map[string]string, debug bool) (*InternalManualSeatRS, error) {

	// call to KAI
	// params := make(map[string]string)
	err := c.Call("transaction", "manual_seat", params, debug)
	if err != nil {
		return nil, err
	}

	// convert to struct
	var vRS ManualSeatRS

	err = json.Unmarshal(c.KAIResponseBody, &vRS)
	if err != nil {
		return nil, err
	}

	// make internal response
	var resp *InternalManualSeatRS

	resp, err = MakeInternalManualSeatRS(vRS)
	if err != nil {
		return nil, err
	}

	// send to caller
	return resp, nil
}

// CallUpdatePax is a function to call KAI "transaction.update_pax" method
func (c *KAIHttpClient) CallUpdatePax(params map[string]string, debug bool) (*InternalUpdatePaxRS, error) {

	// call to KAI
	// params := make(map[string]string)
	err := c.Call("transaction", "update_pax", params, debug)
	if err != nil {
		return nil, err
	}

	// convert to struct
	var vRS UpdatePaxRS

	err = json.Unmarshal(c.KAIResponseBody, &vRS)
	if err != nil {
		return nil, err
	}

	// make internal response
	var resp *InternalUpdatePaxRS

	resp, err = MakeInternalUpdatePaxRS(vRS)
	if err != nil {
		return nil, err
	}

	// send to caller
	return resp, nil
}
