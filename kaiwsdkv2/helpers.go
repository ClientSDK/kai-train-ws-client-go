// Copyright 2018 The ClientSDK Team Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0-style
// license that can be found in the LICENSE file.

// Author: ClientSDK Team (muharihar)

// Helpers
package kaiwsdkv2

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// StdKAIErrorMessage type
type StdKAIErrorMessage struct {
	ErrCode string `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

// TrasformStandardKAIResponse is a function to trasform KAI response to Internal KAI Standard Payload
func TrasformStandardKAIResponse(kaiResponse []byte) []byte {
	result := kaiResponse

	var vRS StdKAIErrorMessage

	err := json.Unmarshal(kaiResponse, &vRS)

	if err != nil {
		ik := strings.Index(string(kaiResponse), ",")
		ic := strings.Index(string(kaiResponse), ":")
		f := string(kaiResponse)[ic+1 : ik]
		t := fmt.Sprintf("\"%s\"", strings.TrimSpace(f))
		errStr := strings.Replace(string(kaiResponse)[:ik], f, t, 1)
		result = []byte(strings.Replace(string(kaiResponse), string(kaiResponse)[:ik], errStr, 1))

		err2 := json.Unmarshal(result, &vRS)
		if err2 != nil {
			log.Fatal(err2)
		}
	}

	return result
}

// DebugHTTP is function to debug HTTP Request/Response
func DebugHTTP(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	} else {
		log.Fatalf("%s\n\n", err)
	}
}

// DebugHTTPRequest is a function to debug HTTP Request
func DebugHTTPRequest(data []byte, err error) {
	fmt.Println("[DEBUG]:: DebugHTTPRequest")
	fmt.Println("=============================")
	fmt.Println("Request: ")
	fmt.Println("-----------------------------")
	DebugHTTP(data, err)
	fmt.Println("-----------------------------")
	fmt.Println("End-Request: ")
}

// DebugHTTPResponse is a function to debug HTTP Response
func DebugHTTPResponse(data []byte, err error) {
	fmt.Println("[DEBUG]:: DebugHTTPResponse")
	fmt.Println("=============================")
	fmt.Println("Response: ")
	fmt.Println("-----------------------------")
	DebugHTTP(data, err)
	fmt.Println("-----------------------------")
	fmt.Println("End-Response: ")
}
