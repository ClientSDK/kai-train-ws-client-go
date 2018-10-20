package kaiwsdkv2

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetDestinationRSOK is a positive test function for "GetDestinationRS" response type -> "data.get_des"
func TestGetDestinationRSOK(t *testing.T) {

	str := `{ "err_code": 0,  "destination": [["AK","ANGKE"],["AKB","AEKLOBA"]]}`

	var vRS GetDestinationRS

	err := json.Unmarshal([]byte(str), &vRS)

	assert.Nil(t, err)
}

// TestGetOriginationRSOK is a positive test function for "GetOriginationRS" response type -> "data.get_org"
func TestGetOriginationRSOK(t *testing.T) {

	str := `{ "err_code": 0,  "origination": [["AK","ANGKE"],["AKB","AEKLOBA"],["AWN","ARJAWINANGUN"]]}`

	var vRS GetOriginationRS

	err := json.Unmarshal([]byte(str), &vRS)

	assert.Nil(t, err)
}

func TestGetPayTypeRSOK(t *testing.T) {

	str := `{ "err_code": 0,  "pay_type": ["TUNAI","000009","EDC BNI"]}`

	var vRS GetPayTypeRS

	err := json.Unmarshal([]byte(str), &vRS)

	assert.Nil(t, err)
}
