# Module/Method DATA:GET_ORG (data.get_org)

Module "Get Origination" (DATA:GET_ORG) is a service method to retreive list of origination (departure location) information from KAI Train (RailTicket) Web Service [[1](https://railticket.kereta-api.co.id/)].

The following are the sections available in this guide.

- [Use Case](#use-case)
- [Prerequisites](#prerequisites)
- [Implementation](#implementation)
- [Build and Running](#build-and-running)

## Use Case
Let’s retrive origination departure location from KAI Train Web Service Endpoint.

![Get Origination Diagram](images/data.get_org.png "Get Origination Diagram")


## Prerequisites

- [KAI Train Web Service Client for Go (kaiwsdkv2 GoLang package) ](https://github.com/ClientSDK/kai-train-ws-client-go)

```Go
go get github.com/ClientSDK/kai-train-ws-client-go/kaiwsdkv2
```

- A Text Editor or an IDE

### KAI Agent requirements
- KAI Train (RailTicket) Agent Credential Account (RQID)
- KAI Train (RailTicket) Web Service Access (IP Whitelist) ( [Production Server](https://railticket.kereta-api.co.id), [Demo Server](http://ws.demo.kai.sqiva.com))

## Implementation

> If you want to skip the basics, you can download the git repo and directly move to the "Build and Running" section by skipping  "Implementation" section.

### Example structure

Go is a complete programming language that supports custom project structures. Let's use the following package structure for this example.

```
get-org
   ├── build_and_run.sh
   └── main.go
```

- Create the above directories in your local machine and also create empty `main.go` and `build_and_run.sh` files.


### Developing the application

Let's make a simple application for getting list of origination (departure location) information using `kaiwsdkv2` package. 

##### Main code for "data.get_org" (main.go)
```go
// Copyright 2018 The ClientSDK Team Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0-style
// license that can be found in the LICENSE file.

// Author: ClientSDK Team (muharihar)

package main

import (
	"crypto/tls"
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

	// Initiate transport with proxy and skip TLS
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxyURL),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Initiate transport without proxy and skip TLS
	// tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }

	// Using Transport
	httpClient := &http.Client{Transport: tr}

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

	// if you want to retreive KAI Origin Response
	// fmt.Println(string(kaiClient.KAIRealResponseBody))

	json, _ := json.Marshal(vRS)
	fmt.Println(string(json))
}

```

##### Bash code for building and running the example application (build_and_run.sh)
```bash
echo "Clean..."
rm ./data.get_org
echo "Build..."
go build -o data.get_org main.go 
echo "Build Done..."
echo "Run..."
./data.get_org > data.get_org-rs.json
echo "Done."

```


## Build and Running

You can build and running by execute the "build_and_run.sh" bash files. 

```bash
   $ sh build_and_run.sh 
```

After the application is running, you will get the json response in `data.get_org-rs.json` files.

## Sample Response

### Sample KAI Response:

```go
    // Get KAI Response Raw from KAIHttpClient Struct 
    KAIHttpClient.KAIRealResponseBody
```

```json
{
    "err_code": 0,
    "origination": [
        [
            "AK",
            "ANGKE"
        ],
        [
            "YK",
            "YOGYAKARTA"
        ]
    ]
}
```

### Sample Internal Response:

```go
    // Get Internal Response Raw from KAIHttpClient Struct 
    KAIHttpClient.KAIResponseBody
```

```json
 {
    "errCode": "0",
    "errMsg": null,
    "return": [
        {
            "OriginCode": "AK",
            "OriginName": "ANGKE"
        },
        {
            "OriginCode": "YK",
            "OriginName": "YOGYAKARTA"
        }
    ]
}
```