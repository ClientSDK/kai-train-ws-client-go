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
