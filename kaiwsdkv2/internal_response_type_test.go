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
	desLen := 2
	destCode := "YK"
	destName := "YOGYAKARTA"

	// test variable
	var vRS InternalGetDestinationRS

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, desLen, len(vRS.Return), "should be equal!")
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
	payLen := 2
	name := "TUNAI"

	// test variable
	var vRS InternalGetPayTypeRS

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, payLen, len(vRS.Return), "should be equal!")
	assert.Equal(t, name, vRS.Return[0].Name, "should be equal")

	assert.Nil(t, err)
}

// TestInternalGetScheduleRSOK is a positive test function for "InternalGetScheduleRSOK" internal response type <- mapping from "information.get_schedule"
func TestInternalGetScheduleRSOK(t *testing.T) {
	// fake schema
	str := `{"errCode": "0","errMsg":null,"return":{"origin":"BD","destination":"GMR","departureDate":"20190919","schedules":[{"trainNo":"10501","trainName":"ARGO PARAHYANGAN PREMIUM","departureTime":"0415","arriveTime":"0725","availSubClass":[{"subClass":"C","seatAvailable":800,"seatClass":"K","adultPrice":100000,"childPrice":0,"infantPrice":0}]},{"trainNo":"710","trainName":"RANGKAS JAYA","departureTime":"0800","arriveTime":"1115","availSubClass":[{"subClass":"C","seatAvailable":424,"seatClass":"K","adultPrice":80000,"childPrice":0,"infantPrice":0}]}]}}`

	// test expected values
	errCode := "0"
	org := "BD"
	des := "GMR"
	schedLen := 2
	trainNo0 := "10501"
	trainName0 := "ARGO PARAHYANGAN PREMIUM"

	// test variable
	var vRS InternalGetScheduleRS

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, org, vRS.Return.Origin, "should be equal")
	assert.Equal(t, des, vRS.Return.Destination, "should be equal")

	assert.Equal(t, schedLen, len(vRS.Return.Schedules), "should be equal!")
	assert.Equal(t, trainNo0, vRS.Return.Schedules[0].TrainNo, "should be equal!")
	assert.Equal(t, trainName0, vRS.Return.Schedules[0].TrainName, "should be equal!")

	assert.Nil(t, err)
}

// TestInternalGetScheduleV2RSOK is a positive test function for "InternalGetScheduleV2RS" internal response type <- mapping from "information.get_schedule_v2"
func TestInternalGetScheduleV2RSOK(t *testing.T) {
	// fake schema
	str := `{"errCode":"0","errMsg":null,"return":{"origin":"BD","destination":"GMR","departureDate":"20190919","schedules":[{"trainNo":"10501","trainName":"ARGO PARAHYANGAN PREMIUM","departureDate":"20190919","arriveDate":"20190919","departureTime":"0415","arriveTime":"0725","availSubClass":[{"subClass":"C","seatAvailable":800,"seatClass":"K","adultPrice":100000,"childPrice":0,"infantPrice":0}]},{"trainNo":"710","trainName":"RANGKAS JAYA","departureDate":"20190919","arriveDate":"20190919","departureTime":"0800","arriveTime":"1115","availSubClass":[{"subClass":"C","seatAvailable":424,"seatClass":"K","adultPrice":80000,"childPrice":0,"infantPrice":0}]},{"trainNo":"77A","trainName":"ARGO GOPAR","departureDate":"20190919","arriveDate":"20190919","departureTime":"1200","arriveTime":"1500","availSubClass":[{"subClass":"A","seatAvailable":49,"seatClass":"E","adultPrice":100000,"childPrice":0,"infantPrice":0},{"subClass":"B","seatAvailable":50,"seatClass":"B","adultPrice":90000,"childPrice":0,"infantPrice":0},{"subClass":"C","seatAvailable":49,"seatClass":"K","adultPrice":60000,"childPrice":0,"infantPrice":0}]},{"trainNo":"P05","trainName":"ARGO PARAHYANGAN","departureDate":"20190919","arriveDate":"20190919","departureTime":"2000","arriveTime":"2300","availSubClass":[{"subClass":"A","seatAvailable":0,"seatClass":"E","adultPrice":200000,"childPrice":0,"infantPrice":0},{"subClass":"B","seatAvailable":0,"seatClass":"B","adultPrice":150000,"childPrice":0,"infantPrice":0},{"subClass":"C","seatAvailable":240,"seatClass":"K","adultPrice":60000,"childPrice":0,"infantPrice":0}]}]}}`

	// test expected values
	errCode := "0"
	org := "BD"
	des := "GMR"
	schedLen := 4
	trainNo0 := "10501"
	trainName0 := "ARGO PARAHYANGAN PREMIUM"
	arvDate0 := "20190919"
	arvTime0 := "0725"

	// test variable
	var vRS InternalGetScheduleV2RS

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, org, vRS.Return.Origin, "should be equal")
	assert.Equal(t, des, vRS.Return.Destination, "should be equal")

	assert.Equal(t, schedLen, len(vRS.Return.Schedules), "should be equal!")
	assert.Equal(t, trainNo0, vRS.Return.Schedules[0].TrainNo, "should be equal!")
	assert.Equal(t, trainName0, vRS.Return.Schedules[0].TrainName, "should be equal!")
	assert.Equal(t, arvDate0, vRS.Return.Schedules[0].ArriveDate, "should be equal!")
	assert.Equal(t, arvTime0, vRS.Return.Schedules[0].ArriveTime, "should be equal!")

	assert.Nil(t, err)
}

// TestInternalGetScheduleLiteRSOK is a positive test function for "InternalGetScheduleLiteRS" internal response type <- mapping from "information.get_schedule_lite"
func TestInternalGetScheduleLiteRSOK(t *testing.T) {
	// fake schema
	str := `{"errCode":"0","errMsg":null,"return":{"origin":"BD","destination":"GMR","departureDate":"20190919","schedules":[{"trainNo":"10501","trainName":"ARGO PARAHYANGAN PREMIUM","departureDate":"20190919","arriveDate":"20190919","departureTime":"0415","arriveTime":"0725","availSubClass":[{"subClass":"C","seatAvailable":0,"seatClass":"K","adultPrice":100000,"childPrice":0,"infantPrice":0}]},{"trainNo":"710","trainName":"RANGKAS JAYA","departureDate":"20190919","arriveDate":"20190919","departureTime":"0800","arriveTime":"1115","availSubClass":[{"subClass":"C","seatAvailable":0,"seatClass":"K","adultPrice":80000,"childPrice":0,"infantPrice":0}]},{"trainNo":"77A","trainName":"ARGO GOPAR","departureDate":"20190919","arriveDate":"20190919","departureTime":"1200","arriveTime":"1500","availSubClass":[{"subClass":"A","seatAvailable":0,"seatClass":"E","adultPrice":100000,"childPrice":0,"infantPrice":0},{"subClass":"B","seatAvailable":0,"seatClass":"B","adultPrice":90000,"childPrice":0,"infantPrice":0},{"subClass":"C","seatAvailable":0,"seatClass":"K","adultPrice":60000,"childPrice":0,"infantPrice":0}]},{"trainNo":"P05","trainName":"ARGO PARAHYANGAN","departureDate":"20190919","arriveDate":"20190919","departureTime":"2000","arriveTime":"2300","availSubClass":[{"subClass":"A","seatAvailable":0,"seatClass":"E","adultPrice":200000,"childPrice":0,"infantPrice":0},{"subClass":"B","seatAvailable":0,"seatClass":"B","adultPrice":150000,"childPrice":0,"infantPrice":0},{"subClass":"C","seatAvailable":0,"seatClass":"K","adultPrice":60000,"childPrice":0,"infantPrice":0}]}]}}`

	// test expected values
	errCode := "0"
	org := "BD"
	des := "GMR"
	schedLen := 4
	trainNo0 := "10501"
	trainName0 := "ARGO PARAHYANGAN PREMIUM"
	arvDate0 := "20190919"
	arvTime0 := "0725"
	var seatAvailable0_0 float64 // = 0

	// test variable
	var vRS InternalGetScheduleLiteRS

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	// test logic
	assert.Equal(t, errCode, vRS.ErrCode, "should be equal!")
	assert.Equal(t, org, vRS.Return.Origin, "should be equal")
	assert.Equal(t, des, vRS.Return.Destination, "should be equal")

	assert.Equal(t, schedLen, len(vRS.Return.Schedules), "should be equal!")
	assert.Equal(t, trainNo0, vRS.Return.Schedules[0].TrainNo, "should be equal!")
	assert.Equal(t, trainName0, vRS.Return.Schedules[0].TrainName, "should be equal!")
	assert.Equal(t, arvDate0, vRS.Return.Schedules[0].ArriveDate, "should be equal!")
	assert.Equal(t, arvTime0, vRS.Return.Schedules[0].ArriveTime, "should be equal!")
	assert.Equal(t, seatAvailable0_0, vRS.Return.Schedules[0].AvailSubClass[0].SeatAvailable, "should be equal!")

	assert.Nil(t, err)
}
