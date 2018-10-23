package kaiwsdkv2

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInternalGetOriginationRSOK is a positive test function for "InternalGetOriginationRS" internal response type <- mapping from "data.get_org"
func TestInternalGetOriginationRSOK(t *testing.T) {
	// fake schema
	str := `{"errCode": "0", "errMsg": null, "return": [{"originCode": "AK", "originName":"ANGKE"},{"originCode": "YK", "originName":"YOGYAKARTA"}]}`

	// test expected values
	errCode := "0"
	orgLen := 2
	originCode := "YK"
	originName := "YOGYAKARTA"

	// test variable
	var vRS InternalGetOriginationRS

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, orgLen, len(vRS.Return), "should be equal!")
	assert.Equal(t, originCode, vRS.Return[1].OriginCode, "should be equal")
	assert.Equal(t, originName, vRS.Return[1].OriginName, "should be equal")

	assert.Nil(t, err)
}
