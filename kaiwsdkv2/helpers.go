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
	ErrMsg  string `jsin:"err_msg"`
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
