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

// TestGetSeatMapRSOK is a positive test function for "GetSeatMapRS" response type -> "information.get_seat_map"
func TestGetSeatMapRSOK(t *testing.T) {
	// fake response
	str := `{ "err_code": 0, "org": "BD", "des": "GMR", "train_no": "10501", "dep_date": "20190919", "seat_map": [["PREMIUM",1,[[1,1,1,"A","C",0],[1,2,1,"B","C",0],[1,3,1,"","",0],[1,4,1,"C","C",0],[1,5,1,"D","C",0],[1,6,1,"E","C",0],[2,1,2,"A","C",0],[2,2,2,"B","C",0],[2,3,2,"","",0],[2,4,2,"C","C",0],[2,5,2,"D","C",0],[2,6,2,"E","C",0],[3,1,3,"A","C",0],[3,2,3,"B","C",0],[3,3,3,"","",0],[3,4,3,"C","C",0],[3,5,3,"D","C",0],[3,6,3,"E","C",0],[4,1,4,"A","C",0],[4,2,4,"B","C",0],[4,3,4,"","",0],[4,4,4,"C","C",0],[4,5,4,"D","C",0],[4,6,4,"E","C",0],[5,1,5,"A","C",0],[5,2,5,"B","C",0],[5,3,5,"","",0],[5,4,5,"C","C",0],[5,5,5,"D","C",0],[5,6,5,"E","C",0],[6,1,6,"A","C",0],[6,2,6,"B","C",0],[6,3,6,"","",0],[6,4,6,"C","C",0],[6,5,6,"D","C",0],[6,6,6,"E","C",0],[7,1,7,"A","C",0],[7,2,7,"B","C",0],[7,3,7,"","",0],[7,4,7,"C","C",0],[7,5,7,"D","C",0],[7,6,7,"E","C",0],[8,1,8,"A","C",0],[8,2,8,"B","C",0],[8,3,8,"","",0],[8,4,8,"C","C",0],[8,5,8,"D","C",0],[8,6,8,"E","C",0],[9,1,9,"A","C",0],[9,2,9,"B","C",0],[9,3,9,"","",0],[9,4,9,"C","C",0],[9,5,9,"D","C",0],[9,6,9,"E","C",0],[10,1,10,"A","C",0],[10,2,10,"B","C",0],[10,3,10,"","",0],[10,4,10,"C","C",0],[10,5,10,"D","C",0],[10,6,10,"E","C",0],[11,1,11,"A","C",0],[11,2,11,"B","C",0],[11,3,11,"","",0],[11,4,11,"C","C",0],[11,5,11,"D","C",0],[11,6,11,"E","C",0],[12,1,12,"A","C",0],[12,2,12,"B","C",0],[12,3,12,"","",0],[12,4,12,"C","C",0],[12,5,12,"D","C",0],[12,6,12,"E","C",0],[13,1,13,"A","C",0],[13,2,13,"B","C",0],[13,3,13,"","",0],[13,4,13,"C","C",0],[13,5,13,"D","C",0],[13,6,13,"E","C",0],[14,1,14,"A","C",0],[14,2,14,"B","C",0],[14,3,14,"","",0],[14,4,14,"C","C",0],[14,5,14,"D","C",0],[14,6,14,"E","C",0],[15,1,15,"A","C",0],[15,2,15,"B","C",0],[15,3,15,"","",0],[15,4,15,"C","C",0],[15,5,15,"D","C",0],[15,6,15,"E","C",0],[16,1,16,"A","C",0],[16,2,16,"B","C",0],[16,3,16,"","",0],[16,4,16,"C","C",0],[16,5,16,"D","C",0],[16,6,16,"E","C",0],[17,1,17,"A","C",0],[17,2,17,"B","C",0],[17,3,17,"","",0],[17,4,17,"C","C",0],[17,5,17,"D","C",0],[17,6,17,"E","C",0]]],["PREMIUM",2,[[1,1,1,"A","C",0],[1,2,1,"B","C",0],[1,3,1,"","",0],[1,4,1,"C","C",0],[1,5,1,"D","C",0],[1,6,1,"E","C",0],[2,1,2,"A","C",0],[2,2,2,"B","C",0],[2,3,2,"","",0],[2,4,2,"C","C",0],[2,5,2,"D","C",0],[2,6,2,"E","C",0],[3,1,3,"A","C",0],[3,2,3,"B","C",0],[3,3,3,"","",0],[3,4,3,"C","C",0],[3,5,3,"D","C",0],[3,6,3,"E","C",0],[4,1,4,"A","C",0],[4,2,4,"B","C",0],[4,3,4,"","",0],[4,4,4,"C","C",0],[4,5,4,"D","C",0],[4,6,4,"E","C",0],[5,1,5,"A","C",0],[5,2,5,"B","C",0],[5,3,5,"","",0],[5,4,5,"C","C",0],[5,5,5,"D","C",0],[5,6,5,"E","C",0],[6,1,6,"A","C",0],[6,2,6,"B","C",0],[6,3,6,"","",0],[6,4,6,"C","C",0],[6,5,6,"D","C",0],[6,6,6,"E","C",0],[7,1,7,"A","C",0],[7,2,7,"B","C",0],[7,3,7,"","",0],[7,4,7,"C","C",0],[7,5,7,"D","C",0],[7,6,7,"E","C",0],[8,1,8,"A","C",0],[8,2,8,"B","C",0],[8,3,8,"","",0],[8,4,8,"C","C",0],[8,5,8,"D","C",0],[8,6,8,"E","C",0],[9,1,9,"A","C",0],[9,2,9,"B","C",0],[9,3,9,"","",0],[9,4,9,"C","C",0],[9,5,9,"D","C",0],[9,6,9,"E","C",0],[10,1,10,"A","C",0],[10,2,10,"B","C",0],[10,3,10,"","",0],[10,4,10,"C","C",0],[10,5,10,"D","C",0],[10,6,10,"E","C",0],[11,1,11,"A","C",0],[11,2,11,"B","C",0],[11,3,11,"","",0],[11,4,11,"C","C",0],[11,5,11,"D","C",0],[11,6,11,"E","C",0],[12,1,12,"A","C",0],[12,2,12,"B","C",0],[12,3,12,"","",0],[12,4,12,"C","C",0],[12,5,12,"D","C",0],[12,6,12,"E","C",0],[13,1,13,"A","C",0],[13,2,13,"B","C",0],[13,3,13,"","",0],[13,4,13,"C","C",0],[13,5,13,"D","C",0],[13,6,13,"E","C",0],[14,1,14,"A","C",0],[14,2,14,"B","C",0],[14,3,14,"","",0],[14,4,14,"C","C",0],[14,5,14,"D","C",0],[14,6,14,"E","C",0],[15,1,15,"A","C",0],[15,2,15,"B","C",0],[15,3,15,"","",0],[15,4,15,"C","C",0],[15,5,15,"D","C",0],[15,6,15,"E","C",0],[16,1,16,"A","C",0],[16,2,16,"B","C",0],[16,3,16,"","",0],[16,4,16,"C","C",0],[16,5,16,"D","C",0],[16,6,16,"E","C",0],[17,1,17,"A","C",0],[17,2,17,"B","C",0],[17,3,17,"","",0],[17,4,17,"C","C",0],[17,5,17,"D","C",0],[17,6,17,"E","C",0]]]]}`

	// test variable
	var vRS GetSeatMapRS
	org := "BD"
	des := "GMR"
	trainNo := "10501"
	lenWagon := 2
	wagonCode := "PREMIUM"
	var wagonNo float64 = 1
	var row float64 = 1
	var column float64 = 1
	var seatRow float64 = 1
	seatColumn := "A"
	subClass := "C"
	var status float64 // = 0

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	wagonInfo0 := reflect.ValueOf(vRS.SeatMaps[0])
	seatInfo := wagonInfo0.Index(2).Interface().([]interface{})
	seatInfo0 := reflect.ValueOf(seatInfo[0])

	// test logic
	assert.Equal(t, org, vRS.Origin, "should be equal!")
	assert.Equal(t, des, vRS.Destination, "should be equal!")
	assert.Equal(t, trainNo, vRS.TrainNo, "should be equal!")
	assert.Equal(t, lenWagon, len(vRS.SeatMaps), "should be equal!")

	assert.Equal(t, wagonCode, wagonInfo0.Index(0).Interface().(string), "should be equal!")
	assert.Equal(t, wagonNo, wagonInfo0.Index(1).Interface().(float64), "should be equal!")

	assert.Equal(t, row, seatInfo0.Index(0).Interface().(float64), "should be equal!")
	assert.Equal(t, column, seatInfo0.Index(1).Interface().(float64), "should be equal!")
	assert.Equal(t, seatRow, seatInfo0.Index(2).Interface().(float64), "should be equal!")
	assert.Equal(t, seatColumn, seatInfo0.Index(3).Interface().(string), "should be equal!")
	assert.Equal(t, subClass, seatInfo0.Index(4).Interface().(string), "should be equal!")
	assert.Equal(t, status, seatInfo0.Index(5).Interface().(float64), "should be equal!")

	assert.Nil(t, err)
}

// TestGetSeatMapPerSubClassRSOK is a positive test function for "GetSeatMapPerSubClassRS" response type -> "information.get_seat_map_per_subclass"
func TestGetSeatMapPerSubClassRSOK(t *testing.T) {
	// fake response
	str := `{ "err_code": 0, "org": "BD", "des": "GMR", "train_no": "10501", "dep_date": "20190919", "seat_map": [["PREMIUM",1,[[1,1,1,"A","C",0],[1,2,1,"B","C",0],[1,4,1,"C","C",0],[1,5,1,"D","C",0],[1,6,1,"E","C",0],[2,1,2,"A","C",0],[2,2,2,"B","C",0],[2,4,2,"C","C",0],[2,5,2,"D","C",0],[2,6,2,"E","C",0],[3,1,3,"A","C",0],[3,2,3,"B","C",0],[3,4,3,"C","C",0],[3,5,3,"D","C",0],[3,6,3,"E","C",0],[4,1,4,"A","C",0],[4,2,4,"B","C",0],[4,4,4,"C","C",0],[4,5,4,"D","C",0],[4,6,4,"E","C",0],[5,1,5,"A","C",0],[5,2,5,"B","C",0],[5,4,5,"C","C",0],[5,5,5,"D","C",0],[5,6,5,"E","C",0],[6,1,6,"A","C",0],[6,2,6,"B","C",0],[6,4,6,"C","C",0],[6,5,6,"D","C",0],[6,6,6,"E","C",0],[7,1,7,"A","C",0],[7,2,7,"B","C",0],[7,4,7,"C","C",0],[7,5,7,"D","C",0],[7,6,7,"E","C",0],[8,1,8,"A","C",0],[8,2,8,"B","C",0],[8,4,8,"C","C",0],[8,5,8,"D","C",0],[8,6,8,"E","C",0],[9,1,9,"A","C",0],[9,2,9,"B","C",0],[9,4,9,"C","C",0],[9,5,9,"D","C",0],[9,6,9,"E","C",0],[10,1,10,"A","C",0],[10,2,10,"B","C",0],[10,4,10,"C","C",0],[10,5,10,"D","C",0],[10,6,10,"E","C",0],[11,1,11,"A","C",0],[11,2,11,"B","C",0],[11,4,11,"C","C",0],[11,5,11,"D","C",0],[11,6,11,"E","C",0],[12,1,12,"A","C",0],[12,2,12,"B","C",0],[12,4,12,"C","C",0],[12,5,12,"D","C",0],[12,6,12,"E","C",0],[13,1,13,"A","C",0],[13,2,13,"B","C",0],[13,4,13,"C","C",0],[13,5,13,"D","C",0],[13,6,13,"E","C",0],[14,1,14,"A","C",0],[14,2,14,"B","C",0],[14,4,14,"C","C",0],[14,5,14,"D","C",0],[14,6,14,"E","C",0],[15,1,15,"A","C",0],[15,2,15,"B","C",0],[15,4,15,"C","C",0],[15,5,15,"D","C",0],[15,6,15,"E","C",0],[16,1,16,"A","C",0],[16,2,16,"B","C",0],[16,4,16,"C","C",0],[16,5,16,"D","C",0],[16,6,16,"E","C",0],[17,1,17,"A","C",0],[17,2,17,"B","C",0],[17,4,17,"C","C",0],[17,5,17,"D","C",0],[17,6,17,"E","C",0]]],["PREMIUM",2,[[1,1,1,"A","C",0],[1,2,1,"B","C",0],[1,4,1,"C","C",0],[1,5,1,"D","C",0],[1,6,1,"E","C",0],[2,1,2,"A","C",0],[2,2,2,"B","C",0],[2,4,2,"C","C",0],[2,5,2,"D","C",0],[2,6,2,"E","C",0],[3,1,3,"A","C",0],[3,2,3,"B","C",0],[3,4,3,"C","C",0],[3,5,3,"D","C",0],[3,6,3,"E","C",0],[4,1,4,"A","C",0],[4,2,4,"B","C",0],[4,4,4,"C","C",0],[4,5,4,"D","C",0],[4,6,4,"E","C",0],[5,1,5,"A","C",0],[5,2,5,"B","C",0],[5,4,5,"C","C",0],[5,5,5,"D","C",0],[5,6,5,"E","C",0],[6,1,6,"A","C",0],[6,2,6,"B","C",0],[6,4,6,"C","C",0],[6,5,6,"D","C",0],[6,6,6,"E","C",0],[7,1,7,"A","C",0],[7,2,7,"B","C",0],[7,4,7,"C","C",0],[7,5,7,"D","C",0],[7,6,7,"E","C",0],[8,1,8,"A","C",0],[8,2,8,"B","C",0],[8,4,8,"C","C",0],[8,5,8,"D","C",0],[8,6,8,"E","C",0],[9,1,9,"A","C",0],[9,2,9,"B","C",0],[9,4,9,"C","C",0],[9,5,9,"D","C",0],[9,6,9,"E","C",0],[10,1,10,"A","C",0],[10,2,10,"B","C",0],[10,4,10,"C","C",0],[10,5,10,"D","C",0],[10,6,10,"E","C",0],[11,1,11,"A","C",0],[11,2,11,"B","C",0],[11,4,11,"C","C",0],[11,5,11,"D","C",0],[11,6,11,"E","C",0],[12,1,12,"A","C",0],[12,2,12,"B","C",0],[12,4,12,"C","C",0],[12,5,12,"D","C",0],[12,6,12,"E","C",0],[13,1,13,"A","C",0],[13,2,13,"B","C",0],[13,4,13,"C","C",0],[13,5,13,"D","C",0],[13,6,13,"E","C",0],[14,1,14,"A","C",0],[14,2,14,"B","C",0],[14,4,14,"C","C",0],[14,5,14,"D","C",0],[14,6,14,"E","C",0],[15,1,15,"A","C",0],[15,2,15,"B","C",0],[15,4,15,"C","C",0],[15,5,15,"D","C",0],[15,6,15,"E","C",0],[16,1,16,"A","C",0],[16,2,16,"B","C",0],[16,4,16,"C","C",0],[16,5,16,"D","C",0],[16,6,16,"E","C",0],[17,1,17,"A","C",0],[17,2,17,"B","C",0],[17,4,17,"C","C",0],[17,5,17,"D","C",0],[17,6,17,"E","C",0]]]] }`

	// test variable
	var vRS GetSeatMapPerSubClassRS
	org := "BD"
	des := "GMR"
	trainNo := "10501"
	lenWagon := 2
	wagonCode := "PREMIUM"
	var wagonNo float64 = 1
	var row float64 = 1
	var column float64 = 1
	var seatRow float64 = 1
	seatColumn := "A"
	subClass := "C"
	var status float64 // = 0

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	wagonInfo0 := reflect.ValueOf(vRS.SeatMaps[0])
	seatInfo := wagonInfo0.Index(2).Interface().([]interface{})
	seatInfo0 := reflect.ValueOf(seatInfo[0])

	// test logic
	assert.Equal(t, org, vRS.Origin, "should be equal!")
	assert.Equal(t, des, vRS.Destination, "should be equal!")
	assert.Equal(t, trainNo, vRS.TrainNo, "should be equal!")
	assert.Equal(t, lenWagon, len(vRS.SeatMaps), "should be equal!")

	assert.Equal(t, wagonCode, wagonInfo0.Index(0).Interface().(string), "should be equal!")
	assert.Equal(t, wagonNo, wagonInfo0.Index(1).Interface().(float64), "should be equal!")

	assert.Equal(t, row, seatInfo0.Index(0).Interface().(float64), "should be equal!")
	assert.Equal(t, column, seatInfo0.Index(1).Interface().(float64), "should be equal!")
	assert.Equal(t, seatRow, seatInfo0.Index(2).Interface().(float64), "should be equal!")
	assert.Equal(t, seatColumn, seatInfo0.Index(3).Interface().(string), "should be equal!")
	assert.Equal(t, subClass, seatInfo0.Index(4).Interface().(string), "should be equal!")
	assert.Equal(t, status, seatInfo0.Index(5).Interface().(float64), "should be equal!")

	assert.Nil(t, err)
}

// TestGetSeatNullRSOK is a positive test function for "GetSeatNullRS" response type -> "information.get_seat_null"
func TestGetSeatNullRSOK(t *testing.T) {
	// fake response
	str := `{ "err_code": 0, "org": "BD", "des": "GMR", "train_no": "10501", "dep_date": "20181101", "seat_null": [["PREMIUM",1,[[1,"A","C",0],[1,"B","C",0],[1,"C","C",0],[1,"D","C",0],[1,"E","C",0],[2,"A","C",0],[2,"B","C",0],[2,"C","C",0],[2,"D","C",0],[2,"E","C",0],[3,"A","C",0],[3,"B","C",0],[3,"C","C",0],[3,"D","C",0],[3,"E","C",0],[4,"A","C",0],[4,"B","C",0],[4,"C","C",0],[4,"D","C",0],[4,"E","C",0],[5,"A","C",0],[5,"B","C",0],[5,"C","C",0],[5,"D","C",0],[5,"E","C",0],[6,"A","C",0],[6,"B","C",0],[6,"C","C",0],[6,"D","C",0],[6,"E","C",0],[7,"A","C",0],[7,"B","C",0],[7,"C","C",0],[7,"D","C",0],[7,"E","C",0],[8,"A","C",0],[8,"B","C",0],[8,"C","C",0],[8,"D","C",0],[8,"E","C",0],[9,"A","C",0],[9,"B","C",0],[9,"C","C",0],[9,"D","C",0],[9,"E","C",0],[10,"A","C",0],[10,"B","C",0],[10,"C","C",0],[10,"D","C",0],[10,"E","C",0],[11,"A","C",0],[11,"B","C",0],[11,"C","C",0],[11,"D","C",0],[11,"E","C",0],[12,"A","C",0],[12,"B","C",0],[12,"C","C",0],[12,"D","C",0],[12,"E","C",0],[13,"A","C",0],[13,"B","C",0],[13,"C","C",0],[13,"D","C",0],[13,"E","C",0],[14,"A","C",0],[14,"B","C",0],[14,"C","C",0],[14,"D","C",0],[14,"E","C",0],[15,"A","C",0],[15,"B","C",0],[15,"C","C",0],[15,"D","C",0],[15,"E","C",0],[16,"A","C",0],[16,"B","C",0],[16,"C","C",0],[16,"D","C",0],[16,"E","C",0],[17,"A","C",0],[17,"B","C",0],[17,"C","C",0],[17,"D","C",0],[17,"E","C",0]]],["PREMIUM",2,[[1,"A","C",0],[1,"B","C",0],[1,"C","C",0],[1,"D","C",0],[1,"E","C",0],[2,"A","C",0],[2,"B","C",0],[2,"C","C",0],[2,"D","C",0],[2,"E","C",0],[3,"A","C",0],[3,"B","C",0],[3,"C","C",0],[3,"D","C",0],[3,"E","C",0],[4,"A","C",0],[4,"B","C",0],[4,"C","C",0],[4,"D","C",0],[4,"E","C",0],[5,"A","C",0],[5,"B","C",0],[5,"C","C",0],[5,"D","C",0],[5,"E","C",0],[6,"A","C",0],[6,"B","C",0],[6,"C","C",0],[6,"D","C",0],[6,"E","C",0],[7,"A","C",0],[7,"B","C",0],[7,"C","C",0],[7,"D","C",0],[7,"E","C",0],[8,"A","C",0],[8,"B","C",0],[8,"C","C",0],[8,"D","C",0],[8,"E","C",0],[9,"A","C",0],[9,"B","C",0],[9,"C","C",0],[9,"D","C",0],[9,"E","C",0],[10,"A","C",0],[10,"B","C",0],[10,"C","C",0],[10,"D","C",0],[10,"E","C",0],[11,"A","C",0],[11,"B","C",0],[11,"C","C",0],[11,"D","C",0],[11,"E","C",0],[12,"A","C",0],[12,"B","C",0],[12,"C","C",0],[12,"D","C",0],[12,"E","C",0],[13,"A","C",0],[13,"B","C",0],[13,"C","C",0],[13,"D","C",0],[13,"E","C",0],[14,"A","C",0],[14,"B","C",0],[14,"C","C",0],[14,"D","C",0],[14,"E","C",0],[15,"A","C",0],[15,"B","C",0],[15,"C","C",0],[15,"D","C",0],[15,"E","C",0],[16,"A","C",0],[16,"B","C",0],[16,"C","C",0],[16,"D","C",0],[16,"E","C",0],[17,"A","C",0],[17,"B","C",0],[17,"C","C",0],[17,"D","C",0],[17,"E","C",0]]]]}`

	// test variable
	var vRS GetSeatNullRS
	org := "BD"
	des := "GMR"
	trainNo := "10501"
	lenWagon := 2
	wagonCode := "PREMIUM"
	var wagonNo float64 = 1
	var seatRow float64 = 1
	seatColumn := "A"
	subClass := "C"
	var status float64 // = 0

	// test function
	err := json.Unmarshal([]byte(str), &vRS)

	wagonInfo0 := reflect.ValueOf(vRS.SeatNulls[0])
	seatInfo := wagonInfo0.Index(2).Interface().([]interface{})
	seatInfo0 := reflect.ValueOf(seatInfo[0])

	// test logic
	assert.Equal(t, org, vRS.Origin, "should be equal!")
	assert.Equal(t, des, vRS.Destination, "should be equal!")
	assert.Equal(t, trainNo, vRS.TrainNo, "should be equal!")
	assert.Equal(t, lenWagon, len(vRS.SeatNulls), "should be equal!")

	assert.Equal(t, wagonCode, wagonInfo0.Index(0).Interface().(string), "should be equal!")
	assert.Equal(t, wagonNo, wagonInfo0.Index(1).Interface().(float64), "should be equal!")

	assert.Equal(t, seatRow, seatInfo0.Index(0).Interface().(float64), "should be equal!")
	assert.Equal(t, seatColumn, seatInfo0.Index(1).Interface().(string), "should be equal!")
	assert.Equal(t, subClass, seatInfo0.Index(2).Interface().(string), "should be equal!")
	assert.Equal(t, status, seatInfo0.Index(3).Interface().(float64), "should be equal!")

	assert.Nil(t, err)
}
