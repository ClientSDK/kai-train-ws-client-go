package kaiwsdkv2

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetDestinationRSOK is a positive test function for "GetDestinationRS" response type -> "data.get_des"
func TestGetDestinationRSOK(t *testing.T) {

	// fake response
	str := `{ "err_code": 0,  "destination": [["AK","ANGKE"],["AKB","AEKLOBA"]]}`

	// test expected values
	var errCode float64 = 0
	desLen := 2
	desItemLen := 2

	// test variable
	var vRS GetDestinationRS

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, desLen, len(vRS.Destinations), "should be equal!")
	assert.Equal(t, desItemLen, len(vRS.Destinations[0]), "should be equal")
	assert.Nil(t, err)
}

// TestGetOriginationRSOK is a positive test function for "GetOriginationRS" response type -> "data.get_org"
func TestGetOriginationRSOK(t *testing.T) {

	// fake response
	str := `{ "err_code": 0,  "origination": [["AK","ANGKE"],["AKB","AEKLOBA"],["AWN","ARJAWINANGUN"]]}`

	// test expected values
	var errCode float64 = 0
	orgLen := 3
	orgItemLen := 2

	// test variable
	var vRS GetOriginationRS

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, orgLen, len(vRS.Originations), "should be equal!")
	assert.Equal(t, orgItemLen, len(vRS.Originations[0]), "should be equal")
	assert.Nil(t, err)
}

// TestGetPayTypeRSOK is a positive test function for "GetPayTypeRS" response type -> "data.get_pay_type"
func TestGetPayTypeRSOK(t *testing.T) {

	// fake response
	str := `{ "err_code": 0,  "pay_type": ["TUNAI","000009","EDC BNI"]}`

	// test expected values
	var errCode float64 = 0
	payTypeLen := 3

	// test variable
	var vRS GetPayTypeRS

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, payTypeLen, len(vRS.PayTypes), "should be equal!")
	assert.Nil(t, err)
}
