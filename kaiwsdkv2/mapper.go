package kaiwsdkv2

import (
	"fmt"
	"reflect"
)

// MakeInternalGetOriginationRS is a function to mapping from GetOriginationRS to InternalGetOriginationRS ("data.get_org")
func MakeInternalGetOriginationRS(input GetOriginationRS) (result *InternalGetOriginationRS, err error) {

	var vRS InternalGetOriginationRS

	vRS.ErrCode = input.ErrCode
	vRS.ErrMsg = input.ErrMsg

	var vReturn []Origination
	for _, v := range input.Originations {
		vReturn = append(vReturn, Origination{OriginCode: fmt.Sprintf("%s", v[0]), OriginName: fmt.Sprintf("%s", v[1])})
	}
	vRS.Return = vReturn

	result = &vRS

	return result, nil
}

// MakeInternalGetDestinationRS is a function to mapping from GetDestinationRS to InternalGetDestinationRS ("data.get_des")
func MakeInternalGetDestinationRS(input GetDestinationRS) (result *InternalGetDestinationRS, err error) {
	var vRS InternalGetDestinationRS

	vRS.ErrCode = input.ErrCode
	vRS.ErrMsg = input.ErrMsg

	var vReturn []Destination
	for _, v := range input.Destinations {
		vReturn = append(vReturn, Destination{DestCode: fmt.Sprintf("%s", v[0]), DestName: fmt.Sprintf("%s", v[1])})
	}
	vRS.Return = vReturn

	result = &vRS

	return result, nil
}

// MakeInternalGetPayTypeRS is a function to mapping from GetPayTypeRS to InternalGetDestinationRS ("data.get_des")
func MakeInternalGetPayTypeRS(input GetPayTypeRS) (result *InternalGetPayTypeRS, err error) {
	var vRS InternalGetPayTypeRS

	vRS.ErrCode = input.ErrCode
	vRS.ErrMsg = input.ErrMsg

	var vReturn []PayType
	for _, v := range input.PayTypes {
		vReturn = append(vReturn, PayType{Name: fmt.Sprintf("%s", v)})
	}
	vRS.Return = vReturn

	result = &vRS

	return result, nil
}

// MakeInternalGetScheduleRS is a function to mapping from GetScheduleRS to InternalGetScheduleRS ("information.get_schedule")
func MakeInternalGetScheduleRS(input GetScheduleRS) (result *InternalGetScheduleRS, err error) {
	var vRS InternalGetScheduleRS

	vRS.ErrCode = input.ErrCode
	vRS.ErrMsg = input.ErrMsg

	var vReturn GetSchedule

	vReturn.Origin = input.Origin
	vReturn.Destination = input.Destination
	vReturn.DepartureDate = input.DepartureDate

	var vArrSchedule []Schedule

	for _, s := range input.Schedules {
		tripInfo := reflect.ValueOf(s)

		vSchedule := Schedule{
			TrainNo:       tripInfo.Index(0).Interface().(string),
			TrainName:     tripInfo.Index(1).Interface().(string),
			DepartureTime: tripInfo.Index(2).Interface().(string),
			ArriveTime:    tripInfo.Index(3).Interface().(string),
		}

		subClass := tripInfo.Index(4).Interface().([]interface{})
		var vArrSubClass []AvailSubClass

		for _, sc := range subClass {

			scInfo := reflect.ValueOf(sc)

			vSubClass := AvailSubClass{
				SubClass:      scInfo.Index(0).Interface().(string),
				SeatAvailable: scInfo.Index(1).Interface().(float64),
				SeatClass:     scInfo.Index(2).Interface().(string),
				AdultPrice:    scInfo.Index(3).Interface().(float64),
				ChildPrice:    scInfo.Index(4).Interface().(float64),
				InfantPrice:   scInfo.Index(5).Interface().(float64),
			}

			vArrSubClass = append(vArrSubClass, vSubClass)
		}

		vSchedule.AvailSubClass = vArrSubClass
		vArrSchedule = append(vArrSchedule, vSchedule)
	}

	vReturn.Schedules = vArrSchedule
	vRS.Return = vReturn

	result = &vRS

	return result, nil
}

// MakeInternalGetScheduleV2RS is a function to mapping from GetScheduleV2RS to InternalGetScheduleV2RS ("information.get_schedule_v2")
func MakeInternalGetScheduleV2RS(input GetScheduleV2RS) (result *InternalGetScheduleV2RS, err error) {
	var vRS InternalGetScheduleV2RS

	vRS.ErrCode = input.ErrCode
	vRS.ErrMsg = input.ErrMsg

	var vReturn GetScheduleV2

	vReturn.Origin = input.Origin
	vReturn.Destination = input.Destination
	vReturn.DepartureDate = input.DepartureDate

	var vArrSchedule []ScheduleV2

	for _, s := range input.Schedules {
		tripInfo := reflect.ValueOf(s)

		vSchedule := ScheduleV2{
			TrainNo:       tripInfo.Index(0).Interface().(string),
			TrainName:     tripInfo.Index(1).Interface().(string),
			DepartureDate: tripInfo.Index(2).Interface().(string),
			ArriveDate:    tripInfo.Index(3).Interface().(string),
			DepartureTime: tripInfo.Index(4).Interface().(string),
			ArriveTime:    tripInfo.Index(5).Interface().(string),
		}

		subClass := tripInfo.Index(6).Interface().([]interface{})
		var vArrSubClass []AvailSubClassV2

		for _, sc := range subClass {

			scInfo := reflect.ValueOf(sc)

			vSubClass := AvailSubClassV2{
				SubClass:      scInfo.Index(0).Interface().(string),
				SeatAvailable: scInfo.Index(1).Interface().(float64),
				SeatClass:     scInfo.Index(2).Interface().(string),
				AdultPrice:    scInfo.Index(3).Interface().(float64),
				ChildPrice:    scInfo.Index(4).Interface().(float64),
				InfantPrice:   scInfo.Index(5).Interface().(float64),
			}

			vArrSubClass = append(vArrSubClass, vSubClass)
		}

		vSchedule.AvailSubClass = vArrSubClass
		vArrSchedule = append(vArrSchedule, vSchedule)
	}

	vReturn.Schedules = vArrSchedule
	vRS.Return = vReturn

	result = &vRS

	return result, nil
}

// MakeInternalGetScheduleLiteRS is a function to mapping from GetScheduleLiteRS to InternalGetScheduleLiteRS ("information.get_schedule_lite")
func MakeInternalGetScheduleLiteRS(input GetScheduleLiteRS) (result *InternalGetScheduleLiteRS, err error) {
	var vRS InternalGetScheduleLiteRS

	vRS.ErrCode = input.ErrCode
	vRS.ErrMsg = input.ErrMsg

	var vReturn GetScheduleLite

	vReturn.Origin = input.Origin
	vReturn.Destination = input.Destination
	vReturn.DepartureDate = input.DepartureDate

	var vArrSchedule []ScheduleLite

	for _, s := range input.Schedules {
		tripInfo := reflect.ValueOf(s)

		vSchedule := ScheduleLite{
			TrainNo:       tripInfo.Index(0).Interface().(string),
			TrainName:     tripInfo.Index(1).Interface().(string),
			DepartureDate: tripInfo.Index(2).Interface().(string),
			ArriveDate:    tripInfo.Index(3).Interface().(string),
			DepartureTime: tripInfo.Index(4).Interface().(string),
			ArriveTime:    tripInfo.Index(5).Interface().(string),
		}

		subClass := tripInfo.Index(6).Interface().([]interface{})
		var vArrSubClass []AvailSubClassLite

		for _, sc := range subClass {

			scInfo := reflect.ValueOf(sc)

			vSubClass := AvailSubClassLite{
				SubClass:      scInfo.Index(0).Interface().(string),
				SeatAvailable: scInfo.Index(1).Interface().(float64),
				SeatClass:     scInfo.Index(2).Interface().(string),
				AdultPrice:    scInfo.Index(3).Interface().(float64),
				ChildPrice:    scInfo.Index(4).Interface().(float64),
				InfantPrice:   scInfo.Index(5).Interface().(float64),
			}

			vArrSubClass = append(vArrSubClass, vSubClass)
		}

		vSchedule.AvailSubClass = vArrSubClass
		vArrSchedule = append(vArrSchedule, vSchedule)
	}

	vReturn.Schedules = vArrSchedule
	vRS.Return = vReturn

	result = &vRS

	return result, nil
}
