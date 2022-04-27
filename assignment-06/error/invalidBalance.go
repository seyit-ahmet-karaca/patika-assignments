package error

import (
	"fmt"
	"karaca/config"
)

type InvalidBalance struct {
	Message string `json:"errorMessage"`
}

func (i *InvalidBalance) Error() string {
	return i.Message
}

func GetInvalidBalance() *InvalidBalance {
	return &InvalidBalance{Message: fmt.Sprintf("balance can not less then %d", config.Get().MinimumBalanceAmount)}
}
