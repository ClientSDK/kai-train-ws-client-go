// Copyright 2018 The ClientSDK Team Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0-style
// license that can be found in the LICENSE file.

// Author: ClientSDK Team (muharihar)

package kaiwsdkv2

// GetDestinationRS represent "data.get_des" response
type GetDestinationRS struct {
	ErrCode      interface{}     `json:"err_code"`
	ErrMsg       interface{}     `json:"err_msg"`
	Destinations [][]interface{} `json:"destination"`
}

// GetOriginationRS represent "data.get_org" response
type GetOriginationRS struct {
	ErrCode      interface{}     `json:"err_code"`
	ErrMsg       interface{}     `json:"err_msg"`
	Originations [][]interface{} `json:"origination"`
}

// GetPayTypeRS represent "data.get_pay_type" response
type GetPayTypeRS struct {
	ErrCode  interface{}   `json:"err_code"`
	ErrMsg   interface{}   `json:"err_msg"`
	PayTypes []interface{} `json:"pay_type"`
}

// GetAgentBalanceRS represent "information.get_agent_balance" response
type GetAgentBalanceRS struct {
	ErrCode      interface{} `json:"err_code"`
	ErrMsg       interface{} `json:"err_msg"`
	AgentCode    string      `json:"agent_code"`
	AgentName    string      `json:"agent_name"`
	AgentBalance float64     `json:"agent_balance"`
}

// GetBalanceRS represent "information.get_balance" response
type GetBalanceRS struct {
	ErrCode     interface{} `json:"err_code"`
	ErrMsg      interface{} `json:"err_msg"`
	BookCode    string      `json:"book_code"`
	NumCode     interface{} `json:"num_code"`
	NormalSales float64     `json:"normal_sales"`
	ExtraFee    float64     `json:"extra_fee"`
	BookBalance float64     `json:"book_balance"`
	Discount    float64     `json:"discount"`
}

// GetBookInfoRS represent "information.get_book_info" response
type GetBookInfoRS struct {
	ErrCode       interface{}     `json:"err_code"`
	ErrMsg        interface{}     `json:"err_msg"`
	BookCode      string          `json:"book_code"`
	NumCode       interface{}     `json:"num_code"`
	Caller        string          `json:"caller"`
	BookTime      string          `json:"book_time"`
	TrainNo       string          `json:"train_no"`
	TrainName     string          `json:"train_name"`
	Origin        string          `json:"org"`
	Destination   string          `json:"des"`
	DepartureDate string          `json:"dep_date"`
	DepartureTime string          `json:"dep_time"`
	ArriveDate    string          `json:"arv_date"`
	ArriveTime    string          `json:"arv_time"`
	SubClass      string          `json:"subclass"`
	Class         string          `json:"class"`
	NormalSales   float64         `json:"normal_sales"`
	ExtraFee      float64         `json:"extra_fee"`
	BookBalance   float64         `json:"book_balance"`
	Discount      float64         `json:"discount"`
	PaxList       [][]interface{} `json:"pax_list"`
}

// GetBookPriceInfoRS represent "information.get_book_price_info" response
type GetBookPriceInfoRS struct {
	ErrCode          interface{} `json:"err_code"`
	ErrMsg           interface{} `json:"err_msg"`
	BookCode         string      `json:"book_code"`
	TotalPriceAdult  float64     `json:"total_price_adult"`
	TotalPriceChild  float64     `json:"total_price_child"`
	TotalPriceInfant float64     `json:"total_price_infant"`
	ExtraFee         float64     `json:"extra_fee"`
	TotalPrice       float64     `json:"total_price"`
}

// GetScheduleLiteRS represent "information.get_schedule_lite" method
type GetScheduleLiteRS struct {
	ErrCode       interface{}   `json:"err_code"`
	ErrMsg        interface{}   `json:"err_msg"`
	Origin        string        `json:"org"`
	Destination   string        `json:"des"`
	DepartureDate string        `json:"dep_date"`
	Schedules     []interface{} `json:"schedule"`
}

// GetScheduleRS represent "information.get_schedule" response
type GetScheduleRS struct {
	ErrCode       interface{}   `json:"err_code"`
	ErrMsg        interface{}   `json:"err_msg"`
	Origin        string        `json:"org"`
	Destination   string        `json:"des"`
	DepartureDate string        `json:"dep_date"`
	Schedules     []interface{} `json:"schedule"`
}

// GetScheduleV2RS represent "information.get_schedule_v2" response
type GetScheduleV2RS struct {
	ErrCode       interface{}   `json:"err_code"`
	ErrMsg        interface{}   `json:"err_msg"`
	Origin        string        `json:"org"`
	Destination   string        `json:"des"`
	DepartureDate string        `json:"dep_date"`
	Schedules     []interface{} `json:"schedule"`
}

// GetSeatMapRS represent "information.get_seat_map" response
type GetSeatMapRS struct {
	ErrCode       interface{}   `json:"err_code"`
	ErrMsg        interface{}   `json:"err_msg"`
	Origin        string        `json:"org"`
	Destination   string        `json:"des"`
	TrainNo       string        `json:"train_no"`
	DepartureDate string        `json:"dep_date"`
	SeatMaps      []interface{} `json:"seat_map"`
}

// GetSeatMapPerSubClassRS represent "information.get_seat_map_per_subclass" response
type GetSeatMapPerSubClassRS struct {
	ErrCode       interface{}   `json:"err_code"`
	ErrMsg        interface{}   `json:"err_msg"`
	Origin        string        `json:"org"`
	Destination   string        `json:"des"`
	TrainNo       string        `json:"train_no"`
	DepartureDate string        `json:"dep_date"`
	SeatMaps      []interface{} `json:"seat_map"`
}

// GetSeatNullRS represent "information.get_seat_null" response
type GetSeatNullRS struct {
	ErrCode       interface{}   `json:"err_code"`
	ErrMsg        interface{}   `json:"err_msg"`
	Origin        string        `json:"org"`
	Destination   string        `json:"des"`
	TrainNo       string        `json:"train_no"`
	DepartureDate string        `json:"dep_date"`
	SeatNulls     []interface{} `json:"seat_null"`
}

// GetSeatNullPerSubClassRS represent "information.get_seat_null_per_subclass" response
type GetSeatNullPerSubClassRS struct {
	ErrCode       interface{}   `json:"err_code"`
	ErrMsg        interface{}   `json:"err_msg"`
	Origin        string        `json:"org"`
	Destination   string        `json:"des"`
	TrainNo       string        `json:"train_no"`
	DepartureDate string        `json:"dep_date"`
	SeatNulls     []interface{} `json:"seat_null"`
}

// BookingRS represent "transaction.booking" response
type BookingRS struct {
	ErrCode       interface{}     `json:"err_code"`
	ErrMsg        interface{}     `json:"err_msg"`
	Origin        string          `json:"org"`
	Destination   string          `json:"des"`
	DepartureDate string          `json:"dep_date"`
	TrainNo       interface{}     `json:"train_no"`
	BookCode      string          `json:"book_code"`
	NumCode       interface{}     `json:"num_code"`
	PaxNums       []interface{}   `json:"pax_num"`
	PaxNames      []interface{}   `json:"pax_name"`
	Seats         [][]interface{} `json:"seat"`
	NormalSales   float64         `json:"normal_sales"`
	ExtraFee      float64         `json:"extra_fee"`
	BookBalance   float64         `json:"book_balance"`
	Discount      float64         `json:"discount"`
}

// BookingWithArvInfoRS represent "transaction.booking_with_arv_info" response
type BookingWithArvInfoRS struct {
	ErrCode       interface{}     `json:"err_code"`
	ErrMsg        interface{}     `json:"err_msg"`
	Origin        string          `json:"org"`
	Destination   string          `json:"des"`
	DepartureDate string          `json:"dep_date"`
	ArriveDate    string          `json:"arv_date"`
	DepartureTime string          `json:"dep_time"`
	ArriveTime    string          `json:"arv_time"`
	TrainNo       interface{}     `json:"train_no"`
	BookCode      string          `json:"book_code"`
	NumCode       interface{}     `json:"num_code"`
	PaxNums       []interface{}   `json:"pax_num"`
	PaxNames      []interface{}   `json:"pax_name"`
	Seats         [][]interface{} `json:"seat"`
	NormalSales   float64         `json:"normal_sales"`
	ExtraFee      float64         `json:"extra_fee"`
	BookBalance   float64         `json:"book_balance"`
	Discount      float64         `json:"discount"`
}

// CancelBookRS represent "transaction.cancel_book" response
type CancelBookRS struct {
	ErrCode      interface{} `json:"err_code"`
	ErrMsg       interface{} `json:"err_msg"`
	BookCode     string      `json:"book_code"`
	Status       string      `json:"status"`
	RefundAmount float64     `json:"refund_amount"`
}

// ManualSeatRS represent "transaction.manual_seat" response
type ManualSeatRS struct {
	ErrCode   interface{}   `json:"err_code"`
	ErrMsg    interface{}   `json:"err_msg"`
	BookCode  string        `json:"book_code"`
	WagonCode string        `json:"wagon_code"`
	WagonNo   int64         `json:"wagon_no"`
	Seats     []interface{} `json:"seat"`
}

// UpdatePaxRS represent "transaction.update_pax" response
type UpdatePaxRS struct {
	ErrCode  interface{}   `json:"err_code"`
	ErrMsg   interface{}   `json:"err_msg"`
	BookCode string        `json:"book_code"`
	PaxNums  []interface{} `json:"pax_num"`
	PaxNames []interface{} `json:"pax_name"`
}

// PaymentRS represent "transaction.payment" response
type PaymentRS struct {
	ErrCode     interface{} `json:"err_code"`
	ErrMsg      interface{} `json:"err_msg"`
	BookCode    string      `json:"book_code"`
	BookBalance float64     `json:"book_balance"`
}

// CancelPaymentRS represent "transaction.cancel_payment" response
type CancelPaymentRS struct {
	ErrCode     interface{} `json:"err_code"`
	ErrMsg      interface{} `json:"err_msg"`
	BookCode    string      `json:"book_code"`
	NormalSales float64     `json:"normal_sales"`
	ExtraFee    float64     `json:"extra_fee"`
	BookBalance float64     `json:"book_balance"`
}
