package kaiwsdkv2

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetDestinationRSOK is a positive test function for "GetDestinationRS" response type -> "data.get_des"
func TestGetDestinationRSOK(t *testing.T) {

	// fake response
	str := `{ "err_code": 0,  "destination": [["AK","ANGKE"],["AKB","AEKLOBA"]]}`

	// test expected values
	var errCode float64 // = 0
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
	var errCode float64 // = 0
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

// TestGetScheduleRSOK is a positive test function for "GetScheduleRS" response type -> "information.get_schedule"
func TestGetScheduleRSOK(t *testing.T) {
	// fake response
	str := `{ "err_code": 0, "org": "BD", "des": "GMR", "dep_date": "20190919", "schedule": [["10501","ARGO PARAHYANGAN PREMIUM","0415","0725",[["C",800,"K",100000,0,0]]],["710","RANGKAS JAYA","0800","1115",[["C",424,"K",80000,0,0]]],["77A","ARGO GOPAR","1200","1500",[["A",49,"E",100000,0,0],["B",50,"B",90000,0,0],["C",49,"K",60000,0,0]]],["P05","ARGO PARAHYANGAN","2000","2300",[["A",0,"E",200000,0,0],["B",0,"B",150000,0,0],["C",240,"K",60000,0,0]]]] }`

	// test variable
	var vRS GetScheduleRS
	org := "BD"
	des := "GMR"
	lenSchedule := 4

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	// test logic
	assert.Equal(t, org, vRS.Origin, "should be equal!")
	assert.Equal(t, des, vRS.Destination, "should be equal!")
	assert.Equal(t, lenSchedule, len(vRS.Schedules), "should be equal!")
	assert.Nil(t, err)
}

// TestGetScheduleV2RSOK is a positive test function for "GetScheduleV2RS" response type -> "information.get_schedule_v2"
func TestGetScheduleV2RSOK(t *testing.T) {
	// fake response
	str := `{ "err_code": 0, "org": "BD", "des": "GMR", "dep_date": "20190919", "schedule": [["10501","ARGO PARAHYANGAN PREMIUM","20190919","20190919","0415","0725",[["C",800,"K",100000,0,0]]],["710","RANGKAS JAYA","20190919","20190919","0800","1115",[["C",424,"K",80000,0,0]]],["77A","ARGO GOPAR","20190919","20190919","1200","1500",[["A",49,"E",100000,0,0],["B",50,"B",90000,0,0],["C",49,"K",60000,0,0]]],["P05","ARGO PARAHYANGAN","20190919","20190919","2000","2300",[["A",0,"E",200000,0,0],["B",0,"B",150000,0,0],["C",240,"K",60000,0,0]]]] }`

	// test variable
	var vRS GetScheduleV2RS
	org := "BD"
	des := "GMR"
	lenSchedule := 4
	arriveDateTrip1 := "20190919"
	arriveTimeTrip1 := "0725"

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	tripInfo := reflect.ValueOf(vRS.Schedules[0])

	// test logic
	assert.Equal(t, org, vRS.Origin, "should be equal!")
	assert.Equal(t, des, vRS.Destination, "should be equal!")
	assert.Equal(t, lenSchedule, len(vRS.Schedules), "should be equal!")
	assert.Equal(t, arriveDateTrip1, tripInfo.Index(3).Interface().(string), "should be equal!")
	assert.Equal(t, arriveTimeTrip1, tripInfo.Index(5).Interface().(string), "should be equal!")
	assert.Nil(t, err)
}

// TestGetScheduleLiteRSOK is a positive test function for "GetScheduleLiteRS" response type -> "information.get_schedule_lite"
func TestGetScheduleLiteRSOK(t *testing.T) {
	// fake response
	str := `{ "err_code": 0, "org": "BD", "des": "GMR", "dep_date": "20190919", "schedule": [["10501","ARGO PARAHYANGAN PREMIUM","20190919","20190919","0415","0725",[["C",0,"K",100000,0,0]]],["710","RANGKAS JAYA","20190919","20190919","0800","1115",[["C",0,"K",80000,0,0]]],["77A","ARGO GOPAR","20190919","20190919","1200","1500",[["A",0,"E",100000,0,0],["B",0,"B",90000,0,0],["C",0,"K",60000,0,0]]],["P05","ARGO PARAHYANGAN","20190919","20190919","2000","2300",[["A",0,"E",200000,0,0],["B",0,"B",150000,0,0],["C",0,"K",60000,0,0]]]] }`

	// test variable
	var vRS GetScheduleLiteRS
	org := "BD"
	des := "GMR"
	lenSchedule := 4
	arriveDateTrip1 := "20190919"
	arriveTimeTrip1 := "0725"
	var seatAvailTrip1Class1 float64 // = 0

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	tripInfo := reflect.ValueOf(vRS.Schedules[0])
	subClassArr := tripInfo.Index(6).Interface().([]interface{})
	subClass := reflect.ValueOf(subClassArr[0])

	// test logic
	assert.Equal(t, org, vRS.Origin, "should be equal!")
	assert.Equal(t, des, vRS.Destination, "should be equal!")
	assert.Equal(t, lenSchedule, len(vRS.Schedules), "should be equal!")
	assert.Equal(t, arriveDateTrip1, tripInfo.Index(3).Interface().(string), "should be equal!")
	assert.Equal(t, arriveTimeTrip1, tripInfo.Index(5).Interface().(string), "should be equal!")
	assert.Equal(t, seatAvailTrip1Class1, subClass.Index(1).Interface().(float64), "should be equal!")
	assert.Nil(t, err)
}
