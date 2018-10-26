// Copyright 2018 The ClientSDK Team Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0-style
// license that can be found in the LICENSE file.

// Author: ClientSDK Team (muharihar)

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/ClientSDK/kai-train-ws-client-go/kaiwsdkv2"
)

const (
	kaiServer = "https://railticket.kereta-api.co.id/" // Production Server
	// kaiServer = "http://ws.demo.kai.sqiva.com/"		// Demo Server
	kaiRQID = "YOUR-KAI-AGENT-CREDENTIAL-RQID"
)

func makeHTTPClient() *http.Client {
	// Access via proxy if needed
	proxyURL, _ := url.Parse("http://proxy-ip-address:proxy-port")
	//proxyURL, _ := url.Parse("http://proxy-user:proxy-password@proxy-ip-address:proxy-port")

	// Initiate http client with transport
	// httpClient := &http.Client{Transport: &http.Transport{}}
	// Using Proxy if needed
	httpClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}

	return httpClient
}

func main() {

	// init http client
	httpClient := makeHTTPClient()

	// Initiate NewKAIHttpClient version 2
	kaiClient, err := kaiwsdkv2.NewKAIHttpClient(httpClient, kaiServer, kaiRQID)
	if err != nil {
		log.Fatal(err)
	}

	// call KAI web service method
	callGetOrigination(kaiClient)
}

func callGetOrigination(kaiClient *kaiwsdkv2.KAIHttpClient) {

	// params := make(map[string]string)
	vRS, err := kaiClient.CallGetOrigination(false)

	if err != nil {
		log.Fatal(err)
	}

	// sample how to Access Response
	// fmt.Println(vRS.ErrCode)
	// fmt.Println(vRS.ErrMsg)
	// fmt.Println(vRS.Return[0].OriginCode)
	// fmt.Println(vRS.Return[0].OriginName)

	json, _ := json.Marshal(vRS)
	fmt.Println(string(json))
}
