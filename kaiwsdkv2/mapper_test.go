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

// TestMakeInternalGetPayTypeRSOK is a positive test function for "MakeInternalGetPayTypeRS" mapper method
func TestMakeInternalGetPayTypeRSOK(t *testing.T) {
	// fake response
	str := `{ "err_code": 0,  "pay_type": ["TUNAI","000009","EDC BNI"]}`

	// transform to standard response
	stdStr := TrasformStandardKAIResponse([]byte(str))

	// test expected values
	errCode := "0"
	ptLen := 3
	ptName := "TUNAI"

	// test variable
	var vRS GetPayTypeRS

	// test function
	err := json.Unmarshal(stdStr, &vRS)

	result, err := MakeInternalGetPayTypeRS(vRS)

	// r, _ := json.Marshal(result)
	// fmt.Println(string(r))

	// test logic
	assert.Equal(t, errCode, result.ErrCode, "sould be equal!")

	assert.Equal(t, ptLen, len(result.Return), "sould be equal!")
	assert.Equal(t, ptName, result.Return[0].Name, "sould be equal!")

	assert.Nil(t, err)
}

// TestMakeInternalGetScheduleRSOK is a positive test function for "MakeInternalGetScheduleRS" mapper method
func TestMakeInternalGetScheduleRSOK(t *testing.T) {
	// fake response
	str := `{ "err_code": 0, "org": "BD", "des": "GMR", "dep_date": "20190919", "schedule": [["10501","ARGO PARAHYANGAN PREMIUM","0415","0725",[["C",800,"K",100000,0,0]]],["710","RANGKAS JAYA","0800","1115",[["C",424,"K",80000,0,0]]],["77A","ARGO GOPAR","1200","1500",[["A",49,"E",100000,0,0],["B",50,"B",90000,0,0],["C",49,"K",60000,0,0]]],["P05","ARGO PARAHYANGAN","2000","2300",[["A",0,"E",200000,0,0],["B",0,"B",150000,0,0],["C",240,"K",60000,0,0]]]] }`

	// transform to standard response
	stdStr := TrasformStandardKAIResponse([]byte(str))

	// test expected values
	errCode := "0"
	orgCode := "BD"
	desCode := "GMR"
	departDate := "20190919"
	schedLen := 4
	tripTrainNo0 := "10501"
	tripTrainName0 := "ARGO PARAHYANGAN PREMIUM"

	// test variable
	var vRS GetScheduleRS

	// test function
	err := json.Unmarshal(stdStr, &vRS)

	result, err := MakeInternalGetScheduleRS(vRS)

	// r, _ := json.Marshal(result)
	// fmt.Println(string(r))

	// test logic
	assert.Equal(t, errCode, result.ErrCode, "sould be equal!")

	assert.Equal(t, orgCode, result.Return.Origin, "sould be equal!")
	assert.Equal(t, desCode, result.Return.Destination, "sould be equal!")
	assert.Equal(t, departDate, result.Return.DepartureDate, "sould be equal!")
	assert.Equal(t, schedLen, len(result.Return.Schedules), "sould be equal!")

	assert.Equal(t, tripTrainNo0, result.Return.Schedules[0].TrainNo, "sould be equal!")
	assert.Equal(t, tripTrainName0, result.Return.Schedules[0].TrainName, "sould be equal!")

	assert.Nil(t, err)
}

// TestMakeInternalGetScheduleV2RSOK is a positive test function for "MakeInternalGetScheduleV2RS" mapper method
func TestMakeInternalGetScheduleV2RSOK(t *testing.T) {
	// fake response
	str := `{ "err_code": 0, "org": "BD", "des": "GMR", "dep_date": "20190919", "schedule": [["10501","ARGO PARAHYANGAN PREMIUM","20190919","20190919","0415","0725",[["C",800,"K",100000,0,0]]],["710","RANGKAS JAYA","20190919","20190919","0800","1115",[["C",424,"K",80000,0,0]]],["77A","ARGO GOPAR","20190919","20190919","1200","1500",[["A",49,"E",100000,0,0],["B",50,"B",90000,0,0],["C",49,"K",60000,0,0]]],["P05","ARGO PARAHYANGAN","20190919","20190919","2000","2300",[["A",0,"E",200000,0,0],["B",0,"B",150000,0,0],["C",240,"K",60000,0,0]]]] }`

	// transform to standard response
	stdStr := TrasformStandardKAIResponse([]byte(str))

	// test expected values
	errCode := "0"
	orgCode := "BD"
	desCode := "GMR"
	departDate := "20190919"
	schedLen := 4
	tripTrainNo0 := "10501"
	tripTrainName0 := "ARGO PARAHYANGAN PREMIUM"
	tripDepDate0 := "20190919"
	tripArvDate0 := "20190919"

	// test variable
	var vRS GetScheduleV2RS

	// test function
	err := json.Unmarshal(stdStr, &vRS)

	result, err := MakeInternalGetScheduleV2RS(vRS)

	// r, _ := json.Marshal(result)
	// fmt.Println(string(r))

	// test logic
	assert.Equal(t, errCode, result.ErrCode, "sould be equal!")

	assert.Equal(t, orgCode, result.Return.Origin, "sould be equal!")
	assert.Equal(t, desCode, result.Return.Destination, "sould be equal!")
	assert.Equal(t, departDate, result.Return.DepartureDate, "sould be equal!")
	assert.Equal(t, schedLen, len(result.Return.Schedules), "sould be equal!")

	assert.Equal(t, tripTrainNo0, result.Return.Schedules[0].TrainNo, "sould be equal!")
	assert.Equal(t, tripTrainName0, result.Return.Schedules[0].TrainName, "sould be equal!")
	assert.Equal(t, tripDepDate0, result.Return.Schedules[0].DepartureDate, "sould be equal!")
	assert.Equal(t, tripArvDate0, result.Return.Schedules[0].ArriveDate, "sould be equal!")

	assert.Nil(t, err)
}

// TestMakeInternalGetScheduleLiteRSOK is a positive test function for "MakeInternalGetScheduleLiteRS" mapper method
func TestMakeInternalGetScheduleLiteRSOK(t *testing.T) {
	// fake response
	str := `{ "err_code": 0, "org": "BD", "des": "GMR", "dep_date": "20190919", "schedule": [["10501","ARGO PARAHYANGAN PREMIUM","20190919","20190919","0415","0725",[["C",0,"K",100000,0,0]]],["710","RANGKAS JAYA","20190919","20190919","0800","1115",[["C",0,"K",80000,0,0]]],["77A","ARGO GOPAR","20190919","20190919","1200","1500",[["A",0,"E",100000,0,0],["B",0,"B",90000,0,0],["C",0,"K",60000,0,0]]],["P05","ARGO PARAHYANGAN","20190919","20190919","2000","2300",[["A",0,"E",200000,0,0],["B",0,"B",150000,0,0],["C",0,"K",60000,0,0]]]] }`

	// transform to standard response
	stdStr := TrasformStandardKAIResponse([]byte(str))

	// test expected values
	errCode := "0"
	orgCode := "BD"
	desCode := "GMR"
	departDate := "20190919"
	schedLen := 4
	tripTrainNo0 := "10501"
	tripTrainName0 := "ARGO PARAHYANGAN PREMIUM"
	tripDepDate0 := "20190919"
	tripArvDate0 := "20190919"
	var tripSeatAvailC0 float64 // = 0

	// test variable
	var vRS GetScheduleLiteRS

	// test function
	err := json.Unmarshal(stdStr, &vRS)

	result, err := MakeInternalGetScheduleLiteRS(vRS)

	// r, _ := json.Marshal(result)
	// fmt.Println(string(r))

	// test logic
	assert.Equal(t, errCode, result.ErrCode, "sould be equal!")

	assert.Equal(t, orgCode, result.Return.Origin, "sould be equal!")
	assert.Equal(t, desCode, result.Return.Destination, "sould be equal!")
	assert.Equal(t, departDate, result.Return.DepartureDate, "sould be equal!")
	assert.Equal(t, schedLen, len(result.Return.Schedules), "sould be equal!")

	assert.Equal(t, tripTrainNo0, result.Return.Schedules[0].TrainNo, "sould be equal!")
	assert.Equal(t, tripTrainName0, result.Return.Schedules[0].TrainName, "sould be equal!")
	assert.Equal(t, tripDepDate0, result.Return.Schedules[0].DepartureDate, "sould be equal!")
	assert.Equal(t, tripArvDate0, result.Return.Schedules[0].ArriveDate, "sould be equal!")

	assert.Equal(t, tripSeatAvailC0, result.Return.Schedules[0].AvailSubClass[0].SeatAvailable, "sould be equal!")

	assert.Nil(t, err)
}
