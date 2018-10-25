// Copyright 2018 The ClientSDK Team Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0-style
// license that can be found in the LICENSE file.

// Author: ClientSDK Team (muharihar)

package kaiwsdkv2

// InternalGetOriginationRS represent "data.get_org" internal response
type InternalGetOriginationRS struct {
	ErrCode interface{}   `json:"errCode"`
	ErrMsg  interface{}   `json:"errMsg"`
	Return  []Origination `json:"return"`
}

// Origination type
type Origination struct {
	OriginCode string `json:"OriginCode"`
	OriginName string `json:"OriginName"`
}

// --

// InternalGetDestinationRS represent "data.get_des" internal response
type InternalGetDestinationRS struct {
	ErrCode interface{}   `json:"errCode"`
	ErrMsg  interface{}   `json:"errMsg"`
	Return  []Destination `json:"return"`
}

// Destination type
type Destination struct {
	DestCode string `json:"DestCode"`
	DestName string `json:"DestName"`
}

// --

// InternalGetPayTypeRS represent "data.get_pay_type" internal response
type InternalGetPayTypeRS struct {
	ErrCode interface{} `json:"errCode"`
	ErrMsg  interface{} `json:"errMsg"`
	Return  []PayType   `json:"return"`
}

// PayType type
type PayType struct {
	Name string `json:"name"`
}

// --

// InternalGetAgentBalanceRS represent "information.get_agent_balance" internal response
type InternalGetAgentBalanceRS struct {
	ErrCode interface{}     `json:"errCode"`
	ErrMsg  interface{}     `json:"errMsg"`
	Return  GetAgentBalance `json:"return"`
}

// GetAgentBalance type
type GetAgentBalance struct {
	AgentCode    string  `json:"agentCode"`
	AgentName    string  `json:"agentName"`
	AgentBalance float64 `json:"agentBalance"`
}

// --

// InternalGetBalanceRS represent "information.get_balance" internal response
type InternalGetBalanceRS struct {
	ErrCode interface{} `json:"errCode"`
	ErrMsg  interface{} `json:"errMsg"`
	Return  GetBalance  `json:"return"`
}

// GetBalance type
type GetBalance struct {
	BookCode    string  `json:"bookCode"`
	NumCode     string  `json:"numCode"`
	NormalSales float64 `json:"normalSales"`
	ExtraFee    float64 `json:"extraFee"`
	BookBalance float64 `json:"bookBalance"`
	Discount    float64 `json:"discount"`
}

// --

// InternalGetBookInfoRS represent "information.get_book_info" internal response
type InternalGetBookInfoRS struct {
	ErrCode interface{} `json:"errCode"`
	ErrMsg  interface{} `json:"errMsg"`
	Return  GetBookInfo `json:"return"`
}

// GetBookInfo type
type GetBookInfo struct {
	BookCode      string    `json:"bookCode"`
	NumCode       string    `json:"numCode"`
	Caller        string    `json:"caller"`
	BookTime      string    `json:"bookTime"`
	TrainNo       string    `json:"trainNo"`
	TrainName     string    `json:"trainName"`
	Origin        string    `json:"origin"`
	Destination   string    `json:"destination"`
	DepartureDate string    `json:"departureDate"`
	DepartureTime string    `json:"departureTime"`
	ArriveDate    string    `json:"arriveDate"`
	ArriveTime    string    `json:"arriveTime"`
	SubClass      string    `json:"subClass"`
	Class         string    `json:"class"`
	NormalSales   float64   `json:"normalSales"`
	ExtraFee      float64   `json:"extraFee"`
	BookBalance   float64   `json:"bookBalance"`
	Discount      float64   `json:"discount"`
	PaxList       []PaxList `json:"paxList"`
}

// PaxList type
type PaxList struct {
	PaxName                   string      `json:"paxName"`
	IdentityNo                string      `json:"IdentityNo"`
	PaxType                   string      `json:"paxType"`
	TicketUnitNo              string      `json:"ticketUnitNo"`
	TicketUnitPrintingCounter interface{} `json:"ticketUnitPrintingCounter"`
	ETicketNo                 string      `json:"eTicketNo"`
	ETicketPrintingCounter    interface{} `json:"eTicketPrintingCounter"`
	Wagon                     string      `json:"wagon"`
	Seat                      string      `json:"seat"`
}

// --

// InternalGetBookPriceInfoRS represent "information.get_book_price_info" internal response
type InternalGetBookPriceInfoRS struct {
	ErrCode interface{}      `json:"errCode"`
	ErrMsg  interface{}      `json:"errMsg"`
	Return  GetBookPriceInfo `json:"return"`
}

// GetBookPriceInfo type
type GetBookPriceInfo struct {
	BookCode         string  `json:"bookCode"`
	TotalPriceAdult  float64 `json:"totalPriceAdult"`
	TotalPriceChild  float64 `json:"totalPriceChild"`
	TotalPriceInfant float64 `json:"totalPriceInfant"`
	ExtraFee         float64 `json:"extraFee"`
	TotalPrice       float64 `json:"totalPrice"`
}

// --

// InternalGetScheduleLiteRS represent "information.get_schedule" internal response
type InternalGetScheduleLiteRS struct {
	ErrCode interface{}     `json:"errCode"`
	ErrMsg  interface{}     `json:"errMsg"`
	Return  GetScheduleLite `json:"return"`
}

// GetScheduleLite type
type GetScheduleLite struct {
	Origin        string         `json:"origin"`
	Destination   string         `json:"destination"`
	DepartureDate string         `json:"departureDate"`
	Schedules     []ScheduleLite `json:"schedules"`
}

// ScheduleLite type
type ScheduleLite struct {
	TrainNo       string              `json:"trainNo"`
	TrainName     string              `json:"trainName"`
	DepartureDate string              `json:"departureDate"`
	ArriveDate    string              `json:"arriveDate"`
	DepartureTime string              `json:"departureTime"`
	ArriveTime    string              `json:"arriveTime"`
	AvailSubClass []AvailSubClassLite `json:"availSubClass"`
}

// AvailSubClassLite type
type AvailSubClassLite struct {
	SubClass      string  `json:"subClass"`
	SeatAvailable float64 `json:"seatAvailable"`
	SeatClass     string  `json:"seatClass"`
	AdultPrice    float64 `json:"adultPrice"`
	ChildPrice    float64 `json:"childPrice"`
	InfantPrice   float64 `json:"infantPrice"`
}

// --

// InternalGetScheduleV2RS represent "information.get_schedule_v2" internal response
type InternalGetScheduleV2RS struct {
	ErrCode interface{}   `json:"errCode"`
	ErrMsg  interface{}   `json:"errMsg"`
	Return  GetScheduleV2 `json:"return"`
}

// GetScheduleV2 type
type GetScheduleV2 struct {
	Origin        string       `json:"origin"`
	Destination   string       `json:"destination"`
	DepartureDate string       `json:"departureDate"`
	Schedules     []ScheduleV2 `json:"schedules"`
}

// ScheduleV2 type
type ScheduleV2 struct {
	TrainNo       string            `json:"trainNo"`
	TrainName     string            `json:"trainName"`
	DepartureDate string            `json:"departureDate"`
	ArriveDate    string            `json:"arriveDate"`
	DepartureTime string            `json:"departureTime"`
	ArriveTime    string            `json:"arriveTime"`
	AvailSubClass []AvailSubClassV2 `json:"availSubClass"`
}

// AvailSubClassV2 type
type AvailSubClassV2 struct {
	SubClass      string  `json:"subClass"`
	SeatAvailable float64 `json:"seatAvailable"`
	SeatClass     string  `json:"seatClass"`
	AdultPrice    float64 `json:"adultPrice"`
	ChildPrice    float64 `json:"childPrice"`
	InfantPrice   float64 `json:"infantPrice"`
}

// --

// InternalGetScheduleRS represent "information.get_schedule" internal response
type InternalGetScheduleRS struct {
	ErrCode interface{} `json:"errCode"`
	ErrMsg  interface{} `json:"errMsg"`
	Return  GetSchedule `json:"return"`
}

// GetSchedule type
type GetSchedule struct {
	Origin        string     `json:"origin"`
	Destination   string     `json:"destination"`
	DepartureDate string     `json:"departureDate"`
	Schedules     []Schedule `json:"schedules"`
}

// Schedule type
type Schedule struct {
	TrainNo       string          `json:"trainNo"`
	TrainName     string          `json:"trainName"`
	DepartureTime string          `json:"departureTime"`
	ArriveTime    string          `json:"arriveTime"`
	AvailSubClass []AvailSubClass `json:"availSubClass"`
}

// AvailSubClass type
type AvailSubClass struct {
	SubClass      string  `json:"subClass"`
	SeatAvailable float64 `json:"seatAvailable"`
	SeatClass     string  `json:"seatClass"`
	AdultPrice    float64 `json:"adultPrice"`
	ChildPrice    float64 `json:"childPrice"`
	InfantPrice   float64 `json:"infantPrice"`
}

// --

// InternalGetSeatMapRS represent "information.get_seat_map" internal response
type InternalGetSeatMapRS struct {
	ErrCode interface{} `json:"errCode"`
	ErrMsg  interface{} `json:"errMsg"`
	Return  GetSeatMap  `json:"return"`
}

// GetSeatMap type
type GetSeatMap struct {
	Origin        string    `json:"origin"`
	Destination   string    `json:"destination"`
	TrainNo       string    `json:"trainNo"`
	DepartureDate string    `json:"departureDate"`
	SeatMaps      []SeatMap `json:"seatMaps"`
}

// SeatMap type
type SeatMap struct {
	WagonCode string  `json:"wagonCode"`
	WagonNo   float64 `json:"wagonNo"`
	Seats     []Seat  `json:"seats"`
}

// Seat type
type Seat struct {
	Row        float64 `json:"row"`
	Column     float64 `json:"column"`
	SeatRow    float64 `json:"seatRow"`
	SeatColumn string  `json:"seatColumn"`
	SubClass   string  `json:"seatClass"`
	Status     float64 `json:"status"`
}

// --

// InternalGetSeatMapPerSubClassRS represent "information.get_seat_map_per_subclass" internal response
type InternalGetSeatMapPerSubClassRS struct {
	ErrCode interface{}           `json:"errCode"`
	ErrMsg  interface{}           `json:"errMsg"`
	Return  GetSeatMapPerSubClass `json:"return"`
}

// GetSeatMapPerSubClass type
type GetSeatMapPerSubClass struct {
	Origin        string    `json:"origin"`
	Destination   string    `json:"destination"`
	TrainNo       string    `json:"trainNo"`
	DepartureDate string    `json:"departureDate"`
	SeatMaps      []SeatMap `json:"seatMaps"`
}

// --

// InternalGetSeatNullRS represent "information.get_seat_null" internal response
type InternalGetSeatNullRS struct {
	ErrCode interface{} `json:"errCode"`
	ErrMsg  interface{} `json:"errMsg"`
	Return  GetSeatNull `json:"return"`
}

// GetSeatNull type
type GetSeatNull struct {
	Origin        string     `json:"origin"`
	Destination   string     `json:"destination"`
	TrainNo       string     `json:"trainNo"`
	DepartureDate string     `json:"departureDate"`
	SeatNulls     []SeatNull `json:"seatNulls"`
}

// SeatNull type
type SeatNull struct {
	WagonCode string  `json:"wagonCode"`
	WagonNo   float64 `json:"wagonNo"`
	Seats     []SeatN `json:"seats"`
}

// SeatN type
type SeatN struct {
	SeatRow    float64 `json:"seatRow"`
	SeatColumn string  `json:"seatColumn"`
	SubClass   string  `json:"seatClass"`
	Status     float64 `json:"status"`
}

// --

// InternalGetSeatNullPerSubClassRS represent "information.get_set_null_per_subclass" internal response
type InternalGetSeatNullPerSubClassRS struct {
	ErrCode interface{}            `json:"errCode"`
	ErrMsg  interface{}            `json:"errMsg"`
	Return  GetSeatNullPerSubClass `json:"return"`
}

// GetSeatNullPerSubClass type
type GetSeatNullPerSubClass struct {
	Origin        string     `json:"origin"`
	Destination   string     `json:"destination"`
	TrainNo       string     `json:"trainNo"`
	DepartureDate string     `json:"departureDate"`
	SeatNulls     []SeatNull `json:"seatNulls"`
}

// --

// InternalBookingRS represent "transaction.booking" internal response
type InternalBookingRS struct {
	ErrCode interface{} `json:"errCode"`
	ErrMsg  interface{} `json:"errMsg"`
	Return  Booking     `json:"return"`
}

// Booking type
type Booking struct {
	Origin        string      `json:"origin"`
	Destination   string      `json:"destination"`
	DepartureDate string      `json:"departureDate"`
	TrainNo       interface{} `json:"trainNo"`
	BookCode      string      `json:"bookCode"`
	NumCode       interface{} `json:"numCode"`
	PaxNums       PaxNum      `json:"paxNums"`
	PaxNames      []PaxName   `json:"paxNames"`
	Seats         []PaxSeat   `json:"seats"`
	NormalSales   float64     `json:"normalSales"`
	ExtraFee      float64     `json:"extraFee"`
	BookBalance   float64     `json:"bookBalance"`
	Discount      float64     `json:"discount"`
}

// PaxNum type
type PaxNum struct {
	AdultCount  int `json:"adultCount"`
	ChildCount  int `json:"childCount"`
	InfantCount int `json:"infantCount"`
}

// PaxName type
type PaxName struct {
	Name string `json:"name"`
}

// PaxSeat type
type PaxSeat struct {
	WagonCode string `json:"wagonCode"`
	WagonNo   string `json:"wagonNo"`
	SeatRow   string `json:"seatRow"`
	SeatCol   string `json:"seatCol"`
}

// --

// InternalBookingWithArvInfoRS represent "transaction.booking_with_arv_info" internal response
type InternalBookingWithArvInfoRS struct {
	ErrCode interface{}        `json:"errCode"`
	ErrMsg  interface{}        `json:"errMsg"`
	Return  BookingWithArvInfo `json:"return"`
}

// BookingWithArvInfo type
type BookingWithArvInfo struct {
	Origin        string      `json:"origin"`
	Destination   string      `json:"destination"`
	DepartureDate string      `json:"departureDate"`
	ArriveDate    string      `json:"arriveDate"`
	DepartureTime string      `json:"departureTime"`
	ArriveTime    string      `json:"arriveTime"`
	TrainNo       interface{} `json:"trainNo"`
	BookCode      string      `json:"bookCode"`
	NumCode       interface{} `json:"numCode"`
	PaxNums       PaxNum      `json:"paxNums"`
	PaxNames      []PaxName   `json:"paxNames"`
	Seats         []PaxSeat   `json:"seats"`
	NormalSales   float64     `json:"normalSales"`
	ExtraFee      float64     `json:"extraFee"`
	BookBalance   float64     `json:"bookBalance"`
	Discount      float64     `json:"discount"`
}

// --

// InternalCancelBookRS represent "transaction.cancel_book" internal response
type InternalCancelBookRS struct {
	ErrCode interface{} `json:"errCode"`
	ErrMsg  interface{} `json:"errMsg"`
	Return  CancelBook  `json:"return"`
}

// CancelBook type
type CancelBook struct {
	BookCode     string  `json:"bookCode"`
	Status       string  `json:"status"`
	RefundAmount float64 `json:"refundAmount"`
}

// --

// InternalManualSeatRS represent "transaction.manual_seat" internal response
type InternalManualSeatRS struct {
	ErrCode interface{} `json:"errCode"`
	ErrMsg  interface{} `json:"errMsg"`
	Return  ManualSeat  `json:"return"`
}

// ManualSeat type
type ManualSeat struct {
	BookCode  string  `json:"bookCode"`
	WagonCode string  `json:"wagonCode"`
	WagonNo   int64   `json:"wagonNo"`
	Seats     []SeatM `json:"seats"`
}

// SeatM type
type SeatM struct {
	SeatNo string `json:"seatNo"`
}

// --

// InternalUpdatePaxRS represent "information.update_pax" internal response
type InternalUpdatePaxRS struct {
	ErrCode interface{} `json:"errCode"`
	ErrMsg  interface{} `json:"errMsg"`
	Return  UpdatePax   `json:"return"`
}

// UpdatePax type
type UpdatePax struct {
	BookCode string    `json:"bookCode"`
	PaxNums  PaxNum    `json:"paxNums"`
	PaxNames []PaxName `json:"paxNames"`
}

// --

// InternalPaymentRS represent "transaction.payment" internal response
type InternalPaymentRS struct {
	ErrCode interface{} `json:"errCode"`
	ErrMsg  interface{} `json:"errMsg"`
	Return  Payment     `json:"return"`
}

// Payment type
type Payment struct {
	BookCode    string  `json:"bookCode"`
	BookBalance float64 `json:"bookBalance"`
}

// --

// InternalCancelPaymentRS represent "transaction.cancel_payment" internal response
type InternalCancelPaymentRS struct {
	ErrCode interface{}   `json:"errCode"`
	ErrMsg  interface{}   `json:"errMsg"`
	Return  CancelPayment `json:"return"`
}

// CancelPayment type
type CancelPayment struct {
	BookCode    string  `json:"bookBode"`
	NormalSales float64 `json:"normalSales"`
	ExtraFee    float64 `json:"extraFee"`
	BookBalance float64 `json:"bookBalance"`
}
