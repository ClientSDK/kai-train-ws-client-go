package kaiwsdkv2

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrasformStandardKAIResponseOK(t *testing.T) {
	// kai response
	str := `{ "err_code": 0,  "destination": [["AK","ANGKE"],["AKB","AEKLOBA"]]}`
	//str := `{ "err_code": "002001", "err_msg": "Definition error, no application defined" }`
	//str := `{ "err_code": 002000, "err_msg": "Invalid Request. IP Address or Requester ID are not registered" }`
	//str := `{ "err_code": 002001, "err_msg": "Definition error, no application defined" }`

	// test variable
	var vRS StdKAIErrorMessage

	// test function
	stdPayload := TrasformStandardKAIResponse([]byte(str))
	err := json.Unmarshal([]byte(stdPayload), &vRS)

	// test logic
	assert.NotEqual(t, "", vRS.ErrCode, "should be not equal!")
	assert.Nil(t, err)
}
