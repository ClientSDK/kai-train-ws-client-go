# Module/Method INFORMATION:GET_SCHEDULE_LITE (information.get_schedule_lite)

Module "Get Schedule" (INFORMATION:GET_SCHEDULE_LITE) is a service method to retreive list of schedule (in lite version) information from KAI Train (RailTicket) Web Service [[1](https://railticket.kereta-api.co.id/)].

The following are the sections available in this guide.

- [Use Case](#use-case)
- [Prerequisites](#prerequisites)
- [Implementation](#implementation)
- [Build and Running](#build-and-running)

## Use Case
Let’s retrieve schedule information from KAI Train Web Service Endpoint.

![Get Schedule Lite Diagram](images/information.get_schedule_lite.png "Get Schedule Lite Diagram")


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
get-schedule-lite
   ├── build_and_run.sh
   └── main.go
```

- Create the above directories in your local machine and also create empty `main.go` and `build_and_run.sh` files.


### Developing the application

Let's make a simple application for getting list of schedule information using `kaiwsdkv2` package. 

##### Main code for "information.get_schedule_lite" (main.go)
```go
package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

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
	callGetScheduleLite(kaiClient)
}

func callGetScheduleLite(kaiClient *kaiwsdkv2.KAIHttpClient) {

	params := make(map[string]string)

	params["org"] = "BD"
	params["des"] = "GMR"
	currentDate := time.Now().Local()
	params["dep_date"] = currentDate.AddDate(0, 0, 7).Format("20060102")

	vRS, err := kaiClient.CallGetScheduleLite(params, false)

	if err != nil {
		log.Fatal(err)
	}

	// sample how to Access Response
	// fmt.Println(vRS.ErrCode)
	// fmt.Println(vRS.ErrMsg)
	// fmt.Println(vRS.Return.Origin)
	// fmt.Println(vRS.Return.Destination)
	// fmt.Println(vRS.Return.DepartureDate)
	// fmt.Println(vRS.Return.Schedules[0].TrainName)
	// fmt.Println(vRS.Return.Schedules[0].AvailSubClass[0].SubClass)

	// if you want to retreive KAI Origin Response
	// fmt.Println(string(kaiClient.KAIRealResponseBody))

	json, _ := json.Marshal(vRS)
	fmt.Println(string(json))
}

```

##### Bash code for building and running the example application (build_and_run.sh)
```bash
echo "Clean..."
rm ./information.get_schedule_lite
echo "Build..."
go build -o information.get_schedule_lite main.go 
echo "Build Done..."
echo "Run..."
./information.get_schedule_lite > information.get_schedule_lite-rs.json
echo "Done."

```


## Build and Running

You can build and running by execute the "build_and_run.sh" bash files. 

```bash
   $ sh build_and_run.sh 
```

After the application is running, you will get the json response in `information.get_schedule_lite-rs.json` files.

## Sample Response

### Sample KAI Response:

```go
    // Get KAI Response Raw from KAIHttpClient Struct 
    KAIHttpClient.KAIRealResponseBody
```

```json
{
    "err_code": 0,
    "org": "BD",
    "des": "GMR",
    "dep_date": "20191201",
    "schedule": [
        [
            "10501",
            "ARGO PARAHYANGAN PREMIUM",
            "20191201",
            "20191201",
            "0415",
            "0725",
            [
                [
                    "C",
                    0,
                    "K",
                    100000,
                    0,
                    0
                ]
            ]
        ],
        [
            "77A",
            "ARGO GOPAR",
            "20191201",
            "20191201",
            "1200",
            "1500",
            [
                [
                    "A",
                    0,
                    "E",
                    100000,
                    0,
                    0
                ],
                [
                    "B",
                    0,
                    "B",
                    90000,
                    0,
                    0
                ],
                [
                    "C",
                    0,
                    "K",
                    60000,
                    0,
                    0
                ]
            ]
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
    "return": {
        "origin": "BD",
        "destination": "GMR",
        "departureDate": "20191201",
        "schedules": [
            {
                "trainNo": "10501",
                "trainName": "ARGO PARAHYANGAN PREMIUM",
                "departureDate": "20191201",
                "arriveDate": "20191201",
                "departureTime": "0415",
                "arriveTime": "0725",
                "availSubClass": [
                    {
                        "subClass": "C",
                        "seatAvailable": 0,
                        "seatClass": "K",
                        "adultPrice": 100000,
                        "childPrice": 0,
                        "infantPrice": 0
                    }
                ]
            },
            {
                "trainNo": "1111",
                "trainName": "ARGO PARAHYANGAN",
                "departureDate": "20191201",
                "arriveDate": "20191201",
                "departureTime": "0500",
                "arriveTime": "0820",
                "availSubClass": [
                    {
                        "subClass": "A",
                        "seatAvailable": 0,
                        "seatClass": "E",
                        "adultPrice": 140000,
                        "childPrice": 0,
                        "infantPrice": 0
                    }
                ]
            },
            {
                "trainNo": "710",
                "trainName": "RANGKAS JAYA",
                "departureDate": "20191201",
                "arriveDate": "20191201",
                "departureTime": "0800",
                "arriveTime": "1115",
                "availSubClass": [
                    {
                        "subClass": "C",
                        "seatAvailable": 0,
                        "seatClass": "K",
                        "adultPrice": 80000,
                        "childPrice": 0,
                        "infantPrice": 0
                    }
                ]
            },
            {
                "trainNo": "77A",
                "trainName": "ARGO GOPAR",
                "departureDate": "20191201",
                "arriveDate": "20191201",
                "departureTime": "1200",
                "arriveTime": "1500",
                "availSubClass": [
                    {
                        "subClass": "A",
                        "seatAvailable": 0,
                        "seatClass": "E",
                        "adultPrice": 100000,
                        "childPrice": 0,
                        "infantPrice": 0
                    },
                    {
                        "subClass": "B",
                        "seatAvailable": 0,
                        "seatClass": "B",
                        "adultPrice": 90000,
                        "childPrice": 0,
                        "infantPrice": 0
                    },
                    {
                        "subClass": "C",
                        "seatAvailable": 0,
                        "seatClass": "K",
                        "adultPrice": 60000,
                        "childPrice": 0,
                        "infantPrice": 0
                    }
                ]
            },
            {
                "trainNo": "P05",
                "trainName": "ARGO PARAHYANGAN",
                "departureDate": "20191201",
                "arriveDate": "20191201",
                "departureTime": "2000",
                "arriveTime": "2300",
                "availSubClass": [
                    {
                        "subClass": "A",
                        "seatAvailable": 0,
                        "seatClass": "E",
                        "adultPrice": 200000,
                        "childPrice": 0,
                        "infantPrice": 0
                    },
                    {
                        "subClass": "B",
                        "seatAvailable": 0,
                        "seatClass": "B",
                        "adultPrice": 150000,
                        "childPrice": 0,
                        "infantPrice": 0
                    },
                    {
                        "subClass": "C",
                        "seatAvailable": 0,
                        "seatClass": "K",
                        "adultPrice": 60000,
                        "childPrice": 0,
                        "infantPrice": 0
                    }
                ]
            }
        ]
    }
}
```