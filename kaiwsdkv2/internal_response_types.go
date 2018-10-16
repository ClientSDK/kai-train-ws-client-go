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
	ErrCode interface{}   `json:"err_code"`
	ErrMsg  interface{}   `json:"err_msg"`
	Return  []Origination `json:"return"`
}

// Origination type
type Origination struct {
	OriginCode string `json:"OriginCode"`
	OriginName string `json:"OriginName"`
}

// InternalGetPayTypeRS represent "data.get_pay_type" internal response
type InternalGetPayTypeRS struct {
	ErrCode interface{} `json:"err_code"`
	ErrMsg  interface{} `json:"err_msg"`
	Return  []PayType   `json:"return"`
}

// PayType type
type PayType struct {
	Name string `json:"name"`
}
