// Copyright 2018 The ClientSDK Team Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0-style
// license that can be found in the LICENSE file.

// Author: ClientSDK Team (muharihar)

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

// MakeInternalGetSeatMapRS is a function to mapping from GetSeatMapRS to InternalGetScheduleLiteRS ("information.get_seat_map")
func MakeInternalGetSeatMapRS(input GetSeatMapRS) (result *InternalGetSeatMapRS, err error) {
	var vRS InternalGetSeatMapRS

	vRS.ErrCode = input.ErrCode
	vRS.ErrMsg = input.ErrMsg

	var vReturn GetSeatMap

	vReturn.Origin = input.Origin
	vReturn.Destination = input.Destination
	vReturn.TrainNo = input.TrainNo
	vReturn.DepartureDate = input.DepartureDate

	var vArrSeatMap []SeatMap

	for _, s := range input.SeatMaps {

		tripInfo := reflect.ValueOf(s)

		vSeatMap := SeatMap{
			WagonCode: tripInfo.Index(0).Interface().(string),
			WagonNo:   tripInfo.Index(1).Interface().(float64),
		}

		seat := tripInfo.Index(2).Interface().([]interface{})
		var vArrSeat []Seat

		for _, sc := range seat {

			scInfo := reflect.ValueOf(sc)

			vSeat := Seat{
				Row:        scInfo.Index(0).Interface().(float64),
				Column:     scInfo.Index(1).Interface().(float64),
				SeatRow:    scInfo.Index(2).Interface().(float64),
				SeatColumn: scInfo.Index(3).Interface().(string),
				SubClass:   scInfo.Index(4).Interface().(string),
				Status:     scInfo.Index(5).Interface().(float64),
			}

			vArrSeat = append(vArrSeat, vSeat)
		}

		vSeatMap.Seats = vArrSeat
		vArrSeatMap = append(vArrSeatMap, vSeatMap)
	}

	vReturn.SeatMaps = vArrSeatMap
	vRS.Return = vReturn

	result = &vRS

	return result, nil
}

// MakeInternalGetSeatMapPerSubClassRS is a function to mapping from GetSeatMapPerSubClassRS to InternalGetSeatMapPerSubClassRS ("information.get_seat_map_per_subclass")
func MakeInternalGetSeatMapPerSubClassRS(input GetSeatMapPerSubClassRS) (result *InternalGetSeatMapPerSubClassRS, err error) {
	var vRS InternalGetSeatMapPerSubClassRS

	vRS.ErrCode = input.ErrCode
	vRS.ErrMsg = input.ErrMsg

	var vReturn GetSeatMapPerSubClass

	vReturn.Origin = input.Origin
	vReturn.Destination = input.Destination
	vReturn.TrainNo = input.TrainNo
	vReturn.DepartureDate = input.DepartureDate

	var vArrSeatMap []SeatMap

	for _, s := range input.SeatMaps {

		tripInfo := reflect.ValueOf(s)

		vSeatMap := SeatMap{
			WagonCode: tripInfo.Index(0).Interface().(string),
			WagonNo:   tripInfo.Index(1).Interface().(float64),
		}

		seat := tripInfo.Index(2).Interface().([]interface{})
		var vArrSeat []Seat

		for _, sc := range seat {

			scInfo := reflect.ValueOf(sc)

			vSeat := Seat{
				Row:        scInfo.Index(0).Interface().(float64),
				Column:     scInfo.Index(1).Interface().(float64),
				SeatRow:    scInfo.Index(2).Interface().(float64),
				SeatColumn: scInfo.Index(3).Interface().(string),
				SubClass:   scInfo.Index(4).Interface().(string),
				Status:     scInfo.Index(5).Interface().(float64),
			}

			vArrSeat = append(vArrSeat, vSeat)
		}

		vSeatMap.Seats = vArrSeat
		vArrSeatMap = append(vArrSeatMap, vSeatMap)
	}

	vReturn.SeatMaps = vArrSeatMap
	vRS.Return = vReturn

	result = &vRS

	return result, nil
}

// MakeInternalGetSeatNullRS is a function to mapping from GetSeatNullRS to InternalGetSeatNullRS ("information.get_seat_null")
func MakeInternalGetSeatNullRS(input GetSeatNullRS) (result *InternalGetSeatNullRS, err error) {
	var vRS InternalGetSeatNullRS

	vRS.ErrCode = input.ErrCode
	vRS.ErrMsg = input.ErrMsg

	var vReturn GetSeatNull

	vReturn.Origin = input.Origin
	vReturn.Destination = input.Destination
	vReturn.TrainNo = input.TrainNo
	vReturn.DepartureDate = input.DepartureDate

	var vArrSeatNull []SeatNull

	for _, s := range input.SeatNulls {

		tripInfo := reflect.ValueOf(s)

		vSeatNull := SeatNull{
			WagonCode: tripInfo.Index(0).Interface().(string),
			WagonNo:   tripInfo.Index(1).Interface().(float64),
		}

		seat := tripInfo.Index(2).Interface().([]interface{})
		var vArrSeatN []SeatN

		for _, sc := range seat {

			scInfo := reflect.ValueOf(sc)

			vSeatN := SeatN{
				SeatRow:    scInfo.Index(0).Interface().(float64),
				SeatColumn: scInfo.Index(1).Interface().(string),
				SubClass:   scInfo.Index(2).Interface().(string),
				Status:     scInfo.Index(3).Interface().(float64),
			}

			vArrSeatN = append(vArrSeatN, vSeatN)
		}

		vSeatNull.Seats = vArrSeatN
		vArrSeatNull = append(vArrSeatNull, vSeatNull)
	}

	vReturn.SeatNulls = vArrSeatNull
	vRS.Return = vReturn

	result = &vRS

	return result, nil
}

// MakeInternalGetSeatNullPerSubClassRS is a function to mapping from GetSeatNullPerSubClassRS to InternalGetSeatNullPerSubClassRS ("information.get_seat_null_per_subclass")
func MakeInternalGetSeatNullPerSubClassRS(input GetSeatNullPerSubClassRS) (result *InternalGetSeatNullPerSubClassRS, err error) {
	var vRS InternalGetSeatNullPerSubClassRS

	vRS.ErrCode = input.ErrCode
	vRS.ErrMsg = input.ErrMsg

	var vReturn GetSeatNullPerSubClass

	vReturn.Origin = input.Origin
	vReturn.Destination = input.Destination
	vReturn.TrainNo = input.TrainNo
	vReturn.DepartureDate = input.DepartureDate

	var vArrSeatNull []SeatNull

	for _, s := range input.SeatNulls {

		tripInfo := reflect.ValueOf(s)

		vSeatNull := SeatNull{
			WagonCode: tripInfo.Index(0).Interface().(string),
			WagonNo:   tripInfo.Index(1).Interface().(float64),
		}

		seat := tripInfo.Index(2).Interface().([]interface{})
		var vArrSeatN []SeatN

		for _, sc := range seat {

			scInfo := reflect.ValueOf(sc)

			vSeatN := SeatN{
				SeatRow:    scInfo.Index(0).Interface().(float64),
				SeatColumn: scInfo.Index(1).Interface().(string),
				SubClass:   scInfo.Index(2).Interface().(string),
				Status:     scInfo.Index(3).Interface().(float64),
			}

			vArrSeatN = append(vArrSeatN, vSeatN)
		}

		vSeatNull.Seats = vArrSeatN
		vArrSeatNull = append(vArrSeatNull, vSeatNull)
	}

	vReturn.SeatNulls = vArrSeatNull
	vRS.Return = vReturn

	result = &vRS

	return result, nil
}

// MakeInternalGetAgentBalanceRS is a function to mapping from GetAgentBalanceRS to InternalGetAgentBalanceRS ("information.get_agent_balance")
func MakeInternalGetAgentBalanceRS(input GetAgentBalanceRS) (result *InternalGetAgentBalanceRS, err error) {

	var vRS InternalGetAgentBalanceRS

	vRS.ErrCode = input.ErrCode
	vRS.ErrMsg = input.ErrMsg

	var vReturn GetAgentBalance

	vReturn.AgentCode = input.AgentCode
	vReturn.AgentName = input.AgentName
	vReturn.AgentBalance = input.AgentBalance

	vRS.Return = vReturn

	result = &vRS

	return result, nil
}

// MakeInternalGetBalanceRS is a function to mapping from GetBalanceRS to InternalGetAgentBalanceRS ("information.get_agent_balance")
func MakeInternalGetBalanceRS(input GetBalanceRS) (result *InternalGetBalanceRS, err error) {

	var vRS InternalGetBalanceRS

	vRS.ErrCode = input.ErrCode
	vRS.ErrMsg = input.ErrMsg

	var vReturn GetBalance

	vReturn.BookCode = input.BookCode
	vReturn.NumCode = fmt.Sprintf("%.f", input.NumCode)
	vReturn.NormalSales = input.NormalSales
	vReturn.ExtraFee = input.ExtraFee
	vReturn.BookBalance = input.BookBalance
	vReturn.Discount = input.Discount

	vRS.Return = vReturn

	result = &vRS

	return result, nil
}

// MakeInternalGetBookInfoRS is a function to mapping from GetBookInfoRS to InternalGetBookInfoRS ("information.get_book_info")
func MakeInternalGetBookInfoRS(input GetBookInfoRS) (result *InternalGetBookInfoRS, err error) {

	var vRS InternalGetBookInfoRS

	vRS.ErrCode = input.ErrCode
	vRS.ErrMsg = input.ErrMsg

	var vReturn GetBookInfo

	vReturn.BookCode = input.BookCode
	vReturn.NumCode = fmt.Sprintf("%.f", input.NumCode)
	vReturn.Caller = input.Caller
	vReturn.BookTime = input.BookTime
	vReturn.TrainNo = input.TrainNo
	vReturn.TrainName = input.TrainName
	vReturn.Origin = input.Origin
	vReturn.Destination = input.Destination
	vReturn.DepartureDate = input.DepartureDate
	vReturn.DepartureTime = input.DepartureTime
	vReturn.ArriveDate = input.ArriveDate
	vReturn.ArriveTime = input.ArriveTime
	vReturn.SubClass = input.SubClass
	vReturn.Class = input.Class
	vReturn.NormalSales = input.NormalSales
	vReturn.ExtraFee = input.ExtraFee
	vReturn.BookBalance = input.BookBalance
	vReturn.Discount = input.Discount
	// PaxList       []PaxList

	var vArrPaxList []PaxList

	for _, v := range input.PaxList {
		vPaxList := PaxList{
			PaxName:                   fmt.Sprintf("%s", v[0]),
			IdentityNo:                fmt.Sprintf("%s", v[1]),
			PaxType:                   fmt.Sprintf("%s", v[2]),
			TicketUnitNo:              fmt.Sprintf("%s", v[3]),
			TicketUnitPrintingCounter: v[4],
			ETicketNo:                 fmt.Sprintf("%s", v[5]),
			ETicketPrintingCounter:    v[6],
			Wagon:                     fmt.Sprintf("%s", v[7]),
			Seat:                      fmt.Sprintf("%s", v[8]),
		}
		vArrPaxList = append(vArrPaxList, vPaxList)
	}

	vReturn.PaxList = vArrPaxList
	vRS.Return = vReturn

	result = &vRS

	return result, nil
}

// MakeInternalGetBookPriceInfoRS is a function to mapping from GetBookPriceInfoRS to InternalGetBookPriceInfoRS ("information.get_book_price_info")
func MakeInternalGetBookPriceInfoRS(input GetBookPriceInfoRS) (result *InternalGetBookPriceInfoRS, err error) {

	var vRS InternalGetBookPriceInfoRS

	vRS.ErrCode = input.ErrCode
	vRS.ErrMsg = input.ErrMsg

	var vReturn GetBookPriceInfo

	vReturn.BookCode = input.BookCode
	vReturn.TotalPriceAdult = input.TotalPriceAdult
	vReturn.TotalPriceChild = input.TotalPriceChild
	vReturn.TotalPriceInfant = input.TotalPriceInfant
	vReturn.ExtraFee = input.ExtraFee
	vReturn.TotalPrice = input.TotalPrice

	vRS.Return = vReturn

	result = &vRS

	return result, nil
}

// MakeInternalBookingRS is a function to mapping from BookingRS to InternalBookingRS ("transaction.booking")
func MakeInternalBookingRS(input BookingRS) (result *InternalBookingRS, err error) {

	var vRS InternalBookingRS

	vRS.ErrCode = input.ErrCode
	vRS.ErrMsg = input.ErrMsg

	var vReturn Booking

	vReturn.Origin = input.Origin
	vReturn.Destination = input.Destination
	vReturn.DepartureDate = input.DepartureDate
	vReturn.TrainNo = input.TrainNo
	vReturn.BookCode = input.BookCode
	vReturn.NumCode = input.NumCode

	// PaxNums       PaxNum
	vReturn.PaxNums = PaxNum{}
	if input.PaxNums != nil {
		vReturn.PaxNums = PaxNum{
			AdultCount:  input.PaxNums[0],
			ChildCount:  input.PaxNums[1],
			InfantCount: input.PaxNums[2],
		}
	}

	// PaxNames      []PaxName
	var vArrPaxName []PaxName
	for _, v := range input.PaxNames {
		vArrPaxName = append(vArrPaxName, PaxName{Name: v})
	}
	vReturn.PaxNames = vArrPaxName

	// Seats         []PaxSeat
	var vArrPaxSeat []PaxSeat
	for _, v := range input.Seats {

		vArrPaxSeat = append(vArrPaxSeat,
			PaxSeat{
				WagonCode: fmt.Sprintf("%s", v[0]),
				WagonNo:   fmt.Sprintf("%s", v[1]),
				SeatRow:   fmt.Sprintf("%s", v[2]),
				SeatCol:   fmt.Sprintf("%s", v[3]),
			})
	}
	vReturn.Seats = vArrPaxSeat

	vReturn.NormalSales = input.NormalSales
	vReturn.ExtraFee = input.ExtraFee
	vReturn.BookBalance = input.BookBalance
	vReturn.Discount = input.Discount

	vRS.Return = vReturn

	result = &vRS

	return result, nil
}
