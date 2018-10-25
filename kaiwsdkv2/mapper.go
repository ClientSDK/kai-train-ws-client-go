package kaiwsdkv2

import "fmt"

// MakeInternalGetOriginationRS is a function to mapping from GetOriginationRS to InternalGetOriginationRS ("information.get_org")
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

// MakeInternalGetDestinationRS is a function to mapping from GetDestinationRS to InternalGetDestinationRS ("information.get_des")
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

// MakeInternalGetPayTypeRS is a function to mapping from GetPayTypeRS to InternalGetDestinationRS ("information.get_des")
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
