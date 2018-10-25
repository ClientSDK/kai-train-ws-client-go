package kaiwsdkv2

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMakeInternalGetOriginationRSOK is a positive test function for "MakeInternalGetOriginationRS" mapper method
func TestMakeInternalGetOriginationRSOK(t *testing.T) {
	// fake response
	str := `{ "err_code": 0,  "origination": [["BD","BANDUNG"],["YK","YOGYAKARTA"]] }`

	// transform to standard response
	stdStr := TrasformStandardKAIResponse([]byte(str))

	// test expected values
	errCode := "0"
	orgLen := 2
	originCode := "BD"
	originName := "BANDUNG"

	// test variable
	var vRS GetOriginationRS

	// test function
	err := json.Unmarshal(stdStr, &vRS)

	result, err := MakeInternalGetOriginationRS(vRS)

	// r, _ := json.Marshal(result)
	// fmt.Println(string(r))

	// test logic
	assert.Equal(t, errCode, result.ErrCode, "sould be equal!")

	assert.Equal(t, orgLen, len(result.Return), "sould be equal!")
	assert.Equal(t, originCode, result.Return[0].OriginCode, "sould be equal!")
	assert.Equal(t, originName, result.Return[0].OriginName, "sould be equal!")

	assert.Nil(t, err)
}

// TestMakeInternalGetDestinationRSOK is a positive test function for "MakeInternalGetOriginationRS" mapper method
func TestMakeInternalGetDestinationRSOK(t *testing.T) {
	// fake response
	str := `{ "err_code": 0,  "destination": [["BD","BANDUNG"],["YK","YOGYAKARTA"]] }`

	// transform to standard response
	stdStr := TrasformStandardKAIResponse([]byte(str))

	// test expected values
	errCode := "0"
	destLen := 2
	destCode := "YK"
	destName := "YOGYAKARTA"

	// test variable
	var vRS GetDestinationRS

	// test function
	err := json.Unmarshal(stdStr, &vRS)

	result, err := MakeInternalGetDestinationRS(vRS)

	// r, _ := json.Marshal(result)
	// fmt.Println(string(r))

	// test logic
	assert.Equal(t, errCode, result.ErrCode, "sould be equal!")

	assert.Equal(t, destLen, len(result.Return), "sould be equal!")
	assert.Equal(t, destCode, result.Return[1].DestCode, "sould be equal!")
	assert.Equal(t, destName, result.Return[1].DestName, "sould be equal!")

	assert.Nil(t, err)
}
