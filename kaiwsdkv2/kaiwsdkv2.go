package kaiwsdkv2

import (
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
