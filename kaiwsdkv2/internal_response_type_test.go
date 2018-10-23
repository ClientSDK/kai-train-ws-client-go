package kaiwsdkv2

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInternalGetOriginationRSOK is a positive test function for "InternalGetOriginationRS" internal response type <- mapping from "data.get_org"
func TestInternalGetOriginationRSOK(t *testing.T) {
	// fake schema
	str := `{"errCode": "0", "errMsg": null, "return": [{"originCode": "BD", "originName":"BANDUNG"},{"originCode": "YK", "originName":"YOGYAKARTA"}]}`

	// test expected values
	errCode := "0"
	orgLen := 2
	originCode := "BD"
	originName := "BANDUNG"

	// test variable
	var vRS InternalGetOriginationRS

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, orgLen, len(vRS.Return), "should be equal!")
	assert.Equal(t, originCode, vRS.Return[0].OriginCode, "should be equal")
	assert.Equal(t, originName, vRS.Return[0].OriginName, "should be equal")

	assert.Nil(t, err)
}

// TestInternalGetDestinationRSOK is a positive test function for "InternalGetDestinationRS" internal response type <- mapping from "data.get_des"
func TestInternalGetDestinationRSOK(t *testing.T) {
	// fake schema
	str := `{"errCode": "0", "errMsg": null, "return": [{"destCode": "BD", "destName":"BANDUNG"},{"destCode": "YK", "destName":"YOGYAKARTA"}]}`

	// test expected values
	errCode := "0"
	orgLen := 2
	destCode := "YK"
	destName := "YOGYAKARTA"

	// test variable
	var vRS InternalGetDestinationRS

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, orgLen, len(vRS.Return), "should be equal!")
	assert.Equal(t, destCode, vRS.Return[1].DestCode, "should be equal")
	assert.Equal(t, destName, vRS.Return[1].DestName, "should be equal")

	assert.Nil(t, err)
}

// TestInternalGetPayTypeRSOK is a positive test function for "InternalGetPayTypeRS" internal response type <- mapping from "data.get_pay_type"
func TestInternalGetPayTypeRSOK(t *testing.T) {
	// fake schema
	str := `{"errCode": "0", "errMsg": null, "return": [{"name": "TUNAI"},{"name": "EDC BNI"}]}`

	// test expected values
	errCode := "0"
	orgLen := 2
	name := "TUNAI"

	// test variable
	var vRS InternalGetPayTypeRS

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, orgLen, len(vRS.Return), "should be equal!")
	assert.Equal(t, name, vRS.Return[0].Name, "should be equal")

	assert.Nil(t, err)
}
