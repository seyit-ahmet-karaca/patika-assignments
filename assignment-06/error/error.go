package error

import (
	"encoding/json"
)

func ParseError(e error) []byte {

	switch e.(type) {
	case *InvalidBalance:
		err := e.(*InvalidBalance)
		errMessage, errJson := json.Marshal(err)
		if errJson != nil {
			return []byte(errJson.Error())
		}
		return errMessage
	case *InvalidData:
		err := e.(*InvalidData)
		errMessage, errJson := json.Marshal(err)
		if errJson != nil {
			return []byte(errJson.Error())
		}
		return errMessage
	case *UsernameNotFound:
		err := e.(*UsernameNotFound)
		errMessage, errJson := json.Marshal(err)
		if errJson != nil {
			return []byte(errJson.Error())
		}
		return errMessage
	case *WalletAlreadyExists:
		err := e.(*WalletAlreadyExists)
		errMessage, errJson := json.Marshal(err)
		if errJson != nil {
			return []byte(errJson.Error())
		}
		return errMessage
	}
	return []byte(e.Error())
}
