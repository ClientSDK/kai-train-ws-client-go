package kaiwsdkv2

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

// InternalGetBalanceRS represent "information.get_balance" internal response
type InternalGetBalanceRS struct {
	ErrCode interface{} `json:"errCode"`
	ErrMsg  interface{} `json:"errMsg"`
	Return  GetBalance  `json:"return"`
}

// GetBalance type
type GetBalance struct {
	BookCode    string      `json:"bookCode"`
	NumCode     interface{} `json:"numCode"`
	NormalSales float64     `json:"normalSales"`
	ExtraFee    float64     `json:"extraFee"`
	BookBalance float64     `json:"bookBalance"`
	Discount    float64     `json:"discount"`
}

// InternalGetBookInfoRS represent "information.get_book_info" internal response
type InternalGetBookInfoRS struct {
	ErrCode interface{} `json:"errCode"`
	ErrMsg  interface{} `json:"errMsg"`
	Return  GetBookInfo `json:"return"`
}

// GetBookInfo type
type GetBookInfo struct {
	BookCode      string      `json:"bookCode"`
	NumCode       interface{} `json:"numCode"`
	Caller        string      `json:"caller"`
	BookTime      string      `json:"bookTime"`
	TrainNo       string      `json:"trainNo"`
	TrainName     string      `json:"trainName"`
	Origin        string      `json:"origin"`
	Destination   string      `json:"destination"`
	DepartureDate string      `json:"departureDate"`
	DepartureTime string      `json:"departureTime"`
	ArriveDate    string      `json:"arriveDate"`
	ArriveTime    string      `json:"arriveTime"`
	Subclass      string      `json:"subclass"`
	Class         string      `json:"class"`
	NormalSales   float64     `json:"normalSales"`
	ExtraFee      float64     `json:"extraFee"`
	BookBalance   float64     `json:"bookBalance"`
	Discount      float64     `json:"discount"`
	PaxList       []PaxList   `json:"paxList"`
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
	ScheduleLites []ScheduleLite `json:"schedule"`
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
