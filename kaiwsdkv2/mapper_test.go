// Copyright 2018 The ClientSDK Team Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0-style
// license that can be found in the LICENSE file.

// Author: ClientSDK Team (muharihar)

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

// TestMakeInternalGetSeatMapRSOK is a positive test function for "MakeInternalGetSeatMapRS" mapper method
func TestMakeInternalGetSeatMapRSOK(t *testing.T) {
	// fake response
	str := `{ "err_code": 0, "org": "BD", "des": "GMR", "train_no": "10501", "dep_date": "20190919", "seat_map": [["PREMIUM",1,[[1,1,1,"A","C",0],[1,2,1,"B","C",0],[1,3,1,"","",0],[1,4,1,"C","C",0],[1,5,1,"D","C",0],[1,6,1,"E","C",0],[2,1,2,"A","C",0],[2,2,2,"B","C",0],[2,3,2,"","",0],[2,4,2,"C","C",0],[2,5,2,"D","C",0],[2,6,2,"E","C",0],[3,1,3,"A","C",0],[3,2,3,"B","C",0],[3,3,3,"","",0],[3,4,3,"C","C",0],[3,5,3,"D","C",0],[3,6,3,"E","C",0],[4,1,4,"A","C",0],[4,2,4,"B","C",0],[4,3,4,"","",0],[4,4,4,"C","C",0],[4,5,4,"D","C",0],[4,6,4,"E","C",0],[5,1,5,"A","C",0],[5,2,5,"B","C",0],[5,3,5,"","",0],[5,4,5,"C","C",0],[5,5,5,"D","C",0],[5,6,5,"E","C",0],[6,1,6,"A","C",0],[6,2,6,"B","C",0],[6,3,6,"","",0],[6,4,6,"C","C",0],[6,5,6,"D","C",0],[6,6,6,"E","C",0],[7,1,7,"A","C",0],[7,2,7,"B","C",0],[7,3,7,"","",0],[7,4,7,"C","C",0],[7,5,7,"D","C",0],[7,6,7,"E","C",0],[8,1,8,"A","C",0],[8,2,8,"B","C",0],[8,3,8,"","",0],[8,4,8,"C","C",0],[8,5,8,"D","C",0],[8,6,8,"E","C",0],[9,1,9,"A","C",0],[9,2,9,"B","C",0],[9,3,9,"","",0],[9,4,9,"C","C",0],[9,5,9,"D","C",0],[9,6,9,"E","C",0],[10,1,10,"A","C",0],[10,2,10,"B","C",0],[10,3,10,"","",0],[10,4,10,"C","C",0],[10,5,10,"D","C",0],[10,6,10,"E","C",0],[11,1,11,"A","C",0],[11,2,11,"B","C",0],[11,3,11,"","",0],[11,4,11,"C","C",0],[11,5,11,"D","C",0],[11,6,11,"E","C",0],[12,1,12,"A","C",0],[12,2,12,"B","C",0],[12,3,12,"","",0],[12,4,12,"C","C",0],[12,5,12,"D","C",0],[12,6,12,"E","C",0],[13,1,13,"A","C",0],[13,2,13,"B","C",0],[13,3,13,"","",0],[13,4,13,"C","C",0],[13,5,13,"D","C",0],[13,6,13,"E","C",0],[14,1,14,"A","C",0],[14,2,14,"B","C",0],[14,3,14,"","",0],[14,4,14,"C","C",0],[14,5,14,"D","C",0],[14,6,14,"E","C",0],[15,1,15,"A","C",0],[15,2,15,"B","C",0],[15,3,15,"","",0],[15,4,15,"C","C",0],[15,5,15,"D","C",0],[15,6,15,"E","C",0],[16,1,16,"A","C",0],[16,2,16,"B","C",0],[16,3,16,"","",0],[16,4,16,"C","C",0],[16,5,16,"D","C",0],[16,6,16,"E","C",0],[17,1,17,"A","C",0],[17,2,17,"B","C",0],[17,3,17,"","",0],[17,4,17,"C","C",0],[17,5,17,"D","C",0],[17,6,17,"E","C",0]]],["PREMIUM",2,[[1,1,1,"A","C",0],[1,2,1,"B","C",0],[1,3,1,"","",0],[1,4,1,"C","C",0],[1,5,1,"D","C",0],[1,6,1,"E","C",0],[2,1,2,"A","C",0],[2,2,2,"B","C",0],[2,3,2,"","",0],[2,4,2,"C","C",0],[2,5,2,"D","C",0],[2,6,2,"E","C",0],[3,1,3,"A","C",0],[3,2,3,"B","C",0],[3,3,3,"","",0],[3,4,3,"C","C",0],[3,5,3,"D","C",0],[3,6,3,"E","C",0],[4,1,4,"A","C",0],[4,2,4,"B","C",0],[4,3,4,"","",0],[4,4,4,"C","C",0],[4,5,4,"D","C",0],[4,6,4,"E","C",0],[5,1,5,"A","C",0],[5,2,5,"B","C",0],[5,3,5,"","",0],[5,4,5,"C","C",0],[5,5,5,"D","C",0],[5,6,5,"E","C",0],[6,1,6,"A","C",0],[6,2,6,"B","C",0],[6,3,6,"","",0],[6,4,6,"C","C",0],[6,5,6,"D","C",0],[6,6,6,"E","C",0],[7,1,7,"A","C",0],[7,2,7,"B","C",0],[7,3,7,"","",0],[7,4,7,"C","C",0],[7,5,7,"D","C",0],[7,6,7,"E","C",0],[8,1,8,"A","C",0],[8,2,8,"B","C",0],[8,3,8,"","",0],[8,4,8,"C","C",0],[8,5,8,"D","C",0],[8,6,8,"E","C",0],[9,1,9,"A","C",0],[9,2,9,"B","C",0],[9,3,9,"","",0],[9,4,9,"C","C",0],[9,5,9,"D","C",0],[9,6,9,"E","C",0],[10,1,10,"A","C",0],[10,2,10,"B","C",0],[10,3,10,"","",0],[10,4,10,"C","C",0],[10,5,10,"D","C",0],[10,6,10,"E","C",0],[11,1,11,"A","C",0],[11,2,11,"B","C",0],[11,3,11,"","",0],[11,4,11,"C","C",0],[11,5,11,"D","C",0],[11,6,11,"E","C",0],[12,1,12,"A","C",0],[12,2,12,"B","C",0],[12,3,12,"","",0],[12,4,12,"C","C",0],[12,5,12,"D","C",0],[12,6,12,"E","C",0],[13,1,13,"A","C",0],[13,2,13,"B","C",0],[13,3,13,"","",0],[13,4,13,"C","C",0],[13,5,13,"D","C",0],[13,6,13,"E","C",0],[14,1,14,"A","C",0],[14,2,14,"B","C",0],[14,3,14,"","",0],[14,4,14,"C","C",0],[14,5,14,"D","C",0],[14,6,14,"E","C",0],[15,1,15,"A","C",0],[15,2,15,"B","C",0],[15,3,15,"","",0],[15,4,15,"C","C",0],[15,5,15,"D","C",0],[15,6,15,"E","C",0],[16,1,16,"A","C",0],[16,2,16,"B","C",0],[16,3,16,"","",0],[16,4,16,"C","C",0],[16,5,16,"D","C",0],[16,6,16,"E","C",0],[17,1,17,"A","C",0],[17,2,17,"B","C",0],[17,3,17,"","",0],[17,4,17,"C","C",0],[17,5,17,"D","C",0],[17,6,17,"E","C",0]]]]}`

	// transform to standard response
	stdStr := TrasformStandardKAIResponse([]byte(str))

	// test expected values
	errCode := "0"
	orgCode := "BD"
	desCode := "GMR"
	departDate := "20190919"
	trainNo := "10501"
	smLen := 2
	wagonCode := "PREMIUM"

	// test variable
	var vRS GetSeatMapRS

	// test function
	err := json.Unmarshal(stdStr, &vRS)

	result, err := MakeInternalGetSeatMapRS(vRS)

	// r, _ := json.Marshal(result)
	// fmt.Println(string(r))

	// test logic
	assert.Equal(t, errCode, result.ErrCode, "sould be equal!")

	assert.Equal(t, orgCode, result.Return.Origin, "sould be equal!")
	assert.Equal(t, desCode, result.Return.Destination, "sould be equal!")
	assert.Equal(t, departDate, result.Return.DepartureDate, "sould be equal!")
	assert.Equal(t, trainNo, result.Return.TrainNo, "sould be equal!")
	assert.Equal(t, smLen, len(result.Return.SeatMaps), "sould be equal!")

	assert.Equal(t, wagonCode, result.Return.SeatMaps[0].WagonCode, "sould be equal!")

	assert.Nil(t, err)
}

// TestMakeInternalGetSeatMapPerSubClassRSOK is a positive test function for "MakeInternalGetSeatMapPerSubClassRS" mapper method
func TestMakeInternalGetSeatMapPerSubClassRSOK(t *testing.T) {
	// fake response
	str := `{ "err_code": 0, "org": "BD", "des": "GMR", "train_no": "10501", "dep_date": "20190919", "seat_map": [["PREMIUM",1,[[1,1,1,"A","C",0],[1,2,1,"B","C",0],[1,4,1,"C","C",0],[1,5,1,"D","C",0],[1,6,1,"E","C",0],[2,1,2,"A","C",0],[2,2,2,"B","C",0],[2,4,2,"C","C",0],[2,5,2,"D","C",0],[2,6,2,"E","C",0],[3,1,3,"A","C",0],[3,2,3,"B","C",0],[3,4,3,"C","C",0],[3,5,3,"D","C",0],[3,6,3,"E","C",0],[4,1,4,"A","C",0],[4,2,4,"B","C",0],[4,4,4,"C","C",0],[4,5,4,"D","C",0],[4,6,4,"E","C",0],[5,1,5,"A","C",0],[5,2,5,"B","C",0],[5,4,5,"C","C",0],[5,5,5,"D","C",0],[5,6,5,"E","C",0],[6,1,6,"A","C",0],[6,2,6,"B","C",0],[6,4,6,"C","C",0],[6,5,6,"D","C",0],[6,6,6,"E","C",0],[7,1,7,"A","C",0],[7,2,7,"B","C",0],[7,4,7,"C","C",0],[7,5,7,"D","C",0],[7,6,7,"E","C",0],[8,1,8,"A","C",0],[8,2,8,"B","C",0],[8,4,8,"C","C",0],[8,5,8,"D","C",0],[8,6,8,"E","C",0],[9,1,9,"A","C",0],[9,2,9,"B","C",0],[9,4,9,"C","C",0],[9,5,9,"D","C",0],[9,6,9,"E","C",0],[10,1,10,"A","C",0],[10,2,10,"B","C",0],[10,4,10,"C","C",0],[10,5,10,"D","C",0],[10,6,10,"E","C",0],[11,1,11,"A","C",0],[11,2,11,"B","C",0],[11,4,11,"C","C",0],[11,5,11,"D","C",0],[11,6,11,"E","C",0],[12,1,12,"A","C",0],[12,2,12,"B","C",0],[12,4,12,"C","C",0],[12,5,12,"D","C",0],[12,6,12,"E","C",0],[13,1,13,"A","C",0],[13,2,13,"B","C",0],[13,4,13,"C","C",0],[13,5,13,"D","C",0],[13,6,13,"E","C",0],[14,1,14,"A","C",0],[14,2,14,"B","C",0],[14,4,14,"C","C",0],[14,5,14,"D","C",0],[14,6,14,"E","C",0],[15,1,15,"A","C",0],[15,2,15,"B","C",0],[15,4,15,"C","C",0],[15,5,15,"D","C",0],[15,6,15,"E","C",0],[16,1,16,"A","C",0],[16,2,16,"B","C",0],[16,4,16,"C","C",0],[16,5,16,"D","C",0],[16,6,16,"E","C",0],[17,1,17,"A","C",0],[17,2,17,"B","C",0],[17,4,17,"C","C",0],[17,5,17,"D","C",0],[17,6,17,"E","C",0]]],["PREMIUM",2,[[1,1,1,"A","C",0],[1,2,1,"B","C",0],[1,4,1,"C","C",0],[1,5,1,"D","C",0],[1,6,1,"E","C",0],[2,1,2,"A","C",0],[2,2,2,"B","C",0],[2,4,2,"C","C",0],[2,5,2,"D","C",0],[2,6,2,"E","C",0],[3,1,3,"A","C",0],[3,2,3,"B","C",0],[3,4,3,"C","C",0],[3,5,3,"D","C",0],[3,6,3,"E","C",0],[4,1,4,"A","C",0],[4,2,4,"B","C",0],[4,4,4,"C","C",0],[4,5,4,"D","C",0],[4,6,4,"E","C",0],[5,1,5,"A","C",0],[5,2,5,"B","C",0],[5,4,5,"C","C",0],[5,5,5,"D","C",0],[5,6,5,"E","C",0],[6,1,6,"A","C",0],[6,2,6,"B","C",0],[6,4,6,"C","C",0],[6,5,6,"D","C",0],[6,6,6,"E","C",0],[7,1,7,"A","C",0],[7,2,7,"B","C",0],[7,4,7,"C","C",0],[7,5,7,"D","C",0],[7,6,7,"E","C",0],[8,1,8,"A","C",0],[8,2,8,"B","C",0],[8,4,8,"C","C",0],[8,5,8,"D","C",0],[8,6,8,"E","C",0],[9,1,9,"A","C",0],[9,2,9,"B","C",0],[9,4,9,"C","C",0],[9,5,9,"D","C",0],[9,6,9,"E","C",0],[10,1,10,"A","C",0],[10,2,10,"B","C",0],[10,4,10,"C","C",0],[10,5,10,"D","C",0],[10,6,10,"E","C",0],[11,1,11,"A","C",0],[11,2,11,"B","C",0],[11,4,11,"C","C",0],[11,5,11,"D","C",0],[11,6,11,"E","C",0],[12,1,12,"A","C",0],[12,2,12,"B","C",0],[12,4,12,"C","C",0],[12,5,12,"D","C",0],[12,6,12,"E","C",0],[13,1,13,"A","C",0],[13,2,13,"B","C",0],[13,4,13,"C","C",0],[13,5,13,"D","C",0],[13,6,13,"E","C",0],[14,1,14,"A","C",0],[14,2,14,"B","C",0],[14,4,14,"C","C",0],[14,5,14,"D","C",0],[14,6,14,"E","C",0],[15,1,15,"A","C",0],[15,2,15,"B","C",0],[15,4,15,"C","C",0],[15,5,15,"D","C",0],[15,6,15,"E","C",0],[16,1,16,"A","C",0],[16,2,16,"B","C",0],[16,4,16,"C","C",0],[16,5,16,"D","C",0],[16,6,16,"E","C",0],[17,1,17,"A","C",0],[17,2,17,"B","C",0],[17,4,17,"C","C",0],[17,5,17,"D","C",0],[17,6,17,"E","C",0]]]] }`

	// transform to standard response
	stdStr := TrasformStandardKAIResponse([]byte(str))

	// test expected values
	errCode := "0"
	orgCode := "BD"
	desCode := "GMR"
	departDate := "20190919"
	trainNo := "10501"
	smLen := 2
	wagonCode := "PREMIUM"
	subClass0_0 := "C"

	// test variable
	var vRS GetSeatMapPerSubClassRS

	// test function
	err := json.Unmarshal(stdStr, &vRS)

	result, err := MakeInternalGetSeatMapPerSubClassRS(vRS)

	// r, _ := json.Marshal(result)
	// fmt.Println(string(r))

	// test logic
	assert.Equal(t, errCode, result.ErrCode, "sould be equal!")

	assert.Equal(t, orgCode, result.Return.Origin, "sould be equal!")
	assert.Equal(t, desCode, result.Return.Destination, "sould be equal!")
	assert.Equal(t, departDate, result.Return.DepartureDate, "sould be equal!")
	assert.Equal(t, trainNo, result.Return.TrainNo, "sould be equal!")
	assert.Equal(t, smLen, len(result.Return.SeatMaps), "sould be equal!")

	assert.Equal(t, wagonCode, result.Return.SeatMaps[0].WagonCode, "sould be equal!")
	assert.Equal(t, subClass0_0, result.Return.SeatMaps[0].Seats[0].SubClass, "sould be equal!")

	assert.Nil(t, err)
}

// TestMakeInternalGetSeatNullRSOK is a positive test function for "MakeInternalGetSeatNullRS" mapper method
func TestMakeInternalGetSeatNullRSOK(t *testing.T) {
	// fake response
	str := `{ "err_code": 0, "org": "BD", "des": "GMR", "train_no": "10501", "dep_date": "20190919", "seat_null": [["PREMIUM",1,[[1,"A","C",0],[1,"B","C",0],[1,"C","C",0],[1,"D","C",0],[1,"E","C",0],[2,"A","C",0],[2,"B","C",0],[2,"C","C",0],[2,"D","C",0],[2,"E","C",0],[3,"A","C",0],[3,"B","C",0],[3,"C","C",0],[3,"D","C",0],[3,"E","C",0],[4,"A","C",0],[4,"B","C",0],[4,"C","C",0],[4,"D","C",0],[4,"E","C",0],[5,"A","C",0],[5,"B","C",0],[5,"C","C",0],[5,"D","C",0],[5,"E","C",0],[6,"A","C",0],[6,"B","C",0],[6,"C","C",0],[6,"D","C",0],[6,"E","C",0],[7,"A","C",0],[7,"B","C",0],[7,"C","C",0],[7,"D","C",0],[7,"E","C",0],[8,"A","C",0],[8,"B","C",0],[8,"C","C",0],[8,"D","C",0],[8,"E","C",0],[9,"A","C",0],[9,"B","C",0],[9,"C","C",0],[9,"D","C",0],[9,"E","C",0],[10,"A","C",0],[10,"B","C",0],[10,"C","C",0],[10,"D","C",0],[10,"E","C",0],[11,"A","C",0],[11,"B","C",0],[11,"C","C",0],[11,"D","C",0],[11,"E","C",0],[12,"A","C",0],[12,"B","C",0],[12,"C","C",0],[12,"D","C",0],[12,"E","C",0],[13,"A","C",0],[13,"B","C",0],[13,"C","C",0],[13,"D","C",0],[13,"E","C",0],[14,"A","C",0],[14,"B","C",0],[14,"C","C",0],[14,"D","C",0],[14,"E","C",0],[15,"A","C",0],[15,"B","C",0],[15,"C","C",0],[15,"D","C",0],[15,"E","C",0],[16,"A","C",0],[16,"B","C",0],[16,"C","C",0],[16,"D","C",0],[16,"E","C",0],[17,"A","C",0],[17,"B","C",0],[17,"C","C",0],[17,"D","C",0],[17,"E","C",0]]],["PREMIUM",2,[[1,"A","C",0],[1,"B","C",0],[1,"C","C",0],[1,"D","C",0],[1,"E","C",0],[2,"A","C",0],[2,"B","C",0],[2,"C","C",0],[2,"D","C",0],[2,"E","C",0],[3,"A","C",0],[3,"B","C",0],[3,"C","C",0],[3,"D","C",0],[3,"E","C",0],[4,"A","C",0],[4,"B","C",0],[4,"C","C",0],[4,"D","C",0],[4,"E","C",0],[5,"A","C",0],[5,"B","C",0],[5,"C","C",0],[5,"D","C",0],[5,"E","C",0],[6,"A","C",0],[6,"B","C",0],[6,"C","C",0],[6,"D","C",0],[6,"E","C",0],[7,"A","C",0],[7,"B","C",0],[7,"C","C",0],[7,"D","C",0],[7,"E","C",0],[8,"A","C",0],[8,"B","C",0],[8,"C","C",0],[8,"D","C",0],[8,"E","C",0],[9,"A","C",0],[9,"B","C",0],[9,"C","C",0],[9,"D","C",0],[9,"E","C",0],[10,"A","C",0],[10,"B","C",0],[10,"C","C",0],[10,"D","C",0],[10,"E","C",0],[11,"A","C",0],[11,"B","C",0],[11,"C","C",0],[11,"D","C",0],[11,"E","C",0],[12,"A","C",0],[12,"B","C",0],[12,"C","C",0],[12,"D","C",0],[12,"E","C",0],[13,"A","C",0],[13,"B","C",0],[13,"C","C",0],[13,"D","C",0],[13,"E","C",0],[14,"A","C",0],[14,"B","C",0],[14,"C","C",0],[14,"D","C",0],[14,"E","C",0],[15,"A","C",0],[15,"B","C",0],[15,"C","C",0],[15,"D","C",0],[15,"E","C",0],[16,"A","C",0],[16,"B","C",0],[16,"C","C",0],[16,"D","C",0],[16,"E","C",0],[17,"A","C",0],[17,"B","C",0],[17,"C","C",0],[17,"D","C",0],[17,"E","C",0]]]]}`

	// transform to standard response
	stdStr := TrasformStandardKAIResponse([]byte(str))

	// test expected values
	errCode := "0"
	orgCode := "BD"
	desCode := "GMR"
	departDate := "20190919"
	trainNo := "10501"
	smLen := 2
	wagonCode := "PREMIUM"
	var statusAvail0_0 float64 // = 0

	// test variable
	var vRS GetSeatNullRS

	// test function
	err := json.Unmarshal(stdStr, &vRS)

	result, err := MakeInternalGetSeatNullRS(vRS)

	// r, _ := json.Marshal(result)
	// fmt.Println(string(r))

	// test logic
	assert.Equal(t, errCode, result.ErrCode, "sould be equal!")

	assert.Equal(t, orgCode, result.Return.Origin, "sould be equal!")
	assert.Equal(t, desCode, result.Return.Destination, "sould be equal!")
	assert.Equal(t, departDate, result.Return.DepartureDate, "sould be equal!")
	assert.Equal(t, trainNo, result.Return.TrainNo, "sould be equal!")
	assert.Equal(t, smLen, len(result.Return.SeatNulls), "sould be equal!")

	assert.Equal(t, wagonCode, result.Return.SeatNulls[0].WagonCode, "sould be equal!")
	assert.Equal(t, statusAvail0_0, result.Return.SeatNulls[0].Seats[0].Status, "sould be equal!")

	assert.Nil(t, err)
}

// TestMakeInternalGetSeatNullPerSubClassRSOK is a positive test function for "MakeInternalGetSeatNullPerSubClassRS" mapper method
func TestMakeInternalGetSeatNullPerSubClassRSOK(t *testing.T) {
	// fake response
	str := `{ "err_code": 0, "org": "BD", "des": "GMR", "train_no": "10501", "dep_date": "20190919", "seat_null": [["PREMIUM",1,[[1,"A","C",0],[1,"B","C",0],[1,"C","C",0],[1,"D","C",0],[1,"E","C",0],[2,"A","C",0],[2,"B","C",0],[2,"C","C",0],[2,"D","C",0],[2,"E","C",0],[3,"A","C",0],[3,"B","C",0],[3,"C","C",0],[3,"D","C",0],[3,"E","C",0],[4,"A","C",0],[4,"B","C",0],[4,"C","C",0],[4,"D","C",0],[4,"E","C",0],[5,"A","C",0],[5,"B","C",0],[5,"C","C",0],[5,"D","C",0],[5,"E","C",0],[6,"A","C",0],[6,"B","C",0],[6,"C","C",0],[6,"D","C",0],[6,"E","C",0],[7,"A","C",0],[7,"B","C",0],[7,"C","C",0],[7,"D","C",0],[7,"E","C",0],[8,"A","C",0],[8,"B","C",0],[8,"C","C",0],[8,"D","C",0],[8,"E","C",0],[9,"A","C",0],[9,"B","C",0],[9,"C","C",0],[9,"D","C",0],[9,"E","C",0],[10,"A","C",0],[10,"B","C",0],[10,"C","C",0],[10,"D","C",0],[10,"E","C",0],[11,"A","C",0],[11,"B","C",0],[11,"C","C",0],[11,"D","C",0],[11,"E","C",0],[12,"A","C",0],[12,"B","C",0],[12,"C","C",0],[12,"D","C",0],[12,"E","C",0],[13,"A","C",0],[13,"B","C",0],[13,"C","C",0],[13,"D","C",0],[13,"E","C",0],[14,"A","C",0],[14,"B","C",0],[14,"C","C",0],[14,"D","C",0],[14,"E","C",0],[15,"A","C",0],[15,"B","C",0],[15,"C","C",0],[15,"D","C",0],[15,"E","C",0],[16,"A","C",0],[16,"B","C",0],[16,"C","C",0],[16,"D","C",0],[16,"E","C",0],[17,"A","C",0],[17,"B","C",0],[17,"C","C",0],[17,"D","C",0],[17,"E","C",0]]],["PREMIUM",2,[[1,"A","C",0],[1,"B","C",0],[1,"C","C",0],[1,"D","C",0],[1,"E","C",0],[2,"A","C",0],[2,"B","C",0],[2,"C","C",0],[2,"D","C",0],[2,"E","C",0],[3,"A","C",0],[3,"B","C",0],[3,"C","C",0],[3,"D","C",0],[3,"E","C",0],[4,"A","C",0],[4,"B","C",0],[4,"C","C",0],[4,"D","C",0],[4,"E","C",0],[5,"A","C",0],[5,"B","C",0],[5,"C","C",0],[5,"D","C",0],[5,"E","C",0],[6,"A","C",0],[6,"B","C",0],[6,"C","C",0],[6,"D","C",0],[6,"E","C",0],[7,"A","C",0],[7,"B","C",0],[7,"C","C",0],[7,"D","C",0],[7,"E","C",0],[8,"A","C",0],[8,"B","C",0],[8,"C","C",0],[8,"D","C",0],[8,"E","C",0],[9,"A","C",0],[9,"B","C",0],[9,"C","C",0],[9,"D","C",0],[9,"E","C",0],[10,"A","C",0],[10,"B","C",0],[10,"C","C",0],[10,"D","C",0],[10,"E","C",0],[11,"A","C",0],[11,"B","C",0],[11,"C","C",0],[11,"D","C",0],[11,"E","C",0],[12,"A","C",0],[12,"B","C",0],[12,"C","C",0],[12,"D","C",0],[12,"E","C",0],[13,"A","C",0],[13,"B","C",0],[13,"C","C",0],[13,"D","C",0],[13,"E","C",0],[14,"A","C",0],[14,"B","C",0],[14,"C","C",0],[14,"D","C",0],[14,"E","C",0],[15,"A","C",0],[15,"B","C",0],[15,"C","C",0],[15,"D","C",0],[15,"E","C",0],[16,"A","C",0],[16,"B","C",0],[16,"C","C",0],[16,"D","C",0],[16,"E","C",0],[17,"A","C",0],[17,"B","C",0],[17,"C","C",0],[17,"D","C",0],[17,"E","C",0]]]]}`

	// transform to standard response
	stdStr := TrasformStandardKAIResponse([]byte(str))

	// test expected values
	errCode := "0"
	orgCode := "BD"
	desCode := "GMR"
	departDate := "20190919"
	trainNo := "10501"
	smLen := 2
	wagonCode := "PREMIUM"
	var statusAvail0_0 float64 // = 0
	subClass0_0 := "C"

	// test variable
	var vRS GetSeatNullPerSubClassRS

	// test function
	err := json.Unmarshal(stdStr, &vRS)

	result, err := MakeInternalGetSeatNullPerSubClassRS(vRS)

	// r, _ := json.Marshal(result)
	// fmt.Println(string(r))

	// test logic
	assert.Equal(t, errCode, result.ErrCode, "sould be equal!")

	assert.Equal(t, orgCode, result.Return.Origin, "sould be equal!")
	assert.Equal(t, desCode, result.Return.Destination, "sould be equal!")
	assert.Equal(t, departDate, result.Return.DepartureDate, "sould be equal!")
	assert.Equal(t, trainNo, result.Return.TrainNo, "sould be equal!")
	assert.Equal(t, smLen, len(result.Return.SeatNulls), "sould be equal!")

	assert.Equal(t, wagonCode, result.Return.SeatNulls[0].WagonCode, "sould be equal!")
	assert.Equal(t, statusAvail0_0, result.Return.SeatNulls[0].Seats[0].Status, "sould be equal!")
	assert.Equal(t, subClass0_0, result.Return.SeatNulls[0].Seats[0].SubClass, "sould be equal!")

	assert.Nil(t, err)
}

// TestMakeInternalGetAgentBalanceRSOK is a positive test function for "MakeInternalGetAgentBalanceRS" mapper method
func TestMakeInternalGetAgentBalanceRSOK(t *testing.T) {
	// fake response
	str := `{ "err_code": 0, "agent_code": "KAI_AGENT_CODE", "agent_name":"KAI_AGENT_NAME","agent_balance":20188102 }`

	// transform to standard response
	stdStr := TrasformStandardKAIResponse([]byte(str))

	// test expected values
	errCode := "0"
	agentCode := "KAI_AGENT_CODE"
	agentName := "KAI_AGENT_NAME"
	var agentBalance float64 = 20188102

	// test variable
	var vRS GetAgentBalanceRS

	// test function
	err := json.Unmarshal(stdStr, &vRS)

	result, err := MakeInternalGetAgentBalanceRS(vRS)

	// r, _ := json.Marshal(result)
	// fmt.Println(string(r))

	// test logic
	assert.Equal(t, errCode, result.ErrCode, "sould be equal!")

	assert.Equal(t, agentCode, result.Return.AgentCode, "sould be equal!")
	assert.Equal(t, agentName, result.Return.AgentName, "sould be equal!")
	assert.Equal(t, agentBalance, result.Return.AgentBalance, "sould be equal!")

	assert.Nil(t, err)
}

// TestMakeInternalGetBalanceRSOK is a positive test function for "MakeInternalGetBalanceRS" mapper method
func TestMakeInternalGetBalanceRSOK(t *testing.T) {
	// fake response
	str := `{ "err_code": 0, "book_code": "ABMNYZ", "num_code": 9998123456789, "normal_sales": 290000, "extra_fee": 0, "book_balance": 282500,"discount":-7500 }`

	// transform to standard response
	stdStr := TrasformStandardKAIResponse([]byte(str))

	// test expected values
	errCode := "0"
	bookCode := "ABMNYZ"
	numCode := "9998123456789"
	var normalSales float64 = 290000

	// test variable
	var vRS GetBalanceRS

	// test function
	err := json.Unmarshal(stdStr, &vRS)

	result, err := MakeInternalGetBalanceRS(vRS)

	// r, _ := json.Marshal(result)
	// fmt.Println(string(r))

	// test logic
	assert.Equal(t, errCode, result.ErrCode, "sould be equal!")

	assert.Equal(t, bookCode, result.Return.BookCode, "sould be equal!")
	assert.Equal(t, numCode, result.Return.NumCode, "sould be equal!")
	assert.Equal(t, normalSales, result.Return.NormalSales, "sould be equal!")

	assert.Nil(t, err)
}
